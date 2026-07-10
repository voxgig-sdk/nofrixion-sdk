package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewAccountEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewBatchEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewBeneficiariesCreateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewBeneficiaryEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewBeneficiaryGroupEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewCardEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewCardCustomerTokenEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewCardPaymentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewCardPublicKeyEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewConsentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewCurrencyEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewDirectDebitBatchSubmitEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewFxRateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewIPaymentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMandateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMerchantEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMerchantAuthorisationSettingEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMerchantDirectDebitMandateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMerchantPayByBankSettingEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMerchantPaymentRequestTemplateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMerchantTokenEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewMetadataEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewNoFrixionVersionEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewOpenBankingEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPayeeverificationEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentAccountEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentAccountMinimalEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentInitiationEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentRequestEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentRequestEventEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentRequestMetricEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentRequestMinimalEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentRequestResultEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPaymentRequestsCreateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPayoutEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPayoutKeysetEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPayoutMetricEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPayoutsCreateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewPayrunEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewReportEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewReportResultEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewRolesCreateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewRuleEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewRuleEventEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewTagEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewTokenEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewTransactionEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewUserEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewUserInviteEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewUserInvitesCreateEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewVirtualEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

var NewWebhookEntityFunc func(client *NofrixionSDK, entopts map[string]any) NofrixionEntity

