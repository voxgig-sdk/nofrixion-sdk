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


// Batch returns a Batch entity bound to this client.
// Idiomatic usage: client.Batch(nil).List(nil, nil) or
// client.Batch(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Batch(data map[string]any) NofrixionEntity {
	return NewBatchEntityFunc(sdk, data)
}


// Beneficiary returns a Beneficiary entity bound to this client.
// Idiomatic usage: client.Beneficiary(nil).List(nil, nil) or
// client.Beneficiary(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Beneficiary(data map[string]any) NofrixionEntity {
	return NewBeneficiaryEntityFunc(sdk, data)
}


// BeneficiaryGroup returns a BeneficiaryGroup entity bound to this client.
// Idiomatic usage: client.BeneficiaryGroup(nil).List(nil, nil) or
// client.BeneficiaryGroup(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) BeneficiaryGroup(data map[string]any) NofrixionEntity {
	return NewBeneficiaryGroupEntityFunc(sdk, data)
}


// Card returns a Card entity bound to this client.
// Idiomatic usage: client.Card(nil).List(nil, nil) or
// client.Card(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Card(data map[string]any) NofrixionEntity {
	return NewCardEntityFunc(sdk, data)
}


// CardCustomerToken returns a CardCustomerToken entity bound to this client.
// Idiomatic usage: client.CardCustomerToken(nil).List(nil, nil) or
// client.CardCustomerToken(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) CardCustomerToken(data map[string]any) NofrixionEntity {
	return NewCardCustomerTokenEntityFunc(sdk, data)
}


// CardPayment returns a CardPayment entity bound to this client.
// Idiomatic usage: client.CardPayment(nil).List(nil, nil) or
// client.CardPayment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) CardPayment(data map[string]any) NofrixionEntity {
	return NewCardPaymentEntityFunc(sdk, data)
}


// CardPublicKey returns a CardPublicKey entity bound to this client.
// Idiomatic usage: client.CardPublicKey(nil).List(nil, nil) or
// client.CardPublicKey(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) CardPublicKey(data map[string]any) NofrixionEntity {
	return NewCardPublicKeyEntityFunc(sdk, data)
}


// Consent returns a Consent entity bound to this client.
// Idiomatic usage: client.Consent(nil).List(nil, nil) or
// client.Consent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Consent(data map[string]any) NofrixionEntity {
	return NewConsentEntityFunc(sdk, data)
}


// Currency returns a Currency entity bound to this client.
// Idiomatic usage: client.Currency(nil).List(nil, nil) or
// client.Currency(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Currency(data map[string]any) NofrixionEntity {
	return NewCurrencyEntityFunc(sdk, data)
}


// DirectDebitBatchSubmit returns a DirectDebitBatchSubmit entity bound to this client.
// Idiomatic usage: client.DirectDebitBatchSubmit(nil).List(nil, nil) or
// client.DirectDebitBatchSubmit(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) DirectDebitBatchSubmit(data map[string]any) NofrixionEntity {
	return NewDirectDebitBatchSubmitEntityFunc(sdk, data)
}


// FxRate returns a FxRate entity bound to this client.
// Idiomatic usage: client.FxRate(nil).List(nil, nil) or
// client.FxRate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) FxRate(data map[string]any) NofrixionEntity {
	return NewFxRateEntityFunc(sdk, data)
}


// IPayment returns a IPayment entity bound to this client.
// Idiomatic usage: client.IPayment(nil).List(nil, nil) or
// client.IPayment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) IPayment(data map[string]any) NofrixionEntity {
	return NewIPaymentEntityFunc(sdk, data)
}


// Mandate returns a Mandate entity bound to this client.
// Idiomatic usage: client.Mandate(nil).List(nil, nil) or
// client.Mandate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Mandate(data map[string]any) NofrixionEntity {
	return NewMandateEntityFunc(sdk, data)
}


// Merchant returns a Merchant entity bound to this client.
// Idiomatic usage: client.Merchant(nil).List(nil, nil) or
// client.Merchant(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Merchant(data map[string]any) NofrixionEntity {
	return NewMerchantEntityFunc(sdk, data)
}


// MerchantAuthorisationSetting returns a MerchantAuthorisationSetting entity bound to this client.
// Idiomatic usage: client.MerchantAuthorisationSetting(nil).List(nil, nil) or
// client.MerchantAuthorisationSetting(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) MerchantAuthorisationSetting(data map[string]any) NofrixionEntity {
	return NewMerchantAuthorisationSettingEntityFunc(sdk, data)
}


// MerchantDirectDebitMandate returns a MerchantDirectDebitMandate entity bound to this client.
// Idiomatic usage: client.MerchantDirectDebitMandate(nil).List(nil, nil) or
// client.MerchantDirectDebitMandate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) MerchantDirectDebitMandate(data map[string]any) NofrixionEntity {
	return NewMerchantDirectDebitMandateEntityFunc(sdk, data)
}


// MerchantPayByBankSetting returns a MerchantPayByBankSetting entity bound to this client.
// Idiomatic usage: client.MerchantPayByBankSetting(nil).List(nil, nil) or
// client.MerchantPayByBankSetting(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) MerchantPayByBankSetting(data map[string]any) NofrixionEntity {
	return NewMerchantPayByBankSettingEntityFunc(sdk, data)
}


// MerchantPaymentRequestTemplate returns a MerchantPaymentRequestTemplate entity bound to this client.
// Idiomatic usage: client.MerchantPaymentRequestTemplate(nil).List(nil, nil) or
// client.MerchantPaymentRequestTemplate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) MerchantPaymentRequestTemplate(data map[string]any) NofrixionEntity {
	return NewMerchantPaymentRequestTemplateEntityFunc(sdk, data)
}


// MerchantToken returns a MerchantToken entity bound to this client.
// Idiomatic usage: client.MerchantToken(nil).List(nil, nil) or
// client.MerchantToken(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) MerchantToken(data map[string]any) NofrixionEntity {
	return NewMerchantTokenEntityFunc(sdk, data)
}


// Metadata returns a Metadata entity bound to this client.
// Idiomatic usage: client.Metadata(nil).List(nil, nil) or
// client.Metadata(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Metadata(data map[string]any) NofrixionEntity {
	return NewMetadataEntityFunc(sdk, data)
}


// NoFrixionVersion returns a NoFrixionVersion entity bound to this client.
// Idiomatic usage: client.NoFrixionVersion(nil).List(nil, nil) or
// client.NoFrixionVersion(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) NoFrixionVersion(data map[string]any) NofrixionEntity {
	return NewNoFrixionVersionEntityFunc(sdk, data)
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


// Payment returns a Payment entity bound to this client.
// Idiomatic usage: client.Payment(nil).List(nil, nil) or
// client.Payment(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Payment(data map[string]any) NofrixionEntity {
	return NewPaymentEntityFunc(sdk, data)
}


// PaymentAccount returns a PaymentAccount entity bound to this client.
// Idiomatic usage: client.PaymentAccount(nil).List(nil, nil) or
// client.PaymentAccount(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PaymentAccount(data map[string]any) NofrixionEntity {
	return NewPaymentAccountEntityFunc(sdk, data)
}


// PaymentAccountMinimal returns a PaymentAccountMinimal entity bound to this client.
// Idiomatic usage: client.PaymentAccountMinimal(nil).List(nil, nil) or
// client.PaymentAccountMinimal(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PaymentAccountMinimal(data map[string]any) NofrixionEntity {
	return NewPaymentAccountMinimalEntityFunc(sdk, data)
}


// PaymentInitiation returns a PaymentInitiation entity bound to this client.
// Idiomatic usage: client.PaymentInitiation(nil).List(nil, nil) or
// client.PaymentInitiation(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PaymentInitiation(data map[string]any) NofrixionEntity {
	return NewPaymentInitiationEntityFunc(sdk, data)
}


// PaymentRequest returns a PaymentRequest entity bound to this client.
// Idiomatic usage: client.PaymentRequest(nil).List(nil, nil) or
// client.PaymentRequest(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PaymentRequest(data map[string]any) NofrixionEntity {
	return NewPaymentRequestEntityFunc(sdk, data)
}


// PaymentRequestEvent returns a PaymentRequestEvent entity bound to this client.
// Idiomatic usage: client.PaymentRequestEvent(nil).List(nil, nil) or
// client.PaymentRequestEvent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PaymentRequestEvent(data map[string]any) NofrixionEntity {
	return NewPaymentRequestEventEntityFunc(sdk, data)
}


// PaymentRequestMetric returns a PaymentRequestMetric entity bound to this client.
// Idiomatic usage: client.PaymentRequestMetric(nil).List(nil, nil) or
// client.PaymentRequestMetric(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PaymentRequestMetric(data map[string]any) NofrixionEntity {
	return NewPaymentRequestMetricEntityFunc(sdk, data)
}


// PaymentRequestMinimal returns a PaymentRequestMinimal entity bound to this client.
// Idiomatic usage: client.PaymentRequestMinimal(nil).List(nil, nil) or
// client.PaymentRequestMinimal(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PaymentRequestMinimal(data map[string]any) NofrixionEntity {
	return NewPaymentRequestMinimalEntityFunc(sdk, data)
}


// PaymentRequestResult returns a PaymentRequestResult entity bound to this client.
// Idiomatic usage: client.PaymentRequestResult(nil).List(nil, nil) or
// client.PaymentRequestResult(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PaymentRequestResult(data map[string]any) NofrixionEntity {
	return NewPaymentRequestResultEntityFunc(sdk, data)
}


// Payout returns a Payout entity bound to this client.
// Idiomatic usage: client.Payout(nil).List(nil, nil) or
// client.Payout(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Payout(data map[string]any) NofrixionEntity {
	return NewPayoutEntityFunc(sdk, data)
}


// PayoutKeyset returns a PayoutKeyset entity bound to this client.
// Idiomatic usage: client.PayoutKeyset(nil).List(nil, nil) or
// client.PayoutKeyset(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PayoutKeyset(data map[string]any) NofrixionEntity {
	return NewPayoutKeysetEntityFunc(sdk, data)
}


// PayoutMetric returns a PayoutMetric entity bound to this client.
// Idiomatic usage: client.PayoutMetric(nil).List(nil, nil) or
// client.PayoutMetric(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) PayoutMetric(data map[string]any) NofrixionEntity {
	return NewPayoutMetricEntityFunc(sdk, data)
}


// Payrun returns a Payrun entity bound to this client.
// Idiomatic usage: client.Payrun(nil).List(nil, nil) or
// client.Payrun(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Payrun(data map[string]any) NofrixionEntity {
	return NewPayrunEntityFunc(sdk, data)
}


// Report returns a Report entity bound to this client.
// Idiomatic usage: client.Report(nil).List(nil, nil) or
// client.Report(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Report(data map[string]any) NofrixionEntity {
	return NewReportEntityFunc(sdk, data)
}


// ReportResult returns a ReportResult entity bound to this client.
// Idiomatic usage: client.ReportResult(nil).List(nil, nil) or
// client.ReportResult(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) ReportResult(data map[string]any) NofrixionEntity {
	return NewReportResultEntityFunc(sdk, data)
}


// Role returns a Role entity bound to this client.
// Idiomatic usage: client.Role(nil).List(nil, nil) or
// client.Role(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Role(data map[string]any) NofrixionEntity {
	return NewRoleEntityFunc(sdk, data)
}


// Rule returns a Rule entity bound to this client.
// Idiomatic usage: client.Rule(nil).List(nil, nil) or
// client.Rule(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) Rule(data map[string]any) NofrixionEntity {
	return NewRuleEntityFunc(sdk, data)
}


// RuleEvent returns a RuleEvent entity bound to this client.
// Idiomatic usage: client.RuleEvent(nil).List(nil, nil) or
// client.RuleEvent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NofrixionSDK) RuleEvent(data map[string]any) NofrixionEntity {
	return NewRuleEventEntityFunc(sdk, data)
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
