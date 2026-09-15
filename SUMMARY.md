# NoFrixion MoneyMoov API

The NoFrixion MoneyMoov API.

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 49 entities and 187 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### [Account](docs/api/account.html)

Results: OK; The payment account was successfully created and is ready to use.; The payment account was created but the is not yet ready to use.; Returns the list of accounts; Returns the list of merchant&#39;s accounts.; Returns a csv file of the accounts for the merchant.; Accepted. The CSV file is being generated.; Returns a list of the transactions for the payment account in CSV format.; The payment account.; No result is returned.; The updated Account record.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `accountBalances`: The various balances for the account.
- `accountID`: ID of the account.
- `accountIdentifications`: The canoncial identifiers for the account.
- `accountName`: Name for the account
- `accountNames`: Optional account names set by the account holder.

### [Batch](docs/api/batch.html)

Results: OK; Returns a list of payout records.

SDK operations: `create`, `load`.

Key fields to recognise:

- `approveUrl`: This field is used when returning a batch payout record to a client. If set it holds the URL
the user needs to visit in order to complete a strong authentication check in order to approve 
the batch payouts.

### [Beneficiary](docs/api/beneficiary.html)

Results: OK; Accepted; The newly created beneficiary.; A BeneficiariesCreateResponse.; Returns the list of beneficiaries.; A BeneficiaryPageResponse containing the list of beneficiaries.; Returns the list of beneficiaries in CSV format.; Accepted. The CSV file is being generated.; The beneficiary record.; No result is returned.; The updated beneficiary record.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `authenticationMethods`: A list of authentication types allowed to authorise the payout.
- `authorisations`: A list of users who have successfully authorised the latest version of the beneficiary.
- `authorisersCompletedCount`: The number of distinct authorisers that have authorised the beneficiary.
- `authorisersRequiredCount`: The number of authorisers required for this beneficiary. Is determined by business settings
on the source account and/or merchant.
- `canAuthorise`: True if the beneficiary can be authorised by the user who loaded it.

### [BeneficiaryGroup](docs/api/beneficiary_group.html)

Results: A BeneficiaryGroupsPageResponse containing the list of beneficiary groups.

SDK operations: `list`.

Key fields to recognise:

- `groupMembers`: The existing group members.
- `groupName`: The descriptive name for the beneficiary group.
- `inserted`: Timestamp indicating when the group was created.
- `lastUpdated`: Timestamp indicating when the group was last updated.
- `merchantID`: Gets or Sets the merchant id.

### [Card](docs/api/card.html)

Results: Returns a card payment response model that indicates the result of the payment attempt.

SDK operations: `create`.

Key fields to recognise:

- `isPayerAuthenticationRequired`: Gets set to true if 3-D Secure payer authentication is required for a payment.
- `isSoftDecline`: Gets set to true if the card processor flagged the transaction as having failed address or
card security number verification. If the payment was for a sale a soft decline will prevent
the transaction being captured. The merchant can set the ignore address verification flag 
if they want to allow soft declines to be accepted or false if not.
- `payerAuthenticationAccessToken`: If a card payment response indicates a 3-D Secure payer authentication is required
this field holds the access token to POST when performing the redirect.
- `payerAuthenticationMerchantData`: If a card payment response indicates a 3-D Secure payer authentication this field may 
get set in order to transfer information back to the &quot;authenticationcallback&quot; method
that gets called automatically after a successful payer authentication attempt.
- `payerAuthenticationUrl`: If a card payment response indicates a 3-D Secure payer authentication is required
this field holds the URL to redirect the payer to their issuing bank.

### [CardCustomerToken](docs/api/card_customer_token.html)

Results: Returns a list of the tokenised cards that have been stored for the customer.; Returns a success response if the token was deleted.

SDK operations: `list`, `load`, `remove`.

Key fields to recognise:

- `cardType`: The type of the tokenised card, for example Visa, MasterCard etc. It&#39;s possible this could
be empty if the card type could not be identified. There is no hard and fast way to
know for sure to know the type of card.
- `customerEmailAddress`: When creating a tokenised card the payer&#39;s email address must be supplied. This is
used as away to group card tokens for an end user. For customer initiated transactions 
it is important that the email address supplied has been verified to belong to the user
initiating the payment.
- `id`: The unique ID of the card token that has been stored for the customer. This is 
the ID to supply when requesting an authorisation on behalf of the customer.

### [CardPayment](docs/api/card_payment.html)

Results: Returns a card payment response model that indicates the result of the refund payment attempt.; Returns a card payment response model that indicates the result of the capture attempt.; Returns a card payment response model that indicates the result of the payment attempt.; Returns a card payment response model that indicates the result of the void payment attempt.; Returns a collection of card payment response models that indicates the result of the void payment request attempt.

SDK operations: `create`.

Key fields to recognise:

- `isPayerAuthenticationRequired`: Gets set to true if 3-D Secure payer authentication is required for a payment.
- `isSoftDecline`: Gets set to true if the card processor flagged the transaction as having failed address or
card security number verification. If the payment was for a sale a soft decline will prevent
the transaction being captured. The merchant can set the ignore address verification flag 
if they want to allow soft declines to be accepted or false if not.
- `payerAuthenticationAccessToken`: If a card payment response indicates a 3-D Secure payer authentication is required
this field holds the access token to POST when performing the redirect.
- `payerAuthenticationMerchantData`: If a card payment response indicates a 3-D Secure payer authentication this field may 
get set in order to transfer information back to the &quot;authenticationcallback&quot; method
that gets called automatically after a successful payer authentication attempt.
- `payerAuthenticationUrl`: If a card payment response indicates a 3-D Secure payer authentication is required
this field holds the URL to redirect the payer to their issuing bank.

### [CardPublicKey](docs/api/card_public_key.html)

Results: Returns a public key that can be used to encrypt sensitive card details.

SDK operations: `load`.

### [Consent](docs/api/consent.html)

Results: OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `authorisationUrl`: The URL the authorising user needs to be redirected to in order to get the open banking
consent token.
- `callbackUrl`: Optional callback URL that the end user performing the open banking authorisation will be redirected to on completion.
- `consentID`: The ID of the open banking consent. Once the consent has been authorised this
is the ID that allows it to be utilised via the open banking APIs to list accounts,
transactions etc.
- `emailAddress`: The email address that identifies the end user that will be authorising the open banking consent request.
- `failureCallbackUrl`: Optional callback URL for open banking consent authorisation failure. 
If the URL is set, the calling application will be redirected here with the error message.
If it’s not set it will fall back on default page that will show the error message.

### [Currency](docs/api/currency.html)

Results: Returns the list of currencies and their static metadata.

SDK operations: `list`.

### [DirectDebitBatchSubmit](docs/api/direct_debit_batch_submit.html)

Results: A DirectDebitBatchSubmitResponse.

SDK operations: `create`.

Key fields to recognise:

- `failedSubmissions`: Dictionary of failed submissions, keyed by the index (1-based) in the original request.
- `successfulSubmissions`: List of successfully submitted direct debit payments.

### [FxRate](docs/api/fx_rate.html)

Results: OK.

SDK operations: `list`, `load`.

Key fields to recognise:

- `exchangeRate`: The price at which the transaction will buy the source currency 
using the destination currency.

### [IPayment](docs/api/i_payment.html)

Results: OK; Returns the payment response.

SDK operations: `create`.

### [Mandate](docs/api/mandate.html)

Results: OK; Mandate information.

SDK operations: `create`, `load`.

Key fields to recognise:

- `accountNumber`: Account number of the customer&#39;s bank account in case of GBP account.
- `addressLine1`: First line of the customer&#39;s address.
- `addressLine2`: Second line of the customer&#39;s address.
- `approvedAt`: Date at which the supplier approved this mandate.
- `city`: Customer&#39;s city.

### [Merchant](docs/api/merchant.html)

Results: Returns the paged list of child merchants successfully.; Returns a list of the merchants the caller has access to.; Returns a list of merchants.; Returns a merchant model.; Returns a list of payout records as a CSV file.; Accepted: The request to export payouts as csv has been accepted for processing but is not yet completed.; Returns a CSV file containing the list of beneficiaries.; Accepted. The CSV file is being generated.; Returns the requested merchant.; Ok; Returns the updated merchant.

SDK operations: `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `accountCurrencies`: The list of currencies that the merchant has accounts for.
- `canHaveTrustAccounts`: Trust accounts are a special type of account that allow the account name
to be trusted for use in statements and verification of payee checks.
- `cardPaymentProcessor`: Name of the card payment processor. Can be &quot;Checkout&quot;, &quot;Stripe&quot; or none.
- `companyID`: The Company ID recorded in the Compliance system.
- `displayQrOnHostedPay`: Indicates if a QR Code containing the payment link should be displayed
on the hosted payment page.

### [MerchantAuthorisationSetting](docs/api/merchant_authorisation_setting.html)

Results: A list of merchant authorisation settings.

SDK operations: `list`.

### [MerchantDirectDebitMandatePage](docs/api/merchant_direct_debit_mandate_page.html)

Results: Mandate collection paged result.

SDK operations: `list`.

Key fields to recognise:

- `approvedAt`: Date at which the supplier approved this mandate.
- `currency`: Currency of this mandate.
- `customerAccountNumber`: Customer&#39;s account number in case of GBP account.
- `customerCity`: Customer&#39;s city of residence.
- `customerCountryCode`: Customer&#39;s country of residence code.

### [MerchantPayByBankSetting](docs/api/merchant_pay_by_bank_setting.html)

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `bankCountryCodes`: The list of country codes representing the banks the country supports.
- `bankID`: ID of the bank to be configured for the merchant.
- `bankName`: Name of the Bank/Institution.
- `businessInstitutionID`: ID that the processor uses to identify the bank (business accounts).
- `currency`: Currency supported by the bank.

### [MerchantPaymentRequestTemplate](docs/api/merchant_payment_request_template.html)

Results: Returns a list of payment request templates for a merchant.; Returns a payment request template for a merchant.; Returns a success response if the template was deleted.; Returns an updated payment request template.

SDK operations: `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `customFields`: A list of custom fields that can be included in the payment request template.
- `defaultFields`: A list of default fields that are included in the payment request template.
These fields are predefined and map to concrete fields in the payment request.

### [MerchantToken](docs/api/merchant_token.html)

Results: OK; Returns the newly created merchant API token.; Returns a list of the API tokens issued for the merchant.; Returns the merchant API token details.; The updated merchant token.

SDK operations: `create`, `list`, `load`, `update`.

Key fields to recognise:

- `authenticationMethods`: A list of authentication types allowed to authorise the merchant token.
- `authorisations`: A list of users who have successfully authorised the latest version of the beneficiary.
- `authorisersCompletedCount`: The number of distinct authorisers that have authorised the merchant token.
- `authorisersRequiredCount`: The number of authorisers required for this merchant token. Is determined by business settings
on the source account and/or merchant.
- `canAuthorise`: True if the merchant token can be authorised by the user who loaded it.

### [Metadata](docs/api/metadata.html)

Results: Returns OK if the problem notification request is accepted.; OK.

SDK operations: `load`.

### [NoFrixionVersion](docs/api/no_frixion_version.html)

Results: Returns a model with fields describing the current API version.

SDK operations: `load`.

### [OpenBanking](docs/api/open_banking.html)

Results: OK.

SDK operations: `create`, `remove`.

### [Payeeverification](docs/api/payeeverification.html)

Results: OK.

SDK operations: `create`.

Key fields to recognise:

- `accountName`: The name of the account to verify
- `accountNumber`: The account number of the account to verify (for CoP checks)
- `iban`: The IBAN of the account to verify (for VoP checks)
- `payeeVerifiedAccountName`: The verified account name of the payee, if available (in case of a close match)
- `result`: The result of the payee verification

### [Payment](docs/api/payment.html)

Results: No result is returned.; OK; Returns the newly created payment request record.; Returns a payment request record.; Returns the updated payment request record.

SDK operations: `create`, `load`, `update`.

Key fields to recognise:

- `amount`: The amount of money to request.
- `amountPending`: Total amount that has been authorised but not settled for this payment request.
- `amountReceived`: Total amount received for this payment request.
- `amountRefunded`: Total amount refunded for this payment request.
- `autoSendReceipt`: If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received.

### [PaymentAccount](docs/api/payment_account.html)

Results: Returns the list of accounts for the merchant.; OK.

SDK operations: `list`.

Key fields to recognise:

- `accountName`: Name for the account
- `accountSupplierName`: The payment account supplier name. A payment account can be supplied by multiple payment processors.
- `availableBalance`: The current available balance of the account. Calculated by subtracting any submitted payments from the current balance.
- `availableBalanceMinorUnits`: The available balance expressed in the currency’s minor units (for example cents, pence).
- `balance`: Balance of the account.

### [PaymentAccountMinimal](docs/api/payment_account_minimal.html)

Results: Returns the list of accounts for the merchant.

SDK operations: `list`.

Key fields to recognise:

- `accountName`: Name for the account
- `availableBalance`: The current available balance of the account. Calculated by subtracting any submitted payments from the current balance.
- `balance`: Balance of the account.
- `balanceMinorUnits`: Balance of the account expressed in the currency’s minor units (for example cents, pence).
- `currency`: Currency of the account in ISO 4217 format

### [PaymentInitiation](docs/api/payment_initiation.html)

Results: Returns a payment initiation response that contains the payment ID and the payment link.

SDK operations: `create`.

Key fields to recognise:

- `paymentInitiationID`: The unique identifier of the payment initiation request.
- `paymentRequestCallbackUrl`: The callback URL that was set when the payment request was created. Payers will be 
redirected to this URL after a successful payment initiation.
- `redirectUrl`: A redirect URL for the user to authorise the payment initiation request at the ASPSP

### [PaymentRequest](docs/api/payment_request.html)

Results: A PaymentRequestsCreateResponse.; Returns the list of payment requests.; Returns the list of payment requests in a CSV file.; Accepted. The CSV file is being generated.; Returns the PDF file of the payment request receipt.; No result is returned.; OK.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `amount`: The amount of money to request.
- `amountPending`: Total amount that has been authorised but not settled for this payment request.
- `amountReceived`: Total amount received for this payment request.
- `amountRefunded`: Total amount refunded for this payment request.
- `autoSendReceipt`: If set to true, a receipt will be automatically sent to the CustomerEmailAddress when payments are received.

### [PaymentRequestEvent](docs/api/payment_request_event.html)

Results: Returns a list of payment request event records.

SDK operations: `list`.

Key fields to recognise:

- `applePayTransactionID`: Transaction ID received in Apple pay token.
- `cardAuthorizationResponseID`: For a successful card authorization this field will hold the response ID. If a capture
needs to be performed this is the ID that must be used.
- `cardExpiryMonth`: For card payment events this field holds the payer&#39;s card expiry month.
- `cardExpiryYear`: For card payment events this field holds the payer&#39;s card expiry year.
- `cardIssuer`: For card payment events this field holds the payer&#39;s card issuer.

### [PaymentRequestMetric](docs/api/payment_request_metric.html)

Results: OK.

SDK operations: `load`.

### [PaymentRequestMinimal](docs/api/payment_request_minimal.html)

Results: Returns a payment request minimal record.

SDK operations: `list`.

Key fields to recognise:

- `amount`: The amount of money to request.
- `amountPending`: The amount of money that was authorised but has not arrived in the account yet.
- `amountReceived`: The amount of money that has been received for this payment request.
- `amountRefunded`: The amount of money that has been refunded for this payment request.
- `countryCode`: The country code associated with the payment.

### [PaymentRequestResult](docs/api/payment_request_result.html)

Results: Returns a payment request result record.

SDK operations: `list`.

Key fields to recognise:

- `amount`: The authorised payment amount.
- `currency`: The authorised payment currency.
- `customerID`: The customer id
- `paymentRequestID`: The ID of the payment request the result is for.
- `payments`: The list of payment attempts that have been received for the payment request.

### [Payout](docs/api/payout.html)

Results: OK; Accepted. Payouts have been submitted for processing.; Accepted. Payout has been submitted for processing.; The newly created payout record.; A PayoutsCreateResponse.; The newly created and submitted payout record.; Returns a list of payout records.; Returns a CSV file of payout records.; Accepted. The CSV file is being generated.; Returns a payout record.; No result is returned.; The result of each payout deletion.; The updated Payout record.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `accountID`: Gets or Sets Account Id of sending account
- `allowIncomplete`: If set to true the payout will get created even if the business validation rules fail.
- `amount`: Gets or Sets payout amount
- `amountMinorUnits`: The payout amount expressed in the currency’s minor units (for example cents, pence).
- `approvePayoutUrl`: This field is used when returning an payout record to a client. If set it holds the URL
the user needs to visit in order to complete a strong authentication check in order to approve 
the payout.

### [PayoutKeysetPage](docs/api/payout_keyset_page.html)

Results: OK.

SDK operations: `list`.

Key fields to recognise:

- `accountID`: Gets or Sets Account Id of sending account
- `amount`: Gets or Sets payout amount
- `amountMinorUnits`: The payout amount expressed in the currency’s minor units (for example cents, pence).
- `approvePayoutUrl`: This field is used when returning an payout record to a client. If set it holds the URL
the user needs to visit in order to complete a strong authentication check in order to approve 
the payout.
- `approverID`: Gets the User ID of person that approved the payout.

### [PayoutMetric](docs/api/payout_metric.html)

Results: OK.

SDK operations: `load`.

### [Payrun](docs/api/payrun.html)

Results: OK.; OK; Accepted.; Returns a payrun record.; No result is returned.; The updated Payrun record.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `authorisations`: A list of the users who have successfully authorised the latest version of the payrun and when.
- `authorisersCompletedCount`: The number of distinct authorisers that have authorised the payrun.
- `authorisersRequiredCount`: The number of authorisers required for this payrun. Is determined by business settings
on the source account and/or merchant.
- `canAuthorise`: True if the payrun can be authorised by the user who loaded it.
- `hasCurrentUserAuthorised`: True if the payrun was loaded for a user and that user has already authorised the latest version of the payrun.

### [Report](docs/api/report.html)

Results: OK.

SDK operations: `update`.

### [ReportResult](docs/api/report_result.html)

Results: OK.

SDK operations: `load`.

### [Role](docs/api/role.html)

Results: OK; A RolesCreateResponse.

SDK operations: `create`.

### [Rule](docs/api/rule.html)

Results: OK; The newly created rule.; Returns a list of the rules for the merchant.; No result is returned.; The updated rule.; The updated rule.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `accountID`: An optional ID of an internal account the counterparty is associated with. If set
it will take precedence over any other destination details set for the counterparty.
- `approveUrl`: If set this property holds the URL an approver needs to visit in order to
complete a strong authentication check in order to approve the rule.
- `authenticationMethods`: A list of authentication types allowed to authorise the payout.
- `authorisations`: A list of the users who have successfully authorised the latest version of the rule and when.
- `authorisersCompletedCount`: The number of distinct authorisers that have authorised the rule.

### [RuleEvent](docs/api/rule_event.html)

Results: Returns a list of the events for the rule.

SDK operations: `list`.

### [Tag](docs/api/tag.html)

Results: OK; The newly created tag.; Returns a list of merchant tags.

SDK operations: `create`, `list`.

### [Token](docs/api/token.html)

Results: OK; Accepted; No result is returned.

SDK operations: `create`, `remove`.

### [Transaction](docs/api/transaction.html)

Results: Indicates the tags have been added to the transaction.; Returns a list of the transactions for the payment account.; OK; Returns the list of the transactions in a CSV file.; Accepted. The CSV file is being generated.; Returns the list of the transactions.; The transaction.; Indicates the tag has been removed from the transaction.

SDK operations: `create`, `list`, `load`, `remove`.

Key fields to recognise:

- `accountID`: The ID of the account the transaction belongs to.
- `accountName`: The name of the account the transaction belongs to.
- `accountSequenceNumber`: The sequence number of transaction on a per account basis. This sequence number is guaranteed to be an arithemtic sequence 
number for all transactions belonging to the same account.
- `amount`: Amount of the transaction. Negative values indicate a pay out debit), positive
values a pay in (credit).
- `amountMinorUnits`: Amount of the transaction expressed in the currency’s minor units (for example cents, pence).

### [User](docs/api/user.html)

Results: OK; Returns a paged list of the users associated with the merchant.; Returns the users associated with the merchant.; Returns a user profile model.; Returns a user model.; A user profile model is returned.; Returns the updated user profile.

SDK operations: `list`, `update`.

Key fields to recognise:

- `clientSessionTimeouts`: The number of seconds a session for this user should last before expiring.
This is based on the user&#39;s role on the merchant.
This is used to set the session timeout in the client. If not set the client&#39;s default
session timeout will be used.
- `userInviteID`: Optional ID of the invite that is being accepted so the user can be assigned a role on a new merchant.

### [UserInvite](docs/api/user_invite.html)

Results: OK; Accepted; Returns an OK response.; A UserInvitesCreateResponse.; Returns a paged list of the user invites associated with merchant.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `initialRoleID`: The role ID to automatically assign to the merchant’s very first user.
Typically set by the compliance team when the first user is invited to a new merchant.
- `inviteeEmailAddress`: Email address of the user being invited.
- `inviteeFirstName`: First Name of the user being invited.
- `inviteeLastName`: Last Name of the user being invited.
- `isAuthorised`: Will be set to true once the invite has met the authorisation requirements.

### [Virtual](docs/api/virtual.html)

Results: OK; The virtual account was successfully created and is ready to use.; The updated Virtual Account record.

SDK operations: `create`, `update`.

Key fields to recognise:

- `accountName`: Name for the account
- `accountSupplierName`: The payment account supplier name. A payment account can be supplied by multiple payment processors.
- `availableBalance`: The current available balance of the account. Calculated by subtracting any submitted payments from the current balance.
- `availableBalanceMinorUnits`: The available balance expressed in the currency’s minor units (for example cents, pence).
- `balance`: Balance of the account.

### [Webhook](docs/api/webhook.html)

Results: OK; The created webhook is returned.; A list of configured webhooks.; The requested webhook.; No result is returned.; The updated Webhook is returned.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `destinationUrl`: The destination URL for the webhook.
- `emailAddress`: The recipient email address(es) for notifications.
- `failedNotificationEmailAddress`: The email address to which notifications about failed webhook deliveries will be sent.
- `merchantID`: The ID of the merchant that the webhook is for.
- `notificationMethod`: The type of notification that will be sent.

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| [Account](docs/api/account.html) | `create` | `POST /api/v1/accounts/{accountID}/{currency}` | Required |
| [Account](docs/api/account.html) | `create` | `POST /api/v1/accounts/{accountID}/statements` | Required |
| [Account](docs/api/account.html) | `create` | `POST /api/v1/accounts` | Required |
| [Account](docs/api/account.html) | `list` | `GET /api/v1/accounts` | Required |
| [Account](docs/api/account.html) | `list` | `GET /api/v1/merchants/{merchantID}/accounts` | Required |
| [Account](docs/api/account.html) | `load` | `GET /api/v1/accounts/export` | Required |
| [Account](docs/api/account.html) | `load` | `GET /api/v1/accounts/{accountID}/transactions/export` | Required |
| [Account](docs/api/account.html) | `load` | `GET /api/v1/accounts/{accountID}/statements/{id}` | Required |
| [Account](docs/api/account.html) | `load` | `GET /api/v1/merchants/{merchantID}/accounts/{accountID}` | Required |
| [Account](docs/api/account.html) | `load` | `GET /api/v1/accounts/{accountID}` | Required |
| [Account](docs/api/account.html) | `load` | `GET /api/v1/openbanking/accounts/{id}` | Required |
| [Account](docs/api/account.html) | `load` | `GET /api/v1/accounts/statements` | Required |
| [Account](docs/api/account.html) | `remove` | `DELETE /api/v1/accounts/archive/{id}` | Required |
| [Account](docs/api/account.html) | `remove` | `DELETE /api/v1/accounts/statements` | Required |
| [Account](docs/api/account.html) | `update` | `PUT /api/v1/accounts/{accountID}/topup/{amount}` | Required |
| [Account](docs/api/account.html) | `update` | `PUT /api/v1/accounts/unarchive/{id}` | Required |
| [Account](docs/api/account.html) | `update` | `PUT /api/v1/accounts/{id}` | Required |
| [Batch](docs/api/batch.html) | `create` | `POST /api/v1/payouts/batch` | Required |
| [Batch](docs/api/batch.html) | `load` | `GET /api/v1/payouts/batch/{id}` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `create` | `POST /api/v1/beneficiaries/authorise/{id}` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `create` | `POST /api/v1/beneficiaries` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `create` | `POST /api/v1/beneficiaries/batchcreate` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `list` | `GET /api/v1/beneficiaries` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `list` | `GET /api/v1/merchants/{merchantID}/beneficiaries` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `load` | `GET /api/v1/beneficiaries/export` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `load` | `GET /api/v1/merchants/{merchantID}/beneficiaries/{id}` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `load` | `GET /api/v1/beneficiaries/{id}` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `remove` | `DELETE /api/v1/beneficiaries/{id}` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `update` | `PUT /api/v1/beneficiaries/disable/{id}` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `update` | `PUT /api/v1/beneficiaries/enable/{id}` | Required |
| [Beneficiary](docs/api/beneficiary.html) | `update` | `PUT /api/v1/beneficiaries/{id}` | Required |
| [BeneficiaryGroup](docs/api/beneficiary_group.html) | `list` | `GET /api/v1/merchants/{merchantID}/beneficiarygroups` | Required |
| [Card](docs/api/card.html) | `create` | `POST /api/v1/paymentrequests/{id}/card` | Required |
| [CardCustomerToken](docs/api/card_customer_token.html) | `list` | `GET /api/v1/paymentrequests/card/customertokens/{merchantID}/{customerEmailAddress}` | Required |
| [CardCustomerToken](docs/api/card_customer_token.html) | `load` | `GET /api/v1/paymentrequests/card/customertokens/{customerEmailAddress}` | Required |
| [CardCustomerToken](docs/api/card_customer_token.html) | `remove` | `DELETE /api/v1/paymentrequests/card/customertokens/removeall/{merchantID}/{customerEmailAddress}` | Required |
| [CardCustomerToken](docs/api/card_customer_token.html) | `remove` | `DELETE /api/v1/paymentrequests/card/customertokens/removeall/{customerEmailAddress}` | Required |
| [CardCustomerToken](docs/api/card_customer_token.html) | `remove` | `DELETE /api/v1/paymentrequests/card/customertokens/{id}` | Required |
| [CardPayment](docs/api/card_payment.html) | `create` | `POST /api/v1/paymentrequests/{id}/card/refund/{partialRefundAmount}` | Required |
| [CardPayment](docs/api/card_payment.html) | `create` | `POST /api/v1/paymentrequests/{id}/card/capture` | Required |
| [CardPayment](docs/api/card_payment.html) | `create` | `POST /api/v1/paymentrequests/{id}/card/paywithtoken` | Required |
| [CardPayment](docs/api/card_payment.html) | `create` | `POST /api/v1/paymentrequests/{id}/card/void` | Required |
| [CardPayment](docs/api/card_payment.html) | `create` | `POST /api/v1/paymentrequests/{id}/card/voidpaymentrequest` | Required |
| [CardPublicKey](docs/api/card_public_key.html) | `load` | `GET /api/v1/paymentrequests/{id}/card/publickey` | Required |
| [Consent](docs/api/consent.html) | `create` | `POST /api/v1/openbanking/consents` | Required |
| [Consent](docs/api/consent.html) | `list` | `GET /api/v1/openbanking/consents/{merchantID}/{email}` | Required |
| [Consent](docs/api/consent.html) | `load` | `GET /api/v1/openbanking/consents/{id}` | Required |
| [Consent](docs/api/consent.html) | `remove` | `DELETE /api/v1/openbanking/consents/{id}` | Required |
| [Consent](docs/api/consent.html) | `update` | `PATCH /api/v1/openbanking/consents/{id}` | Required |
| [Currency](docs/api/currency.html) | `list` | `GET /api/v1/currencies` | Required |
| [DirectDebitBatchSubmit](docs/api/direct_debit_batch_submit.html) | `create` | `POST /api/v1/paymentrequests/directdebit/batchsubmit` | Required |
| [FxRate](docs/api/fx_rate.html) | `list` | `GET /api/v1/payouts/fxallheldrates/{source}/{destination}` | Required |
| [FxRate](docs/api/fx_rate.html) | `load` | `GET /api/v1/payouts/fxheldrate/{source}/{destination}/{validForMinutes}` | Required |
| [IPayment](docs/api/i_payment.html) | `create` | `POST /api/v1/paymentrequests/payondemand` | Required |
| [Mandate](docs/api/mandate.html) | `create` | `POST /api/v1/mandates` | Required |
| [Mandate](docs/api/mandate.html) | `load` | `GET /api/v1/mandates/{id}` | Required |
| [Merchant](docs/api/merchant.html) | `list` | `GET /api/v1/merchants/{merchantID}/childmerchants` | Required |
| [Merchant](docs/api/merchant.html) | `list` | `GET /api/v1/merchants/paged` | Required |
| [Merchant](docs/api/merchant.html) | `list` | `GET /api/v1/merchants` | Required |
| [Merchant](docs/api/merchant.html) | `list` | `GET /api/v1/metadata/whoamimerchant` | Required |
| [Merchant](docs/api/merchant.html) | `list` | `GET /api/v1/metadata/whoamimerchantsigned` | Required |
| [Merchant](docs/api/merchant.html) | `list` | `GET /api/v1/metadata/whoamimerchantwhitelist` | Required |
| [Merchant](docs/api/merchant.html) | `load` | `GET /api/v1/merchants/{merchantID}/payouts/export` | Required |
| [Merchant](docs/api/merchant.html) | `load` | `GET /api/v1/merchants/{merchantID}/beneficiaries/export` | Required |
| [Merchant](docs/api/merchant.html) | `load` | `GET /api/v1/merchants/{merchantID}` | Required |
| [Merchant](docs/api/merchant.html) | `remove` | `DELETE /api/v1/merchants/{merchantId}/users/{userId}` | Required |
| [Merchant](docs/api/merchant.html) | `remove` | `DELETE /api/v1/merchants/{merchantID}/tags/{tagID}` | Required |
| [Merchant](docs/api/merchant.html) | `update` | `PUT /api/v1/merchants/{merchantID}` | Required |
| [Merchant](docs/api/merchant.html) | `update` | `PUT /api/v1/merchants/{merchantId}/suspend` | Required |
| [MerchantAuthorisationSetting](docs/api/merchant_authorisation_setting.html) | `list` | `GET /api/v1/merchants/{merchantID}/authorisationsettings` | Required |
| [MerchantDirectDebitMandatePage](docs/api/merchant_direct_debit_mandate_page.html) | `list` | `GET /api/v1/mandates` | Required |
| [MerchantPayByBankSetting](docs/api/merchant_pay_by_bank_setting.html) | `list` | `GET /api/v1/merchants/{merchantID}/banksettings` | Required |
| [MerchantPaymentRequestTemplate](docs/api/merchant_payment_request_template.html) | `list` | `GET /api/v1/paymentrequests/{merchantID}/templates` | Required |
| [MerchantPaymentRequestTemplate](docs/api/merchant_payment_request_template.html) | `load` | `GET /api/v1/paymentrequests/{merchantID}/templates/{templateID}` | Required |
| [MerchantPaymentRequestTemplate](docs/api/merchant_payment_request_template.html) | `remove` | `DELETE /api/v1/paymentrequests/{merchantID}/templates/{templateID}` | Required |
| [MerchantPaymentRequestTemplate](docs/api/merchant_payment_request_template.html) | `update` | `PUT /api/v1/paymentrequests/{merchantID}/templates/{templateID}` | Required |
| [MerchantToken](docs/api/merchant_token.html) | `create` | `POST /api/v1/tokens` | Required |
| [MerchantToken](docs/api/merchant_token.html) | `list` | `GET /api/v1/merchants/{merchantID}/tokens` | Required |
| [MerchantToken](docs/api/merchant_token.html) | `load` | `GET /api/v1/tokens/{id}` | Required |
| [MerchantToken](docs/api/merchant_token.html) | `update` | `PUT /api/v1/tokens/{id}` | Required |
| [Metadata](docs/api/metadata.html) | `load` | `GET /api/v1/metadata/problemnotification` | Required |
| [Metadata](docs/api/metadata.html) | `load` | `GET /api/v1/metadata/problem` | Required |
| [NoFrixionVersion](docs/api/no_frixion_version.html) | `load` | `GET /api/v1/metadata/version` | Required |
| [OpenBanking](docs/api/open_banking.html) | `create` | `POST /api/v1/openbanking/account/{accountID}/synchronise` | Required |
| [OpenBanking](docs/api/open_banking.html) | `remove` | `DELETE /api/v1/openbanking/consents/{merchantID}/{email}` | Required |
| [OpenBanking](docs/api/open_banking.html) | `remove` | `DELETE /api/v1/openbanking/account/{accountID}` | Required |
| [Payeeverification](docs/api/payeeverification.html) | `create` | `POST /api/v1/openbanking/payeeverification` | Required |
| [Payment](docs/api/payment.html) | `create` | `POST /api/v1/paymentrequests/{id}/directdebit` | Required |
| [Payment](docs/api/payment.html) | `create` | `POST /api/v1/paymentrequests` | Required |
| [Payment](docs/api/payment.html) | `load` | `GET /api/v1/paymentrequests/{id}` | Required |
| [Payment](docs/api/payment.html) | `load` | `GET /api/v1/paymentrequests/getbyorderid/{orderID}` | Required |
| [Payment](docs/api/payment.html) | `update` | `PUT /api/v1/paymentrequests/{id}` | Required |
| [PaymentAccount](docs/api/payment_account.html) | `list` | `GET /api/v1/accounts/paged` | Required |
| [PaymentAccount](docs/api/payment_account.html) | `list` | `GET /api/v1/accounts/{accountID}/virtual` | Required |
| [PaymentAccountMinimal](docs/api/payment_account_minimal.html) | `list` | `GET /api/v1/accounts/minimal` | Required |
| [PaymentInitiation](docs/api/payment_initiation.html) | `create` | `POST /api/v1/paymentrequests/{id}/pisp` | Required |
| [PaymentRequest](docs/api/payment_request.html) | `create` | `POST /api/v1/paymentrequests/batchcreate` | Required |
| [PaymentRequest](docs/api/payment_request.html) | `list` | `GET /api/v1/paymentrequests` | Required |
| [PaymentRequest](docs/api/payment_request.html) | `load` | `GET /api/v1/paymentrequests/export` | Required |
| [PaymentRequest](docs/api/payment_request.html) | `load` | `GET /api/v1/paymentrequests/{id}/receipt` | Required |
| [PaymentRequest](docs/api/payment_request.html) | `remove` | `DELETE /api/v1/paymentrequests/{id}` | Required |
| [PaymentRequest](docs/api/payment_request.html) | `update` | `PUT /api/v1/paymentrequests/{id}/pisp/sandboxcallback` | Required |
| [PaymentRequestEvent](docs/api/payment_request_event.html) | `list` | `GET /api/v1/paymentrequests/{id}/events` | Required |
| [PaymentRequestMetric](docs/api/payment_request_metric.html) | `load` | `GET /api/v1/paymentrequests/metrics` | Required |
| [PaymentRequestMinimal](docs/api/payment_request_minimal.html) | `list` | `GET /api/v1/paymentrequests/{id}/minimal` | Required |
| [PaymentRequestResult](docs/api/payment_request_result.html) | `list` | `GET /api/v1/paymentrequests/{id}/result` | Required |
| [Payout](docs/api/payout.html) | `create` | `POST /api/v1/payouts/batch/submit/{id}` | Required |
| [Payout](docs/api/payout.html) | `create` | `POST /api/v1/payouts/submit/{id}` | Required |
| [Payout](docs/api/payout.html) | `create` | `POST /api/v1/payouts` | Required |
| [Payout](docs/api/payout.html) | `create` | `POST /api/v1/payouts/batchcreate` | Required |
| [Payout](docs/api/payout.html) | `create` | `POST /api/v1/payouts/send` | Required |
| [Payout](docs/api/payout.html) | `create` | `POST /api/v1/payouts/sendbeneficiary` | Required |
| [Payout](docs/api/payout.html) | `list` | `GET /api/v1/payouts` | Required |
| [Payout](docs/api/payout.html) | `list` | `GET /api/v1/accounts/{accountID}/payouts` | Required |
| [Payout](docs/api/payout.html) | `list` | `GET /api/v1/merchants/{merchantID}/payouts` | Required |
| [Payout](docs/api/payout.html) | `load` | `GET /api/v1/payouts/export` | Required |
| [Payout](docs/api/payout.html) | `load` | `GET /api/v1/payouts/fxquote/{source}/{destination}/{amount}` | Required |
| [Payout](docs/api/payout.html) | `load` | `GET /api/v1/payouts/{id}` | Required |
| [Payout](docs/api/payout.html) | `load` | `GET /api/v1/payouts/{id}/proof` | Required |
| [Payout](docs/api/payout.html) | `remove` | `DELETE /api/v1/payouts/{id}` | Required |
| [Payout](docs/api/payout.html) | `remove` | `DELETE /api/v1/payouts/batchdelete` | Required |
| [Payout](docs/api/payout.html) | `update` | `PUT /api/v1/payouts/cancel/{id}` | Required |
| [Payout](docs/api/payout.html) | `update` | `PUT /api/v1/payouts/reject/{id}` | Required |
| [Payout](docs/api/payout.html) | `update` | `PUT /api/v1/payouts/{id}` | Required |
| [PayoutKeysetPage](docs/api/payout_keyset_page.html) | `list` | `GET /api/v1/accounts/{accountID}/payouts/failed` | Required |
| [PayoutKeysetPage](docs/api/payout_keyset_page.html) | `list` | `GET /api/v1/merchants/{merchantID}/payouts/failed` | Required |
| [PayoutKeysetPage](docs/api/payout_keyset_page.html) | `list` | `GET /api/v1/payouts/{merchantID}/failed` | Required |
| [PayoutMetric](docs/api/payout_metric.html) | `load` | `GET /api/v1/payouts/metrics` | Required |
| [Payrun](docs/api/payrun.html) | `create` | `POST /api/v1/payruns/{id}/request-authorisation` | Required |
| [Payrun](docs/api/payrun.html) | `create` | `POST /api/v1/payruns/{id}/submit` | Required |
| [Payrun](docs/api/payrun.html) | `create` | `POST /api/v1/payruns/{merchantID}` | Required |
| [Payrun](docs/api/payrun.html) | `list` | `GET /api/v1/payruns` | Required |
| [Payrun](docs/api/payrun.html) | `load` | `GET /api/v1/payruns/{id}` | Required |
| [Payrun](docs/api/payrun.html) | `remove` | `DELETE /api/v1/payruns/{id}` | Required |
| [Payrun](docs/api/payrun.html) | `remove` | `DELETE /api/v1/payruns/{id}/archive` | Required |
| [Payrun](docs/api/payrun.html) | `update` | `PUT /api/v1/payruns/{id}` | Required |
| [Payrun](docs/api/payrun.html) | `update` | `PUT /api/v1/payruns/{id}/cancel` | Required |
| [Payrun](docs/api/payrun.html) | `update` | `PUT /api/v1/payruns/{id}/reject` | Required |
| [Payrun](docs/api/payrun.html) | `update` | `PUT /api/v1/payruns/{id}/unarchive` | Required |
| [Report](docs/api/report.html) | `update` | `PUT /api/v1/reports/{id}/initiate` | Required |
| [ReportResult](docs/api/report_result.html) | `load` | `GET /api/v1/reports/{id}/result/{statementNumber}` | Required |
| [Role](docs/api/role.html) | `create` | `POST /api/v1/merchants/{merchantID}/roles/batchcreate` | Required |
| [Rule](docs/api/rule.html) | `create` | `POST /api/v1/rules` | Required |
| [Rule](docs/api/rule.html) | `list` | `GET /api/v1/rules` | Required |
| [Rule](docs/api/rule.html) | `load` | `GET /api/v1/rules/{id}` | Required |
| [Rule](docs/api/rule.html) | `remove` | `DELETE /api/v1/rules/{id}` | Required |
| [Rule](docs/api/rule.html) | `update` | `PUT /api/v1/rules/{id}` | Required |
| [Rule](docs/api/rule.html) | `update` | `PUT /api/v1/rules/{id}/disable` | Required |
| [RuleEvent](docs/api/rule_event.html) | `list` | `GET /api/v1/rules/{id}/events` | Required |
| [Tag](docs/api/tag.html) | `create` | `POST /api/v1/merchants/{merchantID}/tags` | Required |
| [Tag](docs/api/tag.html) | `list` | `GET /api/v1/merchants/{merchantID}/tags` | Required |
| [Token](docs/api/token.html) | `create` | `POST /api/v1/tokens/authorise/{id}` | Required |
| [Token](docs/api/token.html) | `remove` | `DELETE /api/v1/tokens/{id}` | Required |
| [Transaction](docs/api/transaction.html) | `create` | `POST /api/v1/transactions/{id}/tags` | Required |
| [Transaction](docs/api/transaction.html) | `list` | `GET /api/v1/accounts/{accountID}/transactions` | Required |
| [Transaction](docs/api/transaction.html) | `list` | `GET /api/v1/openbanking/transactions/{id}/{accountID}` | Required |
| [Transaction](docs/api/transaction.html) | `list` | `GET /api/v1/merchants/{merchantID}/transactions` | Required |
| [Transaction](docs/api/transaction.html) | `list` | `GET /api/v1/transactions` | Required |
| [Transaction](docs/api/transaction.html) | `load` | `GET /api/v1/transactions/{accountID}/export` | Required |
| [Transaction](docs/api/transaction.html) | `load` | `GET /api/v1/transactions/{accountID}` | Required |
| [Transaction](docs/api/transaction.html) | `load` | `GET /api/v1/transactions/{accountID}/from/{sequenceNumber}` | Required |
| [Transaction](docs/api/transaction.html) | `load` | `GET /api/v1/accounts/{accountID}/transactions/{id}` | Required |
| [Transaction](docs/api/transaction.html) | `load` | `GET /api/v1/transactions/detail/{id}` | Required |
| [Transaction](docs/api/transaction.html) | `load` | `GET /api/v1/transactions/{id}/proof` | Required |
| [Transaction](docs/api/transaction.html) | `remove` | `DELETE /api/v1/transactions/{id}/tag` | Required |
| [User](docs/api/user.html) | `list` | `GET /api/v1/user/{merchantID}/userspaged` | Required |
| [User](docs/api/user.html) | `list` | `GET /api/v1/merchants/{merchantID}/users` | Required |
| [User](docs/api/user.html) | `list` | `GET /api/v1/metadata/whoami` | Required |
| [User](docs/api/user.html) | `list` | `GET /api/v1/metadata/whoamitrustedapp` | Required |
| [User](docs/api/user.html) | `list` | `GET /api/v1/user` | Required |
| [User](docs/api/user.html) | `update` | `PUT /api/v1/user/{id}` | Required |
| [UserInvite](docs/api/user_invite.html) | `create` | `POST /api/v1/userinvites/authorise/{id}` | Required |
| [UserInvite](docs/api/user_invite.html) | `create` | `POST /api/v1/userinvites` | Required |
| [UserInvite](docs/api/user_invite.html) | `create` | `POST /api/v1/userinvites/batchcreate` | Required |
| [UserInvite](docs/api/user_invite.html) | `list` | `GET /api/v1/merchants/{merchantID}/userinvitespaged` | Required |
| [UserInvite](docs/api/user_invite.html) | `load` | `GET /api/v1/userinvites/{id}` | Required |
| [UserInvite](docs/api/user_invite.html) | `load` | `GET /api/v1/userinvites/{id}/details` | Required |
| [UserInvite](docs/api/user_invite.html) | `remove` | `DELETE /api/v1/userinvites/{id}` | Required |
| [UserInvite](docs/api/user_invite.html) | `update` | `PUT /api/v1/userinvites/{id}` | Required |
| [Virtual](docs/api/virtual.html) | `create` | `POST /api/v1/accounts/{accountID}/virtual` | Required |
| [Virtual](docs/api/virtual.html) | `update` | `PUT /api/v1/accounts/{accountID}/virtual/{virtualAccountID}` | Required |
| [Webhook](docs/api/webhook.html) | `create` | `POST /api/v1/webhooks` | Required |
| [Webhook](docs/api/webhook.html) | `list` | `GET /api/v1/merchants/{merchantID}/webhooks` | Required |
| [Webhook](docs/api/webhook.html) | `load` | `GET /api/v1/merchants/{merchantID}/webhooks/{id}` | Required |
| [Webhook](docs/api/webhook.html) | `load` | `GET /api/v1/webhooks/{merchantID}` | Required |
| [Webhook](docs/api/webhook.html) | `remove` | `DELETE /api/v1/webhooks/{id}` | Required |
| [Webhook](docs/api/webhook.html) | `update` | `PUT /api/v1/webhooks/{id}` | Required |

## Connect to the API

- v1: `https://api-sandbox.nofrixion.com`

The default credential is sent in the `Authorization` header with the `Bearer` prefix.

JWT Authorization header using the Bearer scheme.&lt;br/&gt;
                      Enter your JWT access token in the text input below.&lt;br/&gt;
                      Example: Bearer eyJhbGciOiJ...

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| [Golang](docs/sdks/go.html) | `go/` | Build from source |
| [Lua](docs/sdks/lua.html) | `lua/` | Build from source |
| [PHP](docs/sdks/php.html) | `php/` | Build from source |
| [Python](docs/sdks/py.html) | `py/` | Build from source |
| [Ruby](docs/sdks/rb.html) | `rb/` | Build from source |
| [TypeScript](docs/sdks/ts.html) | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### [Go CLI](docs/tools/go-cli.html)

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### [Go MCP server](docs/tools/go-mcp.html)

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `nofrixion_list`: List records for an entity. Supported entities: `account`, `beneficiary`, `beneficiary_group`, `card_customer_token`, `consent`, `currency`, `fx_rate`, `merchant`, `merchant_authorisation_setting`, `merchant_direct_debit_mandate_page`, `merchant_pay_by_bank_setting`, `merchant_payment_request_template`, `merchant_token`, `payment_account`, `payment_account_minimal`, `payment_request`, `payment_request_event`, `payment_request_minimal`, `payment_request_result`, `payout`, `payout_keyset_page`, `payrun`, `rule`, `rule_event`, `tag`, `transaction`, `user`, `user_invite`, `webhook`.
- `nofrixion_load`: Load one record for an entity. Supported entities: `account`, `batch`, `beneficiary`, `card_customer_token`, `card_public_key`, `consent`, `fx_rate`, `mandate`, `merchant`, `merchant_payment_request_template`, `merchant_token`, `metadata`, `no_frixion_version`, `payment`, `payment_request`, `payment_request_metric`, `payout`, `payout_metric`, `payrun`, `report_result`, `rule`, `transaction`, `user_invite`, `webhook`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- [`test`](docs/features/test.html): In-memory mock transport for testing without a live server

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the [first-call guide](docs/guides/first-call.html) for the setup sequence.
- Read the [authentication guide](docs/guides/authentication.html) before using protected routes.
- Use the [API reference](docs/api/index.html) for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

