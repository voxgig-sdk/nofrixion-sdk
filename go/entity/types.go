// Typed models for the Nofrixion SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Account is the typed data model for the account entity.
type Account struct {
	AccountId *string `json:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	AccountSupplierName *string `json:"account_supplier_name,omitempty"`
	AccountType *string `json:"account_type,omitempty"`
	AvailableBalance *float64 `json:"available_balance,omitempty"`
	AvailableBalanceMinorUnit *int `json:"available_balance_minor_unit,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnit *int `json:"balance_minor_unit,omitempty"`
	BankName *string `json:"bank_name,omitempty"`
	ConsentId *string `json:"consent_id,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	CreatedByDisplayName *string `json:"created_by_display_name,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"default_payment_rail,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	ExpiryDate *string `json:"expiry_date,omitempty"`
	ExternalAccountIcon *string `json:"external_account_icon,omitempty"`
	Format *string `json:"format,omitempty"`
	FromDate *string `json:"from_date,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsConnectedAccount *bool `json:"is_connected_account,omitempty"`
	IsDefault *bool `json:"is_default,omitempty"`
	IsTrustAccount *bool `json:"is_trust_account,omitempty"`
	IsVirtual *bool `json:"is_virtual,omitempty"`
	LastTransaction *map[string]any `json:"last_transaction,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantName *string `json:"merchant_name,omitempty"`
	PhysicalAccountId *string `json:"physical_account_id,omitempty"`
	RoleID *[]any `json:"role_i_d,omitempty"`
	Rule *[]any `json:"rule,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submitted_payouts_balance,omitempty"`
	SubmittedPayoutsBalanceMinorUnit *int `json:"submitted_payouts_balance_minor_unit,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierPhysicalAccountId *string `json:"supplier_physical_account_id,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplier_sepa_instant_status,omitempty"`
	ToDate *string `json:"to_date,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xero_bank_feed_connection_status,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xero_bank_feed_last_synced_at,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xero_bank_feed_sync_last_failed_at,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xero_bank_feed_sync_last_failure_reason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xero_bank_feed_sync_status,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xero_unsynchronised_transactions_count,omitempty"`
}

// AccountLoadMatch is the typed request payload for Account.LoadTyped.
type AccountLoadMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	Id string `json:"id"`
	MerchantId *string `json:"merchant_id,omitempty"`
}

// AccountListMatch is the typed request payload for Account.ListTyped.
type AccountListMatch struct {
	MerchantId *string `json:"merchant_id,omitempty"`
}

// AccountCreateData is the typed request payload for Account.CreateTyped.
type AccountCreateData struct {
	AccountId *string `json:"account_id,omitempty"`
	Currency *string `json:"currency,omitempty"`
}

// AccountUpdateData is the typed request payload for Account.UpdateTyped.
type AccountUpdateData struct {
	AccountId *string `json:"account_id,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	Id *string `json:"id,omitempty"`
}

// AccountRemoveMatch is the typed request payload for Account.RemoveTyped.
type AccountRemoveMatch struct {
	Id string `json:"id"`
}

// Beneficiary is the typed data model for the beneficiary entity.
type Beneficiary struct {
	ApprovalCallbackUrl *string `json:"approval_callback_url,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BeneficiaryEvent *[]any `json:"beneficiary_event,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency string `json:"currency"`
	Destination *map[string]any `json:"destination,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	LastAuthorised *string `json:"last_authorised,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name string `json:"name"`
	Nonce *string `json:"nonce,omitempty"`
	SourceAccount *[]any `json:"source_account,omitempty"`
	SourceAccountID *[]any `json:"source_account_i_d,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
}

// BeneficiaryLoadMatch is the typed request payload for Beneficiary.LoadTyped.
type BeneficiaryLoadMatch struct {
	Id string `json:"id"`
	MerchantId *string `json:"merchant_id,omitempty"`
}

// BeneficiaryCreateData is the typed request payload for Beneficiary.CreateTyped.
type BeneficiaryCreateData struct {
	Id *string `json:"id,omitempty"`
}

// BeneficiaryUpdateData is the typed request payload for Beneficiary.UpdateTyped.
type BeneficiaryUpdateData struct {
	Id string `json:"id"`
}

// BeneficiaryRemoveMatch is the typed request payload for Beneficiary.RemoveTyped.
type BeneficiaryRemoveMatch struct {
	Id string `json:"id"`
}

// Cancel is the typed data model for the cancel entity.
type Cancel struct {
	AccountId *string `json:"account_id,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	ApprovePayoutUrl *string `json:"approve_payout_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanProcess *bool `json:"can_process,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	ChargeBearer *string `json:"charge_bearer,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserId *string `json:"current_user_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Document *[]any `json:"document,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formatted_fx_destination_amount,omitempty"`
	FormattedSchedule *string `json:"formatted_schedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formatted_schedule_day_only,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formatted_source_account_available_balance,omitempty"`
	FxDestinationAmount *float64 `json:"fx_destination_amount,omitempty"`
	FxDestinationAmountMinorUnit *int `json:"fx_destination_amount_minor_unit,omitempty"`
	FxDestinationCurrency *string `json:"fx_destination_currency,omitempty"`
	FxQuoteExpiresAt *string `json:"fx_quote_expires_at,omitempty"`
	FxQuoteId *string `json:"fx_quote_id,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	FxUseDestinationAmount *bool `json:"fx_use_destination_amount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsFailed *bool `json:"is_failed,omitempty"`
	IsSettled *bool `json:"is_settled,omitempty"`
	IsSubmitted *bool `json:"is_submitted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentRail *string `json:"payment_rail,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PayrunName *string `json:"payrun_name,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"source_account_available_balance,omitempty"`
	SourceAccountAvailableBalanceMinorUnit *int `json:"source_account_available_balance_minor_unit,omitempty"`
	SourceAccountBic *string `json:"source_account_bic,omitempty"`
	SourceAccountCurrency *string `json:"source_account_currency,omitempty"`
	SourceAccountIban *string `json:"source_account_iban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"source_account_identifier"`
	SourceAccountName *string `json:"source_account_name,omitempty"`
	SourceAccountNumber *string `json:"source_account_number,omitempty"`
	SourceAccountSortcode *string `json:"source_account_sortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TopupPayrunId *string `json:"topup_payrun_id,omitempty"`
	TransactedAmount *float64 `json:"transacted_amount,omitempty"`
	TransactedFxAmount *float64 `json:"transacted_fx_amount,omitempty"`
	TransactedFxRate *float64 `json:"transacted_fx_rate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// CancelUpdateData is the typed request payload for Cancel.UpdateTyped.
type CancelUpdateData struct {
	Id string `json:"id"`
}

// Disable is the typed data model for the disable entity.
type Disable struct {
	ApprovalCallbackUrl *string `json:"approval_callback_url,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BeneficiaryEvent *[]any `json:"beneficiary_event,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency string `json:"currency"`
	Destination *map[string]any `json:"destination,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	LastAuthorised *string `json:"last_authorised,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name string `json:"name"`
	Nonce *string `json:"nonce,omitempty"`
	SourceAccount *[]any `json:"source_account,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
}

// DisableUpdateData is the typed request payload for Disable.UpdateTyped.
type DisableUpdateData struct {
	Id string `json:"id"`
}

// Enable is the typed data model for the enable entity.
type Enable struct {
	ApprovalCallbackUrl *string `json:"approval_callback_url,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BeneficiaryEvent *[]any `json:"beneficiary_event,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency string `json:"currency"`
	Destination *map[string]any `json:"destination,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	LastAuthorised *string `json:"last_authorised,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name string `json:"name"`
	Nonce *string `json:"nonce,omitempty"`
	SourceAccount *[]any `json:"source_account,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
}

// EnableUpdateData is the typed request payload for Enable.UpdateTyped.
type EnableUpdateData struct {
	Id string `json:"id"`
}

// Merchant is the typed data model for the merchant entity.
type Merchant struct {
	Reason *string `json:"reason,omitempty"`
}

// MerchantLoadMatch is the typed request payload for Merchant.LoadTyped.
type MerchantLoadMatch struct {
	MerchantId string `json:"merchant_id"`
}

// MerchantUpdateData is the typed request payload for Merchant.UpdateTyped.
type MerchantUpdateData struct {
	Id string `json:"id"`
}

// MerchantRemoveMatch is the typed request payload for Merchant.RemoveTyped.
type MerchantRemoveMatch struct {
	Id *string `json:"id,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	TagId *string `json:"tag_id,omitempty"`
}

// Metadata is the typed data model for the metadata entity.
type Metadata struct {
}

// MetadataLoadMatch is the typed request payload for Metadata.LoadTyped.
type MetadataLoadMatch struct {
}

// NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage is the typed data model for the no_frixion_biz_biz_models_paging_merchant_direct_debit_mandate_page entity.
type NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage struct {
	ApprovedAt *string `json:"approved_at,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomerAccountNumber *string `json:"customer_account_number,omitempty"`
	CustomerCity *string `json:"customer_city,omitempty"`
	CustomerCountryCode *string `json:"customer_country_code,omitempty"`
	CustomerCountryName *string `json:"customer_country_name,omitempty"`
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	CustomerFirstName *string `json:"customer_first_name,omitempty"`
	CustomerIban *string `json:"customer_iban,omitempty"`
	CustomerLastName *string `json:"customer_last_name,omitempty"`
	CustomerSortCode *string `json:"customer_sort_code,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsRecurring *bool `json:"is_recurring,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Reference *string `json:"reference,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplierBankAccountId *string `json:"supplier_bank_account_id,omitempty"`
	SupplierCustomerId *string `json:"supplier_customer_id,omitempty"`
	SupplierMandateId *string `json:"supplier_mandate_id,omitempty"`
	SupplierName *string `json:"supplier_name,omitempty"`
	SupplierStatus *string `json:"supplier_status,omitempty"`
}

// NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageListMatch is the typed request payload for NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage.ListTyped.
type NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageListMatch struct {
	ApprovedAt *string `json:"approved_at,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomerAccountNumber *string `json:"customer_account_number,omitempty"`
	CustomerCity *string `json:"customer_city,omitempty"`
	CustomerCountryCode *string `json:"customer_country_code,omitempty"`
	CustomerCountryName *string `json:"customer_country_name,omitempty"`
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	CustomerFirstName *string `json:"customer_first_name,omitempty"`
	CustomerIban *string `json:"customer_iban,omitempty"`
	CustomerLastName *string `json:"customer_last_name,omitempty"`
	CustomerSortCode *string `json:"customer_sort_code,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsRecurring *bool `json:"is_recurring,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Reference *string `json:"reference,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplierBankAccountId *string `json:"supplier_bank_account_id,omitempty"`
	SupplierCustomerId *string `json:"supplier_customer_id,omitempty"`
	SupplierMandateId *string `json:"supplier_mandate_id,omitempty"`
	SupplierName *string `json:"supplier_name,omitempty"`
	SupplierStatus *string `json:"supplier_status,omitempty"`
}

// NoFrixionBizBizModelsPagingPaymentRequestPage is the typed data model for the no_frixion_biz_biz_models_paging_payment_request_page entity.
type NoFrixionBizBizModelsPagingPaymentRequestPage struct {
	Address *[]any `json:"address,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amount_pending,omitempty"`
	AmountReceived *float64 `json:"amount_received,omitempty"`
	AmountRefunded *float64 `json:"amount_refunded,omitempty"`
	AutoSendReceipt *bool `json:"auto_send_receipt,omitempty"`
	BaseOriginUrl *string `json:"base_origin_url,omitempty"`
	CallbackUrl *string `json:"callback_url,omitempty"`
	CardAuthorizeOnly *bool `json:"card_authorize_only,omitempty"`
	CardCreateToken *bool `json:"card_create_token,omitempty"`
	CardCreateTokenMode *string `json:"card_create_token_mode,omitempty"`
	CardIgnoreCvn *bool `json:"card_ignore_cvn,omitempty"`
	CardProcessorMerchantId *string `json:"card_processor_merchant_id,omitempty"`
	CardStripePaymentIntentId *string `json:"card_stripe_payment_intent_id,omitempty"`
	CardStripePaymentIntentSecret *string `json:"card_stripe_payment_intent_secret,omitempty"`
	CreatedByUser map[string]any `json:"created_by_user"`
	Currency *string `json:"currency,omitempty"`
	CustomField *[]any `json:"custom_field,omitempty"`
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	CustomerName *string `json:"customer_name,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destination_account,omitempty"`
	DirectDebitPayment *map[string]any `json:"direct_debit_payment,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FailureCallbackUrl *string `json:"failure_callback_url,omitempty"`
	FieldDisplaySetting *[]any `json:"field_display_setting,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hosted_pay_checkout_url,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignore_address_verification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"inserted_sortable,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	LightningInvoice *string `json:"lightning_invoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightning_invoice_expires_at,omitempty"`
	MerchantDirectDebitMandateId *string `json:"merchant_direct_debit_mandate_id,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	NotificationEmailAddress *string `json:"notification_email_address,omitempty"`
	NotificationRoleID *[]any `json:"notification_role_i_d,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	PartialPaymentMethod *string `json:"partial_payment_method,omitempty"`
	PartialPaymentStep *string `json:"partial_payment_step,omitempty"`
	PaymentAttempt *[]any `json:"payment_attempt,omitempty"`
	PaymentMethod *[]any `json:"payment_method,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PispAccountId *string `json:"pisp_account_id,omitempty"`
	PriorityBankId *string `json:"priority_bank_id,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSecond *int `json:"sandbox_settle_delay_in_second,omitempty"`
	ShippingAddress *map[string]any `json:"shipping_address,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"success_web_hook_url,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCard *[]any `json:"tokenised_card,omitempty"`
	Transaction *[]any `json:"transaction,omitempty"`
	UseHostedPaymentPage *bool `json:"use_hosted_payment_page,omitempty"`
}

// NoFrixionBizBizModelsPagingPaymentRequestPageListMatch is the typed request payload for NoFrixionBizBizModelsPagingPaymentRequestPage.ListTyped.
type NoFrixionBizBizModelsPagingPaymentRequestPageListMatch struct {
	Address *[]any `json:"address,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amount_pending,omitempty"`
	AmountReceived *float64 `json:"amount_received,omitempty"`
	AmountRefunded *float64 `json:"amount_refunded,omitempty"`
	AutoSendReceipt *bool `json:"auto_send_receipt,omitempty"`
	BaseOriginUrl *string `json:"base_origin_url,omitempty"`
	CallbackUrl *string `json:"callback_url,omitempty"`
	CardAuthorizeOnly *bool `json:"card_authorize_only,omitempty"`
	CardCreateToken *bool `json:"card_create_token,omitempty"`
	CardCreateTokenMode *string `json:"card_create_token_mode,omitempty"`
	CardIgnoreCvn *bool `json:"card_ignore_cvn,omitempty"`
	CardProcessorMerchantId *string `json:"card_processor_merchant_id,omitempty"`
	CardStripePaymentIntentId *string `json:"card_stripe_payment_intent_id,omitempty"`
	CardStripePaymentIntentSecret *string `json:"card_stripe_payment_intent_secret,omitempty"`
	CreatedByUser *map[string]any `json:"created_by_user,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomField *[]any `json:"custom_field,omitempty"`
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	CustomerName *string `json:"customer_name,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destination_account,omitempty"`
	DirectDebitPayment *map[string]any `json:"direct_debit_payment,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FailureCallbackUrl *string `json:"failure_callback_url,omitempty"`
	FieldDisplaySetting *[]any `json:"field_display_setting,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hosted_pay_checkout_url,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignore_address_verification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"inserted_sortable,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	LightningInvoice *string `json:"lightning_invoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightning_invoice_expires_at,omitempty"`
	MerchantDirectDebitMandateId *string `json:"merchant_direct_debit_mandate_id,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	NotificationEmailAddress *string `json:"notification_email_address,omitempty"`
	NotificationRoleID *[]any `json:"notification_role_i_d,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	PartialPaymentMethod *string `json:"partial_payment_method,omitempty"`
	PartialPaymentStep *string `json:"partial_payment_step,omitempty"`
	PaymentAttempt *[]any `json:"payment_attempt,omitempty"`
	PaymentMethod *[]any `json:"payment_method,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PispAccountId *string `json:"pisp_account_id,omitempty"`
	PriorityBankId *string `json:"priority_bank_id,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSecond *int `json:"sandbox_settle_delay_in_second,omitempty"`
	ShippingAddress *map[string]any `json:"shipping_address,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"success_web_hook_url,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCard *[]any `json:"tokenised_card,omitempty"`
	Transaction *[]any `json:"transaction,omitempty"`
	UseHostedPaymentPage *bool `json:"use_hosted_payment_page,omitempty"`
}

// NoFrixionBizBizModelsPagingPayoutPage is the typed data model for the no_frixion_biz_biz_models_paging_payout_page entity.
type NoFrixionBizBizModelsPagingPayoutPage struct {
	AccountId *string `json:"account_id,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	ApprovePayoutUrl *string `json:"approve_payout_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanProcess *bool `json:"can_process,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	ChargeBearer *string `json:"charge_bearer,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserId *string `json:"current_user_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Document *[]any `json:"document,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formatted_fx_destination_amount,omitempty"`
	FormattedSchedule *string `json:"formatted_schedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formatted_schedule_day_only,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formatted_source_account_available_balance,omitempty"`
	FxDestinationAmount *float64 `json:"fx_destination_amount,omitempty"`
	FxDestinationAmountMinorUnit *int `json:"fx_destination_amount_minor_unit,omitempty"`
	FxDestinationCurrency *string `json:"fx_destination_currency,omitempty"`
	FxQuoteExpiresAt *string `json:"fx_quote_expires_at,omitempty"`
	FxQuoteId *string `json:"fx_quote_id,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	FxUseDestinationAmount *bool `json:"fx_use_destination_amount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsFailed *bool `json:"is_failed,omitempty"`
	IsSettled *bool `json:"is_settled,omitempty"`
	IsSubmitted *bool `json:"is_submitted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentRail *string `json:"payment_rail,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PayrunName *string `json:"payrun_name,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"source_account_available_balance,omitempty"`
	SourceAccountAvailableBalanceMinorUnit *int `json:"source_account_available_balance_minor_unit,omitempty"`
	SourceAccountBic *string `json:"source_account_bic,omitempty"`
	SourceAccountCurrency *string `json:"source_account_currency,omitempty"`
	SourceAccountIban *string `json:"source_account_iban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"source_account_identifier"`
	SourceAccountName *string `json:"source_account_name,omitempty"`
	SourceAccountNumber *string `json:"source_account_number,omitempty"`
	SourceAccountSortcode *string `json:"source_account_sortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TopupPayrunId *string `json:"topup_payrun_id,omitempty"`
	TransactedAmount *float64 `json:"transacted_amount,omitempty"`
	TransactedFxAmount *float64 `json:"transacted_fx_amount,omitempty"`
	TransactedFxRate *float64 `json:"transacted_fx_rate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// NoFrixionBizBizModelsPagingPayoutPageListMatch is the typed request payload for NoFrixionBizBizModelsPagingPayoutPage.ListTyped.
type NoFrixionBizBizModelsPagingPayoutPageListMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
}

// NoFrixionBizBizModelsPagingPayrunPage is the typed data model for the no_frixion_biz_biz_models_paging_payrun_page entity.
type NoFrixionBizBizModelsPagingPayrunPage struct {
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisationDate *string `json:"authorisation_date,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanDelete *bool `json:"can_delete,omitempty"`
	CanEdit *bool `json:"can_edit,omitempty"`
	Event *[]any `json:"event,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	Invoice *[]any `json:"invoice,omitempty"`
	InvoicesMinimal *[]any `json:"invoices_minimal,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	LastUpdatedBy map[string]any `json:"last_updated_by"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Payment *[]any `json:"payment,omitempty"`
	Payout *[]any `json:"payout,omitempty"`
	PayoutsCount *int `json:"payouts_count,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	SourceAccount *[]any `json:"source_account,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalEur *float64 `json:"total_eur,omitempty"`
	TotalGbp *float64 `json:"total_gbp,omitempty"`
	TotalUsd *float64 `json:"total_usd,omitempty"`
}

// NoFrixionBizBizModelsPagingPayrunPageListMatch is the typed request payload for NoFrixionBizBizModelsPagingPayrunPage.ListTyped.
type NoFrixionBizBizModelsPagingPayrunPageListMatch struct {
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisationDate *string `json:"authorisation_date,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanDelete *bool `json:"can_delete,omitempty"`
	CanEdit *bool `json:"can_edit,omitempty"`
	Event *[]any `json:"event,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	Invoice *[]any `json:"invoice,omitempty"`
	InvoicesMinimal *[]any `json:"invoices_minimal,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	LastUpdatedBy *map[string]any `json:"last_updated_by,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Payment *[]any `json:"payment,omitempty"`
	Payout *[]any `json:"payout,omitempty"`
	PayoutsCount *int `json:"payouts_count,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	SourceAccount *[]any `json:"source_account,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalEur *float64 `json:"total_eur,omitempty"`
	TotalGbp *float64 `json:"total_gbp,omitempty"`
	TotalUsd *float64 `json:"total_usd,omitempty"`
}

// NoFrixionBizBizModelsPagingRuleEventsPage is the typed data model for the no_frixion_biz_biz_models_paging_rule_events_page entity.
type NoFrixionBizBizModelsPagingRuleEventsPage struct {
	ErrorMessage *string `json:"error_message,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsAuthoriseToEnable *bool `json:"is_authorise_to_enable,omitempty"`
	Message *string `json:"message,omitempty"`
	RawResponse *string `json:"raw_response,omitempty"`
	RuleEventType *string `json:"rule_event_type,omitempty"`
	RuleId *string `json:"rule_id,omitempty"`
	User map[string]any `json:"user"`
}

// NoFrixionBizBizModelsPagingRuleEventsPageListMatch is the typed request payload for NoFrixionBizBizModelsPagingRuleEventsPage.ListTyped.
type NoFrixionBizBizModelsPagingRuleEventsPageListMatch struct {
	RuleId string `json:"rule_id"`
}

// NoFrixionBizBizModelsPagingRulesPage is the typed data model for the no_frixion_biz_biz_models_paging_rules_page entity.
type NoFrixionBizBizModelsPagingRulesPage struct {
	Account *map[string]any `json:"account,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
	ApproveUrl *string `json:"approve_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	Description *string `json:"description,omitempty"`
	EndAt *string `json:"end_at,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsDisabled *bool `json:"is_disabled,omitempty"`
	LastExecutedAt *string `json:"last_executed_at,omitempty"`
	LastRunAtTransactionDate *string `json:"last_run_at_transaction_date,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce string `json:"nonce"`
	OnApprovedWebHookUrl *string `json:"on_approved_web_hook_url,omitempty"`
	OnExecutionErrorWebHookUrl *string `json:"on_execution_error_web_hook_url,omitempty"`
	OnExecutionSuccessWebHookUrl *string `json:"on_execution_success_web_hook_url,omitempty"`
	StartAt *string `json:"start_at,omitempty"`
	Status *string `json:"status,omitempty"`
	SweepAction *map[string]any `json:"sweep_action,omitempty"`
	TimeZoneId *string `json:"time_zone_id,omitempty"`
	TriggerCronExpression *string `json:"trigger_cron_expression,omitempty"`
	TriggerOnPayIn *bool `json:"trigger_on_pay_in,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	WebHookSecret *string `json:"web_hook_secret,omitempty"`
}

// NoFrixionBizBizModelsPagingRulesPageListMatch is the typed request payload for NoFrixionBizBizModelsPagingRulesPage.ListTyped.
type NoFrixionBizBizModelsPagingRulesPageListMatch struct {
	Account *map[string]any `json:"account,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
	ApproveUrl *string `json:"approve_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CreatedBy *map[string]any `json:"created_by,omitempty"`
	Description *string `json:"description,omitempty"`
	EndAt *string `json:"end_at,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsDisabled *bool `json:"is_disabled,omitempty"`
	LastExecutedAt *string `json:"last_executed_at,omitempty"`
	LastRunAtTransactionDate *string `json:"last_run_at_transaction_date,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	OnApprovedWebHookUrl *string `json:"on_approved_web_hook_url,omitempty"`
	OnExecutionErrorWebHookUrl *string `json:"on_execution_error_web_hook_url,omitempty"`
	OnExecutionSuccessWebHookUrl *string `json:"on_execution_success_web_hook_url,omitempty"`
	StartAt *string `json:"start_at,omitempty"`
	Status *string `json:"status,omitempty"`
	SweepAction *map[string]any `json:"sweep_action,omitempty"`
	TimeZoneId *string `json:"time_zone_id,omitempty"`
	TriggerCronExpression *string `json:"trigger_cron_expression,omitempty"`
	TriggerOnPayIn *bool `json:"trigger_on_pay_in,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	WebHookSecret *string `json:"web_hook_secret,omitempty"`
}

// NoFrixionBizBizModelsPaymentsCardPayment is the typed data model for the no_frixion_biz_biz_models_payments_card_payment entity.
type NoFrixionBizBizModelsPaymentsCardPayment struct {
	AuthorizedAmount *string `json:"authorized_amount,omitempty"`
	CurrencyCode *string `json:"currency_code,omitempty"`
	IsPayerAuthenticationRequired *bool `json:"is_payer_authentication_required,omitempty"`
	IsSoftDecline *bool `json:"is_soft_decline,omitempty"`
	PayerAuthenticationAccessToken *string `json:"payer_authentication_access_token,omitempty"`
	PayerAuthenticationMerchantData *string `json:"payer_authentication_merchant_data,omitempty"`
	PayerAuthenticationUrl *string `json:"payer_authentication_url,omitempty"`
	PayerAuthenticationWindowHeight *int `json:"payer_authentication_window_height,omitempty"`
	PayerAuthenticationWindowWidth *int `json:"payer_authentication_window_width,omitempty"`
	PaymentRequestCallbackUrl *string `json:"payment_request_callback_url,omitempty"`
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	RequestId *string `json:"request_id,omitempty"`
	ResponseCode *string `json:"response_code,omitempty"`
	ResponseType *string `json:"response_type,omitempty"`
	Status *string `json:"status,omitempty"`
	ThreeDsRedirectUrl *string `json:"three_ds_redirect_url,omitempty"`
	TransactionId *string `json:"transaction_id,omitempty"`
}

// NoFrixionBizBizModelsPaymentsCardPaymentCreateData is the typed request payload for NoFrixionBizBizModelsPaymentsCardPayment.CreateTyped.
type NoFrixionBizBizModelsPaymentsCardPaymentCreateData struct {
	PartialRefundAmount *float64 `json:"partial_refund_amount,omitempty"`
	PaymentrequestId string `json:"paymentrequest_id"`
}

// NoFrixionBizBizModelsPaymentsCardPublicKey is the typed data model for the no_frixion_biz_biz_models_payments_card_public_key entity.
type NoFrixionBizBizModelsPaymentsCardPublicKey struct {
	Jwt *string `json:"jwt,omitempty"`
}

// NoFrixionBizBizModelsPaymentsCardPublicKeyLoadMatch is the typed request payload for NoFrixionBizBizModelsPaymentsCardPublicKey.LoadTyped.
type NoFrixionBizBizModelsPaymentsCardPublicKeyLoadMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries is the typed data model for the no_frixion_money_moov_api_features_beneficiaries_beneficiaries entity.
type NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries struct {
	Beneficiary *[]any `json:"beneficiary,omitempty"`
	FailedBeneficiary *map[string]any `json:"failed_beneficiary,omitempty"`
}

// NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesCreateData is the typed request payload for NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries.CreateTyped.
type NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesCreateData struct {
	Beneficiary *[]any `json:"beneficiary,omitempty"`
	FailedBeneficiary *map[string]any `json:"failed_beneficiary,omitempty"`
}

// NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment is the typed data model for the no_frixion_money_moov_api_features_payment_requests_payment entity.
type NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment struct {
	FailedPaymentRequest *map[string]any `json:"failed_payment_request,omitempty"`
	PaymentRequest *[]any `json:"payment_request,omitempty"`
}

// NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentCreateData is the typed request payload for NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment.CreateTyped.
type NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentCreateData struct {
	FailedPaymentRequest *map[string]any `json:"failed_payment_request,omitempty"`
	PaymentRequest *[]any `json:"payment_request,omitempty"`
}

// NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate is the typed data model for the no_frixion_money_moov_api_features_permissions_roles_create entity.
type NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate struct {
	FailedRole *map[string]any `json:"failed_role,omitempty"`
	Role *[]any `json:"role,omitempty"`
}

// NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateCreateData is the typed request payload for NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate.CreateTyped.
type NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateCreateData struct {
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovApiFeaturesUserInvitesCreate is the typed data model for the no_frixion_money_moov_api_features_user_invites_create entity.
type NoFrixionMoneyMoovApiFeaturesUserInvitesCreate struct {
	FailedUserInvite *map[string]any `json:"failed_user_invite,omitempty"`
	UserInvite *[]any `json:"user_invite,omitempty"`
}

// NoFrixionMoneyMoovApiFeaturesUserInvitesCreateCreateData is the typed request payload for NoFrixionMoneyMoovApiFeaturesUserInvitesCreate.CreateTyped.
type NoFrixionMoneyMoovApiFeaturesUserInvitesCreateCreateData struct {
	FailedUserInvite *map[string]any `json:"failed_user_invite,omitempty"`
	UserInvite *[]any `json:"user_invite,omitempty"`
}

// NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant is the typed data model for the no_frixion_money_moov_models_authorisation_settings_merchant entity.
type NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant struct {
	AmountLower *float64 `json:"amount_lower,omitempty"`
	AmountUpper *float64 `json:"amount_upper,omitempty"`
	AuthorisationType *string `json:"authorisation_type,omitempty"`
	BeneficiariesOnly *bool `json:"beneficiaries_only,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastEditorCantAuthorise *bool `json:"last_editor_cant_authorise,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	NumberOfAuthoriser *int `json:"number_of_authoriser,omitempty"`
	RoleSetting *[]any `json:"role_setting,omitempty"`
}

// NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantListMatch is the typed request payload for NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant.ListTyped.
type NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsBatchPayout is the typed data model for the no_frixion_money_moov_models_batch_payout entity.
type NoFrixionMoneyMoovModelsBatchPayout struct {
	ApproveUrl *string `json:"approve_url,omitempty"`
	Id *string `json:"id,omitempty"`
	Payout *[]any `json:"payout,omitempty"`
}

// NoFrixionMoneyMoovModelsBatchPayoutLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsBatchPayout.LoadTyped.
type NoFrixionMoneyMoovModelsBatchPayoutLoadMatch struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsBatchPayoutCreateData is the typed request payload for NoFrixionMoneyMoovModelsBatchPayout.CreateTyped.
type NoFrixionMoneyMoovModelsBatchPayoutCreateData struct {
	ApproveUrl *string `json:"approve_url,omitempty"`
	Id *string `json:"id,omitempty"`
	Payout *[]any `json:"payout,omitempty"`
}

// NoFrixionMoneyMoovModelsBeneficiaryGroupPage is the typed data model for the no_frixion_money_moov_models_beneficiary_group_page entity.
type NoFrixionMoneyMoovModelsBeneficiaryGroupPage struct {
	GroupMember *[]any `json:"group_member,omitempty"`
	GroupName string `json:"group_name"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsBeneficiaryGroupPageListMatch is the typed request payload for NoFrixionMoneyMoovModelsBeneficiaryGroupPage.ListTyped.
type NoFrixionMoneyMoovModelsBeneficiaryGroupPageListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsBeneficiaryPage is the typed data model for the no_frixion_money_moov_models_beneficiary_page entity.
type NoFrixionMoneyMoovModelsBeneficiaryPage struct {
	ApprovalCallbackUrl *string `json:"approval_callback_url,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BeneficiaryEvent *[]any `json:"beneficiary_event,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency string `json:"currency"`
	Destination *map[string]any `json:"destination,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	LastAuthorised *string `json:"last_authorised,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name string `json:"name"`
	Nonce *string `json:"nonce,omitempty"`
	SourceAccount *[]any `json:"source_account,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
}

// NoFrixionMoneyMoovModelsBeneficiaryPageListMatch is the typed request payload for NoFrixionMoneyMoovModelsBeneficiaryPage.ListTyped.
type NoFrixionMoneyMoovModelsBeneficiaryPageListMatch struct {
	MerchantId *string `json:"merchant_id,omitempty"`
}

// NoFrixionMoneyMoovModelsCardCustomerToken is the typed data model for the no_frixion_money_moov_models_card_customer_token entity.
type NoFrixionMoneyMoovModelsCardCustomerToken struct {
	CardType *string `json:"card_type,omitempty"`
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	ExpiryMonth *string `json:"expiry_month,omitempty"`
	ExpiryYear *string `json:"expiry_year,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastFourDigit *string `json:"last_four_digit,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MaskedCardNumber *string `json:"masked_card_number,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
}

// NoFrixionMoneyMoovModelsCardCustomerTokenLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsCardCustomerToken.LoadTyped.
type NoFrixionMoneyMoovModelsCardCustomerTokenLoadMatch struct {
	CustomerEmailAddress string `json:"customer_email_address"`
}

// NoFrixionMoneyMoovModelsCardCustomerTokenListMatch is the typed request payload for NoFrixionMoneyMoovModelsCardCustomerToken.ListTyped.
type NoFrixionMoneyMoovModelsCardCustomerTokenListMatch struct {
	CustomerEmailAddress string `json:"customer_email_address"`
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsCardCustomerTokenRemoveMatch is the typed request payload for NoFrixionMoneyMoovModelsCardCustomerToken.RemoveTyped.
type NoFrixionMoneyMoovModelsCardCustomerTokenRemoveMatch struct {
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NoFrixionMoneyMoovModelsCurrencyCurrencyInfo is the typed data model for the no_frixion_money_moov_models_currency_currency_info entity.
type NoFrixionMoneyMoovModelsCurrencyCurrencyInfo struct {
	Code *string `json:"code,omitempty"`
	Decimal *int `json:"decimal,omitempty"`
	IsFiat *bool `json:"is_fiat,omitempty"`
	Iso4217AlphaCode *string `json:"iso4217_alpha_code,omitempty"`
	Iso4217NumericCode *string `json:"iso4217_numeric_code,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NoFrixionMoneyMoovModelsCurrencyCurrencyInfoListMatch is the typed request payload for NoFrixionMoneyMoovModelsCurrencyCurrencyInfo.ListTyped.
type NoFrixionMoneyMoovModelsCurrencyCurrencyInfoListMatch struct {
	Code *string `json:"code,omitempty"`
	Decimal *int `json:"decimal,omitempty"`
	IsFiat *bool `json:"is_fiat,omitempty"`
	Iso4217AlphaCode *string `json:"iso4217_alpha_code,omitempty"`
	Iso4217NumericCode *string `json:"iso4217_numeric_code,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// NoFrixionMoneyMoovModelsDirectDebitBatchSubmit is the typed data model for the no_frixion_money_moov_models_direct_debit_batch_submit entity.
type NoFrixionMoneyMoovModelsDirectDebitBatchSubmit struct {
	FailedSubmission *map[string]any `json:"failed_submission,omitempty"`
	SuccessfulSubmission *[]any `json:"successful_submission,omitempty"`
}

// NoFrixionMoneyMoovModelsDirectDebitBatchSubmitCreateData is the typed request payload for NoFrixionMoneyMoovModelsDirectDebitBatchSubmit.CreateTyped.
type NoFrixionMoneyMoovModelsDirectDebitBatchSubmitCreateData struct {
	FailedSubmission *map[string]any `json:"failed_submission,omitempty"`
	SuccessfulSubmission *[]any `json:"successful_submission,omitempty"`
}

// NoFrixionMoneyMoovModelsFxRate is the typed data model for the no_frixion_money_moov_models_fx_rate entity.
type NoFrixionMoneyMoovModelsFxRate struct {
	DestinationCurrency *string `json:"destination_currency,omitempty"`
	ExchangeRate *float64 `json:"exchange_rate,omitempty"`
	ExpiryTime *string `json:"expiry_time,omitempty"`
	QuoteId *string `json:"quote_id,omitempty"`
	SourceCurrency *string `json:"source_currency,omitempty"`
}

// NoFrixionMoneyMoovModelsFxRateLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsFxRate.LoadTyped.
type NoFrixionMoneyMoovModelsFxRateLoadMatch struct {
	Destination string `json:"destination"`
	Source string `json:"source"`
	ValidForMinute int `json:"valid_for_minute"`
}

// NoFrixionMoneyMoovModelsFxRateListMatch is the typed request payload for NoFrixionMoneyMoovModelsFxRate.ListTyped.
type NoFrixionMoneyMoovModelsFxRateListMatch struct {
	Destination string `json:"destination"`
	Source string `json:"source"`
}

// NoFrixionMoneyMoovModelsIPayment is the typed data model for the no_frixion_money_moov_models_i_payment entity.
type NoFrixionMoneyMoovModelsIPayment struct {
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	ResponseType *string `json:"response_type,omitempty"`
}

// NoFrixionMoneyMoovModelsIPaymentCreateData is the typed request payload for NoFrixionMoneyMoovModelsIPayment.CreateTyped.
type NoFrixionMoneyMoovModelsIPaymentCreateData struct {
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	ResponseType *string `json:"response_type,omitempty"`
}

// NoFrixionMoneyMoovModelsMandatesMandate is the typed data model for the no_frixion_money_moov_models_mandates_mandate entity.
type NoFrixionMoneyMoovModelsMandatesMandate struct {
	AccountNumber *string `json:"account_number,omitempty"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 *string `json:"address_line2,omitempty"`
	ApprovedAt *string `json:"approved_at,omitempty"`
	City string `json:"city"`
	CountryCode string `json:"country_code"`
	Currency *string `json:"currency,omitempty"`
	CustomerAccountNumber *string `json:"customer_account_number,omitempty"`
	CustomerCity *string `json:"customer_city,omitempty"`
	CustomerCountryCode *string `json:"customer_country_code,omitempty"`
	CustomerCountryName *string `json:"customer_country_name,omitempty"`
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	CustomerFirstName *string `json:"customer_first_name,omitempty"`
	CustomerIban *string `json:"customer_iban,omitempty"`
	CustomerLastName *string `json:"customer_last_name,omitempty"`
	CustomerSortCode *string `json:"customer_sort_code,omitempty"`
	EmailAddress string `json:"email_address"`
	FirstName string `json:"first_name"`
	Iban *string `json:"iban,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsRecurring *bool `json:"is_recurring,omitempty"`
	LastName string `json:"last_name"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	PostalCode string `json:"postal_code"`
	Reference *string `json:"reference,omitempty"`
	SortCode *string `json:"sort_code,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplierBankAccountId *string `json:"supplier_bank_account_id,omitempty"`
	SupplierCustomerId *string `json:"supplier_customer_id,omitempty"`
	SupplierMandateId *string `json:"supplier_mandate_id,omitempty"`
	SupplierName *string `json:"supplier_name,omitempty"`
	SupplierStatus *string `json:"supplier_status,omitempty"`
}

// NoFrixionMoneyMoovModelsMandatesMandateLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsMandatesMandate.LoadTyped.
type NoFrixionMoneyMoovModelsMandatesMandateLoadMatch struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsMandatesMandateCreateData is the typed request payload for NoFrixionMoneyMoovModelsMandatesMandate.CreateTyped.
type NoFrixionMoneyMoovModelsMandatesMandateCreateData struct {
	AccountNumber *string `json:"account_number,omitempty"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 *string `json:"address_line2,omitempty"`
	ApprovedAt *string `json:"approved_at,omitempty"`
	City string `json:"city"`
	CountryCode string `json:"country_code"`
	Currency *string `json:"currency,omitempty"`
	CustomerAccountNumber *string `json:"customer_account_number,omitempty"`
	CustomerCity *string `json:"customer_city,omitempty"`
	CustomerCountryCode *string `json:"customer_country_code,omitempty"`
	CustomerCountryName *string `json:"customer_country_name,omitempty"`
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	CustomerFirstName *string `json:"customer_first_name,omitempty"`
	CustomerIban *string `json:"customer_iban,omitempty"`
	CustomerLastName *string `json:"customer_last_name,omitempty"`
	CustomerSortCode *string `json:"customer_sort_code,omitempty"`
	EmailAddress string `json:"email_address"`
	FirstName string `json:"first_name"`
	Iban *string `json:"iban,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsRecurring *bool `json:"is_recurring,omitempty"`
	LastName string `json:"last_name"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	PostalCode string `json:"postal_code"`
	Reference *string `json:"reference,omitempty"`
	SortCode *string `json:"sort_code,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplierBankAccountId *string `json:"supplier_bank_account_id,omitempty"`
	SupplierCustomerId *string `json:"supplier_customer_id,omitempty"`
	SupplierMandateId *string `json:"supplier_mandate_id,omitempty"`
	SupplierName *string `json:"supplier_name,omitempty"`
	SupplierStatus *string `json:"supplier_status,omitempty"`
}

// NoFrixionMoneyMoovModelsMerchant is the typed data model for the no_frixion_money_moov_models_merchant entity.
type NoFrixionMoneyMoovModelsMerchant struct {
	AccountCurrency *[]any `json:"account_currency,omitempty"`
	CanHaveTrustAccount *bool `json:"can_have_trust_account,omitempty"`
	CardPaymentProcessor *string `json:"card_payment_processor,omitempty"`
	CompanyId *string `json:"company_id,omitempty"`
	DisplayQrOnHostedPay *bool `json:"display_qr_on_hosted_pay,omitempty"`
	HostedPayVersion *int `json:"hosted_pay_version,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsBlocked *bool `json:"is_blocked,omitempty"`
	IsExited *bool `json:"is_exited,omitempty"`
	IsSuspended *bool `json:"is_suspended,omitempty"`
	Jurisdiction *string `json:"jurisdiction,omitempty"`
	LogoUrlPng *string `json:"logo_url_png,omitempty"`
	LogoUrlSvg *string `json:"logo_url_svg,omitempty"`
	MerchantCategoryCode *string `json:"merchant_category_code,omitempty"`
	Name *string `json:"name,omitempty"`
	Note *string `json:"note,omitempty"`
	ParentMerchant *map[string]any `json:"parent_merchant,omitempty"`
	PaymentAccount *[]any `json:"payment_account,omitempty"`
	PaymentAccountLimit *int `json:"payment_account_limit,omitempty"`
	ShortName *string `json:"short_name,omitempty"`
	SupportedPaymentMethodsList *[]any `json:"supported_payment_methods_list,omitempty"`
	SuspensionReason *string `json:"suspension_reason,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TimeZoneId *string `json:"time_zone_id,omitempty"`
	TradingName *string `json:"trading_name,omitempty"`
	WebHookLimit *int `json:"web_hook_limit,omitempty"`
	YourRoleName *string `json:"your_role_name,omitempty"`
}

// NoFrixionMoneyMoovModelsMerchantLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsMerchant.LoadTyped.
type NoFrixionMoneyMoovModelsMerchantLoadMatch struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsMerchantListMatch is the typed request payload for NoFrixionMoneyMoovModelsMerchant.ListTyped.
type NoFrixionMoneyMoovModelsMerchantListMatch struct {
	AccountCurrency *[]any `json:"account_currency,omitempty"`
	CanHaveTrustAccount *bool `json:"can_have_trust_account,omitempty"`
	CardPaymentProcessor *string `json:"card_payment_processor,omitempty"`
	CompanyId *string `json:"company_id,omitempty"`
	DisplayQrOnHostedPay *bool `json:"display_qr_on_hosted_pay,omitempty"`
	HostedPayVersion *int `json:"hosted_pay_version,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsBlocked *bool `json:"is_blocked,omitempty"`
	IsExited *bool `json:"is_exited,omitempty"`
	IsSuspended *bool `json:"is_suspended,omitempty"`
	Jurisdiction *string `json:"jurisdiction,omitempty"`
	LogoUrlPng *string `json:"logo_url_png,omitempty"`
	LogoUrlSvg *string `json:"logo_url_svg,omitempty"`
	MerchantCategoryCode *string `json:"merchant_category_code,omitempty"`
	Name *string `json:"name,omitempty"`
	Note *string `json:"note,omitempty"`
	ParentMerchant *map[string]any `json:"parent_merchant,omitempty"`
	PaymentAccount *[]any `json:"payment_account,omitempty"`
	PaymentAccountLimit *int `json:"payment_account_limit,omitempty"`
	ShortName *string `json:"short_name,omitempty"`
	SupportedPaymentMethodsList *[]any `json:"supported_payment_methods_list,omitempty"`
	SuspensionReason *string `json:"suspension_reason,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TimeZoneId *string `json:"time_zone_id,omitempty"`
	TradingName *string `json:"trading_name,omitempty"`
	WebHookLimit *int `json:"web_hook_limit,omitempty"`
	YourRoleName *string `json:"your_role_name,omitempty"`
}

// NoFrixionMoneyMoovModelsMerchantUpdateData is the typed request payload for NoFrixionMoneyMoovModelsMerchant.UpdateTyped.
type NoFrixionMoneyMoovModelsMerchantUpdateData struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsMerchantPage is the typed data model for the no_frixion_money_moov_models_merchant_page entity.
type NoFrixionMoneyMoovModelsMerchantPage struct {
	AccountCurrency *[]any `json:"account_currency,omitempty"`
	CanHaveTrustAccount *bool `json:"can_have_trust_account,omitempty"`
	CardPaymentProcessor *string `json:"card_payment_processor,omitempty"`
	CompanyId *string `json:"company_id,omitempty"`
	DisplayQrOnHostedPay *bool `json:"display_qr_on_hosted_pay,omitempty"`
	HostedPayVersion *int `json:"hosted_pay_version,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsBlocked *bool `json:"is_blocked,omitempty"`
	IsExited *bool `json:"is_exited,omitempty"`
	IsSuspended *bool `json:"is_suspended,omitempty"`
	Jurisdiction *string `json:"jurisdiction,omitempty"`
	LogoUrlPng *string `json:"logo_url_png,omitempty"`
	LogoUrlSvg *string `json:"logo_url_svg,omitempty"`
	MerchantCategoryCode *string `json:"merchant_category_code,omitempty"`
	Name *string `json:"name,omitempty"`
	Note *string `json:"note,omitempty"`
	ParentMerchant *map[string]any `json:"parent_merchant,omitempty"`
	PaymentAccount *[]any `json:"payment_account,omitempty"`
	PaymentAccountLimit *int `json:"payment_account_limit,omitempty"`
	ShortName *string `json:"short_name,omitempty"`
	SupportedPaymentMethodsList *[]any `json:"supported_payment_methods_list,omitempty"`
	SuspensionReason *string `json:"suspension_reason,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TimeZoneId *string `json:"time_zone_id,omitempty"`
	TradingName *string `json:"trading_name,omitempty"`
	WebHookLimit *int `json:"web_hook_limit,omitempty"`
	YourRoleName *string `json:"your_role_name,omitempty"`
}

// NoFrixionMoneyMoovModelsMerchantPageListMatch is the typed request payload for NoFrixionMoneyMoovModelsMerchantPage.ListTyped.
type NoFrixionMoneyMoovModelsMerchantPageListMatch struct {
	MerchantId *string `json:"merchant_id,omitempty"`
}

// NoFrixionMoneyMoovModelsMerchantPayByBankSetting is the typed data model for the no_frixion_money_moov_models_merchant_pay_by_bank_setting entity.
type NoFrixionMoneyMoovModelsMerchantPayByBankSetting struct {
	BankCountryCode *[]any `json:"bank_country_code,omitempty"`
	BankId *string `json:"bank_id,omitempty"`
	BankName *string `json:"bank_name,omitempty"`
	BusinessInstitutionId *string `json:"business_institution_id,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Message *string `json:"message,omitempty"`
	MessageImageUrl *string `json:"message_image_url,omitempty"`
	Order *int `json:"order,omitempty"`
	PersonalInstitutionId *string `json:"personal_institution_id,omitempty"`
	Processor *string `json:"processor,omitempty"`
	WarningHeading *string `json:"warning_heading,omitempty"`
	WarningMessage *string `json:"warning_message,omitempty"`
}

// NoFrixionMoneyMoovModelsMerchantPayByBankSettingListMatch is the typed request payload for NoFrixionMoneyMoovModelsMerchantPayByBankSetting.ListTyped.
type NoFrixionMoneyMoovModelsMerchantPayByBankSettingListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsMerchantToken is the typed data model for the no_frixion_money_moov_models_merchant_token entity.
type NoFrixionMoneyMoovModelsMerchantToken struct {
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	HmacAlgorithm *string `json:"hmac_algorithm,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IpAddressWhitelist *string `json:"ip_address_whitelist,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	LastAuthorised *string `json:"last_authorised,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Nonce string `json:"nonce"`
	PermissionType *[]any `json:"permission_type,omitempty"`
	RequestSignatureVersion *int `json:"request_signature_version,omitempty"`
	SharedSecretAlgorithm *string `json:"shared_secret_algorithm,omitempty"`
	SharedSecretBase64 *string `json:"shared_secret_base64,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NoFrixionMoneyMoovModelsMerchantTokenLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsMerchantToken.LoadTyped.
type NoFrixionMoneyMoovModelsMerchantTokenLoadMatch struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsMerchantTokenCreateData is the typed request payload for NoFrixionMoneyMoovModelsMerchantToken.CreateTyped.
type NoFrixionMoneyMoovModelsMerchantTokenCreateData struct {
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	HmacAlgorithm *string `json:"hmac_algorithm,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IpAddressWhitelist *string `json:"ip_address_whitelist,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	LastAuthorised *string `json:"last_authorised,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Nonce string `json:"nonce"`
	PermissionType *[]any `json:"permission_type,omitempty"`
	RequestSignatureVersion *int `json:"request_signature_version,omitempty"`
	SharedSecretAlgorithm *string `json:"shared_secret_algorithm,omitempty"`
	SharedSecretBase64 *string `json:"shared_secret_base64,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NoFrixionMoneyMoovModelsMerchantTokenUpdateData is the typed request payload for NoFrixionMoneyMoovModelsMerchantToken.UpdateTyped.
type NoFrixionMoneyMoovModelsMerchantTokenUpdateData struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsMerchantTokenPage is the typed data model for the no_frixion_money_moov_models_merchant_token_page entity.
type NoFrixionMoneyMoovModelsMerchantTokenPage struct {
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IpAddressWhitelist *string `json:"ip_address_whitelist,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	LastAuthorised *string `json:"last_authorised,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Nonce string `json:"nonce"`
	PermissionType *[]any `json:"permission_type,omitempty"`
	RequestSignatureVersion *int `json:"request_signature_version,omitempty"`
	SharedSecretAlgorithm *string `json:"shared_secret_algorithm,omitempty"`
	SharedSecretBase64 *string `json:"shared_secret_base64,omitempty"`
	Token *string `json:"token,omitempty"`
}

// NoFrixionMoneyMoovModelsMerchantTokenPageListMatch is the typed request payload for NoFrixionMoneyMoovModelsMerchantTokenPage.ListTyped.
type NoFrixionMoneyMoovModelsMerchantTokenPageListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsNoFrixionVersion is the typed data model for the no_frixion_money_moov_models_no_frixion_version entity.
type NoFrixionMoneyMoovModelsNoFrixionVersion struct {
	BuildVersion *int `json:"build_version,omitempty"`
	MajorVersion *int `json:"major_version,omitempty"`
	MinorVersion *int `json:"minor_version,omitempty"`
	ReleaseName *string `json:"release_name,omitempty"`
}

// NoFrixionMoneyMoovModelsNoFrixionVersionLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsNoFrixionVersion.LoadTyped.
type NoFrixionMoneyMoovModelsNoFrixionVersionLoadMatch struct {
	BuildVersion *int `json:"build_version,omitempty"`
	MajorVersion *int `json:"major_version,omitempty"`
	MinorVersion *int `json:"minor_version,omitempty"`
	ReleaseName *string `json:"release_name,omitempty"`
}

// NoFrixionMoneyMoovModelsOpenBankingAccount is the typed data model for the no_frixion_money_moov_models_open_banking_account entity.
type NoFrixionMoneyMoovModelsOpenBankingAccount struct {
	AccountBalance *[]any `json:"account_balance,omitempty"`
	AccountIdentification *[]any `json:"account_identification,omitempty"`
	AccountName *[]any `json:"account_name,omitempty"`
	AccountType *string `json:"account_type,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	ConsolidatedAccountInformation *map[string]any `json:"consolidated_account_information,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	Detail *string `json:"detail,omitempty"`
	Id *string `json:"id,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	Type *string `json:"type,omitempty"`
	UsageType *string `json:"usage_type,omitempty"`
}

// NoFrixionMoneyMoovModelsOpenBankingAccountLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsOpenBankingAccount.LoadTyped.
type NoFrixionMoneyMoovModelsOpenBankingAccountLoadMatch struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsOpenBankingConsent is the typed data model for the no_frixion_money_moov_models_open_banking_consent entity.
type NoFrixionMoneyMoovModelsOpenBankingConsent struct {
	AuthorisationUrl *string `json:"authorisation_url,omitempty"`
	CallbackUrl *string `json:"callback_url,omitempty"`
	ConsentId *string `json:"consent_id,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	ExpiryDate *string `json:"expiry_date,omitempty"`
	FailureCallbackUrl *string `json:"failure_callback_url,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InstitutionId *string `json:"institution_id,omitempty"`
	IsConnectedAccount *bool `json:"is_connected_account,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Provider *string `json:"provider,omitempty"`
	SuccessWebHookUrl *string `json:"success_web_hook_url,omitempty"`
}

// NoFrixionMoneyMoovModelsOpenBankingConsentLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsOpenBankingConsent.LoadTyped.
type NoFrixionMoneyMoovModelsOpenBankingConsentLoadMatch struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsOpenBankingConsentListMatch is the typed request payload for NoFrixionMoneyMoovModelsOpenBankingConsent.ListTyped.
type NoFrixionMoneyMoovModelsOpenBankingConsentListMatch struct {
	Email string `json:"email"`
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsOpenBankingConsentCreateData is the typed request payload for NoFrixionMoneyMoovModelsOpenBankingConsent.CreateTyped.
type NoFrixionMoneyMoovModelsOpenBankingConsentCreateData struct {
	AuthorisationUrl *string `json:"authorisation_url,omitempty"`
	CallbackUrl *string `json:"callback_url,omitempty"`
	ConsentId *string `json:"consent_id,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	ExpiryDate *string `json:"expiry_date,omitempty"`
	FailureCallbackUrl *string `json:"failure_callback_url,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InstitutionId *string `json:"institution_id,omitempty"`
	IsConnectedAccount *bool `json:"is_connected_account,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Provider *string `json:"provider,omitempty"`
	SuccessWebHookUrl *string `json:"success_web_hook_url,omitempty"`
}

// NoFrixionMoneyMoovModelsOpenBankingConsentUpdateData is the typed request payload for NoFrixionMoneyMoovModelsOpenBankingConsent.UpdateTyped.
type NoFrixionMoneyMoovModelsOpenBankingConsentUpdateData struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsOpenBankingConsentRemoveMatch is the typed request payload for NoFrixionMoneyMoovModelsOpenBankingConsent.RemoveTyped.
type NoFrixionMoneyMoovModelsOpenBankingConsentRemoveMatch struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsOpenBankingTransaction is the typed data model for the no_frixion_money_moov_models_open_banking_transaction entity.
type NoFrixionMoneyMoovModelsOpenBankingTransaction struct {
	AddressDetail *map[string]any `json:"address_detail,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	Balance *map[string]any `json:"balance,omitempty"`
	BookingDateTime *string `json:"booking_date_time,omitempty"`
	ChargeDetail *map[string]any `json:"charge_detail,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrencyExchange *map[string]any `json:"currency_exchange,omitempty"`
	Date *string `json:"date,omitempty"`
	Description *string `json:"description,omitempty"`
	Enrichment *map[string]any `json:"enrichment,omitempty"`
	GrossAmount map[string]any `json:"gross_amount"`
	Id *string `json:"id,omitempty"`
	IsoBankTransactionCode *map[string]any `json:"iso_bank_transaction_code,omitempty"`
	Merchant *map[string]any `json:"merchant,omitempty"`
	PayeeDetail map[string]any `json:"payee_detail"`
	PayerDetail map[string]any `json:"payer_detail"`
	ProprietaryBankTransactionCode *map[string]any `json:"proprietary_bank_transaction_code,omitempty"`
	Reference *string `json:"reference,omitempty"`
	StatementReference *[]any `json:"statement_reference,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplementaryData *any `json:"supplementary_data,omitempty"`
	TransactionAmount map[string]any `json:"transaction_amount"`
	TransactionInformation *[]any `json:"transaction_information,omitempty"`
	TransactionMutability *string `json:"transaction_mutability,omitempty"`
	ValueDateTime *string `json:"value_date_time,omitempty"`
}

// NoFrixionMoneyMoovModelsOpenBankingTransactionListMatch is the typed request payload for NoFrixionMoneyMoovModelsOpenBankingTransaction.ListTyped.
type NoFrixionMoneyMoovModelsOpenBankingTransactionListMatch struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsPayment is the typed data model for the no_frixion_money_moov_models_payment entity.
type NoFrixionMoneyMoovModelsPayment struct {
	Address *[]any `json:"address,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amount_pending,omitempty"`
	AmountReceived *float64 `json:"amount_received,omitempty"`
	AmountRefunded *float64 `json:"amount_refunded,omitempty"`
	AutoSendReceipt *bool `json:"auto_send_receipt,omitempty"`
	BaseOriginUrl *string `json:"base_origin_url,omitempty"`
	CallbackUrl *string `json:"callback_url,omitempty"`
	CardAuthorizeOnly *bool `json:"card_authorize_only,omitempty"`
	CardCreateToken *bool `json:"card_create_token,omitempty"`
	CardCreateTokenMode *string `json:"card_create_token_mode,omitempty"`
	CardIgnoreCvn *bool `json:"card_ignore_cvn,omitempty"`
	CardNoPayerAuthentication *bool `json:"card_no_payer_authentication,omitempty"`
	CardProcessorMerchantId *string `json:"card_processor_merchant_id,omitempty"`
	CardStripePaymentIntentId *string `json:"card_stripe_payment_intent_id,omitempty"`
	CardStripePaymentIntentSecret *string `json:"card_stripe_payment_intent_secret,omitempty"`
	CardTransmitRawDetail *bool `json:"card_transmit_raw_detail,omitempty"`
	CreatedByUser map[string]any `json:"created_by_user"`
	Currency *string `json:"currency,omitempty"`
	CustomField *[]any `json:"custom_field,omitempty"`
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	CustomerName *string `json:"customer_name,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destination_account,omitempty"`
	DirectDebitPayment *map[string]any `json:"direct_debit_payment,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FailureCallbackUrl *string `json:"failure_callback_url,omitempty"`
	FieldDisplaySetting *[]any `json:"field_display_setting,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hosted_pay_checkout_url,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignore_address_verification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"inserted_sortable,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	LightningInvoice *string `json:"lightning_invoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightning_invoice_expires_at,omitempty"`
	MerchantDirectDebitMandateId *string `json:"merchant_direct_debit_mandate_id,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	NotificationEmailAddress *string `json:"notification_email_address,omitempty"`
	NotificationRoleID *[]any `json:"notification_role_i_d,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	PartialPaymentMethod *string `json:"partial_payment_method,omitempty"`
	PartialPaymentStep *string `json:"partial_payment_step,omitempty"`
	PaymentAttempt *[]any `json:"payment_attempt,omitempty"`
	PaymentMethod *[]any `json:"payment_method,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PispAccountId *string `json:"pisp_account_id,omitempty"`
	PriorityBankId *string `json:"priority_bank_id,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSecond *int `json:"sandbox_settle_delay_in_second,omitempty"`
	ShippingAddress *map[string]any `json:"shipping_address,omitempty"`
	ShippingAddressCity *string `json:"shipping_address_city,omitempty"`
	ShippingAddressCountryCode *string `json:"shipping_address_country_code,omitempty"`
	ShippingAddressCounty *string `json:"shipping_address_county,omitempty"`
	ShippingAddressLine1 *string `json:"shipping_address_line1,omitempty"`
	ShippingAddressLine2 *string `json:"shipping_address_line2,omitempty"`
	ShippingAddressPostCode *string `json:"shipping_address_post_code,omitempty"`
	ShippingEmail *string `json:"shipping_email,omitempty"`
	ShippingFirstName *string `json:"shipping_first_name,omitempty"`
	ShippingLastName *string `json:"shipping_last_name,omitempty"`
	ShippingPhone *string `json:"shipping_phone,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"success_web_hook_url,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TagId *[]any `json:"tag_id,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCard *[]any `json:"tokenised_card,omitempty"`
	Transaction *[]any `json:"transaction,omitempty"`
	UseHostedPaymentPage *bool `json:"use_hosted_payment_page,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsPayment.LoadTyped.
type NoFrixionMoneyMoovModelsPaymentLoadMatch struct {
	Id *string `json:"id,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentCreateData is the typed request payload for NoFrixionMoneyMoovModelsPayment.CreateTyped.
type NoFrixionMoneyMoovModelsPaymentCreateData struct {
	Address *[]any `json:"address,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amount_pending,omitempty"`
	AmountReceived *float64 `json:"amount_received,omitempty"`
	AmountRefunded *float64 `json:"amount_refunded,omitempty"`
	AutoSendReceipt *bool `json:"auto_send_receipt,omitempty"`
	BaseOriginUrl *string `json:"base_origin_url,omitempty"`
	CallbackUrl *string `json:"callback_url,omitempty"`
	CardAuthorizeOnly *bool `json:"card_authorize_only,omitempty"`
	CardCreateToken *bool `json:"card_create_token,omitempty"`
	CardCreateTokenMode *string `json:"card_create_token_mode,omitempty"`
	CardIgnoreCvn *bool `json:"card_ignore_cvn,omitempty"`
	CardNoPayerAuthentication *bool `json:"card_no_payer_authentication,omitempty"`
	CardProcessorMerchantId *string `json:"card_processor_merchant_id,omitempty"`
	CardStripePaymentIntentId *string `json:"card_stripe_payment_intent_id,omitempty"`
	CardStripePaymentIntentSecret *string `json:"card_stripe_payment_intent_secret,omitempty"`
	CardTransmitRawDetail *bool `json:"card_transmit_raw_detail,omitempty"`
	CreatedByUser map[string]any `json:"created_by_user"`
	Currency *string `json:"currency,omitempty"`
	CustomField *[]any `json:"custom_field,omitempty"`
	CustomerEmailAddress *string `json:"customer_email_address,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	CustomerName *string `json:"customer_name,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destination_account,omitempty"`
	DirectDebitPayment *map[string]any `json:"direct_debit_payment,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FailureCallbackUrl *string `json:"failure_callback_url,omitempty"`
	FieldDisplaySetting *[]any `json:"field_display_setting,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hosted_pay_checkout_url,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignore_address_verification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"inserted_sortable,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	LightningInvoice *string `json:"lightning_invoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightning_invoice_expires_at,omitempty"`
	MerchantDirectDebitMandateId *string `json:"merchant_direct_debit_mandate_id,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	NotificationEmailAddress *string `json:"notification_email_address,omitempty"`
	NotificationRoleID *[]any `json:"notification_role_i_d,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	PartialPaymentMethod *string `json:"partial_payment_method,omitempty"`
	PartialPaymentStep *string `json:"partial_payment_step,omitempty"`
	PaymentAttempt *[]any `json:"payment_attempt,omitempty"`
	PaymentMethod *[]any `json:"payment_method,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PispAccountId *string `json:"pisp_account_id,omitempty"`
	PriorityBankId *string `json:"priority_bank_id,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSecond *int `json:"sandbox_settle_delay_in_second,omitempty"`
	ShippingAddress *map[string]any `json:"shipping_address,omitempty"`
	ShippingAddressCity *string `json:"shipping_address_city,omitempty"`
	ShippingAddressCountryCode *string `json:"shipping_address_country_code,omitempty"`
	ShippingAddressCounty *string `json:"shipping_address_county,omitempty"`
	ShippingAddressLine1 *string `json:"shipping_address_line1,omitempty"`
	ShippingAddressLine2 *string `json:"shipping_address_line2,omitempty"`
	ShippingAddressPostCode *string `json:"shipping_address_post_code,omitempty"`
	ShippingEmail *string `json:"shipping_email,omitempty"`
	ShippingFirstName *string `json:"shipping_first_name,omitempty"`
	ShippingLastName *string `json:"shipping_last_name,omitempty"`
	ShippingPhone *string `json:"shipping_phone,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"success_web_hook_url,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TagId *[]any `json:"tag_id,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCard *[]any `json:"tokenised_card,omitempty"`
	Transaction *[]any `json:"transaction,omitempty"`
	UseHostedPaymentPage *bool `json:"use_hosted_payment_page,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentUpdateData is the typed request payload for NoFrixionMoneyMoovModelsPayment.UpdateTyped.
type NoFrixionMoneyMoovModelsPaymentUpdateData struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsPaymentAccountMinimalPage is the typed data model for the no_frixion_money_moov_models_payment_account_minimal_page entity.
type NoFrixionMoneyMoovModelsPaymentAccountMinimalPage struct {
	AccountName *string `json:"account_name,omitempty"`
	AvailableBalance *float64 `json:"available_balance,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnit *int `json:"balance_minor_unit,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsConnectedAccount *bool `json:"is_connected_account,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submitted_payouts_balance,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentAccountMinimalPageListMatch is the typed request payload for NoFrixionMoneyMoovModelsPaymentAccountMinimalPage.ListTyped.
type NoFrixionMoneyMoovModelsPaymentAccountMinimalPageListMatch struct {
	AccountName *string `json:"account_name,omitempty"`
	AvailableBalance *float64 `json:"available_balance,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnit *int `json:"balance_minor_unit,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier *map[string]any `json:"identifier,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsConnectedAccount *bool `json:"is_connected_account,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submitted_payouts_balance,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentAccountPage is the typed data model for the no_frixion_money_moov_models_payment_account_page entity.
type NoFrixionMoneyMoovModelsPaymentAccountPage struct {
	AccountName *string `json:"account_name,omitempty"`
	AccountSupplierName *string `json:"account_supplier_name,omitempty"`
	AvailableBalance *float64 `json:"available_balance,omitempty"`
	AvailableBalanceMinorUnit *int `json:"available_balance_minor_unit,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnit *int `json:"balance_minor_unit,omitempty"`
	BankName *string `json:"bank_name,omitempty"`
	ConsentId *string `json:"consent_id,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	CreatedByDisplayName *string `json:"created_by_display_name,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"default_payment_rail,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	ExpiryDate *string `json:"expiry_date,omitempty"`
	ExternalAccountIcon *string `json:"external_account_icon,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsConnectedAccount *bool `json:"is_connected_account,omitempty"`
	IsDefault *bool `json:"is_default,omitempty"`
	IsTrustAccount *bool `json:"is_trust_account,omitempty"`
	IsVirtual *bool `json:"is_virtual,omitempty"`
	LastTransaction *map[string]any `json:"last_transaction,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantName *string `json:"merchant_name,omitempty"`
	PhysicalAccountId *string `json:"physical_account_id,omitempty"`
	Rule *[]any `json:"rule,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submitted_payouts_balance,omitempty"`
	SubmittedPayoutsBalanceMinorUnit *int `json:"submitted_payouts_balance_minor_unit,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplier_sepa_instant_status,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xero_bank_feed_connection_status,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xero_bank_feed_last_synced_at,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xero_bank_feed_sync_last_failed_at,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xero_bank_feed_sync_last_failure_reason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xero_bank_feed_sync_status,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xero_unsynchronised_transactions_count,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentAccountPageListMatch is the typed request payload for NoFrixionMoneyMoovModelsPaymentAccountPage.ListTyped.
type NoFrixionMoneyMoovModelsPaymentAccountPageListMatch struct {
	AccountId *string `json:"account_id,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentInitiation is the typed data model for the no_frixion_money_moov_models_payment_initiation entity.
type NoFrixionMoneyMoovModelsPaymentInitiation struct {
	PaymentInitiationId *string `json:"payment_initiation_id,omitempty"`
	PaymentRequestCallbackUrl *string `json:"payment_request_callback_url,omitempty"`
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	RedirectUrl *string `json:"redirect_url,omitempty"`
	ResponseType *string `json:"response_type,omitempty"`
	SpecificErrorMessage *string `json:"specific_error_message,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentInitiationCreateData is the typed request payload for NoFrixionMoneyMoovModelsPaymentInitiation.CreateTyped.
type NoFrixionMoneyMoovModelsPaymentInitiationCreateData struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// NoFrixionMoneyMoovModelsPaymentRequestEvent is the typed data model for the no_frixion_money_moov_models_payment_request_event entity.
type NoFrixionMoneyMoovModelsPaymentRequestEvent struct {
	Amount float64 `json:"amount"`
	ApplePayTransactionId *string `json:"apple_pay_transaction_id,omitempty"`
	CardAuthorizationResponseId *string `json:"card_authorization_response_id,omitempty"`
	CardExpiryMonth *int `json:"card_expiry_month,omitempty"`
	CardExpiryYear *int `json:"card_expiry_year,omitempty"`
	CardIssuer *string `json:"card_issuer,omitempty"`
	CardIssuerCountry *string `json:"card_issuer_country,omitempty"`
	CardLastFourDigit *string `json:"card_last_four_digit,omitempty"`
	CardRequestId *string `json:"card_request_id,omitempty"`
	CardScheme *string `json:"card_scheme,omitempty"`
	CardTokenCustomerId *string `json:"card_token_customer_id,omitempty"`
	CardTransactionId *string `json:"card_transaction_id,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DirectDebitPaymentId *string `json:"direct_debit_payment_id,omitempty"`
	DirectDebitPaymentReference *string `json:"direct_debit_payment_reference,omitempty"`
	DrirectDebitMandateId *string `json:"drirect_debit_mandate_id,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
	ErrorReason *string `json:"error_reason,omitempty"`
	EventType *string `json:"event_type,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LightningInvoice *string `json:"lightning_invoice,omitempty"`
	LightningRHash *string `json:"lightning_r_hash,omitempty"`
	OriginUrl *string `json:"origin_url,omitempty"`
	PaymentMethodType *string `json:"payment_method_type,omitempty"`
	PaymentProcessorName *string `json:"payment_processor_name,omitempty"`
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	PispBankStatus *string `json:"pisp_bank_status,omitempty"`
	PispPaymentInitiationId *string `json:"pisp_payment_initiation_id,omitempty"`
	PispPaymentInstitutionName *string `json:"pisp_payment_institution_name,omitempty"`
	PispPaymentServiceProviderId *string `json:"pisp_payment_service_provider_id,omitempty"`
	PispRedirectUrl *string `json:"pisp_redirect_url,omitempty"`
	ReconciledTransactionId *string `json:"reconciled_transaction_id,omitempty"`
	RefundPayoutId *string `json:"refund_payout_id,omitempty"`
	Status *string `json:"status,omitempty"`
	WalletName *string `json:"wallet_name,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentRequestEventListMatch is the typed request payload for NoFrixionMoneyMoovModelsPaymentRequestEvent.ListTyped.
type NoFrixionMoneyMoovModelsPaymentRequestEventListMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// NoFrixionMoneyMoovModelsPaymentRequestMetric is the typed data model for the no_frixion_money_moov_models_payment_request_metric entity.
type NoFrixionMoneyMoovModelsPaymentRequestMetric struct {
	All *int `json:"all,omitempty"`
	Authorized *int `json:"authorized,omitempty"`
	Paid *int `json:"paid,omitempty"`
	PartiallyPaid *int `json:"partially_paid,omitempty"`
	TotalAmountsByCurrency *map[string]any `json:"total_amounts_by_currency,omitempty"`
	Unpaid *int `json:"unpaid,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentRequestMetricLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsPaymentRequestMetric.LoadTyped.
type NoFrixionMoneyMoovModelsPaymentRequestMetricLoadMatch struct {
	All *int `json:"all,omitempty"`
	Authorized *int `json:"authorized,omitempty"`
	Paid *int `json:"paid,omitempty"`
	PartiallyPaid *int `json:"partially_paid,omitempty"`
	TotalAmountsByCurrency *map[string]any `json:"total_amounts_by_currency,omitempty"`
	Unpaid *int `json:"unpaid,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentRequestMinimal is the typed data model for the no_frixion_money_moov_models_payment_request_minimal entity.
type NoFrixionMoneyMoovModelsPaymentRequestMinimal struct {
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amount_pending,omitempty"`
	AmountReceived *float64 `json:"amount_received,omitempty"`
	AmountRefunded *float64 `json:"amount_refunded,omitempty"`
	CallbackUrl *string `json:"callback_url,omitempty"`
	CardStripePaymentIntentSecret *string `json:"card_stripe_payment_intent_secret,omitempty"`
	CountryCode *string `json:"country_code,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomFieldsToDisplay *[]any `json:"custom_fields_to_display,omitempty"`
	Description *string `json:"description,omitempty"`
	DueDate *string `json:"due_date,omitempty"`
	FieldDisplaySetting *[]any `json:"field_display_setting,omitempty"`
	GooglePayMerchantId *string `json:"google_pay_merchant_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantLogoUrlPng *string `json:"merchant_logo_url_png,omitempty"`
	MerchantLogoUrlSvg *string `json:"merchant_logo_url_svg,omitempty"`
	MerchantName *string `json:"merchant_name,omitempty"`
	MerchantShortName *string `json:"merchant_short_name,omitempty"`
	PartialPaymentMethod *string `json:"partial_payment_method,omitempty"`
	PaymentAttempt *[]any `json:"payment_attempt,omitempty"`
	PaymentMethodsList *[]any `json:"payment_methods_list,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentProcessorKey *string `json:"payment_processor_key,omitempty"`
	PispError *string `json:"pisp_error,omitempty"`
	PriorityBankId *string `json:"priority_bank_id,omitempty"`
	Status *string `json:"status,omitempty"`
	StripeAccountId *string `json:"stripe_account_id,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentRequestMinimalListMatch is the typed request payload for NoFrixionMoneyMoovModelsPaymentRequestMinimal.ListTyped.
type NoFrixionMoneyMoovModelsPaymentRequestMinimalListMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// NoFrixionMoneyMoovModelsPaymentRequestResult is the typed data model for the no_frixion_money_moov_models_payment_request_result entity.
type NoFrixionMoneyMoovModelsPaymentRequestResult struct {
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amount_pending,omitempty"`
	AmountReceived *float64 `json:"amount_received,omitempty"`
	AmountRefunded *float64 `json:"amount_refunded,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomerId *string `json:"customer_id,omitempty"`
	Payment *[]any `json:"payment,omitempty"`
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	PispAuthorization *[]any `json:"pisp_authorization,omitempty"`
	RequestedAmount *float64 `json:"requested_amount,omitempty"`
	Result *string `json:"result,omitempty"`
}

// NoFrixionMoneyMoovModelsPaymentRequestResultListMatch is the typed request payload for NoFrixionMoneyMoovModelsPaymentRequestResult.ListTyped.
type NoFrixionMoneyMoovModelsPaymentRequestResultListMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment is the typed data model for the no_frixion_money_moov_models_payment_requests_merchant_payment entity.
type NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment struct {
	Description string `json:"description"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name string `json:"name"`
	Template map[string]any `json:"template"`
}

// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentListMatch is the typed request payload for NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment.ListTyped.
type NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 is the typed data model for the no_frixion_money_moov_models_payment_requests_merchant_payment2 entity.
type NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 struct {
	Description string `json:"description"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name string `json:"name"`
	Template map[string]any `json:"template"`
}

// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2LoadMatch is the typed request payload for NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2.LoadTyped.
type NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2LoadMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
	TemplateId string `json:"template_id"`
}

// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 is the typed data model for the no_frixion_money_moov_models_payment_requests_merchant_payment3 entity.
type NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 struct {
	Description string `json:"description"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name string `json:"name"`
	Template map[string]any `json:"template"`
}

// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3UpdateData is the typed request payload for NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3.UpdateTyped.
type NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3UpdateData struct {
	PaymentrequestId string `json:"paymentrequest_id"`
	TemplateId string `json:"template_id"`
}

// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 is the typed data model for the no_frixion_money_moov_models_payment_requests_merchant_payment4 entity.
type NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 struct {
}

// NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4RemoveMatch is the typed request payload for NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4.RemoveTyped.
type NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4RemoveMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
	TemplateId string `json:"template_id"`
}

// NoFrixionMoneyMoovModelsPayoutKeysetPage is the typed data model for the no_frixion_money_moov_models_payout_keyset_page entity.
type NoFrixionMoneyMoovModelsPayoutKeysetPage struct {
	AccountId *string `json:"account_id,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	ApprovePayoutUrl *string `json:"approve_payout_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanProcess *bool `json:"can_process,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	ChargeBearer *string `json:"charge_bearer,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserId *string `json:"current_user_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Document *[]any `json:"document,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formatted_fx_destination_amount,omitempty"`
	FormattedSchedule *string `json:"formatted_schedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formatted_schedule_day_only,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formatted_source_account_available_balance,omitempty"`
	FxDestinationAmount *float64 `json:"fx_destination_amount,omitempty"`
	FxDestinationAmountMinorUnit *int `json:"fx_destination_amount_minor_unit,omitempty"`
	FxDestinationCurrency *string `json:"fx_destination_currency,omitempty"`
	FxQuoteExpiresAt *string `json:"fx_quote_expires_at,omitempty"`
	FxQuoteId *string `json:"fx_quote_id,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	FxUseDestinationAmount *bool `json:"fx_use_destination_amount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsFailed *bool `json:"is_failed,omitempty"`
	IsSettled *bool `json:"is_settled,omitempty"`
	IsSubmitted *bool `json:"is_submitted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentRail *string `json:"payment_rail,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PayrunName *string `json:"payrun_name,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"source_account_available_balance,omitempty"`
	SourceAccountAvailableBalanceMinorUnit *int `json:"source_account_available_balance_minor_unit,omitempty"`
	SourceAccountBic *string `json:"source_account_bic,omitempty"`
	SourceAccountCurrency *string `json:"source_account_currency,omitempty"`
	SourceAccountIban *string `json:"source_account_iban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"source_account_identifier"`
	SourceAccountName *string `json:"source_account_name,omitempty"`
	SourceAccountNumber *string `json:"source_account_number,omitempty"`
	SourceAccountSortcode *string `json:"source_account_sortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TopupPayrunId *string `json:"topup_payrun_id,omitempty"`
	TransactedAmount *float64 `json:"transacted_amount,omitempty"`
	TransactedFxAmount *float64 `json:"transacted_fx_amount,omitempty"`
	TransactedFxRate *float64 `json:"transacted_fx_rate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// NoFrixionMoneyMoovModelsPayoutKeysetPageListMatch is the typed request payload for NoFrixionMoneyMoovModelsPayoutKeysetPage.ListTyped.
type NoFrixionMoneyMoovModelsPayoutKeysetPageListMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
}

// NoFrixionMoneyMoovModelsPayoutMetric is the typed data model for the no_frixion_money_moov_models_payout_metric entity.
type NoFrixionMoneyMoovModelsPayoutMetric struct {
	All *float64 `json:"all,omitempty"`
	Failed *float64 `json:"failed,omitempty"`
	InProgress *float64 `json:"in_progress,omitempty"`
	Paid *float64 `json:"paid,omitempty"`
	PendingApproval *float64 `json:"pending_approval,omitempty"`
	Scheduled *float64 `json:"scheduled,omitempty"`
	TotalAmountsByCurrency *map[string]any `json:"total_amounts_by_currency,omitempty"`
}

// NoFrixionMoneyMoovModelsPayoutMetricLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsPayoutMetric.LoadTyped.
type NoFrixionMoneyMoovModelsPayoutMetricLoadMatch struct {
	All *float64 `json:"all,omitempty"`
	Failed *float64 `json:"failed,omitempty"`
	InProgress *float64 `json:"in_progress,omitempty"`
	Paid *float64 `json:"paid,omitempty"`
	PendingApproval *float64 `json:"pending_approval,omitempty"`
	Scheduled *float64 `json:"scheduled,omitempty"`
	TotalAmountsByCurrency *map[string]any `json:"total_amounts_by_currency,omitempty"`
}

// NoFrixionMoneyMoovModelsPayoutsPayoutsCreate is the typed data model for the no_frixion_money_moov_models_payouts_payouts_create entity.
type NoFrixionMoneyMoovModelsPayoutsPayoutsCreate struct {
	FailedPayout *map[string]any `json:"failed_payout,omitempty"`
	Payout *[]any `json:"payout,omitempty"`
}

// NoFrixionMoneyMoovModelsPayoutsPayoutsCreateCreateData is the typed request payload for NoFrixionMoneyMoovModelsPayoutsPayoutsCreate.CreateTyped.
type NoFrixionMoneyMoovModelsPayoutsPayoutsCreateCreateData struct {
	FailedPayout *map[string]any `json:"failed_payout,omitempty"`
	Payout *[]any `json:"payout,omitempty"`
}

// NoFrixionMoneyMoovModelsPayrun is the typed data model for the no_frixion_money_moov_models_payrun entity.
type NoFrixionMoneyMoovModelsPayrun struct {
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisationDate *string `json:"authorisation_date,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanDelete *bool `json:"can_delete,omitempty"`
	CanEdit *bool `json:"can_edit,omitempty"`
	Event *[]any `json:"event,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	Invoice *[]any `json:"invoice,omitempty"`
	InvoicesMinimal *[]any `json:"invoices_minimal,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	LastUpdatedBy map[string]any `json:"last_updated_by"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Payment *[]any `json:"payment,omitempty"`
	Payout *[]any `json:"payout,omitempty"`
	PayoutsCount *int `json:"payouts_count,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	ScheduledDate *string `json:"scheduled_date,omitempty"`
	SourceAccount *[]any `json:"source_account,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalEur *float64 `json:"total_eur,omitempty"`
	TotalGbp *float64 `json:"total_gbp,omitempty"`
	TotalUsd *float64 `json:"total_usd,omitempty"`
}

// NoFrixionMoneyMoovModelsPayrunLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsPayrun.LoadTyped.
type NoFrixionMoneyMoovModelsPayrunLoadMatch struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsPayrunCreateData is the typed request payload for NoFrixionMoneyMoovModelsPayrun.CreateTyped.
type NoFrixionMoneyMoovModelsPayrunCreateData struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsPayrunUpdateData is the typed request payload for NoFrixionMoneyMoovModelsPayrun.UpdateTyped.
type NoFrixionMoneyMoovModelsPayrunUpdateData struct {
	Id *string `json:"id,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
}

// NoFrixionMoneyMoovModelsReportResult is the typed data model for the no_frixion_money_moov_models_report_result entity.
type NoFrixionMoneyMoovModelsReportResult struct {
	Content *string `json:"content,omitempty"`
	ContentType *string `json:"content_type,omitempty"`
	LastCompletedAt *string `json:"last_completed_at,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	ReportName *string `json:"report_name,omitempty"`
	ReportType *string `json:"report_type,omitempty"`
	StatementNumber *int `json:"statement_number,omitempty"`
}

// NoFrixionMoneyMoovModelsReportResultLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsReportResult.LoadTyped.
type NoFrixionMoneyMoovModelsReportResultLoadMatch struct {
	Id int `json:"id"`
	ReportId string `json:"report_id"`
}

// NoFrixionMoneyMoovModelsRule is the typed data model for the no_frixion_money_moov_models_rule entity.
type NoFrixionMoneyMoovModelsRule struct {
	Account *map[string]any `json:"account,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
	ApproveUrl *string `json:"approve_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	Description *string `json:"description,omitempty"`
	EndAt *string `json:"end_at,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsDisabled *bool `json:"is_disabled,omitempty"`
	LastExecutedAt *string `json:"last_executed_at,omitempty"`
	LastRunAtTransactionDate *string `json:"last_run_at_transaction_date,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce string `json:"nonce"`
	OnApprovedWebHookUrl *string `json:"on_approved_web_hook_url,omitempty"`
	OnExecutionErrorWebHookUrl *string `json:"on_execution_error_web_hook_url,omitempty"`
	OnExecutionSuccessWebHookUrl *string `json:"on_execution_success_web_hook_url,omitempty"`
	StartAt *string `json:"start_at,omitempty"`
	Status *string `json:"status,omitempty"`
	SweepAction *map[string]any `json:"sweep_action,omitempty"`
	TimeZoneId *string `json:"time_zone_id,omitempty"`
	TriggerCronExpression *string `json:"trigger_cron_expression,omitempty"`
	TriggerOnPayIn *bool `json:"trigger_on_pay_in,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	WebHookSecret *string `json:"web_hook_secret,omitempty"`
}

// NoFrixionMoneyMoovModelsRuleLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsRule.LoadTyped.
type NoFrixionMoneyMoovModelsRuleLoadMatch struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsRuleCreateData is the typed request payload for NoFrixionMoneyMoovModelsRule.CreateTyped.
type NoFrixionMoneyMoovModelsRuleCreateData struct {
	Account *map[string]any `json:"account,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
	ApproveUrl *string `json:"approve_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	Description *string `json:"description,omitempty"`
	EndAt *string `json:"end_at,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsDisabled *bool `json:"is_disabled,omitempty"`
	LastExecutedAt *string `json:"last_executed_at,omitempty"`
	LastRunAtTransactionDate *string `json:"last_run_at_transaction_date,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce string `json:"nonce"`
	OnApprovedWebHookUrl *string `json:"on_approved_web_hook_url,omitempty"`
	OnExecutionErrorWebHookUrl *string `json:"on_execution_error_web_hook_url,omitempty"`
	OnExecutionSuccessWebHookUrl *string `json:"on_execution_success_web_hook_url,omitempty"`
	StartAt *string `json:"start_at,omitempty"`
	Status *string `json:"status,omitempty"`
	SweepAction *map[string]any `json:"sweep_action,omitempty"`
	TimeZoneId *string `json:"time_zone_id,omitempty"`
	TriggerCronExpression *string `json:"trigger_cron_expression,omitempty"`
	TriggerOnPayIn *bool `json:"trigger_on_pay_in,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	WebHookSecret *string `json:"web_hook_secret,omitempty"`
}

// NoFrixionMoneyMoovModelsRuleUpdateData is the typed request payload for NoFrixionMoneyMoovModelsRule.UpdateTyped.
type NoFrixionMoneyMoovModelsRuleUpdateData struct {
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsTransaction is the typed data model for the no_frixion_money_moov_models_transaction entity.
type NoFrixionMoneyMoovModelsTransaction struct {
	AccountId *string `json:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	AccountSequenceNumber *int `json:"account_sequence_number,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnit *int `json:"balance_minor_unit,omitempty"`
	Counterparty *map[string]any `json:"counterparty,omitempty"`
	CounterpartySummary *string `json:"counterparty_summary,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	FxAmount *float64 `json:"fx_amount,omitempty"`
	FxCurrency *string `json:"fx_currency,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	PaymentRequestCustomField *map[string]any `json:"payment_request_custom_field,omitempty"`
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	PayoutId *string `json:"payout_id,omitempty"`
	RawReference *string `json:"raw_reference,omitempty"`
	RuleId *string `json:"rule_id,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TransactionDate *string `json:"transaction_date,omitempty"`
	Type *string `json:"type,omitempty"`
	VirtualIban *string `json:"virtual_iban,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// NoFrixionMoneyMoovModelsTransactionLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsTransaction.LoadTyped.
type NoFrixionMoneyMoovModelsTransactionLoadMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	Id string `json:"id"`
}

// NoFrixionMoneyMoovModelsTransactionPage is the typed data model for the no_frixion_money_moov_models_transaction_page entity.
type NoFrixionMoneyMoovModelsTransactionPage struct {
	AccountId *string `json:"account_id,omitempty"`
	AccountName *string `json:"account_name,omitempty"`
	AccountSequenceNumber *int `json:"account_sequence_number,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnit *int `json:"balance_minor_unit,omitempty"`
	Content *[]any `json:"content,omitempty"`
	Counterparty *map[string]any `json:"counterparty,omitempty"`
	CounterpartySummary *string `json:"counterparty_summary,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Description *string `json:"description,omitempty"`
	FxAmount *float64 `json:"fx_amount,omitempty"`
	FxCurrency *string `json:"fx_currency,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	PageNumber *int `json:"page_number,omitempty"`
	PageSize *int `json:"page_size,omitempty"`
	PaymentRequestCustomField *map[string]any `json:"payment_request_custom_field,omitempty"`
	PaymentRequestId *string `json:"payment_request_id,omitempty"`
	PayoutId *string `json:"payout_id,omitempty"`
	RawReference *string `json:"raw_reference,omitempty"`
	RuleId *string `json:"rule_id,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TotalPage *int `json:"total_page,omitempty"`
	TotalSize *int `json:"total_size,omitempty"`
	TransactionDate *string `json:"transaction_date,omitempty"`
	Type *string `json:"type,omitempty"`
	VirtualIban *string `json:"virtual_iban,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// NoFrixionMoneyMoovModelsTransactionPageLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsTransactionPage.LoadTyped.
type NoFrixionMoneyMoovModelsTransactionPageLoadMatch struct {
	AccountId string `json:"account_id"`
}

// NoFrixionMoneyMoovModelsTransactionPageListMatch is the typed request payload for NoFrixionMoneyMoovModelsTransactionPage.ListTyped.
type NoFrixionMoneyMoovModelsTransactionPageListMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
}

// NoFrixionMoneyMoovModelsUserInvite is the typed data model for the no_frixion_money_moov_models_user_invite entity.
type NoFrixionMoneyMoovModelsUserInvite struct {
	AuthorisationStatus *map[string]any `json:"authorisation_status,omitempty"`
	Id *string `json:"id,omitempty"`
	InitialRoleId *string `json:"initial_role_id,omitempty"`
	InviteeEmailAddress *string `json:"invitee_email_address,omitempty"`
	InviteeFirstName *string `json:"invitee_first_name,omitempty"`
	InviteeLastName *string `json:"invitee_last_name,omitempty"`
	InviterEmailAddress *string `json:"inviter_email_address,omitempty"`
	InviterFirstName *string `json:"inviter_first_name,omitempty"`
	InviterLastName *string `json:"inviter_last_name,omitempty"`
	IsAuthorised *bool `json:"is_authorised,omitempty"`
	IsInviteeRegistered *bool `json:"is_invitee_registered,omitempty"`
	LastInvited *string `json:"last_invited,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantName *string `json:"merchant_name,omitempty"`
	Message *string `json:"message,omitempty"`
	RegistrationUrl *string `json:"registration_url,omitempty"`
	SendInviteEmail *bool `json:"send_invite_email,omitempty"`
	Status *string `json:"status,omitempty"`
	User map[string]any `json:"user"`
	UserId *string `json:"user_id,omitempty"`
}

// NoFrixionMoneyMoovModelsUserInviteLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsUserInvite.LoadTyped.
type NoFrixionMoneyMoovModelsUserInviteLoadMatch struct {
	Id *string `json:"id,omitempty"`
	UserinviteId *string `json:"userinvite_id,omitempty"`
}

// NoFrixionMoneyMoovModelsUserInviteCreateData is the typed request payload for NoFrixionMoneyMoovModelsUserInvite.CreateTyped.
type NoFrixionMoneyMoovModelsUserInviteCreateData struct {
	AuthorisationStatus *map[string]any `json:"authorisation_status,omitempty"`
	Id *string `json:"id,omitempty"`
	InitialRoleId *string `json:"initial_role_id,omitempty"`
	InviteeEmailAddress *string `json:"invitee_email_address,omitempty"`
	InviteeFirstName *string `json:"invitee_first_name,omitempty"`
	InviteeLastName *string `json:"invitee_last_name,omitempty"`
	InviterEmailAddress *string `json:"inviter_email_address,omitempty"`
	InviterFirstName *string `json:"inviter_first_name,omitempty"`
	InviterLastName *string `json:"inviter_last_name,omitempty"`
	IsAuthorised *bool `json:"is_authorised,omitempty"`
	IsInviteeRegistered *bool `json:"is_invitee_registered,omitempty"`
	LastInvited *string `json:"last_invited,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantName *string `json:"merchant_name,omitempty"`
	Message *string `json:"message,omitempty"`
	RegistrationUrl *string `json:"registration_url,omitempty"`
	SendInviteEmail *bool `json:"send_invite_email,omitempty"`
	Status *string `json:"status,omitempty"`
	User map[string]any `json:"user"`
	UserId *string `json:"user_id,omitempty"`
}

// NoFrixionMoneyMoovModelsUserInvitePage is the typed data model for the no_frixion_money_moov_models_user_invite_page entity.
type NoFrixionMoneyMoovModelsUserInvitePage struct {
	AuthorisationStatus *map[string]any `json:"authorisation_status,omitempty"`
	Id *string `json:"id,omitempty"`
	InitialRoleId *string `json:"initial_role_id,omitempty"`
	InviteeEmailAddress *string `json:"invitee_email_address,omitempty"`
	InviteeFirstName *string `json:"invitee_first_name,omitempty"`
	InviteeLastName *string `json:"invitee_last_name,omitempty"`
	InviterEmailAddress *string `json:"inviter_email_address,omitempty"`
	InviterFirstName *string `json:"inviter_first_name,omitempty"`
	InviterLastName *string `json:"inviter_last_name,omitempty"`
	IsAuthorised *bool `json:"is_authorised,omitempty"`
	IsInviteeRegistered *bool `json:"is_invitee_registered,omitempty"`
	LastInvited *string `json:"last_invited,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantName *string `json:"merchant_name,omitempty"`
	Message *string `json:"message,omitempty"`
	RegistrationUrl *string `json:"registration_url,omitempty"`
	Status *string `json:"status,omitempty"`
	User map[string]any `json:"user"`
	UserId *string `json:"user_id,omitempty"`
}

// NoFrixionMoneyMoovModelsUserInvitePageListMatch is the typed request payload for NoFrixionMoneyMoovModelsUserInvitePage.ListTyped.
type NoFrixionMoneyMoovModelsUserInvitePageListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsUserPage is the typed data model for the no_frixion_money_moov_models_user_page entity.
type NoFrixionMoneyMoovModelsUserPage struct {
	ClientSessionTimeout *[]any `json:"client_session_timeout,omitempty"`
	EmailAddress string `json:"email_address"`
	FirstName string `json:"first_name"`
	Id *string `json:"id,omitempty"`
	LastName string `json:"last_name"`
	PasskeyAdded *bool `json:"passkey_added,omitempty"`
	Permission *map[string]any `json:"permission,omitempty"`
	RolesWithScope *[]any `json:"roles_with_scope,omitempty"`
	TwoFactorEnabled *bool `json:"two_factor_enabled,omitempty"`
}

// NoFrixionMoneyMoovModelsUserPageListMatch is the typed request payload for NoFrixionMoneyMoovModelsUserPage.ListTyped.
type NoFrixionMoneyMoovModelsUserPageListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsWebhook is the typed data model for the no_frixion_money_moov_models_webhook entity.
type NoFrixionMoneyMoovModelsWebhook struct {
	DestinationUrl *string `json:"destination_url,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	FailedNotificationEmailAddress *string `json:"failed_notification_email_address,omitempty"`
	Id *string `json:"id,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	NotificationMethod *string `json:"notification_method,omitempty"`
	ResourceType *[]any `json:"resource_type,omitempty"`
	Retry *bool `json:"retry,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Version *int `json:"version,omitempty"`
}

// NoFrixionMoneyMoovModelsWebhookLoadMatch is the typed request payload for NoFrixionMoneyMoovModelsWebhook.LoadTyped.
type NoFrixionMoneyMoovModelsWebhookLoadMatch struct {
	Id string `json:"id"`
	MerchantId *string `json:"merchant_id,omitempty"`
}

// NoFrixionMoneyMoovModelsWebhookListMatch is the typed request payload for NoFrixionMoneyMoovModelsWebhook.ListTyped.
type NoFrixionMoneyMoovModelsWebhookListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// NoFrixionMoneyMoovModelsWebhookCreateData is the typed request payload for NoFrixionMoneyMoovModelsWebhook.CreateTyped.
type NoFrixionMoneyMoovModelsWebhookCreateData struct {
	DestinationUrl *string `json:"destination_url,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	FailedNotificationEmailAddress *string `json:"failed_notification_email_address,omitempty"`
	Id *string `json:"id,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	NotificationMethod *string `json:"notification_method,omitempty"`
	ResourceType *[]any `json:"resource_type,omitempty"`
	Retry *bool `json:"retry,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Version *int `json:"version,omitempty"`
}

// NoFrixionMoneyMoovModelsWebhookUpdateData is the typed request payload for NoFrixionMoneyMoovModelsWebhook.UpdateTyped.
type NoFrixionMoneyMoovModelsWebhookUpdateData struct {
	Id string `json:"id"`
}

// OpenBanking is the typed data model for the open_banking entity.
type OpenBanking struct {
}

// OpenBankingCreateData is the typed request payload for OpenBanking.CreateTyped.
type OpenBankingCreateData struct {
	AccountId string `json:"account_id"`
}

// OpenBankingRemoveMatch is the typed request payload for OpenBanking.RemoveTyped.
type OpenBankingRemoveMatch struct {
	Email *string `json:"email,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
}

// Payeeverification is the typed data model for the payeeverification entity.
type Payeeverification struct {
	AccountName string `json:"account_name"`
	AccountNumber *string `json:"account_number,omitempty"`
	Iban string `json:"iban"`
	PayeeVerifiedAccountName *string `json:"payee_verified_account_name,omitempty"`
	Result *string `json:"result,omitempty"`
	SecondaryIdentification *string `json:"secondary_identification,omitempty"`
	SortCode *string `json:"sort_code,omitempty"`
}

// PayeeverificationCreateData is the typed request payload for Payeeverification.CreateTyped.
type PayeeverificationCreateData struct {
	AccountName string `json:"account_name"`
	AccountNumber *string `json:"account_number,omitempty"`
	Iban string `json:"iban"`
	PayeeVerifiedAccountName *string `json:"payee_verified_account_name,omitempty"`
	Result *string `json:"result,omitempty"`
	SecondaryIdentification *string `json:"secondary_identification,omitempty"`
	SortCode *string `json:"sort_code,omitempty"`
}

// PaymentRequest is the typed data model for the payment_request entity.
type PaymentRequest struct {
	Amount *float64 `json:"amount,omitempty"`
	DoSimulateSettlementFailure *bool `json:"do_simulate_settlement_failure,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
	Institution *string `json:"institution,omitempty"`
	PaymentInitiationId *string `json:"payment_initiation_id,omitempty"`
}

// PaymentRequestLoadMatch is the typed request payload for PaymentRequest.LoadTyped.
type PaymentRequestLoadMatch struct {
	PaymentrequestId *string `json:"paymentrequest_id,omitempty"`
}

// PaymentRequestCreateData is the typed request payload for PaymentRequest.CreateTyped.
type PaymentRequestCreateData struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// PaymentRequestUpdateData is the typed request payload for PaymentRequest.UpdateTyped.
type PaymentRequestUpdateData struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// PaymentRequestRemoveMatch is the typed request payload for PaymentRequest.RemoveTyped.
type PaymentRequestRemoveMatch struct {
	Id string `json:"id"`
}

// Payout is the typed data model for the payout entity.
type Payout struct {
	AccountId *string `json:"account_id,omitempty"`
	AllowIncomplete *bool `json:"allow_incomplete,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	ApprovePayoutUrl *string `json:"approve_payout_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	BeneficiaryId *string `json:"beneficiary_id,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanProcess *bool `json:"can_process,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	ChargeBearer *string `json:"charge_bearer,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserId *string `json:"current_user_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Document *[]any `json:"document,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formatted_fx_destination_amount,omitempty"`
	FormattedSchedule *string `json:"formatted_schedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formatted_schedule_day_only,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formatted_source_account_available_balance,omitempty"`
	FxDestinationAmount *float64 `json:"fx_destination_amount,omitempty"`
	FxDestinationAmountMinorUnit *int `json:"fx_destination_amount_minor_unit,omitempty"`
	FxDestinationCurrency *string `json:"fx_destination_currency,omitempty"`
	FxQuoteExpiresAt *string `json:"fx_quote_expires_at,omitempty"`
	FxQuoteId *string `json:"fx_quote_id,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	FxUseDestinationAmount *bool `json:"fx_use_destination_amount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsFailed *bool `json:"is_failed,omitempty"`
	IsSettled *bool `json:"is_settled,omitempty"`
	IsSubmitted *bool `json:"is_submitted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentRail *string `json:"payment_rail,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PayrunName *string `json:"payrun_name,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"source_account_available_balance,omitempty"`
	SourceAccountAvailableBalanceMinorUnit *int `json:"source_account_available_balance_minor_unit,omitempty"`
	SourceAccountBic *string `json:"source_account_bic,omitempty"`
	SourceAccountCurrency *string `json:"source_account_currency,omitempty"`
	SourceAccountIban *string `json:"source_account_iban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"source_account_identifier"`
	SourceAccountName *string `json:"source_account_name,omitempty"`
	SourceAccountNumber *string `json:"source_account_number,omitempty"`
	SourceAccountSortcode *string `json:"source_account_sortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TagId *[]any `json:"tag_id,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TopupPayrunId *string `json:"topup_payrun_id,omitempty"`
	TransactedAmount *float64 `json:"transacted_amount,omitempty"`
	TransactedFxAmount *float64 `json:"transacted_fx_amount,omitempty"`
	TransactedFxRate *float64 `json:"transacted_fx_rate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// PayoutLoadMatch is the typed request payload for Payout.LoadTyped.
type PayoutLoadMatch struct {
	Amount *float64 `json:"amount,omitempty"`
	Destination *string `json:"destination,omitempty"`
	Source *string `json:"source,omitempty"`
	Id *string `json:"id,omitempty"`
}

// PayoutCreateData is the typed request payload for Payout.CreateTyped.
type PayoutCreateData struct {
	Id *string `json:"id,omitempty"`
}

// PayoutUpdateData is the typed request payload for Payout.UpdateTyped.
type PayoutUpdateData struct {
	Id string `json:"id"`
}

// PayoutRemoveMatch is the typed request payload for Payout.RemoveTyped.
type PayoutRemoveMatch struct {
	Id string `json:"id"`
}

// Payrun is the typed data model for the payrun entity.
type Payrun struct {
	Id *string `json:"id,omitempty"`
	Note *string `json:"note,omitempty"`
	ScheduledDate *string `json:"scheduled_date,omitempty"`
}

// PayrunCreateData is the typed request payload for Payrun.CreateTyped.
type PayrunCreateData struct {
	Id string `json:"id"`
}

// PayrunUpdateData is the typed request payload for Payrun.UpdateTyped.
type PayrunUpdateData struct {
	Id string `json:"id"`
}

// PayrunRemoveMatch is the typed request payload for Payrun.RemoveTyped.
type PayrunRemoveMatch struct {
	Id string `json:"id"`
}

// Reject is the typed data model for the reject entity.
type Reject struct {
	AccountId *string `json:"account_id,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	ApprovePayoutUrl *string `json:"approve_payout_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanProcess *bool `json:"can_process,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	ChargeBearer *string `json:"charge_bearer,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserId *string `json:"current_user_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Document *[]any `json:"document,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formatted_fx_destination_amount,omitempty"`
	FormattedSchedule *string `json:"formatted_schedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formatted_schedule_day_only,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formatted_source_account_available_balance,omitempty"`
	FxDestinationAmount *float64 `json:"fx_destination_amount,omitempty"`
	FxDestinationAmountMinorUnit *int `json:"fx_destination_amount_minor_unit,omitempty"`
	FxDestinationCurrency *string `json:"fx_destination_currency,omitempty"`
	FxQuoteExpiresAt *string `json:"fx_quote_expires_at,omitempty"`
	FxQuoteId *string `json:"fx_quote_id,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	FxUseDestinationAmount *bool `json:"fx_use_destination_amount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsFailed *bool `json:"is_failed,omitempty"`
	IsSettled *bool `json:"is_settled,omitempty"`
	IsSubmitted *bool `json:"is_submitted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentRail *string `json:"payment_rail,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PayrunName *string `json:"payrun_name,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"source_account_available_balance,omitempty"`
	SourceAccountAvailableBalanceMinorUnit *int `json:"source_account_available_balance_minor_unit,omitempty"`
	SourceAccountBic *string `json:"source_account_bic,omitempty"`
	SourceAccountCurrency *string `json:"source_account_currency,omitempty"`
	SourceAccountIban *string `json:"source_account_iban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"source_account_identifier"`
	SourceAccountName *string `json:"source_account_name,omitempty"`
	SourceAccountNumber *string `json:"source_account_number,omitempty"`
	SourceAccountSortcode *string `json:"source_account_sortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TopupPayrunId *string `json:"topup_payrun_id,omitempty"`
	TransactedAmount *float64 `json:"transacted_amount,omitempty"`
	TransactedFxAmount *float64 `json:"transacted_fx_amount,omitempty"`
	TransactedFxRate *float64 `json:"transacted_fx_rate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// RejectUpdateData is the typed request payload for Reject.UpdateTyped.
type RejectUpdateData struct {
	Id string `json:"id"`
}

// Report is the typed data model for the report entity.
type Report struct {
}

// ReportUpdateData is the typed request payload for Report.UpdateTyped.
type ReportUpdateData struct {
	Id string `json:"id"`
}

// Rule is the typed data model for the rule entity.
type Rule struct {
}

// RuleUpdateData is the typed request payload for Rule.UpdateTyped.
type RuleUpdateData struct {
	Id string `json:"id"`
}

// RuleRemoveMatch is the typed request payload for Rule.RemoveTyped.
type RuleRemoveMatch struct {
	Id string `json:"id"`
}

// Send is the typed data model for the send entity.
type Send struct {
	AccountId *string `json:"account_id,omitempty"`
	AllowIncomplete *bool `json:"allow_incomplete,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	ApprovePayoutUrl *string `json:"approve_payout_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	BeneficiaryId *string `json:"beneficiary_id,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanProcess *bool `json:"can_process,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	ChargeBearer *string `json:"charge_bearer,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserId *string `json:"current_user_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Document *[]any `json:"document,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formatted_fx_destination_amount,omitempty"`
	FormattedSchedule *string `json:"formatted_schedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formatted_schedule_day_only,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formatted_source_account_available_balance,omitempty"`
	FxDestinationAmount *float64 `json:"fx_destination_amount,omitempty"`
	FxDestinationAmountMinorUnit *int `json:"fx_destination_amount_minor_unit,omitempty"`
	FxDestinationCurrency *string `json:"fx_destination_currency,omitempty"`
	FxQuoteExpiresAt *string `json:"fx_quote_expires_at,omitempty"`
	FxQuoteId *string `json:"fx_quote_id,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	FxUseDestinationAmount *bool `json:"fx_use_destination_amount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsFailed *bool `json:"is_failed,omitempty"`
	IsSettled *bool `json:"is_settled,omitempty"`
	IsSubmitted *bool `json:"is_submitted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentRail *string `json:"payment_rail,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PayrunName *string `json:"payrun_name,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"source_account_available_balance,omitempty"`
	SourceAccountAvailableBalanceMinorUnit *int `json:"source_account_available_balance_minor_unit,omitempty"`
	SourceAccountBic *string `json:"source_account_bic,omitempty"`
	SourceAccountCurrency *string `json:"source_account_currency,omitempty"`
	SourceAccountIban *string `json:"source_account_iban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"source_account_identifier"`
	SourceAccountName *string `json:"source_account_name,omitempty"`
	SourceAccountNumber *string `json:"source_account_number,omitempty"`
	SourceAccountSortcode *string `json:"source_account_sortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TagId *[]any `json:"tag_id,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TopupPayrunId *string `json:"topup_payrun_id,omitempty"`
	TransactedAmount *float64 `json:"transacted_amount,omitempty"`
	TransactedFxAmount *float64 `json:"transacted_fx_amount,omitempty"`
	TransactedFxRate *float64 `json:"transacted_fx_rate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// SendCreateData is the typed request payload for Send.CreateTyped.
type SendCreateData struct {
	AccountId *string `json:"account_id,omitempty"`
	AllowIncomplete *bool `json:"allow_incomplete,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	ApprovePayoutUrl *string `json:"approve_payout_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	BeneficiaryId *string `json:"beneficiary_id,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanProcess *bool `json:"can_process,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	ChargeBearer *string `json:"charge_bearer,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserId *string `json:"current_user_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Document *[]any `json:"document,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formatted_fx_destination_amount,omitempty"`
	FormattedSchedule *string `json:"formatted_schedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formatted_schedule_day_only,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formatted_source_account_available_balance,omitempty"`
	FxDestinationAmount *float64 `json:"fx_destination_amount,omitempty"`
	FxDestinationAmountMinorUnit *int `json:"fx_destination_amount_minor_unit,omitempty"`
	FxDestinationCurrency *string `json:"fx_destination_currency,omitempty"`
	FxQuoteExpiresAt *string `json:"fx_quote_expires_at,omitempty"`
	FxQuoteId *string `json:"fx_quote_id,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	FxUseDestinationAmount *bool `json:"fx_use_destination_amount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsFailed *bool `json:"is_failed,omitempty"`
	IsSettled *bool `json:"is_settled,omitempty"`
	IsSubmitted *bool `json:"is_submitted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentRail *string `json:"payment_rail,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PayrunName *string `json:"payrun_name,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"source_account_available_balance,omitempty"`
	SourceAccountAvailableBalanceMinorUnit *int `json:"source_account_available_balance_minor_unit,omitempty"`
	SourceAccountBic *string `json:"source_account_bic,omitempty"`
	SourceAccountCurrency *string `json:"source_account_currency,omitempty"`
	SourceAccountIban *string `json:"source_account_iban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"source_account_identifier"`
	SourceAccountName *string `json:"source_account_name,omitempty"`
	SourceAccountNumber *string `json:"source_account_number,omitempty"`
	SourceAccountSortcode *string `json:"source_account_sortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TagId *[]any `json:"tag_id,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TopupPayrunId *string `json:"topup_payrun_id,omitempty"`
	TransactedAmount *float64 `json:"transacted_amount,omitempty"`
	TransactedFxAmount *float64 `json:"transacted_fx_amount,omitempty"`
	TransactedFxRate *float64 `json:"transacted_fx_rate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// Sendbeneficiary is the typed data model for the sendbeneficiary entity.
type Sendbeneficiary struct {
	AccountId *string `json:"account_id,omitempty"`
	AllowIncomplete *bool `json:"allow_incomplete,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	ApprovePayoutUrl *string `json:"approve_payout_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	BeneficiaryId *string `json:"beneficiary_id,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanProcess *bool `json:"can_process,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	ChargeBearer *string `json:"charge_bearer,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserId *string `json:"current_user_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Document *[]any `json:"document,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formatted_fx_destination_amount,omitempty"`
	FormattedSchedule *string `json:"formatted_schedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formatted_schedule_day_only,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formatted_source_account_available_balance,omitempty"`
	FxDestinationAmount *float64 `json:"fx_destination_amount,omitempty"`
	FxDestinationAmountMinorUnit *int `json:"fx_destination_amount_minor_unit,omitempty"`
	FxDestinationCurrency *string `json:"fx_destination_currency,omitempty"`
	FxQuoteExpiresAt *string `json:"fx_quote_expires_at,omitempty"`
	FxQuoteId *string `json:"fx_quote_id,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	FxUseDestinationAmount *bool `json:"fx_use_destination_amount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsFailed *bool `json:"is_failed,omitempty"`
	IsSettled *bool `json:"is_settled,omitempty"`
	IsSubmitted *bool `json:"is_submitted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentRail *string `json:"payment_rail,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PayrunName *string `json:"payrun_name,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"source_account_available_balance,omitempty"`
	SourceAccountAvailableBalanceMinorUnit *int `json:"source_account_available_balance_minor_unit,omitempty"`
	SourceAccountBic *string `json:"source_account_bic,omitempty"`
	SourceAccountCurrency *string `json:"source_account_currency,omitempty"`
	SourceAccountIban *string `json:"source_account_iban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"source_account_identifier"`
	SourceAccountName *string `json:"source_account_name,omitempty"`
	SourceAccountNumber *string `json:"source_account_number,omitempty"`
	SourceAccountSortcode *string `json:"source_account_sortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TagId *[]any `json:"tag_id,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TopupPayrunId *string `json:"topup_payrun_id,omitempty"`
	TransactedAmount *float64 `json:"transacted_amount,omitempty"`
	TransactedFxAmount *float64 `json:"transacted_fx_amount,omitempty"`
	TransactedFxRate *float64 `json:"transacted_fx_rate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// SendbeneficiaryCreateData is the typed request payload for Sendbeneficiary.CreateTyped.
type SendbeneficiaryCreateData struct {
	AccountId *string `json:"account_id,omitempty"`
	AllowIncomplete *bool `json:"allow_incomplete,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnit *int `json:"amount_minor_unit,omitempty"`
	ApprovePayoutUrl *string `json:"approve_payout_url,omitempty"`
	ApproverId *string `json:"approver_id,omitempty"`
	AuthenticationMethod *[]any `json:"authentication_method,omitempty"`
	Authorisation *[]any `json:"authorisation,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisers_completed_count,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisers_required_count,omitempty"`
	BatchPayoutId *string `json:"batch_payout_id,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	BeneficiaryId *string `json:"beneficiary_id,omitempty"`
	CanAuthorise *bool `json:"can_authorise,omitempty"`
	CanProcess *bool `json:"can_process,omitempty"`
	CanUpdate *bool `json:"can_update,omitempty"`
	ChargeBearer *string `json:"charge_bearer,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	CreatedByEmailAddress *string `json:"created_by_email_address,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserId *string `json:"current_user_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Document *[]any `json:"document,omitempty"`
	Event *[]any `json:"event,omitempty"`
	FormattedAmount *string `json:"formatted_amount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formatted_fx_destination_amount,omitempty"`
	FormattedSchedule *string `json:"formatted_schedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formatted_schedule_day_only,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formatted_source_account_available_balance,omitempty"`
	FxDestinationAmount *float64 `json:"fx_destination_amount,omitempty"`
	FxDestinationAmountMinorUnit *int `json:"fx_destination_amount_minor_unit,omitempty"`
	FxDestinationCurrency *string `json:"fx_destination_currency,omitempty"`
	FxQuoteExpiresAt *string `json:"fx_quote_expires_at,omitempty"`
	FxQuoteId *string `json:"fx_quote_id,omitempty"`
	FxRate *float64 `json:"fx_rate,omitempty"`
	FxUseDestinationAmount *bool `json:"fx_use_destination_amount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"has_current_user_authorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceId *string `json:"invoice_id,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsFailed *bool `json:"is_failed,omitempty"`
	IsSettled *bool `json:"is_settled,omitempty"`
	IsSubmitted *bool `json:"is_submitted,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantTokenDescription *string `json:"merchant_token_description,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"payment_processor,omitempty"`
	PaymentRail *string `json:"payment_rail,omitempty"`
	PayrunId *string `json:"payrun_id,omitempty"`
	PayrunName *string `json:"payrun_name,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"schedule_date,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"source_account_available_balance,omitempty"`
	SourceAccountAvailableBalanceMinorUnit *int `json:"source_account_available_balance_minor_unit,omitempty"`
	SourceAccountBic *string `json:"source_account_bic,omitempty"`
	SourceAccountCurrency *string `json:"source_account_currency,omitempty"`
	SourceAccountIban *string `json:"source_account_iban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"source_account_identifier"`
	SourceAccountName *string `json:"source_account_name,omitempty"`
	SourceAccountNumber *string `json:"source_account_number,omitempty"`
	SourceAccountSortcode *string `json:"source_account_sortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tag *[]any `json:"tag,omitempty"`
	TagId *[]any `json:"tag_id,omitempty"`
	TheirReference *string `json:"their_reference,omitempty"`
	TopupPayrunId *string `json:"topup_payrun_id,omitempty"`
	TransactedAmount *float64 `json:"transacted_amount,omitempty"`
	TransactedFxAmount *float64 `json:"transacted_fx_amount,omitempty"`
	TransactedFxRate *float64 `json:"transacted_fx_rate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	YourReference *string `json:"your_reference,omitempty"`
}

// Tag is the typed data model for the tag entity.
type Tag struct {
	ColourHex *string `json:"colour_hex,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	MerchantId string `json:"merchant_id"`
	Name string `json:"name"`
}

// TagListMatch is the typed request payload for Tag.ListTyped.
type TagListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// TagCreateData is the typed request payload for Tag.CreateTyped.
type TagCreateData struct {
	MerchantId string `json:"merchant_id"`
}

// Token is the typed data model for the token entity.
type Token struct {
}

// TokenCreateData is the typed request payload for Token.CreateTyped.
type TokenCreateData struct {
	Id string `json:"id"`
}

// TokenRemoveMatch is the typed request payload for Token.RemoveTyped.
type TokenRemoveMatch struct {
	Id string `json:"id"`
}

// Transaction is the typed data model for the transaction entity.
type Transaction struct {
}

// TransactionLoadMatch is the typed request payload for Transaction.LoadTyped.
type TransactionLoadMatch struct {
	SequenceNumber int `json:"sequence_number"`
	TransactionId string `json:"transaction_id"`
}

// TransactionCreateData is the typed request payload for Transaction.CreateTyped.
type TransactionCreateData struct {
	Id string `json:"id"`
}

// TransactionRemoveMatch is the typed request payload for Transaction.RemoveTyped.
type TransactionRemoveMatch struct {
	Id string `json:"id"`
}

// User is the typed data model for the user entity.
type User struct {
	ClientSessionTimeout *[]any `json:"client_session_timeout,omitempty"`
	EmailAddress string `json:"email_address"`
	FirstName string `json:"first_name"`
	Id *string `json:"id,omitempty"`
	LastName string `json:"last_name"`
	PasskeyAdded *bool `json:"passkey_added,omitempty"`
	Permission *map[string]any `json:"permission,omitempty"`
	Profile *string `json:"profile,omitempty"`
	RolesWithScope *[]any `json:"roles_with_scope,omitempty"`
	TwoFactorEnabled *bool `json:"two_factor_enabled,omitempty"`
	UserInviteId *string `json:"user_invite_id,omitempty"`
}

// UserListMatch is the typed request payload for User.ListTyped.
type UserListMatch struct {
	MerchantId *string `json:"merchant_id,omitempty"`
}

// UserUpdateData is the typed request payload for User.UpdateTyped.
type UserUpdateData struct {
	Id string `json:"id"`
}

// UserInvite is the typed data model for the user_invite entity.
type UserInvite struct {
}

// UserInviteCreateData is the typed request payload for UserInvite.CreateTyped.
type UserInviteCreateData struct {
	Id string `json:"id"`
}

// UserInviteUpdateData is the typed request payload for UserInvite.UpdateTyped.
type UserInviteUpdateData struct {
	Id string `json:"id"`
}

// UserInviteRemoveMatch is the typed request payload for UserInvite.RemoveTyped.
type UserInviteRemoveMatch struct {
	Id string `json:"id"`
}

// Virtual is the typed data model for the virtual entity.
type Virtual struct {
	AccountName *string `json:"account_name,omitempty"`
	AccountSupplierName *string `json:"account_supplier_name,omitempty"`
	AvailableBalance *float64 `json:"available_balance,omitempty"`
	AvailableBalanceMinorUnit *int `json:"available_balance_minor_unit,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnit *int `json:"balance_minor_unit,omitempty"`
	BankName *string `json:"bank_name,omitempty"`
	ConsentId *string `json:"consent_id,omitempty"`
	CreatedBy map[string]any `json:"created_by"`
	CreatedByDisplayName *string `json:"created_by_display_name,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"default_payment_rail,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	ExpiryDate *string `json:"expiry_date,omitempty"`
	ExternalAccountIcon *string `json:"external_account_icon,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"is_archived,omitempty"`
	IsConnectedAccount *bool `json:"is_connected_account,omitempty"`
	IsDefault *bool `json:"is_default,omitempty"`
	IsTrustAccount *bool `json:"is_trust_account,omitempty"`
	IsVirtual *bool `json:"is_virtual,omitempty"`
	LastTransaction *map[string]any `json:"last_transaction,omitempty"`
	LastUpdated *string `json:"last_updated,omitempty"`
	MerchantId *string `json:"merchant_id,omitempty"`
	MerchantName *string `json:"merchant_name,omitempty"`
	Name string `json:"name"`
	PhysicalAccountId *string `json:"physical_account_id,omitempty"`
	Rule *[]any `json:"rule,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submitted_payouts_balance,omitempty"`
	SubmittedPayoutsBalanceMinorUnit *int `json:"submitted_payouts_balance_minor_unit,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplier_sepa_instant_status,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xero_bank_feed_connection_status,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xero_bank_feed_last_synced_at,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xero_bank_feed_sync_last_failed_at,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xero_bank_feed_sync_last_failure_reason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xero_bank_feed_sync_status,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xero_unsynchronised_transactions_count,omitempty"`
}

// VirtualCreateData is the typed request payload for Virtual.CreateTyped.
type VirtualCreateData struct {
	AccountId string `json:"account_id"`
}

// VirtualUpdateData is the typed request payload for Virtual.UpdateTyped.
type VirtualUpdateData struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
}

// Webhook is the typed data model for the webhook entity.
type Webhook struct {
}

// WebhookRemoveMatch is the typed request payload for Webhook.RemoveTyped.
type WebhookRemoveMatch struct {
	Id string `json:"id"`
}

// Whoami is the typed data model for the whoami entity.
type Whoami struct {
	ClientSessionTimeout *[]any `json:"client_session_timeout,omitempty"`
	EmailAddress string `json:"email_address"`
	FirstName string `json:"first_name"`
	Id *string `json:"id,omitempty"`
	LastName string `json:"last_name"`
	PasskeyAdded *bool `json:"passkey_added,omitempty"`
	Permission *map[string]any `json:"permission,omitempty"`
	RolesWithScope *[]any `json:"roles_with_scope,omitempty"`
	TwoFactorEnabled *bool `json:"two_factor_enabled,omitempty"`
}

// WhoamiListMatch is the typed request payload for Whoami.ListTyped.
type WhoamiListMatch struct {
	ClientSessionTimeout *[]any `json:"client_session_timeout,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	PasskeyAdded *bool `json:"passkey_added,omitempty"`
	Permission *map[string]any `json:"permission,omitempty"`
	RolesWithScope *[]any `json:"roles_with_scope,omitempty"`
	TwoFactorEnabled *bool `json:"two_factor_enabled,omitempty"`
}

// Whoamitrustedapp is the typed data model for the whoamitrustedapp entity.
type Whoamitrustedapp struct {
	ClientSessionTimeout *[]any `json:"client_session_timeout,omitempty"`
	EmailAddress string `json:"email_address"`
	FirstName string `json:"first_name"`
	Id *string `json:"id,omitempty"`
	LastName string `json:"last_name"`
	PasskeyAdded *bool `json:"passkey_added,omitempty"`
	Permission *map[string]any `json:"permission,omitempty"`
	RolesWithScope *[]any `json:"roles_with_scope,omitempty"`
	TwoFactorEnabled *bool `json:"two_factor_enabled,omitempty"`
}

// WhoamitrustedappListMatch is the typed request payload for Whoamitrustedapp.ListTyped.
type WhoamitrustedappListMatch struct {
	ClientSessionTimeout *[]any `json:"client_session_timeout,omitempty"`
	EmailAddress *string `json:"email_address,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	PasskeyAdded *bool `json:"passkey_added,omitempty"`
	Permission *map[string]any `json:"permission,omitempty"`
	RolesWithScope *[]any `json:"roles_with_scope,omitempty"`
	TwoFactorEnabled *bool `json:"two_factor_enabled,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
