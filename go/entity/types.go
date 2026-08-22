// Typed models for the Nofrixion SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/nofrixion-sdk/go/core"
)

// Account is the typed data model for the account entity.
type Account struct {
	AccountBalances *[]any `json:"accountBalances,omitempty"`
	AccountID *string `json:"accountID,omitempty"`
	AccountIdentifications *[]any `json:"accountIdentifications,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	AccountNames *[]any `json:"accountNames,omitempty"`
	AccountSupplierName *string `json:"accountSupplierName,omitempty"`
	AccountType *string `json:"accountType,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	AvailableBalanceMinorUnits *int `json:"availableBalanceMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	ConsolidatedAccountInformation *map[string]any `json:"consolidatedAccountInformation,omitempty"`
	CreatedBy map[string]any `json:"createdBy"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"defaultPaymentRail,omitempty"`
	Description *string `json:"description,omitempty"`
	Details *string `json:"details,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	ExternalAccountIcon *string `json:"externalAccountIcon,omitempty"`
	Format *string `json:"format,omitempty"`
	FromDate *string `json:"fromDate,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsTrustAccount *bool `json:"isTrustAccount,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	LastTransaction *map[string]any `json:"lastTransaction,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	PhysicalAccountID *string `json:"physicalAccountID,omitempty"`
	RoleIDs *[]any `json:"roleIDs,omitempty"`
	Rules *[]any `json:"rules,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
	SubmittedPayoutsBalanceMinorUnits *int `json:"submittedPayoutsBalanceMinorUnits,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierPhysicalAccountID *string `json:"supplierPhysicalAccountID,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplierSepaInstantStatus,omitempty"`
	ToDate *string `json:"toDate,omitempty"`
	Type *string `json:"type,omitempty"`
	UsageType *string `json:"usageType,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xeroBankFeedConnectionStatus,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xeroBankFeedLastSyncedAt,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xeroBankFeedSyncLastFailedAt,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xeroBankFeedSyncLastFailureReason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xeroBankFeedSyncStatus,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xeroUnsynchronisedTransactionsCount,omitempty"`
}

// AccountLoadMatch is the typed request payload for Account.LoadTyped.
type AccountLoadMatch struct {
	AccountId *string `json:"account_id,omitempty"`
	Id string `json:"id"`
	MerchantId *string `json:"merchant_id,omitempty"`
}

// AccountListMatch is the typed request payload for Account.ListTyped.
type AccountListMatch struct {
	AccountBalances *[]any `json:"accountBalances,omitempty"`
	AccountID *string `json:"accountID,omitempty"`
	AccountIdentifications *[]any `json:"accountIdentifications,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	AccountNames *[]any `json:"accountNames,omitempty"`
	AccountSupplierName *string `json:"accountSupplierName,omitempty"`
	AccountType *string `json:"accountType,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	AvailableBalanceMinorUnits *int `json:"availableBalanceMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	ConsolidatedAccountInformation *map[string]any `json:"consolidatedAccountInformation,omitempty"`
	CreatedBy *map[string]any `json:"createdBy,omitempty"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"defaultPaymentRail,omitempty"`
	Description *string `json:"description,omitempty"`
	Details *string `json:"details,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	ExternalAccountIcon *string `json:"externalAccountIcon,omitempty"`
	Format *string `json:"format,omitempty"`
	FromDate *string `json:"fromDate,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier *map[string]any `json:"identifier,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsTrustAccount *bool `json:"isTrustAccount,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	LastTransaction *map[string]any `json:"lastTransaction,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	PhysicalAccountID *string `json:"physicalAccountID,omitempty"`
	RoleIDs *[]any `json:"roleIDs,omitempty"`
	Rules *[]any `json:"rules,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
	SubmittedPayoutsBalanceMinorUnits *int `json:"submittedPayoutsBalanceMinorUnits,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierPhysicalAccountID *string `json:"supplierPhysicalAccountID,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplierSepaInstantStatus,omitempty"`
	ToDate *string `json:"toDate,omitempty"`
	Type *string `json:"type,omitempty"`
	UsageType *string `json:"usageType,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xeroBankFeedConnectionStatus,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xeroBankFeedLastSyncedAt,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xeroBankFeedSyncLastFailedAt,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xeroBankFeedSyncLastFailureReason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xeroBankFeedSyncStatus,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xeroUnsynchronisedTransactionsCount,omitempty"`
}

// AccountCreateData is the typed request payload for Account.CreateTyped.
type AccountCreateData struct {
	AccountId string `json:"account_id"`
	Currency string `json:"currency"`
	AccountBalances *[]any `json:"accountBalances,omitempty"`
	AccountID *string `json:"accountID,omitempty"`
	AccountIdentifications *[]any `json:"accountIdentifications,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	AccountNames *[]any `json:"accountNames,omitempty"`
	AccountSupplierName *string `json:"accountSupplierName,omitempty"`
	AccountType *string `json:"accountType,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	AvailableBalanceMinorUnits *int `json:"availableBalanceMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	ConsolidatedAccountInformation *map[string]any `json:"consolidatedAccountInformation,omitempty"`
	CreatedBy map[string]any `json:"createdBy"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	DefaultPaymentRail *string `json:"defaultPaymentRail,omitempty"`
	Description *string `json:"description,omitempty"`
	Details *string `json:"details,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	ExternalAccountIcon *string `json:"externalAccountIcon,omitempty"`
	Format *string `json:"format,omitempty"`
	FromDate *string `json:"fromDate,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsTrustAccount *bool `json:"isTrustAccount,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	LastTransaction *map[string]any `json:"lastTransaction,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	PhysicalAccountID *string `json:"physicalAccountID,omitempty"`
	RoleIDs *[]any `json:"roleIDs,omitempty"`
	Rules *[]any `json:"rules,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
	SubmittedPayoutsBalanceMinorUnits *int `json:"submittedPayoutsBalanceMinorUnits,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierPhysicalAccountID *string `json:"supplierPhysicalAccountID,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplierSepaInstantStatus,omitempty"`
	ToDate *string `json:"toDate,omitempty"`
	Type *string `json:"type,omitempty"`
	UsageType *string `json:"usageType,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xeroBankFeedConnectionStatus,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xeroBankFeedLastSyncedAt,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xeroBankFeedSyncLastFailedAt,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xeroBankFeedSyncLastFailureReason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xeroBankFeedSyncStatus,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xeroUnsynchronisedTransactionsCount,omitempty"`
}

// AccountUpdateData is the typed request payload for Account.UpdateTyped.
type AccountUpdateData struct {
	Id string `json:"id"`
	AccountBalances *[]any `json:"accountBalances,omitempty"`
	AccountID *string `json:"accountID,omitempty"`
	AccountIdentifications *[]any `json:"accountIdentifications,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	AccountNames *[]any `json:"accountNames,omitempty"`
	AccountSupplierName *string `json:"accountSupplierName,omitempty"`
	AccountType *string `json:"accountType,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	AvailableBalanceMinorUnits *int `json:"availableBalanceMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	ConsolidatedAccountInformation *map[string]any `json:"consolidatedAccountInformation,omitempty"`
	CreatedBy *map[string]any `json:"createdBy,omitempty"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"defaultPaymentRail,omitempty"`
	Description *string `json:"description,omitempty"`
	Details *string `json:"details,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	ExternalAccountIcon *string `json:"externalAccountIcon,omitempty"`
	Format *string `json:"format,omitempty"`
	FromDate *string `json:"fromDate,omitempty"`
	Identifier *map[string]any `json:"identifier,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsTrustAccount *bool `json:"isTrustAccount,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	LastTransaction *map[string]any `json:"lastTransaction,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	PhysicalAccountID *string `json:"physicalAccountID,omitempty"`
	RoleIDs *[]any `json:"roleIDs,omitempty"`
	Rules *[]any `json:"rules,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
	SubmittedPayoutsBalanceMinorUnits *int `json:"submittedPayoutsBalanceMinorUnits,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierPhysicalAccountID *string `json:"supplierPhysicalAccountID,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplierSepaInstantStatus,omitempty"`
	ToDate *string `json:"toDate,omitempty"`
	Type *string `json:"type,omitempty"`
	UsageType *string `json:"usageType,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xeroBankFeedConnectionStatus,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xeroBankFeedLastSyncedAt,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xeroBankFeedSyncLastFailedAt,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xeroBankFeedSyncLastFailureReason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xeroBankFeedSyncStatus,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xeroUnsynchronisedTransactionsCount,omitempty"`
}

// AccountRemoveMatch is the typed request payload for Account.RemoveTyped.
type AccountRemoveMatch struct {
	Id string `json:"id"`
}

// Batch is the typed data model for the batch entity.
type Batch struct {
	ApproveUrl *string `json:"approveUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
}

// BatchLoadMatch is the typed request payload for Batch.LoadTyped.
type BatchLoadMatch struct {
	Id string `json:"id"`
}

// BatchCreateData is the typed request payload for Batch.CreateTyped.
type BatchCreateData struct {
	ApproveUrl *string `json:"approveUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
}

// Beneficiary is the typed data model for the beneficiary entity.
type Beneficiary struct {
	ApprovalCallbackUrl *string `json:"approvalCallbackUrl,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	Beneficiaries *[]any `json:"beneficiaries,omitempty"`
	BeneficiaryEvents *[]any `json:"beneficiaryEvents,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	CreatedBy map[string]any `json:"createdBy"`
	CreatedByEmailAddress *string `json:"createdByEmailAddress,omitempty"`
	Currency string `json:"currency"`
	Destination *map[string]any `json:"destination,omitempty"`
	FailedBeneficiaries *map[string]any `json:"failedBeneficiaries,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	LastAuthorised *string `json:"lastAuthorised,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name string `json:"name"`
	Nonce *string `json:"nonce,omitempty"`
	SourceAccountIDs *[]any `json:"sourceAccountIDs,omitempty"`
	SourceAccounts *[]any `json:"sourceAccounts,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
}

// BeneficiaryLoadMatch is the typed request payload for Beneficiary.LoadTyped.
type BeneficiaryLoadMatch struct {
	Id string `json:"id"`
	MerchantId *string `json:"merchant_id,omitempty"`
}

// BeneficiaryListMatch is the typed request payload for Beneficiary.ListTyped.
type BeneficiaryListMatch struct {
	ApprovalCallbackUrl *string `json:"approvalCallbackUrl,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	Beneficiaries *[]any `json:"beneficiaries,omitempty"`
	BeneficiaryEvents *[]any `json:"beneficiaryEvents,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	CreatedBy *map[string]any `json:"createdBy,omitempty"`
	CreatedByEmailAddress *string `json:"createdByEmailAddress,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	FailedBeneficiaries *map[string]any `json:"failedBeneficiaries,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	LastAuthorised *string `json:"lastAuthorised,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	SourceAccountIDs *[]any `json:"sourceAccountIDs,omitempty"`
	SourceAccounts *[]any `json:"sourceAccounts,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
}

// BeneficiaryCreateData is the typed request payload for Beneficiary.CreateTyped.
type BeneficiaryCreateData struct {
	Id string `json:"id"`
	ApprovalCallbackUrl *string `json:"approvalCallbackUrl,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	Beneficiaries *[]any `json:"beneficiaries,omitempty"`
	BeneficiaryEvents *[]any `json:"beneficiaryEvents,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	CreatedBy map[string]any `json:"createdBy"`
	CreatedByEmailAddress *string `json:"createdByEmailAddress,omitempty"`
	Currency string `json:"currency"`
	Destination *map[string]any `json:"destination,omitempty"`
	FailedBeneficiaries *map[string]any `json:"failedBeneficiaries,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	LastAuthorised *string `json:"lastAuthorised,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name string `json:"name"`
	Nonce *string `json:"nonce,omitempty"`
	SourceAccountIDs *[]any `json:"sourceAccountIDs,omitempty"`
	SourceAccounts *[]any `json:"sourceAccounts,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
}

// BeneficiaryUpdateData is the typed request payload for Beneficiary.UpdateTyped.
type BeneficiaryUpdateData struct {
	Id string `json:"id"`
	ApprovalCallbackUrl *string `json:"approvalCallbackUrl,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	Beneficiaries *[]any `json:"beneficiaries,omitempty"`
	BeneficiaryEvents *[]any `json:"beneficiaryEvents,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	CreatedBy *map[string]any `json:"createdBy,omitempty"`
	CreatedByEmailAddress *string `json:"createdByEmailAddress,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	FailedBeneficiaries *map[string]any `json:"failedBeneficiaries,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	LastAuthorised *string `json:"lastAuthorised,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	SourceAccountIDs *[]any `json:"sourceAccountIDs,omitempty"`
	SourceAccounts *[]any `json:"sourceAccounts,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
}

// BeneficiaryRemoveMatch is the typed request payload for Beneficiary.RemoveTyped.
type BeneficiaryRemoveMatch struct {
	Id string `json:"id"`
}

// BeneficiaryGroup is the typed data model for the beneficiary_group entity.
type BeneficiaryGroup struct {
	GroupMembers *[]any `json:"groupMembers,omitempty"`
	GroupName string `json:"groupName"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID string `json:"merchantID"`
}

// BeneficiaryGroupListMatch is the typed request payload for BeneficiaryGroup.ListTyped.
type BeneficiaryGroupListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// Card is the typed data model for the card entity.
type Card struct {
	AuthorizedAmount *string `json:"authorizedAmount,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
	IsPayerAuthenticationRequired *bool `json:"isPayerAuthenticationRequired,omitempty"`
	IsSoftDecline *bool `json:"isSoftDecline,omitempty"`
	PayerAuthenticationAccessToken *string `json:"payerAuthenticationAccessToken,omitempty"`
	PayerAuthenticationMerchantData *string `json:"payerAuthenticationMerchantData,omitempty"`
	PayerAuthenticationUrl *string `json:"payerAuthenticationUrl,omitempty"`
	PayerAuthenticationWindowHeight *int `json:"payerAuthenticationWindowHeight,omitempty"`
	PayerAuthenticationWindowWidth *int `json:"payerAuthenticationWindowWidth,omitempty"`
	PaymentRequestCallbackUrl *string `json:"paymentRequestCallbackUrl,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	Status *string `json:"status,omitempty"`
	ThreeDSRedirectUrl *string `json:"threeDSRedirectUrl,omitempty"`
	TransactionID *string `json:"transactionID,omitempty"`
}

// CardCreateData is the typed request payload for Card.CreateTyped.
type CardCreateData struct {
	PaymentrequestId string `json:"paymentrequest_id"`
	AuthorizedAmount *string `json:"authorizedAmount,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
	IsPayerAuthenticationRequired *bool `json:"isPayerAuthenticationRequired,omitempty"`
	IsSoftDecline *bool `json:"isSoftDecline,omitempty"`
	PayerAuthenticationAccessToken *string `json:"payerAuthenticationAccessToken,omitempty"`
	PayerAuthenticationMerchantData *string `json:"payerAuthenticationMerchantData,omitempty"`
	PayerAuthenticationUrl *string `json:"payerAuthenticationUrl,omitempty"`
	PayerAuthenticationWindowHeight *int `json:"payerAuthenticationWindowHeight,omitempty"`
	PayerAuthenticationWindowWidth *int `json:"payerAuthenticationWindowWidth,omitempty"`
	PaymentRequestCallbackUrl *string `json:"paymentRequestCallbackUrl,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	Status *string `json:"status,omitempty"`
	ThreeDSRedirectUrl *string `json:"threeDSRedirectUrl,omitempty"`
	TransactionID *string `json:"transactionID,omitempty"`
}

// CardCustomerToken is the typed data model for the card_customer_token entity.
type CardCustomerToken struct {
	CardType *string `json:"cardType,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	ExpiryMonth *string `json:"expiryMonth,omitempty"`
	ExpiryYear *string `json:"expiryYear,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastFourDigits *string `json:"lastFourDigits,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MaskedCardNumber *string `json:"maskedCardNumber,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
}

// CardCustomerTokenLoadMatch is the typed request payload for CardCustomerToken.LoadTyped.
type CardCustomerTokenLoadMatch struct {
	CustomerEmailAddress string `json:"customer_email_address"`
}

// CardCustomerTokenListMatch is the typed request payload for CardCustomerToken.ListTyped.
type CardCustomerTokenListMatch struct {
	CustomerEmailAddress string `json:"customer_email_address"`
	MerchantId string `json:"merchant_id"`
}

// CardCustomerTokenRemoveMatch is the typed request payload for CardCustomerToken.RemoveTyped.
type CardCustomerTokenRemoveMatch struct {
	Id string `json:"id"`
}

// CardPayment is the typed data model for the card_payment entity.
type CardPayment struct {
	AuthorizedAmount *string `json:"authorizedAmount,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
	IsPayerAuthenticationRequired *bool `json:"isPayerAuthenticationRequired,omitempty"`
	IsSoftDecline *bool `json:"isSoftDecline,omitempty"`
	PayerAuthenticationAccessToken *string `json:"payerAuthenticationAccessToken,omitempty"`
	PayerAuthenticationMerchantData *string `json:"payerAuthenticationMerchantData,omitempty"`
	PayerAuthenticationUrl *string `json:"payerAuthenticationUrl,omitempty"`
	PayerAuthenticationWindowHeight *int `json:"payerAuthenticationWindowHeight,omitempty"`
	PayerAuthenticationWindowWidth *int `json:"payerAuthenticationWindowWidth,omitempty"`
	PaymentRequestCallbackUrl *string `json:"paymentRequestCallbackUrl,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	Status *string `json:"status,omitempty"`
	ThreeDSRedirectUrl *string `json:"threeDSRedirectUrl,omitempty"`
	TransactionID *string `json:"transactionID,omitempty"`
}

// CardPaymentCreateData is the typed request payload for CardPayment.CreateTyped.
type CardPaymentCreateData struct {
	PartialRefundAmount *float64 `json:"partial_refund_amount,omitempty"`
	PaymentrequestId string `json:"paymentrequest_id"`
	AuthorizedAmount *string `json:"authorizedAmount,omitempty"`
	CurrencyCode *string `json:"currencyCode,omitempty"`
	IsPayerAuthenticationRequired *bool `json:"isPayerAuthenticationRequired,omitempty"`
	IsSoftDecline *bool `json:"isSoftDecline,omitempty"`
	PayerAuthenticationAccessToken *string `json:"payerAuthenticationAccessToken,omitempty"`
	PayerAuthenticationMerchantData *string `json:"payerAuthenticationMerchantData,omitempty"`
	PayerAuthenticationUrl *string `json:"payerAuthenticationUrl,omitempty"`
	PayerAuthenticationWindowHeight *int `json:"payerAuthenticationWindowHeight,omitempty"`
	PayerAuthenticationWindowWidth *int `json:"payerAuthenticationWindowWidth,omitempty"`
	PaymentRequestCallbackUrl *string `json:"paymentRequestCallbackUrl,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	RequestID *string `json:"requestID,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	Status *string `json:"status,omitempty"`
	ThreeDSRedirectUrl *string `json:"threeDSRedirectUrl,omitempty"`
	TransactionID *string `json:"transactionID,omitempty"`
}

// CardPublicKey is the typed data model for the card_public_key entity.
type CardPublicKey struct {
	Jwt *string `json:"jwt,omitempty"`
}

// CardPublicKeyLoadMatch is the typed request payload for CardPublicKey.LoadTyped.
type CardPublicKeyLoadMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// Consent is the typed data model for the consent entity.
type Consent struct {
	AuthorisationUrl *string `json:"authorisationUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InstitutionID *string `json:"institutionID,omitempty"`
	IsConnectedAccounts *bool `json:"isConnectedAccounts,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Provider *string `json:"provider,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
}

// ConsentLoadMatch is the typed request payload for Consent.LoadTyped.
type ConsentLoadMatch struct {
	Id string `json:"id"`
}

// ConsentListMatch is the typed request payload for Consent.ListTyped.
type ConsentListMatch struct {
	Email string `json:"email"`
	MerchantId string `json:"merchant_id"`
}

// ConsentCreateData is the typed request payload for Consent.CreateTyped.
type ConsentCreateData struct {
	AuthorisationUrl *string `json:"authorisationUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InstitutionID *string `json:"institutionID,omitempty"`
	IsConnectedAccounts *bool `json:"isConnectedAccounts,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Provider *string `json:"provider,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
}

// ConsentUpdateData is the typed request payload for Consent.UpdateTyped.
type ConsentUpdateData struct {
	Id string `json:"id"`
	AuthorisationUrl *string `json:"authorisationUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InstitutionID *string `json:"institutionID,omitempty"`
	IsConnectedAccounts *bool `json:"isConnectedAccounts,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Provider *string `json:"provider,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
}

// ConsentRemoveMatch is the typed request payload for Consent.RemoveTyped.
type ConsentRemoveMatch struct {
	Id string `json:"id"`
}

// Currency is the typed data model for the currency entity.
type Currency struct {
	Code *string `json:"code,omitempty"`
	Decimals *int `json:"decimals,omitempty"`
	IsFiat *bool `json:"isFiat,omitempty"`
	Iso4217AlphaCode *string `json:"iso4217AlphaCode,omitempty"`
	Iso4217NumericCode *string `json:"iso4217NumericCode,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// CurrencyListMatch is the typed request payload for Currency.ListTyped.
type CurrencyListMatch struct {
	Code *string `json:"code,omitempty"`
	Decimals *int `json:"decimals,omitempty"`
	IsFiat *bool `json:"isFiat,omitempty"`
	Iso4217AlphaCode *string `json:"iso4217AlphaCode,omitempty"`
	Iso4217NumericCode *string `json:"iso4217NumericCode,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
}

// DirectDebitBatchSubmit is the typed data model for the direct_debit_batch_submit entity.
type DirectDebitBatchSubmit struct {
	FailedSubmissions *map[string]any `json:"failedSubmissions,omitempty"`
	SuccessfulSubmissions *[]any `json:"successfulSubmissions,omitempty"`
}

// DirectDebitBatchSubmitCreateData is the typed request payload for DirectDebitBatchSubmit.CreateTyped.
type DirectDebitBatchSubmitCreateData struct {
	FailedSubmissions *map[string]any `json:"failedSubmissions,omitempty"`
	SuccessfulSubmissions *[]any `json:"successfulSubmissions,omitempty"`
}

// FxRate is the typed data model for the fx_rate entity.
type FxRate struct {
	DestinationCurrency *string `json:"destinationCurrency,omitempty"`
	ExchangeRate *float64 `json:"exchangeRate,omitempty"`
	ExpiryTime *string `json:"expiryTime,omitempty"`
	QuoteID *string `json:"quoteID,omitempty"`
	SourceCurrency *string `json:"sourceCurrency,omitempty"`
}

// FxRateLoadMatch is the typed request payload for FxRate.LoadTyped.
type FxRateLoadMatch struct {
	Destination string `json:"destination"`
	Source string `json:"source"`
	ValidForMinute int `json:"valid_for_minute"`
}

// FxRateListMatch is the typed request payload for FxRate.ListTyped.
type FxRateListMatch struct {
	Destination string `json:"destination"`
	Source string `json:"source"`
}

// IPayment is the typed data model for the i_payment entity.
type IPayment struct {
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
}

// IPaymentCreateData is the typed request payload for IPayment.CreateTyped.
type IPaymentCreateData struct {
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
}

// Mandate is the typed data model for the mandate entity.
type Mandate struct {
	AccountNumber *string `json:"accountNumber,omitempty"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	ApprovedAt *string `json:"approvedAt,omitempty"`
	City string `json:"city"`
	CountryCode string `json:"countryCode"`
	Currency *string `json:"currency,omitempty"`
	CustomerAccountNumber *string `json:"customerAccountNumber,omitempty"`
	CustomerCity *string `json:"customerCity,omitempty"`
	CustomerCountryCode *string `json:"customerCountryCode,omitempty"`
	CustomerCountryName *string `json:"customerCountryName,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerFirstName *string `json:"customerFirstName,omitempty"`
	CustomerIban *string `json:"customerIban,omitempty"`
	CustomerLastName *string `json:"customerLastName,omitempty"`
	CustomerSortCode *string `json:"customerSortCode,omitempty"`
	EmailAddress string `json:"emailAddress"`
	FirstName string `json:"firstName"`
	Iban *string `json:"iban,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsRecurring *bool `json:"isRecurring,omitempty"`
	LastName string `json:"lastName"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	PostalCode string `json:"postalCode"`
	Reference *string `json:"reference,omitempty"`
	SortCode *string `json:"sortCode,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplierBankAccountID *string `json:"supplierBankAccountID,omitempty"`
	SupplierCustomerID *string `json:"supplierCustomerID,omitempty"`
	SupplierMandateID *string `json:"supplierMandateID,omitempty"`
	SupplierName *string `json:"supplierName,omitempty"`
	SupplierStatus *string `json:"supplierStatus,omitempty"`
}

// MandateLoadMatch is the typed request payload for Mandate.LoadTyped.
type MandateLoadMatch struct {
	Id string `json:"id"`
}

// MandateCreateData is the typed request payload for Mandate.CreateTyped.
type MandateCreateData struct {
	AccountNumber *string `json:"accountNumber,omitempty"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2,omitempty"`
	ApprovedAt *string `json:"approvedAt,omitempty"`
	City string `json:"city"`
	CountryCode string `json:"countryCode"`
	Currency *string `json:"currency,omitempty"`
	CustomerAccountNumber *string `json:"customerAccountNumber,omitempty"`
	CustomerCity *string `json:"customerCity,omitempty"`
	CustomerCountryCode *string `json:"customerCountryCode,omitempty"`
	CustomerCountryName *string `json:"customerCountryName,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerFirstName *string `json:"customerFirstName,omitempty"`
	CustomerIban *string `json:"customerIban,omitempty"`
	CustomerLastName *string `json:"customerLastName,omitempty"`
	CustomerSortCode *string `json:"customerSortCode,omitempty"`
	EmailAddress string `json:"emailAddress"`
	FirstName string `json:"firstName"`
	Iban *string `json:"iban,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsRecurring *bool `json:"isRecurring,omitempty"`
	LastName string `json:"lastName"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	PostalCode string `json:"postalCode"`
	Reference *string `json:"reference,omitempty"`
	SortCode *string `json:"sortCode,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplierBankAccountID *string `json:"supplierBankAccountID,omitempty"`
	SupplierCustomerID *string `json:"supplierCustomerID,omitempty"`
	SupplierMandateID *string `json:"supplierMandateID,omitempty"`
	SupplierName *string `json:"supplierName,omitempty"`
	SupplierStatus *string `json:"supplierStatus,omitempty"`
}

// Merchant is the typed data model for the merchant entity.
type Merchant struct {
	AccountCurrencies *[]any `json:"accountCurrencies,omitempty"`
	CanHaveTrustAccounts *bool `json:"canHaveTrustAccounts,omitempty"`
	CardPaymentProcessor *string `json:"cardPaymentProcessor,omitempty"`
	CompanyID *string `json:"companyID,omitempty"`
	DisplayQrOnHostedPay *bool `json:"displayQrOnHostedPay,omitempty"`
	HostedPayVersion *int `json:"hostedPayVersion,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	IsExited *bool `json:"isExited,omitempty"`
	IsSuspended *bool `json:"isSuspended,omitempty"`
	Jurisdiction *string `json:"jurisdiction,omitempty"`
	LogoUrlPng *string `json:"logoUrlPng,omitempty"`
	LogoUrlSvg *string `json:"logoUrlSvg,omitempty"`
	MerchantCategoryCode *string `json:"merchantCategoryCode,omitempty"`
	Name *string `json:"name,omitempty"`
	Notes *string `json:"notes,omitempty"`
	ParentMerchant *map[string]any `json:"parentMerchant,omitempty"`
	PaymentAccountLimit *int `json:"paymentAccountLimit,omitempty"`
	PaymentAccounts *[]any `json:"paymentAccounts,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	SupportedPaymentMethodsList *[]any `json:"supportedPaymentMethodsList,omitempty"`
	SuspensionReason *string `json:"suspensionReason,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	TradingName *string `json:"tradingName,omitempty"`
	WebHookLimit *int `json:"webHookLimit,omitempty"`
	YourRoleName *string `json:"yourRoleName,omitempty"`
}

// MerchantLoadMatch is the typed request payload for Merchant.LoadTyped.
type MerchantLoadMatch struct {
	Id string `json:"id"`
}

// MerchantListMatch is the typed request payload for Merchant.ListTyped.
type MerchantListMatch struct {
	AccountCurrencies *[]any `json:"accountCurrencies,omitempty"`
	CanHaveTrustAccounts *bool `json:"canHaveTrustAccounts,omitempty"`
	CardPaymentProcessor *string `json:"cardPaymentProcessor,omitempty"`
	CompanyID *string `json:"companyID,omitempty"`
	DisplayQrOnHostedPay *bool `json:"displayQrOnHostedPay,omitempty"`
	HostedPayVersion *int `json:"hostedPayVersion,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	IsExited *bool `json:"isExited,omitempty"`
	IsSuspended *bool `json:"isSuspended,omitempty"`
	Jurisdiction *string `json:"jurisdiction,omitempty"`
	LogoUrlPng *string `json:"logoUrlPng,omitempty"`
	LogoUrlSvg *string `json:"logoUrlSvg,omitempty"`
	MerchantCategoryCode *string `json:"merchantCategoryCode,omitempty"`
	Name *string `json:"name,omitempty"`
	Notes *string `json:"notes,omitempty"`
	ParentMerchant *map[string]any `json:"parentMerchant,omitempty"`
	PaymentAccountLimit *int `json:"paymentAccountLimit,omitempty"`
	PaymentAccounts *[]any `json:"paymentAccounts,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	SupportedPaymentMethodsList *[]any `json:"supportedPaymentMethodsList,omitempty"`
	SuspensionReason *string `json:"suspensionReason,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	TradingName *string `json:"tradingName,omitempty"`
	WebHookLimit *int `json:"webHookLimit,omitempty"`
	YourRoleName *string `json:"yourRoleName,omitempty"`
}

// MerchantUpdateData is the typed request payload for Merchant.UpdateTyped.
type MerchantUpdateData struct {
	Id string `json:"id"`
	AccountCurrencies *[]any `json:"accountCurrencies,omitempty"`
	CanHaveTrustAccounts *bool `json:"canHaveTrustAccounts,omitempty"`
	CardPaymentProcessor *string `json:"cardPaymentProcessor,omitempty"`
	CompanyID *string `json:"companyID,omitempty"`
	DisplayQrOnHostedPay *bool `json:"displayQrOnHostedPay,omitempty"`
	HostedPayVersion *int `json:"hostedPayVersion,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsBlocked *bool `json:"isBlocked,omitempty"`
	IsExited *bool `json:"isExited,omitempty"`
	IsSuspended *bool `json:"isSuspended,omitempty"`
	Jurisdiction *string `json:"jurisdiction,omitempty"`
	LogoUrlPng *string `json:"logoUrlPng,omitempty"`
	LogoUrlSvg *string `json:"logoUrlSvg,omitempty"`
	MerchantCategoryCode *string `json:"merchantCategoryCode,omitempty"`
	Name *string `json:"name,omitempty"`
	Notes *string `json:"notes,omitempty"`
	ParentMerchant *map[string]any `json:"parentMerchant,omitempty"`
	PaymentAccountLimit *int `json:"paymentAccountLimit,omitempty"`
	PaymentAccounts *[]any `json:"paymentAccounts,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	SupportedPaymentMethodsList *[]any `json:"supportedPaymentMethodsList,omitempty"`
	SuspensionReason *string `json:"suspensionReason,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	TradingName *string `json:"tradingName,omitempty"`
	WebHookLimit *int `json:"webHookLimit,omitempty"`
	YourRoleName *string `json:"yourRoleName,omitempty"`
}

// MerchantRemoveMatch is the typed request payload for Merchant.RemoveTyped.
type MerchantRemoveMatch struct {
	Id string `json:"id"`
	UserId string `json:"user_id"`
}

// MerchantAuthorisationSetting is the typed data model for the merchant_authorisation_setting entity.
type MerchantAuthorisationSetting struct {
	AmountLower *float64 `json:"amountLower,omitempty"`
	AmountUpper *float64 `json:"amountUpper,omitempty"`
	AuthorisationType *string `json:"authorisationType,omitempty"`
	BeneficiariesOnly *bool `json:"beneficiariesOnly,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastEditorCantAuthorise *bool `json:"lastEditorCantAuthorise,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	NumberOfAuthorisers *int `json:"numberOfAuthorisers,omitempty"`
	RoleSettings *[]any `json:"roleSettings,omitempty"`
}

// MerchantAuthorisationSettingListMatch is the typed request payload for MerchantAuthorisationSetting.ListTyped.
type MerchantAuthorisationSettingListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// MerchantDirectDebitMandatePage is the typed data model for the merchant_direct_debit_mandate_page entity.
type MerchantDirectDebitMandatePage struct {
	ApprovedAt *string `json:"approvedAt,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomerAccountNumber *string `json:"customerAccountNumber,omitempty"`
	CustomerCity *string `json:"customerCity,omitempty"`
	CustomerCountryCode *string `json:"customerCountryCode,omitempty"`
	CustomerCountryName *string `json:"customerCountryName,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerFirstName *string `json:"customerFirstName,omitempty"`
	CustomerIban *string `json:"customerIban,omitempty"`
	CustomerLastName *string `json:"customerLastName,omitempty"`
	CustomerSortCode *string `json:"customerSortCode,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsRecurring *bool `json:"isRecurring,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Reference *string `json:"reference,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplierBankAccountID *string `json:"supplierBankAccountID,omitempty"`
	SupplierCustomerID *string `json:"supplierCustomerID,omitempty"`
	SupplierMandateID *string `json:"supplierMandateID,omitempty"`
	SupplierName *string `json:"supplierName,omitempty"`
	SupplierStatus *string `json:"supplierStatus,omitempty"`
}

// MerchantDirectDebitMandatePageListMatch is the typed request payload for MerchantDirectDebitMandatePage.ListTyped.
type MerchantDirectDebitMandatePageListMatch struct {
	ApprovedAt *string `json:"approvedAt,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomerAccountNumber *string `json:"customerAccountNumber,omitempty"`
	CustomerCity *string `json:"customerCity,omitempty"`
	CustomerCountryCode *string `json:"customerCountryCode,omitempty"`
	CustomerCountryName *string `json:"customerCountryName,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerFirstName *string `json:"customerFirstName,omitempty"`
	CustomerIban *string `json:"customerIban,omitempty"`
	CustomerLastName *string `json:"customerLastName,omitempty"`
	CustomerSortCode *string `json:"customerSortCode,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsRecurring *bool `json:"isRecurring,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Reference *string `json:"reference,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplierBankAccountID *string `json:"supplierBankAccountID,omitempty"`
	SupplierCustomerID *string `json:"supplierCustomerID,omitempty"`
	SupplierMandateID *string `json:"supplierMandateID,omitempty"`
	SupplierName *string `json:"supplierName,omitempty"`
	SupplierStatus *string `json:"supplierStatus,omitempty"`
}

// MerchantPayByBankSetting is the typed data model for the merchant_pay_by_bank_setting entity.
type MerchantPayByBankSetting struct {
	BankCountryCodes *[]any `json:"bankCountryCodes,omitempty"`
	BankID *string `json:"bankID,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	BusinessInstitutionID *string `json:"businessInstitutionID,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Message *string `json:"message,omitempty"`
	MessageImageUrl *string `json:"messageImageUrl,omitempty"`
	Order *int `json:"order,omitempty"`
	PersonalInstitutionID *string `json:"personalInstitutionID,omitempty"`
	Processor *string `json:"processor,omitempty"`
	WarningHeading *string `json:"warningHeading,omitempty"`
	WarningMessage *string `json:"warningMessage,omitempty"`
}

// MerchantPayByBankSettingListMatch is the typed request payload for MerchantPayByBankSetting.ListTyped.
type MerchantPayByBankSettingListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// MerchantPaymentRequestTemplate is the typed data model for the merchant_payment_request_template entity.
type MerchantPaymentRequestTemplate struct {
	BankPaymentOptions *map[string]any `json:"bankPaymentOptions,omitempty"`
	CardPaymentAddressOptions *map[string]any `json:"cardPaymentAddressOptions,omitempty"`
	CardPaymentCaptureOptions *map[string]any `json:"cardPaymentCaptureOptions,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	DefaultFields *[]any `json:"defaultFields,omitempty"`
	Description string `json:"description"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name string `json:"name"`
	NotificationOptions *map[string]any `json:"notificationOptions,omitempty"`
	PaymentMethods *map[string]any `json:"paymentMethods,omitempty"`
	PaymentTerms *map[string]any `json:"paymentTerms,omitempty"`
	PriorityBankOptions *map[string]any `json:"priorityBankOptions,omitempty"`
	Template map[string]any `json:"template"`
}

// MerchantPaymentRequestTemplateLoadMatch is the typed request payload for MerchantPaymentRequestTemplate.LoadTyped.
type MerchantPaymentRequestTemplateLoadMatch struct {
	Id string `json:"id"`
	PaymentrequestId string `json:"paymentrequest_id"`
}

// MerchantPaymentRequestTemplateListMatch is the typed request payload for MerchantPaymentRequestTemplate.ListTyped.
type MerchantPaymentRequestTemplateListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// MerchantPaymentRequestTemplateUpdateData is the typed request payload for MerchantPaymentRequestTemplate.UpdateTyped.
type MerchantPaymentRequestTemplateUpdateData struct {
	Id string `json:"id"`
	PaymentrequestId string `json:"paymentrequest_id"`
	BankPaymentOptions *map[string]any `json:"bankPaymentOptions,omitempty"`
	CardPaymentAddressOptions *map[string]any `json:"cardPaymentAddressOptions,omitempty"`
	CardPaymentCaptureOptions *map[string]any `json:"cardPaymentCaptureOptions,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	DefaultFields *[]any `json:"defaultFields,omitempty"`
	Description *string `json:"description,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	NotificationOptions *map[string]any `json:"notificationOptions,omitempty"`
	PaymentMethods *map[string]any `json:"paymentMethods,omitempty"`
	PaymentTerms *map[string]any `json:"paymentTerms,omitempty"`
	PriorityBankOptions *map[string]any `json:"priorityBankOptions,omitempty"`
	Template *map[string]any `json:"template,omitempty"`
}

// MerchantPaymentRequestTemplateRemoveMatch is the typed request payload for MerchantPaymentRequestTemplate.RemoveTyped.
type MerchantPaymentRequestTemplateRemoveMatch struct {
	Id string `json:"id"`
	PaymentrequestId string `json:"paymentrequest_id"`
}

// MerchantToken is the typed data model for the merchant_token entity.
type MerchantToken struct {
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	HmacAlgorithm *string `json:"hmacAlgorithm,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IpAddressWhitelist *string `json:"ipAddressWhitelist,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	LastAuthorised *string `json:"lastAuthorised,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Nonce string `json:"nonce"`
	PermissionTypes *[]any `json:"permissionTypes,omitempty"`
	RequestSignatureVersion *int `json:"requestSignatureVersion,omitempty"`
	SharedSecretAlgorithm *string `json:"sharedSecretAlgorithm,omitempty"`
	SharedSecretBase64 *string `json:"sharedSecretBase64,omitempty"`
	Token *string `json:"token,omitempty"`
}

// MerchantTokenLoadMatch is the typed request payload for MerchantToken.LoadTyped.
type MerchantTokenLoadMatch struct {
	Id string `json:"id"`
}

// MerchantTokenListMatch is the typed request payload for MerchantToken.ListTyped.
type MerchantTokenListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// MerchantTokenCreateData is the typed request payload for MerchantToken.CreateTyped.
type MerchantTokenCreateData struct {
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	HmacAlgorithm *string `json:"hmacAlgorithm,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IpAddressWhitelist *string `json:"ipAddressWhitelist,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	LastAuthorised *string `json:"lastAuthorised,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Nonce string `json:"nonce"`
	PermissionTypes *[]any `json:"permissionTypes,omitempty"`
	RequestSignatureVersion *int `json:"requestSignatureVersion,omitempty"`
	SharedSecretAlgorithm *string `json:"sharedSecretAlgorithm,omitempty"`
	SharedSecretBase64 *string `json:"sharedSecretBase64,omitempty"`
	Token *string `json:"token,omitempty"`
}

// MerchantTokenUpdateData is the typed request payload for MerchantToken.UpdateTyped.
type MerchantTokenUpdateData struct {
	Id string `json:"id"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpiresAt *string `json:"expiresAt,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	HmacAlgorithm *string `json:"hmacAlgorithm,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IpAddressWhitelist *string `json:"ipAddressWhitelist,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	LastAuthorised *string `json:"lastAuthorised,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PermissionTypes *[]any `json:"permissionTypes,omitempty"`
	RequestSignatureVersion *int `json:"requestSignatureVersion,omitempty"`
	SharedSecretAlgorithm *string `json:"sharedSecretAlgorithm,omitempty"`
	SharedSecretBase64 *string `json:"sharedSecretBase64,omitempty"`
	Token *string `json:"token,omitempty"`
}

// Metadata is the typed data model for the metadata entity.
type Metadata struct {
}

// MetadataLoadMatch is the typed request payload for Metadata.LoadTyped.
type MetadataLoadMatch struct {
}

// NoFrixionVersion is the typed data model for the no_frixion_version entity.
type NoFrixionVersion struct {
	BuildVersion *int `json:"buildVersion,omitempty"`
	MajorVersion *int `json:"majorVersion,omitempty"`
	MinorVersion *int `json:"minorVersion,omitempty"`
	ReleaseName *string `json:"releaseName,omitempty"`
}

// NoFrixionVersionLoadMatch is the typed request payload for NoFrixionVersion.LoadTyped.
type NoFrixionVersionLoadMatch struct {
	BuildVersion *int `json:"buildVersion,omitempty"`
	MajorVersion *int `json:"majorVersion,omitempty"`
	MinorVersion *int `json:"minorVersion,omitempty"`
	ReleaseName *string `json:"releaseName,omitempty"`
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
	AccountId string `json:"account_id"`
}

// Payeeverification is the typed data model for the payeeverification entity.
type Payeeverification struct {
	AccountName string `json:"accountName"`
	AccountNumber *string `json:"accountNumber,omitempty"`
	Iban string `json:"iban"`
	PayeeVerifiedAccountName *string `json:"payeeVerifiedAccountName,omitempty"`
	Result *string `json:"result,omitempty"`
	SecondaryIdentification *string `json:"secondaryIdentification,omitempty"`
	SortCode *string `json:"sortCode,omitempty"`
}

// PayeeverificationCreateData is the typed request payload for Payeeverification.CreateTyped.
type PayeeverificationCreateData struct {
	AccountName string `json:"accountName"`
	AccountNumber *string `json:"accountNumber,omitempty"`
	Iban string `json:"iban"`
	PayeeVerifiedAccountName *string `json:"payeeVerifiedAccountName,omitempty"`
	Result *string `json:"result,omitempty"`
	SecondaryIdentification *string `json:"secondaryIdentification,omitempty"`
	SortCode *string `json:"sortCode,omitempty"`
}

// Payment is the typed data model for the payment entity.
type Payment struct {
	Addresses *[]any `json:"addresses,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	AutoSendReceipt *bool `json:"autoSendReceipt,omitempty"`
	BaseOriginUrl *string `json:"baseOriginUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	CardAuthorizeOnly *bool `json:"cardAuthorizeOnly,omitempty"`
	CardCreateToken *bool `json:"cardCreateToken,omitempty"`
	CardCreateTokenMode *string `json:"cardCreateTokenMode,omitempty"`
	CardIgnoreCVN *bool `json:"cardIgnoreCVN,omitempty"`
	CardNoPayerAuthentication *bool `json:"cardNoPayerAuthentication,omitempty"`
	CardProcessorMerchantID *string `json:"cardProcessorMerchantID,omitempty"`
	CardStripePaymentIntentID *string `json:"cardStripePaymentIntentID,omitempty"`
	CardStripePaymentIntentSecret *string `json:"cardStripePaymentIntentSecret,omitempty"`
	CardTransmitRawDetails *bool `json:"cardTransmitRawDetails,omitempty"`
	CreatedByUser map[string]any `json:"createdByUser"`
	Currency *string `json:"currency,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerID *string `json:"customerID,omitempty"`
	CustomerName *string `json:"customerName,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destinationAccount,omitempty"`
	DirectDebitPayment *map[string]any `json:"directDebitPayment,omitempty"`
	DueDate *string `json:"dueDate,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	FieldDisplaySettings *[]any `json:"fieldDisplaySettings,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hostedPayCheckoutUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignoreAddressVerification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"insertedSortable,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LightningInvoice *string `json:"lightningInvoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightningInvoiceExpiresAt,omitempty"`
	MerchantDirectDebitMandateID *string `json:"merchantDirectDebitMandateID,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	NotificationEmailAddresses *string `json:"notificationEmailAddresses,omitempty"`
	NotificationRoleIDs *[]any `json:"notificationRoleIDs,omitempty"`
	OrderID *string `json:"orderID,omitempty"`
	PartialPaymentMethod *string `json:"partialPaymentMethod,omitempty"`
	PartialPaymentSteps *string `json:"partialPaymentSteps,omitempty"`
	PaymentAttempts *[]any `json:"paymentAttempts,omitempty"`
	PaymentMethods *[]any `json:"paymentMethods,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PispAccountID *string `json:"pispAccountID,omitempty"`
	PriorityBankID *string `json:"priorityBankID,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSeconds *int `json:"sandboxSettleDelayInSeconds,omitempty"`
	ShippingAddress *map[string]any `json:"shippingAddress,omitempty"`
	ShippingAddressCity *string `json:"shippingAddressCity,omitempty"`
	ShippingAddressCountryCode *string `json:"shippingAddressCountryCode,omitempty"`
	ShippingAddressCounty *string `json:"shippingAddressCounty,omitempty"`
	ShippingAddressLine1 *string `json:"shippingAddressLine1,omitempty"`
	ShippingAddressLine2 *string `json:"shippingAddressLine2,omitempty"`
	ShippingAddressPostCode *string `json:"shippingAddressPostCode,omitempty"`
	ShippingEmail *string `json:"shippingEmail,omitempty"`
	ShippingFirstName *string `json:"shippingFirstName,omitempty"`
	ShippingLastName *string `json:"shippingLastName,omitempty"`
	ShippingPhone *string `json:"shippingPhone,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
	TagIds *[]any `json:"tagIds,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCards *[]any `json:"tokenisedCards,omitempty"`
	Transactions *[]any `json:"transactions,omitempty"`
	UseHostedPaymentPage *bool `json:"useHostedPaymentPage,omitempty"`
}

// PaymentLoadMatch is the typed request payload for Payment.LoadTyped.
type PaymentLoadMatch struct {
	Id string `json:"id"`
}

// PaymentCreateData is the typed request payload for Payment.CreateTyped.
type PaymentCreateData struct {
	Addresses *[]any `json:"addresses,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	AutoSendReceipt *bool `json:"autoSendReceipt,omitempty"`
	BaseOriginUrl *string `json:"baseOriginUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	CardAuthorizeOnly *bool `json:"cardAuthorizeOnly,omitempty"`
	CardCreateToken *bool `json:"cardCreateToken,omitempty"`
	CardCreateTokenMode *string `json:"cardCreateTokenMode,omitempty"`
	CardIgnoreCVN *bool `json:"cardIgnoreCVN,omitempty"`
	CardNoPayerAuthentication *bool `json:"cardNoPayerAuthentication,omitempty"`
	CardProcessorMerchantID *string `json:"cardProcessorMerchantID,omitempty"`
	CardStripePaymentIntentID *string `json:"cardStripePaymentIntentID,omitempty"`
	CardStripePaymentIntentSecret *string `json:"cardStripePaymentIntentSecret,omitempty"`
	CardTransmitRawDetails *bool `json:"cardTransmitRawDetails,omitempty"`
	CreatedByUser map[string]any `json:"createdByUser"`
	Currency *string `json:"currency,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerID *string `json:"customerID,omitempty"`
	CustomerName *string `json:"customerName,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destinationAccount,omitempty"`
	DirectDebitPayment *map[string]any `json:"directDebitPayment,omitempty"`
	DueDate *string `json:"dueDate,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	FieldDisplaySettings *[]any `json:"fieldDisplaySettings,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hostedPayCheckoutUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignoreAddressVerification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"insertedSortable,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LightningInvoice *string `json:"lightningInvoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightningInvoiceExpiresAt,omitempty"`
	MerchantDirectDebitMandateID *string `json:"merchantDirectDebitMandateID,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	NotificationEmailAddresses *string `json:"notificationEmailAddresses,omitempty"`
	NotificationRoleIDs *[]any `json:"notificationRoleIDs,omitempty"`
	OrderID *string `json:"orderID,omitempty"`
	PartialPaymentMethod *string `json:"partialPaymentMethod,omitempty"`
	PartialPaymentSteps *string `json:"partialPaymentSteps,omitempty"`
	PaymentAttempts *[]any `json:"paymentAttempts,omitempty"`
	PaymentMethods *[]any `json:"paymentMethods,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PispAccountID *string `json:"pispAccountID,omitempty"`
	PriorityBankID *string `json:"priorityBankID,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSeconds *int `json:"sandboxSettleDelayInSeconds,omitempty"`
	ShippingAddress *map[string]any `json:"shippingAddress,omitempty"`
	ShippingAddressCity *string `json:"shippingAddressCity,omitempty"`
	ShippingAddressCountryCode *string `json:"shippingAddressCountryCode,omitempty"`
	ShippingAddressCounty *string `json:"shippingAddressCounty,omitempty"`
	ShippingAddressLine1 *string `json:"shippingAddressLine1,omitempty"`
	ShippingAddressLine2 *string `json:"shippingAddressLine2,omitempty"`
	ShippingAddressPostCode *string `json:"shippingAddressPostCode,omitempty"`
	ShippingEmail *string `json:"shippingEmail,omitempty"`
	ShippingFirstName *string `json:"shippingFirstName,omitempty"`
	ShippingLastName *string `json:"shippingLastName,omitempty"`
	ShippingPhone *string `json:"shippingPhone,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
	TagIds *[]any `json:"tagIds,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCards *[]any `json:"tokenisedCards,omitempty"`
	Transactions *[]any `json:"transactions,omitempty"`
	UseHostedPaymentPage *bool `json:"useHostedPaymentPage,omitempty"`
}

// PaymentUpdateData is the typed request payload for Payment.UpdateTyped.
type PaymentUpdateData struct {
	Id string `json:"id"`
	Addresses *[]any `json:"addresses,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	AutoSendReceipt *bool `json:"autoSendReceipt,omitempty"`
	BaseOriginUrl *string `json:"baseOriginUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	CardAuthorizeOnly *bool `json:"cardAuthorizeOnly,omitempty"`
	CardCreateToken *bool `json:"cardCreateToken,omitempty"`
	CardCreateTokenMode *string `json:"cardCreateTokenMode,omitempty"`
	CardIgnoreCVN *bool `json:"cardIgnoreCVN,omitempty"`
	CardNoPayerAuthentication *bool `json:"cardNoPayerAuthentication,omitempty"`
	CardProcessorMerchantID *string `json:"cardProcessorMerchantID,omitempty"`
	CardStripePaymentIntentID *string `json:"cardStripePaymentIntentID,omitempty"`
	CardStripePaymentIntentSecret *string `json:"cardStripePaymentIntentSecret,omitempty"`
	CardTransmitRawDetails *bool `json:"cardTransmitRawDetails,omitempty"`
	CreatedByUser *map[string]any `json:"createdByUser,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerID *string `json:"customerID,omitempty"`
	CustomerName *string `json:"customerName,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destinationAccount,omitempty"`
	DirectDebitPayment *map[string]any `json:"directDebitPayment,omitempty"`
	DueDate *string `json:"dueDate,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	FieldDisplaySettings *[]any `json:"fieldDisplaySettings,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hostedPayCheckoutUrl,omitempty"`
	IgnoreAddressVerification *bool `json:"ignoreAddressVerification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"insertedSortable,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LightningInvoice *string `json:"lightningInvoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightningInvoiceExpiresAt,omitempty"`
	MerchantDirectDebitMandateID *string `json:"merchantDirectDebitMandateID,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	NotificationEmailAddresses *string `json:"notificationEmailAddresses,omitempty"`
	NotificationRoleIDs *[]any `json:"notificationRoleIDs,omitempty"`
	OrderID *string `json:"orderID,omitempty"`
	PartialPaymentMethod *string `json:"partialPaymentMethod,omitempty"`
	PartialPaymentSteps *string `json:"partialPaymentSteps,omitempty"`
	PaymentAttempts *[]any `json:"paymentAttempts,omitempty"`
	PaymentMethods *[]any `json:"paymentMethods,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PispAccountID *string `json:"pispAccountID,omitempty"`
	PriorityBankID *string `json:"priorityBankID,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSeconds *int `json:"sandboxSettleDelayInSeconds,omitempty"`
	ShippingAddress *map[string]any `json:"shippingAddress,omitempty"`
	ShippingAddressCity *string `json:"shippingAddressCity,omitempty"`
	ShippingAddressCountryCode *string `json:"shippingAddressCountryCode,omitempty"`
	ShippingAddressCounty *string `json:"shippingAddressCounty,omitempty"`
	ShippingAddressLine1 *string `json:"shippingAddressLine1,omitempty"`
	ShippingAddressLine2 *string `json:"shippingAddressLine2,omitempty"`
	ShippingAddressPostCode *string `json:"shippingAddressPostCode,omitempty"`
	ShippingEmail *string `json:"shippingEmail,omitempty"`
	ShippingFirstName *string `json:"shippingFirstName,omitempty"`
	ShippingLastName *string `json:"shippingLastName,omitempty"`
	ShippingPhone *string `json:"shippingPhone,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
	TagIds *[]any `json:"tagIds,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCards *[]any `json:"tokenisedCards,omitempty"`
	Transactions *[]any `json:"transactions,omitempty"`
	UseHostedPaymentPage *bool `json:"useHostedPaymentPage,omitempty"`
}

// PaymentAccount is the typed data model for the payment_account entity.
type PaymentAccount struct {
	AccountName *string `json:"accountName,omitempty"`
	AccountSupplierName *string `json:"accountSupplierName,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	AvailableBalanceMinorUnits *int `json:"availableBalanceMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	CreatedBy map[string]any `json:"createdBy"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"defaultPaymentRail,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	ExternalAccountIcon *string `json:"externalAccountIcon,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsTrustAccount *bool `json:"isTrustAccount,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	LastTransaction *map[string]any `json:"lastTransaction,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	PhysicalAccountID *string `json:"physicalAccountID,omitempty"`
	Rules *[]any `json:"rules,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
	SubmittedPayoutsBalanceMinorUnits *int `json:"submittedPayoutsBalanceMinorUnits,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplierSepaInstantStatus,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xeroBankFeedConnectionStatus,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xeroBankFeedLastSyncedAt,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xeroBankFeedSyncLastFailedAt,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xeroBankFeedSyncLastFailureReason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xeroBankFeedSyncStatus,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xeroUnsynchronisedTransactionsCount,omitempty"`
}

// PaymentAccountListMatch is the typed request payload for PaymentAccount.ListTyped.
type PaymentAccountListMatch struct {
	AccountName *string `json:"accountName,omitempty"`
	AccountSupplierName *string `json:"accountSupplierName,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	AvailableBalanceMinorUnits *int `json:"availableBalanceMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	CreatedBy *map[string]any `json:"createdBy,omitempty"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"defaultPaymentRail,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	ExternalAccountIcon *string `json:"externalAccountIcon,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier *map[string]any `json:"identifier,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsTrustAccount *bool `json:"isTrustAccount,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	LastTransaction *map[string]any `json:"lastTransaction,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	PhysicalAccountID *string `json:"physicalAccountID,omitempty"`
	Rules *[]any `json:"rules,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
	SubmittedPayoutsBalanceMinorUnits *int `json:"submittedPayoutsBalanceMinorUnits,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplierSepaInstantStatus,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xeroBankFeedConnectionStatus,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xeroBankFeedLastSyncedAt,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xeroBankFeedSyncLastFailedAt,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xeroBankFeedSyncLastFailureReason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xeroBankFeedSyncStatus,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xeroUnsynchronisedTransactionsCount,omitempty"`
}

// PaymentAccountMinimal is the typed data model for the payment_account_minimal entity.
type PaymentAccountMinimal struct {
	AccountName *string `json:"accountName,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
}

// PaymentAccountMinimalListMatch is the typed request payload for PaymentAccountMinimal.ListTyped.
type PaymentAccountMinimalListMatch struct {
	AccountName *string `json:"accountName,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier *map[string]any `json:"identifier,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
}

// PaymentInitiation is the typed data model for the payment_initiation entity.
type PaymentInitiation struct {
	PaymentInitiationID *string `json:"paymentInitiationID,omitempty"`
	PaymentRequestCallbackUrl *string `json:"paymentRequestCallbackUrl,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	SpecificErrorMessage *string `json:"specificErrorMessage,omitempty"`
}

// PaymentInitiationCreateData is the typed request payload for PaymentInitiation.CreateTyped.
type PaymentInitiationCreateData struct {
	PaymentrequestId string `json:"paymentrequest_id"`
	PaymentInitiationID *string `json:"paymentInitiationID,omitempty"`
	PaymentRequestCallbackUrl *string `json:"paymentRequestCallbackUrl,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	ResponseType *string `json:"responseType,omitempty"`
	SpecificErrorMessage *string `json:"specificErrorMessage,omitempty"`
}

// PaymentRequest is the typed data model for the payment_request entity.
type PaymentRequest struct {
	Addresses *[]any `json:"addresses,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	AutoSendReceipt *bool `json:"autoSendReceipt,omitempty"`
	BaseOriginUrl *string `json:"baseOriginUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	CardAuthorizeOnly *bool `json:"cardAuthorizeOnly,omitempty"`
	CardCreateToken *bool `json:"cardCreateToken,omitempty"`
	CardCreateTokenMode *string `json:"cardCreateTokenMode,omitempty"`
	CardIgnoreCVN *bool `json:"cardIgnoreCVN,omitempty"`
	CardProcessorMerchantID *string `json:"cardProcessorMerchantID,omitempty"`
	CardStripePaymentIntentID *string `json:"cardStripePaymentIntentID,omitempty"`
	CardStripePaymentIntentSecret *string `json:"cardStripePaymentIntentSecret,omitempty"`
	CreatedByUser map[string]any `json:"createdByUser"`
	Currency *string `json:"currency,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerID *string `json:"customerID,omitempty"`
	CustomerName *string `json:"customerName,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destinationAccount,omitempty"`
	DirectDebitPayment *map[string]any `json:"directDebitPayment,omitempty"`
	DoSimulateSettlementFailure *bool `json:"doSimulateSettlementFailure,omitempty"`
	DueDate *string `json:"dueDate,omitempty"`
	ErrorDescription *string `json:"errorDescription,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailedPaymentRequests *map[string]any `json:"failedPaymentRequests,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	FieldDisplaySettings *[]any `json:"fieldDisplaySettings,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hostedPayCheckoutUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignoreAddressVerification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"insertedSortable,omitempty"`
	Institution *string `json:"institution,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LightningInvoice *string `json:"lightningInvoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightningInvoiceExpiresAt,omitempty"`
	MerchantDirectDebitMandateID *string `json:"merchantDirectDebitMandateID,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	NotificationEmailAddresses *string `json:"notificationEmailAddresses,omitempty"`
	NotificationRoleIDs *[]any `json:"notificationRoleIDs,omitempty"`
	OrderID *string `json:"orderID,omitempty"`
	PartialPaymentMethod *string `json:"partialPaymentMethod,omitempty"`
	PartialPaymentSteps *string `json:"partialPaymentSteps,omitempty"`
	PaymentAttempts *[]any `json:"paymentAttempts,omitempty"`
	PaymentInitiationID *string `json:"paymentInitiationID,omitempty"`
	PaymentMethods *[]any `json:"paymentMethods,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRequests *[]any `json:"paymentRequests,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PispAccountID *string `json:"pispAccountID,omitempty"`
	PriorityBankID *string `json:"priorityBankID,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSeconds *int `json:"sandboxSettleDelayInSeconds,omitempty"`
	ShippingAddress *map[string]any `json:"shippingAddress,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCards *[]any `json:"tokenisedCards,omitempty"`
	Transactions *[]any `json:"transactions,omitempty"`
	UseHostedPaymentPage *bool `json:"useHostedPaymentPage,omitempty"`
}

// PaymentRequestLoadMatch is the typed request payload for PaymentRequest.LoadTyped.
type PaymentRequestLoadMatch struct {
	Addresses *[]any `json:"addresses,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	AutoSendReceipt *bool `json:"autoSendReceipt,omitempty"`
	BaseOriginUrl *string `json:"baseOriginUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	CardAuthorizeOnly *bool `json:"cardAuthorizeOnly,omitempty"`
	CardCreateToken *bool `json:"cardCreateToken,omitempty"`
	CardCreateTokenMode *string `json:"cardCreateTokenMode,omitempty"`
	CardIgnoreCVN *bool `json:"cardIgnoreCVN,omitempty"`
	CardProcessorMerchantID *string `json:"cardProcessorMerchantID,omitempty"`
	CardStripePaymentIntentID *string `json:"cardStripePaymentIntentID,omitempty"`
	CardStripePaymentIntentSecret *string `json:"cardStripePaymentIntentSecret,omitempty"`
	CreatedByUser *map[string]any `json:"createdByUser,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerID *string `json:"customerID,omitempty"`
	CustomerName *string `json:"customerName,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destinationAccount,omitempty"`
	DirectDebitPayment *map[string]any `json:"directDebitPayment,omitempty"`
	DoSimulateSettlementFailure *bool `json:"doSimulateSettlementFailure,omitempty"`
	DueDate *string `json:"dueDate,omitempty"`
	ErrorDescription *string `json:"errorDescription,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailedPaymentRequests *map[string]any `json:"failedPaymentRequests,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	FieldDisplaySettings *[]any `json:"fieldDisplaySettings,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hostedPayCheckoutUrl,omitempty"`
	Id string `json:"id"`
	IgnoreAddressVerification *bool `json:"ignoreAddressVerification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"insertedSortable,omitempty"`
	Institution *string `json:"institution,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LightningInvoice *string `json:"lightningInvoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightningInvoiceExpiresAt,omitempty"`
	MerchantDirectDebitMandateID *string `json:"merchantDirectDebitMandateID,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	NotificationEmailAddresses *string `json:"notificationEmailAddresses,omitempty"`
	NotificationRoleIDs *[]any `json:"notificationRoleIDs,omitempty"`
	OrderID *string `json:"orderID,omitempty"`
	PartialPaymentMethod *string `json:"partialPaymentMethod,omitempty"`
	PartialPaymentSteps *string `json:"partialPaymentSteps,omitempty"`
	PaymentAttempts *[]any `json:"paymentAttempts,omitempty"`
	PaymentInitiationID *string `json:"paymentInitiationID,omitempty"`
	PaymentMethods *[]any `json:"paymentMethods,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRequests *[]any `json:"paymentRequests,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PispAccountID *string `json:"pispAccountID,omitempty"`
	PriorityBankID *string `json:"priorityBankID,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSeconds *int `json:"sandboxSettleDelayInSeconds,omitempty"`
	ShippingAddress *map[string]any `json:"shippingAddress,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCards *[]any `json:"tokenisedCards,omitempty"`
	Transactions *[]any `json:"transactions,omitempty"`
	UseHostedPaymentPage *bool `json:"useHostedPaymentPage,omitempty"`
}

// PaymentRequestListMatch is the typed request payload for PaymentRequest.ListTyped.
type PaymentRequestListMatch struct {
	Addresses *[]any `json:"addresses,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	AutoSendReceipt *bool `json:"autoSendReceipt,omitempty"`
	BaseOriginUrl *string `json:"baseOriginUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	CardAuthorizeOnly *bool `json:"cardAuthorizeOnly,omitempty"`
	CardCreateToken *bool `json:"cardCreateToken,omitempty"`
	CardCreateTokenMode *string `json:"cardCreateTokenMode,omitempty"`
	CardIgnoreCVN *bool `json:"cardIgnoreCVN,omitempty"`
	CardProcessorMerchantID *string `json:"cardProcessorMerchantID,omitempty"`
	CardStripePaymentIntentID *string `json:"cardStripePaymentIntentID,omitempty"`
	CardStripePaymentIntentSecret *string `json:"cardStripePaymentIntentSecret,omitempty"`
	CreatedByUser *map[string]any `json:"createdByUser,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerID *string `json:"customerID,omitempty"`
	CustomerName *string `json:"customerName,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destinationAccount,omitempty"`
	DirectDebitPayment *map[string]any `json:"directDebitPayment,omitempty"`
	DoSimulateSettlementFailure *bool `json:"doSimulateSettlementFailure,omitempty"`
	DueDate *string `json:"dueDate,omitempty"`
	ErrorDescription *string `json:"errorDescription,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailedPaymentRequests *map[string]any `json:"failedPaymentRequests,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	FieldDisplaySettings *[]any `json:"fieldDisplaySettings,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hostedPayCheckoutUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignoreAddressVerification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"insertedSortable,omitempty"`
	Institution *string `json:"institution,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LightningInvoice *string `json:"lightningInvoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightningInvoiceExpiresAt,omitempty"`
	MerchantDirectDebitMandateID *string `json:"merchantDirectDebitMandateID,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	NotificationEmailAddresses *string `json:"notificationEmailAddresses,omitempty"`
	NotificationRoleIDs *[]any `json:"notificationRoleIDs,omitempty"`
	OrderID *string `json:"orderID,omitempty"`
	PartialPaymentMethod *string `json:"partialPaymentMethod,omitempty"`
	PartialPaymentSteps *string `json:"partialPaymentSteps,omitempty"`
	PaymentAttempts *[]any `json:"paymentAttempts,omitempty"`
	PaymentInitiationID *string `json:"paymentInitiationID,omitempty"`
	PaymentMethods *[]any `json:"paymentMethods,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRequests *[]any `json:"paymentRequests,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PispAccountID *string `json:"pispAccountID,omitempty"`
	PriorityBankID *string `json:"priorityBankID,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSeconds *int `json:"sandboxSettleDelayInSeconds,omitempty"`
	ShippingAddress *map[string]any `json:"shippingAddress,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCards *[]any `json:"tokenisedCards,omitempty"`
	Transactions *[]any `json:"transactions,omitempty"`
	UseHostedPaymentPage *bool `json:"useHostedPaymentPage,omitempty"`
}

// PaymentRequestCreateData is the typed request payload for PaymentRequest.CreateTyped.
type PaymentRequestCreateData struct {
	Addresses *[]any `json:"addresses,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	AutoSendReceipt *bool `json:"autoSendReceipt,omitempty"`
	BaseOriginUrl *string `json:"baseOriginUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	CardAuthorizeOnly *bool `json:"cardAuthorizeOnly,omitempty"`
	CardCreateToken *bool `json:"cardCreateToken,omitempty"`
	CardCreateTokenMode *string `json:"cardCreateTokenMode,omitempty"`
	CardIgnoreCVN *bool `json:"cardIgnoreCVN,omitempty"`
	CardProcessorMerchantID *string `json:"cardProcessorMerchantID,omitempty"`
	CardStripePaymentIntentID *string `json:"cardStripePaymentIntentID,omitempty"`
	CardStripePaymentIntentSecret *string `json:"cardStripePaymentIntentSecret,omitempty"`
	CreatedByUser map[string]any `json:"createdByUser"`
	Currency *string `json:"currency,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerID *string `json:"customerID,omitempty"`
	CustomerName *string `json:"customerName,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destinationAccount,omitempty"`
	DirectDebitPayment *map[string]any `json:"directDebitPayment,omitempty"`
	DoSimulateSettlementFailure *bool `json:"doSimulateSettlementFailure,omitempty"`
	DueDate *string `json:"dueDate,omitempty"`
	ErrorDescription *string `json:"errorDescription,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailedPaymentRequests *map[string]any `json:"failedPaymentRequests,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	FieldDisplaySettings *[]any `json:"fieldDisplaySettings,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hostedPayCheckoutUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignoreAddressVerification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"insertedSortable,omitempty"`
	Institution *string `json:"institution,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LightningInvoice *string `json:"lightningInvoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightningInvoiceExpiresAt,omitempty"`
	MerchantDirectDebitMandateID *string `json:"merchantDirectDebitMandateID,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	NotificationEmailAddresses *string `json:"notificationEmailAddresses,omitempty"`
	NotificationRoleIDs *[]any `json:"notificationRoleIDs,omitempty"`
	OrderID *string `json:"orderID,omitempty"`
	PartialPaymentMethod *string `json:"partialPaymentMethod,omitempty"`
	PartialPaymentSteps *string `json:"partialPaymentSteps,omitempty"`
	PaymentAttempts *[]any `json:"paymentAttempts,omitempty"`
	PaymentInitiationID *string `json:"paymentInitiationID,omitempty"`
	PaymentMethods *[]any `json:"paymentMethods,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRequests *[]any `json:"paymentRequests,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PispAccountID *string `json:"pispAccountID,omitempty"`
	PriorityBankID *string `json:"priorityBankID,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSeconds *int `json:"sandboxSettleDelayInSeconds,omitempty"`
	ShippingAddress *map[string]any `json:"shippingAddress,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCards *[]any `json:"tokenisedCards,omitempty"`
	Transactions *[]any `json:"transactions,omitempty"`
	UseHostedPaymentPage *bool `json:"useHostedPaymentPage,omitempty"`
}

// PaymentRequestUpdateData is the typed request payload for PaymentRequest.UpdateTyped.
type PaymentRequestUpdateData struct {
	PaymentrequestId string `json:"paymentrequest_id"`
	Addresses *[]any `json:"addresses,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	AutoSendReceipt *bool `json:"autoSendReceipt,omitempty"`
	BaseOriginUrl *string `json:"baseOriginUrl,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	CardAuthorizeOnly *bool `json:"cardAuthorizeOnly,omitempty"`
	CardCreateToken *bool `json:"cardCreateToken,omitempty"`
	CardCreateTokenMode *string `json:"cardCreateTokenMode,omitempty"`
	CardIgnoreCVN *bool `json:"cardIgnoreCVN,omitempty"`
	CardProcessorMerchantID *string `json:"cardProcessorMerchantID,omitempty"`
	CardStripePaymentIntentID *string `json:"cardStripePaymentIntentID,omitempty"`
	CardStripePaymentIntentSecret *string `json:"cardStripePaymentIntentSecret,omitempty"`
	CreatedByUser *map[string]any `json:"createdByUser,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomFields *[]any `json:"customFields,omitempty"`
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	CustomerID *string `json:"customerID,omitempty"`
	CustomerName *string `json:"customerName,omitempty"`
	Description *string `json:"description,omitempty"`
	DestinationAccount *map[string]any `json:"destinationAccount,omitempty"`
	DirectDebitPayment *map[string]any `json:"directDebitPayment,omitempty"`
	DoSimulateSettlementFailure *bool `json:"doSimulateSettlementFailure,omitempty"`
	DueDate *string `json:"dueDate,omitempty"`
	ErrorDescription *string `json:"errorDescription,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailedPaymentRequests *map[string]any `json:"failedPaymentRequests,omitempty"`
	FailureCallbackUrl *string `json:"failureCallbackUrl,omitempty"`
	FieldDisplaySettings *[]any `json:"fieldDisplaySettings,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	HostedPayCheckoutUrl *string `json:"hostedPayCheckoutUrl,omitempty"`
	Id *string `json:"id,omitempty"`
	IgnoreAddressVerification *bool `json:"ignoreAddressVerification,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InsertedSortable *string `json:"insertedSortable,omitempty"`
	Institution *string `json:"institution,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LightningInvoice *string `json:"lightningInvoice,omitempty"`
	LightningInvoiceExpiresAt *string `json:"lightningInvoiceExpiresAt,omitempty"`
	MerchantDirectDebitMandateID *string `json:"merchantDirectDebitMandateID,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	NotificationEmailAddresses *string `json:"notificationEmailAddresses,omitempty"`
	NotificationRoleIDs *[]any `json:"notificationRoleIDs,omitempty"`
	OrderID *string `json:"orderID,omitempty"`
	PartialPaymentMethod *string `json:"partialPaymentMethod,omitempty"`
	PartialPaymentSteps *string `json:"partialPaymentSteps,omitempty"`
	PaymentAttempts *[]any `json:"paymentAttempts,omitempty"`
	PaymentInitiationID *string `json:"paymentInitiationID,omitempty"`
	PaymentMethods *[]any `json:"paymentMethods,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRequests *[]any `json:"paymentRequests,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PispAccountID *string `json:"pispAccountID,omitempty"`
	PriorityBankID *string `json:"priorityBankID,omitempty"`
	Result *map[string]any `json:"result,omitempty"`
	SandboxSettleDelayInSeconds *int `json:"sandboxSettleDelayInSeconds,omitempty"`
	ShippingAddress *map[string]any `json:"shippingAddress,omitempty"`
	Status *string `json:"status,omitempty"`
	SuccessWebHookUrl *string `json:"successWebHookUrl,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	Title *string `json:"title,omitempty"`
	TokenisedCards *[]any `json:"tokenisedCards,omitempty"`
	Transactions *[]any `json:"transactions,omitempty"`
	UseHostedPaymentPage *bool `json:"useHostedPaymentPage,omitempty"`
}

// PaymentRequestRemoveMatch is the typed request payload for PaymentRequest.RemoveTyped.
type PaymentRequestRemoveMatch struct {
	Id string `json:"id"`
}

// PaymentRequestEvent is the typed data model for the payment_request_event entity.
type PaymentRequestEvent struct {
	Amount float64 `json:"amount"`
	ApplePayTransactionID *string `json:"applePayTransactionID,omitempty"`
	CardAuthorizationResponseID *string `json:"cardAuthorizationResponseID,omitempty"`
	CardExpiryMonth *int `json:"cardExpiryMonth,omitempty"`
	CardExpiryYear *int `json:"cardExpiryYear,omitempty"`
	CardIssuer *string `json:"cardIssuer,omitempty"`
	CardIssuerCountry *string `json:"cardIssuerCountry,omitempty"`
	CardLastFourDigits *string `json:"cardLastFourDigits,omitempty"`
	CardRequestID *string `json:"cardRequestID,omitempty"`
	CardScheme *string `json:"cardScheme,omitempty"`
	CardTokenCustomerID *string `json:"cardTokenCustomerID,omitempty"`
	CardTransactionID *string `json:"cardTransactionID,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DirectDebitPaymentID *string `json:"directDebitPaymentID,omitempty"`
	DirectDebitPaymentReference *string `json:"directDebitPaymentReference,omitempty"`
	DrirectDebitMandateID *string `json:"drirectDebitMandateID,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorReason *string `json:"errorReason,omitempty"`
	EventType *string `json:"eventType,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	LightningInvoice *string `json:"lightningInvoice,omitempty"`
	LightningRHash *string `json:"lightningRHash,omitempty"`
	OriginUrl *string `json:"originUrl,omitempty"`
	PaymentMethodType *string `json:"paymentMethodType,omitempty"`
	PaymentProcessorName *string `json:"paymentProcessorName,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	PispBankStatus *string `json:"pispBankStatus,omitempty"`
	PispPaymentInitiationID *string `json:"pispPaymentInitiationID,omitempty"`
	PispPaymentInstitutionName *string `json:"pispPaymentInstitutionName,omitempty"`
	PispPaymentServiceProviderID *string `json:"pispPaymentServiceProviderID,omitempty"`
	PispRedirectUrl *string `json:"pispRedirectUrl,omitempty"`
	ReconciledTransactionID *string `json:"reconciledTransactionID,omitempty"`
	RefundPayoutID *string `json:"refundPayoutID,omitempty"`
	Status *string `json:"status,omitempty"`
	WalletName *string `json:"walletName,omitempty"`
}

// PaymentRequestEventListMatch is the typed request payload for PaymentRequestEvent.ListTyped.
type PaymentRequestEventListMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// PaymentRequestMetric is the typed data model for the payment_request_metric entity.
type PaymentRequestMetric struct {
}

// PaymentRequestMetricLoadMatch is the typed request payload for PaymentRequestMetric.LoadTyped.
type PaymentRequestMetricLoadMatch struct {
}

// PaymentRequestMinimal is the typed data model for the payment_request_minimal entity.
type PaymentRequestMinimal struct {
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	CardStripePaymentIntentSecret *string `json:"cardStripePaymentIntentSecret,omitempty"`
	CountryCode *string `json:"countryCode,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomFieldsToDisplay *[]any `json:"customFieldsToDisplay,omitempty"`
	Description *string `json:"description,omitempty"`
	DueDate *string `json:"dueDate,omitempty"`
	FieldDisplaySettings *[]any `json:"fieldDisplaySettings,omitempty"`
	GooglePayMerchantID *string `json:"googlePayMerchantID,omitempty"`
	Id *string `json:"id,omitempty"`
	Jwk *string `json:"jwk,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantLogoUrlPng *string `json:"merchantLogoUrlPng,omitempty"`
	MerchantLogoUrlSvg *string `json:"merchantLogoUrlSvg,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	MerchantShortName *string `json:"merchantShortName,omitempty"`
	PartialPaymentMethod *string `json:"partialPaymentMethod,omitempty"`
	PaymentAttempts *[]any `json:"paymentAttempts,omitempty"`
	PaymentMethodsList *[]any `json:"paymentMethodsList,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentProcessorKey *string `json:"paymentProcessorKey,omitempty"`
	PispError *string `json:"pispError,omitempty"`
	PriorityBankID *string `json:"priorityBankID,omitempty"`
	Status *string `json:"status,omitempty"`
	StripeAccountID *string `json:"stripeAccountID,omitempty"`
	Title *string `json:"title,omitempty"`
}

// PaymentRequestMinimalListMatch is the typed request payload for PaymentRequestMinimal.ListTyped.
type PaymentRequestMinimalListMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// PaymentRequestResult is the typed data model for the payment_request_result entity.
type PaymentRequestResult struct {
	Amount *float64 `json:"amount,omitempty"`
	AmountPending *float64 `json:"amountPending,omitempty"`
	AmountReceived *float64 `json:"amountReceived,omitempty"`
	AmountRefunded *float64 `json:"amountRefunded,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CustomerID *string `json:"customerID,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	Payments *[]any `json:"payments,omitempty"`
	PispAuthorizations *[]any `json:"pispAuthorizations,omitempty"`
	RequestedAmount *float64 `json:"requestedAmount,omitempty"`
	Result *string `json:"result,omitempty"`
}

// PaymentRequestResultListMatch is the typed request payload for PaymentRequestResult.ListTyped.
type PaymentRequestResultListMatch struct {
	PaymentrequestId string `json:"paymentrequest_id"`
}

// Payout is the typed data model for the payout entity.
type Payout struct {
	AccountID *string `json:"accountID,omitempty"`
	AllowIncomplete *bool `json:"allowIncomplete,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnits *int `json:"amountMinorUnits,omitempty"`
	ApprovePayoutUrl *string `json:"approvePayoutUrl,omitempty"`
	ApproverID *string `json:"approverID,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	BatchPayoutID *string `json:"batchPayoutID,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	BeneficiaryID *string `json:"beneficiaryID,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanProcess *bool `json:"canProcess,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	ChargeBearer *string `json:"chargeBearer,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedByEmailAddress *string `json:"createdByEmailAddress,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserID *string `json:"currentUserID,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Documents *[]any `json:"documents,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailedPayouts *map[string]any `json:"failedPayouts,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formattedFxDestinationAmount,omitempty"`
	FormattedSchedule *string `json:"formattedSchedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formattedScheduleDayOnly,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formattedSourceAccountAvailableBalance,omitempty"`
	FxDestinationAmount *float64 `json:"fxDestinationAmount,omitempty"`
	FxDestinationAmountMinorUnits *int `json:"fxDestinationAmountMinorUnits,omitempty"`
	FxDestinationCurrency *string `json:"fxDestinationCurrency,omitempty"`
	FxQuoteExpiresAt *string `json:"fxQuoteExpiresAt,omitempty"`
	FxQuoteID *string `json:"fxQuoteID,omitempty"`
	FxRate *float64 `json:"fxRate,omitempty"`
	FxUseDestinationAmount *bool `json:"fxUseDestinationAmount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceID *string `json:"invoiceID,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsFailed *bool `json:"isFailed,omitempty"`
	IsSettled *bool `json:"isSettled,omitempty"`
	IsSubmitted *bool `json:"isSubmitted,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRail *string `json:"paymentRail,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PayrunName *string `json:"payrunName,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"scheduleDate,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"sourceAccountAvailableBalance,omitempty"`
	SourceAccountAvailableBalanceMinorUnits *int `json:"sourceAccountAvailableBalanceMinorUnits,omitempty"`
	SourceAccountBic *string `json:"sourceAccountBic,omitempty"`
	SourceAccountCurrency *string `json:"sourceAccountCurrency,omitempty"`
	SourceAccountIban *string `json:"sourceAccountIban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"sourceAccountIdentifier"`
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	SourceAccountNumber *string `json:"sourceAccountNumber,omitempty"`
	SourceAccountSortcode *string `json:"sourceAccountSortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	TagIds *[]any `json:"tagIds,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
	TopupPayrunID *string `json:"topupPayrunID,omitempty"`
	TransactedAmount *float64 `json:"transactedAmount,omitempty"`
	TransactedFxAmount *float64 `json:"transactedFxAmount,omitempty"`
	TransactedFxRate *float64 `json:"transactedFxRate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserID *string `json:"userID,omitempty"`
	YourReference *string `json:"yourReference,omitempty"`
}

// PayoutLoadMatch is the typed request payload for Payout.LoadTyped.
type PayoutLoadMatch struct {
	Id string `json:"id"`
}

// PayoutListMatch is the typed request payload for Payout.ListTyped.
type PayoutListMatch struct {
	AccountID *string `json:"accountID,omitempty"`
	AllowIncomplete *bool `json:"allowIncomplete,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnits *int `json:"amountMinorUnits,omitempty"`
	ApprovePayoutUrl *string `json:"approvePayoutUrl,omitempty"`
	ApproverID *string `json:"approverID,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	BatchPayoutID *string `json:"batchPayoutID,omitempty"`
	Beneficiary *map[string]any `json:"beneficiary,omitempty"`
	BeneficiaryID *string `json:"beneficiaryID,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanProcess *bool `json:"canProcess,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	ChargeBearer *string `json:"chargeBearer,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedByEmailAddress *string `json:"createdByEmailAddress,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserID *string `json:"currentUserID,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Documents *[]any `json:"documents,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailedPayouts *map[string]any `json:"failedPayouts,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formattedFxDestinationAmount,omitempty"`
	FormattedSchedule *string `json:"formattedSchedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formattedScheduleDayOnly,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formattedSourceAccountAvailableBalance,omitempty"`
	FxDestinationAmount *float64 `json:"fxDestinationAmount,omitempty"`
	FxDestinationAmountMinorUnits *int `json:"fxDestinationAmountMinorUnits,omitempty"`
	FxDestinationCurrency *string `json:"fxDestinationCurrency,omitempty"`
	FxQuoteExpiresAt *string `json:"fxQuoteExpiresAt,omitempty"`
	FxQuoteID *string `json:"fxQuoteID,omitempty"`
	FxRate *float64 `json:"fxRate,omitempty"`
	FxUseDestinationAmount *bool `json:"fxUseDestinationAmount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceID *string `json:"invoiceID,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsFailed *bool `json:"isFailed,omitempty"`
	IsSettled *bool `json:"isSettled,omitempty"`
	IsSubmitted *bool `json:"isSubmitted,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRail *string `json:"paymentRail,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PayrunName *string `json:"payrunName,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"scheduleDate,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"sourceAccountAvailableBalance,omitempty"`
	SourceAccountAvailableBalanceMinorUnits *int `json:"sourceAccountAvailableBalanceMinorUnits,omitempty"`
	SourceAccountBic *string `json:"sourceAccountBic,omitempty"`
	SourceAccountCurrency *string `json:"sourceAccountCurrency,omitempty"`
	SourceAccountIban *string `json:"sourceAccountIban,omitempty"`
	SourceAccountIdentifier *map[string]any `json:"sourceAccountIdentifier,omitempty"`
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	SourceAccountNumber *string `json:"sourceAccountNumber,omitempty"`
	SourceAccountSortcode *string `json:"sourceAccountSortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	TagIds *[]any `json:"tagIds,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
	TopupPayrunID *string `json:"topupPayrunID,omitempty"`
	TransactedAmount *float64 `json:"transactedAmount,omitempty"`
	TransactedFxAmount *float64 `json:"transactedFxAmount,omitempty"`
	TransactedFxRate *float64 `json:"transactedFxRate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserID *string `json:"userID,omitempty"`
	YourReference *string `json:"yourReference,omitempty"`
}

// PayoutCreateData is the typed request payload for Payout.CreateTyped.
type PayoutCreateData struct {
	Id string `json:"id"`
	AccountID *string `json:"accountID,omitempty"`
	AllowIncomplete *bool `json:"allowIncomplete,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnits *int `json:"amountMinorUnits,omitempty"`
	ApprovePayoutUrl *string `json:"approvePayoutUrl,omitempty"`
	ApproverID *string `json:"approverID,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	BatchPayoutID *string `json:"batchPayoutID,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	BeneficiaryID *string `json:"beneficiaryID,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanProcess *bool `json:"canProcess,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	ChargeBearer *string `json:"chargeBearer,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedByEmailAddress *string `json:"createdByEmailAddress,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserID *string `json:"currentUserID,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Documents *[]any `json:"documents,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailedPayouts *map[string]any `json:"failedPayouts,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formattedFxDestinationAmount,omitempty"`
	FormattedSchedule *string `json:"formattedSchedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formattedScheduleDayOnly,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formattedSourceAccountAvailableBalance,omitempty"`
	FxDestinationAmount *float64 `json:"fxDestinationAmount,omitempty"`
	FxDestinationAmountMinorUnits *int `json:"fxDestinationAmountMinorUnits,omitempty"`
	FxDestinationCurrency *string `json:"fxDestinationCurrency,omitempty"`
	FxQuoteExpiresAt *string `json:"fxQuoteExpiresAt,omitempty"`
	FxQuoteID *string `json:"fxQuoteID,omitempty"`
	FxRate *float64 `json:"fxRate,omitempty"`
	FxUseDestinationAmount *bool `json:"fxUseDestinationAmount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceID *string `json:"invoiceID,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsFailed *bool `json:"isFailed,omitempty"`
	IsSettled *bool `json:"isSettled,omitempty"`
	IsSubmitted *bool `json:"isSubmitted,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRail *string `json:"paymentRail,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PayrunName *string `json:"payrunName,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"scheduleDate,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"sourceAccountAvailableBalance,omitempty"`
	SourceAccountAvailableBalanceMinorUnits *int `json:"sourceAccountAvailableBalanceMinorUnits,omitempty"`
	SourceAccountBic *string `json:"sourceAccountBic,omitempty"`
	SourceAccountCurrency *string `json:"sourceAccountCurrency,omitempty"`
	SourceAccountIban *string `json:"sourceAccountIban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"sourceAccountIdentifier"`
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	SourceAccountNumber *string `json:"sourceAccountNumber,omitempty"`
	SourceAccountSortcode *string `json:"sourceAccountSortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	TagIds *[]any `json:"tagIds,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
	TopupPayrunID *string `json:"topupPayrunID,omitempty"`
	TransactedAmount *float64 `json:"transactedAmount,omitempty"`
	TransactedFxAmount *float64 `json:"transactedFxAmount,omitempty"`
	TransactedFxRate *float64 `json:"transactedFxRate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserID *string `json:"userID,omitempty"`
	YourReference *string `json:"yourReference,omitempty"`
}

// PayoutUpdateData is the typed request payload for Payout.UpdateTyped.
type PayoutUpdateData struct {
	Id string `json:"id"`
	AccountID *string `json:"accountID,omitempty"`
	AllowIncomplete *bool `json:"allowIncomplete,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnits *int `json:"amountMinorUnits,omitempty"`
	ApprovePayoutUrl *string `json:"approvePayoutUrl,omitempty"`
	ApproverID *string `json:"approverID,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	BatchPayoutID *string `json:"batchPayoutID,omitempty"`
	Beneficiary *map[string]any `json:"beneficiary,omitempty"`
	BeneficiaryID *string `json:"beneficiaryID,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanProcess *bool `json:"canProcess,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	ChargeBearer *string `json:"chargeBearer,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedByEmailAddress *string `json:"createdByEmailAddress,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserID *string `json:"currentUserID,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Documents *[]any `json:"documents,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FailedPayouts *map[string]any `json:"failedPayouts,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formattedFxDestinationAmount,omitempty"`
	FormattedSchedule *string `json:"formattedSchedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formattedScheduleDayOnly,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formattedSourceAccountAvailableBalance,omitempty"`
	FxDestinationAmount *float64 `json:"fxDestinationAmount,omitempty"`
	FxDestinationAmountMinorUnits *int `json:"fxDestinationAmountMinorUnits,omitempty"`
	FxDestinationCurrency *string `json:"fxDestinationCurrency,omitempty"`
	FxQuoteExpiresAt *string `json:"fxQuoteExpiresAt,omitempty"`
	FxQuoteID *string `json:"fxQuoteID,omitempty"`
	FxRate *float64 `json:"fxRate,omitempty"`
	FxUseDestinationAmount *bool `json:"fxUseDestinationAmount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceID *string `json:"invoiceID,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsFailed *bool `json:"isFailed,omitempty"`
	IsSettled *bool `json:"isSettled,omitempty"`
	IsSubmitted *bool `json:"isSubmitted,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRail *string `json:"paymentRail,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PayrunName *string `json:"payrunName,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"scheduleDate,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"sourceAccountAvailableBalance,omitempty"`
	SourceAccountAvailableBalanceMinorUnits *int `json:"sourceAccountAvailableBalanceMinorUnits,omitempty"`
	SourceAccountBic *string `json:"sourceAccountBic,omitempty"`
	SourceAccountCurrency *string `json:"sourceAccountCurrency,omitempty"`
	SourceAccountIban *string `json:"sourceAccountIban,omitempty"`
	SourceAccountIdentifier *map[string]any `json:"sourceAccountIdentifier,omitempty"`
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	SourceAccountNumber *string `json:"sourceAccountNumber,omitempty"`
	SourceAccountSortcode *string `json:"sourceAccountSortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	TagIds *[]any `json:"tagIds,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
	TopupPayrunID *string `json:"topupPayrunID,omitempty"`
	TransactedAmount *float64 `json:"transactedAmount,omitempty"`
	TransactedFxAmount *float64 `json:"transactedFxAmount,omitempty"`
	TransactedFxRate *float64 `json:"transactedFxRate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserID *string `json:"userID,omitempty"`
	YourReference *string `json:"yourReference,omitempty"`
}

// PayoutRemoveMatch is the typed request payload for Payout.RemoveTyped.
type PayoutRemoveMatch struct {
	Id string `json:"id"`
}

// PayoutKeysetPage is the typed data model for the payout_keyset_page entity.
type PayoutKeysetPage struct {
	AccountID *string `json:"accountID,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnits *int `json:"amountMinorUnits,omitempty"`
	ApprovePayoutUrl *string `json:"approvePayoutUrl,omitempty"`
	ApproverID *string `json:"approverID,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	BatchPayoutID *string `json:"batchPayoutID,omitempty"`
	Beneficiary map[string]any `json:"beneficiary"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanProcess *bool `json:"canProcess,omitempty"`
	CanUpdate *bool `json:"canUpdate,omitempty"`
	ChargeBearer *string `json:"chargeBearer,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedByEmailAddress *string `json:"createdByEmailAddress,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrentUserID *string `json:"currentUserID,omitempty"`
	Description *string `json:"description,omitempty"`
	Destination *map[string]any `json:"destination,omitempty"`
	Documents *[]any `json:"documents,omitempty"`
	Events *[]any `json:"events,omitempty"`
	FormattedAmount *string `json:"formattedAmount,omitempty"`
	FormattedFxDestinationAmount *string `json:"formattedFxDestinationAmount,omitempty"`
	FormattedSchedule *string `json:"formattedSchedule,omitempty"`
	FormattedScheduleDayOnly *string `json:"formattedScheduleDayOnly,omitempty"`
	FormattedSourceAccountAvailableBalance *string `json:"formattedSourceAccountAvailableBalance,omitempty"`
	FxDestinationAmount *float64 `json:"fxDestinationAmount,omitempty"`
	FxDestinationAmountMinorUnits *int `json:"fxDestinationAmountMinorUnits,omitempty"`
	FxDestinationCurrency *string `json:"fxDestinationCurrency,omitempty"`
	FxQuoteExpiresAt *string `json:"fxQuoteExpiresAt,omitempty"`
	FxQuoteID *string `json:"fxQuoteID,omitempty"`
	FxRate *float64 `json:"fxRate,omitempty"`
	FxUseDestinationAmount *bool `json:"fxUseDestinationAmount,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	InvoiceID *string `json:"invoiceID,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsFailed *bool `json:"isFailed,omitempty"`
	IsSettled *bool `json:"isSettled,omitempty"`
	IsSubmitted *bool `json:"isSubmitted,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantTokenDescription *string `json:"merchantTokenDescription,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentRail *string `json:"paymentRail,omitempty"`
	PayrunID *string `json:"payrunID,omitempty"`
	PayrunName *string `json:"payrunName,omitempty"`
	Rule *map[string]any `json:"rule,omitempty"`
	ScheduleDate *string `json:"scheduleDate,omitempty"`
	Scheduled *bool `json:"scheduled,omitempty"`
	SourceAccountAvailableBalance *float64 `json:"sourceAccountAvailableBalance,omitempty"`
	SourceAccountAvailableBalanceMinorUnits *int `json:"sourceAccountAvailableBalanceMinorUnits,omitempty"`
	SourceAccountBic *string `json:"sourceAccountBic,omitempty"`
	SourceAccountCurrency *string `json:"sourceAccountCurrency,omitempty"`
	SourceAccountIban *string `json:"sourceAccountIban,omitempty"`
	SourceAccountIdentifier map[string]any `json:"sourceAccountIdentifier"`
	SourceAccountName *string `json:"sourceAccountName,omitempty"`
	SourceAccountNumber *string `json:"sourceAccountNumber,omitempty"`
	SourceAccountSortcode *string `json:"sourceAccountSortcode,omitempty"`
	Status *string `json:"status,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
	TopupPayrunID *string `json:"topupPayrunID,omitempty"`
	TransactedAmount *float64 `json:"transactedAmount,omitempty"`
	TransactedFxAmount *float64 `json:"transactedFxAmount,omitempty"`
	TransactedFxRate *float64 `json:"transactedFxRate,omitempty"`
	Type *string `json:"type,omitempty"`
	UserID *string `json:"userID,omitempty"`
	YourReference *string `json:"yourReference,omitempty"`
}

// PayoutKeysetPageListMatch is the typed request payload for PayoutKeysetPage.ListTyped.
type PayoutKeysetPageListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// PayoutMetric is the typed data model for the payout_metric entity.
type PayoutMetric struct {
}

// PayoutMetricLoadMatch is the typed request payload for PayoutMetric.LoadTyped.
type PayoutMetricLoadMatch struct {
}

// Payrun is the typed data model for the payrun entity.
type Payrun struct {
	AuthorisationDate *string `json:"authorisationDate,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	BatchPayoutID *string `json:"batchPayoutID,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanEdit *bool `json:"canEdit,omitempty"`
	Events *[]any `json:"events,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	Invoices *[]any `json:"invoices,omitempty"`
	InvoicesMinimal *[]any `json:"invoicesMinimal,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LastUpdatedBy map[string]any `json:"lastUpdatedBy"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Payments *[]any `json:"payments,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
	PayoutsCount *int `json:"payoutsCount,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ScheduleDate *string `json:"scheduleDate,omitempty"`
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	SourceAccounts *[]any `json:"sourceAccounts,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalEur *float64 `json:"totalEur,omitempty"`
	TotalGbp *float64 `json:"totalGbp,omitempty"`
	TotalUsd *float64 `json:"totalUsd,omitempty"`
}

// PayrunLoadMatch is the typed request payload for Payrun.LoadTyped.
type PayrunLoadMatch struct {
	Id string `json:"id"`
}

// PayrunListMatch is the typed request payload for Payrun.ListTyped.
type PayrunListMatch struct {
	AuthorisationDate *string `json:"authorisationDate,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	BatchPayoutID *string `json:"batchPayoutID,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanEdit *bool `json:"canEdit,omitempty"`
	Events *[]any `json:"events,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	Invoices *[]any `json:"invoices,omitempty"`
	InvoicesMinimal *[]any `json:"invoicesMinimal,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LastUpdatedBy *map[string]any `json:"lastUpdatedBy,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Payments *[]any `json:"payments,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
	PayoutsCount *int `json:"payoutsCount,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ScheduleDate *string `json:"scheduleDate,omitempty"`
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	SourceAccounts *[]any `json:"sourceAccounts,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalEur *float64 `json:"totalEur,omitempty"`
	TotalGbp *float64 `json:"totalGbp,omitempty"`
	TotalUsd *float64 `json:"totalUsd,omitempty"`
}

// PayrunCreateData is the typed request payload for Payrun.CreateTyped.
type PayrunCreateData struct {
	Id string `json:"id"`
	AuthorisationDate *string `json:"authorisationDate,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	BatchPayoutID *string `json:"batchPayoutID,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanEdit *bool `json:"canEdit,omitempty"`
	Events *[]any `json:"events,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	Invoices *[]any `json:"invoices,omitempty"`
	InvoicesMinimal *[]any `json:"invoicesMinimal,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LastUpdatedBy map[string]any `json:"lastUpdatedBy"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Payments *[]any `json:"payments,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
	PayoutsCount *int `json:"payoutsCount,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ScheduleDate *string `json:"scheduleDate,omitempty"`
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	SourceAccounts *[]any `json:"sourceAccounts,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalEur *float64 `json:"totalEur,omitempty"`
	TotalGbp *float64 `json:"totalGbp,omitempty"`
	TotalUsd *float64 `json:"totalUsd,omitempty"`
}

// PayrunUpdateData is the typed request payload for Payrun.UpdateTyped.
type PayrunUpdateData struct {
	Id string `json:"id"`
	AuthorisationDate *string `json:"authorisationDate,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	BatchPayoutID *string `json:"batchPayoutID,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CanDelete *bool `json:"canDelete,omitempty"`
	CanEdit *bool `json:"canEdit,omitempty"`
	Events *[]any `json:"events,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	Invoices *[]any `json:"invoices,omitempty"`
	InvoicesMinimal *[]any `json:"invoicesMinimal,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	LastUpdatedBy *map[string]any `json:"lastUpdatedBy,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Payments *[]any `json:"payments,omitempty"`
	Payouts *[]any `json:"payouts,omitempty"`
	PayoutsCount *int `json:"payoutsCount,omitempty"`
	Reason *string `json:"reason,omitempty"`
	ScheduleDate *string `json:"scheduleDate,omitempty"`
	ScheduledDate *string `json:"scheduledDate,omitempty"`
	SourceAccounts *[]any `json:"sourceAccounts,omitempty"`
	Status *string `json:"status,omitempty"`
	TotalEur *float64 `json:"totalEur,omitempty"`
	TotalGbp *float64 `json:"totalGbp,omitempty"`
	TotalUsd *float64 `json:"totalUsd,omitempty"`
}

// PayrunRemoveMatch is the typed request payload for Payrun.RemoveTyped.
type PayrunRemoveMatch struct {
	Id string `json:"id"`
}

// Report is the typed data model for the report entity.
type Report struct {
}

// ReportUpdateData is the typed request payload for Report.UpdateTyped.
type ReportUpdateData struct {
	Id string `json:"id"`
}

// ReportResult is the typed data model for the report_result entity.
type ReportResult struct {
	ContentType *string `json:"contentType,omitempty"`
	Contents *string `json:"contents,omitempty"`
	LastCompletedAt *string `json:"lastCompletedAt,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	ReportName *string `json:"reportName,omitempty"`
	ReportType *string `json:"reportType,omitempty"`
	StatementNumber *int `json:"statementNumber,omitempty"`
}

// ReportResultLoadMatch is the typed request payload for ReportResult.LoadTyped.
type ReportResultLoadMatch struct {
	Id int `json:"id"`
	ReportId string `json:"report_id"`
}

// Role is the typed data model for the role entity.
type Role struct {
	FailedRoles *map[string]any `json:"failedRoles,omitempty"`
	Roles *[]any `json:"roles,omitempty"`
}

// RoleCreateData is the typed request payload for Role.CreateTyped.
type RoleCreateData struct {
	MerchantId string `json:"merchant_id"`
	FailedRoles *map[string]any `json:"failedRoles,omitempty"`
	Roles *[]any `json:"roles,omitempty"`
}

// Rule is the typed data model for the rule entity.
type Rule struct {
	Account *map[string]any `json:"account,omitempty"`
	AccountID *string `json:"accountID,omitempty"`
	ApproveUrl *string `json:"approveUrl,omitempty"`
	ApproverID *string `json:"approverID,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CreatedBy map[string]any `json:"createdBy"`
	Description *string `json:"description,omitempty"`
	EndAt *string `json:"endAt,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsDisabled *bool `json:"isDisabled,omitempty"`
	LastExecutedAt *string `json:"lastExecutedAt,omitempty"`
	LastRunAtTransactionDate *string `json:"lastRunAtTransactionDate,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce string `json:"nonce"`
	OnApprovedWebHookUrl *string `json:"onApprovedWebHookUrl,omitempty"`
	OnExecutionErrorWebHookUrl *string `json:"onExecutionErrorWebHookUrl,omitempty"`
	OnExecutionSuccessWebHookUrl *string `json:"onExecutionSuccessWebHookUrl,omitempty"`
	StartAt *string `json:"startAt,omitempty"`
	Status *string `json:"status,omitempty"`
	SweepAction *map[string]any `json:"sweepAction,omitempty"`
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	TriggerCronExpression *string `json:"triggerCronExpression,omitempty"`
	TriggerOnPayIn *bool `json:"triggerOnPayIn,omitempty"`
	UserID *string `json:"userID,omitempty"`
	WebHookSecret *string `json:"webHookSecret,omitempty"`
}

// RuleLoadMatch is the typed request payload for Rule.LoadTyped.
type RuleLoadMatch struct {
	Id string `json:"id"`
}

// RuleListMatch is the typed request payload for Rule.ListTyped.
type RuleListMatch struct {
	Account *map[string]any `json:"account,omitempty"`
	AccountID *string `json:"accountID,omitempty"`
	ApproveUrl *string `json:"approveUrl,omitempty"`
	ApproverID *string `json:"approverID,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CreatedBy *map[string]any `json:"createdBy,omitempty"`
	Description *string `json:"description,omitempty"`
	EndAt *string `json:"endAt,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsDisabled *bool `json:"isDisabled,omitempty"`
	LastExecutedAt *string `json:"lastExecutedAt,omitempty"`
	LastRunAtTransactionDate *string `json:"lastRunAtTransactionDate,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	OnApprovedWebHookUrl *string `json:"onApprovedWebHookUrl,omitempty"`
	OnExecutionErrorWebHookUrl *string `json:"onExecutionErrorWebHookUrl,omitempty"`
	OnExecutionSuccessWebHookUrl *string `json:"onExecutionSuccessWebHookUrl,omitempty"`
	StartAt *string `json:"startAt,omitempty"`
	Status *string `json:"status,omitempty"`
	SweepAction *map[string]any `json:"sweepAction,omitempty"`
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	TriggerCronExpression *string `json:"triggerCronExpression,omitempty"`
	TriggerOnPayIn *bool `json:"triggerOnPayIn,omitempty"`
	UserID *string `json:"userID,omitempty"`
	WebHookSecret *string `json:"webHookSecret,omitempty"`
}

// RuleCreateData is the typed request payload for Rule.CreateTyped.
type RuleCreateData struct {
	Account *map[string]any `json:"account,omitempty"`
	AccountID *string `json:"accountID,omitempty"`
	ApproveUrl *string `json:"approveUrl,omitempty"`
	ApproverID *string `json:"approverID,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CreatedBy map[string]any `json:"createdBy"`
	Description *string `json:"description,omitempty"`
	EndAt *string `json:"endAt,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsDisabled *bool `json:"isDisabled,omitempty"`
	LastExecutedAt *string `json:"lastExecutedAt,omitempty"`
	LastRunAtTransactionDate *string `json:"lastRunAtTransactionDate,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce string `json:"nonce"`
	OnApprovedWebHookUrl *string `json:"onApprovedWebHookUrl,omitempty"`
	OnExecutionErrorWebHookUrl *string `json:"onExecutionErrorWebHookUrl,omitempty"`
	OnExecutionSuccessWebHookUrl *string `json:"onExecutionSuccessWebHookUrl,omitempty"`
	StartAt *string `json:"startAt,omitempty"`
	Status *string `json:"status,omitempty"`
	SweepAction *map[string]any `json:"sweepAction,omitempty"`
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	TriggerCronExpression *string `json:"triggerCronExpression,omitempty"`
	TriggerOnPayIn *bool `json:"triggerOnPayIn,omitempty"`
	UserID *string `json:"userID,omitempty"`
	WebHookSecret *string `json:"webHookSecret,omitempty"`
}

// RuleUpdateData is the typed request payload for Rule.UpdateTyped.
type RuleUpdateData struct {
	Id string `json:"id"`
	Account *map[string]any `json:"account,omitempty"`
	AccountID *string `json:"accountID,omitempty"`
	ApproveUrl *string `json:"approveUrl,omitempty"`
	ApproverID *string `json:"approverID,omitempty"`
	AuthenticationMethods *[]any `json:"authenticationMethods,omitempty"`
	Authorisations *[]any `json:"authorisations,omitempty"`
	AuthorisersCompletedCount *int `json:"authorisersCompletedCount,omitempty"`
	AuthorisersRequiredCount *int `json:"authorisersRequiredCount,omitempty"`
	CanAuthorise *bool `json:"canAuthorise,omitempty"`
	CreatedBy *map[string]any `json:"createdBy,omitempty"`
	Description *string `json:"description,omitempty"`
	EndAt *string `json:"endAt,omitempty"`
	HasCurrentUserAuthorised *bool `json:"hasCurrentUserAuthorised,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsDisabled *bool `json:"isDisabled,omitempty"`
	LastExecutedAt *string `json:"lastExecutedAt,omitempty"`
	LastRunAtTransactionDate *string `json:"lastRunAtTransactionDate,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	Name *string `json:"name,omitempty"`
	Nonce *string `json:"nonce,omitempty"`
	OnApprovedWebHookUrl *string `json:"onApprovedWebHookUrl,omitempty"`
	OnExecutionErrorWebHookUrl *string `json:"onExecutionErrorWebHookUrl,omitempty"`
	OnExecutionSuccessWebHookUrl *string `json:"onExecutionSuccessWebHookUrl,omitempty"`
	StartAt *string `json:"startAt,omitempty"`
	Status *string `json:"status,omitempty"`
	SweepAction *map[string]any `json:"sweepAction,omitempty"`
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	TriggerCronExpression *string `json:"triggerCronExpression,omitempty"`
	TriggerOnPayIn *bool `json:"triggerOnPayIn,omitempty"`
	UserID *string `json:"userID,omitempty"`
	WebHookSecret *string `json:"webHookSecret,omitempty"`
}

// RuleRemoveMatch is the typed request payload for Rule.RemoveTyped.
type RuleRemoveMatch struct {
	Id string `json:"id"`
}

// RuleEvent is the typed data model for the rule_event entity.
type RuleEvent struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsAuthoriseToEnable *bool `json:"isAuthoriseToEnable,omitempty"`
	Message *string `json:"message,omitempty"`
	RawResponse *string `json:"rawResponse,omitempty"`
	RuleEventType *string `json:"ruleEventType,omitempty"`
	RuleID *string `json:"ruleID,omitempty"`
	User map[string]any `json:"user"`
}

// RuleEventListMatch is the typed request payload for RuleEvent.ListTyped.
type RuleEventListMatch struct {
	Id string `json:"id"`
}

// Tag is the typed data model for the tag entity.
type Tag struct {
	ColourHex *string `json:"colourHex,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	MerchantID string `json:"merchantID"`
	Name string `json:"name"`
}

// TagListMatch is the typed request payload for Tag.ListTyped.
type TagListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// TagCreateData is the typed request payload for Tag.CreateTyped.
type TagCreateData struct {
	MerchantId string `json:"merchant_id"`
	ColourHex *string `json:"colourHex,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	MerchantID string `json:"merchantID"`
	Name string `json:"name"`
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
	AccountID *string `json:"accountID,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	AccountSequenceNumber *int `json:"accountSequenceNumber,omitempty"`
	AddressDetails *map[string]any `json:"addressDetails,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnits *int `json:"amountMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BookingDateTime *string `json:"bookingDateTime,omitempty"`
	ChargeDetails *map[string]any `json:"chargeDetails,omitempty"`
	Content *[]any `json:"content,omitempty"`
	Counterparty *map[string]any `json:"counterparty,omitempty"`
	CounterpartySummary *string `json:"counterpartySummary,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrencyExchange *map[string]any `json:"currencyExchange,omitempty"`
	Date *string `json:"date,omitempty"`
	Description *string `json:"description,omitempty"`
	Enrichment *map[string]any `json:"enrichment,omitempty"`
	FxAmount *float64 `json:"fxAmount,omitempty"`
	FxCurrency *string `json:"fxCurrency,omitempty"`
	FxRate *float64 `json:"fxRate,omitempty"`
	GrossAmount map[string]any `json:"grossAmount"`
	Id *string `json:"id,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsoBankTransactionCode *map[string]any `json:"isoBankTransactionCode,omitempty"`
	Merchant *map[string]any `json:"merchant,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	PageNumber *int `json:"pageNumber,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
	PayeeDetails map[string]any `json:"payeeDetails"`
	PayerDetails map[string]any `json:"payerDetails"`
	PaymentRequestCustomFields *map[string]any `json:"paymentRequestCustomFields,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	PayoutID *string `json:"payoutID,omitempty"`
	ProprietaryBankTransactionCode *map[string]any `json:"proprietaryBankTransactionCode,omitempty"`
	RawReference *string `json:"rawReference,omitempty"`
	Reference *string `json:"reference,omitempty"`
	RuleID *string `json:"ruleID,omitempty"`
	StatementReferences *[]any `json:"statementReferences,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplementaryData *any `json:"supplementaryData,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
	TotalPages *int `json:"totalPages,omitempty"`
	TotalSize *int `json:"totalSize,omitempty"`
	TransactionAmount map[string]any `json:"transactionAmount"`
	TransactionDate *string `json:"transactionDate,omitempty"`
	TransactionInformation *[]any `json:"transactionInformation,omitempty"`
	TransactionMutability *string `json:"transactionMutability,omitempty"`
	Type *string `json:"type,omitempty"`
	ValueDateTime *string `json:"valueDateTime,omitempty"`
	VirtualIBAN *string `json:"virtualIBAN,omitempty"`
	YourReference *string `json:"yourReference,omitempty"`
}

// TransactionLoadMatch is the typed request payload for Transaction.LoadTyped.
type TransactionLoadMatch struct {
	Id string `json:"id"`
}

// TransactionListMatch is the typed request payload for Transaction.ListTyped.
type TransactionListMatch struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
}

// TransactionCreateData is the typed request payload for Transaction.CreateTyped.
type TransactionCreateData struct {
	Id string `json:"id"`
	AccountID *string `json:"accountID,omitempty"`
	AccountName *string `json:"accountName,omitempty"`
	AccountSequenceNumber *int `json:"accountSequenceNumber,omitempty"`
	AddressDetails *map[string]any `json:"addressDetails,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	AmountMinorUnits *int `json:"amountMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BookingDateTime *string `json:"bookingDateTime,omitempty"`
	ChargeDetails *map[string]any `json:"chargeDetails,omitempty"`
	Content *[]any `json:"content,omitempty"`
	Counterparty *map[string]any `json:"counterparty,omitempty"`
	CounterpartySummary *string `json:"counterpartySummary,omitempty"`
	Currency *string `json:"currency,omitempty"`
	CurrencyExchange *map[string]any `json:"currencyExchange,omitempty"`
	Date *string `json:"date,omitempty"`
	Description *string `json:"description,omitempty"`
	Enrichment *map[string]any `json:"enrichment,omitempty"`
	FxAmount *float64 `json:"fxAmount,omitempty"`
	FxCurrency *string `json:"fxCurrency,omitempty"`
	FxRate *float64 `json:"fxRate,omitempty"`
	GrossAmount map[string]any `json:"grossAmount"`
	Inserted *string `json:"inserted,omitempty"`
	IsoBankTransactionCode *map[string]any `json:"isoBankTransactionCode,omitempty"`
	Merchant *map[string]any `json:"merchant,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	PageNumber *int `json:"pageNumber,omitempty"`
	PageSize *int `json:"pageSize,omitempty"`
	PayeeDetails map[string]any `json:"payeeDetails"`
	PayerDetails map[string]any `json:"payerDetails"`
	PaymentRequestCustomFields *map[string]any `json:"paymentRequestCustomFields,omitempty"`
	PaymentRequestID *string `json:"paymentRequestID,omitempty"`
	PayoutID *string `json:"payoutID,omitempty"`
	ProprietaryBankTransactionCode *map[string]any `json:"proprietaryBankTransactionCode,omitempty"`
	RawReference *string `json:"rawReference,omitempty"`
	Reference *string `json:"reference,omitempty"`
	RuleID *string `json:"ruleID,omitempty"`
	StatementReferences *[]any `json:"statementReferences,omitempty"`
	Status *string `json:"status,omitempty"`
	SupplementaryData *any `json:"supplementaryData,omitempty"`
	Tags *[]any `json:"tags,omitempty"`
	TheirReference *string `json:"theirReference,omitempty"`
	TotalPages *int `json:"totalPages,omitempty"`
	TotalSize *int `json:"totalSize,omitempty"`
	TransactionAmount map[string]any `json:"transactionAmount"`
	TransactionDate *string `json:"transactionDate,omitempty"`
	TransactionInformation *[]any `json:"transactionInformation,omitempty"`
	TransactionMutability *string `json:"transactionMutability,omitempty"`
	Type *string `json:"type,omitempty"`
	ValueDateTime *string `json:"valueDateTime,omitempty"`
	VirtualIBAN *string `json:"virtualIBAN,omitempty"`
	YourReference *string `json:"yourReference,omitempty"`
}

// TransactionRemoveMatch is the typed request payload for Transaction.RemoveTyped.
type TransactionRemoveMatch struct {
	Id string `json:"id"`
}

// User is the typed data model for the user entity.
type User struct {
	ClientSessionTimeouts *[]any `json:"clientSessionTimeouts,omitempty"`
	EmailAddress string `json:"emailAddress"`
	FirstName string `json:"firstName"`
	Id *string `json:"id,omitempty"`
	LastName string `json:"lastName"`
	PasskeyAdded *bool `json:"passkeyAdded,omitempty"`
	Permissions *map[string]any `json:"permissions,omitempty"`
	Profile *string `json:"profile,omitempty"`
	RolesWithScope *[]any `json:"rolesWithScope,omitempty"`
	TwoFactorEnabled *bool `json:"twoFactorEnabled,omitempty"`
	UserInviteID *string `json:"userInviteID,omitempty"`
}

// UserListMatch is the typed request payload for User.ListTyped.
type UserListMatch struct {
	ClientSessionTimeouts *[]any `json:"clientSessionTimeouts,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	PasskeyAdded *bool `json:"passkeyAdded,omitempty"`
	Permissions *map[string]any `json:"permissions,omitempty"`
	Profile *string `json:"profile,omitempty"`
	RolesWithScope *[]any `json:"rolesWithScope,omitempty"`
	TwoFactorEnabled *bool `json:"twoFactorEnabled,omitempty"`
	UserInviteID *string `json:"userInviteID,omitempty"`
}

// UserUpdateData is the typed request payload for User.UpdateTyped.
type UserUpdateData struct {
	Id string `json:"id"`
	ClientSessionTimeouts *[]any `json:"clientSessionTimeouts,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	PasskeyAdded *bool `json:"passkeyAdded,omitempty"`
	Permissions *map[string]any `json:"permissions,omitempty"`
	Profile *string `json:"profile,omitempty"`
	RolesWithScope *[]any `json:"rolesWithScope,omitempty"`
	TwoFactorEnabled *bool `json:"twoFactorEnabled,omitempty"`
	UserInviteID *string `json:"userInviteID,omitempty"`
}

// UserInvite is the typed data model for the user_invite entity.
type UserInvite struct {
	AuthorisationStatus *map[string]any `json:"authorisationStatus,omitempty"`
	FailedUserInvites *map[string]any `json:"failedUserInvites,omitempty"`
	Id *string `json:"id,omitempty"`
	InitialRoleID *string `json:"initialRoleID,omitempty"`
	InviteeEmailAddress *string `json:"inviteeEmailAddress,omitempty"`
	InviteeFirstName *string `json:"inviteeFirstName,omitempty"`
	InviteeLastName *string `json:"inviteeLastName,omitempty"`
	InviterEmailAddress *string `json:"inviterEmailAddress,omitempty"`
	InviterFirstName *string `json:"inviterFirstName,omitempty"`
	InviterLastName *string `json:"inviterLastName,omitempty"`
	IsAuthorised *bool `json:"isAuthorised,omitempty"`
	IsInviteeRegistered *bool `json:"isInviteeRegistered,omitempty"`
	LastInvited *string `json:"lastInvited,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Message *string `json:"message,omitempty"`
	RegistrationUrl *string `json:"registrationUrl,omitempty"`
	SendInviteEmail *bool `json:"sendInviteEmail,omitempty"`
	Status *string `json:"status,omitempty"`
	User map[string]any `json:"user"`
	UserID *string `json:"userID,omitempty"`
	UserInvites *[]any `json:"userInvites,omitempty"`
}

// UserInviteLoadMatch is the typed request payload for UserInvite.LoadTyped.
type UserInviteLoadMatch struct {
	Id string `json:"id"`
}

// UserInviteListMatch is the typed request payload for UserInvite.ListTyped.
type UserInviteListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// UserInviteCreateData is the typed request payload for UserInvite.CreateTyped.
type UserInviteCreateData struct {
	Id string `json:"id"`
	AuthorisationStatus *map[string]any `json:"authorisationStatus,omitempty"`
	FailedUserInvites *map[string]any `json:"failedUserInvites,omitempty"`
	InitialRoleID *string `json:"initialRoleID,omitempty"`
	InviteeEmailAddress *string `json:"inviteeEmailAddress,omitempty"`
	InviteeFirstName *string `json:"inviteeFirstName,omitempty"`
	InviteeLastName *string `json:"inviteeLastName,omitempty"`
	InviterEmailAddress *string `json:"inviterEmailAddress,omitempty"`
	InviterFirstName *string `json:"inviterFirstName,omitempty"`
	InviterLastName *string `json:"inviterLastName,omitempty"`
	IsAuthorised *bool `json:"isAuthorised,omitempty"`
	IsInviteeRegistered *bool `json:"isInviteeRegistered,omitempty"`
	LastInvited *string `json:"lastInvited,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Message *string `json:"message,omitempty"`
	RegistrationUrl *string `json:"registrationUrl,omitempty"`
	SendInviteEmail *bool `json:"sendInviteEmail,omitempty"`
	Status *string `json:"status,omitempty"`
	User map[string]any `json:"user"`
	UserID *string `json:"userID,omitempty"`
	UserInvites *[]any `json:"userInvites,omitempty"`
}

// UserInviteUpdateData is the typed request payload for UserInvite.UpdateTyped.
type UserInviteUpdateData struct {
	Id string `json:"id"`
	AuthorisationStatus *map[string]any `json:"authorisationStatus,omitempty"`
	FailedUserInvites *map[string]any `json:"failedUserInvites,omitempty"`
	InitialRoleID *string `json:"initialRoleID,omitempty"`
	InviteeEmailAddress *string `json:"inviteeEmailAddress,omitempty"`
	InviteeFirstName *string `json:"inviteeFirstName,omitempty"`
	InviteeLastName *string `json:"inviteeLastName,omitempty"`
	InviterEmailAddress *string `json:"inviterEmailAddress,omitempty"`
	InviterFirstName *string `json:"inviterFirstName,omitempty"`
	InviterLastName *string `json:"inviterLastName,omitempty"`
	IsAuthorised *bool `json:"isAuthorised,omitempty"`
	IsInviteeRegistered *bool `json:"isInviteeRegistered,omitempty"`
	LastInvited *string `json:"lastInvited,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Message *string `json:"message,omitempty"`
	RegistrationUrl *string `json:"registrationUrl,omitempty"`
	SendInviteEmail *bool `json:"sendInviteEmail,omitempty"`
	Status *string `json:"status,omitempty"`
	User *map[string]any `json:"user,omitempty"`
	UserID *string `json:"userID,omitempty"`
	UserInvites *[]any `json:"userInvites,omitempty"`
}

// UserInviteRemoveMatch is the typed request payload for UserInvite.RemoveTyped.
type UserInviteRemoveMatch struct {
	Id string `json:"id"`
}

// Virtual is the typed data model for the virtual entity.
type Virtual struct {
	AccountName *string `json:"accountName,omitempty"`
	AccountSupplierName *string `json:"accountSupplierName,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	AvailableBalanceMinorUnits *int `json:"availableBalanceMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	CreatedBy map[string]any `json:"createdBy"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"defaultPaymentRail,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	ExternalAccountIcon *string `json:"externalAccountIcon,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsTrustAccount *bool `json:"isTrustAccount,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	LastTransaction *map[string]any `json:"lastTransaction,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Name string `json:"name"`
	PhysicalAccountID *string `json:"physicalAccountID,omitempty"`
	Rules *[]any `json:"rules,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
	SubmittedPayoutsBalanceMinorUnits *int `json:"submittedPayoutsBalanceMinorUnits,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplierSepaInstantStatus,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xeroBankFeedConnectionStatus,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xeroBankFeedLastSyncedAt,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xeroBankFeedSyncLastFailedAt,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xeroBankFeedSyncLastFailureReason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xeroBankFeedSyncStatus,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xeroUnsynchronisedTransactionsCount,omitempty"`
}

// VirtualCreateData is the typed request payload for Virtual.CreateTyped.
type VirtualCreateData struct {
	AccountId string `json:"account_id"`
	AccountName *string `json:"accountName,omitempty"`
	AccountSupplierName *string `json:"accountSupplierName,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	AvailableBalanceMinorUnits *int `json:"availableBalanceMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	CreatedBy map[string]any `json:"createdBy"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"defaultPaymentRail,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	ExternalAccountIcon *string `json:"externalAccountIcon,omitempty"`
	Id *string `json:"id,omitempty"`
	Identifier map[string]any `json:"identifier"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsTrustAccount *bool `json:"isTrustAccount,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	LastTransaction *map[string]any `json:"lastTransaction,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Name string `json:"name"`
	PhysicalAccountID *string `json:"physicalAccountID,omitempty"`
	Rules *[]any `json:"rules,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
	SubmittedPayoutsBalanceMinorUnits *int `json:"submittedPayoutsBalanceMinorUnits,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplierSepaInstantStatus,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xeroBankFeedConnectionStatus,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xeroBankFeedLastSyncedAt,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xeroBankFeedSyncLastFailedAt,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xeroBankFeedSyncLastFailureReason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xeroBankFeedSyncStatus,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xeroUnsynchronisedTransactionsCount,omitempty"`
}

// VirtualUpdateData is the typed request payload for Virtual.UpdateTyped.
type VirtualUpdateData struct {
	AccountId string `json:"account_id"`
	Id string `json:"id"`
	AccountName *string `json:"accountName,omitempty"`
	AccountSupplierName *string `json:"accountSupplierName,omitempty"`
	AvailableBalance *float64 `json:"availableBalance,omitempty"`
	AvailableBalanceMinorUnits *int `json:"availableBalanceMinorUnits,omitempty"`
	Balance *float64 `json:"balance,omitempty"`
	BalanceMinorUnits *int `json:"balanceMinorUnits,omitempty"`
	BankName *string `json:"bankName,omitempty"`
	ConsentID *string `json:"consentID,omitempty"`
	CreatedBy *map[string]any `json:"createdBy,omitempty"`
	CreatedByDisplayName *string `json:"createdByDisplayName,omitempty"`
	Currency *string `json:"currency,omitempty"`
	DefaultPaymentRail *string `json:"defaultPaymentRail,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	ExpiryDate *string `json:"expiryDate,omitempty"`
	ExternalAccountIcon *string `json:"externalAccountIcon,omitempty"`
	Identifier *map[string]any `json:"identifier,omitempty"`
	Inserted *string `json:"inserted,omitempty"`
	IsArchived *bool `json:"isArchived,omitempty"`
	IsConnectedAccount *bool `json:"isConnectedAccount,omitempty"`
	IsDefault *bool `json:"isDefault,omitempty"`
	IsTrustAccount *bool `json:"isTrustAccount,omitempty"`
	IsVirtual *bool `json:"isVirtual,omitempty"`
	LastTransaction *map[string]any `json:"lastTransaction,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	MerchantName *string `json:"merchantName,omitempty"`
	Name *string `json:"name,omitempty"`
	PhysicalAccountID *string `json:"physicalAccountID,omitempty"`
	Rules *[]any `json:"rules,omitempty"`
	SubmittedPayoutsBalance *float64 `json:"submittedPayoutsBalance,omitempty"`
	SubmittedPayoutsBalanceMinorUnits *int `json:"submittedPayoutsBalanceMinorUnits,omitempty"`
	Summary *string `json:"summary,omitempty"`
	SupplierSepaInstantStatus *string `json:"supplierSepaInstantStatus,omitempty"`
	XeroBankFeedConnectionStatus *string `json:"xeroBankFeedConnectionStatus,omitempty"`
	XeroBankFeedLastSyncedAt *string `json:"xeroBankFeedLastSyncedAt,omitempty"`
	XeroBankFeedSyncLastFailedAt *string `json:"xeroBankFeedSyncLastFailedAt,omitempty"`
	XeroBankFeedSyncLastFailureReason *string `json:"xeroBankFeedSyncLastFailureReason,omitempty"`
	XeroBankFeedSyncStatus *string `json:"xeroBankFeedSyncStatus,omitempty"`
	XeroUnsynchronisedTransactionsCount *int `json:"xeroUnsynchronisedTransactionsCount,omitempty"`
}

// Webhook is the typed data model for the webhook entity.
type Webhook struct {
	DestinationUrl *string `json:"destinationUrl,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	FailedNotificationEmailAddress *string `json:"failedNotificationEmailAddress,omitempty"`
	Id *string `json:"id,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	NotificationMethod *string `json:"notificationMethod,omitempty"`
	ResourceTypes *[]any `json:"resourceTypes,omitempty"`
	Retry *bool `json:"retry,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Version *int `json:"version,omitempty"`
}

// WebhookLoadMatch is the typed request payload for Webhook.LoadTyped.
type WebhookLoadMatch struct {
	Id string `json:"id"`
	MerchantId *string `json:"merchant_id,omitempty"`
}

// WebhookListMatch is the typed request payload for Webhook.ListTyped.
type WebhookListMatch struct {
	MerchantId string `json:"merchant_id"`
}

// WebhookCreateData is the typed request payload for Webhook.CreateTyped.
type WebhookCreateData struct {
	DestinationUrl *string `json:"destinationUrl,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	FailedNotificationEmailAddress *string `json:"failedNotificationEmailAddress,omitempty"`
	Id *string `json:"id,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	NotificationMethod *string `json:"notificationMethod,omitempty"`
	ResourceTypes *[]any `json:"resourceTypes,omitempty"`
	Retry *bool `json:"retry,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Version *int `json:"version,omitempty"`
}

// WebhookUpdateData is the typed request payload for Webhook.UpdateTyped.
type WebhookUpdateData struct {
	Id string `json:"id"`
	DestinationUrl *string `json:"destinationUrl,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	FailedNotificationEmailAddress *string `json:"failedNotificationEmailAddress,omitempty"`
	IsActive *bool `json:"isActive,omitempty"`
	MerchantID *string `json:"merchantID,omitempty"`
	NotificationMethod *string `json:"notificationMethod,omitempty"`
	ResourceTypes *[]any `json:"resourceTypes,omitempty"`
	Retry *bool `json:"retry,omitempty"`
	Secret *string `json:"secret,omitempty"`
	Version *int `json:"version,omitempty"`
}

// WebhookRemoveMatch is the typed request payload for Webhook.RemoveTyped.
type WebhookRemoveMatch struct {
	Id string `json:"id"`
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

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
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

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
