# Nofrixion SDK

require_relative 'utility/struct/voxgig_struct'
require_relative 'core/utility_type'
require_relative 'core/spec'
require_relative 'core/helpers'

# Load utility registration
require_relative 'utility/register'

# Load config and features
require_relative 'config'
require_relative 'feature/base_feature'
require_relative 'features'

# Load typed models (Struct value objects).
require_relative 'Nofrixion_types'


class NofrixionSDK
  attr_accessor :mode, :features, :options

  def initialize(options = {})
    @mode = "live"
    @features = []
    @options = nil

    utility = NofrixionUtility.new
    @_utility = utility

    config = NofrixionConfig.make_config

    @_rootctx = utility.make_context.call({
      "client" => self,
      "utility" => utility,
      "config" => config,
      "options" => options || {},
      "shared" => {},
    }, nil)

    @options = utility.make_options.call(@_rootctx)

    if VoxgigStruct.getpath(@options, "feature.test.active") == true
      @mode = "test"
    end

    @_rootctx.options = @options

    # Add features from config.
    feature_opts = NofrixionHelpers.to_map(VoxgigStruct.getprop(@options, "feature"))
    if feature_opts
      items = VoxgigStruct.items(feature_opts)
      if items
        items.each do |item|
          fname = item[0]
          fopts = NofrixionHelpers.to_map(item[1])
          if fopts && fopts["active"] == true
            utility.feature_add.call(@_rootctx, NofrixionFeatures.make_feature(fname))
          end
        end
      end
    end

    # Add extension features.
    extend_val = VoxgigStruct.getprop(@options, "extend")
    if extend_val.is_a?(Array)
      extend_val.each do |f|
        if f.respond_to?(:get_name)
          utility.feature_add.call(@_rootctx, f)
        end
      end
    end

    # Initialize features.
    @features.each do |f|
      utility.feature_init.call(@_rootctx, f)
    end

    utility.feature_hook.call(@_rootctx, "PostConstruct")
  end

  def options_map
    out = VoxgigStruct.clone(@options)
    out.is_a?(Hash) ? out : {}
  end

  def get_utility
    NofrixionUtility.copy(@_utility)
  end

  def get_root_ctx
    @_rootctx
  end

  def prepare(fetchargs = {})
    utility = @_utility
    fetchargs ||= {}

    ctrl = NofrixionHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "prepare",
      "ctrl" => ctrl,
    }, @_rootctx)

    opts = @options
    path = VoxgigStruct.getprop(fetchargs, "path") || ""
    path = "" unless path.is_a?(String)
    method_val = VoxgigStruct.getprop(fetchargs, "method") || "GET"
    method_val = "GET" unless method_val.is_a?(String)
    params = NofrixionHelpers.to_map(VoxgigStruct.getprop(fetchargs, "params")) || {}
    query = NofrixionHelpers.to_map(VoxgigStruct.getprop(fetchargs, "query")) || {}
    headers = utility.prepare_headers.call(ctx)

    base = VoxgigStruct.getprop(opts, "base") || ""
    base = "" unless base.is_a?(String)
    prefix = VoxgigStruct.getprop(opts, "prefix") || ""
    prefix = "" unless prefix.is_a?(String)
    suffix = VoxgigStruct.getprop(opts, "suffix") || ""
    suffix = "" unless suffix.is_a?(String)

    ctx.spec = NofrixionSpec.new({
      "base" => base, "prefix" => prefix, "suffix" => suffix,
      "path" => path, "method" => method_val,
      "params" => params, "query" => query, "headers" => headers,
      "body" => VoxgigStruct.getprop(fetchargs, "body"),
      "step" => "start",
    })

    # Merge user-provided headers.
    uh = VoxgigStruct.getprop(fetchargs, "headers")
    if uh.is_a?(Hash)
      uh.each { |k, v| ctx.spec.headers[k] = v }
    end

    _, err = utility.prepare_auth.call(ctx)
    raise err if err

    # make_fetch_def returns a (fetchdef, err) tuple; destructure it and
    # return just the fetchdef Hash (raising on error) so callers — including
    # direct(), which indexes fetchdef["url"] — receive a Hash, mirroring the
    # ts/py prepare().
    fetchdef, fd_err = utility.make_fetch_def.call(ctx)
    raise fd_err if fd_err

    fetchdef
  end

  def direct(fetchargs = {})
    utility = @_utility

    # direct() is the raw-HTTP escape hatch: it always returns a result hash
    # ({ "ok" => ..., ... }) and never raises. prepare() raises on error, so
    # trap that and surface it in the hash.
    begin
      fetchdef = prepare(fetchargs)
    rescue NofrixionError => err
      return { "ok" => false, "err" => err }
    end

    fetchargs ||= {}
    ctrl = NofrixionHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "direct",
      "ctrl" => ctrl,
    }, @_rootctx)

    url = fetchdef["url"] || ""
    fetched, fetch_err = utility.fetcher.call(ctx, url, fetchdef)

    return { "ok" => false, "err" => fetch_err } if fetch_err

    if fetched.nil?
      return {
        "ok" => false,
        "err" => ctx.make_error("direct_no_response", "response: undefined"),
      }
    end

    if fetched.is_a?(Hash)
      status = NofrixionHelpers.to_int(VoxgigStruct.getprop(fetched, "status"))
      headers = VoxgigStruct.getprop(fetched, "headers") || {}

      # No-body responses (204, 304) and explicit zero content-length must
      # skip JSON parsing — calling json() on an empty body errors.
      content_length = headers.is_a?(Hash) ? headers["content-length"] : nil
      no_body = status == 204 || status == 304 || content_length.to_s == "0"

      json_data = nil
      unless no_body
        jf = VoxgigStruct.getprop(fetched, "json")
        if jf.is_a?(Proc)
          begin
            json_data = jf.call
          rescue StandardError
            # Non-JSON body — leave data nil, keep status/headers.
            json_data = nil
          end
        end
      end

      return {
        "ok" => status >= 200 && status < 300,
        "status" => status,
        "headers" => headers,
        "data" => json_data,
      }
    end

    return {
      "ok" => false,
      "err" => ctx.make_error("direct_invalid", "invalid response type"),
    }
  end


  # Canonical facade: client.Account.list / client.Account.load({ "id" => ... })
  def Account(data = nil)
    require_relative 'entity/account_entity'
    AccountEntity.new(self, data)
  end


  # Canonical facade: client.Batch.list / client.Batch.load({ "id" => ... })
  def Batch(data = nil)
    require_relative 'entity/batch_entity'
    BatchEntity.new(self, data)
  end


  # Canonical facade: client.Beneficiary.list / client.Beneficiary.load({ "id" => ... })
  def Beneficiary(data = nil)
    require_relative 'entity/beneficiary_entity'
    BeneficiaryEntity.new(self, data)
  end


  # Canonical facade: client.BeneficiaryGroup.list / client.BeneficiaryGroup.load({ "id" => ... })
  def BeneficiaryGroup(data = nil)
    require_relative 'entity/beneficiary_group_entity'
    BeneficiaryGroupEntity.new(self, data)
  end


  # Canonical facade: client.Card.list / client.Card.load({ "id" => ... })
  def Card(data = nil)
    require_relative 'entity/card_entity'
    CardEntity.new(self, data)
  end


  # Canonical facade: client.CardCustomerToken.list / client.CardCustomerToken.load({ "id" => ... })
  def CardCustomerToken(data = nil)
    require_relative 'entity/card_customer_token_entity'
    CardCustomerTokenEntity.new(self, data)
  end


  # Canonical facade: client.CardPayment.list / client.CardPayment.load({ "id" => ... })
  def CardPayment(data = nil)
    require_relative 'entity/card_payment_entity'
    CardPaymentEntity.new(self, data)
  end


  # Canonical facade: client.CardPublicKey.list / client.CardPublicKey.load({ "id" => ... })
  def CardPublicKey(data = nil)
    require_relative 'entity/card_public_key_entity'
    CardPublicKeyEntity.new(self, data)
  end


  # Canonical facade: client.Consent.list / client.Consent.load({ "id" => ... })
  def Consent(data = nil)
    require_relative 'entity/consent_entity'
    ConsentEntity.new(self, data)
  end


  # Canonical facade: client.Currency.list / client.Currency.load({ "id" => ... })
  def Currency(data = nil)
    require_relative 'entity/currency_entity'
    CurrencyEntity.new(self, data)
  end


  # Canonical facade: client.DirectDebitBatchSubmit.list / client.DirectDebitBatchSubmit.load({ "id" => ... })
  def DirectDebitBatchSubmit(data = nil)
    require_relative 'entity/direct_debit_batch_submit_entity'
    DirectDebitBatchSubmitEntity.new(self, data)
  end


  # Canonical facade: client.FxRate.list / client.FxRate.load({ "id" => ... })
  def FxRate(data = nil)
    require_relative 'entity/fx_rate_entity'
    FxRateEntity.new(self, data)
  end


  # Canonical facade: client.IPayment.list / client.IPayment.load({ "id" => ... })
  def IPayment(data = nil)
    require_relative 'entity/i_payment_entity'
    IPaymentEntity.new(self, data)
  end


  # Canonical facade: client.Mandate.list / client.Mandate.load({ "id" => ... })
  def Mandate(data = nil)
    require_relative 'entity/mandate_entity'
    MandateEntity.new(self, data)
  end


  # Canonical facade: client.Merchant.list / client.Merchant.load({ "id" => ... })
  def Merchant(data = nil)
    require_relative 'entity/merchant_entity'
    MerchantEntity.new(self, data)
  end


  # Canonical facade: client.MerchantAuthorisationSetting.list / client.MerchantAuthorisationSetting.load({ "id" => ... })
  def MerchantAuthorisationSetting(data = nil)
    require_relative 'entity/merchant_authorisation_setting_entity'
    MerchantAuthorisationSettingEntity.new(self, data)
  end


  # Canonical facade: client.MerchantDirectDebitMandatePage.list / client.MerchantDirectDebitMandatePage.load({ "id" => ... })
  def MerchantDirectDebitMandatePage(data = nil)
    require_relative 'entity/merchant_direct_debit_mandate_page_entity'
    MerchantDirectDebitMandatePageEntity.new(self, data)
  end


  # Canonical facade: client.MerchantPayByBankSetting.list / client.MerchantPayByBankSetting.load({ "id" => ... })
  def MerchantPayByBankSetting(data = nil)
    require_relative 'entity/merchant_pay_by_bank_setting_entity'
    MerchantPayByBankSettingEntity.new(self, data)
  end


  # Canonical facade: client.MerchantPaymentRequestTemplate.list / client.MerchantPaymentRequestTemplate.load({ "id" => ... })
  def MerchantPaymentRequestTemplate(data = nil)
    require_relative 'entity/merchant_payment_request_template_entity'
    MerchantPaymentRequestTemplateEntity.new(self, data)
  end


  # Canonical facade: client.MerchantToken.list / client.MerchantToken.load({ "id" => ... })
  def MerchantToken(data = nil)
    require_relative 'entity/merchant_token_entity'
    MerchantTokenEntity.new(self, data)
  end


  # Canonical facade: client.Metadata.list / client.Metadata.load({ "id" => ... })
  def Metadata(data = nil)
    require_relative 'entity/metadata_entity'
    MetadataEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionVersion.list / client.NoFrixionVersion.load({ "id" => ... })
  def NoFrixionVersion(data = nil)
    require_relative 'entity/no_frixion_version_entity'
    NoFrixionVersionEntity.new(self, data)
  end


  # Canonical facade: client.OpenBanking.list / client.OpenBanking.load({ "id" => ... })
  def OpenBanking(data = nil)
    require_relative 'entity/open_banking_entity'
    OpenBankingEntity.new(self, data)
  end


  # Canonical facade: client.Payeeverification.list / client.Payeeverification.load({ "id" => ... })
  def Payeeverification(data = nil)
    require_relative 'entity/payeeverification_entity'
    PayeeverificationEntity.new(self, data)
  end


  # Canonical facade: client.Payment.list / client.Payment.load({ "id" => ... })
  def Payment(data = nil)
    require_relative 'entity/payment_entity'
    PaymentEntity.new(self, data)
  end


  # Canonical facade: client.PaymentAccount.list / client.PaymentAccount.load({ "id" => ... })
  def PaymentAccount(data = nil)
    require_relative 'entity/payment_account_entity'
    PaymentAccountEntity.new(self, data)
  end


  # Canonical facade: client.PaymentAccountMinimal.list / client.PaymentAccountMinimal.load({ "id" => ... })
  def PaymentAccountMinimal(data = nil)
    require_relative 'entity/payment_account_minimal_entity'
    PaymentAccountMinimalEntity.new(self, data)
  end


  # Canonical facade: client.PaymentInitiation.list / client.PaymentInitiation.load({ "id" => ... })
  def PaymentInitiation(data = nil)
    require_relative 'entity/payment_initiation_entity'
    PaymentInitiationEntity.new(self, data)
  end


  # Canonical facade: client.PaymentRequest.list / client.PaymentRequest.load({ "id" => ... })
  def PaymentRequest(data = nil)
    require_relative 'entity/payment_request_entity'
    PaymentRequestEntity.new(self, data)
  end


  # Canonical facade: client.PaymentRequestEvent.list / client.PaymentRequestEvent.load({ "id" => ... })
  def PaymentRequestEvent(data = nil)
    require_relative 'entity/payment_request_event_entity'
    PaymentRequestEventEntity.new(self, data)
  end


  # Canonical facade: client.PaymentRequestMetric.list / client.PaymentRequestMetric.load({ "id" => ... })
  def PaymentRequestMetric(data = nil)
    require_relative 'entity/payment_request_metric_entity'
    PaymentRequestMetricEntity.new(self, data)
  end


  # Canonical facade: client.PaymentRequestMinimal.list / client.PaymentRequestMinimal.load({ "id" => ... })
  def PaymentRequestMinimal(data = nil)
    require_relative 'entity/payment_request_minimal_entity'
    PaymentRequestMinimalEntity.new(self, data)
  end


  # Canonical facade: client.PaymentRequestResult.list / client.PaymentRequestResult.load({ "id" => ... })
  def PaymentRequestResult(data = nil)
    require_relative 'entity/payment_request_result_entity'
    PaymentRequestResultEntity.new(self, data)
  end


  # Canonical facade: client.Payout.list / client.Payout.load({ "id" => ... })
  def Payout(data = nil)
    require_relative 'entity/payout_entity'
    PayoutEntity.new(self, data)
  end


  # Canonical facade: client.PayoutKeysetPage.list / client.PayoutKeysetPage.load({ "id" => ... })
  def PayoutKeysetPage(data = nil)
    require_relative 'entity/payout_keyset_page_entity'
    PayoutKeysetPageEntity.new(self, data)
  end


  # Canonical facade: client.PayoutMetric.list / client.PayoutMetric.load({ "id" => ... })
  def PayoutMetric(data = nil)
    require_relative 'entity/payout_metric_entity'
    PayoutMetricEntity.new(self, data)
  end


  # Canonical facade: client.Payrun.list / client.Payrun.load({ "id" => ... })
  def Payrun(data = nil)
    require_relative 'entity/payrun_entity'
    PayrunEntity.new(self, data)
  end


  # Canonical facade: client.Report.list / client.Report.load({ "id" => ... })
  def Report(data = nil)
    require_relative 'entity/report_entity'
    ReportEntity.new(self, data)
  end


  # Canonical facade: client.ReportResult.list / client.ReportResult.load({ "id" => ... })
  def ReportResult(data = nil)
    require_relative 'entity/report_result_entity'
    ReportResultEntity.new(self, data)
  end


  # Canonical facade: client.Role.list / client.Role.load({ "id" => ... })
  def Role(data = nil)
    require_relative 'entity/role_entity'
    RoleEntity.new(self, data)
  end


  # Canonical facade: client.Rule.list / client.Rule.load({ "id" => ... })
  def Rule(data = nil)
    require_relative 'entity/rule_entity'
    RuleEntity.new(self, data)
  end


  # Canonical facade: client.RuleEvent.list / client.RuleEvent.load({ "id" => ... })
  def RuleEvent(data = nil)
    require_relative 'entity/rule_event_entity'
    RuleEventEntity.new(self, data)
  end


  # Canonical facade: client.Tag.list / client.Tag.load({ "id" => ... })
  def Tag(data = nil)
    require_relative 'entity/tag_entity'
    TagEntity.new(self, data)
  end


  # Canonical facade: client.Token.list / client.Token.load({ "id" => ... })
  def Token(data = nil)
    require_relative 'entity/token_entity'
    TokenEntity.new(self, data)
  end


  # Canonical facade: client.Transaction.list / client.Transaction.load({ "id" => ... })
  def Transaction(data = nil)
    require_relative 'entity/transaction_entity'
    TransactionEntity.new(self, data)
  end


  # Canonical facade: client.User.list / client.User.load({ "id" => ... })
  def User(data = nil)
    require_relative 'entity/user_entity'
    UserEntity.new(self, data)
  end


  # Canonical facade: client.UserInvite.list / client.UserInvite.load({ "id" => ... })
  def UserInvite(data = nil)
    require_relative 'entity/user_invite_entity'
    UserInviteEntity.new(self, data)
  end


  # Canonical facade: client.Virtual.list / client.Virtual.load({ "id" => ... })
  def Virtual(data = nil)
    require_relative 'entity/virtual_entity'
    VirtualEntity.new(self, data)
  end


  # Canonical facade: client.Webhook.list / client.Webhook.load({ "id" => ... })
  def Webhook(data = nil)
    require_relative 'entity/webhook_entity'
    WebhookEntity.new(self, data)
  end



  def self.test(testopts = nil, sdkopts = nil)
    sdkopts = sdkopts || {}
    sdkopts = VoxgigStruct.clone(sdkopts)
    sdkopts = {} unless sdkopts.is_a?(Hash)

    testopts = testopts || {}
    testopts = VoxgigStruct.clone(testopts)
    testopts = {} unless testopts.is_a?(Hash)
    testopts["active"] = true

    VoxgigStruct.setpath(sdkopts, "feature.test", testopts)

    sdk = NofrixionSDK.new(sdkopts)
    sdk.mode = "test"
    sdk
  end
end
