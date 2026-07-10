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
	core.NewBatchEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewBatchEntity(client, entopts)
	}
	core.NewBeneficiaryEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewBeneficiaryEntity(client, entopts)
	}
	core.NewBeneficiaryGroupEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewBeneficiaryGroupEntity(client, entopts)
	}
	core.NewCardEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewCardEntity(client, entopts)
	}
	core.NewCardCustomerTokenEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewCardCustomerTokenEntity(client, entopts)
	}
	core.NewCardPaymentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewCardPaymentEntity(client, entopts)
	}
	core.NewCardPublicKeyEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewCardPublicKeyEntity(client, entopts)
	}
	core.NewConsentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewConsentEntity(client, entopts)
	}
	core.NewCurrencyEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewCurrencyEntity(client, entopts)
	}
	core.NewDirectDebitBatchSubmitEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewDirectDebitBatchSubmitEntity(client, entopts)
	}
	core.NewFxRateEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewFxRateEntity(client, entopts)
	}
	core.NewIPaymentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewIPaymentEntity(client, entopts)
	}
	core.NewMandateEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMandateEntity(client, entopts)
	}
	core.NewMerchantEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMerchantEntity(client, entopts)
	}
	core.NewMerchantAuthorisationSettingEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMerchantAuthorisationSettingEntity(client, entopts)
	}
	core.NewMerchantDirectDebitMandatePageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMerchantDirectDebitMandatePageEntity(client, entopts)
	}
	core.NewMerchantPayByBankSettingEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMerchantPayByBankSettingEntity(client, entopts)
	}
	core.NewMerchantPaymentRequestTemplateEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMerchantPaymentRequestTemplateEntity(client, entopts)
	}
	core.NewMerchantTokenEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMerchantTokenEntity(client, entopts)
	}
	core.NewMetadataEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewMetadataEntity(client, entopts)
	}
	core.NewNoFrixionVersionEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewNoFrixionVersionEntity(client, entopts)
	}
	core.NewOpenBankingEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewOpenBankingEntity(client, entopts)
	}
	core.NewPayeeverificationEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPayeeverificationEntity(client, entopts)
	}
	core.NewPaymentEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentEntity(client, entopts)
	}
	core.NewPaymentAccountEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentAccountEntity(client, entopts)
	}
	core.NewPaymentAccountMinimalEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentAccountMinimalEntity(client, entopts)
	}
	core.NewPaymentInitiationEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentInitiationEntity(client, entopts)
	}
	core.NewPaymentRequestEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentRequestEntity(client, entopts)
	}
	core.NewPaymentRequestEventEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentRequestEventEntity(client, entopts)
	}
	core.NewPaymentRequestMetricEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentRequestMetricEntity(client, entopts)
	}
	core.NewPaymentRequestMinimalEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentRequestMinimalEntity(client, entopts)
	}
	core.NewPaymentRequestResultEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPaymentRequestResultEntity(client, entopts)
	}
	core.NewPayoutEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPayoutEntity(client, entopts)
	}
	core.NewPayoutKeysetPageEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPayoutKeysetPageEntity(client, entopts)
	}
	core.NewPayoutMetricEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPayoutMetricEntity(client, entopts)
	}
	core.NewPayrunEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewPayrunEntity(client, entopts)
	}
	core.NewReportEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewReportEntity(client, entopts)
	}
	core.NewReportResultEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewReportResultEntity(client, entopts)
	}
	core.NewRoleEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewRoleEntity(client, entopts)
	}
	core.NewRuleEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewRuleEntity(client, entopts)
	}
	core.NewRuleEventEntityFunc = func(client *core.NofrixionSDK, entopts map[string]any) core.NofrixionEntity {
		return entity.NewRuleEventEntity(client, entopts)
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
