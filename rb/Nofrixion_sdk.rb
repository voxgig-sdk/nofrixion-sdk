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


  # Canonical facade: client.Beneficiary.list / client.Beneficiary.load({ "id" => ... })
  def Beneficiary(data = nil)
    require_relative 'entity/beneficiary_entity'
    BeneficiaryEntity.new(self, data)
  end


  # Canonical facade: client.Cancel.list / client.Cancel.load({ "id" => ... })
  def Cancel(data = nil)
    require_relative 'entity/cancel_entity'
    CancelEntity.new(self, data)
  end


  # Canonical facade: client.Disable.list / client.Disable.load({ "id" => ... })
  def Disable(data = nil)
    require_relative 'entity/disable_entity'
    DisableEntity.new(self, data)
  end


  # Canonical facade: client.Enable.list / client.Enable.load({ "id" => ... })
  def Enable(data = nil)
    require_relative 'entity/enable_entity'
    EnableEntity.new(self, data)
  end


  # Canonical facade: client.Merchant.list / client.Merchant.load({ "id" => ... })
  def Merchant(data = nil)
    require_relative 'entity/merchant_entity'
    MerchantEntity.new(self, data)
  end


  # Canonical facade: client.Metadata.list / client.Metadata.load({ "id" => ... })
  def Metadata(data = nil)
    require_relative 'entity/metadata_entity'
    MetadataEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage.list / client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage.load({ "id" => ... })
  def NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(data = nil)
    require_relative 'entity/no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page_entity'
    NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionBizBizModelsPagingPaymentRequestPage.list / client.NoFrixionBizBizModelsPagingPaymentRequestPage.load({ "id" => ... })
  def NoFrixionBizBizModelsPagingPaymentRequestPage(data = nil)
    require_relative 'entity/no_frixion_biz_biz_models_paging_payment_request_page_entity'
    NoFrixionBizBizModelsPagingPaymentRequestPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionBizBizModelsPagingPayoutPage.list / client.NoFrixionBizBizModelsPagingPayoutPage.load({ "id" => ... })
  def NoFrixionBizBizModelsPagingPayoutPage(data = nil)
    require_relative 'entity/no_frixion_biz_biz_models_paging_payout_page_entity'
    NoFrixionBizBizModelsPagingPayoutPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionBizBizModelsPagingPayrunPage.list / client.NoFrixionBizBizModelsPagingPayrunPage.load({ "id" => ... })
  def NoFrixionBizBizModelsPagingPayrunPage(data = nil)
    require_relative 'entity/no_frixion_biz_biz_models_paging_payrun_page_entity'
    NoFrixionBizBizModelsPagingPayrunPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionBizBizModelsPagingRuleEventsPage.list / client.NoFrixionBizBizModelsPagingRuleEventsPage.load({ "id" => ... })
  def NoFrixionBizBizModelsPagingRuleEventsPage(data = nil)
    require_relative 'entity/no_frixion_biz_biz_models_paging_rule_events_page_entity'
    NoFrixionBizBizModelsPagingRuleEventsPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionBizBizModelsPagingRulesPage.list / client.NoFrixionBizBizModelsPagingRulesPage.load({ "id" => ... })
  def NoFrixionBizBizModelsPagingRulesPage(data = nil)
    require_relative 'entity/no_frixion_biz_biz_models_paging_rules_page_entity'
    NoFrixionBizBizModelsPagingRulesPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionBizBizModelsPaymentsCardPayment.list / client.NoFrixionBizBizModelsPaymentsCardPayment.load({ "id" => ... })
  def NoFrixionBizBizModelsPaymentsCardPayment(data = nil)
    require_relative 'entity/no_frixion_biz_biz_models_payments_card_payment_entity'
    NoFrixionBizBizModelsPaymentsCardPaymentEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionBizBizModelsPaymentsCardPublicKey.list / client.NoFrixionBizBizModelsPaymentsCardPublicKey.load({ "id" => ... })
  def NoFrixionBizBizModelsPaymentsCardPublicKey(data = nil)
    require_relative 'entity/no_frixion_biz_biz_models_payments_card_public_key_entity'
    NoFrixionBizBizModelsPaymentsCardPublicKeyEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries.list / client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries.load({ "id" => ... })
  def NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(data = nil)
    require_relative 'entity/no_frixion_money_moov_api_features_beneficiaries_beneficiaries_entity'
    NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment.list / client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment.load({ "id" => ... })
  def NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(data = nil)
    require_relative 'entity/no_frixion_money_moov_api_features_payment_requests_payment_entity'
    NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate.list / client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate.load({ "id" => ... })
  def NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(data = nil)
    require_relative 'entity/no_frixion_money_moov_api_features_permissions_roles_create_entity'
    NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate.list / client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate.load({ "id" => ... })
  def NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(data = nil)
    require_relative 'entity/no_frixion_money_moov_api_features_user_invites_create_entity'
    NoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant.list / client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_authorisation_settings_merchant_entity'
    NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsBatchPayout.list / client.NoFrixionMoneyMoovModelsBatchPayout.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsBatchPayout(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_batch_payout_entity'
    NoFrixionMoneyMoovModelsBatchPayoutEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage.list / client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsBeneficiaryGroupPage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_beneficiary_group_page_entity'
    NoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsBeneficiaryPage.list / client.NoFrixionMoneyMoovModelsBeneficiaryPage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsBeneficiaryPage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_beneficiary_page_entity'
    NoFrixionMoneyMoovModelsBeneficiaryPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsCardCustomerToken.list / client.NoFrixionMoneyMoovModelsCardCustomerToken.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsCardCustomerToken(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_card_customer_token_entity'
    NoFrixionMoneyMoovModelsCardCustomerTokenEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo.list / client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_currency_currency_info_entity'
    NoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit.list / client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_direct_debit_batch_submit_entity'
    NoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsFxRate.list / client.NoFrixionMoneyMoovModelsFxRate.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsFxRate(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_fx_rate_entity'
    NoFrixionMoneyMoovModelsFxRateEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsIPayment.list / client.NoFrixionMoneyMoovModelsIPayment.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsIPayment(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_i_payment_entity'
    NoFrixionMoneyMoovModelsIPaymentEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsMandatesMandate.list / client.NoFrixionMoneyMoovModelsMandatesMandate.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsMandatesMandate(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_mandates_mandate_entity'
    NoFrixionMoneyMoovModelsMandatesMandateEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsMerchant.list / client.NoFrixionMoneyMoovModelsMerchant.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsMerchant(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_merchant_entity'
    NoFrixionMoneyMoovModelsMerchantEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsMerchantPage.list / client.NoFrixionMoneyMoovModelsMerchantPage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsMerchantPage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_merchant_page_entity'
    NoFrixionMoneyMoovModelsMerchantPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting.list / client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsMerchantPayByBankSetting(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_merchant_pay_by_bank_setting_entity'
    NoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsMerchantToken.list / client.NoFrixionMoneyMoovModelsMerchantToken.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsMerchantToken(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_merchant_token_entity'
    NoFrixionMoneyMoovModelsMerchantTokenEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsMerchantTokenPage.list / client.NoFrixionMoneyMoovModelsMerchantTokenPage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsMerchantTokenPage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_merchant_token_page_entity'
    NoFrixionMoneyMoovModelsMerchantTokenPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsNoFrixionVersion.list / client.NoFrixionMoneyMoovModelsNoFrixionVersion.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsNoFrixionVersion(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_no_frixion_version_entity'
    NoFrixionMoneyMoovModelsNoFrixionVersionEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsOpenBankingAccount.list / client.NoFrixionMoneyMoovModelsOpenBankingAccount.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsOpenBankingAccount(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_open_banking_account_entity'
    NoFrixionMoneyMoovModelsOpenBankingAccountEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsOpenBankingConsent.list / client.NoFrixionMoneyMoovModelsOpenBankingConsent.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsOpenBankingConsent(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_open_banking_consent_entity'
    NoFrixionMoneyMoovModelsOpenBankingConsentEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsOpenBankingTransaction.list / client.NoFrixionMoneyMoovModelsOpenBankingTransaction.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsOpenBankingTransaction(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_open_banking_transaction_entity'
    NoFrixionMoneyMoovModelsOpenBankingTransactionEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPayment.list / client.NoFrixionMoneyMoovModelsPayment.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPayment(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_entity'
    NoFrixionMoneyMoovModelsPaymentEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage.list / client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_account_minimal_page_entity'
    NoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentAccountPage.list / client.NoFrixionMoneyMoovModelsPaymentAccountPage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentAccountPage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_account_page_entity'
    NoFrixionMoneyMoovModelsPaymentAccountPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentInitiation.list / client.NoFrixionMoneyMoovModelsPaymentInitiation.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentInitiation(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_initiation_entity'
    NoFrixionMoneyMoovModelsPaymentInitiationEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentRequestEvent.list / client.NoFrixionMoneyMoovModelsPaymentRequestEvent.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentRequestEvent(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_request_event_entity'
    NoFrixionMoneyMoovModelsPaymentRequestEventEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentRequestMetric.list / client.NoFrixionMoneyMoovModelsPaymentRequestMetric.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentRequestMetric(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_request_metric_entity'
    NoFrixionMoneyMoovModelsPaymentRequestMetricEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentRequestMinimal.list / client.NoFrixionMoneyMoovModelsPaymentRequestMinimal.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentRequestMinimal(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_request_minimal_entity'
    NoFrixionMoneyMoovModelsPaymentRequestMinimalEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentRequestResult.list / client.NoFrixionMoneyMoovModelsPaymentRequestResult.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentRequestResult(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_request_result_entity'
    NoFrixionMoneyMoovModelsPaymentRequestResultEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment.list / client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_requests_merchant_payment_entity'
    NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2.list / client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_requests_merchant_payment2_entity'
    NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3.list / client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_requests_merchant_payment3_entity'
    NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4.list / client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payment_requests_merchant_payment4_entity'
    NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPayoutKeysetPage.list / client.NoFrixionMoneyMoovModelsPayoutKeysetPage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPayoutKeysetPage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payout_keyset_page_entity'
    NoFrixionMoneyMoovModelsPayoutKeysetPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPayoutMetric.list / client.NoFrixionMoneyMoovModelsPayoutMetric.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPayoutMetric(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payout_metric_entity'
    NoFrixionMoneyMoovModelsPayoutMetricEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate.list / client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payouts_payouts_create_entity'
    NoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsPayrun.list / client.NoFrixionMoneyMoovModelsPayrun.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsPayrun(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_payrun_entity'
    NoFrixionMoneyMoovModelsPayrunEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsReportResult.list / client.NoFrixionMoneyMoovModelsReportResult.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsReportResult(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_report_result_entity'
    NoFrixionMoneyMoovModelsReportResultEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsRule.list / client.NoFrixionMoneyMoovModelsRule.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsRule(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_rule_entity'
    NoFrixionMoneyMoovModelsRuleEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsTransaction.list / client.NoFrixionMoneyMoovModelsTransaction.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsTransaction(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_transaction_entity'
    NoFrixionMoneyMoovModelsTransactionEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsTransactionPage.list / client.NoFrixionMoneyMoovModelsTransactionPage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsTransactionPage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_transaction_page_entity'
    NoFrixionMoneyMoovModelsTransactionPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsUserInvite.list / client.NoFrixionMoneyMoovModelsUserInvite.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsUserInvite(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_user_invite_entity'
    NoFrixionMoneyMoovModelsUserInviteEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsUserInvitePage.list / client.NoFrixionMoneyMoovModelsUserInvitePage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsUserInvitePage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_user_invite_page_entity'
    NoFrixionMoneyMoovModelsUserInvitePageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsUserPage.list / client.NoFrixionMoneyMoovModelsUserPage.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsUserPage(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_user_page_entity'
    NoFrixionMoneyMoovModelsUserPageEntity.new(self, data)
  end


  # Canonical facade: client.NoFrixionMoneyMoovModelsWebhook.list / client.NoFrixionMoneyMoovModelsWebhook.load({ "id" => ... })
  def NoFrixionMoneyMoovModelsWebhook(data = nil)
    require_relative 'entity/no_frixion_money_moov_models_webhook_entity'
    NoFrixionMoneyMoovModelsWebhookEntity.new(self, data)
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


  # Canonical facade: client.PaymentRequest.list / client.PaymentRequest.load({ "id" => ... })
  def PaymentRequest(data = nil)
    require_relative 'entity/payment_request_entity'
    PaymentRequestEntity.new(self, data)
  end


  # Canonical facade: client.Payout.list / client.Payout.load({ "id" => ... })
  def Payout(data = nil)
    require_relative 'entity/payout_entity'
    PayoutEntity.new(self, data)
  end


  # Canonical facade: client.Payrun.list / client.Payrun.load({ "id" => ... })
  def Payrun(data = nil)
    require_relative 'entity/payrun_entity'
    PayrunEntity.new(self, data)
  end


  # Canonical facade: client.Reject.list / client.Reject.load({ "id" => ... })
  def Reject(data = nil)
    require_relative 'entity/reject_entity'
    RejectEntity.new(self, data)
  end


  # Canonical facade: client.Report.list / client.Report.load({ "id" => ... })
  def Report(data = nil)
    require_relative 'entity/report_entity'
    ReportEntity.new(self, data)
  end


  # Canonical facade: client.Rule.list / client.Rule.load({ "id" => ... })
  def Rule(data = nil)
    require_relative 'entity/rule_entity'
    RuleEntity.new(self, data)
  end


  # Canonical facade: client.Send.list / client.Send.load({ "id" => ... })
  def Send(data = nil)
    require_relative 'entity/send_entity'
    SendEntity.new(self, data)
  end


  # Canonical facade: client.Sendbeneficiary.list / client.Sendbeneficiary.load({ "id" => ... })
  def Sendbeneficiary(data = nil)
    require_relative 'entity/sendbeneficiary_entity'
    SendbeneficiaryEntity.new(self, data)
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


  # Canonical facade: client.Whoami.list / client.Whoami.load({ "id" => ... })
  def Whoami(data = nil)
    require_relative 'entity/whoami_entity'
    WhoamiEntity.new(self, data)
  end


  # Canonical facade: client.Whoamitrustedapp.list / client.Whoamitrustedapp.load({ "id" => ... })
  def Whoamitrustedapp(data = nil)
    require_relative 'entity/whoamitrustedapp_entity'
    WhoamitrustedappEntity.new(self, data)
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
