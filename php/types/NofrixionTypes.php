<?php
declare(strict_types=1);

// Typed models for the Nofrixion SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Account entity data model. */
class Account
{
    public ?array $accountBalances = null;
    public ?string $accountID = null;
    public ?array $accountIdentifications = null;
    public ?string $accountName = null;
    public ?array $accountNames = null;
    public ?string $accountSupplierName = null;
    public ?string $accountType = null;
    public ?float $availableBalance = null;
    public ?int $availableBalanceMinorUnits = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $bankName = null;
    public ?string $consentID = null;
    public ?array $consolidatedAccountInformation = null;
    public array $createdBy;
    public ?string $createdByDisplayName = null;
    public ?string $currency = null;
    public ?string $defaultPaymentRail = null;
    public ?string $description = null;
    public ?string $details = null;
    public ?string $displayName = null;
    public ?string $expiryDate = null;
    public ?string $externalAccountIcon = null;
    public ?string $format = null;
    public ?string $fromDate = null;
    public ?string $id = null;
    public array $identifier;
    public ?string $inserted = null;
    public ?bool $isArchived = null;
    public ?bool $isConnectedAccount = null;
    public ?bool $isDefault = null;
    public ?bool $isTrustAccount = null;
    public ?bool $isVirtual = null;
    public ?array $lastTransaction = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public ?string $nickname = null;
    public ?string $physicalAccountID = null;
    public ?array $roleIDs = null;
    public ?array $rules = null;
    public ?float $submittedPayoutsBalance = null;
    public ?int $submittedPayoutsBalanceMinorUnits = null;
    public ?string $summary = null;
    public ?string $supplierPhysicalAccountID = null;
    public ?string $supplierSepaInstantStatus = null;
    public ?string $toDate = null;
    public ?string $type = null;
    public ?string $usageType = null;
    public ?string $xeroBankFeedConnectionStatus = null;
    public ?string $xeroBankFeedLastSyncedAt = null;
    public ?string $xeroBankFeedSyncLastFailedAt = null;
    public ?string $xeroBankFeedSyncLastFailureReason = null;
    public ?string $xeroBankFeedSyncStatus = null;
    public ?int $xeroUnsynchronisedTransactionsCount = null;
}

/** Request payload for Account#load. */
class AccountLoadMatch
{
    public ?string $account_id = null;
    public string $id;
    public ?string $merchant_id = null;
}

/** Request payload for Account#list. */
class AccountListMatch
{
    public ?string $merchant_id = null;
}

/** Request payload for Account#create. */
class AccountCreateData
{
    public ?string $account_id = null;
    public ?string $currency = null;
    public ?array $accountBalances = null;
    public ?string $accountID = null;
    public ?array $accountIdentifications = null;
    public ?string $accountName = null;
    public ?array $accountNames = null;
    public ?string $accountSupplierName = null;
    public ?string $accountType = null;
    public ?float $availableBalance = null;
    public ?int $availableBalanceMinorUnits = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $bankName = null;
    public ?string $consentID = null;
    public ?array $consolidatedAccountInformation = null;
    public array $createdBy;
    public ?string $createdByDisplayName = null;
    public ?string $defaultPaymentRail = null;
    public ?string $description = null;
    public ?string $details = null;
    public ?string $displayName = null;
    public ?string $expiryDate = null;
    public ?string $externalAccountIcon = null;
    public ?string $format = null;
    public ?string $fromDate = null;
    public ?string $id = null;
    public array $identifier;
    public ?string $inserted = null;
    public ?bool $isArchived = null;
    public ?bool $isConnectedAccount = null;
    public ?bool $isDefault = null;
    public ?bool $isTrustAccount = null;
    public ?bool $isVirtual = null;
    public ?array $lastTransaction = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public ?string $nickname = null;
    public ?string $physicalAccountID = null;
    public ?array $roleIDs = null;
    public ?array $rules = null;
    public ?float $submittedPayoutsBalance = null;
    public ?int $submittedPayoutsBalanceMinorUnits = null;
    public ?string $summary = null;
    public ?string $supplierPhysicalAccountID = null;
    public ?string $supplierSepaInstantStatus = null;
    public ?string $toDate = null;
    public ?string $type = null;
    public ?string $usageType = null;
    public ?string $xeroBankFeedConnectionStatus = null;
    public ?string $xeroBankFeedLastSyncedAt = null;
    public ?string $xeroBankFeedSyncLastFailedAt = null;
    public ?string $xeroBankFeedSyncLastFailureReason = null;
    public ?string $xeroBankFeedSyncStatus = null;
    public ?int $xeroUnsynchronisedTransactionsCount = null;
}

/** Request payload for Account#update. */
class AccountUpdateData
{
    public ?string $account_id = null;
    public ?float $amount = null;
    public ?string $id = null;
    public ?array $accountBalances = null;
    public ?string $accountID = null;
    public ?array $accountIdentifications = null;
    public ?string $accountName = null;
    public ?array $accountNames = null;
    public ?string $accountSupplierName = null;
    public ?string $accountType = null;
    public ?float $availableBalance = null;
    public ?int $availableBalanceMinorUnits = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $bankName = null;
    public ?string $consentID = null;
    public ?array $consolidatedAccountInformation = null;
    public ?array $createdBy = null;
    public ?string $createdByDisplayName = null;
    public ?string $currency = null;
    public ?string $defaultPaymentRail = null;
    public ?string $description = null;
    public ?string $details = null;
    public ?string $displayName = null;
    public ?string $expiryDate = null;
    public ?string $externalAccountIcon = null;
    public ?string $format = null;
    public ?string $fromDate = null;
    public ?array $identifier = null;
    public ?string $inserted = null;
    public ?bool $isArchived = null;
    public ?bool $isConnectedAccount = null;
    public ?bool $isDefault = null;
    public ?bool $isTrustAccount = null;
    public ?bool $isVirtual = null;
    public ?array $lastTransaction = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public ?string $nickname = null;
    public ?string $physicalAccountID = null;
    public ?array $roleIDs = null;
    public ?array $rules = null;
    public ?float $submittedPayoutsBalance = null;
    public ?int $submittedPayoutsBalanceMinorUnits = null;
    public ?string $summary = null;
    public ?string $supplierPhysicalAccountID = null;
    public ?string $supplierSepaInstantStatus = null;
    public ?string $toDate = null;
    public ?string $type = null;
    public ?string $usageType = null;
    public ?string $xeroBankFeedConnectionStatus = null;
    public ?string $xeroBankFeedLastSyncedAt = null;
    public ?string $xeroBankFeedSyncLastFailedAt = null;
    public ?string $xeroBankFeedSyncLastFailureReason = null;
    public ?string $xeroBankFeedSyncStatus = null;
    public ?int $xeroUnsynchronisedTransactionsCount = null;
}

/** Request payload for Account#remove. */
class AccountRemoveMatch
{
    public string $id;
}

/** Batch entity data model. */
class Batch
{
    public ?string $approveUrl = null;
    public ?string $id = null;
    public ?array $payouts = null;
}

/** Request payload for Batch#load. */
class BatchLoadMatch
{
    public string $id;
}

/** Request payload for Batch#create. */
class BatchCreateData
{
    public ?string $approveUrl = null;
    public ?string $id = null;
    public ?array $payouts = null;
}

/** Beneficiary entity data model. */
class Beneficiary
{
    public ?string $approvalCallbackUrl = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?array $beneficiaries = null;
    public ?array $beneficiaryEvents = null;
    public ?bool $canAuthorise = null;
    public ?bool $canUpdate = null;
    public array $createdBy;
    public ?string $createdByEmailAddress = null;
    public string $currency;
    public ?array $destination = null;
    public ?array $failedBeneficiaries = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isEnabled = null;
    public ?string $lastAuthorised = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public string $name;
    public ?string $nonce = null;
    public ?array $sourceAccountIDs = null;
    public ?array $sourceAccounts = null;
    public ?string $theirReference = null;
}

/** Request payload for Beneficiary#load. */
class BeneficiaryLoadMatch
{
    public string $id;
    public ?string $merchant_id = null;
}

/** Request payload for Beneficiary#list. */
class BeneficiaryListMatch
{
    public ?string $merchant_id = null;
}

/** Request payload for Beneficiary#create. */
class BeneficiaryCreateData
{
    public ?string $id = null;
    public ?string $approvalCallbackUrl = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?array $beneficiaries = null;
    public ?array $beneficiaryEvents = null;
    public ?bool $canAuthorise = null;
    public ?bool $canUpdate = null;
    public array $createdBy;
    public ?string $createdByEmailAddress = null;
    public string $currency;
    public ?array $destination = null;
    public ?array $failedBeneficiaries = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $inserted = null;
    public ?bool $isEnabled = null;
    public ?string $lastAuthorised = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public string $name;
    public ?string $nonce = null;
    public ?array $sourceAccountIDs = null;
    public ?array $sourceAccounts = null;
    public ?string $theirReference = null;
}

/** Request payload for Beneficiary#update. */
class BeneficiaryUpdateData
{
    public string $id;
    public ?string $approvalCallbackUrl = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?array $beneficiaries = null;
    public ?array $beneficiaryEvents = null;
    public ?bool $canAuthorise = null;
    public ?bool $canUpdate = null;
    public ?array $createdBy = null;
    public ?string $createdByEmailAddress = null;
    public ?string $currency = null;
    public ?array $destination = null;
    public ?array $failedBeneficiaries = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $inserted = null;
    public ?bool $isEnabled = null;
    public ?string $lastAuthorised = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?array $sourceAccountIDs = null;
    public ?array $sourceAccounts = null;
    public ?string $theirReference = null;
}

/** Request payload for Beneficiary#remove. */
class BeneficiaryRemoveMatch
{
    public string $id;
}

/** BeneficiaryGroup entity data model. */
class BeneficiaryGroup
{
    public ?array $groupMembers = null;
    public string $groupName;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $lastUpdated = null;
    public string $merchantID;
}

/** Request payload for BeneficiaryGroup#list. */
class BeneficiaryGroupListMatch
{
    public string $merchant_id;
}

/** Card entity data model. */
class Card
{
    public ?string $authorizedAmount = null;
    public ?string $currencyCode = null;
    public ?bool $isPayerAuthenticationRequired = null;
    public ?bool $isSoftDecline = null;
    public ?string $payerAuthenticationAccessToken = null;
    public ?string $payerAuthenticationMerchantData = null;
    public ?string $payerAuthenticationUrl = null;
    public ?int $payerAuthenticationWindowHeight = null;
    public ?int $payerAuthenticationWindowWidth = null;
    public ?string $paymentRequestCallbackUrl = null;
    public ?string $paymentRequestID = null;
    public ?string $requestID = null;
    public ?string $responseCode = null;
    public ?string $responseType = null;
    public ?string $status = null;
    public ?string $threeDSRedirectUrl = null;
    public ?string $transactionID = null;
}

/** Request payload for Card#create. */
class CardCreateData
{
    public string $paymentrequest_id;
    public ?string $authorizedAmount = null;
    public ?string $currencyCode = null;
    public ?bool $isPayerAuthenticationRequired = null;
    public ?bool $isSoftDecline = null;
    public ?string $payerAuthenticationAccessToken = null;
    public ?string $payerAuthenticationMerchantData = null;
    public ?string $payerAuthenticationUrl = null;
    public ?int $payerAuthenticationWindowHeight = null;
    public ?int $payerAuthenticationWindowWidth = null;
    public ?string $paymentRequestCallbackUrl = null;
    public ?string $paymentRequestID = null;
    public ?string $requestID = null;
    public ?string $responseCode = null;
    public ?string $responseType = null;
    public ?string $status = null;
    public ?string $threeDSRedirectUrl = null;
    public ?string $transactionID = null;
}

/** CardCustomerToken entity data model. */
class CardCustomerToken
{
    public ?string $cardType = null;
    public ?string $customerEmailAddress = null;
    public ?string $expiryMonth = null;
    public ?string $expiryYear = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $lastFourDigits = null;
    public ?string $lastUpdated = null;
    public ?string $maskedCardNumber = null;
    public ?string $merchantID = null;
    public ?string $paymentRequestID = null;
}

/** Request payload for CardCustomerToken#load. */
class CardCustomerTokenLoadMatch
{
    public string $customer_email_address;
}

/** Request payload for CardCustomerToken#list. */
class CardCustomerTokenListMatch
{
    public string $customer_email_address;
    public string $merchant_id;
}

/** Request payload for CardCustomerToken#remove. */
class CardCustomerTokenRemoveMatch
{
    public ?string $customer_email_address = null;
    public ?string $merchant_id = null;
    public ?string $id = null;
}

/** CardPayment entity data model. */
class CardPayment
{
    public ?string $authorizedAmount = null;
    public ?string $currencyCode = null;
    public ?bool $isPayerAuthenticationRequired = null;
    public ?bool $isSoftDecline = null;
    public ?string $payerAuthenticationAccessToken = null;
    public ?string $payerAuthenticationMerchantData = null;
    public ?string $payerAuthenticationUrl = null;
    public ?int $payerAuthenticationWindowHeight = null;
    public ?int $payerAuthenticationWindowWidth = null;
    public ?string $paymentRequestCallbackUrl = null;
    public ?string $paymentRequestID = null;
    public ?string $requestID = null;
    public ?string $responseCode = null;
    public ?string $responseType = null;
    public ?string $status = null;
    public ?string $threeDSRedirectUrl = null;
    public ?string $transactionID = null;
}

/** Request payload for CardPayment#create. */
class CardPaymentCreateData
{
    public ?float $partial_refund_amount = null;
    public string $paymentrequest_id;
    public ?string $authorizedAmount = null;
    public ?string $currencyCode = null;
    public ?bool $isPayerAuthenticationRequired = null;
    public ?bool $isSoftDecline = null;
    public ?string $payerAuthenticationAccessToken = null;
    public ?string $payerAuthenticationMerchantData = null;
    public ?string $payerAuthenticationUrl = null;
    public ?int $payerAuthenticationWindowHeight = null;
    public ?int $payerAuthenticationWindowWidth = null;
    public ?string $paymentRequestCallbackUrl = null;
    public ?string $paymentRequestID = null;
    public ?string $requestID = null;
    public ?string $responseCode = null;
    public ?string $responseType = null;
    public ?string $status = null;
    public ?string $threeDSRedirectUrl = null;
    public ?string $transactionID = null;
}

/** CardPublicKey entity data model. */
class CardPublicKey
{
    public ?string $jwt = null;
}

/** Request payload for CardPublicKey#load. */
class CardPublicKeyLoadMatch
{
    public string $paymentrequest_id;
}

/** Consent entity data model. */
class Consent
{
    public ?string $authorisationUrl = null;
    public ?string $callbackUrl = null;
    public ?string $consentID = null;
    public ?string $emailAddress = null;
    public ?string $expiryDate = null;
    public ?string $failureCallbackUrl = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $institutionID = null;
    public ?bool $isConnectedAccounts = null;
    public ?bool $isEnabled = null;
    public ?string $merchantID = null;
    public ?string $provider = null;
    public ?string $successWebHookUrl = null;
}

/** Request payload for Consent#load. */
class ConsentLoadMatch
{
    public string $id;
}

/** Request payload for Consent#list. */
class ConsentListMatch
{
    public string $email;
    public string $merchant_id;
}

/** Request payload for Consent#create. */
class ConsentCreateData
{
    public ?string $authorisationUrl = null;
    public ?string $callbackUrl = null;
    public ?string $consentID = null;
    public ?string $emailAddress = null;
    public ?string $expiryDate = null;
    public ?string $failureCallbackUrl = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $institutionID = null;
    public ?bool $isConnectedAccounts = null;
    public ?bool $isEnabled = null;
    public ?string $merchantID = null;
    public ?string $provider = null;
    public ?string $successWebHookUrl = null;
}

/** Request payload for Consent#update. */
class ConsentUpdateData
{
    public string $id;
    public ?string $authorisationUrl = null;
    public ?string $callbackUrl = null;
    public ?string $consentID = null;
    public ?string $emailAddress = null;
    public ?string $expiryDate = null;
    public ?string $failureCallbackUrl = null;
    public ?string $inserted = null;
    public ?string $institutionID = null;
    public ?bool $isConnectedAccounts = null;
    public ?bool $isEnabled = null;
    public ?string $merchantID = null;
    public ?string $provider = null;
    public ?string $successWebHookUrl = null;
}

/** Request payload for Consent#remove. */
class ConsentRemoveMatch
{
    public string $id;
}

/** Currency entity data model. */
class Currency
{
    public ?string $code = null;
    public ?int $decimals = null;
    public ?bool $isFiat = null;
    public ?string $iso4217AlphaCode = null;
    public ?string $iso4217NumericCode = null;
    public ?string $symbol = null;
}

/** Request payload for Currency#list. */
class CurrencyListMatch
{
    public ?string $code = null;
    public ?int $decimals = null;
    public ?bool $isFiat = null;
    public ?string $iso4217AlphaCode = null;
    public ?string $iso4217NumericCode = null;
    public ?string $symbol = null;
}

/** DirectDebitBatchSubmit entity data model. */
class DirectDebitBatchSubmit
{
    public ?array $failedSubmissions = null;
    public ?array $successfulSubmissions = null;
}

/** Request payload for DirectDebitBatchSubmit#create. */
class DirectDebitBatchSubmitCreateData
{
    public ?array $failedSubmissions = null;
    public ?array $successfulSubmissions = null;
}

/** FxRate entity data model. */
class FxRate
{
    public ?string $destinationCurrency = null;
    public ?float $exchangeRate = null;
    public ?string $expiryTime = null;
    public ?string $quoteID = null;
    public ?string $sourceCurrency = null;
}

/** Request payload for FxRate#load. */
class FxRateLoadMatch
{
    public string $destination;
    public string $source;
    public int $valid_for_minute;
}

/** Request payload for FxRate#list. */
class FxRateListMatch
{
    public string $destination;
    public string $source;
}

/** IPayment entity data model. */
class IPayment
{
    public ?string $paymentRequestID = null;
    public ?string $responseType = null;
}

/** Request payload for IPayment#create. */
class IPaymentCreateData
{
    public ?string $paymentRequestID = null;
    public ?string $responseType = null;
}

/** Mandate entity data model. */
class Mandate
{
    public ?string $accountNumber = null;
    public string $addressLine1;
    public ?string $addressLine2 = null;
    public ?string $approvedAt = null;
    public string $city;
    public string $countryCode;
    public ?string $currency = null;
    public ?string $customerAccountNumber = null;
    public ?string $customerCity = null;
    public ?string $customerCountryCode = null;
    public ?string $customerCountryName = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerFirstName = null;
    public ?string $customerIban = null;
    public ?string $customerLastName = null;
    public ?string $customerSortCode = null;
    public string $emailAddress;
    public string $firstName;
    public ?string $iban = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isRecurring = null;
    public string $lastName;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public string $postalCode;
    public ?string $reference = null;
    public ?string $sortCode = null;
    public ?string $status = null;
    public ?string $supplierBankAccountID = null;
    public ?string $supplierCustomerID = null;
    public ?string $supplierMandateID = null;
    public ?string $supplierName = null;
    public ?string $supplierStatus = null;
}

/** Request payload for Mandate#load. */
class MandateLoadMatch
{
    public string $id;
}

/** Request payload for Mandate#create. */
class MandateCreateData
{
    public ?string $accountNumber = null;
    public string $addressLine1;
    public ?string $addressLine2 = null;
    public ?string $approvedAt = null;
    public string $city;
    public string $countryCode;
    public ?string $currency = null;
    public ?string $customerAccountNumber = null;
    public ?string $customerCity = null;
    public ?string $customerCountryCode = null;
    public ?string $customerCountryName = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerFirstName = null;
    public ?string $customerIban = null;
    public ?string $customerLastName = null;
    public ?string $customerSortCode = null;
    public string $emailAddress;
    public string $firstName;
    public ?string $iban = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isRecurring = null;
    public string $lastName;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public string $postalCode;
    public ?string $reference = null;
    public ?string $sortCode = null;
    public ?string $status = null;
    public ?string $supplierBankAccountID = null;
    public ?string $supplierCustomerID = null;
    public ?string $supplierMandateID = null;
    public ?string $supplierName = null;
    public ?string $supplierStatus = null;
}

/** Merchant entity data model. */
class Merchant
{
    public ?array $accountCurrencies = null;
    public ?bool $canHaveTrustAccounts = null;
    public ?string $cardPaymentProcessor = null;
    public ?string $companyID = null;
    public ?bool $displayQrOnHostedPay = null;
    public ?int $hostedPayVersion = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isBlocked = null;
    public ?bool $isExited = null;
    public ?bool $isSuspended = null;
    public ?string $jurisdiction = null;
    public ?string $logoUrlPng = null;
    public ?string $logoUrlSvg = null;
    public ?string $merchantCategoryCode = null;
    public ?string $name = null;
    public ?string $notes = null;
    public ?array $parentMerchant = null;
    public ?int $paymentAccountLimit = null;
    public ?array $paymentAccounts = null;
    public ?string $reason = null;
    public ?string $shortName = null;
    public ?array $supportedPaymentMethodsList = null;
    public ?string $suspensionReason = null;
    public ?array $tags = null;
    public ?string $timeZoneId = null;
    public ?string $tradingName = null;
    public ?int $webHookLimit = null;
    public ?string $yourRoleName = null;
}

/** Request payload for Merchant#load. */
class MerchantLoadMatch
{
    public string $id;
}

/** Request payload for Merchant#list. */
class MerchantListMatch
{
    public ?array $accountCurrencies = null;
    public ?bool $canHaveTrustAccounts = null;
    public ?string $cardPaymentProcessor = null;
    public ?string $companyID = null;
    public ?bool $displayQrOnHostedPay = null;
    public ?int $hostedPayVersion = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isBlocked = null;
    public ?bool $isExited = null;
    public ?bool $isSuspended = null;
    public ?string $jurisdiction = null;
    public ?string $logoUrlPng = null;
    public ?string $logoUrlSvg = null;
    public ?string $merchantCategoryCode = null;
    public ?string $name = null;
    public ?string $notes = null;
    public ?array $parentMerchant = null;
    public ?int $paymentAccountLimit = null;
    public ?array $paymentAccounts = null;
    public ?string $reason = null;
    public ?string $shortName = null;
    public ?array $supportedPaymentMethodsList = null;
    public ?string $suspensionReason = null;
    public ?array $tags = null;
    public ?string $timeZoneId = null;
    public ?string $tradingName = null;
    public ?int $webHookLimit = null;
    public ?string $yourRoleName = null;
}

/** Request payload for Merchant#update. */
class MerchantUpdateData
{
    public string $id;
    public ?array $accountCurrencies = null;
    public ?bool $canHaveTrustAccounts = null;
    public ?string $cardPaymentProcessor = null;
    public ?string $companyID = null;
    public ?bool $displayQrOnHostedPay = null;
    public ?int $hostedPayVersion = null;
    public ?string $inserted = null;
    public ?bool $isBlocked = null;
    public ?bool $isExited = null;
    public ?bool $isSuspended = null;
    public ?string $jurisdiction = null;
    public ?string $logoUrlPng = null;
    public ?string $logoUrlSvg = null;
    public ?string $merchantCategoryCode = null;
    public ?string $name = null;
    public ?string $notes = null;
    public ?array $parentMerchant = null;
    public ?int $paymentAccountLimit = null;
    public ?array $paymentAccounts = null;
    public ?string $reason = null;
    public ?string $shortName = null;
    public ?array $supportedPaymentMethodsList = null;
    public ?string $suspensionReason = null;
    public ?array $tags = null;
    public ?string $timeZoneId = null;
    public ?string $tradingName = null;
    public ?int $webHookLimit = null;
    public ?string $yourRoleName = null;
}

/** Request payload for Merchant#remove. */
class MerchantRemoveMatch
{
    public ?string $id = null;
    public ?string $user_id = null;
    public ?string $merchant_id = null;
    public ?string $tag_id = null;
}

/** MerchantAuthorisationSetting entity data model. */
class MerchantAuthorisationSetting
{
    public ?float $amountLower = null;
    public ?float $amountUpper = null;
    public ?string $authorisationType = null;
    public ?bool $beneficiariesOnly = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $lastEditorCantAuthorise = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?int $numberOfAuthorisers = null;
    public ?array $roleSettings = null;
}

/** Request payload for MerchantAuthorisationSetting#list. */
class MerchantAuthorisationSettingListMatch
{
    public string $merchant_id;
}

/** MerchantDirectDebitMandatePage entity data model. */
class MerchantDirectDebitMandatePage
{
    public ?string $approvedAt = null;
    public ?string $currency = null;
    public ?string $customerAccountNumber = null;
    public ?string $customerCity = null;
    public ?string $customerCountryCode = null;
    public ?string $customerCountryName = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerFirstName = null;
    public ?string $customerIban = null;
    public ?string $customerLastName = null;
    public ?string $customerSortCode = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isRecurring = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $reference = null;
    public ?string $status = null;
    public ?string $supplierBankAccountID = null;
    public ?string $supplierCustomerID = null;
    public ?string $supplierMandateID = null;
    public ?string $supplierName = null;
    public ?string $supplierStatus = null;
}

/** Request payload for MerchantDirectDebitMandatePage#list. */
class MerchantDirectDebitMandatePageListMatch
{
    public ?string $approvedAt = null;
    public ?string $currency = null;
    public ?string $customerAccountNumber = null;
    public ?string $customerCity = null;
    public ?string $customerCountryCode = null;
    public ?string $customerCountryName = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerFirstName = null;
    public ?string $customerIban = null;
    public ?string $customerLastName = null;
    public ?string $customerSortCode = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isRecurring = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $reference = null;
    public ?string $status = null;
    public ?string $supplierBankAccountID = null;
    public ?string $supplierCustomerID = null;
    public ?string $supplierMandateID = null;
    public ?string $supplierName = null;
    public ?string $supplierStatus = null;
}

/** MerchantPayByBankSetting entity data model. */
class MerchantPayByBankSetting
{
    public ?array $bankCountryCodes = null;
    public ?string $bankID = null;
    public ?string $bankName = null;
    public ?string $businessInstitutionID = null;
    public ?string $currency = null;
    public ?string $logo = null;
    public ?string $message = null;
    public ?string $messageImageUrl = null;
    public ?int $order = null;
    public ?string $personalInstitutionID = null;
    public ?string $processor = null;
    public ?string $warningHeading = null;
    public ?string $warningMessage = null;
}

/** Request payload for MerchantPayByBankSetting#list. */
class MerchantPayByBankSettingListMatch
{
    public string $merchant_id;
}

/** MerchantPaymentRequestTemplate entity data model. */
class MerchantPaymentRequestTemplate
{
    public ?array $bankPaymentOptions = null;
    public ?array $cardPaymentAddressOptions = null;
    public ?array $cardPaymentCaptureOptions = null;
    public ?array $customFields = null;
    public ?array $defaultFields = null;
    public string $description;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public string $name;
    public ?array $notificationOptions = null;
    public ?array $paymentMethods = null;
    public ?array $paymentTerms = null;
    public ?array $priorityBankOptions = null;
    public array $template;
}

/** Request payload for MerchantPaymentRequestTemplate#load. */
class MerchantPaymentRequestTemplateLoadMatch
{
    public string $id;
    public string $paymentrequest_id;
}

/** Request payload for MerchantPaymentRequestTemplate#list. */
class MerchantPaymentRequestTemplateListMatch
{
    public string $merchant_id;
}

/** Request payload for MerchantPaymentRequestTemplate#update. */
class MerchantPaymentRequestTemplateUpdateData
{
    public string $id;
    public string $paymentrequest_id;
    public ?array $bankPaymentOptions = null;
    public ?array $cardPaymentAddressOptions = null;
    public ?array $cardPaymentCaptureOptions = null;
    public ?array $customFields = null;
    public ?array $defaultFields = null;
    public ?string $description = null;
    public ?string $inserted = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $name = null;
    public ?array $notificationOptions = null;
    public ?array $paymentMethods = null;
    public ?array $paymentTerms = null;
    public ?array $priorityBankOptions = null;
    public ?array $template = null;
}

/** Request payload for MerchantPaymentRequestTemplate#remove. */
class MerchantPaymentRequestTemplateRemoveMatch
{
    public string $id;
    public string $paymentrequest_id;
}

/** MerchantToken entity data model. */
class MerchantToken
{
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?bool $canAuthorise = null;
    public ?string $description = null;
    public ?string $expiresAt = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $hmacAlgorithm = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $ipAddressWhitelist = null;
    public ?bool $isArchived = null;
    public ?bool $isEnabled = null;
    public ?string $lastAuthorised = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public string $nonce;
    public ?array $permissionTypes = null;
    public ?int $requestSignatureVersion = null;
    public ?string $sharedSecretAlgorithm = null;
    public ?string $sharedSecretBase64 = null;
    public ?string $token = null;
}

/** Request payload for MerchantToken#load. */
class MerchantTokenLoadMatch
{
    public string $id;
}

/** Request payload for MerchantToken#list. */
class MerchantTokenListMatch
{
    public string $merchant_id;
}

/** Request payload for MerchantToken#create. */
class MerchantTokenCreateData
{
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?bool $canAuthorise = null;
    public ?string $description = null;
    public ?string $expiresAt = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $hmacAlgorithm = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $ipAddressWhitelist = null;
    public ?bool $isArchived = null;
    public ?bool $isEnabled = null;
    public ?string $lastAuthorised = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public string $nonce;
    public ?array $permissionTypes = null;
    public ?int $requestSignatureVersion = null;
    public ?string $sharedSecretAlgorithm = null;
    public ?string $sharedSecretBase64 = null;
    public ?string $token = null;
}

/** Request payload for MerchantToken#update. */
class MerchantTokenUpdateData
{
    public string $id;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?bool $canAuthorise = null;
    public ?string $description = null;
    public ?string $expiresAt = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $hmacAlgorithm = null;
    public ?string $inserted = null;
    public ?string $ipAddressWhitelist = null;
    public ?bool $isArchived = null;
    public ?bool $isEnabled = null;
    public ?string $lastAuthorised = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $nonce = null;
    public ?array $permissionTypes = null;
    public ?int $requestSignatureVersion = null;
    public ?string $sharedSecretAlgorithm = null;
    public ?string $sharedSecretBase64 = null;
    public ?string $token = null;
}

/** Metadata entity data model. */
class Metadata
{
}

/** Request payload for Metadata#load. */
class MetadataLoadMatch
{
}

/** NoFrixionVersion entity data model. */
class NoFrixionVersion
{
    public ?int $buildVersion = null;
    public ?int $majorVersion = null;
    public ?int $minorVersion = null;
    public ?string $releaseName = null;
}

/** Request payload for NoFrixionVersion#load. */
class NoFrixionVersionLoadMatch
{
    public ?int $buildVersion = null;
    public ?int $majorVersion = null;
    public ?int $minorVersion = null;
    public ?string $releaseName = null;
}

/** OpenBanking entity data model. */
class OpenBanking
{
}

/** Request payload for OpenBanking#create. */
class OpenBankingCreateData
{
    public string $account_id;
}

/** Request payload for OpenBanking#remove. */
class OpenBankingRemoveMatch
{
    public ?string $email = null;
    public ?string $merchant_id = null;
    public ?string $account_id = null;
}

/** Payeeverification entity data model. */
class Payeeverification
{
    public string $accountName;
    public ?string $accountNumber = null;
    public string $iban;
    public ?string $payeeVerifiedAccountName = null;
    public ?string $result = null;
    public ?string $secondaryIdentification = null;
    public ?string $sortCode = null;
}

/** Request payload for Payeeverification#create. */
class PayeeverificationCreateData
{
    public string $accountName;
    public ?string $accountNumber = null;
    public string $iban;
    public ?string $payeeVerifiedAccountName = null;
    public ?string $result = null;
    public ?string $secondaryIdentification = null;
    public ?string $sortCode = null;
}

/** Payment entity data model. */
class Payment
{
    public ?array $addresses = null;
    public ?float $amount = null;
    public ?float $amountPending = null;
    public ?float $amountReceived = null;
    public ?float $amountRefunded = null;
    public ?bool $autoSendReceipt = null;
    public ?string $baseOriginUrl = null;
    public ?string $callbackUrl = null;
    public ?bool $cardAuthorizeOnly = null;
    public ?bool $cardCreateToken = null;
    public ?string $cardCreateTokenMode = null;
    public ?bool $cardIgnoreCVN = null;
    public ?bool $cardNoPayerAuthentication = null;
    public ?string $cardProcessorMerchantID = null;
    public ?string $cardStripePaymentIntentID = null;
    public ?string $cardStripePaymentIntentSecret = null;
    public ?bool $cardTransmitRawDetails = null;
    public array $createdByUser;
    public ?string $currency = null;
    public ?array $customFields = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerID = null;
    public ?string $customerName = null;
    public ?string $description = null;
    public ?array $destinationAccount = null;
    public ?array $directDebitPayment = null;
    public ?string $dueDate = null;
    public ?array $events = null;
    public ?string $failureCallbackUrl = null;
    public ?array $fieldDisplaySettings = null;
    public ?string $formattedAmount = null;
    public ?string $hostedPayCheckoutUrl = null;
    public ?string $id = null;
    public ?bool $ignoreAddressVerification = null;
    public ?string $inserted = null;
    public ?string $insertedSortable = null;
    public ?bool $isArchived = null;
    public ?string $jwk = null;
    public ?string $lastUpdated = null;
    public ?string $lightningInvoice = null;
    public ?string $lightningInvoiceExpiresAt = null;
    public ?string $merchantDirectDebitMandateID = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $notificationEmailAddresses = null;
    public ?array $notificationRoleIDs = null;
    public ?string $orderID = null;
    public ?string $partialPaymentMethod = null;
    public ?string $partialPaymentSteps = null;
    public ?array $paymentAttempts = null;
    public ?array $paymentMethods = null;
    public ?string $paymentProcessor = null;
    public ?string $payrunID = null;
    public ?string $pispAccountID = null;
    public ?string $priorityBankID = null;
    public ?array $result = null;
    public ?int $sandboxSettleDelayInSeconds = null;
    public ?array $shippingAddress = null;
    public ?string $shippingAddressCity = null;
    public ?string $shippingAddressCountryCode = null;
    public ?string $shippingAddressCounty = null;
    public ?string $shippingAddressLine1 = null;
    public ?string $shippingAddressLine2 = null;
    public ?string $shippingAddressPostCode = null;
    public ?string $shippingEmail = null;
    public ?string $shippingFirstName = null;
    public ?string $shippingLastName = null;
    public ?string $shippingPhone = null;
    public ?string $status = null;
    public ?string $successWebHookUrl = null;
    public ?array $tagIds = null;
    public ?array $tags = null;
    public ?string $title = null;
    public ?array $tokenisedCards = null;
    public ?array $transactions = null;
    public ?bool $useHostedPaymentPage = null;
}

/** Request payload for Payment#load. */
class PaymentLoadMatch
{
    public ?string $id = null;
    public ?string $order_id = null;
}

/** Request payload for Payment#create. */
class PaymentCreateData
{
    public ?array $addresses = null;
    public ?float $amount = null;
    public ?float $amountPending = null;
    public ?float $amountReceived = null;
    public ?float $amountRefunded = null;
    public ?bool $autoSendReceipt = null;
    public ?string $baseOriginUrl = null;
    public ?string $callbackUrl = null;
    public ?bool $cardAuthorizeOnly = null;
    public ?bool $cardCreateToken = null;
    public ?string $cardCreateTokenMode = null;
    public ?bool $cardIgnoreCVN = null;
    public ?bool $cardNoPayerAuthentication = null;
    public ?string $cardProcessorMerchantID = null;
    public ?string $cardStripePaymentIntentID = null;
    public ?string $cardStripePaymentIntentSecret = null;
    public ?bool $cardTransmitRawDetails = null;
    public array $createdByUser;
    public ?string $currency = null;
    public ?array $customFields = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerID = null;
    public ?string $customerName = null;
    public ?string $description = null;
    public ?array $destinationAccount = null;
    public ?array $directDebitPayment = null;
    public ?string $dueDate = null;
    public ?array $events = null;
    public ?string $failureCallbackUrl = null;
    public ?array $fieldDisplaySettings = null;
    public ?string $formattedAmount = null;
    public ?string $hostedPayCheckoutUrl = null;
    public ?string $id = null;
    public ?bool $ignoreAddressVerification = null;
    public ?string $inserted = null;
    public ?string $insertedSortable = null;
    public ?bool $isArchived = null;
    public ?string $jwk = null;
    public ?string $lastUpdated = null;
    public ?string $lightningInvoice = null;
    public ?string $lightningInvoiceExpiresAt = null;
    public ?string $merchantDirectDebitMandateID = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $notificationEmailAddresses = null;
    public ?array $notificationRoleIDs = null;
    public ?string $orderID = null;
    public ?string $partialPaymentMethod = null;
    public ?string $partialPaymentSteps = null;
    public ?array $paymentAttempts = null;
    public ?array $paymentMethods = null;
    public ?string $paymentProcessor = null;
    public ?string $payrunID = null;
    public ?string $pispAccountID = null;
    public ?string $priorityBankID = null;
    public ?array $result = null;
    public ?int $sandboxSettleDelayInSeconds = null;
    public ?array $shippingAddress = null;
    public ?string $shippingAddressCity = null;
    public ?string $shippingAddressCountryCode = null;
    public ?string $shippingAddressCounty = null;
    public ?string $shippingAddressLine1 = null;
    public ?string $shippingAddressLine2 = null;
    public ?string $shippingAddressPostCode = null;
    public ?string $shippingEmail = null;
    public ?string $shippingFirstName = null;
    public ?string $shippingLastName = null;
    public ?string $shippingPhone = null;
    public ?string $status = null;
    public ?string $successWebHookUrl = null;
    public ?array $tagIds = null;
    public ?array $tags = null;
    public ?string $title = null;
    public ?array $tokenisedCards = null;
    public ?array $transactions = null;
    public ?bool $useHostedPaymentPage = null;
}

/** Request payload for Payment#update. */
class PaymentUpdateData
{
    public string $id;
    public ?array $addresses = null;
    public ?float $amount = null;
    public ?float $amountPending = null;
    public ?float $amountReceived = null;
    public ?float $amountRefunded = null;
    public ?bool $autoSendReceipt = null;
    public ?string $baseOriginUrl = null;
    public ?string $callbackUrl = null;
    public ?bool $cardAuthorizeOnly = null;
    public ?bool $cardCreateToken = null;
    public ?string $cardCreateTokenMode = null;
    public ?bool $cardIgnoreCVN = null;
    public ?bool $cardNoPayerAuthentication = null;
    public ?string $cardProcessorMerchantID = null;
    public ?string $cardStripePaymentIntentID = null;
    public ?string $cardStripePaymentIntentSecret = null;
    public ?bool $cardTransmitRawDetails = null;
    public ?array $createdByUser = null;
    public ?string $currency = null;
    public ?array $customFields = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerID = null;
    public ?string $customerName = null;
    public ?string $description = null;
    public ?array $destinationAccount = null;
    public ?array $directDebitPayment = null;
    public ?string $dueDate = null;
    public ?array $events = null;
    public ?string $failureCallbackUrl = null;
    public ?array $fieldDisplaySettings = null;
    public ?string $formattedAmount = null;
    public ?string $hostedPayCheckoutUrl = null;
    public ?bool $ignoreAddressVerification = null;
    public ?string $inserted = null;
    public ?string $insertedSortable = null;
    public ?bool $isArchived = null;
    public ?string $jwk = null;
    public ?string $lastUpdated = null;
    public ?string $lightningInvoice = null;
    public ?string $lightningInvoiceExpiresAt = null;
    public ?string $merchantDirectDebitMandateID = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $notificationEmailAddresses = null;
    public ?array $notificationRoleIDs = null;
    public ?string $orderID = null;
    public ?string $partialPaymentMethod = null;
    public ?string $partialPaymentSteps = null;
    public ?array $paymentAttempts = null;
    public ?array $paymentMethods = null;
    public ?string $paymentProcessor = null;
    public ?string $payrunID = null;
    public ?string $pispAccountID = null;
    public ?string $priorityBankID = null;
    public ?array $result = null;
    public ?int $sandboxSettleDelayInSeconds = null;
    public ?array $shippingAddress = null;
    public ?string $shippingAddressCity = null;
    public ?string $shippingAddressCountryCode = null;
    public ?string $shippingAddressCounty = null;
    public ?string $shippingAddressLine1 = null;
    public ?string $shippingAddressLine2 = null;
    public ?string $shippingAddressPostCode = null;
    public ?string $shippingEmail = null;
    public ?string $shippingFirstName = null;
    public ?string $shippingLastName = null;
    public ?string $shippingPhone = null;
    public ?string $status = null;
    public ?string $successWebHookUrl = null;
    public ?array $tagIds = null;
    public ?array $tags = null;
    public ?string $title = null;
    public ?array $tokenisedCards = null;
    public ?array $transactions = null;
    public ?bool $useHostedPaymentPage = null;
}

/** PaymentAccount entity data model. */
class PaymentAccount
{
    public ?string $accountName = null;
    public ?string $accountSupplierName = null;
    public ?float $availableBalance = null;
    public ?int $availableBalanceMinorUnits = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $bankName = null;
    public ?string $consentID = null;
    public array $createdBy;
    public ?string $createdByDisplayName = null;
    public ?string $currency = null;
    public ?string $defaultPaymentRail = null;
    public ?string $displayName = null;
    public ?string $expiryDate = null;
    public ?string $externalAccountIcon = null;
    public ?string $id = null;
    public array $identifier;
    public ?string $inserted = null;
    public ?bool $isArchived = null;
    public ?bool $isConnectedAccount = null;
    public ?bool $isDefault = null;
    public ?bool $isTrustAccount = null;
    public ?bool $isVirtual = null;
    public ?array $lastTransaction = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public ?string $physicalAccountID = null;
    public ?array $rules = null;
    public ?float $submittedPayoutsBalance = null;
    public ?int $submittedPayoutsBalanceMinorUnits = null;
    public ?string $summary = null;
    public ?string $supplierSepaInstantStatus = null;
    public ?string $xeroBankFeedConnectionStatus = null;
    public ?string $xeroBankFeedLastSyncedAt = null;
    public ?string $xeroBankFeedSyncLastFailedAt = null;
    public ?string $xeroBankFeedSyncLastFailureReason = null;
    public ?string $xeroBankFeedSyncStatus = null;
    public ?int $xeroUnsynchronisedTransactionsCount = null;
}

/** Request payload for PaymentAccount#list. */
class PaymentAccountListMatch
{
    public ?string $account_id = null;
}

/** PaymentAccountMinimal entity data model. */
class PaymentAccountMinimal
{
    public ?string $accountName = null;
    public ?float $availableBalance = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $currency = null;
    public ?string $id = null;
    public array $identifier;
    public ?bool $isArchived = null;
    public ?bool $isConnectedAccount = null;
    public ?string $merchantID = null;
    public ?float $submittedPayoutsBalance = null;
}

/** Request payload for PaymentAccountMinimal#list. */
class PaymentAccountMinimalListMatch
{
    public ?string $accountName = null;
    public ?float $availableBalance = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $currency = null;
    public ?string $id = null;
    public ?array $identifier = null;
    public ?bool $isArchived = null;
    public ?bool $isConnectedAccount = null;
    public ?string $merchantID = null;
    public ?float $submittedPayoutsBalance = null;
}

/** PaymentInitiation entity data model. */
class PaymentInitiation
{
    public ?string $paymentInitiationID = null;
    public ?string $paymentRequestCallbackUrl = null;
    public ?string $paymentRequestID = null;
    public ?string $redirectUrl = null;
    public ?string $responseType = null;
    public ?string $specificErrorMessage = null;
}

/** Request payload for PaymentInitiation#create. */
class PaymentInitiationCreateData
{
    public string $paymentrequest_id;
    public ?string $paymentInitiationID = null;
    public ?string $paymentRequestCallbackUrl = null;
    public ?string $paymentRequestID = null;
    public ?string $redirectUrl = null;
    public ?string $responseType = null;
    public ?string $specificErrorMessage = null;
}

/** PaymentRequest entity data model. */
class PaymentRequest
{
    public ?array $addresses = null;
    public ?float $amount = null;
    public ?float $amountPending = null;
    public ?float $amountReceived = null;
    public ?float $amountRefunded = null;
    public ?bool $autoSendReceipt = null;
    public ?string $baseOriginUrl = null;
    public ?string $callbackUrl = null;
    public ?bool $cardAuthorizeOnly = null;
    public ?bool $cardCreateToken = null;
    public ?string $cardCreateTokenMode = null;
    public ?bool $cardIgnoreCVN = null;
    public ?string $cardProcessorMerchantID = null;
    public ?string $cardStripePaymentIntentID = null;
    public ?string $cardStripePaymentIntentSecret = null;
    public array $createdByUser;
    public ?string $currency = null;
    public ?array $customFields = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerID = null;
    public ?string $customerName = null;
    public ?string $description = null;
    public ?array $destinationAccount = null;
    public ?array $directDebitPayment = null;
    public ?bool $doSimulateSettlementFailure = null;
    public ?string $dueDate = null;
    public ?string $errorDescription = null;
    public ?array $events = null;
    public ?array $failedPaymentRequests = null;
    public ?string $failureCallbackUrl = null;
    public ?array $fieldDisplaySettings = null;
    public ?string $formattedAmount = null;
    public ?string $hostedPayCheckoutUrl = null;
    public ?string $id = null;
    public ?bool $ignoreAddressVerification = null;
    public ?string $inserted = null;
    public ?string $insertedSortable = null;
    public ?string $institution = null;
    public ?bool $isArchived = null;
    public ?string $jwk = null;
    public ?string $lastUpdated = null;
    public ?string $lightningInvoice = null;
    public ?string $lightningInvoiceExpiresAt = null;
    public ?string $merchantDirectDebitMandateID = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $notificationEmailAddresses = null;
    public ?array $notificationRoleIDs = null;
    public ?string $orderID = null;
    public ?string $partialPaymentMethod = null;
    public ?string $partialPaymentSteps = null;
    public ?array $paymentAttempts = null;
    public ?string $paymentInitiationID = null;
    public ?array $paymentMethods = null;
    public ?string $paymentProcessor = null;
    public ?array $paymentRequests = null;
    public ?string $payrunID = null;
    public ?string $pispAccountID = null;
    public ?string $priorityBankID = null;
    public ?array $result = null;
    public ?int $sandboxSettleDelayInSeconds = null;
    public ?array $shippingAddress = null;
    public ?string $status = null;
    public ?string $successWebHookUrl = null;
    public ?array $tags = null;
    public ?string $title = null;
    public ?array $tokenisedCards = null;
    public ?array $transactions = null;
    public ?bool $useHostedPaymentPage = null;
}

/** Request payload for PaymentRequest#load. */
class PaymentRequestLoadMatch
{
    public ?string $paymentrequest_id = null;
}

/** Request payload for PaymentRequest#list. */
class PaymentRequestListMatch
{
    public ?array $addresses = null;
    public ?float $amount = null;
    public ?float $amountPending = null;
    public ?float $amountReceived = null;
    public ?float $amountRefunded = null;
    public ?bool $autoSendReceipt = null;
    public ?string $baseOriginUrl = null;
    public ?string $callbackUrl = null;
    public ?bool $cardAuthorizeOnly = null;
    public ?bool $cardCreateToken = null;
    public ?string $cardCreateTokenMode = null;
    public ?bool $cardIgnoreCVN = null;
    public ?string $cardProcessorMerchantID = null;
    public ?string $cardStripePaymentIntentID = null;
    public ?string $cardStripePaymentIntentSecret = null;
    public ?array $createdByUser = null;
    public ?string $currency = null;
    public ?array $customFields = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerID = null;
    public ?string $customerName = null;
    public ?string $description = null;
    public ?array $destinationAccount = null;
    public ?array $directDebitPayment = null;
    public ?bool $doSimulateSettlementFailure = null;
    public ?string $dueDate = null;
    public ?string $errorDescription = null;
    public ?array $events = null;
    public ?array $failedPaymentRequests = null;
    public ?string $failureCallbackUrl = null;
    public ?array $fieldDisplaySettings = null;
    public ?string $formattedAmount = null;
    public ?string $hostedPayCheckoutUrl = null;
    public ?string $id = null;
    public ?bool $ignoreAddressVerification = null;
    public ?string $inserted = null;
    public ?string $insertedSortable = null;
    public ?string $institution = null;
    public ?bool $isArchived = null;
    public ?string $jwk = null;
    public ?string $lastUpdated = null;
    public ?string $lightningInvoice = null;
    public ?string $lightningInvoiceExpiresAt = null;
    public ?string $merchantDirectDebitMandateID = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $notificationEmailAddresses = null;
    public ?array $notificationRoleIDs = null;
    public ?string $orderID = null;
    public ?string $partialPaymentMethod = null;
    public ?string $partialPaymentSteps = null;
    public ?array $paymentAttempts = null;
    public ?string $paymentInitiationID = null;
    public ?array $paymentMethods = null;
    public ?string $paymentProcessor = null;
    public ?array $paymentRequests = null;
    public ?string $payrunID = null;
    public ?string $pispAccountID = null;
    public ?string $priorityBankID = null;
    public ?array $result = null;
    public ?int $sandboxSettleDelayInSeconds = null;
    public ?array $shippingAddress = null;
    public ?string $status = null;
    public ?string $successWebHookUrl = null;
    public ?array $tags = null;
    public ?string $title = null;
    public ?array $tokenisedCards = null;
    public ?array $transactions = null;
    public ?bool $useHostedPaymentPage = null;
}

/** Request payload for PaymentRequest#create. */
class PaymentRequestCreateData
{
    public ?string $paymentrequest_id = null;
    public ?array $addresses = null;
    public ?float $amount = null;
    public ?float $amountPending = null;
    public ?float $amountReceived = null;
    public ?float $amountRefunded = null;
    public ?bool $autoSendReceipt = null;
    public ?string $baseOriginUrl = null;
    public ?string $callbackUrl = null;
    public ?bool $cardAuthorizeOnly = null;
    public ?bool $cardCreateToken = null;
    public ?string $cardCreateTokenMode = null;
    public ?bool $cardIgnoreCVN = null;
    public ?string $cardProcessorMerchantID = null;
    public ?string $cardStripePaymentIntentID = null;
    public ?string $cardStripePaymentIntentSecret = null;
    public array $createdByUser;
    public ?string $currency = null;
    public ?array $customFields = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerID = null;
    public ?string $customerName = null;
    public ?string $description = null;
    public ?array $destinationAccount = null;
    public ?array $directDebitPayment = null;
    public ?bool $doSimulateSettlementFailure = null;
    public ?string $dueDate = null;
    public ?string $errorDescription = null;
    public ?array $events = null;
    public ?array $failedPaymentRequests = null;
    public ?string $failureCallbackUrl = null;
    public ?array $fieldDisplaySettings = null;
    public ?string $formattedAmount = null;
    public ?string $hostedPayCheckoutUrl = null;
    public ?string $id = null;
    public ?bool $ignoreAddressVerification = null;
    public ?string $inserted = null;
    public ?string $insertedSortable = null;
    public ?string $institution = null;
    public ?bool $isArchived = null;
    public ?string $jwk = null;
    public ?string $lastUpdated = null;
    public ?string $lightningInvoice = null;
    public ?string $lightningInvoiceExpiresAt = null;
    public ?string $merchantDirectDebitMandateID = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $notificationEmailAddresses = null;
    public ?array $notificationRoleIDs = null;
    public ?string $orderID = null;
    public ?string $partialPaymentMethod = null;
    public ?string $partialPaymentSteps = null;
    public ?array $paymentAttempts = null;
    public ?string $paymentInitiationID = null;
    public ?array $paymentMethods = null;
    public ?string $paymentProcessor = null;
    public ?array $paymentRequests = null;
    public ?string $payrunID = null;
    public ?string $pispAccountID = null;
    public ?string $priorityBankID = null;
    public ?array $result = null;
    public ?int $sandboxSettleDelayInSeconds = null;
    public ?array $shippingAddress = null;
    public ?string $status = null;
    public ?string $successWebHookUrl = null;
    public ?array $tags = null;
    public ?string $title = null;
    public ?array $tokenisedCards = null;
    public ?array $transactions = null;
    public ?bool $useHostedPaymentPage = null;
}

/** Request payload for PaymentRequest#update. */
class PaymentRequestUpdateData
{
    public string $paymentrequest_id;
    public ?array $addresses = null;
    public ?float $amount = null;
    public ?float $amountPending = null;
    public ?float $amountReceived = null;
    public ?float $amountRefunded = null;
    public ?bool $autoSendReceipt = null;
    public ?string $baseOriginUrl = null;
    public ?string $callbackUrl = null;
    public ?bool $cardAuthorizeOnly = null;
    public ?bool $cardCreateToken = null;
    public ?string $cardCreateTokenMode = null;
    public ?bool $cardIgnoreCVN = null;
    public ?string $cardProcessorMerchantID = null;
    public ?string $cardStripePaymentIntentID = null;
    public ?string $cardStripePaymentIntentSecret = null;
    public ?array $createdByUser = null;
    public ?string $currency = null;
    public ?array $customFields = null;
    public ?string $customerEmailAddress = null;
    public ?string $customerID = null;
    public ?string $customerName = null;
    public ?string $description = null;
    public ?array $destinationAccount = null;
    public ?array $directDebitPayment = null;
    public ?bool $doSimulateSettlementFailure = null;
    public ?string $dueDate = null;
    public ?string $errorDescription = null;
    public ?array $events = null;
    public ?array $failedPaymentRequests = null;
    public ?string $failureCallbackUrl = null;
    public ?array $fieldDisplaySettings = null;
    public ?string $formattedAmount = null;
    public ?string $hostedPayCheckoutUrl = null;
    public ?string $id = null;
    public ?bool $ignoreAddressVerification = null;
    public ?string $inserted = null;
    public ?string $insertedSortable = null;
    public ?string $institution = null;
    public ?bool $isArchived = null;
    public ?string $jwk = null;
    public ?string $lastUpdated = null;
    public ?string $lightningInvoice = null;
    public ?string $lightningInvoiceExpiresAt = null;
    public ?string $merchantDirectDebitMandateID = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $notificationEmailAddresses = null;
    public ?array $notificationRoleIDs = null;
    public ?string $orderID = null;
    public ?string $partialPaymentMethod = null;
    public ?string $partialPaymentSteps = null;
    public ?array $paymentAttempts = null;
    public ?string $paymentInitiationID = null;
    public ?array $paymentMethods = null;
    public ?string $paymentProcessor = null;
    public ?array $paymentRequests = null;
    public ?string $payrunID = null;
    public ?string $pispAccountID = null;
    public ?string $priorityBankID = null;
    public ?array $result = null;
    public ?int $sandboxSettleDelayInSeconds = null;
    public ?array $shippingAddress = null;
    public ?string $status = null;
    public ?string $successWebHookUrl = null;
    public ?array $tags = null;
    public ?string $title = null;
    public ?array $tokenisedCards = null;
    public ?array $transactions = null;
    public ?bool $useHostedPaymentPage = null;
}

/** Request payload for PaymentRequest#remove. */
class PaymentRequestRemoveMatch
{
    public string $id;
}

/** PaymentRequestEvent entity data model. */
class PaymentRequestEvent
{
    public float $amount;
    public ?string $applePayTransactionID = null;
    public ?string $cardAuthorizationResponseID = null;
    public ?int $cardExpiryMonth = null;
    public ?int $cardExpiryYear = null;
    public ?string $cardIssuer = null;
    public ?string $cardIssuerCountry = null;
    public ?string $cardLastFourDigits = null;
    public ?string $cardRequestID = null;
    public ?string $cardScheme = null;
    public ?string $cardTokenCustomerID = null;
    public ?string $cardTransactionID = null;
    public ?string $currency = null;
    public ?string $directDebitPaymentID = null;
    public ?string $directDebitPaymentReference = null;
    public ?string $drirectDebitMandateID = null;
    public ?string $errorMessage = null;
    public ?string $errorReason = null;
    public ?string $eventType = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $lightningInvoice = null;
    public ?string $lightningRHash = null;
    public ?string $originUrl = null;
    public ?string $paymentMethodType = null;
    public ?string $paymentProcessorName = null;
    public ?string $paymentRequestID = null;
    public ?string $pispBankStatus = null;
    public ?string $pispPaymentInitiationID = null;
    public ?string $pispPaymentInstitutionName = null;
    public ?string $pispPaymentServiceProviderID = null;
    public ?string $pispRedirectUrl = null;
    public ?string $reconciledTransactionID = null;
    public ?string $refundPayoutID = null;
    public ?string $status = null;
    public ?string $walletName = null;
}

/** Request payload for PaymentRequestEvent#list. */
class PaymentRequestEventListMatch
{
    public string $paymentrequest_id;
}

/** PaymentRequestMetric entity data model. */
class PaymentRequestMetric
{
}

/** Request payload for PaymentRequestMetric#load. */
class PaymentRequestMetricLoadMatch
{
}

/** PaymentRequestMinimal entity data model. */
class PaymentRequestMinimal
{
    public ?float $amount = null;
    public ?float $amountPending = null;
    public ?float $amountReceived = null;
    public ?float $amountRefunded = null;
    public ?string $callbackUrl = null;
    public ?string $cardStripePaymentIntentSecret = null;
    public ?string $countryCode = null;
    public ?string $currency = null;
    public ?array $customFieldsToDisplay = null;
    public ?string $description = null;
    public ?string $dueDate = null;
    public ?array $fieldDisplaySettings = null;
    public ?string $googlePayMerchantID = null;
    public ?string $id = null;
    public ?string $jwk = null;
    public ?string $merchantID = null;
    public ?string $merchantLogoUrlPng = null;
    public ?string $merchantLogoUrlSvg = null;
    public ?string $merchantName = null;
    public ?string $merchantShortName = null;
    public ?string $partialPaymentMethod = null;
    public ?array $paymentAttempts = null;
    public ?array $paymentMethodsList = null;
    public ?string $paymentProcessor = null;
    public ?string $paymentProcessorKey = null;
    public ?string $pispError = null;
    public ?string $priorityBankID = null;
    public ?string $status = null;
    public ?string $stripeAccountID = null;
    public ?string $title = null;
}

/** Request payload for PaymentRequestMinimal#list. */
class PaymentRequestMinimalListMatch
{
    public string $paymentrequest_id;
}

/** PaymentRequestResult entity data model. */
class PaymentRequestResult
{
    public ?float $amount = null;
    public ?float $amountPending = null;
    public ?float $amountReceived = null;
    public ?float $amountRefunded = null;
    public ?string $currency = null;
    public ?string $customerID = null;
    public ?string $paymentRequestID = null;
    public ?array $payments = null;
    public ?array $pispAuthorizations = null;
    public ?float $requestedAmount = null;
    public ?string $result = null;
}

/** Request payload for PaymentRequestResult#list. */
class PaymentRequestResultListMatch
{
    public string $paymentrequest_id;
}

/** Payout entity data model. */
class Payout
{
    public ?string $accountID = null;
    public ?bool $allowIncomplete = null;
    public ?float $amount = null;
    public ?int $amountMinorUnits = null;
    public ?string $approvePayoutUrl = null;
    public ?string $approverID = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?string $batchPayoutID = null;
    public array $beneficiary;
    public ?string $beneficiaryID = null;
    public ?bool $canAuthorise = null;
    public ?bool $canProcess = null;
    public ?bool $canUpdate = null;
    public ?string $chargeBearer = null;
    public ?string $createdBy = null;
    public ?string $createdByEmailAddress = null;
    public ?string $currency = null;
    public ?string $currentUserID = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $documents = null;
    public ?array $events = null;
    public ?array $failedPayouts = null;
    public ?string $formattedAmount = null;
    public ?string $formattedFxDestinationAmount = null;
    public ?string $formattedSchedule = null;
    public ?string $formattedScheduleDayOnly = null;
    public ?string $formattedSourceAccountAvailableBalance = null;
    public ?float $fxDestinationAmount = null;
    public ?int $fxDestinationAmountMinorUnits = null;
    public ?string $fxDestinationCurrency = null;
    public ?string $fxQuoteExpiresAt = null;
    public ?string $fxQuoteID = null;
    public ?float $fxRate = null;
    public ?bool $fxUseDestinationAmount = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoiceID = null;
    public ?bool $isArchived = null;
    public ?bool $isFailed = null;
    public ?bool $isSettled = null;
    public ?bool $isSubmitted = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $nonce = null;
    public ?string $paymentProcessor = null;
    public ?string $paymentRail = null;
    public ?array $payouts = null;
    public ?string $payrunID = null;
    public ?string $payrunName = null;
    public ?string $reason = null;
    public ?array $rule = null;
    public ?string $scheduleDate = null;
    public ?bool $scheduled = null;
    public ?float $sourceAccountAvailableBalance = null;
    public ?int $sourceAccountAvailableBalanceMinorUnits = null;
    public ?string $sourceAccountBic = null;
    public ?string $sourceAccountCurrency = null;
    public ?string $sourceAccountIban = null;
    public array $sourceAccountIdentifier;
    public ?string $sourceAccountName = null;
    public ?string $sourceAccountNumber = null;
    public ?string $sourceAccountSortcode = null;
    public ?string $status = null;
    public ?array $tagIds = null;
    public ?array $tags = null;
    public ?string $theirReference = null;
    public ?string $topupPayrunID = null;
    public ?float $transactedAmount = null;
    public ?float $transactedFxAmount = null;
    public ?float $transactedFxRate = null;
    public ?string $type = null;
    public ?string $userID = null;
    public ?string $yourReference = null;
}

/** Request payload for Payout#load. */
class PayoutLoadMatch
{
    public ?float $amount = null;
    public ?string $destination = null;
    public ?string $source = null;
    public ?string $id = null;
}

/** Request payload for Payout#list. */
class PayoutListMatch
{
    public ?string $account_id = null;
    public ?string $merchant_id = null;
}

/** Request payload for Payout#create. */
class PayoutCreateData
{
    public ?string $id = null;
    public ?string $accountID = null;
    public ?bool $allowIncomplete = null;
    public ?float $amount = null;
    public ?int $amountMinorUnits = null;
    public ?string $approvePayoutUrl = null;
    public ?string $approverID = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?string $batchPayoutID = null;
    public array $beneficiary;
    public ?string $beneficiaryID = null;
    public ?bool $canAuthorise = null;
    public ?bool $canProcess = null;
    public ?bool $canUpdate = null;
    public ?string $chargeBearer = null;
    public ?string $createdBy = null;
    public ?string $createdByEmailAddress = null;
    public ?string $currency = null;
    public ?string $currentUserID = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $documents = null;
    public ?array $events = null;
    public ?array $failedPayouts = null;
    public ?string $formattedAmount = null;
    public ?string $formattedFxDestinationAmount = null;
    public ?string $formattedSchedule = null;
    public ?string $formattedScheduleDayOnly = null;
    public ?string $formattedSourceAccountAvailableBalance = null;
    public ?float $fxDestinationAmount = null;
    public ?int $fxDestinationAmountMinorUnits = null;
    public ?string $fxDestinationCurrency = null;
    public ?string $fxQuoteExpiresAt = null;
    public ?string $fxQuoteID = null;
    public ?float $fxRate = null;
    public ?bool $fxUseDestinationAmount = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $inserted = null;
    public ?string $invoiceID = null;
    public ?bool $isArchived = null;
    public ?bool $isFailed = null;
    public ?bool $isSettled = null;
    public ?bool $isSubmitted = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $nonce = null;
    public ?string $paymentProcessor = null;
    public ?string $paymentRail = null;
    public ?array $payouts = null;
    public ?string $payrunID = null;
    public ?string $payrunName = null;
    public ?string $reason = null;
    public ?array $rule = null;
    public ?string $scheduleDate = null;
    public ?bool $scheduled = null;
    public ?float $sourceAccountAvailableBalance = null;
    public ?int $sourceAccountAvailableBalanceMinorUnits = null;
    public ?string $sourceAccountBic = null;
    public ?string $sourceAccountCurrency = null;
    public ?string $sourceAccountIban = null;
    public array $sourceAccountIdentifier;
    public ?string $sourceAccountName = null;
    public ?string $sourceAccountNumber = null;
    public ?string $sourceAccountSortcode = null;
    public ?string $status = null;
    public ?array $tagIds = null;
    public ?array $tags = null;
    public ?string $theirReference = null;
    public ?string $topupPayrunID = null;
    public ?float $transactedAmount = null;
    public ?float $transactedFxAmount = null;
    public ?float $transactedFxRate = null;
    public ?string $type = null;
    public ?string $userID = null;
    public ?string $yourReference = null;
}

/** Request payload for Payout#update. */
class PayoutUpdateData
{
    public string $id;
    public ?string $accountID = null;
    public ?bool $allowIncomplete = null;
    public ?float $amount = null;
    public ?int $amountMinorUnits = null;
    public ?string $approvePayoutUrl = null;
    public ?string $approverID = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?string $batchPayoutID = null;
    public ?array $beneficiary = null;
    public ?string $beneficiaryID = null;
    public ?bool $canAuthorise = null;
    public ?bool $canProcess = null;
    public ?bool $canUpdate = null;
    public ?string $chargeBearer = null;
    public ?string $createdBy = null;
    public ?string $createdByEmailAddress = null;
    public ?string $currency = null;
    public ?string $currentUserID = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $documents = null;
    public ?array $events = null;
    public ?array $failedPayouts = null;
    public ?string $formattedAmount = null;
    public ?string $formattedFxDestinationAmount = null;
    public ?string $formattedSchedule = null;
    public ?string $formattedScheduleDayOnly = null;
    public ?string $formattedSourceAccountAvailableBalance = null;
    public ?float $fxDestinationAmount = null;
    public ?int $fxDestinationAmountMinorUnits = null;
    public ?string $fxDestinationCurrency = null;
    public ?string $fxQuoteExpiresAt = null;
    public ?string $fxQuoteID = null;
    public ?float $fxRate = null;
    public ?bool $fxUseDestinationAmount = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $inserted = null;
    public ?string $invoiceID = null;
    public ?bool $isArchived = null;
    public ?bool $isFailed = null;
    public ?bool $isSettled = null;
    public ?bool $isSubmitted = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $nonce = null;
    public ?string $paymentProcessor = null;
    public ?string $paymentRail = null;
    public ?array $payouts = null;
    public ?string $payrunID = null;
    public ?string $payrunName = null;
    public ?string $reason = null;
    public ?array $rule = null;
    public ?string $scheduleDate = null;
    public ?bool $scheduled = null;
    public ?float $sourceAccountAvailableBalance = null;
    public ?int $sourceAccountAvailableBalanceMinorUnits = null;
    public ?string $sourceAccountBic = null;
    public ?string $sourceAccountCurrency = null;
    public ?string $sourceAccountIban = null;
    public ?array $sourceAccountIdentifier = null;
    public ?string $sourceAccountName = null;
    public ?string $sourceAccountNumber = null;
    public ?string $sourceAccountSortcode = null;
    public ?string $status = null;
    public ?array $tagIds = null;
    public ?array $tags = null;
    public ?string $theirReference = null;
    public ?string $topupPayrunID = null;
    public ?float $transactedAmount = null;
    public ?float $transactedFxAmount = null;
    public ?float $transactedFxRate = null;
    public ?string $type = null;
    public ?string $userID = null;
    public ?string $yourReference = null;
}

/** Request payload for Payout#remove. */
class PayoutRemoveMatch
{
    public string $id;
}

/** PayoutKeysetPage entity data model. */
class PayoutKeysetPage
{
    public ?string $accountID = null;
    public ?float $amount = null;
    public ?int $amountMinorUnits = null;
    public ?string $approvePayoutUrl = null;
    public ?string $approverID = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?string $batchPayoutID = null;
    public array $beneficiary;
    public ?bool $canAuthorise = null;
    public ?bool $canProcess = null;
    public ?bool $canUpdate = null;
    public ?string $chargeBearer = null;
    public ?string $createdBy = null;
    public ?string $createdByEmailAddress = null;
    public ?string $currency = null;
    public ?string $currentUserID = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $documents = null;
    public ?array $events = null;
    public ?string $formattedAmount = null;
    public ?string $formattedFxDestinationAmount = null;
    public ?string $formattedSchedule = null;
    public ?string $formattedScheduleDayOnly = null;
    public ?string $formattedSourceAccountAvailableBalance = null;
    public ?float $fxDestinationAmount = null;
    public ?int $fxDestinationAmountMinorUnits = null;
    public ?string $fxDestinationCurrency = null;
    public ?string $fxQuoteExpiresAt = null;
    public ?string $fxQuoteID = null;
    public ?float $fxRate = null;
    public ?bool $fxUseDestinationAmount = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoiceID = null;
    public ?bool $isArchived = null;
    public ?bool $isFailed = null;
    public ?bool $isSettled = null;
    public ?bool $isSubmitted = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantTokenDescription = null;
    public ?string $nonce = null;
    public ?string $paymentProcessor = null;
    public ?string $paymentRail = null;
    public ?string $payrunID = null;
    public ?string $payrunName = null;
    public ?array $rule = null;
    public ?string $scheduleDate = null;
    public ?bool $scheduled = null;
    public ?float $sourceAccountAvailableBalance = null;
    public ?int $sourceAccountAvailableBalanceMinorUnits = null;
    public ?string $sourceAccountBic = null;
    public ?string $sourceAccountCurrency = null;
    public ?string $sourceAccountIban = null;
    public array $sourceAccountIdentifier;
    public ?string $sourceAccountName = null;
    public ?string $sourceAccountNumber = null;
    public ?string $sourceAccountSortcode = null;
    public ?string $status = null;
    public ?array $tags = null;
    public ?string $theirReference = null;
    public ?string $topupPayrunID = null;
    public ?float $transactedAmount = null;
    public ?float $transactedFxAmount = null;
    public ?float $transactedFxRate = null;
    public ?string $type = null;
    public ?string $userID = null;
    public ?string $yourReference = null;
}

/** Request payload for PayoutKeysetPage#list. */
class PayoutKeysetPageListMatch
{
    public ?string $account_id = null;
    public ?string $merchant_id = null;
}

/** PayoutMetric entity data model. */
class PayoutMetric
{
}

/** Request payload for PayoutMetric#load. */
class PayoutMetricLoadMatch
{
}

/** Payrun entity data model. */
class Payrun
{
    public ?string $authorisationDate = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?string $batchPayoutID = null;
    public ?bool $canAuthorise = null;
    public ?bool $canDelete = null;
    public ?bool $canEdit = null;
    public ?array $events = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?array $invoices = null;
    public ?array $invoicesMinimal = null;
    public ?bool $isArchived = null;
    public ?string $lastUpdated = null;
    public array $lastUpdatedBy;
    public ?string $merchantID = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?string $notes = null;
    public ?array $payments = null;
    public ?array $payouts = null;
    public ?int $payoutsCount = null;
    public ?string $reason = null;
    public ?string $scheduleDate = null;
    public ?string $scheduledDate = null;
    public ?array $sourceAccounts = null;
    public ?string $status = null;
    public ?float $totalEur = null;
    public ?float $totalGbp = null;
    public ?float $totalUsd = null;
}

/** Request payload for Payrun#load. */
class PayrunLoadMatch
{
    public string $id;
}

/** Request payload for Payrun#list. */
class PayrunListMatch
{
    public ?string $authorisationDate = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?string $batchPayoutID = null;
    public ?bool $canAuthorise = null;
    public ?bool $canDelete = null;
    public ?bool $canEdit = null;
    public ?array $events = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?array $invoices = null;
    public ?array $invoicesMinimal = null;
    public ?bool $isArchived = null;
    public ?string $lastUpdated = null;
    public ?array $lastUpdatedBy = null;
    public ?string $merchantID = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?string $notes = null;
    public ?array $payments = null;
    public ?array $payouts = null;
    public ?int $payoutsCount = null;
    public ?string $reason = null;
    public ?string $scheduleDate = null;
    public ?string $scheduledDate = null;
    public ?array $sourceAccounts = null;
    public ?string $status = null;
    public ?float $totalEur = null;
    public ?float $totalGbp = null;
    public ?float $totalUsd = null;
}

/** Request payload for Payrun#create. */
class PayrunCreateData
{
    public string $id;
    public ?string $authorisationDate = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?string $batchPayoutID = null;
    public ?bool $canAuthorise = null;
    public ?bool $canDelete = null;
    public ?bool $canEdit = null;
    public ?array $events = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $inserted = null;
    public ?array $invoices = null;
    public ?array $invoicesMinimal = null;
    public ?bool $isArchived = null;
    public ?string $lastUpdated = null;
    public array $lastUpdatedBy;
    public ?string $merchantID = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?string $notes = null;
    public ?array $payments = null;
    public ?array $payouts = null;
    public ?int $payoutsCount = null;
    public ?string $reason = null;
    public ?string $scheduleDate = null;
    public ?string $scheduledDate = null;
    public ?array $sourceAccounts = null;
    public ?string $status = null;
    public ?float $totalEur = null;
    public ?float $totalGbp = null;
    public ?float $totalUsd = null;
}

/** Request payload for Payrun#update. */
class PayrunUpdateData
{
    public string $id;
    public ?string $authorisationDate = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?string $batchPayoutID = null;
    public ?bool $canAuthorise = null;
    public ?bool $canDelete = null;
    public ?bool $canEdit = null;
    public ?array $events = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $inserted = null;
    public ?array $invoices = null;
    public ?array $invoicesMinimal = null;
    public ?bool $isArchived = null;
    public ?string $lastUpdated = null;
    public ?array $lastUpdatedBy = null;
    public ?string $merchantID = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?string $notes = null;
    public ?array $payments = null;
    public ?array $payouts = null;
    public ?int $payoutsCount = null;
    public ?string $reason = null;
    public ?string $scheduleDate = null;
    public ?string $scheduledDate = null;
    public ?array $sourceAccounts = null;
    public ?string $status = null;
    public ?float $totalEur = null;
    public ?float $totalGbp = null;
    public ?float $totalUsd = null;
}

/** Request payload for Payrun#remove. */
class PayrunRemoveMatch
{
    public string $id;
}

/** Report entity data model. */
class Report
{
}

/** Request payload for Report#update. */
class ReportUpdateData
{
    public string $id;
}

/** ReportResult entity data model. */
class ReportResult
{
    public ?string $contentType = null;
    public ?string $contents = null;
    public ?string $lastCompletedAt = null;
    public ?string $merchantID = null;
    public ?string $reportName = null;
    public ?string $reportType = null;
    public ?int $statementNumber = null;
}

/** Request payload for ReportResult#load. */
class ReportResultLoadMatch
{
    public int $id;
    public string $report_id;
}

/** Role entity data model. */
class Role
{
    public ?array $failedRoles = null;
    public ?array $roles = null;
}

/** Request payload for Role#create. */
class RoleCreateData
{
    public string $merchant_id;
    public ?array $failedRoles = null;
    public ?array $roles = null;
}

/** Rule entity data model. */
class Rule
{
    public ?array $account = null;
    public ?string $accountID = null;
    public ?string $approveUrl = null;
    public ?string $approverID = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?bool $canAuthorise = null;
    public array $createdBy;
    public ?string $description = null;
    public ?string $endAt = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isDisabled = null;
    public ?string $lastExecutedAt = null;
    public ?string $lastRunAtTransactionDate = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $name = null;
    public string $nonce;
    public ?string $onApprovedWebHookUrl = null;
    public ?string $onExecutionErrorWebHookUrl = null;
    public ?string $onExecutionSuccessWebHookUrl = null;
    public ?string $startAt = null;
    public ?string $status = null;
    public ?array $sweepAction = null;
    public ?string $timeZoneId = null;
    public ?string $triggerCronExpression = null;
    public ?bool $triggerOnPayIn = null;
    public ?string $userID = null;
    public ?string $webHookSecret = null;
}

/** Request payload for Rule#load. */
class RuleLoadMatch
{
    public string $id;
}

/** Request payload for Rule#list. */
class RuleListMatch
{
    public ?array $account = null;
    public ?string $accountID = null;
    public ?string $approveUrl = null;
    public ?string $approverID = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?bool $canAuthorise = null;
    public ?array $createdBy = null;
    public ?string $description = null;
    public ?string $endAt = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isDisabled = null;
    public ?string $lastExecutedAt = null;
    public ?string $lastRunAtTransactionDate = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?string $onApprovedWebHookUrl = null;
    public ?string $onExecutionErrorWebHookUrl = null;
    public ?string $onExecutionSuccessWebHookUrl = null;
    public ?string $startAt = null;
    public ?string $status = null;
    public ?array $sweepAction = null;
    public ?string $timeZoneId = null;
    public ?string $triggerCronExpression = null;
    public ?bool $triggerOnPayIn = null;
    public ?string $userID = null;
    public ?string $webHookSecret = null;
}

/** Request payload for Rule#create. */
class RuleCreateData
{
    public ?array $account = null;
    public ?string $accountID = null;
    public ?string $approveUrl = null;
    public ?string $approverID = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?bool $canAuthorise = null;
    public array $createdBy;
    public ?string $description = null;
    public ?string $endAt = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isDisabled = null;
    public ?string $lastExecutedAt = null;
    public ?string $lastRunAtTransactionDate = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $name = null;
    public string $nonce;
    public ?string $onApprovedWebHookUrl = null;
    public ?string $onExecutionErrorWebHookUrl = null;
    public ?string $onExecutionSuccessWebHookUrl = null;
    public ?string $startAt = null;
    public ?string $status = null;
    public ?array $sweepAction = null;
    public ?string $timeZoneId = null;
    public ?string $triggerCronExpression = null;
    public ?bool $triggerOnPayIn = null;
    public ?string $userID = null;
    public ?string $webHookSecret = null;
}

/** Request payload for Rule#update. */
class RuleUpdateData
{
    public string $id;
    public ?array $account = null;
    public ?string $accountID = null;
    public ?string $approveUrl = null;
    public ?string $approverID = null;
    public ?array $authenticationMethods = null;
    public ?array $authorisations = null;
    public ?int $authorisersCompletedCount = null;
    public ?int $authorisersRequiredCount = null;
    public ?bool $canAuthorise = null;
    public ?array $createdBy = null;
    public ?string $description = null;
    public ?string $endAt = null;
    public ?bool $hasCurrentUserAuthorised = null;
    public ?string $inserted = null;
    public ?bool $isDisabled = null;
    public ?string $lastExecutedAt = null;
    public ?string $lastRunAtTransactionDate = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?string $onApprovedWebHookUrl = null;
    public ?string $onExecutionErrorWebHookUrl = null;
    public ?string $onExecutionSuccessWebHookUrl = null;
    public ?string $startAt = null;
    public ?string $status = null;
    public ?array $sweepAction = null;
    public ?string $timeZoneId = null;
    public ?string $triggerCronExpression = null;
    public ?bool $triggerOnPayIn = null;
    public ?string $userID = null;
    public ?string $webHookSecret = null;
}

/** Request payload for Rule#remove. */
class RuleRemoveMatch
{
    public string $id;
}

/** RuleEvent entity data model. */
class RuleEvent
{
    public ?string $errorMessage = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $isAuthoriseToEnable = null;
    public ?string $message = null;
    public ?string $rawResponse = null;
    public ?string $ruleEventType = null;
    public ?string $ruleID = null;
    public array $user;
}

/** Request payload for RuleEvent#list. */
class RuleEventListMatch
{
    public string $id;
}

/** Tag entity data model. */
class Tag
{
    public ?string $colourHex = null;
    public ?string $description = null;
    public ?string $id = null;
    public string $merchantID;
    public string $name;
}

/** Request payload for Tag#list. */
class TagListMatch
{
    public string $merchant_id;
}

/** Request payload for Tag#create. */
class TagCreateData
{
    public string $merchant_id;
    public ?string $colourHex = null;
    public ?string $description = null;
    public ?string $id = null;
    public string $merchantID;
    public string $name;
}

/** Token entity data model. */
class Token
{
}

/** Request payload for Token#create. */
class TokenCreateData
{
    public string $id;
}

/** Request payload for Token#remove. */
class TokenRemoveMatch
{
    public string $id;
}

/** Transaction entity data model. */
class Transaction
{
    public ?string $accountID = null;
    public ?string $accountName = null;
    public ?int $accountSequenceNumber = null;
    public ?array $addressDetails = null;
    public ?float $amount = null;
    public ?int $amountMinorUnits = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $bookingDateTime = null;
    public ?array $chargeDetails = null;
    public ?array $content = null;
    public ?array $counterparty = null;
    public ?string $counterpartySummary = null;
    public ?string $currency = null;
    public ?array $currencyExchange = null;
    public ?string $date = null;
    public ?string $description = null;
    public ?array $enrichment = null;
    public ?float $fxAmount = null;
    public ?string $fxCurrency = null;
    public ?float $fxRate = null;
    public array $grossAmount;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?array $isoBankTransactionCode = null;
    public ?array $merchant = null;
    public ?string $merchantID = null;
    public ?int $pageNumber = null;
    public ?int $pageSize = null;
    public array $payeeDetails;
    public array $payerDetails;
    public ?array $paymentRequestCustomFields = null;
    public ?string $paymentRequestID = null;
    public ?string $payoutID = null;
    public ?array $proprietaryBankTransactionCode = null;
    public ?string $rawReference = null;
    public ?string $reference = null;
    public ?string $ruleID = null;
    public ?array $statementReferences = null;
    public ?string $status = null;
    public mixed $supplementaryData = null;
    public ?array $tags = null;
    public ?string $theirReference = null;
    public ?int $totalPages = null;
    public ?int $totalSize = null;
    public array $transactionAmount;
    public ?string $transactionDate = null;
    public ?array $transactionInformation = null;
    public ?string $transactionMutability = null;
    public ?string $type = null;
    public ?string $valueDateTime = null;
    public ?string $virtualIBAN = null;
    public ?string $yourReference = null;
}

/** Request payload for Transaction#load. */
class TransactionLoadMatch
{
    public ?string $id = null;
    public ?int $sequence_number = null;
    public ?string $transaction_id = null;
    public ?string $account_id = null;
}

/** Request payload for Transaction#list. */
class TransactionListMatch
{
    public ?string $account_id = null;
    public ?string $id = null;
    public ?string $merchant_id = null;
}

/** Request payload for Transaction#create. */
class TransactionCreateData
{
    public string $id;
    public ?string $accountID = null;
    public ?string $accountName = null;
    public ?int $accountSequenceNumber = null;
    public ?array $addressDetails = null;
    public ?float $amount = null;
    public ?int $amountMinorUnits = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $bookingDateTime = null;
    public ?array $chargeDetails = null;
    public ?array $content = null;
    public ?array $counterparty = null;
    public ?string $counterpartySummary = null;
    public ?string $currency = null;
    public ?array $currencyExchange = null;
    public ?string $date = null;
    public ?string $description = null;
    public ?array $enrichment = null;
    public ?float $fxAmount = null;
    public ?string $fxCurrency = null;
    public ?float $fxRate = null;
    public array $grossAmount;
    public ?string $inserted = null;
    public ?array $isoBankTransactionCode = null;
    public ?array $merchant = null;
    public ?string $merchantID = null;
    public ?int $pageNumber = null;
    public ?int $pageSize = null;
    public array $payeeDetails;
    public array $payerDetails;
    public ?array $paymentRequestCustomFields = null;
    public ?string $paymentRequestID = null;
    public ?string $payoutID = null;
    public ?array $proprietaryBankTransactionCode = null;
    public ?string $rawReference = null;
    public ?string $reference = null;
    public ?string $ruleID = null;
    public ?array $statementReferences = null;
    public ?string $status = null;
    public mixed $supplementaryData = null;
    public ?array $tags = null;
    public ?string $theirReference = null;
    public ?int $totalPages = null;
    public ?int $totalSize = null;
    public array $transactionAmount;
    public ?string $transactionDate = null;
    public ?array $transactionInformation = null;
    public ?string $transactionMutability = null;
    public ?string $type = null;
    public ?string $valueDateTime = null;
    public ?string $virtualIBAN = null;
    public ?string $yourReference = null;
}

/** Request payload for Transaction#remove. */
class TransactionRemoveMatch
{
    public string $id;
}

/** User entity data model. */
class User
{
    public ?array $clientSessionTimeouts = null;
    public string $emailAddress;
    public string $firstName;
    public ?string $id = null;
    public string $lastName;
    public ?bool $passkeyAdded = null;
    public ?array $permissions = null;
    public ?string $profile = null;
    public ?array $rolesWithScope = null;
    public ?bool $twoFactorEnabled = null;
    public ?string $userInviteID = null;
}

/** Request payload for User#list. */
class UserListMatch
{
    public ?string $merchant_id = null;
}

/** Request payload for User#update. */
class UserUpdateData
{
    public string $id;
    public ?array $clientSessionTimeouts = null;
    public ?string $emailAddress = null;
    public ?string $firstName = null;
    public ?string $lastName = null;
    public ?bool $passkeyAdded = null;
    public ?array $permissions = null;
    public ?string $profile = null;
    public ?array $rolesWithScope = null;
    public ?bool $twoFactorEnabled = null;
    public ?string $userInviteID = null;
}

/** UserInvite entity data model. */
class UserInvite
{
    public ?array $authorisationStatus = null;
    public ?array $failedUserInvites = null;
    public ?string $id = null;
    public ?string $initialRoleID = null;
    public ?string $inviteeEmailAddress = null;
    public ?string $inviteeFirstName = null;
    public ?string $inviteeLastName = null;
    public ?string $inviterEmailAddress = null;
    public ?string $inviterFirstName = null;
    public ?string $inviterLastName = null;
    public ?bool $isAuthorised = null;
    public ?bool $isInviteeRegistered = null;
    public ?string $lastInvited = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public ?string $message = null;
    public ?string $registrationUrl = null;
    public ?bool $sendInviteEmail = null;
    public ?string $status = null;
    public array $user;
    public ?string $userID = null;
    public ?array $userInvites = null;
}

/** Request payload for UserInvite#load. */
class UserInviteLoadMatch
{
    public ?string $id = null;
    public ?string $userinvite_id = null;
}

/** Request payload for UserInvite#list. */
class UserInviteListMatch
{
    public string $merchant_id;
}

/** Request payload for UserInvite#create. */
class UserInviteCreateData
{
    public ?string $id = null;
    public ?array $authorisationStatus = null;
    public ?array $failedUserInvites = null;
    public ?string $initialRoleID = null;
    public ?string $inviteeEmailAddress = null;
    public ?string $inviteeFirstName = null;
    public ?string $inviteeLastName = null;
    public ?string $inviterEmailAddress = null;
    public ?string $inviterFirstName = null;
    public ?string $inviterLastName = null;
    public ?bool $isAuthorised = null;
    public ?bool $isInviteeRegistered = null;
    public ?string $lastInvited = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public ?string $message = null;
    public ?string $registrationUrl = null;
    public ?bool $sendInviteEmail = null;
    public ?string $status = null;
    public array $user;
    public ?string $userID = null;
    public ?array $userInvites = null;
}

/** Request payload for UserInvite#update. */
class UserInviteUpdateData
{
    public string $id;
    public ?array $authorisationStatus = null;
    public ?array $failedUserInvites = null;
    public ?string $initialRoleID = null;
    public ?string $inviteeEmailAddress = null;
    public ?string $inviteeFirstName = null;
    public ?string $inviteeLastName = null;
    public ?string $inviterEmailAddress = null;
    public ?string $inviterFirstName = null;
    public ?string $inviterLastName = null;
    public ?bool $isAuthorised = null;
    public ?bool $isInviteeRegistered = null;
    public ?string $lastInvited = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public ?string $message = null;
    public ?string $registrationUrl = null;
    public ?bool $sendInviteEmail = null;
    public ?string $status = null;
    public ?array $user = null;
    public ?string $userID = null;
    public ?array $userInvites = null;
}

/** Request payload for UserInvite#remove. */
class UserInviteRemoveMatch
{
    public string $id;
}

/** Virtual entity data model. */
class Virtual
{
    public ?string $accountName = null;
    public ?string $accountSupplierName = null;
    public ?float $availableBalance = null;
    public ?int $availableBalanceMinorUnits = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $bankName = null;
    public ?string $consentID = null;
    public array $createdBy;
    public ?string $createdByDisplayName = null;
    public ?string $currency = null;
    public ?string $defaultPaymentRail = null;
    public ?string $displayName = null;
    public ?string $expiryDate = null;
    public ?string $externalAccountIcon = null;
    public ?string $id = null;
    public array $identifier;
    public ?string $inserted = null;
    public ?bool $isArchived = null;
    public ?bool $isConnectedAccount = null;
    public ?bool $isDefault = null;
    public ?bool $isTrustAccount = null;
    public ?bool $isVirtual = null;
    public ?array $lastTransaction = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public string $name;
    public ?string $physicalAccountID = null;
    public ?array $rules = null;
    public ?float $submittedPayoutsBalance = null;
    public ?int $submittedPayoutsBalanceMinorUnits = null;
    public ?string $summary = null;
    public ?string $supplierSepaInstantStatus = null;
    public ?string $xeroBankFeedConnectionStatus = null;
    public ?string $xeroBankFeedLastSyncedAt = null;
    public ?string $xeroBankFeedSyncLastFailedAt = null;
    public ?string $xeroBankFeedSyncLastFailureReason = null;
    public ?string $xeroBankFeedSyncStatus = null;
    public ?int $xeroUnsynchronisedTransactionsCount = null;
}

/** Request payload for Virtual#create. */
class VirtualCreateData
{
    public string $account_id;
    public ?string $accountName = null;
    public ?string $accountSupplierName = null;
    public ?float $availableBalance = null;
    public ?int $availableBalanceMinorUnits = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $bankName = null;
    public ?string $consentID = null;
    public array $createdBy;
    public ?string $createdByDisplayName = null;
    public ?string $currency = null;
    public ?string $defaultPaymentRail = null;
    public ?string $displayName = null;
    public ?string $expiryDate = null;
    public ?string $externalAccountIcon = null;
    public ?string $id = null;
    public array $identifier;
    public ?string $inserted = null;
    public ?bool $isArchived = null;
    public ?bool $isConnectedAccount = null;
    public ?bool $isDefault = null;
    public ?bool $isTrustAccount = null;
    public ?bool $isVirtual = null;
    public ?array $lastTransaction = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public string $name;
    public ?string $physicalAccountID = null;
    public ?array $rules = null;
    public ?float $submittedPayoutsBalance = null;
    public ?int $submittedPayoutsBalanceMinorUnits = null;
    public ?string $summary = null;
    public ?string $supplierSepaInstantStatus = null;
    public ?string $xeroBankFeedConnectionStatus = null;
    public ?string $xeroBankFeedLastSyncedAt = null;
    public ?string $xeroBankFeedSyncLastFailedAt = null;
    public ?string $xeroBankFeedSyncLastFailureReason = null;
    public ?string $xeroBankFeedSyncStatus = null;
    public ?int $xeroUnsynchronisedTransactionsCount = null;
}

/** Request payload for Virtual#update. */
class VirtualUpdateData
{
    public string $account_id;
    public string $id;
    public ?string $accountName = null;
    public ?string $accountSupplierName = null;
    public ?float $availableBalance = null;
    public ?int $availableBalanceMinorUnits = null;
    public ?float $balance = null;
    public ?int $balanceMinorUnits = null;
    public ?string $bankName = null;
    public ?string $consentID = null;
    public ?array $createdBy = null;
    public ?string $createdByDisplayName = null;
    public ?string $currency = null;
    public ?string $defaultPaymentRail = null;
    public ?string $displayName = null;
    public ?string $expiryDate = null;
    public ?string $externalAccountIcon = null;
    public ?array $identifier = null;
    public ?string $inserted = null;
    public ?bool $isArchived = null;
    public ?bool $isConnectedAccount = null;
    public ?bool $isDefault = null;
    public ?bool $isTrustAccount = null;
    public ?bool $isVirtual = null;
    public ?array $lastTransaction = null;
    public ?string $lastUpdated = null;
    public ?string $merchantID = null;
    public ?string $merchantName = null;
    public ?string $name = null;
    public ?string $physicalAccountID = null;
    public ?array $rules = null;
    public ?float $submittedPayoutsBalance = null;
    public ?int $submittedPayoutsBalanceMinorUnits = null;
    public ?string $summary = null;
    public ?string $supplierSepaInstantStatus = null;
    public ?string $xeroBankFeedConnectionStatus = null;
    public ?string $xeroBankFeedLastSyncedAt = null;
    public ?string $xeroBankFeedSyncLastFailedAt = null;
    public ?string $xeroBankFeedSyncLastFailureReason = null;
    public ?string $xeroBankFeedSyncStatus = null;
    public ?int $xeroUnsynchronisedTransactionsCount = null;
}

/** Webhook entity data model. */
class Webhook
{
    public ?string $destinationUrl = null;
    public ?string $emailAddress = null;
    public ?string $failedNotificationEmailAddress = null;
    public ?string $id = null;
    public ?bool $isActive = null;
    public ?string $merchantID = null;
    public ?string $notificationMethod = null;
    public ?array $resourceTypes = null;
    public ?bool $retry = null;
    public ?string $secret = null;
    public ?int $version = null;
}

/** Request payload for Webhook#load. */
class WebhookLoadMatch
{
    public string $id;
    public ?string $merchant_id = null;
}

/** Request payload for Webhook#list. */
class WebhookListMatch
{
    public string $merchant_id;
}

/** Request payload for Webhook#create. */
class WebhookCreateData
{
    public ?string $destinationUrl = null;
    public ?string $emailAddress = null;
    public ?string $failedNotificationEmailAddress = null;
    public ?string $id = null;
    public ?bool $isActive = null;
    public ?string $merchantID = null;
    public ?string $notificationMethod = null;
    public ?array $resourceTypes = null;
    public ?bool $retry = null;
    public ?string $secret = null;
    public ?int $version = null;
}

/** Request payload for Webhook#update. */
class WebhookUpdateData
{
    public string $id;
    public ?string $destinationUrl = null;
    public ?string $emailAddress = null;
    public ?string $failedNotificationEmailAddress = null;
    public ?bool $isActive = null;
    public ?string $merchantID = null;
    public ?string $notificationMethod = null;
    public ?array $resourceTypes = null;
    public ?bool $retry = null;
    public ?string $secret = null;
    public ?int $version = null;
}

/** Request payload for Webhook#remove. */
class WebhookRemoveMatch
{
    public string $id;
}

