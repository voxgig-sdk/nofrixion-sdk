package core

import (
	"fmt"

	vs "github.com/voxgig-sdk/nofrixion-sdk/go/utility/struct"
)

type NofrixionSDK struct {
	Mode     string
	options  map[string]any
	utility  *Utility
	Features []Feature
	rootctx  *Context
}

func NewNofrixionSDK(options map[string]any) *NofrixionSDK {
	sdk := &NofrixionSDK{
		Mode:     "live",
		Features: []Feature{},
	}

	sdk.utility = NewUtility()

	config := MakeConfig()

	sdk.rootctx = sdk.utility.MakeContext(map[string]any{
		"client":  sdk,
		"utility": sdk.utility,
		"config":  config,
		"options": options,
		"shared":  map[string]any{},
	}, nil)

	sdk.options = sdk.utility.MakeOptions(sdk.rootctx)

	if vs.GetPath([]any{"feature", "test", "active"}, sdk.options) == true {
		sdk.Mode = "test"
	}

	sdk.rootctx.Options = sdk.options

	// Add features from config.
	featureOpts := ToMapAny(vs.GetProp(sdk.options, "feature"))
	if featureOpts != nil {
		for _, item := range vs.Items(featureOpts) {
			fname, _ := item[0].(string)
			fopts := ToMapAny(item[1])
			if fopts != nil {
				if active, ok := fopts["active"]; ok {
					if ab, ok := active.(bool); ok && ab {
						sdk.utility.FeatureAdd(sdk.rootctx, makeFeature(fname))
					}
				}
			}
		}
	}

	// Add extension features.
	if extend := vs.GetProp(sdk.options, "extend"); extend != nil {
		if extList, ok := extend.([]any); ok {
			for _, f := range extList {
				if feat, ok := f.(Feature); ok {
					sdk.utility.FeatureAdd(sdk.rootctx, feat)
				}
			}
		}
	}

	// Initialize features.
	for _, f := range sdk.Features {
		sdk.utility.FeatureInit(sdk.rootctx, f)
	}

	sdk.utility.FeatureHook(sdk.rootctx, "PostConstruct")

	return sdk
}

func (sdk *NofrixionSDK) OptionsMap() map[string]any {
	out := vs.Clone(sdk.options)
	if om, ok := out.(map[string]any); ok {
		return om
	}
	return map[string]any{}
}

func (sdk *NofrixionSDK) GetUtility() *Utility {
	return CopyUtility(sdk.utility)
}

func (sdk *NofrixionSDK) GetRootCtx() *Context {
	return sdk.rootctx
}

func (sdk *NofrixionSDK) Prepare(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "prepare",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	options := sdk.options

	path, _ := vs.GetProp(fetchargs, "path").(string)
	method, _ := vs.GetProp(fetchargs, "method").(string)
	if method == "" {
		method = "GET"
	}

	params := ToMapAny(vs.GetProp(fetchargs, "params"))
	if params == nil {
		params = map[string]any{}
	}
	query := ToMapAny(vs.GetProp(fetchargs, "query"))
	if query == nil {
		query = map[string]any{}
	}

	headers := utility.PrepareHeaders(ctx)

	base, _ := vs.GetProp(options, "base").(string)
	prefix, _ := vs.GetProp(options, "prefix").(string)
	suffix, _ := vs.GetProp(options, "suffix").(string)

	ctx.Spec = NewSpec(map[string]any{
		"base":    base,
		"prefix":  prefix,
		"suffix":  suffix,
		"path":    path,
		"method":  method,
		"params":  params,
		"query":   query,
		"headers": headers,
		"body":    vs.GetProp(fetchargs, "body"),
		"step":    "start",
	})

	// Merge user-provided headers.
	if uh := vs.GetProp(fetchargs, "headers"); uh != nil {
		if uhm, ok := uh.(map[string]any); ok {
			for k, v := range uhm {
				ctx.Spec.Headers[k] = v
			}
		}
	}

	_, err := utility.PrepareAuth(ctx)
	if err != nil {
		return nil, err
	}

	return utility.MakeFetchDef(ctx)
}

func (sdk *NofrixionSDK) Direct(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	fetchdef, err := sdk.Prepare(fetchargs)
	if err != nil {
		return map[string]any{"ok": false, "err": err}, nil
	}

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "direct",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	url, _ := fetchdef["url"].(string)
	fetched, fetchErr := utility.Fetcher(ctx, url, fetchdef)

	if fetchErr != nil {
		return map[string]any{"ok": false, "err": fetchErr}, nil
	}

	if fetched == nil {
		return map[string]any{
			"ok":  false,
			"err": ctx.MakeError("direct_no_response", "response: undefined"),
		}, nil
	}

	if fm, ok := fetched.(map[string]any); ok {
		status := ToInt(vs.GetProp(fm, "status"))
		headers := vs.GetProp(fm, "headers")

		// No-body responses (204, 304) and explicit zero content-length
		// must skip JSON parsing — calling json() on an empty body errors.
		var contentLength string
		if hm, ok := headers.(map[string]any); ok {
			if cl, ok := hm["content-length"]; ok {
				contentLength = fmt.Sprintf("%v", cl)
			}
		}
		noBody := status == 204 || status == 304 || contentLength == "0"

		var jsonData any
		if !noBody {
			if jf := vs.GetProp(fm, "json"); jf != nil {
				if f, ok := jf.(func() any); ok {
					// f() returns nil on parse error in our fetcher.
					jsonData = f()
				}
			}
		}

		return map[string]any{
			"ok":      status >= 200 && status < 300,
			"status":  status,
			"headers": headers,
			"data":    jsonData,
		}, nil
	}

	return map[string]any{"ok": false, "err": ctx.MakeError("direct_invalid", "invalid response type")}, nil
}


// Account returns a Account entity bound to this client.
// Idiomatic usage: client.Account(nil).List(nil, nil) or
// client.Account(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Account(data map[string]any) NofrixionEntity {
	return NewAccountEntityFunc(sdk, data)
}


// Beneficiary returns a Beneficiary entity bound to this client.
// Idiomatic usage: client.Beneficiary(nil).List(nil, nil) or
// client.Beneficiary(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Beneficiary(data map[string]any) NofrixionEntity {
	return NewBeneficiaryEntityFunc(sdk, data)
}


// Cancel returns a Cancel entity bound to this client.
// Idiomatic usage: client.Cancel(nil).List(nil, nil) or
// client.Cancel(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Cancel(data map[string]any) NofrixionEntity {
	return NewCancelEntityFunc(sdk, data)
}


// Disable returns a Disable entity bound to this client.
// Idiomatic usage: client.Disable(nil).List(nil, nil) or
// client.Disable(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Disable(data map[string]any) NofrixionEntity {
	return NewDisableEntityFunc(sdk, data)
}


// Enable returns a Enable entity bound to this client.
// Idiomatic usage: client.Enable(nil).List(nil, nil) or
// client.Enable(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Enable(data map[string]any) NofrixionEntity {
	return NewEnableEntityFunc(sdk, data)
}


// Merchant returns a Merchant entity bound to this client.
// Idiomatic usage: client.Merchant(nil).List(nil, nil) or
// client.Merchant(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Merchant(data map[string]any) NofrixionEntity {
	return NewMerchantEntityFunc(sdk, data)
}


// Metadata returns a Metadata entity bound to this client.
// Idiomatic usage: client.Metadata(nil).List(nil, nil) or
// client.Metadata(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Metadata(data map[string]any) NofrixionEntity {
	return NewMetadataEntityFunc(sdk, data)
}


// NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage returns a NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage entity bound to this client.
// Idiomatic usage: client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(nil).List(nil, nil) or
// client.NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(data map[string]any) NofrixionEntity {
	return NewNoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageEntityFunc(sdk, data)
}


// NoFrixionBizBizModelsPagingPaymentRequestPage returns a NoFrixionBizBizModelsPagingPaymentRequestPage entity bound to this client.
// Idiomatic usage: client.NoFrixionBizBizModelsPagingPaymentRequestPage(nil).List(nil, nil) or
// client.NoFrixionBizBizModelsPagingPaymentRequestPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionBizBizModelsPagingPaymentRequestPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionBizBizModelsPagingPaymentRequestPageEntityFunc(sdk, data)
}


// NoFrixionBizBizModelsPagingPayoutPage returns a NoFrixionBizBizModelsPagingPayoutPage entity bound to this client.
// Idiomatic usage: client.NoFrixionBizBizModelsPagingPayoutPage(nil).List(nil, nil) or
// client.NoFrixionBizBizModelsPagingPayoutPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionBizBizModelsPagingPayoutPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionBizBizModelsPagingPayoutPageEntityFunc(sdk, data)
}


// NoFrixionBizBizModelsPagingPayrunPage returns a NoFrixionBizBizModelsPagingPayrunPage entity bound to this client.
// Idiomatic usage: client.NoFrixionBizBizModelsPagingPayrunPage(nil).List(nil, nil) or
// client.NoFrixionBizBizModelsPagingPayrunPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionBizBizModelsPagingPayrunPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionBizBizModelsPagingPayrunPageEntityFunc(sdk, data)
}


// NoFrixionBizBizModelsPagingRuleEventsPage returns a NoFrixionBizBizModelsPagingRuleEventsPage entity bound to this client.
// Idiomatic usage: client.NoFrixionBizBizModelsPagingRuleEventsPage(nil).List(nil, nil) or
// client.NoFrixionBizBizModelsPagingRuleEventsPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionBizBizModelsPagingRuleEventsPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionBizBizModelsPagingRuleEventsPageEntityFunc(sdk, data)
}


// NoFrixionBizBizModelsPagingRulesPage returns a NoFrixionBizBizModelsPagingRulesPage entity bound to this client.
// Idiomatic usage: client.NoFrixionBizBizModelsPagingRulesPage(nil).List(nil, nil) or
// client.NoFrixionBizBizModelsPagingRulesPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionBizBizModelsPagingRulesPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionBizBizModelsPagingRulesPageEntityFunc(sdk, data)
}


// NoFrixionBizBizModelsPaymentsCardPayment returns a NoFrixionBizBizModelsPaymentsCardPayment entity bound to this client.
// Idiomatic usage: client.NoFrixionBizBizModelsPaymentsCardPayment(nil).List(nil, nil) or
// client.NoFrixionBizBizModelsPaymentsCardPayment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionBizBizModelsPaymentsCardPayment(data map[string]any) NofrixionEntity {
	return NewNoFrixionBizBizModelsPaymentsCardPaymentEntityFunc(sdk, data)
}


// NoFrixionBizBizModelsPaymentsCardPublicKey returns a NoFrixionBizBizModelsPaymentsCardPublicKey entity bound to this client.
// Idiomatic usage: client.NoFrixionBizBizModelsPaymentsCardPublicKey(nil).List(nil, nil) or
// client.NoFrixionBizBizModelsPaymentsCardPublicKey(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionBizBizModelsPaymentsCardPublicKey(data map[string]any) NofrixionEntity {
	return NewNoFrixionBizBizModelsPaymentsCardPublicKeyEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries returns a NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment returns a NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate returns a NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovApiFeaturesUserInvitesCreate returns a NoFrixionMoneyMoovApiFeaturesUserInvitesCreate entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovApiFeaturesUserInvitesCreateEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant returns a NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsAuthorisationSettingsMerchantEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsBatchPayout returns a NoFrixionMoneyMoovModelsBatchPayout entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsBatchPayout(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsBatchPayout(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsBatchPayout(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsBatchPayoutEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsBeneficiaryGroupPage returns a NoFrixionMoneyMoovModelsBeneficiaryGroupPage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsBeneficiaryGroupPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsBeneficiaryGroupPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsBeneficiaryGroupPageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsBeneficiaryPage returns a NoFrixionMoneyMoovModelsBeneficiaryPage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsBeneficiaryPage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsBeneficiaryPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsBeneficiaryPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsBeneficiaryPageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsCardCustomerToken returns a NoFrixionMoneyMoovModelsCardCustomerToken entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsCardCustomerToken(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsCardCustomerToken(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsCardCustomerToken(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsCardCustomerTokenEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsCurrencyCurrencyInfo returns a NoFrixionMoneyMoovModelsCurrencyCurrencyInfo entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsCurrencyCurrencyInfoEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsDirectDebitBatchSubmit returns a NoFrixionMoneyMoovModelsDirectDebitBatchSubmit entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsDirectDebitBatchSubmitEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsFxRate returns a NoFrixionMoneyMoovModelsFxRate entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsFxRate(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsFxRate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsFxRate(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsFxRateEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsIPayment returns a NoFrixionMoneyMoovModelsIPayment entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsIPayment(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsIPayment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsIPayment(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsIPaymentEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsMandatesMandate returns a NoFrixionMoneyMoovModelsMandatesMandate entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsMandatesMandate(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsMandatesMandate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsMandatesMandate(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsMandatesMandateEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsMerchant returns a NoFrixionMoneyMoovModelsMerchant entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsMerchant(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsMerchant(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsMerchant(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsMerchantEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsMerchantPage returns a NoFrixionMoneyMoovModelsMerchantPage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsMerchantPage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsMerchantPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsMerchantPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsMerchantPageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsMerchantPayByBankSetting returns a NoFrixionMoneyMoovModelsMerchantPayByBankSetting entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsMerchantPayByBankSetting(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsMerchantPayByBankSetting(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsMerchantPayByBankSettingEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsMerchantToken returns a NoFrixionMoneyMoovModelsMerchantToken entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsMerchantToken(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsMerchantToken(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsMerchantToken(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsMerchantTokenEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsMerchantTokenPage returns a NoFrixionMoneyMoovModelsMerchantTokenPage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsMerchantTokenPage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsMerchantTokenPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsMerchantTokenPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsMerchantTokenPageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsNoFrixionVersion returns a NoFrixionMoneyMoovModelsNoFrixionVersion entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsNoFrixionVersion(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsNoFrixionVersion(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsNoFrixionVersion(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsNoFrixionVersionEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsOpenBankingAccount returns a NoFrixionMoneyMoovModelsOpenBankingAccount entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsOpenBankingAccount(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsOpenBankingAccount(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsOpenBankingAccount(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsOpenBankingAccountEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsOpenBankingConsent returns a NoFrixionMoneyMoovModelsOpenBankingConsent entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsOpenBankingConsent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsOpenBankingConsent(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsOpenBankingConsentEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsOpenBankingTransaction returns a NoFrixionMoneyMoovModelsOpenBankingTransaction entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsOpenBankingTransaction(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsOpenBankingTransaction(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsOpenBankingTransaction(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsOpenBankingTransactionEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPayment returns a NoFrixionMoneyMoovModelsPayment entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPayment(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPayment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPayment(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentAccountMinimalPage returns a NoFrixionMoneyMoovModelsPaymentAccountMinimalPage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentAccountMinimalPageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentAccountPage returns a NoFrixionMoneyMoovModelsPaymentAccountPage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentAccountPage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentAccountPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentAccountPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentAccountPageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentInitiation returns a NoFrixionMoneyMoovModelsPaymentInitiation entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentInitiation(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentInitiation(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentInitiation(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentInitiationEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentRequestEvent returns a NoFrixionMoneyMoovModelsPaymentRequestEvent entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentRequestEvent(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentRequestEvent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentRequestEvent(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentRequestEventEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentRequestMetric returns a NoFrixionMoneyMoovModelsPaymentRequestMetric entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentRequestMetric(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentRequestMetric(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentRequestMetric(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentRequestMetricEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentRequestMinimal returns a NoFrixionMoneyMoovModelsPaymentRequestMinimal entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentRequestMinimal(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentRequestMinimal(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentRequestMinimal(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentRequestMinimalEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentRequestResult returns a NoFrixionMoneyMoovModelsPaymentRequestResult entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentRequestResult(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentRequestResult(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentRequestResult(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentRequestResultEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment returns a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 returns a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2EntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 returns a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3EntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 returns a NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4EntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPayoutKeysetPage returns a NoFrixionMoneyMoovModelsPayoutKeysetPage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPayoutKeysetPage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPayoutKeysetPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPayoutKeysetPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPayoutKeysetPageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPayoutMetric returns a NoFrixionMoneyMoovModelsPayoutMetric entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPayoutMetric(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPayoutMetric(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPayoutMetric(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPayoutMetricEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPayoutsPayoutsCreate returns a NoFrixionMoneyMoovModelsPayoutsPayoutsCreate entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPayoutsPayoutsCreateEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsPayrun returns a NoFrixionMoneyMoovModelsPayrun entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsPayrun(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsPayrun(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsPayrun(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsPayrunEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsReportResult returns a NoFrixionMoneyMoovModelsReportResult entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsReportResult(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsReportResult(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsReportResult(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsReportResultEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsRule returns a NoFrixionMoneyMoovModelsRule entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsRule(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsRule(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsRule(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsRuleEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsTransaction returns a NoFrixionMoneyMoovModelsTransaction entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsTransaction(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsTransaction(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsTransaction(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsTransactionEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsTransactionPage returns a NoFrixionMoneyMoovModelsTransactionPage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsTransactionPage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsTransactionPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsTransactionPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsTransactionPageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsUserInvite returns a NoFrixionMoneyMoovModelsUserInvite entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsUserInvite(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsUserInvite(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsUserInvite(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsUserInviteEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsUserInvitePage returns a NoFrixionMoneyMoovModelsUserInvitePage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsUserInvitePage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsUserInvitePage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsUserInvitePage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsUserInvitePageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsUserPage returns a NoFrixionMoneyMoovModelsUserPage entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsUserPage(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsUserPage(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsUserPage(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsUserPageEntityFunc(sdk, data)
}


// NoFrixionMoneyMoovModelsWebhook returns a NoFrixionMoneyMoovModelsWebhook entity bound to this client.
// Idiomatic usage: client.NoFrixionMoneyMoovModelsWebhook(nil).List(nil, nil) or
// client.NoFrixionMoneyMoovModelsWebhook(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionMoneyMoovModelsWebhook(data map[string]any) NofrixionEntity {
	return NewNoFrixionMoneyMoovModelsWebhookEntityFunc(sdk, data)
}


// OpenBanking returns a OpenBanking entity bound to this client.
// Idiomatic usage: client.OpenBanking(nil).List(nil, nil) or
// client.OpenBanking(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) OpenBanking(data map[string]any) NofrixionEntity {
	return NewOpenBankingEntityFunc(sdk, data)
}


// Payeeverification returns a Payeeverification entity bound to this client.
// Idiomatic usage: client.Payeeverification(nil).List(nil, nil) or
// client.Payeeverification(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Payeeverification(data map[string]any) NofrixionEntity {
	return NewPayeeverificationEntityFunc(sdk, data)
}


// PaymentRequest returns a PaymentRequest entity bound to this client.
// Idiomatic usage: client.PaymentRequest(nil).List(nil, nil) or
// client.PaymentRequest(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PaymentRequest(data map[string]any) NofrixionEntity {
	return NewPaymentRequestEntityFunc(sdk, data)
}


// Payout returns a Payout entity bound to this client.
// Idiomatic usage: client.Payout(nil).List(nil, nil) or
// client.Payout(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Payout(data map[string]any) NofrixionEntity {
	return NewPayoutEntityFunc(sdk, data)
}


// Payrun returns a Payrun entity bound to this client.
// Idiomatic usage: client.Payrun(nil).List(nil, nil) or
// client.Payrun(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Payrun(data map[string]any) NofrixionEntity {
	return NewPayrunEntityFunc(sdk, data)
}


// Reject returns a Reject entity bound to this client.
// Idiomatic usage: client.Reject(nil).List(nil, nil) or
// client.Reject(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Reject(data map[string]any) NofrixionEntity {
	return NewRejectEntityFunc(sdk, data)
}


// Report returns a Report entity bound to this client.
// Idiomatic usage: client.Report(nil).List(nil, nil) or
// client.Report(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Report(data map[string]any) NofrixionEntity {
	return NewReportEntityFunc(sdk, data)
}


// Rule returns a Rule entity bound to this client.
// Idiomatic usage: client.Rule(nil).List(nil, nil) or
// client.Rule(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Rule(data map[string]any) NofrixionEntity {
	return NewRuleEntityFunc(sdk, data)
}


// Send returns a Send entity bound to this client.
// Idiomatic usage: client.Send(nil).List(nil, nil) or
// client.Send(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Send(data map[string]any) NofrixionEntity {
	return NewSendEntityFunc(sdk, data)
}


// Sendbeneficiary returns a Sendbeneficiary entity bound to this client.
// Idiomatic usage: client.Sendbeneficiary(nil).List(nil, nil) or
// client.Sendbeneficiary(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Sendbeneficiary(data map[string]any) NofrixionEntity {
	return NewSendbeneficiaryEntityFunc(sdk, data)
}


// Tag returns a Tag entity bound to this client.
// Idiomatic usage: client.Tag(nil).List(nil, nil) or
// client.Tag(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Tag(data map[string]any) NofrixionEntity {
	return NewTagEntityFunc(sdk, data)
}


// Token returns a Token entity bound to this client.
// Idiomatic usage: client.Token(nil).List(nil, nil) or
// client.Token(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Token(data map[string]any) NofrixionEntity {
	return NewTokenEntityFunc(sdk, data)
}


// Transaction returns a Transaction entity bound to this client.
// Idiomatic usage: client.Transaction(nil).List(nil, nil) or
// client.Transaction(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Transaction(data map[string]any) NofrixionEntity {
	return NewTransactionEntityFunc(sdk, data)
}


// User returns a User entity bound to this client.
// Idiomatic usage: client.User(nil).List(nil, nil) or
// client.User(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) User(data map[string]any) NofrixionEntity {
	return NewUserEntityFunc(sdk, data)
}


// UserInvite returns a UserInvite entity bound to this client.
// Idiomatic usage: client.UserInvite(nil).List(nil, nil) or
// client.UserInvite(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) UserInvite(data map[string]any) NofrixionEntity {
	return NewUserInviteEntityFunc(sdk, data)
}


// Virtual returns a Virtual entity bound to this client.
// Idiomatic usage: client.Virtual(nil).List(nil, nil) or
// client.Virtual(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Virtual(data map[string]any) NofrixionEntity {
	return NewVirtualEntityFunc(sdk, data)
}


// Webhook returns a Webhook entity bound to this client.
// Idiomatic usage: client.Webhook(nil).List(nil, nil) or
// client.Webhook(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Webhook(data map[string]any) NofrixionEntity {
	return NewWebhookEntityFunc(sdk, data)
}


// Whoami returns a Whoami entity bound to this client.
// Idiomatic usage: client.Whoami(nil).List(nil, nil) or
// client.Whoami(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Whoami(data map[string]any) NofrixionEntity {
	return NewWhoamiEntityFunc(sdk, data)
}


// Whoamitrustedapp returns a Whoamitrustedapp entity bound to this client.
// Idiomatic usage: client.Whoamitrustedapp(nil).List(nil, nil) or
// client.Whoamitrustedapp(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Whoamitrustedapp(data map[string]any) NofrixionEntity {
	return NewWhoamitrustedappEntityFunc(sdk, data)
}



func TestSDK(testopts map[string]any, sdkopts map[string]any) *NofrixionSDK {
	if sdkopts == nil {
		sdkopts = map[string]any{}
	}
	sdkopts = vs.Clone(sdkopts).(map[string]any)

	if testopts == nil {
		testopts = map[string]any{}
	}
	testopts = vs.Clone(testopts).(map[string]any)
	testopts["active"] = true

	vs.SetPath(sdkopts, []any{"feature", "test"}, testopts)

	sdk := NewNofrixionSDK(sdkopts)
	sdk.Mode = "test"

	return sdk
}
