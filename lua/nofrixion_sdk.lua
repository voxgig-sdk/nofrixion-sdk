-- Nofrixion SDK

local vs = require("utility.struct.struct")
local Utility = require("core.utility_type")
local Spec = require("core.spec")
local helpers = require("core.helpers")

-- Load utility registration (populates Utility._registrar)
require("utility.register")

-- Load features
local BaseFeature = require("feature.base_feature")
local features_factory = require("features")


local NofrixionSDK = {}
NofrixionSDK.__index = NofrixionSDK


local function _make_feature(name)
  local factory = features_factory[name]
  if factory ~= nil then
    return factory()
  end
  return features_factory.base()
end

NofrixionSDK._make_feature = _make_feature


function NofrixionSDK.new(options)
  local self = setmetatable({}, NofrixionSDK)
  self.mode = "live"
  self.features = {}
  self.options = nil

  local utility = Utility.new()
  self._utility = utility

  local config = require("config")()

  self._rootctx = utility.make_context({
    client = self,
    utility = utility,
    config = config,
    options = options or {},
    shared = {},
  }, nil)

  self.options = utility.make_options(self._rootctx)

  if vs.getpath(self.options, "feature.test.active") == true then
    self.mode = "test"
  end

  self._rootctx.options = self.options

  -- Add features from config.
  local feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
  if feature_opts ~= nil then
    local feature_items = vs.items(feature_opts)
    if feature_items ~= nil then
      for _, item in ipairs(feature_items) do
        local fname = item[1]
        local fopts = helpers.to_map(item[2])
        if fopts ~= nil and fopts["active"] == true then
          utility.feature_add(self._rootctx, _make_feature(fname))
        end
      end
    end
  end

  -- Add extension features.
  local extend = vs.getprop(self.options, "extend")
  if type(extend) == "table" then
    for _, f in ipairs(extend) do
      if type(f) == "table" and type(f.get_name) == "function" then
        utility.feature_add(self._rootctx, f)
      end
    end
  end

  -- Initialize features.
  for _, f in ipairs(self.features) do
    utility.feature_init(self._rootctx, f)
  end

  utility.feature_hook(self._rootctx, "PostConstruct")

  -- #BuildFeatures

  return self
end


function NofrixionSDK:options_map()
  local out = vs.clone(self.options)
  if type(out) == "table" then
    return out
  end
  return {}
end


function NofrixionSDK:get_utility()
  return Utility.copy(self._utility)
end


function NofrixionSDK:get_root_ctx()
  return self._rootctx
end


function NofrixionSDK:prepare(fetchargs)
  local utility = self._utility

  fetchargs = fetchargs or {}

  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "prepare",
    ctrl = ctrl,
  }, self._rootctx)

  local options = self.options

  local path = vs.getprop(fetchargs, "path") or ""
  if type(path) ~= "string" then path = "" end

  local method = vs.getprop(fetchargs, "method") or "GET"
  if type(method) ~= "string" then method = "GET" end

  local params = helpers.to_map(vs.getprop(fetchargs, "params")) or {}
  local query = helpers.to_map(vs.getprop(fetchargs, "query")) or {}

  local headers = utility.prepare_headers(ctx)

  local base = vs.getprop(options, "base") or ""
  if type(base) ~= "string" then base = "" end
  local prefix = vs.getprop(options, "prefix") or ""
  if type(prefix) ~= "string" then prefix = "" end
  local suffix = vs.getprop(options, "suffix") or ""
  if type(suffix) ~= "string" then suffix = "" end

  ctx.spec = Spec.new({
    base = base,
    prefix = prefix,
    suffix = suffix,
    path = path,
    method = method,
    params = params,
    query = query,
    headers = headers,
    body = vs.getprop(fetchargs, "body"),
    step = "start",
  })

  -- Merge user-provided headers.
  local uh = vs.getprop(fetchargs, "headers")
  if type(uh) == "table" then
    for k, v in pairs(uh) do
      ctx.spec.headers[k] = v
    end
  end

  local _, err = utility.prepare_auth(ctx)
  if err ~= nil then
    return nil, err
  end

  return utility.make_fetch_def(ctx)
end


function NofrixionSDK:direct(fetchargs)
  local utility = self._utility

  local fetchdef, err = self:prepare(fetchargs)
  if err ~= nil then
    return { ok = false, err = err }, nil
  end

  fetchargs = fetchargs or {}
  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "direct",
    ctrl = ctrl,
  }, self._rootctx)

  local url = fetchdef["url"] or ""
  local fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

  if fetch_err ~= nil then
    return { ok = false, err = fetch_err }, nil
  end

  if fetched == nil then
    return {
      ok = false,
      err = ctx:make_error("direct_no_response", "response: undefined"),
    }, nil
  end

  if type(fetched) == "table" then
    local status = helpers.to_int(vs.getprop(fetched, "status"))
    local headers = vs.getprop(fetched, "headers") or {}

    -- No-body responses (204, 304) and explicit zero content-length
    -- must skip JSON parsing — calling json() on an empty body errors.
    local content_length = nil
    if type(headers) == "table" then
      content_length = headers["content-length"]
    end
    local no_body = status == 204 or status == 304 or tostring(content_length) == "0"

    local json_data = nil
    if not no_body then
      local jf = vs.getprop(fetched, "json")
      if type(jf) == "function" then
        local ok, result = pcall(jf)
        if ok then
          json_data = result
        end
        -- Non-JSON body: json_data stays nil, status/headers preserved.
      end
    end

    return {
      ok = status >= 200 and status < 300,
      status = status,
      headers = headers,
      data = json_data,
    }, nil
  end

  return {
    ok = false,
    err = ctx:make_error("direct_invalid", "invalid response type"),
  }, nil
end



-- Idiomatic facade: client:Account():list() / client:Account():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Account(data)
  local EntityMod = require("entity.account_entity")
  if data == nil then
    if self._account == nil then
      self._account = EntityMod.new(self, nil)
    end
    return self._account
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Batch():list() / client:Batch():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Batch(data)
  local EntityMod = require("entity.batch_entity")
  if data == nil then
    if self._batch == nil then
      self._batch = EntityMod.new(self, nil)
    end
    return self._batch
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Beneficiary():list() / client:Beneficiary():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Beneficiary(data)
  local EntityMod = require("entity.beneficiary_entity")
  if data == nil then
    if self._beneficiary == nil then
      self._beneficiary = EntityMod.new(self, nil)
    end
    return self._beneficiary
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:BeneficiaryGroup():list() / client:BeneficiaryGroup():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:BeneficiaryGroup(data)
  local EntityMod = require("entity.beneficiary_group_entity")
  if data == nil then
    if self._beneficiary_group == nil then
      self._beneficiary_group = EntityMod.new(self, nil)
    end
    return self._beneficiary_group
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Card():list() / client:Card():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Card(data)
  local EntityMod = require("entity.card_entity")
  if data == nil then
    if self._card == nil then
      self._card = EntityMod.new(self, nil)
    end
    return self._card
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CardCustomerToken():list() / client:CardCustomerToken():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:CardCustomerToken(data)
  local EntityMod = require("entity.card_customer_token_entity")
  if data == nil then
    if self._card_customer_token == nil then
      self._card_customer_token = EntityMod.new(self, nil)
    end
    return self._card_customer_token
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CardPayment():list() / client:CardPayment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:CardPayment(data)
  local EntityMod = require("entity.card_payment_entity")
  if data == nil then
    if self._card_payment == nil then
      self._card_payment = EntityMod.new(self, nil)
    end
    return self._card_payment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:CardPublicKey():list() / client:CardPublicKey():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:CardPublicKey(data)
  local EntityMod = require("entity.card_public_key_entity")
  if data == nil then
    if self._card_public_key == nil then
      self._card_public_key = EntityMod.new(self, nil)
    end
    return self._card_public_key
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Consent():list() / client:Consent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Consent(data)
  local EntityMod = require("entity.consent_entity")
  if data == nil then
    if self._consent == nil then
      self._consent = EntityMod.new(self, nil)
    end
    return self._consent
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Currency():list() / client:Currency():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Currency(data)
  local EntityMod = require("entity.currency_entity")
  if data == nil then
    if self._currency == nil then
      self._currency = EntityMod.new(self, nil)
    end
    return self._currency
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:DirectDebitBatchSubmit():list() / client:DirectDebitBatchSubmit():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:DirectDebitBatchSubmit(data)
  local EntityMod = require("entity.direct_debit_batch_submit_entity")
  if data == nil then
    if self._direct_debit_batch_submit == nil then
      self._direct_debit_batch_submit = EntityMod.new(self, nil)
    end
    return self._direct_debit_batch_submit
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:FxRate():list() / client:FxRate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:FxRate(data)
  local EntityMod = require("entity.fx_rate_entity")
  if data == nil then
    if self._fx_rate == nil then
      self._fx_rate = EntityMod.new(self, nil)
    end
    return self._fx_rate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:IPayment():list() / client:IPayment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:IPayment(data)
  local EntityMod = require("entity.i_payment_entity")
  if data == nil then
    if self._i_payment == nil then
      self._i_payment = EntityMod.new(self, nil)
    end
    return self._i_payment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Mandate():list() / client:Mandate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Mandate(data)
  local EntityMod = require("entity.mandate_entity")
  if data == nil then
    if self._mandate == nil then
      self._mandate = EntityMod.new(self, nil)
    end
    return self._mandate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Merchant():list() / client:Merchant():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Merchant(data)
  local EntityMod = require("entity.merchant_entity")
  if data == nil then
    if self._merchant == nil then
      self._merchant = EntityMod.new(self, nil)
    end
    return self._merchant
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MerchantAuthorisationSetting():list() / client:MerchantAuthorisationSetting():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:MerchantAuthorisationSetting(data)
  local EntityMod = require("entity.merchant_authorisation_setting_entity")
  if data == nil then
    if self._merchant_authorisation_setting == nil then
      self._merchant_authorisation_setting = EntityMod.new(self, nil)
    end
    return self._merchant_authorisation_setting
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MerchantDirectDebitMandate():list() / client:MerchantDirectDebitMandate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:MerchantDirectDebitMandate(data)
  local EntityMod = require("entity.merchant_direct_debit_mandate_entity")
  if data == nil then
    if self._merchant_direct_debit_mandate == nil then
      self._merchant_direct_debit_mandate = EntityMod.new(self, nil)
    end
    return self._merchant_direct_debit_mandate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MerchantPayByBankSetting():list() / client:MerchantPayByBankSetting():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:MerchantPayByBankSetting(data)
  local EntityMod = require("entity.merchant_pay_by_bank_setting_entity")
  if data == nil then
    if self._merchant_pay_by_bank_setting == nil then
      self._merchant_pay_by_bank_setting = EntityMod.new(self, nil)
    end
    return self._merchant_pay_by_bank_setting
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MerchantPaymentRequestTemplate():list() / client:MerchantPaymentRequestTemplate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:MerchantPaymentRequestTemplate(data)
  local EntityMod = require("entity.merchant_payment_request_template_entity")
  if data == nil then
    if self._merchant_payment_request_template == nil then
      self._merchant_payment_request_template = EntityMod.new(self, nil)
    end
    return self._merchant_payment_request_template
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MerchantToken():list() / client:MerchantToken():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:MerchantToken(data)
  local EntityMod = require("entity.merchant_token_entity")
  if data == nil then
    if self._merchant_token == nil then
      self._merchant_token = EntityMod.new(self, nil)
    end
    return self._merchant_token
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Metadata():list() / client:Metadata():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Metadata(data)
  local EntityMod = require("entity.metadata_entity")
  if data == nil then
    if self._metadata == nil then
      self._metadata = EntityMod.new(self, nil)
    end
    return self._metadata
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionVersion():list() / client:NoFrixionVersion():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionVersion(data)
  local EntityMod = require("entity.no_frixion_version_entity")
  if data == nil then
    if self._no_frixion_version == nil then
      self._no_frixion_version = EntityMod.new(self, nil)
    end
    return self._no_frixion_version
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:OpenBanking():list() / client:OpenBanking():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:OpenBanking(data)
  local EntityMod = require("entity.open_banking_entity")
  if data == nil then
    if self._open_banking == nil then
      self._open_banking = EntityMod.new(self, nil)
    end
    return self._open_banking
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Payeeverification():list() / client:Payeeverification():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Payeeverification(data)
  local EntityMod = require("entity.payeeverification_entity")
  if data == nil then
    if self._payeeverification == nil then
      self._payeeverification = EntityMod.new(self, nil)
    end
    return self._payeeverification
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Payment():list() / client:Payment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Payment(data)
  local EntityMod = require("entity.payment_entity")
  if data == nil then
    if self._payment == nil then
      self._payment = EntityMod.new(self, nil)
    end
    return self._payment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PaymentAccount():list() / client:PaymentAccount():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PaymentAccount(data)
  local EntityMod = require("entity.payment_account_entity")
  if data == nil then
    if self._payment_account == nil then
      self._payment_account = EntityMod.new(self, nil)
    end
    return self._payment_account
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PaymentAccountMinimal():list() / client:PaymentAccountMinimal():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PaymentAccountMinimal(data)
  local EntityMod = require("entity.payment_account_minimal_entity")
  if data == nil then
    if self._payment_account_minimal == nil then
      self._payment_account_minimal = EntityMod.new(self, nil)
    end
    return self._payment_account_minimal
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PaymentInitiation():list() / client:PaymentInitiation():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PaymentInitiation(data)
  local EntityMod = require("entity.payment_initiation_entity")
  if data == nil then
    if self._payment_initiation == nil then
      self._payment_initiation = EntityMod.new(self, nil)
    end
    return self._payment_initiation
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PaymentRequest():list() / client:PaymentRequest():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PaymentRequest(data)
  local EntityMod = require("entity.payment_request_entity")
  if data == nil then
    if self._payment_request == nil then
      self._payment_request = EntityMod.new(self, nil)
    end
    return self._payment_request
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PaymentRequestEvent():list() / client:PaymentRequestEvent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PaymentRequestEvent(data)
  local EntityMod = require("entity.payment_request_event_entity")
  if data == nil then
    if self._payment_request_event == nil then
      self._payment_request_event = EntityMod.new(self, nil)
    end
    return self._payment_request_event
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PaymentRequestMetric():list() / client:PaymentRequestMetric():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PaymentRequestMetric(data)
  local EntityMod = require("entity.payment_request_metric_entity")
  if data == nil then
    if self._payment_request_metric == nil then
      self._payment_request_metric = EntityMod.new(self, nil)
    end
    return self._payment_request_metric
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PaymentRequestMinimal():list() / client:PaymentRequestMinimal():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PaymentRequestMinimal(data)
  local EntityMod = require("entity.payment_request_minimal_entity")
  if data == nil then
    if self._payment_request_minimal == nil then
      self._payment_request_minimal = EntityMod.new(self, nil)
    end
    return self._payment_request_minimal
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PaymentRequestResult():list() / client:PaymentRequestResult():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PaymentRequestResult(data)
  local EntityMod = require("entity.payment_request_result_entity")
  if data == nil then
    if self._payment_request_result == nil then
      self._payment_request_result = EntityMod.new(self, nil)
    end
    return self._payment_request_result
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Payout():list() / client:Payout():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Payout(data)
  local EntityMod = require("entity.payout_entity")
  if data == nil then
    if self._payout == nil then
      self._payout = EntityMod.new(self, nil)
    end
    return self._payout
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PayoutKeyset():list() / client:PayoutKeyset():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PayoutKeyset(data)
  local EntityMod = require("entity.payout_keyset_entity")
  if data == nil then
    if self._payout_keyset == nil then
      self._payout_keyset = EntityMod.new(self, nil)
    end
    return self._payout_keyset
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PayoutMetric():list() / client:PayoutMetric():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:PayoutMetric(data)
  local EntityMod = require("entity.payout_metric_entity")
  if data == nil then
    if self._payout_metric == nil then
      self._payout_metric = EntityMod.new(self, nil)
    end
    return self._payout_metric
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Payrun():list() / client:Payrun():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Payrun(data)
  local EntityMod = require("entity.payrun_entity")
  if data == nil then
    if self._payrun == nil then
      self._payrun = EntityMod.new(self, nil)
    end
    return self._payrun
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Report():list() / client:Report():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Report(data)
  local EntityMod = require("entity.report_entity")
  if data == nil then
    if self._report == nil then
      self._report = EntityMod.new(self, nil)
    end
    return self._report
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ReportResult():list() / client:ReportResult():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:ReportResult(data)
  local EntityMod = require("entity.report_result_entity")
  if data == nil then
    if self._report_result == nil then
      self._report_result = EntityMod.new(self, nil)
    end
    return self._report_result
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Role():list() / client:Role():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Role(data)
  local EntityMod = require("entity.role_entity")
  if data == nil then
    if self._role == nil then
      self._role = EntityMod.new(self, nil)
    end
    return self._role
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Rule():list() / client:Rule():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Rule(data)
  local EntityMod = require("entity.rule_entity")
  if data == nil then
    if self._rule == nil then
      self._rule = EntityMod.new(self, nil)
    end
    return self._rule
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:RuleEvent():list() / client:RuleEvent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:RuleEvent(data)
  local EntityMod = require("entity.rule_event_entity")
  if data == nil then
    if self._rule_event == nil then
      self._rule_event = EntityMod.new(self, nil)
    end
    return self._rule_event
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Tag():list() / client:Tag():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Tag(data)
  local EntityMod = require("entity.tag_entity")
  if data == nil then
    if self._tag == nil then
      self._tag = EntityMod.new(self, nil)
    end
    return self._tag
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Token():list() / client:Token():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Token(data)
  local EntityMod = require("entity.token_entity")
  if data == nil then
    if self._token == nil then
      self._token = EntityMod.new(self, nil)
    end
    return self._token
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Transaction():list() / client:Transaction():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Transaction(data)
  local EntityMod = require("entity.transaction_entity")
  if data == nil then
    if self._transaction == nil then
      self._transaction = EntityMod.new(self, nil)
    end
    return self._transaction
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:User():list() / client:User():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:User(data)
  local EntityMod = require("entity.user_entity")
  if data == nil then
    if self._user == nil then
      self._user = EntityMod.new(self, nil)
    end
    return self._user
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:UserInvite():list() / client:UserInvite():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:UserInvite(data)
  local EntityMod = require("entity.user_invite_entity")
  if data == nil then
    if self._user_invite == nil then
      self._user_invite = EntityMod.new(self, nil)
    end
    return self._user_invite
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Virtual():list() / client:Virtual():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Virtual(data)
  local EntityMod = require("entity.virtual_entity")
  if data == nil then
    if self._virtual == nil then
      self._virtual = EntityMod.new(self, nil)
    end
    return self._virtual
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Webhook():list() / client:Webhook():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Webhook(data)
  local EntityMod = require("entity.webhook_entity")
  if data == nil then
    if self._webhook == nil then
      self._webhook = EntityMod.new(self, nil)
    end
    return self._webhook
  end
  return EntityMod.new(self, data)
end




function NofrixionSDK.test(testopts, sdkopts)
  sdkopts = sdkopts or {}
  sdkopts = vs.clone(sdkopts)
  if type(sdkopts) ~= "table" then
    sdkopts = {}
  end

  testopts = testopts or {}
  testopts = vs.clone(testopts)
  if type(testopts) ~= "table" then
    testopts = {}
  end
  testopts["active"] = true

  vs.setpath(sdkopts, "feature.test", testopts)

  local sdk = NofrixionSDK.new(sdkopts)
  sdk.mode = "test"

  return sdk
end


return NofrixionSDK
