# Typed models for the Nofrixion SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class AccountRequired(TypedDict):
    createdBy: dict
    identifier: dict


class Account(AccountRequired, total=False):
    accountBalances: list
    accountID: str
    accountIdentifications: list
    accountName: str
    accountNames: list
    accountSupplierName: str
    accountType: str
    availableBalance: float
    availableBalanceMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bankName: str
    consentID: str
    consolidatedAccountInformation: dict
    createdByDisplayName: str
    currency: str
    defaultPaymentRail: str
    description: str
    details: str
    displayName: str
    expiryDate: str
    externalAccountIcon: str
    format: str
    fromDate: str
    id: str
    inserted: str
    isArchived: bool
    isConnectedAccount: bool
    isDefault: bool
    isTrustAccount: bool
    isVirtual: bool
    lastTransaction: dict
    lastUpdated: str
    merchantID: str
    merchantName: str
    nickname: str
    physicalAccountID: str
    roleIDs: list
    rules: list
    submittedPayoutsBalance: float
    submittedPayoutsBalanceMinorUnits: int
    summary: str
    supplierPhysicalAccountID: str
    supplierSepaInstantStatus: str
    toDate: str
    type: str
    usageType: str
    xeroBankFeedConnectionStatus: str
    xeroBankFeedLastSyncedAt: str
    xeroBankFeedSyncLastFailedAt: str
    xeroBankFeedSyncLastFailureReason: str
    xeroBankFeedSyncStatus: str
    xeroUnsynchronisedTransactionsCount: int


class AccountLoadMatchRequired(TypedDict):
    id: str


class AccountLoadMatch(AccountLoadMatchRequired, total=False):
    account_id: str
    merchant_id: str


class AccountListMatch(TypedDict, total=False):
    accountBalances: list
    accountID: str
    accountIdentifications: list
    accountName: str
    accountNames: list
    accountSupplierName: str
    accountType: str
    availableBalance: float
    availableBalanceMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bankName: str
    consentID: str
    consolidatedAccountInformation: dict
    createdBy: dict
    createdByDisplayName: str
    currency: str
    defaultPaymentRail: str
    description: str
    details: str
    displayName: str
    expiryDate: str
    externalAccountIcon: str
    format: str
    fromDate: str
    id: str
    identifier: dict
    inserted: str
    isArchived: bool
    isConnectedAccount: bool
    isDefault: bool
    isTrustAccount: bool
    isVirtual: bool
    lastTransaction: dict
    lastUpdated: str
    merchantID: str
    merchantName: str
    nickname: str
    physicalAccountID: str
    roleIDs: list
    rules: list
    submittedPayoutsBalance: float
    submittedPayoutsBalanceMinorUnits: int
    summary: str
    supplierPhysicalAccountID: str
    supplierSepaInstantStatus: str
    toDate: str
    type: str
    usageType: str
    xeroBankFeedConnectionStatus: str
    xeroBankFeedLastSyncedAt: str
    xeroBankFeedSyncLastFailedAt: str
    xeroBankFeedSyncLastFailureReason: str
    xeroBankFeedSyncStatus: str
    xeroUnsynchronisedTransactionsCount: int


class AccountCreateDataRequired(TypedDict):
    account_id: str
    currency: str
    createdBy: dict
    identifier: dict


class AccountCreateData(AccountCreateDataRequired, total=False):
    accountBalances: list
    accountID: str
    accountIdentifications: list
    accountName: str
    accountNames: list
    accountSupplierName: str
    accountType: str
    availableBalance: float
    availableBalanceMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bankName: str
    consentID: str
    consolidatedAccountInformation: dict
    createdByDisplayName: str
    defaultPaymentRail: str
    description: str
    details: str
    displayName: str
    expiryDate: str
    externalAccountIcon: str
    format: str
    fromDate: str
    id: str
    inserted: str
    isArchived: bool
    isConnectedAccount: bool
    isDefault: bool
    isTrustAccount: bool
    isVirtual: bool
    lastTransaction: dict
    lastUpdated: str
    merchantID: str
    merchantName: str
    nickname: str
    physicalAccountID: str
    roleIDs: list
    rules: list
    submittedPayoutsBalance: float
    submittedPayoutsBalanceMinorUnits: int
    summary: str
    supplierPhysicalAccountID: str
    supplierSepaInstantStatus: str
    toDate: str
    type: str
    usageType: str
    xeroBankFeedConnectionStatus: str
    xeroBankFeedLastSyncedAt: str
    xeroBankFeedSyncLastFailedAt: str
    xeroBankFeedSyncLastFailureReason: str
    xeroBankFeedSyncStatus: str
    xeroUnsynchronisedTransactionsCount: int


class AccountUpdateDataRequired(TypedDict):
    id: str


class AccountUpdateData(AccountUpdateDataRequired, total=False):
    accountBalances: list
    accountID: str
    accountIdentifications: list
    accountName: str
    accountNames: list
    accountSupplierName: str
    accountType: str
    availableBalance: float
    availableBalanceMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bankName: str
    consentID: str
    consolidatedAccountInformation: dict
    createdBy: dict
    createdByDisplayName: str
    currency: str
    defaultPaymentRail: str
    description: str
    details: str
    displayName: str
    expiryDate: str
    externalAccountIcon: str
    format: str
    fromDate: str
    identifier: dict
    inserted: str
    isArchived: bool
    isConnectedAccount: bool
    isDefault: bool
    isTrustAccount: bool
    isVirtual: bool
    lastTransaction: dict
    lastUpdated: str
    merchantID: str
    merchantName: str
    nickname: str
    physicalAccountID: str
    roleIDs: list
    rules: list
    submittedPayoutsBalance: float
    submittedPayoutsBalanceMinorUnits: int
    summary: str
    supplierPhysicalAccountID: str
    supplierSepaInstantStatus: str
    toDate: str
    type: str
    usageType: str
    xeroBankFeedConnectionStatus: str
    xeroBankFeedLastSyncedAt: str
    xeroBankFeedSyncLastFailedAt: str
    xeroBankFeedSyncLastFailureReason: str
    xeroBankFeedSyncStatus: str
    xeroUnsynchronisedTransactionsCount: int


class AccountRemoveMatch(TypedDict):
    id: str


class Batch(TypedDict, total=False):
    approveUrl: str
    id: str
    payouts: list


class BatchLoadMatch(TypedDict):
    id: str


class BatchCreateData(TypedDict, total=False):
    approveUrl: str
    id: str
    payouts: list


class BeneficiaryRequired(TypedDict):
    createdBy: dict
    currency: str
    name: str


class Beneficiary(BeneficiaryRequired, total=False):
    approvalCallbackUrl: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    beneficiaries: list
    beneficiaryEvents: list
    canAuthorise: bool
    canUpdate: bool
    createdByEmailAddress: str
    destination: dict
    failedBeneficiaries: dict
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    isEnabled: bool
    lastAuthorised: str
    lastUpdated: str
    merchantID: str
    nonce: str
    sourceAccountIDs: list
    sourceAccounts: list
    theirReference: str


class BeneficiaryLoadMatchRequired(TypedDict):
    id: str


class BeneficiaryLoadMatch(BeneficiaryLoadMatchRequired, total=False):
    merchant_id: str


class BeneficiaryListMatch(TypedDict, total=False):
    approvalCallbackUrl: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    beneficiaries: list
    beneficiaryEvents: list
    canAuthorise: bool
    canUpdate: bool
    createdBy: dict
    createdByEmailAddress: str
    currency: str
    destination: dict
    failedBeneficiaries: dict
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    isEnabled: bool
    lastAuthorised: str
    lastUpdated: str
    merchantID: str
    name: str
    nonce: str
    sourceAccountIDs: list
    sourceAccounts: list
    theirReference: str


class BeneficiaryCreateDataRequired(TypedDict):
    id: str
    createdBy: dict
    currency: str
    name: str


class BeneficiaryCreateData(BeneficiaryCreateDataRequired, total=False):
    approvalCallbackUrl: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    beneficiaries: list
    beneficiaryEvents: list
    canAuthorise: bool
    canUpdate: bool
    createdByEmailAddress: str
    destination: dict
    failedBeneficiaries: dict
    hasCurrentUserAuthorised: bool
    inserted: str
    isEnabled: bool
    lastAuthorised: str
    lastUpdated: str
    merchantID: str
    nonce: str
    sourceAccountIDs: list
    sourceAccounts: list
    theirReference: str


class BeneficiaryUpdateDataRequired(TypedDict):
    id: str


class BeneficiaryUpdateData(BeneficiaryUpdateDataRequired, total=False):
    approvalCallbackUrl: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    beneficiaries: list
    beneficiaryEvents: list
    canAuthorise: bool
    canUpdate: bool
    createdBy: dict
    createdByEmailAddress: str
    currency: str
    destination: dict
    failedBeneficiaries: dict
    hasCurrentUserAuthorised: bool
    inserted: str
    isEnabled: bool
    lastAuthorised: str
    lastUpdated: str
    merchantID: str
    name: str
    nonce: str
    sourceAccountIDs: list
    sourceAccounts: list
    theirReference: str


class BeneficiaryRemoveMatch(TypedDict):
    id: str


class BeneficiaryGroupRequired(TypedDict):
    groupName: str
    merchantID: str


class BeneficiaryGroup(BeneficiaryGroupRequired, total=False):
    groupMembers: list
    id: str
    inserted: str
    lastUpdated: str


class BeneficiaryGroupListMatch(TypedDict):
    merchant_id: str


class Card(TypedDict, total=False):
    authorizedAmount: str
    currencyCode: str
    isPayerAuthenticationRequired: bool
    isSoftDecline: bool
    payerAuthenticationAccessToken: str
    payerAuthenticationMerchantData: str
    payerAuthenticationUrl: str
    payerAuthenticationWindowHeight: int
    payerAuthenticationWindowWidth: int
    paymentRequestCallbackUrl: str
    paymentRequestID: str
    requestID: str
    responseCode: str
    responseType: str
    status: str
    threeDSRedirectUrl: str
    transactionID: str


class CardCreateDataRequired(TypedDict):
    paymentrequest_id: str


class CardCreateData(CardCreateDataRequired, total=False):
    authorizedAmount: str
    currencyCode: str
    isPayerAuthenticationRequired: bool
    isSoftDecline: bool
    payerAuthenticationAccessToken: str
    payerAuthenticationMerchantData: str
    payerAuthenticationUrl: str
    payerAuthenticationWindowHeight: int
    payerAuthenticationWindowWidth: int
    paymentRequestCallbackUrl: str
    paymentRequestID: str
    requestID: str
    responseCode: str
    responseType: str
    status: str
    threeDSRedirectUrl: str
    transactionID: str


class CardCustomerToken(TypedDict, total=False):
    cardType: str
    customerEmailAddress: str
    expiryMonth: str
    expiryYear: str
    id: str
    inserted: str
    lastFourDigits: str
    lastUpdated: str
    maskedCardNumber: str
    merchantID: str
    paymentRequestID: str


class CardCustomerTokenLoadMatch(TypedDict):
    customer_email_address: str


class CardCustomerTokenListMatch(TypedDict):
    customer_email_address: str
    merchant_id: str


class CardCustomerTokenRemoveMatch(TypedDict):
    id: str


class CardPayment(TypedDict, total=False):
    authorizedAmount: str
    currencyCode: str
    isPayerAuthenticationRequired: bool
    isSoftDecline: bool
    payerAuthenticationAccessToken: str
    payerAuthenticationMerchantData: str
    payerAuthenticationUrl: str
    payerAuthenticationWindowHeight: int
    payerAuthenticationWindowWidth: int
    paymentRequestCallbackUrl: str
    paymentRequestID: str
    requestID: str
    responseCode: str
    responseType: str
    status: str
    threeDSRedirectUrl: str
    transactionID: str


class CardPaymentCreateDataRequired(TypedDict):
    paymentrequest_id: str


class CardPaymentCreateData(CardPaymentCreateDataRequired, total=False):
    partial_refund_amount: float
    authorizedAmount: str
    currencyCode: str
    isPayerAuthenticationRequired: bool
    isSoftDecline: bool
    payerAuthenticationAccessToken: str
    payerAuthenticationMerchantData: str
    payerAuthenticationUrl: str
    payerAuthenticationWindowHeight: int
    payerAuthenticationWindowWidth: int
    paymentRequestCallbackUrl: str
    paymentRequestID: str
    requestID: str
    responseCode: str
    responseType: str
    status: str
    threeDSRedirectUrl: str
    transactionID: str


class CardPublicKey(TypedDict, total=False):
    jwt: str


class CardPublicKeyLoadMatch(TypedDict):
    paymentrequest_id: str


class Consent(TypedDict, total=False):
    authorisationUrl: str
    callbackUrl: str
    consentID: str
    emailAddress: str
    expiryDate: str
    failureCallbackUrl: str
    id: str
    inserted: str
    institutionID: str
    isConnectedAccounts: bool
    isEnabled: bool
    merchantID: str
    provider: str
    successWebHookUrl: str


class ConsentLoadMatch(TypedDict):
    id: str


class ConsentListMatch(TypedDict):
    email: str
    merchant_id: str


class ConsentCreateData(TypedDict, total=False):
    authorisationUrl: str
    callbackUrl: str
    consentID: str
    emailAddress: str
    expiryDate: str
    failureCallbackUrl: str
    id: str
    inserted: str
    institutionID: str
    isConnectedAccounts: bool
    isEnabled: bool
    merchantID: str
    provider: str
    successWebHookUrl: str


class ConsentUpdateDataRequired(TypedDict):
    id: str


class ConsentUpdateData(ConsentUpdateDataRequired, total=False):
    authorisationUrl: str
    callbackUrl: str
    consentID: str
    emailAddress: str
    expiryDate: str
    failureCallbackUrl: str
    inserted: str
    institutionID: str
    isConnectedAccounts: bool
    isEnabled: bool
    merchantID: str
    provider: str
    successWebHookUrl: str


class ConsentRemoveMatch(TypedDict):
    id: str


class Currency(TypedDict, total=False):
    code: str
    decimals: int
    isFiat: bool
    iso4217AlphaCode: str
    iso4217NumericCode: str
    symbol: str


class CurrencyListMatch(TypedDict, total=False):
    code: str
    decimals: int
    isFiat: bool
    iso4217AlphaCode: str
    iso4217NumericCode: str
    symbol: str


class DirectDebitBatchSubmit(TypedDict, total=False):
    failedSubmissions: dict
    successfulSubmissions: list


class DirectDebitBatchSubmitCreateData(TypedDict, total=False):
    failedSubmissions: dict
    successfulSubmissions: list


class FxRate(TypedDict, total=False):
    destinationCurrency: str
    exchangeRate: float
    expiryTime: str
    quoteID: str
    sourceCurrency: str


class FxRateLoadMatch(TypedDict):
    destination: str
    source: str
    valid_for_minute: int


class FxRateListMatch(TypedDict):
    destination: str
    source: str


class IPayment(TypedDict, total=False):
    paymentRequestID: str
    responseType: str


class IPaymentCreateData(TypedDict, total=False):
    paymentRequestID: str
    responseType: str


class MandateRequired(TypedDict):
    addressLine1: str
    city: str
    countryCode: str
    emailAddress: str
    firstName: str
    lastName: str
    postalCode: str


class Mandate(MandateRequired, total=False):
    accountNumber: str
    addressLine2: str
    approvedAt: str
    currency: str
    customerAccountNumber: str
    customerCity: str
    customerCountryCode: str
    customerCountryName: str
    customerEmailAddress: str
    customerFirstName: str
    customerIban: str
    customerLastName: str
    customerSortCode: str
    iban: str
    id: str
    inserted: str
    isRecurring: bool
    lastUpdated: str
    merchantID: str
    reference: str
    sortCode: str
    status: str
    supplierBankAccountID: str
    supplierCustomerID: str
    supplierMandateID: str
    supplierName: str
    supplierStatus: str


class MandateLoadMatch(TypedDict):
    id: str


class MandateCreateDataRequired(TypedDict):
    addressLine1: str
    city: str
    countryCode: str
    emailAddress: str
    firstName: str
    lastName: str
    postalCode: str


class MandateCreateData(MandateCreateDataRequired, total=False):
    accountNumber: str
    addressLine2: str
    approvedAt: str
    currency: str
    customerAccountNumber: str
    customerCity: str
    customerCountryCode: str
    customerCountryName: str
    customerEmailAddress: str
    customerFirstName: str
    customerIban: str
    customerLastName: str
    customerSortCode: str
    iban: str
    id: str
    inserted: str
    isRecurring: bool
    lastUpdated: str
    merchantID: str
    reference: str
    sortCode: str
    status: str
    supplierBankAccountID: str
    supplierCustomerID: str
    supplierMandateID: str
    supplierName: str
    supplierStatus: str


class Merchant(TypedDict, total=False):
    accountCurrencies: list
    canHaveTrustAccounts: bool
    cardPaymentProcessor: str
    companyID: str
    displayQrOnHostedPay: bool
    hostedPayVersion: int
    id: str
    inserted: str
    isBlocked: bool
    isExited: bool
    isSuspended: bool
    jurisdiction: str
    logoUrlPng: str
    logoUrlSvg: str
    merchantCategoryCode: str
    name: str
    notes: str
    parentMerchant: dict
    paymentAccountLimit: int
    paymentAccounts: list
    reason: str
    shortName: str
    supportedPaymentMethodsList: list
    suspensionReason: str
    tags: list
    timeZoneId: str
    tradingName: str
    webHookLimit: int
    yourRoleName: str


class MerchantLoadMatch(TypedDict):
    id: str


class MerchantListMatch(TypedDict, total=False):
    accountCurrencies: list
    canHaveTrustAccounts: bool
    cardPaymentProcessor: str
    companyID: str
    displayQrOnHostedPay: bool
    hostedPayVersion: int
    id: str
    inserted: str
    isBlocked: bool
    isExited: bool
    isSuspended: bool
    jurisdiction: str
    logoUrlPng: str
    logoUrlSvg: str
    merchantCategoryCode: str
    name: str
    notes: str
    parentMerchant: dict
    paymentAccountLimit: int
    paymentAccounts: list
    reason: str
    shortName: str
    supportedPaymentMethodsList: list
    suspensionReason: str
    tags: list
    timeZoneId: str
    tradingName: str
    webHookLimit: int
    yourRoleName: str


class MerchantUpdateDataRequired(TypedDict):
    id: str


class MerchantUpdateData(MerchantUpdateDataRequired, total=False):
    accountCurrencies: list
    canHaveTrustAccounts: bool
    cardPaymentProcessor: str
    companyID: str
    displayQrOnHostedPay: bool
    hostedPayVersion: int
    inserted: str
    isBlocked: bool
    isExited: bool
    isSuspended: bool
    jurisdiction: str
    logoUrlPng: str
    logoUrlSvg: str
    merchantCategoryCode: str
    name: str
    notes: str
    parentMerchant: dict
    paymentAccountLimit: int
    paymentAccounts: list
    reason: str
    shortName: str
    supportedPaymentMethodsList: list
    suspensionReason: str
    tags: list
    timeZoneId: str
    tradingName: str
    webHookLimit: int
    yourRoleName: str


class MerchantRemoveMatch(TypedDict):
    id: str
    user_id: str


class MerchantAuthorisationSetting(TypedDict, total=False):
    amountLower: float
    amountUpper: float
    authorisationType: str
    beneficiariesOnly: bool
    id: str
    inserted: str
    lastEditorCantAuthorise: bool
    lastUpdated: str
    merchantID: str
    numberOfAuthorisers: int
    roleSettings: list


class MerchantAuthorisationSettingListMatch(TypedDict):
    merchant_id: str


class MerchantDirectDebitMandatePage(TypedDict, total=False):
    approvedAt: str
    currency: str
    customerAccountNumber: str
    customerCity: str
    customerCountryCode: str
    customerCountryName: str
    customerEmailAddress: str
    customerFirstName: str
    customerIban: str
    customerLastName: str
    customerSortCode: str
    id: str
    inserted: str
    isRecurring: bool
    lastUpdated: str
    merchantID: str
    reference: str
    status: str
    supplierBankAccountID: str
    supplierCustomerID: str
    supplierMandateID: str
    supplierName: str
    supplierStatus: str


class MerchantDirectDebitMandatePageListMatch(TypedDict, total=False):
    approvedAt: str
    currency: str
    customerAccountNumber: str
    customerCity: str
    customerCountryCode: str
    customerCountryName: str
    customerEmailAddress: str
    customerFirstName: str
    customerIban: str
    customerLastName: str
    customerSortCode: str
    id: str
    inserted: str
    isRecurring: bool
    lastUpdated: str
    merchantID: str
    reference: str
    status: str
    supplierBankAccountID: str
    supplierCustomerID: str
    supplierMandateID: str
    supplierName: str
    supplierStatus: str


class MerchantPayByBankSetting(TypedDict, total=False):
    bankCountryCodes: list
    bankID: str
    bankName: str
    businessInstitutionID: str
    currency: str
    logo: str
    message: str
    messageImageUrl: str
    order: int
    personalInstitutionID: str
    processor: str
    warningHeading: str
    warningMessage: str


class MerchantPayByBankSettingListMatch(TypedDict):
    merchant_id: str


class MerchantPaymentRequestTemplateRequired(TypedDict):
    description: str
    name: str
    template: dict


class MerchantPaymentRequestTemplate(MerchantPaymentRequestTemplateRequired, total=False):
    bankPaymentOptions: dict
    cardPaymentAddressOptions: dict
    cardPaymentCaptureOptions: dict
    customFields: list
    defaultFields: list
    id: str
    inserted: str
    lastUpdated: str
    merchantID: str
    notificationOptions: dict
    paymentMethods: dict
    paymentTerms: dict
    priorityBankOptions: dict


class MerchantPaymentRequestTemplateLoadMatch(TypedDict):
    id: str
    paymentrequest_id: str


class MerchantPaymentRequestTemplateListMatch(TypedDict):
    merchant_id: str


class MerchantPaymentRequestTemplateUpdateDataRequired(TypedDict):
    id: str
    paymentrequest_id: str


class MerchantPaymentRequestTemplateUpdateData(MerchantPaymentRequestTemplateUpdateDataRequired, total=False):
    bankPaymentOptions: dict
    cardPaymentAddressOptions: dict
    cardPaymentCaptureOptions: dict
    customFields: list
    defaultFields: list
    description: str
    inserted: str
    lastUpdated: str
    merchantID: str
    name: str
    notificationOptions: dict
    paymentMethods: dict
    paymentTerms: dict
    priorityBankOptions: dict
    template: dict


class MerchantPaymentRequestTemplateRemoveMatch(TypedDict):
    id: str
    paymentrequest_id: str


class MerchantTokenRequired(TypedDict):
    nonce: str


class MerchantToken(MerchantTokenRequired, total=False):
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    canAuthorise: bool
    description: str
    expiresAt: str
    hasCurrentUserAuthorised: bool
    hmacAlgorithm: str
    id: str
    inserted: str
    ipAddressWhitelist: str
    isArchived: bool
    isEnabled: bool
    lastAuthorised: str
    lastUpdated: str
    merchantID: str
    permissionTypes: list
    requestSignatureVersion: int
    sharedSecretAlgorithm: str
    sharedSecretBase64: str
    token: str


class MerchantTokenLoadMatch(TypedDict):
    id: str


class MerchantTokenListMatch(TypedDict):
    merchant_id: str


class MerchantTokenCreateDataRequired(TypedDict):
    nonce: str


class MerchantTokenCreateData(MerchantTokenCreateDataRequired, total=False):
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    canAuthorise: bool
    description: str
    expiresAt: str
    hasCurrentUserAuthorised: bool
    hmacAlgorithm: str
    id: str
    inserted: str
    ipAddressWhitelist: str
    isArchived: bool
    isEnabled: bool
    lastAuthorised: str
    lastUpdated: str
    merchantID: str
    permissionTypes: list
    requestSignatureVersion: int
    sharedSecretAlgorithm: str
    sharedSecretBase64: str
    token: str


class MerchantTokenUpdateDataRequired(TypedDict):
    id: str


class MerchantTokenUpdateData(MerchantTokenUpdateDataRequired, total=False):
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    canAuthorise: bool
    description: str
    expiresAt: str
    hasCurrentUserAuthorised: bool
    hmacAlgorithm: str
    inserted: str
    ipAddressWhitelist: str
    isArchived: bool
    isEnabled: bool
    lastAuthorised: str
    lastUpdated: str
    merchantID: str
    nonce: str
    permissionTypes: list
    requestSignatureVersion: int
    sharedSecretAlgorithm: str
    sharedSecretBase64: str
    token: str


class Metadata(TypedDict):
    pass


class MetadataLoadMatch(TypedDict):
    pass


class NoFrixionVersion(TypedDict, total=False):
    buildVersion: int
    majorVersion: int
    minorVersion: int
    releaseName: str


class NoFrixionVersionLoadMatch(TypedDict, total=False):
    buildVersion: int
    majorVersion: int
    minorVersion: int
    releaseName: str


class OpenBanking(TypedDict):
    pass


class OpenBankingCreateData(TypedDict):
    account_id: str


class OpenBankingRemoveMatch(TypedDict):
    account_id: str


class PayeeverificationRequired(TypedDict):
    accountName: str
    iban: str


class Payeeverification(PayeeverificationRequired, total=False):
    accountNumber: str
    payeeVerifiedAccountName: str
    result: str
    secondaryIdentification: str
    sortCode: str


class PayeeverificationCreateDataRequired(TypedDict):
    accountName: str
    iban: str


class PayeeverificationCreateData(PayeeverificationCreateDataRequired, total=False):
    accountNumber: str
    payeeVerifiedAccountName: str
    result: str
    secondaryIdentification: str
    sortCode: str


class PaymentRequired(TypedDict):
    createdByUser: dict


class Payment(PaymentRequired, total=False):
    addresses: list
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    autoSendReceipt: bool
    baseOriginUrl: str
    callbackUrl: str
    cardAuthorizeOnly: bool
    cardCreateToken: bool
    cardCreateTokenMode: str
    cardIgnoreCVN: bool
    cardNoPayerAuthentication: bool
    cardProcessorMerchantID: str
    cardStripePaymentIntentID: str
    cardStripePaymentIntentSecret: str
    cardTransmitRawDetails: bool
    currency: str
    customFields: list
    customerEmailAddress: str
    customerID: str
    customerName: str
    description: str
    destinationAccount: dict
    directDebitPayment: dict
    dueDate: str
    events: list
    failureCallbackUrl: str
    fieldDisplaySettings: list
    formattedAmount: str
    hostedPayCheckoutUrl: str
    id: str
    ignoreAddressVerification: bool
    inserted: str
    insertedSortable: str
    isArchived: bool
    jwk: str
    lastUpdated: str
    lightningInvoice: str
    lightningInvoiceExpiresAt: str
    merchantDirectDebitMandateID: str
    merchantID: str
    merchantTokenDescription: str
    notificationEmailAddresses: str
    notificationRoleIDs: list
    orderID: str
    partialPaymentMethod: str
    partialPaymentSteps: str
    paymentAttempts: list
    paymentMethods: list
    paymentProcessor: str
    payrunID: str
    pispAccountID: str
    priorityBankID: str
    result: dict
    sandboxSettleDelayInSeconds: int
    shippingAddress: dict
    shippingAddressCity: str
    shippingAddressCountryCode: str
    shippingAddressCounty: str
    shippingAddressLine1: str
    shippingAddressLine2: str
    shippingAddressPostCode: str
    shippingEmail: str
    shippingFirstName: str
    shippingLastName: str
    shippingPhone: str
    status: str
    successWebHookUrl: str
    tagIds: list
    tags: list
    title: str
    tokenisedCards: list
    transactions: list
    useHostedPaymentPage: bool


class PaymentLoadMatch(TypedDict):
    id: str


class PaymentCreateDataRequired(TypedDict):
    createdByUser: dict


class PaymentCreateData(PaymentCreateDataRequired, total=False):
    addresses: list
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    autoSendReceipt: bool
    baseOriginUrl: str
    callbackUrl: str
    cardAuthorizeOnly: bool
    cardCreateToken: bool
    cardCreateTokenMode: str
    cardIgnoreCVN: bool
    cardNoPayerAuthentication: bool
    cardProcessorMerchantID: str
    cardStripePaymentIntentID: str
    cardStripePaymentIntentSecret: str
    cardTransmitRawDetails: bool
    currency: str
    customFields: list
    customerEmailAddress: str
    customerID: str
    customerName: str
    description: str
    destinationAccount: dict
    directDebitPayment: dict
    dueDate: str
    events: list
    failureCallbackUrl: str
    fieldDisplaySettings: list
    formattedAmount: str
    hostedPayCheckoutUrl: str
    id: str
    ignoreAddressVerification: bool
    inserted: str
    insertedSortable: str
    isArchived: bool
    jwk: str
    lastUpdated: str
    lightningInvoice: str
    lightningInvoiceExpiresAt: str
    merchantDirectDebitMandateID: str
    merchantID: str
    merchantTokenDescription: str
    notificationEmailAddresses: str
    notificationRoleIDs: list
    orderID: str
    partialPaymentMethod: str
    partialPaymentSteps: str
    paymentAttempts: list
    paymentMethods: list
    paymentProcessor: str
    payrunID: str
    pispAccountID: str
    priorityBankID: str
    result: dict
    sandboxSettleDelayInSeconds: int
    shippingAddress: dict
    shippingAddressCity: str
    shippingAddressCountryCode: str
    shippingAddressCounty: str
    shippingAddressLine1: str
    shippingAddressLine2: str
    shippingAddressPostCode: str
    shippingEmail: str
    shippingFirstName: str
    shippingLastName: str
    shippingPhone: str
    status: str
    successWebHookUrl: str
    tagIds: list
    tags: list
    title: str
    tokenisedCards: list
    transactions: list
    useHostedPaymentPage: bool


class PaymentUpdateDataRequired(TypedDict):
    id: str


class PaymentUpdateData(PaymentUpdateDataRequired, total=False):
    addresses: list
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    autoSendReceipt: bool
    baseOriginUrl: str
    callbackUrl: str
    cardAuthorizeOnly: bool
    cardCreateToken: bool
    cardCreateTokenMode: str
    cardIgnoreCVN: bool
    cardNoPayerAuthentication: bool
    cardProcessorMerchantID: str
    cardStripePaymentIntentID: str
    cardStripePaymentIntentSecret: str
    cardTransmitRawDetails: bool
    createdByUser: dict
    currency: str
    customFields: list
    customerEmailAddress: str
    customerID: str
    customerName: str
    description: str
    destinationAccount: dict
    directDebitPayment: dict
    dueDate: str
    events: list
    failureCallbackUrl: str
    fieldDisplaySettings: list
    formattedAmount: str
    hostedPayCheckoutUrl: str
    ignoreAddressVerification: bool
    inserted: str
    insertedSortable: str
    isArchived: bool
    jwk: str
    lastUpdated: str
    lightningInvoice: str
    lightningInvoiceExpiresAt: str
    merchantDirectDebitMandateID: str
    merchantID: str
    merchantTokenDescription: str
    notificationEmailAddresses: str
    notificationRoleIDs: list
    orderID: str
    partialPaymentMethod: str
    partialPaymentSteps: str
    paymentAttempts: list
    paymentMethods: list
    paymentProcessor: str
    payrunID: str
    pispAccountID: str
    priorityBankID: str
    result: dict
    sandboxSettleDelayInSeconds: int
    shippingAddress: dict
    shippingAddressCity: str
    shippingAddressCountryCode: str
    shippingAddressCounty: str
    shippingAddressLine1: str
    shippingAddressLine2: str
    shippingAddressPostCode: str
    shippingEmail: str
    shippingFirstName: str
    shippingLastName: str
    shippingPhone: str
    status: str
    successWebHookUrl: str
    tagIds: list
    tags: list
    title: str
    tokenisedCards: list
    transactions: list
    useHostedPaymentPage: bool


class PaymentAccountRequired(TypedDict):
    createdBy: dict
    identifier: dict


class PaymentAccount(PaymentAccountRequired, total=False):
    accountName: str
    accountSupplierName: str
    availableBalance: float
    availableBalanceMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bankName: str
    consentID: str
    createdByDisplayName: str
    currency: str
    defaultPaymentRail: str
    displayName: str
    expiryDate: str
    externalAccountIcon: str
    id: str
    inserted: str
    isArchived: bool
    isConnectedAccount: bool
    isDefault: bool
    isTrustAccount: bool
    isVirtual: bool
    lastTransaction: dict
    lastUpdated: str
    merchantID: str
    merchantName: str
    physicalAccountID: str
    rules: list
    submittedPayoutsBalance: float
    submittedPayoutsBalanceMinorUnits: int
    summary: str
    supplierSepaInstantStatus: str
    xeroBankFeedConnectionStatus: str
    xeroBankFeedLastSyncedAt: str
    xeroBankFeedSyncLastFailedAt: str
    xeroBankFeedSyncLastFailureReason: str
    xeroBankFeedSyncStatus: str
    xeroUnsynchronisedTransactionsCount: int


class PaymentAccountListMatch(TypedDict, total=False):
    accountName: str
    accountSupplierName: str
    availableBalance: float
    availableBalanceMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bankName: str
    consentID: str
    createdBy: dict
    createdByDisplayName: str
    currency: str
    defaultPaymentRail: str
    displayName: str
    expiryDate: str
    externalAccountIcon: str
    id: str
    identifier: dict
    inserted: str
    isArchived: bool
    isConnectedAccount: bool
    isDefault: bool
    isTrustAccount: bool
    isVirtual: bool
    lastTransaction: dict
    lastUpdated: str
    merchantID: str
    merchantName: str
    physicalAccountID: str
    rules: list
    submittedPayoutsBalance: float
    submittedPayoutsBalanceMinorUnits: int
    summary: str
    supplierSepaInstantStatus: str
    xeroBankFeedConnectionStatus: str
    xeroBankFeedLastSyncedAt: str
    xeroBankFeedSyncLastFailedAt: str
    xeroBankFeedSyncLastFailureReason: str
    xeroBankFeedSyncStatus: str
    xeroUnsynchronisedTransactionsCount: int


class PaymentAccountMinimalRequired(TypedDict):
    identifier: dict


class PaymentAccountMinimal(PaymentAccountMinimalRequired, total=False):
    accountName: str
    availableBalance: float
    balance: float
    balanceMinorUnits: int
    currency: str
    id: str
    isArchived: bool
    isConnectedAccount: bool
    merchantID: str
    submittedPayoutsBalance: float


class PaymentAccountMinimalListMatch(TypedDict, total=False):
    accountName: str
    availableBalance: float
    balance: float
    balanceMinorUnits: int
    currency: str
    id: str
    identifier: dict
    isArchived: bool
    isConnectedAccount: bool
    merchantID: str
    submittedPayoutsBalance: float


class PaymentInitiation(TypedDict, total=False):
    paymentInitiationID: str
    paymentRequestCallbackUrl: str
    paymentRequestID: str
    redirectUrl: str
    responseType: str
    specificErrorMessage: str


class PaymentInitiationCreateDataRequired(TypedDict):
    paymentrequest_id: str


class PaymentInitiationCreateData(PaymentInitiationCreateDataRequired, total=False):
    paymentInitiationID: str
    paymentRequestCallbackUrl: str
    paymentRequestID: str
    redirectUrl: str
    responseType: str
    specificErrorMessage: str


class PaymentRequestRequired(TypedDict):
    createdByUser: dict


class PaymentRequest(PaymentRequestRequired, total=False):
    addresses: list
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    autoSendReceipt: bool
    baseOriginUrl: str
    callbackUrl: str
    cardAuthorizeOnly: bool
    cardCreateToken: bool
    cardCreateTokenMode: str
    cardIgnoreCVN: bool
    cardProcessorMerchantID: str
    cardStripePaymentIntentID: str
    cardStripePaymentIntentSecret: str
    currency: str
    customFields: list
    customerEmailAddress: str
    customerID: str
    customerName: str
    description: str
    destinationAccount: dict
    directDebitPayment: dict
    doSimulateSettlementFailure: bool
    dueDate: str
    errorDescription: str
    events: list
    failedPaymentRequests: dict
    failureCallbackUrl: str
    fieldDisplaySettings: list
    formattedAmount: str
    hostedPayCheckoutUrl: str
    id: str
    ignoreAddressVerification: bool
    inserted: str
    insertedSortable: str
    institution: str
    isArchived: bool
    jwk: str
    lastUpdated: str
    lightningInvoice: str
    lightningInvoiceExpiresAt: str
    merchantDirectDebitMandateID: str
    merchantID: str
    merchantTokenDescription: str
    notificationEmailAddresses: str
    notificationRoleIDs: list
    orderID: str
    partialPaymentMethod: str
    partialPaymentSteps: str
    paymentAttempts: list
    paymentInitiationID: str
    paymentMethods: list
    paymentProcessor: str
    paymentRequests: list
    payrunID: str
    pispAccountID: str
    priorityBankID: str
    result: dict
    sandboxSettleDelayInSeconds: int
    shippingAddress: dict
    status: str
    successWebHookUrl: str
    tags: list
    title: str
    tokenisedCards: list
    transactions: list
    useHostedPaymentPage: bool


class PaymentRequestLoadMatchRequired(TypedDict):
    id: str


class PaymentRequestLoadMatch(PaymentRequestLoadMatchRequired, total=False):
    addresses: list
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    autoSendReceipt: bool
    baseOriginUrl: str
    callbackUrl: str
    cardAuthorizeOnly: bool
    cardCreateToken: bool
    cardCreateTokenMode: str
    cardIgnoreCVN: bool
    cardProcessorMerchantID: str
    cardStripePaymentIntentID: str
    cardStripePaymentIntentSecret: str
    createdByUser: dict
    currency: str
    customFields: list
    customerEmailAddress: str
    customerID: str
    customerName: str
    description: str
    destinationAccount: dict
    directDebitPayment: dict
    doSimulateSettlementFailure: bool
    dueDate: str
    errorDescription: str
    events: list
    failedPaymentRequests: dict
    failureCallbackUrl: str
    fieldDisplaySettings: list
    formattedAmount: str
    hostedPayCheckoutUrl: str
    ignoreAddressVerification: bool
    inserted: str
    insertedSortable: str
    institution: str
    isArchived: bool
    jwk: str
    lastUpdated: str
    lightningInvoice: str
    lightningInvoiceExpiresAt: str
    merchantDirectDebitMandateID: str
    merchantID: str
    merchantTokenDescription: str
    notificationEmailAddresses: str
    notificationRoleIDs: list
    orderID: str
    partialPaymentMethod: str
    partialPaymentSteps: str
    paymentAttempts: list
    paymentInitiationID: str
    paymentMethods: list
    paymentProcessor: str
    paymentRequests: list
    payrunID: str
    pispAccountID: str
    priorityBankID: str
    result: dict
    sandboxSettleDelayInSeconds: int
    shippingAddress: dict
    status: str
    successWebHookUrl: str
    tags: list
    title: str
    tokenisedCards: list
    transactions: list
    useHostedPaymentPage: bool


class PaymentRequestListMatch(TypedDict, total=False):
    addresses: list
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    autoSendReceipt: bool
    baseOriginUrl: str
    callbackUrl: str
    cardAuthorizeOnly: bool
    cardCreateToken: bool
    cardCreateTokenMode: str
    cardIgnoreCVN: bool
    cardProcessorMerchantID: str
    cardStripePaymentIntentID: str
    cardStripePaymentIntentSecret: str
    createdByUser: dict
    currency: str
    customFields: list
    customerEmailAddress: str
    customerID: str
    customerName: str
    description: str
    destinationAccount: dict
    directDebitPayment: dict
    doSimulateSettlementFailure: bool
    dueDate: str
    errorDescription: str
    events: list
    failedPaymentRequests: dict
    failureCallbackUrl: str
    fieldDisplaySettings: list
    formattedAmount: str
    hostedPayCheckoutUrl: str
    id: str
    ignoreAddressVerification: bool
    inserted: str
    insertedSortable: str
    institution: str
    isArchived: bool
    jwk: str
    lastUpdated: str
    lightningInvoice: str
    lightningInvoiceExpiresAt: str
    merchantDirectDebitMandateID: str
    merchantID: str
    merchantTokenDescription: str
    notificationEmailAddresses: str
    notificationRoleIDs: list
    orderID: str
    partialPaymentMethod: str
    partialPaymentSteps: str
    paymentAttempts: list
    paymentInitiationID: str
    paymentMethods: list
    paymentProcessor: str
    paymentRequests: list
    payrunID: str
    pispAccountID: str
    priorityBankID: str
    result: dict
    sandboxSettleDelayInSeconds: int
    shippingAddress: dict
    status: str
    successWebHookUrl: str
    tags: list
    title: str
    tokenisedCards: list
    transactions: list
    useHostedPaymentPage: bool


class PaymentRequestCreateDataRequired(TypedDict):
    createdByUser: dict


class PaymentRequestCreateData(PaymentRequestCreateDataRequired, total=False):
    addresses: list
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    autoSendReceipt: bool
    baseOriginUrl: str
    callbackUrl: str
    cardAuthorizeOnly: bool
    cardCreateToken: bool
    cardCreateTokenMode: str
    cardIgnoreCVN: bool
    cardProcessorMerchantID: str
    cardStripePaymentIntentID: str
    cardStripePaymentIntentSecret: str
    currency: str
    customFields: list
    customerEmailAddress: str
    customerID: str
    customerName: str
    description: str
    destinationAccount: dict
    directDebitPayment: dict
    doSimulateSettlementFailure: bool
    dueDate: str
    errorDescription: str
    events: list
    failedPaymentRequests: dict
    failureCallbackUrl: str
    fieldDisplaySettings: list
    formattedAmount: str
    hostedPayCheckoutUrl: str
    id: str
    ignoreAddressVerification: bool
    inserted: str
    insertedSortable: str
    institution: str
    isArchived: bool
    jwk: str
    lastUpdated: str
    lightningInvoice: str
    lightningInvoiceExpiresAt: str
    merchantDirectDebitMandateID: str
    merchantID: str
    merchantTokenDescription: str
    notificationEmailAddresses: str
    notificationRoleIDs: list
    orderID: str
    partialPaymentMethod: str
    partialPaymentSteps: str
    paymentAttempts: list
    paymentInitiationID: str
    paymentMethods: list
    paymentProcessor: str
    paymentRequests: list
    payrunID: str
    pispAccountID: str
    priorityBankID: str
    result: dict
    sandboxSettleDelayInSeconds: int
    shippingAddress: dict
    status: str
    successWebHookUrl: str
    tags: list
    title: str
    tokenisedCards: list
    transactions: list
    useHostedPaymentPage: bool


class PaymentRequestUpdateDataRequired(TypedDict):
    paymentrequest_id: str


class PaymentRequestUpdateData(PaymentRequestUpdateDataRequired, total=False):
    addresses: list
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    autoSendReceipt: bool
    baseOriginUrl: str
    callbackUrl: str
    cardAuthorizeOnly: bool
    cardCreateToken: bool
    cardCreateTokenMode: str
    cardIgnoreCVN: bool
    cardProcessorMerchantID: str
    cardStripePaymentIntentID: str
    cardStripePaymentIntentSecret: str
    createdByUser: dict
    currency: str
    customFields: list
    customerEmailAddress: str
    customerID: str
    customerName: str
    description: str
    destinationAccount: dict
    directDebitPayment: dict
    doSimulateSettlementFailure: bool
    dueDate: str
    errorDescription: str
    events: list
    failedPaymentRequests: dict
    failureCallbackUrl: str
    fieldDisplaySettings: list
    formattedAmount: str
    hostedPayCheckoutUrl: str
    id: str
    ignoreAddressVerification: bool
    inserted: str
    insertedSortable: str
    institution: str
    isArchived: bool
    jwk: str
    lastUpdated: str
    lightningInvoice: str
    lightningInvoiceExpiresAt: str
    merchantDirectDebitMandateID: str
    merchantID: str
    merchantTokenDescription: str
    notificationEmailAddresses: str
    notificationRoleIDs: list
    orderID: str
    partialPaymentMethod: str
    partialPaymentSteps: str
    paymentAttempts: list
    paymentInitiationID: str
    paymentMethods: list
    paymentProcessor: str
    paymentRequests: list
    payrunID: str
    pispAccountID: str
    priorityBankID: str
    result: dict
    sandboxSettleDelayInSeconds: int
    shippingAddress: dict
    status: str
    successWebHookUrl: str
    tags: list
    title: str
    tokenisedCards: list
    transactions: list
    useHostedPaymentPage: bool


class PaymentRequestRemoveMatch(TypedDict):
    id: str


class PaymentRequestEventRequired(TypedDict):
    amount: float


class PaymentRequestEvent(PaymentRequestEventRequired, total=False):
    applePayTransactionID: str
    cardAuthorizationResponseID: str
    cardExpiryMonth: int
    cardExpiryYear: int
    cardIssuer: str
    cardIssuerCountry: str
    cardLastFourDigits: str
    cardRequestID: str
    cardScheme: str
    cardTokenCustomerID: str
    cardTransactionID: str
    currency: str
    directDebitPaymentID: str
    directDebitPaymentReference: str
    drirectDebitMandateID: str
    errorMessage: str
    errorReason: str
    eventType: str
    id: str
    inserted: str
    lightningInvoice: str
    lightningRHash: str
    originUrl: str
    paymentMethodType: str
    paymentProcessorName: str
    paymentRequestID: str
    pispBankStatus: str
    pispPaymentInitiationID: str
    pispPaymentInstitutionName: str
    pispPaymentServiceProviderID: str
    pispRedirectUrl: str
    reconciledTransactionID: str
    refundPayoutID: str
    status: str
    walletName: str


class PaymentRequestEventListMatch(TypedDict):
    paymentrequest_id: str


class PaymentRequestMetric(TypedDict):
    pass


class PaymentRequestMetricLoadMatch(TypedDict):
    pass


class PaymentRequestMinimal(TypedDict, total=False):
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    callbackUrl: str
    cardStripePaymentIntentSecret: str
    countryCode: str
    currency: str
    customFieldsToDisplay: list
    description: str
    dueDate: str
    fieldDisplaySettings: list
    googlePayMerchantID: str
    id: str
    jwk: str
    merchantID: str
    merchantLogoUrlPng: str
    merchantLogoUrlSvg: str
    merchantName: str
    merchantShortName: str
    partialPaymentMethod: str
    paymentAttempts: list
    paymentMethodsList: list
    paymentProcessor: str
    paymentProcessorKey: str
    pispError: str
    priorityBankID: str
    status: str
    stripeAccountID: str
    title: str


class PaymentRequestMinimalListMatch(TypedDict):
    paymentrequest_id: str


class PaymentRequestResult(TypedDict, total=False):
    amount: float
    amountPending: float
    amountReceived: float
    amountRefunded: float
    currency: str
    customerID: str
    paymentRequestID: str
    payments: list
    pispAuthorizations: list
    requestedAmount: float
    result: str


class PaymentRequestResultListMatch(TypedDict):
    paymentrequest_id: str


class PayoutRequired(TypedDict):
    beneficiary: dict
    sourceAccountIdentifier: dict


class Payout(PayoutRequired, total=False):
    accountID: str
    allowIncomplete: bool
    amount: float
    amountMinorUnits: int
    approvePayoutUrl: str
    approverID: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    batchPayoutID: str
    beneficiaryID: str
    canAuthorise: bool
    canProcess: bool
    canUpdate: bool
    chargeBearer: str
    createdBy: str
    createdByEmailAddress: str
    currency: str
    currentUserID: str
    description: str
    destination: dict
    documents: list
    events: list
    failedPayouts: dict
    formattedAmount: str
    formattedFxDestinationAmount: str
    formattedSchedule: str
    formattedScheduleDayOnly: str
    formattedSourceAccountAvailableBalance: str
    fxDestinationAmount: float
    fxDestinationAmountMinorUnits: int
    fxDestinationCurrency: str
    fxQuoteExpiresAt: str
    fxQuoteID: str
    fxRate: float
    fxUseDestinationAmount: bool
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    invoiceID: str
    isArchived: bool
    isFailed: bool
    isSettled: bool
    isSubmitted: bool
    lastUpdated: str
    merchantID: str
    merchantTokenDescription: str
    nonce: str
    paymentProcessor: str
    paymentRail: str
    payouts: list
    payrunID: str
    payrunName: str
    reason: str
    rule: dict
    scheduleDate: str
    scheduled: bool
    sourceAccountAvailableBalance: float
    sourceAccountAvailableBalanceMinorUnits: int
    sourceAccountBic: str
    sourceAccountCurrency: str
    sourceAccountIban: str
    sourceAccountName: str
    sourceAccountNumber: str
    sourceAccountSortcode: str
    status: str
    tagIds: list
    tags: list
    theirReference: str
    topupPayrunID: str
    transactedAmount: float
    transactedFxAmount: float
    transactedFxRate: float
    type: str
    userID: str
    yourReference: str


class PayoutLoadMatch(TypedDict):
    id: str


class PayoutListMatch(TypedDict, total=False):
    accountID: str
    allowIncomplete: bool
    amount: float
    amountMinorUnits: int
    approvePayoutUrl: str
    approverID: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    batchPayoutID: str
    beneficiary: dict
    beneficiaryID: str
    canAuthorise: bool
    canProcess: bool
    canUpdate: bool
    chargeBearer: str
    createdBy: str
    createdByEmailAddress: str
    currency: str
    currentUserID: str
    description: str
    destination: dict
    documents: list
    events: list
    failedPayouts: dict
    formattedAmount: str
    formattedFxDestinationAmount: str
    formattedSchedule: str
    formattedScheduleDayOnly: str
    formattedSourceAccountAvailableBalance: str
    fxDestinationAmount: float
    fxDestinationAmountMinorUnits: int
    fxDestinationCurrency: str
    fxQuoteExpiresAt: str
    fxQuoteID: str
    fxRate: float
    fxUseDestinationAmount: bool
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    invoiceID: str
    isArchived: bool
    isFailed: bool
    isSettled: bool
    isSubmitted: bool
    lastUpdated: str
    merchantID: str
    merchantTokenDescription: str
    nonce: str
    paymentProcessor: str
    paymentRail: str
    payouts: list
    payrunID: str
    payrunName: str
    reason: str
    rule: dict
    scheduleDate: str
    scheduled: bool
    sourceAccountAvailableBalance: float
    sourceAccountAvailableBalanceMinorUnits: int
    sourceAccountBic: str
    sourceAccountCurrency: str
    sourceAccountIban: str
    sourceAccountIdentifier: dict
    sourceAccountName: str
    sourceAccountNumber: str
    sourceAccountSortcode: str
    status: str
    tagIds: list
    tags: list
    theirReference: str
    topupPayrunID: str
    transactedAmount: float
    transactedFxAmount: float
    transactedFxRate: float
    type: str
    userID: str
    yourReference: str


class PayoutCreateDataRequired(TypedDict):
    id: str
    beneficiary: dict
    sourceAccountIdentifier: dict


class PayoutCreateData(PayoutCreateDataRequired, total=False):
    accountID: str
    allowIncomplete: bool
    amount: float
    amountMinorUnits: int
    approvePayoutUrl: str
    approverID: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    batchPayoutID: str
    beneficiaryID: str
    canAuthorise: bool
    canProcess: bool
    canUpdate: bool
    chargeBearer: str
    createdBy: str
    createdByEmailAddress: str
    currency: str
    currentUserID: str
    description: str
    destination: dict
    documents: list
    events: list
    failedPayouts: dict
    formattedAmount: str
    formattedFxDestinationAmount: str
    formattedSchedule: str
    formattedScheduleDayOnly: str
    formattedSourceAccountAvailableBalance: str
    fxDestinationAmount: float
    fxDestinationAmountMinorUnits: int
    fxDestinationCurrency: str
    fxQuoteExpiresAt: str
    fxQuoteID: str
    fxRate: float
    fxUseDestinationAmount: bool
    hasCurrentUserAuthorised: bool
    inserted: str
    invoiceID: str
    isArchived: bool
    isFailed: bool
    isSettled: bool
    isSubmitted: bool
    lastUpdated: str
    merchantID: str
    merchantTokenDescription: str
    nonce: str
    paymentProcessor: str
    paymentRail: str
    payouts: list
    payrunID: str
    payrunName: str
    reason: str
    rule: dict
    scheduleDate: str
    scheduled: bool
    sourceAccountAvailableBalance: float
    sourceAccountAvailableBalanceMinorUnits: int
    sourceAccountBic: str
    sourceAccountCurrency: str
    sourceAccountIban: str
    sourceAccountName: str
    sourceAccountNumber: str
    sourceAccountSortcode: str
    status: str
    tagIds: list
    tags: list
    theirReference: str
    topupPayrunID: str
    transactedAmount: float
    transactedFxAmount: float
    transactedFxRate: float
    type: str
    userID: str
    yourReference: str


class PayoutUpdateDataRequired(TypedDict):
    id: str


class PayoutUpdateData(PayoutUpdateDataRequired, total=False):
    accountID: str
    allowIncomplete: bool
    amount: float
    amountMinorUnits: int
    approvePayoutUrl: str
    approverID: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    batchPayoutID: str
    beneficiary: dict
    beneficiaryID: str
    canAuthorise: bool
    canProcess: bool
    canUpdate: bool
    chargeBearer: str
    createdBy: str
    createdByEmailAddress: str
    currency: str
    currentUserID: str
    description: str
    destination: dict
    documents: list
    events: list
    failedPayouts: dict
    formattedAmount: str
    formattedFxDestinationAmount: str
    formattedSchedule: str
    formattedScheduleDayOnly: str
    formattedSourceAccountAvailableBalance: str
    fxDestinationAmount: float
    fxDestinationAmountMinorUnits: int
    fxDestinationCurrency: str
    fxQuoteExpiresAt: str
    fxQuoteID: str
    fxRate: float
    fxUseDestinationAmount: bool
    hasCurrentUserAuthorised: bool
    inserted: str
    invoiceID: str
    isArchived: bool
    isFailed: bool
    isSettled: bool
    isSubmitted: bool
    lastUpdated: str
    merchantID: str
    merchantTokenDescription: str
    nonce: str
    paymentProcessor: str
    paymentRail: str
    payouts: list
    payrunID: str
    payrunName: str
    reason: str
    rule: dict
    scheduleDate: str
    scheduled: bool
    sourceAccountAvailableBalance: float
    sourceAccountAvailableBalanceMinorUnits: int
    sourceAccountBic: str
    sourceAccountCurrency: str
    sourceAccountIban: str
    sourceAccountIdentifier: dict
    sourceAccountName: str
    sourceAccountNumber: str
    sourceAccountSortcode: str
    status: str
    tagIds: list
    tags: list
    theirReference: str
    topupPayrunID: str
    transactedAmount: float
    transactedFxAmount: float
    transactedFxRate: float
    type: str
    userID: str
    yourReference: str


class PayoutRemoveMatch(TypedDict):
    id: str


class PayoutKeysetPageRequired(TypedDict):
    beneficiary: dict
    sourceAccountIdentifier: dict


class PayoutKeysetPage(PayoutKeysetPageRequired, total=False):
    accountID: str
    amount: float
    amountMinorUnits: int
    approvePayoutUrl: str
    approverID: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    batchPayoutID: str
    canAuthorise: bool
    canProcess: bool
    canUpdate: bool
    chargeBearer: str
    createdBy: str
    createdByEmailAddress: str
    currency: str
    currentUserID: str
    description: str
    destination: dict
    documents: list
    events: list
    formattedAmount: str
    formattedFxDestinationAmount: str
    formattedSchedule: str
    formattedScheduleDayOnly: str
    formattedSourceAccountAvailableBalance: str
    fxDestinationAmount: float
    fxDestinationAmountMinorUnits: int
    fxDestinationCurrency: str
    fxQuoteExpiresAt: str
    fxQuoteID: str
    fxRate: float
    fxUseDestinationAmount: bool
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    invoiceID: str
    isArchived: bool
    isFailed: bool
    isSettled: bool
    isSubmitted: bool
    lastUpdated: str
    merchantID: str
    merchantTokenDescription: str
    nonce: str
    paymentProcessor: str
    paymentRail: str
    payrunID: str
    payrunName: str
    rule: dict
    scheduleDate: str
    scheduled: bool
    sourceAccountAvailableBalance: float
    sourceAccountAvailableBalanceMinorUnits: int
    sourceAccountBic: str
    sourceAccountCurrency: str
    sourceAccountIban: str
    sourceAccountName: str
    sourceAccountNumber: str
    sourceAccountSortcode: str
    status: str
    tags: list
    theirReference: str
    topupPayrunID: str
    transactedAmount: float
    transactedFxAmount: float
    transactedFxRate: float
    type: str
    userID: str
    yourReference: str


class PayoutKeysetPageListMatch(TypedDict):
    merchant_id: str


class PayoutMetric(TypedDict):
    pass


class PayoutMetricLoadMatch(TypedDict):
    pass


class PayrunRequired(TypedDict):
    lastUpdatedBy: dict


class Payrun(PayrunRequired, total=False):
    authorisationDate: str
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    batchPayoutID: str
    canAuthorise: bool
    canDelete: bool
    canEdit: bool
    events: list
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    invoices: list
    invoicesMinimal: list
    isArchived: bool
    lastUpdated: str
    merchantID: str
    name: str
    nonce: str
    notes: str
    payments: list
    payouts: list
    payoutsCount: int
    reason: str
    scheduleDate: str
    scheduledDate: str
    sourceAccounts: list
    status: str
    totalEur: float
    totalGbp: float
    totalUsd: float


class PayrunLoadMatch(TypedDict):
    id: str


class PayrunListMatch(TypedDict, total=False):
    authorisationDate: str
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    batchPayoutID: str
    canAuthorise: bool
    canDelete: bool
    canEdit: bool
    events: list
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    invoices: list
    invoicesMinimal: list
    isArchived: bool
    lastUpdated: str
    lastUpdatedBy: dict
    merchantID: str
    name: str
    nonce: str
    notes: str
    payments: list
    payouts: list
    payoutsCount: int
    reason: str
    scheduleDate: str
    scheduledDate: str
    sourceAccounts: list
    status: str
    totalEur: float
    totalGbp: float
    totalUsd: float


class PayrunCreateDataRequired(TypedDict):
    id: str
    lastUpdatedBy: dict


class PayrunCreateData(PayrunCreateDataRequired, total=False):
    authorisationDate: str
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    batchPayoutID: str
    canAuthorise: bool
    canDelete: bool
    canEdit: bool
    events: list
    hasCurrentUserAuthorised: bool
    inserted: str
    invoices: list
    invoicesMinimal: list
    isArchived: bool
    lastUpdated: str
    merchantID: str
    name: str
    nonce: str
    notes: str
    payments: list
    payouts: list
    payoutsCount: int
    reason: str
    scheduleDate: str
    scheduledDate: str
    sourceAccounts: list
    status: str
    totalEur: float
    totalGbp: float
    totalUsd: float


class PayrunUpdateDataRequired(TypedDict):
    id: str


class PayrunUpdateData(PayrunUpdateDataRequired, total=False):
    authorisationDate: str
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    batchPayoutID: str
    canAuthorise: bool
    canDelete: bool
    canEdit: bool
    events: list
    hasCurrentUserAuthorised: bool
    inserted: str
    invoices: list
    invoicesMinimal: list
    isArchived: bool
    lastUpdated: str
    lastUpdatedBy: dict
    merchantID: str
    name: str
    nonce: str
    notes: str
    payments: list
    payouts: list
    payoutsCount: int
    reason: str
    scheduleDate: str
    scheduledDate: str
    sourceAccounts: list
    status: str
    totalEur: float
    totalGbp: float
    totalUsd: float


class PayrunRemoveMatch(TypedDict):
    id: str


class Report(TypedDict):
    pass


class ReportUpdateData(TypedDict):
    id: str


class ReportResult(TypedDict, total=False):
    contentType: str
    contents: str
    lastCompletedAt: str
    merchantID: str
    reportName: str
    reportType: str
    statementNumber: int


class ReportResultLoadMatch(TypedDict):
    id: int
    report_id: str


class Role(TypedDict, total=False):
    failedRoles: dict
    roles: list


class RoleCreateDataRequired(TypedDict):
    merchant_id: str


class RoleCreateData(RoleCreateDataRequired, total=False):
    failedRoles: dict
    roles: list


class RuleRequired(TypedDict):
    createdBy: dict
    nonce: str


class Rule(RuleRequired, total=False):
    account: dict
    accountID: str
    approveUrl: str
    approverID: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    canAuthorise: bool
    description: str
    endAt: str
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    isDisabled: bool
    lastExecutedAt: str
    lastRunAtTransactionDate: str
    lastUpdated: str
    merchantID: str
    name: str
    onApprovedWebHookUrl: str
    onExecutionErrorWebHookUrl: str
    onExecutionSuccessWebHookUrl: str
    startAt: str
    status: str
    sweepAction: dict
    timeZoneId: str
    triggerCronExpression: str
    triggerOnPayIn: bool
    userID: str
    webHookSecret: str


class RuleLoadMatch(TypedDict):
    id: str


class RuleListMatch(TypedDict, total=False):
    account: dict
    accountID: str
    approveUrl: str
    approverID: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    canAuthorise: bool
    createdBy: dict
    description: str
    endAt: str
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    isDisabled: bool
    lastExecutedAt: str
    lastRunAtTransactionDate: str
    lastUpdated: str
    merchantID: str
    name: str
    nonce: str
    onApprovedWebHookUrl: str
    onExecutionErrorWebHookUrl: str
    onExecutionSuccessWebHookUrl: str
    startAt: str
    status: str
    sweepAction: dict
    timeZoneId: str
    triggerCronExpression: str
    triggerOnPayIn: bool
    userID: str
    webHookSecret: str


class RuleCreateDataRequired(TypedDict):
    createdBy: dict
    nonce: str


class RuleCreateData(RuleCreateDataRequired, total=False):
    account: dict
    accountID: str
    approveUrl: str
    approverID: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    canAuthorise: bool
    description: str
    endAt: str
    hasCurrentUserAuthorised: bool
    id: str
    inserted: str
    isDisabled: bool
    lastExecutedAt: str
    lastRunAtTransactionDate: str
    lastUpdated: str
    merchantID: str
    name: str
    onApprovedWebHookUrl: str
    onExecutionErrorWebHookUrl: str
    onExecutionSuccessWebHookUrl: str
    startAt: str
    status: str
    sweepAction: dict
    timeZoneId: str
    triggerCronExpression: str
    triggerOnPayIn: bool
    userID: str
    webHookSecret: str


class RuleUpdateDataRequired(TypedDict):
    id: str


class RuleUpdateData(RuleUpdateDataRequired, total=False):
    account: dict
    accountID: str
    approveUrl: str
    approverID: str
    authenticationMethods: list
    authorisations: list
    authorisersCompletedCount: int
    authorisersRequiredCount: int
    canAuthorise: bool
    createdBy: dict
    description: str
    endAt: str
    hasCurrentUserAuthorised: bool
    inserted: str
    isDisabled: bool
    lastExecutedAt: str
    lastRunAtTransactionDate: str
    lastUpdated: str
    merchantID: str
    name: str
    nonce: str
    onApprovedWebHookUrl: str
    onExecutionErrorWebHookUrl: str
    onExecutionSuccessWebHookUrl: str
    startAt: str
    status: str
    sweepAction: dict
    timeZoneId: str
    triggerCronExpression: str
    triggerOnPayIn: bool
    userID: str
    webHookSecret: str


class RuleRemoveMatch(TypedDict):
    id: str


class RuleEventRequired(TypedDict):
    user: dict


class RuleEvent(RuleEventRequired, total=False):
    errorMessage: str
    id: str
    inserted: str
    isAuthoriseToEnable: bool
    message: str
    rawResponse: str
    ruleEventType: str
    ruleID: str


class RuleEventListMatch(TypedDict):
    id: str


class TagRequired(TypedDict):
    merchantID: str
    name: str


class Tag(TagRequired, total=False):
    colourHex: str
    description: str
    id: str


class TagListMatch(TypedDict):
    merchant_id: str


class TagCreateDataRequired(TypedDict):
    merchant_id: str
    merchantID: str
    name: str


class TagCreateData(TagCreateDataRequired, total=False):
    colourHex: str
    description: str
    id: str


class Token(TypedDict):
    pass


class TokenCreateData(TypedDict):
    id: str


class TokenRemoveMatch(TypedDict):
    id: str


class TransactionRequired(TypedDict):
    grossAmount: dict
    payeeDetails: dict
    payerDetails: dict
    transactionAmount: dict


class Transaction(TransactionRequired, total=False):
    accountID: str
    accountName: str
    accountSequenceNumber: int
    addressDetails: dict
    amount: float
    amountMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bookingDateTime: str
    chargeDetails: dict
    content: list
    counterparty: dict
    counterpartySummary: str
    currency: str
    currencyExchange: dict
    date: str
    description: str
    enrichment: dict
    fxAmount: float
    fxCurrency: str
    fxRate: float
    id: str
    inserted: str
    isoBankTransactionCode: dict
    merchant: dict
    merchantID: str
    pageNumber: int
    pageSize: int
    paymentRequestCustomFields: dict
    paymentRequestID: str
    payoutID: str
    proprietaryBankTransactionCode: dict
    rawReference: str
    reference: str
    ruleID: str
    statementReferences: list
    status: str
    supplementaryData: Any
    tags: list
    theirReference: str
    totalPages: int
    totalSize: int
    transactionDate: str
    transactionInformation: list
    transactionMutability: str
    type: str
    valueDateTime: str
    virtualIBAN: str
    yourReference: str


class TransactionLoadMatch(TypedDict):
    id: str


class TransactionListMatch(TypedDict):
    account_id: str
    id: str


class TransactionCreateDataRequired(TypedDict):
    id: str
    grossAmount: dict
    payeeDetails: dict
    payerDetails: dict
    transactionAmount: dict


class TransactionCreateData(TransactionCreateDataRequired, total=False):
    accountID: str
    accountName: str
    accountSequenceNumber: int
    addressDetails: dict
    amount: float
    amountMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bookingDateTime: str
    chargeDetails: dict
    content: list
    counterparty: dict
    counterpartySummary: str
    currency: str
    currencyExchange: dict
    date: str
    description: str
    enrichment: dict
    fxAmount: float
    fxCurrency: str
    fxRate: float
    inserted: str
    isoBankTransactionCode: dict
    merchant: dict
    merchantID: str
    pageNumber: int
    pageSize: int
    paymentRequestCustomFields: dict
    paymentRequestID: str
    payoutID: str
    proprietaryBankTransactionCode: dict
    rawReference: str
    reference: str
    ruleID: str
    statementReferences: list
    status: str
    supplementaryData: Any
    tags: list
    theirReference: str
    totalPages: int
    totalSize: int
    transactionDate: str
    transactionInformation: list
    transactionMutability: str
    type: str
    valueDateTime: str
    virtualIBAN: str
    yourReference: str


class TransactionRemoveMatch(TypedDict):
    id: str


class UserRequired(TypedDict):
    emailAddress: str
    firstName: str
    lastName: str


class User(UserRequired, total=False):
    clientSessionTimeouts: list
    id: str
    passkeyAdded: bool
    permissions: dict
    profile: str
    rolesWithScope: list
    twoFactorEnabled: bool
    userInviteID: str


class UserListMatch(TypedDict, total=False):
    clientSessionTimeouts: list
    emailAddress: str
    firstName: str
    id: str
    lastName: str
    passkeyAdded: bool
    permissions: dict
    profile: str
    rolesWithScope: list
    twoFactorEnabled: bool
    userInviteID: str


class UserUpdateDataRequired(TypedDict):
    id: str


class UserUpdateData(UserUpdateDataRequired, total=False):
    clientSessionTimeouts: list
    emailAddress: str
    firstName: str
    lastName: str
    passkeyAdded: bool
    permissions: dict
    profile: str
    rolesWithScope: list
    twoFactorEnabled: bool
    userInviteID: str


class UserInviteRequired(TypedDict):
    user: dict


class UserInvite(UserInviteRequired, total=False):
    authorisationStatus: dict
    failedUserInvites: dict
    id: str
    initialRoleID: str
    inviteeEmailAddress: str
    inviteeFirstName: str
    inviteeLastName: str
    inviterEmailAddress: str
    inviterFirstName: str
    inviterLastName: str
    isAuthorised: bool
    isInviteeRegistered: bool
    lastInvited: str
    merchantID: str
    merchantName: str
    message: str
    registrationUrl: str
    sendInviteEmail: bool
    status: str
    userID: str
    userInvites: list


class UserInviteLoadMatch(TypedDict):
    id: str


class UserInviteListMatch(TypedDict):
    merchant_id: str


class UserInviteCreateDataRequired(TypedDict):
    id: str
    user: dict


class UserInviteCreateData(UserInviteCreateDataRequired, total=False):
    authorisationStatus: dict
    failedUserInvites: dict
    initialRoleID: str
    inviteeEmailAddress: str
    inviteeFirstName: str
    inviteeLastName: str
    inviterEmailAddress: str
    inviterFirstName: str
    inviterLastName: str
    isAuthorised: bool
    isInviteeRegistered: bool
    lastInvited: str
    merchantID: str
    merchantName: str
    message: str
    registrationUrl: str
    sendInviteEmail: bool
    status: str
    userID: str
    userInvites: list


class UserInviteUpdateDataRequired(TypedDict):
    id: str


class UserInviteUpdateData(UserInviteUpdateDataRequired, total=False):
    authorisationStatus: dict
    failedUserInvites: dict
    initialRoleID: str
    inviteeEmailAddress: str
    inviteeFirstName: str
    inviteeLastName: str
    inviterEmailAddress: str
    inviterFirstName: str
    inviterLastName: str
    isAuthorised: bool
    isInviteeRegistered: bool
    lastInvited: str
    merchantID: str
    merchantName: str
    message: str
    registrationUrl: str
    sendInviteEmail: bool
    status: str
    user: dict
    userID: str
    userInvites: list


class UserInviteRemoveMatch(TypedDict):
    id: str


class VirtualRequired(TypedDict):
    createdBy: dict
    identifier: dict
    name: str


class Virtual(VirtualRequired, total=False):
    accountName: str
    accountSupplierName: str
    availableBalance: float
    availableBalanceMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bankName: str
    consentID: str
    createdByDisplayName: str
    currency: str
    defaultPaymentRail: str
    displayName: str
    expiryDate: str
    externalAccountIcon: str
    id: str
    inserted: str
    isArchived: bool
    isConnectedAccount: bool
    isDefault: bool
    isTrustAccount: bool
    isVirtual: bool
    lastTransaction: dict
    lastUpdated: str
    merchantID: str
    merchantName: str
    physicalAccountID: str
    rules: list
    submittedPayoutsBalance: float
    submittedPayoutsBalanceMinorUnits: int
    summary: str
    supplierSepaInstantStatus: str
    xeroBankFeedConnectionStatus: str
    xeroBankFeedLastSyncedAt: str
    xeroBankFeedSyncLastFailedAt: str
    xeroBankFeedSyncLastFailureReason: str
    xeroBankFeedSyncStatus: str
    xeroUnsynchronisedTransactionsCount: int


class VirtualCreateDataRequired(TypedDict):
    account_id: str
    createdBy: dict
    identifier: dict
    name: str


class VirtualCreateData(VirtualCreateDataRequired, total=False):
    accountName: str
    accountSupplierName: str
    availableBalance: float
    availableBalanceMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bankName: str
    consentID: str
    createdByDisplayName: str
    currency: str
    defaultPaymentRail: str
    displayName: str
    expiryDate: str
    externalAccountIcon: str
    id: str
    inserted: str
    isArchived: bool
    isConnectedAccount: bool
    isDefault: bool
    isTrustAccount: bool
    isVirtual: bool
    lastTransaction: dict
    lastUpdated: str
    merchantID: str
    merchantName: str
    physicalAccountID: str
    rules: list
    submittedPayoutsBalance: float
    submittedPayoutsBalanceMinorUnits: int
    summary: str
    supplierSepaInstantStatus: str
    xeroBankFeedConnectionStatus: str
    xeroBankFeedLastSyncedAt: str
    xeroBankFeedSyncLastFailedAt: str
    xeroBankFeedSyncLastFailureReason: str
    xeroBankFeedSyncStatus: str
    xeroUnsynchronisedTransactionsCount: int


class VirtualUpdateDataRequired(TypedDict):
    account_id: str
    id: str


class VirtualUpdateData(VirtualUpdateDataRequired, total=False):
    accountName: str
    accountSupplierName: str
    availableBalance: float
    availableBalanceMinorUnits: int
    balance: float
    balanceMinorUnits: int
    bankName: str
    consentID: str
    createdBy: dict
    createdByDisplayName: str
    currency: str
    defaultPaymentRail: str
    displayName: str
    expiryDate: str
    externalAccountIcon: str
    identifier: dict
    inserted: str
    isArchived: bool
    isConnectedAccount: bool
    isDefault: bool
    isTrustAccount: bool
    isVirtual: bool
    lastTransaction: dict
    lastUpdated: str
    merchantID: str
    merchantName: str
    name: str
    physicalAccountID: str
    rules: list
    submittedPayoutsBalance: float
    submittedPayoutsBalanceMinorUnits: int
    summary: str
    supplierSepaInstantStatus: str
    xeroBankFeedConnectionStatus: str
    xeroBankFeedLastSyncedAt: str
    xeroBankFeedSyncLastFailedAt: str
    xeroBankFeedSyncLastFailureReason: str
    xeroBankFeedSyncStatus: str
    xeroUnsynchronisedTransactionsCount: int


class Webhook(TypedDict, total=False):
    destinationUrl: str
    emailAddress: str
    failedNotificationEmailAddress: str
    id: str
    isActive: bool
    merchantID: str
    notificationMethod: str
    resourceTypes: list
    retry: bool
    secret: str
    version: int


class WebhookLoadMatchRequired(TypedDict):
    id: str


class WebhookLoadMatch(WebhookLoadMatchRequired, total=False):
    merchant_id: str


class WebhookListMatch(TypedDict):
    merchant_id: str


class WebhookCreateData(TypedDict, total=False):
    destinationUrl: str
    emailAddress: str
    failedNotificationEmailAddress: str
    id: str
    isActive: bool
    merchantID: str
    notificationMethod: str
    resourceTypes: list
    retry: bool
    secret: str
    version: int


class WebhookUpdateDataRequired(TypedDict):
    id: str


class WebhookUpdateData(WebhookUpdateDataRequired, total=False):
    destinationUrl: str
    emailAddress: str
    failedNotificationEmailAddress: str
    isActive: bool
    merchantID: str
    notificationMethod: str
    resourceTypes: list
    retry: bool
    secret: str
    version: int


class WebhookRemoveMatch(TypedDict):
    id: str
