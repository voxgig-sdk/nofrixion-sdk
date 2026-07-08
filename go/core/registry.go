package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewAccountEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewBeneficiaryEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewCancelEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewDisableEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewEnableEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMerchantEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMetadataEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionBizBizModelsPagingPaymentRequestPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionBizBizModelsPagingPayoutPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionBizBizModelsPagingPayrunPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionBizBizModelsPagingRuleEventsPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionBizBizModelsPagingRulesPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionBizBizModelsPaymentsCardPaymentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionBizBizModelsPaymentsCardPublicKeyEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsBatchPayoutEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsBeneficiaryGroupPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsBeneficiaryPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsCardCustomerTokenEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsFxRateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsIPaymentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsMandatesMandateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsMerchantEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsMerchantPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsMerchantPayByBankSettingEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsMerchantTokenEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsMerchantTokenPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsNoFrixionVersionEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsOpenBankingAccountEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsOpenBankingConsentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsOpenBankingTransactionEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentAccountPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentInitiationEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentRequestEventEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentRequestMetricEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentRequestMinimalEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentRequestResultEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2EntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3EntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4EntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPayoutKeysetPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPayoutMetricEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsPayrunEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsReportResultEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsRuleEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsTransactionEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsTransactionPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsUserInviteEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsUserInvitePageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsUserPageEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionMoneyMoovModelsWebhookEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewOpenBankingEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPayeeverificationEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentRequestEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPayoutEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPayrunEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewRejectEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewReportEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewRuleEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewSendEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewSendbeneficiaryEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewTagEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewTokenEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewTransactionEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewUserEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewUserInviteEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewVirtualEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewWebhookEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewWhoamiEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewWhoamitrustedappEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

