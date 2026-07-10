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


-- Idiomatic facade: client:Cancel():list() / client:Cancel():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Cancel(data)
  local EntityMod = require("entity.cancel_entity")
  if data == nil then
    if self._cancel == nil then
      self._cancel = EntityMod.new(self, nil)
    end
    return self._cancel
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Disable():list() / client:Disable():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Disable(data)
  local EntityMod = require("entity.disable_entity")
  if data == nil then
    if self._disable == nil then
      self._disable = EntityMod.new(self, nil)
    end
    return self._disable
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Enable():list() / client:Enable():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Enable(data)
  local EntityMod = require("entity.enable_entity")
  if data == nil then
    if self._enable == nil then
      self._enable = EntityMod.new(self, nil)
    end
    return self._enable
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


-- Idiomatic facade: client:NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage():list() / client:NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(data)
  local EntityMod = require("entity.no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page_entity")
  if data == nil then
    if self._no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page == nil then
      self._no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionBizBizModelsPagingPaymentRequestPage():list() / client:NoFrixionBizBizModelsPagingPaymentRequestPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionBizBizModelsPagingPaymentRequestPage(data)
  local EntityMod = require("entity.no_frixion_biz_biz_models_paging_payment_request_page_entity")
  if data == nil then
    if self._no_frixion_biz_biz_models_paging_payment_request_page == nil then
      self._no_frixion_biz_biz_models_paging_payment_request_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_biz_biz_models_paging_payment_request_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionBizBizModelsPagingPayoutPage():list() / client:NoFrixionBizBizModelsPagingPayoutPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionBizBizModelsPagingPayoutPage(data)
  local EntityMod = require("entity.no_frixion_biz_biz_models_paging_payout_page_entity")
  if data == nil then
    if self._no_frixion_biz_biz_models_paging_payout_page == nil then
      self._no_frixion_biz_biz_models_paging_payout_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_biz_biz_models_paging_payout_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionBizBizModelsPagingPayrunPage():list() / client:NoFrixionBizBizModelsPagingPayrunPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionBizBizModelsPagingPayrunPage(data)
  local EntityMod = require("entity.no_frixion_biz_biz_models_paging_payrun_page_entity")
  if data == nil then
    if self._no_frixion_biz_biz_models_paging_payrun_page == nil then
      self._no_frixion_biz_biz_models_paging_payrun_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_biz_biz_models_paging_payrun_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionBizBizModelsPagingRuleEventsPage():list() / client:NoFrixionBizBizModelsPagingRuleEventsPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionBizBizModelsPagingRuleEventsPage(data)
  local EntityMod = require("entity.no_frixion_biz_biz_models_paging_rule_events_page_entity")
  if data == nil then
    if self._no_frixion_biz_biz_models_paging_rule_events_page == nil then
      self._no_frixion_biz_biz_models_paging_rule_events_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_biz_biz_models_paging_rule_events_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionBizBizModelsPagingRulesPage():list() / client:NoFrixionBizBizModelsPagingRulesPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionBizBizModelsPagingRulesPage(data)
  local EntityMod = require("entity.no_frixion_biz_biz_models_paging_rules_page_entity")
  if data == nil then
    if self._no_frixion_biz_biz_models_paging_rules_page == nil then
      self._no_frixion_biz_biz_models_paging_rules_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_biz_biz_models_paging_rules_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionBizBizModelsPaymentsCardPayment():list() / client:NoFrixionBizBizModelsPaymentsCardPayment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionBizBizModelsPaymentsCardPayment(data)
  local EntityMod = require("entity.no_frixion_biz_biz_models_payments_card_payment_entity")
  if data == nil then
    if self._no_frixion_biz_biz_models_payments_card_payment == nil then
      self._no_frixion_biz_biz_models_payments_card_payment = EntityMod.new(self, nil)
    end
    return self._no_frixion_biz_biz_models_payments_card_payment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionBizBizModelsPaymentsCardPublicKey():list() / client:NoFrixionBizBizModelsPaymentsCardPublicKey():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionBizBizModelsPaymentsCardPublicKey(data)
  local EntityMod = require("entity.no_frixion_biz_biz_models_payments_card_public_key_entity")
  if data == nil then
    if self._no_frixion_biz_biz_models_payments_card_public_key == nil then
      self._no_frixion_biz_biz_models_payments_card_public_key = EntityMod.new(self, nil)
    end
    return self._no_frixion_biz_biz_models_payments_card_public_key
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries():list() / client:NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(data)
  local EntityMod = require("entity.no_frixion_money_moov_api_features_beneficiaries_beneficiaries_entity")
  if data == nil then
    if self._no_frixion_money_moov_api_features_beneficiaries_beneficiaries == nil then
      self._no_frixion_money_moov_api_features_beneficiaries_beneficiaries = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_api_features_beneficiaries_beneficiaries
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment():list() / client:NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(data)
  local EntityMod = require("entity.no_frixion_money_moov_api_features_payment_requests_payment_entity")
  if data == nil then
    if self._no_frixion_money_moov_api_features_payment_requests_payment == nil then
      self._no_frixion_money_moov_api_features_payment_requests_payment = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_api_features_payment_requests_payment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate():list() / client:NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(data)
  local EntityMod = require("entity.no_frixion_money_moov_api_features_permissions_roles_create_entity")
  if data == nil then
    if self._no_frixion_money_moov_api_features_permissions_roles_create == nil then
      self._no_frixion_money_moov_api_features_permissions_roles_create = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_api_features_permissions_roles_create
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovApiFeaturesUserInvitesCreate():list() / client:NoFrixionMoneyMoovApiFeaturesUserInvitesCreate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(data)
  local EntityMod = require("entity.no_frixion_money_moov_api_features_user_invites_create_entity")
  if data == nil then
    if self._no_frixion_money_moov_api_features_user_invites_create == nil then
      self._no_frixion_money_moov_api_features_user_invites_create = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_api_features_user_invites_create
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant():list() / client:NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_authorisation_settings_merchant_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_authorisation_settings_merchant == nil then
      self._no_frixion_money_moov_models_authorisation_settings_merchant = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_authorisation_settings_merchant
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsBatchPayout():list() / client:NoFrixionMoneyMoovModelsBatchPayout():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsBatchPayout(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_batch_payout_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_batch_payout == nil then
      self._no_frixion_money_moov_models_batch_payout = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_batch_payout
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsBeneficiaryGroupPage():list() / client:NoFrixionMoneyMoovModelsBeneficiaryGroupPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsBeneficiaryGroupPage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_beneficiary_group_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_beneficiary_group_page == nil then
      self._no_frixion_money_moov_models_beneficiary_group_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_beneficiary_group_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsBeneficiaryPage():list() / client:NoFrixionMoneyMoovModelsBeneficiaryPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsBeneficiaryPage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_beneficiary_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_beneficiary_page == nil then
      self._no_frixion_money_moov_models_beneficiary_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_beneficiary_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsCardCustomerToken():list() / client:NoFrixionMoneyMoovModelsCardCustomerToken():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsCardCustomerToken(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_card_customer_token_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_card_customer_token == nil then
      self._no_frixion_money_moov_models_card_customer_token = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_card_customer_token
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsCurrencyCurrencyInfo():list() / client:NoFrixionMoneyMoovModelsCurrencyCurrencyInfo():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_currency_currency_info_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_currency_currency_info == nil then
      self._no_frixion_money_moov_models_currency_currency_info = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_currency_currency_info
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsDirectDebitBatchSubmit():list() / client:NoFrixionMoneyMoovModelsDirectDebitBatchSubmit():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_direct_debit_batch_submit_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_direct_debit_batch_submit == nil then
      self._no_frixion_money_moov_models_direct_debit_batch_submit = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_direct_debit_batch_submit
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsFxRate():list() / client:NoFrixionMoneyMoovModelsFxRate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsFxRate(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_fx_rate_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_fx_rate == nil then
      self._no_frixion_money_moov_models_fx_rate = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_fx_rate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsIPayment():list() / client:NoFrixionMoneyMoovModelsIPayment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsIPayment(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_i_payment_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_i_payment == nil then
      self._no_frixion_money_moov_models_i_payment = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_i_payment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsMandatesMandate():list() / client:NoFrixionMoneyMoovModelsMandatesMandate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsMandatesMandate(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_mandates_mandate_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_mandates_mandate == nil then
      self._no_frixion_money_moov_models_mandates_mandate = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_mandates_mandate
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsMerchant():list() / client:NoFrixionMoneyMoovModelsMerchant():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsMerchant(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_merchant_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_merchant == nil then
      self._no_frixion_money_moov_models_merchant = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_merchant
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsMerchantPage():list() / client:NoFrixionMoneyMoovModelsMerchantPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsMerchantPage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_merchant_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_merchant_page == nil then
      self._no_frixion_money_moov_models_merchant_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_merchant_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsMerchantPayByBankSetting():list() / client:NoFrixionMoneyMoovModelsMerchantPayByBankSetting():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsMerchantPayByBankSetting(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_merchant_pay_by_bank_setting_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_merchant_pay_by_bank_setting == nil then
      self._no_frixion_money_moov_models_merchant_pay_by_bank_setting = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_merchant_pay_by_bank_setting
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsMerchantToken():list() / client:NoFrixionMoneyMoovModelsMerchantToken():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsMerchantToken(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_merchant_token_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_merchant_token == nil then
      self._no_frixion_money_moov_models_merchant_token = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_merchant_token
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsMerchantTokenPage():list() / client:NoFrixionMoneyMoovModelsMerchantTokenPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsMerchantTokenPage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_merchant_token_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_merchant_token_page == nil then
      self._no_frixion_money_moov_models_merchant_token_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_merchant_token_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsNoFrixionVersion():list() / client:NoFrixionMoneyMoovModelsNoFrixionVersion():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsNoFrixionVersion(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_no_frixion_version_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_no_frixion_version == nil then
      self._no_frixion_money_moov_models_no_frixion_version = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_no_frixion_version
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsOpenBankingAccount():list() / client:NoFrixionMoneyMoovModelsOpenBankingAccount():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsOpenBankingAccount(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_open_banking_account_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_open_banking_account == nil then
      self._no_frixion_money_moov_models_open_banking_account = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_open_banking_account
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsOpenBankingConsent():list() / client:NoFrixionMoneyMoovModelsOpenBankingConsent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsOpenBankingConsent(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_open_banking_consent_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_open_banking_consent == nil then
      self._no_frixion_money_moov_models_open_banking_consent = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_open_banking_consent
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsOpenBankingTransaction():list() / client:NoFrixionMoneyMoovModelsOpenBankingTransaction():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsOpenBankingTransaction(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_open_banking_transaction_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_open_banking_transaction == nil then
      self._no_frixion_money_moov_models_open_banking_transaction = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_open_banking_transaction
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPayment():list() / client:NoFrixionMoneyMoovModelsPayment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPayment(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment == nil then
      self._no_frixion_money_moov_models_payment = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentAccountMinimalPage():list() / client:NoFrixionMoneyMoovModelsPaymentAccountMinimalPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_account_minimal_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_account_minimal_page == nil then
      self._no_frixion_money_moov_models_payment_account_minimal_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_account_minimal_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentAccountPage():list() / client:NoFrixionMoneyMoovModelsPaymentAccountPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentAccountPage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_account_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_account_page == nil then
      self._no_frixion_money_moov_models_payment_account_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_account_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentInitiation():list() / client:NoFrixionMoneyMoovModelsPaymentInitiation():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentInitiation(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_initiation_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_initiation == nil then
      self._no_frixion_money_moov_models_payment_initiation = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_initiation
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentRequestEvent():list() / client:NoFrixionMoneyMoovModelsPaymentRequestEvent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentRequestEvent(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_request_event_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_request_event == nil then
      self._no_frixion_money_moov_models_payment_request_event = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_request_event
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentRequestMetric():list() / client:NoFrixionMoneyMoovModelsPaymentRequestMetric():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentRequestMetric(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_request_metric_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_request_metric == nil then
      self._no_frixion_money_moov_models_payment_request_metric = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_request_metric
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentRequestMinimal():list() / client:NoFrixionMoneyMoovModelsPaymentRequestMinimal():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentRequestMinimal(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_request_minimal_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_request_minimal == nil then
      self._no_frixion_money_moov_models_payment_request_minimal = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_request_minimal
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentRequestResult():list() / client:NoFrixionMoneyMoovModelsPaymentRequestResult():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentRequestResult(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_request_result_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_request_result == nil then
      self._no_frixion_money_moov_models_payment_request_result = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_request_result
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment():list() / client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_requests_merchant_payment_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_requests_merchant_payment == nil then
      self._no_frixion_money_moov_models_payment_requests_merchant_payment = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_requests_merchant_payment
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2():list() / client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_requests_merchant_payment2_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_requests_merchant_payment2 == nil then
      self._no_frixion_money_moov_models_payment_requests_merchant_payment2 = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_requests_merchant_payment2
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3():list() / client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_requests_merchant_payment3_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_requests_merchant_payment3 == nil then
      self._no_frixion_money_moov_models_payment_requests_merchant_payment3 = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_requests_merchant_payment3
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4():list() / client:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payment_requests_merchant_payment4_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payment_requests_merchant_payment4 == nil then
      self._no_frixion_money_moov_models_payment_requests_merchant_payment4 = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payment_requests_merchant_payment4
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPayoutKeysetPage():list() / client:NoFrixionMoneyMoovModelsPayoutKeysetPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPayoutKeysetPage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payout_keyset_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payout_keyset_page == nil then
      self._no_frixion_money_moov_models_payout_keyset_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payout_keyset_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPayoutMetric():list() / client:NoFrixionMoneyMoovModelsPayoutMetric():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPayoutMetric(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payout_metric_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payout_metric == nil then
      self._no_frixion_money_moov_models_payout_metric = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payout_metric
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPayoutsPayoutsCreate():list() / client:NoFrixionMoneyMoovModelsPayoutsPayoutsCreate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payouts_payouts_create_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payouts_payouts_create == nil then
      self._no_frixion_money_moov_models_payouts_payouts_create = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payouts_payouts_create
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsPayrun():list() / client:NoFrixionMoneyMoovModelsPayrun():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsPayrun(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_payrun_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_payrun == nil then
      self._no_frixion_money_moov_models_payrun = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_payrun
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsReportResult():list() / client:NoFrixionMoneyMoovModelsReportResult():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsReportResult(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_report_result_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_report_result == nil then
      self._no_frixion_money_moov_models_report_result = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_report_result
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsRule():list() / client:NoFrixionMoneyMoovModelsRule():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsRule(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_rule_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_rule == nil then
      self._no_frixion_money_moov_models_rule = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_rule
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsTransaction():list() / client:NoFrixionMoneyMoovModelsTransaction():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsTransaction(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_transaction_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_transaction == nil then
      self._no_frixion_money_moov_models_transaction = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_transaction
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsTransactionPage():list() / client:NoFrixionMoneyMoovModelsTransactionPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsTransactionPage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_transaction_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_transaction_page == nil then
      self._no_frixion_money_moov_models_transaction_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_transaction_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsUserInvite():list() / client:NoFrixionMoneyMoovModelsUserInvite():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsUserInvite(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_user_invite_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_user_invite == nil then
      self._no_frixion_money_moov_models_user_invite = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_user_invite
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsUserInvitePage():list() / client:NoFrixionMoneyMoovModelsUserInvitePage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsUserInvitePage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_user_invite_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_user_invite_page == nil then
      self._no_frixion_money_moov_models_user_invite_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_user_invite_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsUserPage():list() / client:NoFrixionMoneyMoovModelsUserPage():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsUserPage(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_user_page_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_user_page == nil then
      self._no_frixion_money_moov_models_user_page = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_user_page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:NoFrixionMoneyMoovModelsWebhook():list() / client:NoFrixionMoneyMoovModelsWebhook():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:NoFrixionMoneyMoovModelsWebhook(data)
  local EntityMod = require("entity.no_frixion_money_moov_models_webhook_entity")
  if data == nil then
    if self._no_frixion_money_moov_models_webhook == nil then
      self._no_frixion_money_moov_models_webhook = EntityMod.new(self, nil)
    end
    return self._no_frixion_money_moov_models_webhook
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


-- Idiomatic facade: client:Reject():list() / client:Reject():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Reject(data)
  local EntityMod = require("entity.reject_entity")
  if data == nil then
    if self._reject == nil then
      self._reject = EntityMod.new(self, nil)
    end
    return self._reject
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


-- Idiomatic facade: client:Send():list() / client:Send():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Send(data)
  local EntityMod = require("entity.send_entity")
  if data == nil then
    if self._send == nil then
      self._send = EntityMod.new(self, nil)
    end
    return self._send
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Sendbeneficiary():list() / client:Sendbeneficiary():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Sendbeneficiary(data)
  local EntityMod = require("entity.sendbeneficiary_entity")
  if data == nil then
    if self._sendbeneficiary == nil then
      self._sendbeneficiary = EntityMod.new(self, nil)
    end
    return self._sendbeneficiary
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


-- Idiomatic facade: client:Whoami():list() / client:Whoami():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Whoami(data)
  local EntityMod = require("entity.whoami_entity")
  if data == nil then
    if self._whoami == nil then
      self._whoami = EntityMod.new(self, nil)
    end
    return self._whoami
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Whoamitrustedapp():list() / client:Whoamitrustedapp():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function NofrixionSDK:Whoamitrustedapp(data)
  local EntityMod = require("entity.whoamitrustedapp_entity")
  if data == nil then
    if self._whoamitrustedapp == nil then
      self._whoamitrustedapp = EntityMod.new(self, nil)
    end
    return self._whoamitrustedapp
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
