// Typed models for the Nofrixion SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Account {
  accountBalances?: any[]
  accountID?: string
  accountIdentifications?: any[]
  accountName?: string
  accountNames?: any[]
  accountSupplierName?: string
  accountType?: string
  availableBalance?: number
  availableBalanceMinorUnits?: number
  balance?: number
  balanceMinorUnits?: number
  bankName?: string
  consentID?: string
  consolidatedAccountInformation?: Record<string, any>
  createdBy: Record<string, any>
  createdByDisplayName?: string
  currency?: string
  defaultPaymentRail?: string
  description?: string
  details?: string
  displayName?: string
  expiryDate?: string
  externalAccountIcon?: string
  format?: string
  fromDate?: string
  id?: string
  identifier: Record<string, any>
  inserted?: string
  isArchived?: boolean
  isConnectedAccount?: boolean
  isDefault?: boolean
  isTrustAccount?: boolean
  isVirtual?: boolean
  lastTransaction?: Record<string, any>
  lastUpdated?: string
  merchantID?: string
  merchantName?: string
  nickname?: string
  physicalAccountID?: string
  roleIDs?: any[]
  rules?: any[]
  submittedPayoutsBalance?: number
  submittedPayoutsBalanceMinorUnits?: number
  summary?: string
  supplierPhysicalAccountID?: string
  supplierSepaInstantStatus?: string
  toDate?: string
  type?: string
  usageType?: string
  xeroBankFeedConnectionStatus?: string
  xeroBankFeedLastSyncedAt?: string
  xeroBankFeedSyncLastFailedAt?: string
  xeroBankFeedSyncLastFailureReason?: string
  xeroBankFeedSyncStatus?: string
  xeroUnsynchronisedTransactionsCount?: number
}

export interface AccountLoadMatch {
  account_id?: string
  id: string
  merchant_id?: string

  // Selects a custom action instead of the plain load:
  //   'export' | 'statement' | 'transaction_export'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface AccountListMatch {
  merchant_id?: string
}

export interface AccountCreateData {
  account_id?: string
  currency?: string
  accountBalances?: any[]
  accountID?: string
  accountIdentifications?: any[]
  accountName?: string
  accountNames?: any[]
  accountSupplierName?: string
  accountType?: string
  availableBalance?: number
  availableBalanceMinorUnits?: number
  balance?: number
  balanceMinorUnits?: number
  bankName?: string
  consentID?: string
  consolidatedAccountInformation?: Record<string, any>
  createdBy: Record<string, any>
  createdByDisplayName?: string
  defaultPaymentRail?: string
  description?: string
  details?: string
  displayName?: string
  expiryDate?: string
  externalAccountIcon?: string
  format?: string
  fromDate?: string
  id?: string
  identifier: Record<string, any>
  inserted?: string
  isArchived?: boolean
  isConnectedAccount?: boolean
  isDefault?: boolean
  isTrustAccount?: boolean
  isVirtual?: boolean
  lastTransaction?: Record<string, any>
  lastUpdated?: string
  merchantID?: string
  merchantName?: string
  nickname?: string
  physicalAccountID?: string
  roleIDs?: any[]
  rules?: any[]
  submittedPayoutsBalance?: number
  submittedPayoutsBalanceMinorUnits?: number
  summary?: string
  supplierPhysicalAccountID?: string
  supplierSepaInstantStatus?: string
  toDate?: string
  type?: string
  usageType?: string
  xeroBankFeedConnectionStatus?: string
  xeroBankFeedLastSyncedAt?: string
  xeroBankFeedSyncLastFailedAt?: string
  xeroBankFeedSyncLastFailureReason?: string
  xeroBankFeedSyncStatus?: string
  xeroUnsynchronisedTransactionsCount?: number

  // Selects a custom action instead of the plain create:
  //   'statement'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface AccountUpdateData {
  account_id?: string
  amount?: number
  id?: string
  accountBalances?: any[]
  accountID?: string
  accountIdentifications?: any[]
  accountName?: string
  accountNames?: any[]
  accountSupplierName?: string
  accountType?: string
  availableBalance?: number
  availableBalanceMinorUnits?: number
  balance?: number
  balanceMinorUnits?: number
  bankName?: string
  consentID?: string
  consolidatedAccountInformation?: Record<string, any>
  createdBy?: Record<string, any>
  createdByDisplayName?: string
  currency?: string
  defaultPaymentRail?: string
  description?: string
  details?: string
  displayName?: string
  expiryDate?: string
  externalAccountIcon?: string
  format?: string
  fromDate?: string
  identifier?: Record<string, any>
  inserted?: string
  isArchived?: boolean
  isConnectedAccount?: boolean
  isDefault?: boolean
  isTrustAccount?: boolean
  isVirtual?: boolean
  lastTransaction?: Record<string, any>
  lastUpdated?: string
  merchantID?: string
  merchantName?: string
  nickname?: string
  physicalAccountID?: string
  roleIDs?: any[]
  rules?: any[]
  submittedPayoutsBalance?: number
  submittedPayoutsBalanceMinorUnits?: number
  summary?: string
  supplierPhysicalAccountID?: string
  supplierSepaInstantStatus?: string
  toDate?: string
  type?: string
  usageType?: string
  xeroBankFeedConnectionStatus?: string
  xeroBankFeedLastSyncedAt?: string
  xeroBankFeedSyncLastFailedAt?: string
  xeroBankFeedSyncLastFailureReason?: string
  xeroBankFeedSyncStatus?: string
  xeroUnsynchronisedTransactionsCount?: number
}

export interface AccountRemoveMatch {
  id: string

  // Selects a custom action instead of the plain remove:
  //   'statement'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface Batch {
  approveUrl?: string
  id?: string
  payouts?: any[]
}

export interface BatchLoadMatch {
  id: string
}

export interface BatchCreateData {
  approveUrl?: string
  id?: string
  payouts?: any[]
}

export interface Beneficiary {
  approvalCallbackUrl?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  beneficiaries?: any[]
  beneficiaryEvents?: any[]
  canAuthorise?: boolean
  canUpdate?: boolean
  createdBy: Record<string, any>
  createdByEmailAddress?: string
  currency: string
  destination?: Record<string, any>
  failedBeneficiaries?: Record<string, any>
  hasCurrentUserAuthorised?: boolean
  id?: string
  inserted?: string
  isEnabled?: boolean
  lastAuthorised?: string
  lastUpdated?: string
  merchantID?: string
  name: string
  nonce?: string
  sourceAccountIDs?: any[]
  sourceAccounts?: any[]
  theirReference?: string
}

export interface BeneficiaryLoadMatch {
  id: string
  merchant_id?: string

  // Selects a custom action instead of the plain load:
  //   'export'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface BeneficiaryListMatch {
  merchant_id?: string
}

export interface BeneficiaryCreateData {
  id?: string
  approvalCallbackUrl?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  beneficiaries?: any[]
  beneficiaryEvents?: any[]
  canAuthorise?: boolean
  canUpdate?: boolean
  createdBy: Record<string, any>
  createdByEmailAddress?: string
  currency: string
  destination?: Record<string, any>
  failedBeneficiaries?: Record<string, any>
  hasCurrentUserAuthorised?: boolean
  inserted?: string
  isEnabled?: boolean
  lastAuthorised?: string
  lastUpdated?: string
  merchantID?: string
  name: string
  nonce?: string
  sourceAccountIDs?: any[]
  sourceAccounts?: any[]
  theirReference?: string

  // Selects a custom action instead of the plain create:
  //   'batchcreate'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface BeneficiaryUpdateData {
  id: string
  approvalCallbackUrl?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  beneficiaries?: any[]
  beneficiaryEvents?: any[]
  canAuthorise?: boolean
  canUpdate?: boolean
  createdBy?: Record<string, any>
  createdByEmailAddress?: string
  currency?: string
  destination?: Record<string, any>
  failedBeneficiaries?: Record<string, any>
  hasCurrentUserAuthorised?: boolean
  inserted?: string
  isEnabled?: boolean
  lastAuthorised?: string
  lastUpdated?: string
  merchantID?: string
  name?: string
  nonce?: string
  sourceAccountIDs?: any[]
  sourceAccounts?: any[]
  theirReference?: string
}

export interface BeneficiaryRemoveMatch {
  id: string
}

export interface BeneficiaryGroup {
  groupMembers?: any[]
  groupName: string
  id?: string
  inserted?: string
  lastUpdated?: string
  merchantID: string
}

export interface BeneficiaryGroupListMatch {
  merchant_id: string
}

export interface Card {
  authorizedAmount?: string
  currencyCode?: string
  isPayerAuthenticationRequired?: boolean
  isSoftDecline?: boolean
  payerAuthenticationAccessToken?: string
  payerAuthenticationMerchantData?: string
  payerAuthenticationUrl?: string
  payerAuthenticationWindowHeight?: number
  payerAuthenticationWindowWidth?: number
  paymentRequestCallbackUrl?: string
  paymentRequestID?: string
  requestID?: string
  responseCode?: string
  responseType?: string
  status?: string
  threeDSRedirectUrl?: string
  transactionID?: string
}

export interface CardCreateData {
  paymentrequest_id: string
  authorizedAmount?: string
  currencyCode?: string
  isPayerAuthenticationRequired?: boolean
  isSoftDecline?: boolean
  payerAuthenticationAccessToken?: string
  payerAuthenticationMerchantData?: string
  payerAuthenticationUrl?: string
  payerAuthenticationWindowHeight?: number
  payerAuthenticationWindowWidth?: number
  paymentRequestCallbackUrl?: string
  paymentRequestID?: string
  requestID?: string
  responseCode?: string
  responseType?: string
  status?: string
  threeDSRedirectUrl?: string
  transactionID?: string
}

export interface CardCustomerToken {
  cardType?: string
  customerEmailAddress?: string
  expiryMonth?: string
  expiryYear?: string
  id?: string
  inserted?: string
  lastFourDigits?: string
  lastUpdated?: string
  maskedCardNumber?: string
  merchantID?: string
  paymentRequestID?: string
}

export interface CardCustomerTokenLoadMatch {
  customer_email_address: string
}

export interface CardCustomerTokenListMatch {
  customer_email_address: string
  merchant_id: string
}

export interface CardCustomerTokenRemoveMatch {
  customer_email_address?: string
  merchant_id?: string
  id?: string
}

export interface CardPayment {
  authorizedAmount?: string
  currencyCode?: string
  isPayerAuthenticationRequired?: boolean
  isSoftDecline?: boolean
  payerAuthenticationAccessToken?: string
  payerAuthenticationMerchantData?: string
  payerAuthenticationUrl?: string
  payerAuthenticationWindowHeight?: number
  payerAuthenticationWindowWidth?: number
  paymentRequestCallbackUrl?: string
  paymentRequestID?: string
  requestID?: string
  responseCode?: string
  responseType?: string
  status?: string
  threeDSRedirectUrl?: string
  transactionID?: string
}

export interface CardPaymentCreateData {
  partial_refund_amount?: number
  paymentrequest_id: string
  authorizedAmount?: string
  currencyCode?: string
  isPayerAuthenticationRequired?: boolean
  isSoftDecline?: boolean
  payerAuthenticationAccessToken?: string
  payerAuthenticationMerchantData?: string
  payerAuthenticationUrl?: string
  payerAuthenticationWindowHeight?: number
  payerAuthenticationWindowWidth?: number
  paymentRequestCallbackUrl?: string
  paymentRequestID?: string
  requestID?: string
  responseCode?: string
  responseType?: string
  status?: string
  threeDSRedirectUrl?: string
  transactionID?: string
}

export interface CardPublicKey {
  jwt?: string
}

export interface CardPublicKeyLoadMatch {
  paymentrequest_id: string
}

export interface Consent {
  authorisationUrl?: string
  callbackUrl?: string
  consentID?: string
  emailAddress?: string
  expiryDate?: string
  failureCallbackUrl?: string
  id?: string
  inserted?: string
  institutionID?: string
  isConnectedAccounts?: boolean
  isEnabled?: boolean
  merchantID?: string
  provider?: string
  successWebHookUrl?: string
}

export interface ConsentLoadMatch {
  id: string
}

export interface ConsentListMatch {
  email: string
  merchant_id: string
}

export interface ConsentCreateData {
  authorisationUrl?: string
  callbackUrl?: string
  consentID?: string
  emailAddress?: string
  expiryDate?: string
  failureCallbackUrl?: string
  id?: string
  inserted?: string
  institutionID?: string
  isConnectedAccounts?: boolean
  isEnabled?: boolean
  merchantID?: string
  provider?: string
  successWebHookUrl?: string
}

export interface ConsentUpdateData {
  id: string
  authorisationUrl?: string
  callbackUrl?: string
  consentID?: string
  emailAddress?: string
  expiryDate?: string
  failureCallbackUrl?: string
  inserted?: string
  institutionID?: string
  isConnectedAccounts?: boolean
  isEnabled?: boolean
  merchantID?: string
  provider?: string
  successWebHookUrl?: string
}

export interface ConsentRemoveMatch {
  id: string
}

export interface Currency {
  code?: string
  decimals?: number
  isFiat?: boolean
  iso4217AlphaCode?: string
  iso4217NumericCode?: string
  symbol?: string
}

export interface CurrencyListMatch {
  code?: string
  decimals?: number
  isFiat?: boolean
  iso4217AlphaCode?: string
  iso4217NumericCode?: string
  symbol?: string
}

export interface DirectDebitBatchSubmit {
  failedSubmissions?: Record<string, any>
  successfulSubmissions?: any[]
}

export interface DirectDebitBatchSubmitCreateData {
  failedSubmissions?: Record<string, any>
  successfulSubmissions?: any[]
}

export interface FxRate {
  destinationCurrency?: string
  exchangeRate?: number
  expiryTime?: string
  quoteID?: string
  sourceCurrency?: string
}

export interface FxRateLoadMatch {
  destination: string
  source: string
  valid_for_minute: number
}

export interface FxRateListMatch {
  destination: string
  source: string
}

export interface IPayment {
  paymentRequestID?: string
  responseType?: string
}

export interface IPaymentCreateData {
  paymentRequestID?: string
  responseType?: string
}

export interface Mandate {
  accountNumber?: string
  addressLine1: string
  addressLine2?: string
  approvedAt?: string
  city: string
  countryCode: string
  currency?: string
  customerAccountNumber?: string
  customerCity?: string
  customerCountryCode?: string
  customerCountryName?: string
  customerEmailAddress?: string
  customerFirstName?: string
  customerIban?: string
  customerLastName?: string
  customerSortCode?: string
  emailAddress: string
  firstName: string
  iban?: string
  id?: string
  inserted?: string
  isRecurring?: boolean
  lastName: string
  lastUpdated?: string
  merchantID?: string
  postalCode: string
  reference?: string
  sortCode?: string
  status?: string
  supplierBankAccountID?: string
  supplierCustomerID?: string
  supplierMandateID?: string
  supplierName?: string
  supplierStatus?: string
}

export interface MandateLoadMatch {
  id: string
}

export interface MandateCreateData {
  accountNumber?: string
  addressLine1: string
  addressLine2?: string
  approvedAt?: string
  city: string
  countryCode: string
  currency?: string
  customerAccountNumber?: string
  customerCity?: string
  customerCountryCode?: string
  customerCountryName?: string
  customerEmailAddress?: string
  customerFirstName?: string
  customerIban?: string
  customerLastName?: string
  customerSortCode?: string
  emailAddress: string
  firstName: string
  iban?: string
  id?: string
  inserted?: string
  isRecurring?: boolean
  lastName: string
  lastUpdated?: string
  merchantID?: string
  postalCode: string
  reference?: string
  sortCode?: string
  status?: string
  supplierBankAccountID?: string
  supplierCustomerID?: string
  supplierMandateID?: string
  supplierName?: string
  supplierStatus?: string
}

export interface Merchant {
  accountCurrencies?: any[]
  canHaveTrustAccounts?: boolean
  cardPaymentProcessor?: string
  companyID?: string
  displayQrOnHostedPay?: boolean
  hostedPayVersion?: number
  id?: string
  inserted?: string
  isBlocked?: boolean
  isExited?: boolean
  isSuspended?: boolean
  jurisdiction?: string
  logoUrlPng?: string
  logoUrlSvg?: string
  merchantCategoryCode?: string
  name?: string
  notes?: string
  parentMerchant?: Record<string, any>
  paymentAccountLimit?: number
  paymentAccounts?: any[]
  reason?: string
  shortName?: string
  supportedPaymentMethodsList?: any[]
  suspensionReason?: string
  tags?: any[]
  timeZoneId?: string
  tradingName?: string
  webHookLimit?: number
  yourRoleName?: string
}

export interface MerchantLoadMatch {
  id: string

  // Selects a custom action instead of the plain load:
  //   'beneficiary_export' | 'payout_export'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface MerchantListMatch {
  accountCurrencies?: any[]
  canHaveTrustAccounts?: boolean
  cardPaymentProcessor?: string
  companyID?: string
  displayQrOnHostedPay?: boolean
  hostedPayVersion?: number
  id?: string
  inserted?: string
  isBlocked?: boolean
  isExited?: boolean
  isSuspended?: boolean
  jurisdiction?: string
  logoUrlPng?: string
  logoUrlSvg?: string
  merchantCategoryCode?: string
  name?: string
  notes?: string
  parentMerchant?: Record<string, any>
  paymentAccountLimit?: number
  paymentAccounts?: any[]
  reason?: string
  shortName?: string
  supportedPaymentMethodsList?: any[]
  suspensionReason?: string
  tags?: any[]
  timeZoneId?: string
  tradingName?: string
  webHookLimit?: number
  yourRoleName?: string

  // Selects a custom action instead of the plain list:
  //   'childmerchant' | 'paged'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface MerchantUpdateData {
  id: string
  accountCurrencies?: any[]
  canHaveTrustAccounts?: boolean
  cardPaymentProcessor?: string
  companyID?: string
  displayQrOnHostedPay?: boolean
  hostedPayVersion?: number
  inserted?: string
  isBlocked?: boolean
  isExited?: boolean
  isSuspended?: boolean
  jurisdiction?: string
  logoUrlPng?: string
  logoUrlSvg?: string
  merchantCategoryCode?: string
  name?: string
  notes?: string
  parentMerchant?: Record<string, any>
  paymentAccountLimit?: number
  paymentAccounts?: any[]
  reason?: string
  shortName?: string
  supportedPaymentMethodsList?: any[]
  suspensionReason?: string
  tags?: any[]
  timeZoneId?: string
  tradingName?: string
  webHookLimit?: number
  yourRoleName?: string

  // Selects a custom action instead of the plain update:
  //   'suspend'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface MerchantRemoveMatch {
  id?: string
  user_id?: string
  merchant_id?: string
  tag_id?: string
}

export interface MerchantAuthorisationSetting {
  amountLower?: number
  amountUpper?: number
  authorisationType?: string
  beneficiariesOnly?: boolean
  id?: string
  inserted?: string
  lastEditorCantAuthorise?: boolean
  lastUpdated?: string
  merchantID?: string
  numberOfAuthorisers?: number
  roleSettings?: any[]
}

export interface MerchantAuthorisationSettingListMatch {
  merchant_id: string
}

export interface MerchantDirectDebitMandatePage {
  approvedAt?: string
  currency?: string
  customerAccountNumber?: string
  customerCity?: string
  customerCountryCode?: string
  customerCountryName?: string
  customerEmailAddress?: string
  customerFirstName?: string
  customerIban?: string
  customerLastName?: string
  customerSortCode?: string
  id?: string
  inserted?: string
  isRecurring?: boolean
  lastUpdated?: string
  merchantID?: string
  reference?: string
  status?: string
  supplierBankAccountID?: string
  supplierCustomerID?: string
  supplierMandateID?: string
  supplierName?: string
  supplierStatus?: string
}

export interface MerchantDirectDebitMandatePageListMatch {
  approvedAt?: string
  currency?: string
  customerAccountNumber?: string
  customerCity?: string
  customerCountryCode?: string
  customerCountryName?: string
  customerEmailAddress?: string
  customerFirstName?: string
  customerIban?: string
  customerLastName?: string
  customerSortCode?: string
  id?: string
  inserted?: string
  isRecurring?: boolean
  lastUpdated?: string
  merchantID?: string
  reference?: string
  status?: string
  supplierBankAccountID?: string
  supplierCustomerID?: string
  supplierMandateID?: string
  supplierName?: string
  supplierStatus?: string
}

export interface MerchantPayByBankSetting {
  bankCountryCodes?: any[]
  bankID?: string
  bankName?: string
  businessInstitutionID?: string
  currency?: string
  logo?: string
  message?: string
  messageImageUrl?: string
  order?: number
  personalInstitutionID?: string
  processor?: string
  warningHeading?: string
  warningMessage?: string
}

export interface MerchantPayByBankSettingListMatch {
  merchant_id: string
}

export interface MerchantPaymentRequestTemplate {
  bankPaymentOptions?: Record<string, any>
  cardPaymentAddressOptions?: Record<string, any>
  cardPaymentCaptureOptions?: Record<string, any>
  customFields?: any[]
  defaultFields?: any[]
  description: string
  id?: string
  inserted?: string
  lastUpdated?: string
  merchantID?: string
  name: string
  notificationOptions?: Record<string, any>
  paymentMethods?: Record<string, any>
  paymentTerms?: Record<string, any>
  priorityBankOptions?: Record<string, any>
  template: Record<string, any>
}

export interface MerchantPaymentRequestTemplateLoadMatch {
  id: string
  paymentrequest_id: string
}

export interface MerchantPaymentRequestTemplateListMatch {
  merchant_id: string
}

export interface MerchantPaymentRequestTemplateUpdateData {
  id: string
  paymentrequest_id: string
  bankPaymentOptions?: Record<string, any>
  cardPaymentAddressOptions?: Record<string, any>
  cardPaymentCaptureOptions?: Record<string, any>
  customFields?: any[]
  defaultFields?: any[]
  description?: string
  inserted?: string
  lastUpdated?: string
  merchantID?: string
  name?: string
  notificationOptions?: Record<string, any>
  paymentMethods?: Record<string, any>
  paymentTerms?: Record<string, any>
  priorityBankOptions?: Record<string, any>
  template?: Record<string, any>
}

export interface MerchantPaymentRequestTemplateRemoveMatch {
  id: string
  paymentrequest_id: string
}

export interface MerchantToken {
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  canAuthorise?: boolean
  description?: string
  expiresAt?: string
  hasCurrentUserAuthorised?: boolean
  hmacAlgorithm?: string
  id?: string
  inserted?: string
  ipAddressWhitelist?: string
  isArchived?: boolean
  isEnabled?: boolean
  lastAuthorised?: string
  lastUpdated?: string
  merchantID?: string
  nonce: string
  permissionTypes?: any[]
  requestSignatureVersion?: number
  sharedSecretAlgorithm?: string
  sharedSecretBase64?: string
  token?: string
}

export interface MerchantTokenLoadMatch {
  id: string
}

export interface MerchantTokenListMatch {
  merchant_id: string
}

export interface MerchantTokenCreateData {
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  canAuthorise?: boolean
  description?: string
  expiresAt?: string
  hasCurrentUserAuthorised?: boolean
  hmacAlgorithm?: string
  id?: string
  inserted?: string
  ipAddressWhitelist?: string
  isArchived?: boolean
  isEnabled?: boolean
  lastAuthorised?: string
  lastUpdated?: string
  merchantID?: string
  nonce: string
  permissionTypes?: any[]
  requestSignatureVersion?: number
  sharedSecretAlgorithm?: string
  sharedSecretBase64?: string
  token?: string
}

export interface MerchantTokenUpdateData {
  id: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  canAuthorise?: boolean
  description?: string
  expiresAt?: string
  hasCurrentUserAuthorised?: boolean
  hmacAlgorithm?: string
  inserted?: string
  ipAddressWhitelist?: string
  isArchived?: boolean
  isEnabled?: boolean
  lastAuthorised?: string
  lastUpdated?: string
  merchantID?: string
  nonce?: string
  permissionTypes?: any[]
  requestSignatureVersion?: number
  sharedSecretAlgorithm?: string
  sharedSecretBase64?: string
  token?: string
}

export interface Metadata {
}

export interface MetadataLoadMatch {

  // Selects a custom action instead of the plain load:
  //   'problem' | 'problemnotification'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface NoFrixionVersion {
  buildVersion?: number
  majorVersion?: number
  minorVersion?: number
  releaseName?: string
}

export interface NoFrixionVersionLoadMatch {
  buildVersion?: number
  majorVersion?: number
  minorVersion?: number
  releaseName?: string
}

export interface OpenBanking {
}

export interface OpenBankingCreateData {
  account_id: string
}

export interface OpenBankingRemoveMatch {
  email?: string
  merchant_id?: string
  account_id?: string
}

export interface Payeeverification {
  accountName: string
  accountNumber?: string
  iban: string
  payeeVerifiedAccountName?: string
  result?: string
  secondaryIdentification?: string
  sortCode?: string
}

export interface PayeeverificationCreateData {
  accountName: string
  accountNumber?: string
  iban: string
  payeeVerifiedAccountName?: string
  result?: string
  secondaryIdentification?: string
  sortCode?: string
}

export interface Payment {
  addresses?: any[]
  amount?: number
  amountPending?: number
  amountReceived?: number
  amountRefunded?: number
  autoSendReceipt?: boolean
  baseOriginUrl?: string
  callbackUrl?: string
  cardAuthorizeOnly?: boolean
  cardCreateToken?: boolean
  cardCreateTokenMode?: string
  cardIgnoreCVN?: boolean
  cardNoPayerAuthentication?: boolean
  cardProcessorMerchantID?: string
  cardStripePaymentIntentID?: string
  cardStripePaymentIntentSecret?: string
  cardTransmitRawDetails?: boolean
  createdByUser: Record<string, any>
  currency?: string
  customFields?: any[]
  customerEmailAddress?: string
  customerID?: string
  customerName?: string
  description?: string
  destinationAccount?: Record<string, any>
  directDebitPayment?: Record<string, any>
  dueDate?: string
  events?: any[]
  failureCallbackUrl?: string
  fieldDisplaySettings?: any[]
  formattedAmount?: string
  hostedPayCheckoutUrl?: string
  id?: string
  ignoreAddressVerification?: boolean
  inserted?: string
  insertedSortable?: string
  isArchived?: boolean
  jwk?: string
  lastUpdated?: string
  lightningInvoice?: string
  lightningInvoiceExpiresAt?: string
  merchantDirectDebitMandateID?: string
  merchantID?: string
  merchantTokenDescription?: string
  notificationEmailAddresses?: string
  notificationRoleIDs?: any[]
  orderID?: string
  partialPaymentMethod?: string
  partialPaymentSteps?: string
  paymentAttempts?: any[]
  paymentMethods?: any[]
  paymentProcessor?: string
  payrunID?: string
  pispAccountID?: string
  priorityBankID?: string
  result?: Record<string, any>
  sandboxSettleDelayInSeconds?: number
  shippingAddress?: Record<string, any>
  shippingAddressCity?: string
  shippingAddressCountryCode?: string
  shippingAddressCounty?: string
  shippingAddressLine1?: string
  shippingAddressLine2?: string
  shippingAddressPostCode?: string
  shippingEmail?: string
  shippingFirstName?: string
  shippingLastName?: string
  shippingPhone?: string
  status?: string
  successWebHookUrl?: string
  tagIds?: any[]
  tags?: any[]
  title?: string
  tokenisedCards?: any[]
  transactions?: any[]
  useHostedPaymentPage?: boolean
}

export interface PaymentLoadMatch {
  id?: string
  order_id?: string
}

export interface PaymentCreateData {
  addresses?: any[]
  amount?: number
  amountPending?: number
  amountReceived?: number
  amountRefunded?: number
  autoSendReceipt?: boolean
  baseOriginUrl?: string
  callbackUrl?: string
  cardAuthorizeOnly?: boolean
  cardCreateToken?: boolean
  cardCreateTokenMode?: string
  cardIgnoreCVN?: boolean
  cardNoPayerAuthentication?: boolean
  cardProcessorMerchantID?: string
  cardStripePaymentIntentID?: string
  cardStripePaymentIntentSecret?: string
  cardTransmitRawDetails?: boolean
  createdByUser: Record<string, any>
  currency?: string
  customFields?: any[]
  customerEmailAddress?: string
  customerID?: string
  customerName?: string
  description?: string
  destinationAccount?: Record<string, any>
  directDebitPayment?: Record<string, any>
  dueDate?: string
  events?: any[]
  failureCallbackUrl?: string
  fieldDisplaySettings?: any[]
  formattedAmount?: string
  hostedPayCheckoutUrl?: string
  id?: string
  ignoreAddressVerification?: boolean
  inserted?: string
  insertedSortable?: string
  isArchived?: boolean
  jwk?: string
  lastUpdated?: string
  lightningInvoice?: string
  lightningInvoiceExpiresAt?: string
  merchantDirectDebitMandateID?: string
  merchantID?: string
  merchantTokenDescription?: string
  notificationEmailAddresses?: string
  notificationRoleIDs?: any[]
  orderID?: string
  partialPaymentMethod?: string
  partialPaymentSteps?: string
  paymentAttempts?: any[]
  paymentMethods?: any[]
  paymentProcessor?: string
  payrunID?: string
  pispAccountID?: string
  priorityBankID?: string
  result?: Record<string, any>
  sandboxSettleDelayInSeconds?: number
  shippingAddress?: Record<string, any>
  shippingAddressCity?: string
  shippingAddressCountryCode?: string
  shippingAddressCounty?: string
  shippingAddressLine1?: string
  shippingAddressLine2?: string
  shippingAddressPostCode?: string
  shippingEmail?: string
  shippingFirstName?: string
  shippingLastName?: string
  shippingPhone?: string
  status?: string
  successWebHookUrl?: string
  tagIds?: any[]
  tags?: any[]
  title?: string
  tokenisedCards?: any[]
  transactions?: any[]
  useHostedPaymentPage?: boolean
}

export interface PaymentUpdateData {
  id: string
  addresses?: any[]
  amount?: number
  amountPending?: number
  amountReceived?: number
  amountRefunded?: number
  autoSendReceipt?: boolean
  baseOriginUrl?: string
  callbackUrl?: string
  cardAuthorizeOnly?: boolean
  cardCreateToken?: boolean
  cardCreateTokenMode?: string
  cardIgnoreCVN?: boolean
  cardNoPayerAuthentication?: boolean
  cardProcessorMerchantID?: string
  cardStripePaymentIntentID?: string
  cardStripePaymentIntentSecret?: string
  cardTransmitRawDetails?: boolean
  createdByUser?: Record<string, any>
  currency?: string
  customFields?: any[]
  customerEmailAddress?: string
  customerID?: string
  customerName?: string
  description?: string
  destinationAccount?: Record<string, any>
  directDebitPayment?: Record<string, any>
  dueDate?: string
  events?: any[]
  failureCallbackUrl?: string
  fieldDisplaySettings?: any[]
  formattedAmount?: string
  hostedPayCheckoutUrl?: string
  ignoreAddressVerification?: boolean
  inserted?: string
  insertedSortable?: string
  isArchived?: boolean
  jwk?: string
  lastUpdated?: string
  lightningInvoice?: string
  lightningInvoiceExpiresAt?: string
  merchantDirectDebitMandateID?: string
  merchantID?: string
  merchantTokenDescription?: string
  notificationEmailAddresses?: string
  notificationRoleIDs?: any[]
  orderID?: string
  partialPaymentMethod?: string
  partialPaymentSteps?: string
  paymentAttempts?: any[]
  paymentMethods?: any[]
  paymentProcessor?: string
  payrunID?: string
  pispAccountID?: string
  priorityBankID?: string
  result?: Record<string, any>
  sandboxSettleDelayInSeconds?: number
  shippingAddress?: Record<string, any>
  shippingAddressCity?: string
  shippingAddressCountryCode?: string
  shippingAddressCounty?: string
  shippingAddressLine1?: string
  shippingAddressLine2?: string
  shippingAddressPostCode?: string
  shippingEmail?: string
  shippingFirstName?: string
  shippingLastName?: string
  shippingPhone?: string
  status?: string
  successWebHookUrl?: string
  tagIds?: any[]
  tags?: any[]
  title?: string
  tokenisedCards?: any[]
  transactions?: any[]
  useHostedPaymentPage?: boolean
}

export interface PaymentAccount {
  accountName?: string
  accountSupplierName?: string
  availableBalance?: number
  availableBalanceMinorUnits?: number
  balance?: number
  balanceMinorUnits?: number
  bankName?: string
  consentID?: string
  createdBy: Record<string, any>
  createdByDisplayName?: string
  currency?: string
  defaultPaymentRail?: string
  displayName?: string
  expiryDate?: string
  externalAccountIcon?: string
  id?: string
  identifier: Record<string, any>
  inserted?: string
  isArchived?: boolean
  isConnectedAccount?: boolean
  isDefault?: boolean
  isTrustAccount?: boolean
  isVirtual?: boolean
  lastTransaction?: Record<string, any>
  lastUpdated?: string
  merchantID?: string
  merchantName?: string
  physicalAccountID?: string
  rules?: any[]
  submittedPayoutsBalance?: number
  submittedPayoutsBalanceMinorUnits?: number
  summary?: string
  supplierSepaInstantStatus?: string
  xeroBankFeedConnectionStatus?: string
  xeroBankFeedLastSyncedAt?: string
  xeroBankFeedSyncLastFailedAt?: string
  xeroBankFeedSyncLastFailureReason?: string
  xeroBankFeedSyncStatus?: string
  xeroUnsynchronisedTransactionsCount?: number
}

export interface PaymentAccountListMatch {
  account_id?: string
}

export interface PaymentAccountMinimal {
  accountName?: string
  availableBalance?: number
  balance?: number
  balanceMinorUnits?: number
  currency?: string
  id?: string
  identifier: Record<string, any>
  isArchived?: boolean
  isConnectedAccount?: boolean
  merchantID?: string
  submittedPayoutsBalance?: number
}

export interface PaymentAccountMinimalListMatch {
  accountName?: string
  availableBalance?: number
  balance?: number
  balanceMinorUnits?: number
  currency?: string
  id?: string
  identifier?: Record<string, any>
  isArchived?: boolean
  isConnectedAccount?: boolean
  merchantID?: string
  submittedPayoutsBalance?: number
}

export interface PaymentInitiation {
  paymentInitiationID?: string
  paymentRequestCallbackUrl?: string
  paymentRequestID?: string
  redirectUrl?: string
  responseType?: string
  specificErrorMessage?: string
}

export interface PaymentInitiationCreateData {
  paymentrequest_id: string
  paymentInitiationID?: string
  paymentRequestCallbackUrl?: string
  paymentRequestID?: string
  redirectUrl?: string
  responseType?: string
  specificErrorMessage?: string
}

export interface PaymentRequest {
  addresses?: any[]
  amount?: number
  amountPending?: number
  amountReceived?: number
  amountRefunded?: number
  autoSendReceipt?: boolean
  baseOriginUrl?: string
  callbackUrl?: string
  cardAuthorizeOnly?: boolean
  cardCreateToken?: boolean
  cardCreateTokenMode?: string
  cardIgnoreCVN?: boolean
  cardProcessorMerchantID?: string
  cardStripePaymentIntentID?: string
  cardStripePaymentIntentSecret?: string
  createdByUser: Record<string, any>
  currency?: string
  customFields?: any[]
  customerEmailAddress?: string
  customerID?: string
  customerName?: string
  description?: string
  destinationAccount?: Record<string, any>
  directDebitPayment?: Record<string, any>
  doSimulateSettlementFailure?: boolean
  dueDate?: string
  errorDescription?: string
  events?: any[]
  failedPaymentRequests?: Record<string, any>
  failureCallbackUrl?: string
  fieldDisplaySettings?: any[]
  formattedAmount?: string
  hostedPayCheckoutUrl?: string
  id?: string
  ignoreAddressVerification?: boolean
  inserted?: string
  insertedSortable?: string
  institution?: string
  isArchived?: boolean
  jwk?: string
  lastUpdated?: string
  lightningInvoice?: string
  lightningInvoiceExpiresAt?: string
  merchantDirectDebitMandateID?: string
  merchantID?: string
  merchantTokenDescription?: string
  notificationEmailAddresses?: string
  notificationRoleIDs?: any[]
  orderID?: string
  partialPaymentMethod?: string
  partialPaymentSteps?: string
  paymentAttempts?: any[]
  paymentInitiationID?: string
  paymentMethods?: any[]
  paymentProcessor?: string
  paymentRequests?: any[]
  payrunID?: string
  pispAccountID?: string
  priorityBankID?: string
  result?: Record<string, any>
  sandboxSettleDelayInSeconds?: number
  shippingAddress?: Record<string, any>
  status?: string
  successWebHookUrl?: string
  tags?: any[]
  title?: string
  tokenisedCards?: any[]
  transactions?: any[]
  useHostedPaymentPage?: boolean
}

export interface PaymentRequestLoadMatch {
  paymentrequest_id?: string
}

export interface PaymentRequestListMatch {
  addresses?: any[]
  amount?: number
  amountPending?: number
  amountReceived?: number
  amountRefunded?: number
  autoSendReceipt?: boolean
  baseOriginUrl?: string
  callbackUrl?: string
  cardAuthorizeOnly?: boolean
  cardCreateToken?: boolean
  cardCreateTokenMode?: string
  cardIgnoreCVN?: boolean
  cardProcessorMerchantID?: string
  cardStripePaymentIntentID?: string
  cardStripePaymentIntentSecret?: string
  createdByUser?: Record<string, any>
  currency?: string
  customFields?: any[]
  customerEmailAddress?: string
  customerID?: string
  customerName?: string
  description?: string
  destinationAccount?: Record<string, any>
  directDebitPayment?: Record<string, any>
  doSimulateSettlementFailure?: boolean
  dueDate?: string
  errorDescription?: string
  events?: any[]
  failedPaymentRequests?: Record<string, any>
  failureCallbackUrl?: string
  fieldDisplaySettings?: any[]
  formattedAmount?: string
  hostedPayCheckoutUrl?: string
  id?: string
  ignoreAddressVerification?: boolean
  inserted?: string
  insertedSortable?: string
  institution?: string
  isArchived?: boolean
  jwk?: string
  lastUpdated?: string
  lightningInvoice?: string
  lightningInvoiceExpiresAt?: string
  merchantDirectDebitMandateID?: string
  merchantID?: string
  merchantTokenDescription?: string
  notificationEmailAddresses?: string
  notificationRoleIDs?: any[]
  orderID?: string
  partialPaymentMethod?: string
  partialPaymentSteps?: string
  paymentAttempts?: any[]
  paymentInitiationID?: string
  paymentMethods?: any[]
  paymentProcessor?: string
  paymentRequests?: any[]
  payrunID?: string
  pispAccountID?: string
  priorityBankID?: string
  result?: Record<string, any>
  sandboxSettleDelayInSeconds?: number
  shippingAddress?: Record<string, any>
  status?: string
  successWebHookUrl?: string
  tags?: any[]
  title?: string
  tokenisedCards?: any[]
  transactions?: any[]
  useHostedPaymentPage?: boolean
}

export interface PaymentRequestCreateData {
  paymentrequest_id?: string
  addresses?: any[]
  amount?: number
  amountPending?: number
  amountReceived?: number
  amountRefunded?: number
  autoSendReceipt?: boolean
  baseOriginUrl?: string
  callbackUrl?: string
  cardAuthorizeOnly?: boolean
  cardCreateToken?: boolean
  cardCreateTokenMode?: string
  cardIgnoreCVN?: boolean
  cardProcessorMerchantID?: string
  cardStripePaymentIntentID?: string
  cardStripePaymentIntentSecret?: string
  createdByUser: Record<string, any>
  currency?: string
  customFields?: any[]
  customerEmailAddress?: string
  customerID?: string
  customerName?: string
  description?: string
  destinationAccount?: Record<string, any>
  directDebitPayment?: Record<string, any>
  doSimulateSettlementFailure?: boolean
  dueDate?: string
  errorDescription?: string
  events?: any[]
  failedPaymentRequests?: Record<string, any>
  failureCallbackUrl?: string
  fieldDisplaySettings?: any[]
  formattedAmount?: string
  hostedPayCheckoutUrl?: string
  id?: string
  ignoreAddressVerification?: boolean
  inserted?: string
  insertedSortable?: string
  institution?: string
  isArchived?: boolean
  jwk?: string
  lastUpdated?: string
  lightningInvoice?: string
  lightningInvoiceExpiresAt?: string
  merchantDirectDebitMandateID?: string
  merchantID?: string
  merchantTokenDescription?: string
  notificationEmailAddresses?: string
  notificationRoleIDs?: any[]
  orderID?: string
  partialPaymentMethod?: string
  partialPaymentSteps?: string
  paymentAttempts?: any[]
  paymentInitiationID?: string
  paymentMethods?: any[]
  paymentProcessor?: string
  paymentRequests?: any[]
  payrunID?: string
  pispAccountID?: string
  priorityBankID?: string
  result?: Record<string, any>
  sandboxSettleDelayInSeconds?: number
  shippingAddress?: Record<string, any>
  status?: string
  successWebHookUrl?: string
  tags?: any[]
  title?: string
  tokenisedCards?: any[]
  transactions?: any[]
  useHostedPaymentPage?: boolean
}

export interface PaymentRequestUpdateData {
  paymentrequest_id: string
  addresses?: any[]
  amount?: number
  amountPending?: number
  amountReceived?: number
  amountRefunded?: number
  autoSendReceipt?: boolean
  baseOriginUrl?: string
  callbackUrl?: string
  cardAuthorizeOnly?: boolean
  cardCreateToken?: boolean
  cardCreateTokenMode?: string
  cardIgnoreCVN?: boolean
  cardProcessorMerchantID?: string
  cardStripePaymentIntentID?: string
  cardStripePaymentIntentSecret?: string
  createdByUser?: Record<string, any>
  currency?: string
  customFields?: any[]
  customerEmailAddress?: string
  customerID?: string
  customerName?: string
  description?: string
  destinationAccount?: Record<string, any>
  directDebitPayment?: Record<string, any>
  doSimulateSettlementFailure?: boolean
  dueDate?: string
  errorDescription?: string
  events?: any[]
  failedPaymentRequests?: Record<string, any>
  failureCallbackUrl?: string
  fieldDisplaySettings?: any[]
  formattedAmount?: string
  hostedPayCheckoutUrl?: string
  id?: string
  ignoreAddressVerification?: boolean
  inserted?: string
  insertedSortable?: string
  institution?: string
  isArchived?: boolean
  jwk?: string
  lastUpdated?: string
  lightningInvoice?: string
  lightningInvoiceExpiresAt?: string
  merchantDirectDebitMandateID?: string
  merchantID?: string
  merchantTokenDescription?: string
  notificationEmailAddresses?: string
  notificationRoleIDs?: any[]
  orderID?: string
  partialPaymentMethod?: string
  partialPaymentSteps?: string
  paymentAttempts?: any[]
  paymentInitiationID?: string
  paymentMethods?: any[]
  paymentProcessor?: string
  paymentRequests?: any[]
  payrunID?: string
  pispAccountID?: string
  priorityBankID?: string
  result?: Record<string, any>
  sandboxSettleDelayInSeconds?: number
  shippingAddress?: Record<string, any>
  status?: string
  successWebHookUrl?: string
  tags?: any[]
  title?: string
  tokenisedCards?: any[]
  transactions?: any[]
  useHostedPaymentPage?: boolean
}

export interface PaymentRequestRemoveMatch {
  id: string
}

export interface PaymentRequestEvent {
  amount: number
  applePayTransactionID?: string
  cardAuthorizationResponseID?: string
  cardExpiryMonth?: number
  cardExpiryYear?: number
  cardIssuer?: string
  cardIssuerCountry?: string
  cardLastFourDigits?: string
  cardRequestID?: string
  cardScheme?: string
  cardTokenCustomerID?: string
  cardTransactionID?: string
  currency?: string
  directDebitPaymentID?: string
  directDebitPaymentReference?: string
  drirectDebitMandateID?: string
  errorMessage?: string
  errorReason?: string
  eventType?: string
  id?: string
  inserted?: string
  lightningInvoice?: string
  lightningRHash?: string
  originUrl?: string
  paymentMethodType?: string
  paymentProcessorName?: string
  paymentRequestID?: string
  pispBankStatus?: string
  pispPaymentInitiationID?: string
  pispPaymentInstitutionName?: string
  pispPaymentServiceProviderID?: string
  pispRedirectUrl?: string
  reconciledTransactionID?: string
  refundPayoutID?: string
  status?: string
  walletName?: string
}

export interface PaymentRequestEventListMatch {
  paymentrequest_id: string
}

export interface PaymentRequestMetric {
}

export interface PaymentRequestMetricLoadMatch {
}

export interface PaymentRequestMinimal {
  amount?: number
  amountPending?: number
  amountReceived?: number
  amountRefunded?: number
  callbackUrl?: string
  cardStripePaymentIntentSecret?: string
  countryCode?: string
  currency?: string
  customFieldsToDisplay?: any[]
  description?: string
  dueDate?: string
  fieldDisplaySettings?: any[]
  googlePayMerchantID?: string
  id?: string
  jwk?: string
  merchantID?: string
  merchantLogoUrlPng?: string
  merchantLogoUrlSvg?: string
  merchantName?: string
  merchantShortName?: string
  partialPaymentMethod?: string
  paymentAttempts?: any[]
  paymentMethodsList?: any[]
  paymentProcessor?: string
  paymentProcessorKey?: string
  pispError?: string
  priorityBankID?: string
  status?: string
  stripeAccountID?: string
  title?: string
}

export interface PaymentRequestMinimalListMatch {
  paymentrequest_id: string
}

export interface PaymentRequestResult {
  amount?: number
  amountPending?: number
  amountReceived?: number
  amountRefunded?: number
  currency?: string
  customerID?: string
  paymentRequestID?: string
  payments?: any[]
  pispAuthorizations?: any[]
  requestedAmount?: number
  result?: string
}

export interface PaymentRequestResultListMatch {
  paymentrequest_id: string
}

export interface Payout {
  accountID?: string
  allowIncomplete?: boolean
  amount?: number
  amountMinorUnits?: number
  approvePayoutUrl?: string
  approverID?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  batchPayoutID?: string
  beneficiary: Record<string, any>
  beneficiaryID?: string
  canAuthorise?: boolean
  canProcess?: boolean
  canUpdate?: boolean
  chargeBearer?: string
  createdBy?: string
  createdByEmailAddress?: string
  currency?: string
  currentUserID?: string
  description?: string
  destination?: Record<string, any>
  documents?: any[]
  events?: any[]
  failedPayouts?: Record<string, any>
  formattedAmount?: string
  formattedFxDestinationAmount?: string
  formattedSchedule?: string
  formattedScheduleDayOnly?: string
  formattedSourceAccountAvailableBalance?: string
  fxDestinationAmount?: number
  fxDestinationAmountMinorUnits?: number
  fxDestinationCurrency?: string
  fxQuoteExpiresAt?: string
  fxQuoteID?: string
  fxRate?: number
  fxUseDestinationAmount?: boolean
  hasCurrentUserAuthorised?: boolean
  id?: string
  inserted?: string
  invoiceID?: string
  isArchived?: boolean
  isFailed?: boolean
  isSettled?: boolean
  isSubmitted?: boolean
  lastUpdated?: string
  merchantID?: string
  merchantTokenDescription?: string
  nonce?: string
  paymentProcessor?: string
  paymentRail?: string
  payouts?: any[]
  payrunID?: string
  payrunName?: string
  reason?: string
  rule?: Record<string, any>
  scheduleDate?: string
  scheduled?: boolean
  sourceAccountAvailableBalance?: number
  sourceAccountAvailableBalanceMinorUnits?: number
  sourceAccountBic?: string
  sourceAccountCurrency?: string
  sourceAccountIban?: string
  sourceAccountIdentifier: Record<string, any>
  sourceAccountName?: string
  sourceAccountNumber?: string
  sourceAccountSortcode?: string
  status?: string
  tagIds?: any[]
  tags?: any[]
  theirReference?: string
  topupPayrunID?: string
  transactedAmount?: number
  transactedFxAmount?: number
  transactedFxRate?: number
  type?: string
  userID?: string
  yourReference?: string
}

export interface PayoutLoadMatch {
  amount?: number
  destination?: string
  source?: string
  id?: string

  // Selects a custom action instead of the plain load:
  //   'export' | 'proof'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface PayoutListMatch {
  account_id?: string
  merchant_id?: string
}

export interface PayoutCreateData {
  id?: string
  accountID?: string
  allowIncomplete?: boolean
  amount?: number
  amountMinorUnits?: number
  approvePayoutUrl?: string
  approverID?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  batchPayoutID?: string
  beneficiary: Record<string, any>
  beneficiaryID?: string
  canAuthorise?: boolean
  canProcess?: boolean
  canUpdate?: boolean
  chargeBearer?: string
  createdBy?: string
  createdByEmailAddress?: string
  currency?: string
  currentUserID?: string
  description?: string
  destination?: Record<string, any>
  documents?: any[]
  events?: any[]
  failedPayouts?: Record<string, any>
  formattedAmount?: string
  formattedFxDestinationAmount?: string
  formattedSchedule?: string
  formattedScheduleDayOnly?: string
  formattedSourceAccountAvailableBalance?: string
  fxDestinationAmount?: number
  fxDestinationAmountMinorUnits?: number
  fxDestinationCurrency?: string
  fxQuoteExpiresAt?: string
  fxQuoteID?: string
  fxRate?: number
  fxUseDestinationAmount?: boolean
  hasCurrentUserAuthorised?: boolean
  inserted?: string
  invoiceID?: string
  isArchived?: boolean
  isFailed?: boolean
  isSettled?: boolean
  isSubmitted?: boolean
  lastUpdated?: string
  merchantID?: string
  merchantTokenDescription?: string
  nonce?: string
  paymentProcessor?: string
  paymentRail?: string
  payouts?: any[]
  payrunID?: string
  payrunName?: string
  reason?: string
  rule?: Record<string, any>
  scheduleDate?: string
  scheduled?: boolean
  sourceAccountAvailableBalance?: number
  sourceAccountAvailableBalanceMinorUnits?: number
  sourceAccountBic?: string
  sourceAccountCurrency?: string
  sourceAccountIban?: string
  sourceAccountIdentifier: Record<string, any>
  sourceAccountName?: string
  sourceAccountNumber?: string
  sourceAccountSortcode?: string
  status?: string
  tagIds?: any[]
  tags?: any[]
  theirReference?: string
  topupPayrunID?: string
  transactedAmount?: number
  transactedFxAmount?: number
  transactedFxRate?: number
  type?: string
  userID?: string
  yourReference?: string

  // Selects a custom action instead of the plain create:
  //   'batchcreate' | 'send' | 'sendbeneficiary'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface PayoutUpdateData {
  id: string
  accountID?: string
  allowIncomplete?: boolean
  amount?: number
  amountMinorUnits?: number
  approvePayoutUrl?: string
  approverID?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  batchPayoutID?: string
  beneficiary?: Record<string, any>
  beneficiaryID?: string
  canAuthorise?: boolean
  canProcess?: boolean
  canUpdate?: boolean
  chargeBearer?: string
  createdBy?: string
  createdByEmailAddress?: string
  currency?: string
  currentUserID?: string
  description?: string
  destination?: Record<string, any>
  documents?: any[]
  events?: any[]
  failedPayouts?: Record<string, any>
  formattedAmount?: string
  formattedFxDestinationAmount?: string
  formattedSchedule?: string
  formattedScheduleDayOnly?: string
  formattedSourceAccountAvailableBalance?: string
  fxDestinationAmount?: number
  fxDestinationAmountMinorUnits?: number
  fxDestinationCurrency?: string
  fxQuoteExpiresAt?: string
  fxQuoteID?: string
  fxRate?: number
  fxUseDestinationAmount?: boolean
  hasCurrentUserAuthorised?: boolean
  inserted?: string
  invoiceID?: string
  isArchived?: boolean
  isFailed?: boolean
  isSettled?: boolean
  isSubmitted?: boolean
  lastUpdated?: string
  merchantID?: string
  merchantTokenDescription?: string
  nonce?: string
  paymentProcessor?: string
  paymentRail?: string
  payouts?: any[]
  payrunID?: string
  payrunName?: string
  reason?: string
  rule?: Record<string, any>
  scheduleDate?: string
  scheduled?: boolean
  sourceAccountAvailableBalance?: number
  sourceAccountAvailableBalanceMinorUnits?: number
  sourceAccountBic?: string
  sourceAccountCurrency?: string
  sourceAccountIban?: string
  sourceAccountIdentifier?: Record<string, any>
  sourceAccountName?: string
  sourceAccountNumber?: string
  sourceAccountSortcode?: string
  status?: string
  tagIds?: any[]
  tags?: any[]
  theirReference?: string
  topupPayrunID?: string
  transactedAmount?: number
  transactedFxAmount?: number
  transactedFxRate?: number
  type?: string
  userID?: string
  yourReference?: string
}

export interface PayoutRemoveMatch {
  id: string

  // Selects a custom action instead of the plain remove:
  //   'batchdelete'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface PayoutKeysetPage {
  accountID?: string
  amount?: number
  amountMinorUnits?: number
  approvePayoutUrl?: string
  approverID?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  batchPayoutID?: string
  beneficiary: Record<string, any>
  canAuthorise?: boolean
  canProcess?: boolean
  canUpdate?: boolean
  chargeBearer?: string
  createdBy?: string
  createdByEmailAddress?: string
  currency?: string
  currentUserID?: string
  description?: string
  destination?: Record<string, any>
  documents?: any[]
  events?: any[]
  formattedAmount?: string
  formattedFxDestinationAmount?: string
  formattedSchedule?: string
  formattedScheduleDayOnly?: string
  formattedSourceAccountAvailableBalance?: string
  fxDestinationAmount?: number
  fxDestinationAmountMinorUnits?: number
  fxDestinationCurrency?: string
  fxQuoteExpiresAt?: string
  fxQuoteID?: string
  fxRate?: number
  fxUseDestinationAmount?: boolean
  hasCurrentUserAuthorised?: boolean
  id?: string
  inserted?: string
  invoiceID?: string
  isArchived?: boolean
  isFailed?: boolean
  isSettled?: boolean
  isSubmitted?: boolean
  lastUpdated?: string
  merchantID?: string
  merchantTokenDescription?: string
  nonce?: string
  paymentProcessor?: string
  paymentRail?: string
  payrunID?: string
  payrunName?: string
  rule?: Record<string, any>
  scheduleDate?: string
  scheduled?: boolean
  sourceAccountAvailableBalance?: number
  sourceAccountAvailableBalanceMinorUnits?: number
  sourceAccountBic?: string
  sourceAccountCurrency?: string
  sourceAccountIban?: string
  sourceAccountIdentifier: Record<string, any>
  sourceAccountName?: string
  sourceAccountNumber?: string
  sourceAccountSortcode?: string
  status?: string
  tags?: any[]
  theirReference?: string
  topupPayrunID?: string
  transactedAmount?: number
  transactedFxAmount?: number
  transactedFxRate?: number
  type?: string
  userID?: string
  yourReference?: string
}

export interface PayoutKeysetPageListMatch {
  account_id?: string
  merchant_id?: string
}

export interface PayoutMetric {
}

export interface PayoutMetricLoadMatch {
}

export interface Payrun {
  authorisationDate?: string
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  batchPayoutID?: string
  canAuthorise?: boolean
  canDelete?: boolean
  canEdit?: boolean
  events?: any[]
  hasCurrentUserAuthorised?: boolean
  id?: string
  inserted?: string
  invoices?: any[]
  invoicesMinimal?: any[]
  isArchived?: boolean
  lastUpdated?: string
  lastUpdatedBy: Record<string, any>
  merchantID?: string
  name?: string
  nonce?: string
  notes?: string
  payments?: any[]
  payouts?: any[]
  payoutsCount?: number
  reason?: string
  scheduleDate?: string
  scheduledDate?: string
  sourceAccounts?: any[]
  status?: string
  totalEur?: number
  totalGbp?: number
  totalUsd?: number
}

export interface PayrunLoadMatch {
  id: string
}

export interface PayrunListMatch {
  authorisationDate?: string
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  batchPayoutID?: string
  canAuthorise?: boolean
  canDelete?: boolean
  canEdit?: boolean
  events?: any[]
  hasCurrentUserAuthorised?: boolean
  id?: string
  inserted?: string
  invoices?: any[]
  invoicesMinimal?: any[]
  isArchived?: boolean
  lastUpdated?: string
  lastUpdatedBy?: Record<string, any>
  merchantID?: string
  name?: string
  nonce?: string
  notes?: string
  payments?: any[]
  payouts?: any[]
  payoutsCount?: number
  reason?: string
  scheduleDate?: string
  scheduledDate?: string
  sourceAccounts?: any[]
  status?: string
  totalEur?: number
  totalGbp?: number
  totalUsd?: number
}

export interface PayrunCreateData {
  id: string
  authorisationDate?: string
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  batchPayoutID?: string
  canAuthorise?: boolean
  canDelete?: boolean
  canEdit?: boolean
  events?: any[]
  hasCurrentUserAuthorised?: boolean
  inserted?: string
  invoices?: any[]
  invoicesMinimal?: any[]
  isArchived?: boolean
  lastUpdated?: string
  lastUpdatedBy: Record<string, any>
  merchantID?: string
  name?: string
  nonce?: string
  notes?: string
  payments?: any[]
  payouts?: any[]
  payoutsCount?: number
  reason?: string
  scheduleDate?: string
  scheduledDate?: string
  sourceAccounts?: any[]
  status?: string
  totalEur?: number
  totalGbp?: number
  totalUsd?: number

  // Selects a custom action instead of the plain create:
  //   'request_authorisation' | 'submit'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface PayrunUpdateData {
  id: string
  authorisationDate?: string
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  batchPayoutID?: string
  canAuthorise?: boolean
  canDelete?: boolean
  canEdit?: boolean
  events?: any[]
  hasCurrentUserAuthorised?: boolean
  inserted?: string
  invoices?: any[]
  invoicesMinimal?: any[]
  isArchived?: boolean
  lastUpdated?: string
  lastUpdatedBy?: Record<string, any>
  merchantID?: string
  name?: string
  nonce?: string
  notes?: string
  payments?: any[]
  payouts?: any[]
  payoutsCount?: number
  reason?: string
  scheduleDate?: string
  scheduledDate?: string
  sourceAccounts?: any[]
  status?: string
  totalEur?: number
  totalGbp?: number
  totalUsd?: number

  // Selects a custom action instead of the plain update:
  //   'cancel' | 'reject' | 'unarchive'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface PayrunRemoveMatch {
  id: string

  // Selects a custom action instead of the plain remove:
  //   'archive'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface Report {
}

export interface ReportUpdateData {
  id: string

  // Selects a custom action instead of the plain update:
  //   'initiate'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface ReportResult {
  contentType?: string
  contents?: string
  lastCompletedAt?: string
  merchantID?: string
  reportName?: string
  reportType?: string
  statementNumber?: number
}

export interface ReportResultLoadMatch {
  id: number
  report_id: string
}

export interface Role {
  failedRoles?: Record<string, any>
  roles?: any[]
}

export interface RoleCreateData {
  merchant_id: string
  failedRoles?: Record<string, any>
  roles?: any[]

  // Selects a custom action instead of the plain create:
  //   'batchcreate'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface Rule {
  account?: Record<string, any>
  accountID?: string
  approveUrl?: string
  approverID?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  canAuthorise?: boolean
  createdBy: Record<string, any>
  description?: string
  endAt?: string
  hasCurrentUserAuthorised?: boolean
  id?: string
  inserted?: string
  isDisabled?: boolean
  lastExecutedAt?: string
  lastRunAtTransactionDate?: string
  lastUpdated?: string
  merchantID?: string
  name?: string
  nonce: string
  onApprovedWebHookUrl?: string
  onExecutionErrorWebHookUrl?: string
  onExecutionSuccessWebHookUrl?: string
  startAt?: string
  status?: string
  sweepAction?: Record<string, any>
  timeZoneId?: string
  triggerCronExpression?: string
  triggerOnPayIn?: boolean
  userID?: string
  webHookSecret?: string
}

export interface RuleLoadMatch {
  id: string
}

export interface RuleListMatch {
  account?: Record<string, any>
  accountID?: string
  approveUrl?: string
  approverID?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  canAuthorise?: boolean
  createdBy?: Record<string, any>
  description?: string
  endAt?: string
  hasCurrentUserAuthorised?: boolean
  id?: string
  inserted?: string
  isDisabled?: boolean
  lastExecutedAt?: string
  lastRunAtTransactionDate?: string
  lastUpdated?: string
  merchantID?: string
  name?: string
  nonce?: string
  onApprovedWebHookUrl?: string
  onExecutionErrorWebHookUrl?: string
  onExecutionSuccessWebHookUrl?: string
  startAt?: string
  status?: string
  sweepAction?: Record<string, any>
  timeZoneId?: string
  triggerCronExpression?: string
  triggerOnPayIn?: boolean
  userID?: string
  webHookSecret?: string
}

export interface RuleCreateData {
  account?: Record<string, any>
  accountID?: string
  approveUrl?: string
  approverID?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  canAuthorise?: boolean
  createdBy: Record<string, any>
  description?: string
  endAt?: string
  hasCurrentUserAuthorised?: boolean
  id?: string
  inserted?: string
  isDisabled?: boolean
  lastExecutedAt?: string
  lastRunAtTransactionDate?: string
  lastUpdated?: string
  merchantID?: string
  name?: string
  nonce: string
  onApprovedWebHookUrl?: string
  onExecutionErrorWebHookUrl?: string
  onExecutionSuccessWebHookUrl?: string
  startAt?: string
  status?: string
  sweepAction?: Record<string, any>
  timeZoneId?: string
  triggerCronExpression?: string
  triggerOnPayIn?: boolean
  userID?: string
  webHookSecret?: string
}

export interface RuleUpdateData {
  id: string
  account?: Record<string, any>
  accountID?: string
  approveUrl?: string
  approverID?: string
  authenticationMethods?: any[]
  authorisations?: any[]
  authorisersCompletedCount?: number
  authorisersRequiredCount?: number
  canAuthorise?: boolean
  createdBy?: Record<string, any>
  description?: string
  endAt?: string
  hasCurrentUserAuthorised?: boolean
  inserted?: string
  isDisabled?: boolean
  lastExecutedAt?: string
  lastRunAtTransactionDate?: string
  lastUpdated?: string
  merchantID?: string
  name?: string
  nonce?: string
  onApprovedWebHookUrl?: string
  onExecutionErrorWebHookUrl?: string
  onExecutionSuccessWebHookUrl?: string
  startAt?: string
  status?: string
  sweepAction?: Record<string, any>
  timeZoneId?: string
  triggerCronExpression?: string
  triggerOnPayIn?: boolean
  userID?: string
  webHookSecret?: string

  // Selects a custom action instead of the plain update:
  //   'disable'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface RuleRemoveMatch {
  id: string
}

export interface RuleEvent {
  errorMessage?: string
  id?: string
  inserted?: string
  isAuthoriseToEnable?: boolean
  message?: string
  rawResponse?: string
  ruleEventType?: string
  ruleID?: string
  user: Record<string, any>
}

export interface RuleEventListMatch {
  id: string
}

export interface Tag {
  colourHex?: string
  description?: string
  id?: string
  merchantID: string
  name: string
}

export interface TagListMatch {
  merchant_id: string
}

export interface TagCreateData {
  merchant_id: string
  colourHex?: string
  description?: string
  id?: string
  merchantID: string
  name: string
}

export interface Token {
}

export interface TokenCreateData {
  id: string
}

export interface TokenRemoveMatch {
  id: string
}

export interface Transaction {
  accountID?: string
  accountName?: string
  accountSequenceNumber?: number
  addressDetails?: Record<string, any>
  amount?: number
  amountMinorUnits?: number
  balance?: number
  balanceMinorUnits?: number
  bookingDateTime?: string
  chargeDetails?: Record<string, any>
  content?: any[]
  counterparty?: Record<string, any>
  counterpartySummary?: string
  currency?: string
  currencyExchange?: Record<string, any>
  date?: string
  description?: string
  enrichment?: Record<string, any>
  fxAmount?: number
  fxCurrency?: string
  fxRate?: number
  grossAmount: Record<string, any>
  id?: string
  inserted?: string
  isoBankTransactionCode?: Record<string, any>
  merchant?: Record<string, any>
  merchantID?: string
  pageNumber?: number
  pageSize?: number
  payeeDetails: Record<string, any>
  payerDetails: Record<string, any>
  paymentRequestCustomFields?: Record<string, any>
  paymentRequestID?: string
  payoutID?: string
  proprietaryBankTransactionCode?: Record<string, any>
  rawReference?: string
  reference?: string
  ruleID?: string
  statementReferences?: any[]
  status?: string
  supplementaryData?: any
  tags?: any[]
  theirReference?: string
  totalPages?: number
  totalSize?: number
  transactionAmount: Record<string, any>
  transactionDate?: string
  transactionInformation?: any[]
  transactionMutability?: string
  type?: string
  valueDateTime?: string
  virtualIBAN?: string
  yourReference?: string
}

export interface TransactionLoadMatch {
  id?: string
  sequence_number?: number
  transaction_id?: string
  account_id?: string

  // Selects a custom action instead of the plain load:
  //   'export' | 'proof'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface TransactionListMatch {
  account_id?: string
  id?: string
  merchant_id?: string
}

export interface TransactionCreateData {
  id: string
  accountID?: string
  accountName?: string
  accountSequenceNumber?: number
  addressDetails?: Record<string, any>
  amount?: number
  amountMinorUnits?: number
  balance?: number
  balanceMinorUnits?: number
  bookingDateTime?: string
  chargeDetails?: Record<string, any>
  content?: any[]
  counterparty?: Record<string, any>
  counterpartySummary?: string
  currency?: string
  currencyExchange?: Record<string, any>
  date?: string
  description?: string
  enrichment?: Record<string, any>
  fxAmount?: number
  fxCurrency?: string
  fxRate?: number
  grossAmount: Record<string, any>
  inserted?: string
  isoBankTransactionCode?: Record<string, any>
  merchant?: Record<string, any>
  merchantID?: string
  pageNumber?: number
  pageSize?: number
  payeeDetails: Record<string, any>
  payerDetails: Record<string, any>
  paymentRequestCustomFields?: Record<string, any>
  paymentRequestID?: string
  payoutID?: string
  proprietaryBankTransactionCode?: Record<string, any>
  rawReference?: string
  reference?: string
  ruleID?: string
  statementReferences?: any[]
  status?: string
  supplementaryData?: any
  tags?: any[]
  theirReference?: string
  totalPages?: number
  totalSize?: number
  transactionAmount: Record<string, any>
  transactionDate?: string
  transactionInformation?: any[]
  transactionMutability?: string
  type?: string
  valueDateTime?: string
  virtualIBAN?: string
  yourReference?: string

  // Selects a custom action instead of the plain create:
  //   'tag'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface TransactionRemoveMatch {
  id: string

  // Selects a custom action instead of the plain remove:
  //   'tag'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface User {
  clientSessionTimeouts?: any[]
  emailAddress: string
  firstName: string
  id?: string
  lastName: string
  passkeyAdded?: boolean
  permissions?: Record<string, any>
  profile?: string
  rolesWithScope?: any[]
  twoFactorEnabled?: boolean
  userInviteID?: string
}

export interface UserListMatch {
  merchant_id?: string

  // Selects a custom action instead of the plain list:
  //   'userspaged'
  // The remaining keys are that action's own payload.
  $action?: string
  [action: string]: any
}

export interface UserUpdateData {
  id: string
  clientSessionTimeouts?: any[]
  emailAddress?: string
  firstName?: string
  lastName?: string
  passkeyAdded?: boolean
  permissions?: Record<string, any>
  profile?: string
  rolesWithScope?: any[]
  twoFactorEnabled?: boolean
  userInviteID?: string
}

export interface UserInvite {
  authorisationStatus?: Record<string, any>
  failedUserInvites?: Record<string, any>
  id?: string
  initialRoleID?: string
  inviteeEmailAddress?: string
  inviteeFirstName?: string
  inviteeLastName?: string
  inviterEmailAddress?: string
  inviterFirstName?: string
  inviterLastName?: string
  isAuthorised?: boolean
  isInviteeRegistered?: boolean
  lastInvited?: string
  merchantID?: string
  merchantName?: string
  message?: string
  registrationUrl?: string
  sendInviteEmail?: boolean
  status?: string
  user: Record<string, any>
  userID?: string
  userInvites?: any[]
}

export interface UserInviteLoadMatch {
  id?: string
  userinvite_id?: string
}

export interface UserInviteListMatch {
  merchant_id: string
}

export interface UserInviteCreateData {
  id?: string
  authorisationStatus?: Record<string, any>
  failedUserInvites?: Record<string, any>
  initialRoleID?: string
  inviteeEmailAddress?: string
  inviteeFirstName?: string
  inviteeLastName?: string
  inviterEmailAddress?: string
  inviterFirstName?: string
  inviterLastName?: string
  isAuthorised?: boolean
  isInviteeRegistered?: boolean
  lastInvited?: string
  merchantID?: string
  merchantName?: string
  message?: string
  registrationUrl?: string
  sendInviteEmail?: boolean
  status?: string
  user: Record<string, any>
  userID?: string
  userInvites?: any[]
}

export interface UserInviteUpdateData {
  id: string
  authorisationStatus?: Record<string, any>
  failedUserInvites?: Record<string, any>
  initialRoleID?: string
  inviteeEmailAddress?: string
  inviteeFirstName?: string
  inviteeLastName?: string
  inviterEmailAddress?: string
  inviterFirstName?: string
  inviterLastName?: string
  isAuthorised?: boolean
  isInviteeRegistered?: boolean
  lastInvited?: string
  merchantID?: string
  merchantName?: string
  message?: string
  registrationUrl?: string
  sendInviteEmail?: boolean
  status?: string
  user?: Record<string, any>
  userID?: string
  userInvites?: any[]
}

export interface UserInviteRemoveMatch {
  id: string
}

export interface Virtual {
  accountName?: string
  accountSupplierName?: string
  availableBalance?: number
  availableBalanceMinorUnits?: number
  balance?: number
  balanceMinorUnits?: number
  bankName?: string
  consentID?: string
  createdBy: Record<string, any>
  createdByDisplayName?: string
  currency?: string
  defaultPaymentRail?: string
  displayName?: string
  expiryDate?: string
  externalAccountIcon?: string
  id?: string
  identifier: Record<string, any>
  inserted?: string
  isArchived?: boolean
  isConnectedAccount?: boolean
  isDefault?: boolean
  isTrustAccount?: boolean
  isVirtual?: boolean
  lastTransaction?: Record<string, any>
  lastUpdated?: string
  merchantID?: string
  merchantName?: string
  name: string
  physicalAccountID?: string
  rules?: any[]
  submittedPayoutsBalance?: number
  submittedPayoutsBalanceMinorUnits?: number
  summary?: string
  supplierSepaInstantStatus?: string
  xeroBankFeedConnectionStatus?: string
  xeroBankFeedLastSyncedAt?: string
  xeroBankFeedSyncLastFailedAt?: string
  xeroBankFeedSyncLastFailureReason?: string
  xeroBankFeedSyncStatus?: string
  xeroUnsynchronisedTransactionsCount?: number
}

export interface VirtualCreateData {
  account_id: string
  accountName?: string
  accountSupplierName?: string
  availableBalance?: number
  availableBalanceMinorUnits?: number
  balance?: number
  balanceMinorUnits?: number
  bankName?: string
  consentID?: string
  createdBy: Record<string, any>
  createdByDisplayName?: string
  currency?: string
  defaultPaymentRail?: string
  displayName?: string
  expiryDate?: string
  externalAccountIcon?: string
  id?: string
  identifier: Record<string, any>
  inserted?: string
  isArchived?: boolean
  isConnectedAccount?: boolean
  isDefault?: boolean
  isTrustAccount?: boolean
  isVirtual?: boolean
  lastTransaction?: Record<string, any>
  lastUpdated?: string
  merchantID?: string
  merchantName?: string
  name: string
  physicalAccountID?: string
  rules?: any[]
  submittedPayoutsBalance?: number
  submittedPayoutsBalanceMinorUnits?: number
  summary?: string
  supplierSepaInstantStatus?: string
  xeroBankFeedConnectionStatus?: string
  xeroBankFeedLastSyncedAt?: string
  xeroBankFeedSyncLastFailedAt?: string
  xeroBankFeedSyncLastFailureReason?: string
  xeroBankFeedSyncStatus?: string
  xeroUnsynchronisedTransactionsCount?: number
}

export interface VirtualUpdateData {
  account_id: string
  id: string
  accountName?: string
  accountSupplierName?: string
  availableBalance?: number
  availableBalanceMinorUnits?: number
  balance?: number
  balanceMinorUnits?: number
  bankName?: string
  consentID?: string
  createdBy?: Record<string, any>
  createdByDisplayName?: string
  currency?: string
  defaultPaymentRail?: string
  displayName?: string
  expiryDate?: string
  externalAccountIcon?: string
  identifier?: Record<string, any>
  inserted?: string
  isArchived?: boolean
  isConnectedAccount?: boolean
  isDefault?: boolean
  isTrustAccount?: boolean
  isVirtual?: boolean
  lastTransaction?: Record<string, any>
  lastUpdated?: string
  merchantID?: string
  merchantName?: string
  name?: string
  physicalAccountID?: string
  rules?: any[]
  submittedPayoutsBalance?: number
  submittedPayoutsBalanceMinorUnits?: number
  summary?: string
  supplierSepaInstantStatus?: string
  xeroBankFeedConnectionStatus?: string
  xeroBankFeedLastSyncedAt?: string
  xeroBankFeedSyncLastFailedAt?: string
  xeroBankFeedSyncLastFailureReason?: string
  xeroBankFeedSyncStatus?: string
  xeroUnsynchronisedTransactionsCount?: number
}

export interface Webhook {
  destinationUrl?: string
  emailAddress?: string
  failedNotificationEmailAddress?: string
  id?: string
  isActive?: boolean
  merchantID?: string
  notificationMethod?: string
  resourceTypes?: any[]
  retry?: boolean
  secret?: string
  version?: number
}

export interface WebhookLoadMatch {
  id: string
  merchant_id?: string
}

export interface WebhookListMatch {
  merchant_id: string
}

export interface WebhookCreateData {
  destinationUrl?: string
  emailAddress?: string
  failedNotificationEmailAddress?: string
  id?: string
  isActive?: boolean
  merchantID?: string
  notificationMethod?: string
  resourceTypes?: any[]
  retry?: boolean
  secret?: string
  version?: number
}

export interface WebhookUpdateData {
  id: string
  destinationUrl?: string
  emailAddress?: string
  failedNotificationEmailAddress?: string
  isActive?: boolean
  merchantID?: string
  notificationMethod?: string
  resourceTypes?: any[]
  retry?: boolean
  secret?: string
  version?: number
}

export interface WebhookRemoveMatch {
  id: string
}

