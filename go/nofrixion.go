package voxgignofrixionsdk

import (
	"github.com/voxgig-sdk/nofrixion-sdk/go/core"
	"github.com/voxgig-sdk/nofrixion-sdk/go/entity"
	"github.com/voxgig-sdk/nofrixion-sdk/go/feature"
	_ "github.com/voxgig-sdk/nofrixion-sdk/go/utility"
)

// Type aliases preserve external API.
type NofrixionSDK = core.NofrixionSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type NofrixionEntity = core.NofrixionEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type NofrixionError = core.NofrixionError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewAccountEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewAccountEntity(client, entopts)
	}
	core.NewBeneficiaryEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewBeneficiaryEntity(client, entopts)
	}
	core.NewCancelEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewCancelEntity(client, entopts)
	}
	core.NewDisableEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewDisableEntity(client, entopts)
	}
	core.NewEnableEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewEnableEntity(client, entopts)
	}
	core.NewMerchantEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMerchantEntity(client, entopts)
	}
	core.NewMetadataEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMetadataEntity(client, entopts)
	}
	core.NewNoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntity(client, entopts)
	}
	core.NewNoFrixionBizBizModelsPagingPaymentRequestPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionBizBizModelsPagingPaymentRequestPageEntity(client, entopts)
	}
	core.NewNoFrixionBizBizModelsPagingPayoutPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionBizBizModelsPagingPayoutPageEntity(client, entopts)
	}
	core.NewNoFrixionBizBizModelsPagingPayrunPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionBizBizModelsPagingPayrunPageEntity(client, entopts)
	}
	core.NewNoFrixionBizBizModelsPagingRuleEventsPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionBizBizModelsPagingRuleEventsPageEntity(client, entopts)
	}
	core.NewNoFrixionBizBizModelsPagingRulesPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionBizBizModelsPagingRulesPageEntity(client, entopts)
	}
	core.NewNoFrixionBizBizModelsPaymentsCardPaymentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionBizBizModelsPaymentsCardPaymentEntity(client, entopts)
	}
	core.NewNoFrixionBizBizModelsPaymentsCardPublicKeyEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionBizBizModelsPaymentsCardPublicKeyEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsBatchPayoutEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsBatchPayoutEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsBeneficiaryGroupPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsBeneficiaryGroupPageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsBeneficiaryPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsBeneficiaryPageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsCardCustomerTokenEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsCardCustomerTokenEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsFxRateEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsFxRateEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsIPaymentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsIPaymentEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsMandatesMandateEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsMandatesMandateEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsMerchantEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsMerchantEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsMerchantPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsMerchantPageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsMerchantPayByBankSettingEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsMerchantPayByBankSettingEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsMerchantTokenEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsMerchantTokenEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsMerchantTokenPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsMerchantTokenPageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsNoFrixionVersionEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsNoFrixionVersionEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsOpenBankingAccountEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsOpenBankingAccountEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsOpenBankingConsentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsOpenBankingConsentEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsOpenBankingTransactionEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsOpenBankingTransactionEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentAccountPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentAccountPageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentInitiationEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentInitiationEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentRequestEventEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentRequestEventEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentRequestMetricEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentRequestMetricEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentRequestMinimalEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentRequestMinimalEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentRequestResultEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentRequestResultEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2EntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Entity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3EntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Entity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4EntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4Entity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPayoutKeysetPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPayoutKeysetPageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPayoutMetricEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPayoutMetricEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsPayrunEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsPayrunEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsReportResultEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsReportResultEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsRuleEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsRuleEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsTransactionEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsTransactionEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsTransactionPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsTransactionPageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsUserInviteEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsUserInviteEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsUserInvitePageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsUserInvitePageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsUserPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsUserPageEntity(client, entopts)
	}
	core.NewNoFrixionMoneyMoovModelsWebhookEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionMoneyMoovModelsWebhookEntity(client, entopts)
	}
	core.NewOpenBankingEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewOpenBankingEntity(client, entopts)
	}
	core.NewPayeeverificationEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPayeeverificationEntity(client, entopts)
	}
	core.NewPaymentRequestEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentRequestEntity(client, entopts)
	}
	core.NewPayoutEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPayoutEntity(client, entopts)
	}
	core.NewPayrunEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPayrunEntity(client, entopts)
	}
	core.NewRejectEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewRejectEntity(client, entopts)
	}
	core.NewReportEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewReportEntity(client, entopts)
	}
	core.NewRuleEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewRuleEntity(client, entopts)
	}
	core.NewSendEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewSendEntity(client, entopts)
	}
	core.NewSendbeneficiaryEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewSendbeneficiaryEntity(client, entopts)
	}
	core.NewTagEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewTagEntity(client, entopts)
	}
	core.NewTokenEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewTokenEntity(client, entopts)
	}
	core.NewTransactionEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewTransactionEntity(client, entopts)
	}
	core.NewUserEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewUserEntity(client, entopts)
	}
	core.NewUserInviteEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewUserInviteEntity(client, entopts)
	}
	core.NewVirtualEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewVirtualEntity(client, entopts)
	}
	core.NewWebhookEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewWebhookEntity(client, entopts)
	}
	core.NewWhoamiEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewWhoamiEntity(client, entopts)
	}
	core.NewWhoamitrustedappEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewWhoamitrustedappEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewNofrixionSDK = core.NewNofrixionSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewNofrixionSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *NofrixionSDK  { return NewNofrixionSDK(nil) }
func Test() *NofrixionSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
