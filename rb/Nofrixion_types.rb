# frozen_string_literal: true

# Typed models for the Nofrixion SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Account entity data model.
#
# @!attribute [rw] account_balance
#   @return [Array, nil]
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] account_identification
#   @return [Array, nil]
#
# @!attribute [rw] account_name
#   @return [String, nil]
#
# @!attribute [rw] account_supplier_name
#   @return [String, nil]
#
# @!attribute [rw] account_type
#   @return [String, nil]
#
# @!attribute [rw] available_balance
#   @return [Float, nil]
#
# @!attribute [rw] available_balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] bank_name
#   @return [String, nil]
#
# @!attribute [rw] consent_id
#   @return [String, nil]
#
# @!attribute [rw] consolidated_account_information
#   @return [Hash, nil]
#
# @!attribute [rw] created_by
#   @return [Hash]
#
# @!attribute [rw] created_by_display_name
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] default_payment_rail
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] detail
#   @return [String, nil]
#
# @!attribute [rw] display_name
#   @return [String, nil]
#
# @!attribute [rw] expiry_date
#   @return [String, nil]
#
# @!attribute [rw] external_account_icon
#   @return [String, nil]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] from_date
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_connected_account
#   @return [Boolean, nil]
#
# @!attribute [rw] is_default
#   @return [Boolean, nil]
#
# @!attribute [rw] is_trust_account
#   @return [Boolean, nil]
#
# @!attribute [rw] is_virtual
#   @return [Boolean, nil]
#
# @!attribute [rw] last_transaction
#   @return [Hash, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_name
#   @return [String, nil]
#
# @!attribute [rw] nickname
#   @return [String, nil]
#
# @!attribute [rw] physical_account_id
#   @return [String, nil]
#
# @!attribute [rw] role_i_d
#   @return [Array, nil]
#
# @!attribute [rw] rule
#   @return [Array, nil]
#
# @!attribute [rw] submitted_payouts_balance
#   @return [Float, nil]
#
# @!attribute [rw] submitted_payouts_balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplier_physical_account_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_sepa_instant_status
#   @return [String, nil]
#
# @!attribute [rw] to_date
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] usage_type
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_connection_status
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_last_synced_at
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_sync_last_failed_at
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_sync_last_failure_reason
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_sync_status
#   @return [String, nil]
#
# @!attribute [rw] xero_unsynchronised_transactions_count
#   @return [Integer, nil]
Account = Struct.new(
  :account_balance,
  :account_id,
  :account_identification,
  :account_name,
  :account_supplier_name,
  :account_type,
  :available_balance,
  :available_balance_minor_unit,
  :balance,
  :balance_minor_unit,
  :bank_name,
  :consent_id,
  :consolidated_account_information,
  :created_by,
  :created_by_display_name,
  :currency,
  :default_payment_rail,
  :description,
  :detail,
  :display_name,
  :expiry_date,
  :external_account_icon,
  :format,
  :from_date,
  :id,
  :identifier,
  :inserted,
  :is_archived,
  :is_connected_account,
  :is_default,
  :is_trust_account,
  :is_virtual,
  :last_transaction,
  :last_updated,
  :merchant_id,
  :merchant_name,
  :nickname,
  :physical_account_id,
  :role_i_d,
  :rule,
  :submitted_payouts_balance,
  :submitted_payouts_balance_minor_unit,
  :summary,
  :supplier_physical_account_id,
  :supplier_sepa_instant_status,
  :to_date,
  :type,
  :usage_type,
  :xero_bank_feed_connection_status,
  :xero_bank_feed_last_synced_at,
  :xero_bank_feed_sync_last_failed_at,
  :xero_bank_feed_sync_last_failure_reason,
  :xero_bank_feed_sync_status,
  :xero_unsynchronised_transactions_count,
  keyword_init: true
)

# Request payload for Account#load.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
AccountLoadMatch = Struct.new(
  :account_id,
  :id,
  :merchant_id,
  keyword_init: true
)

# Request payload for Account#list.
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
AccountListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for Account#create.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
AccountCreateData = Struct.new(
  :account_id,
  :currency,
  keyword_init: true
)

# Request payload for Account#update.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
AccountUpdateData = Struct.new(
  :account_id,
  :amount,
  :id,
  keyword_init: true
)

# Request payload for Account#remove.
#
# @!attribute [rw] id
#   @return [String]
AccountRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Batch entity data model.
#
# @!attribute [rw] approve_url
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] payout
#   @return [Array, nil]
Batch = Struct.new(
  :approve_url,
  :id,
  :payout,
  keyword_init: true
)

# Request payload for Batch#load.
#
# @!attribute [rw] id
#   @return [String]
BatchLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Batch#create.
#
# @!attribute [rw] approve_url
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] payout
#   @return [Array, nil]
BatchCreateData = Struct.new(
  :approve_url,
  :id,
  :payout,
  keyword_init: true
)

# Beneficiary entity data model.
#
# @!attribute [rw] approval_callback_url
#   @return [String, nil]
#
# @!attribute [rw] authentication_method
#   @return [Array, nil]
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] beneficiary
#   @return [Array, nil]
#
# @!attribute [rw] beneficiary_event
#   @return [Array, nil]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] can_update
#   @return [Boolean, nil]
#
# @!attribute [rw] created_by
#   @return [Hash]
#
# @!attribute [rw] created_by_email_address
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] failed_beneficiary
#   @return [Hash, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] last_authorised
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] source_account
#   @return [Array, nil]
#
# @!attribute [rw] source_account_i_d
#   @return [Array, nil]
#
# @!attribute [rw] their_reference
#   @return [String, nil]
Beneficiary = Struct.new(
  :approval_callback_url,
  :authentication_method,
  :authorisation,
  :authorisers_completed_count,
  :authorisers_required_count,
  :beneficiary,
  :beneficiary_event,
  :can_authorise,
  :can_update,
  :created_by,
  :created_by_email_address,
  :currency,
  :destination,
  :failed_beneficiary,
  :has_current_user_authorised,
  :id,
  :inserted,
  :is_enabled,
  :last_authorised,
  :last_updated,
  :merchant_id,
  :name,
  :nonce,
  :source_account,
  :source_account_i_d,
  :their_reference,
  keyword_init: true
)

# Request payload for Beneficiary#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
BeneficiaryLoadMatch = Struct.new(
  :id,
  :merchant_id,
  keyword_init: true
)

# Request payload for Beneficiary#list.
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
BeneficiaryListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for Beneficiary#create.
#
# @!attribute [rw] id
#   @return [String, nil]
BeneficiaryCreateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Beneficiary#update.
#
# @!attribute [rw] id
#   @return [String]
BeneficiaryUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Beneficiary#remove.
#
# @!attribute [rw] id
#   @return [String]
BeneficiaryRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# BeneficiaryGroup entity data model.
#
# @!attribute [rw] group_member
#   @return [Array, nil]
#
# @!attribute [rw] group_name
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String]
BeneficiaryGroup = Struct.new(
  :group_member,
  :group_name,
  :id,
  :inserted,
  :last_updated,
  :merchant_id,
  keyword_init: true
)

# Request payload for BeneficiaryGroup#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
BeneficiaryGroupListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Card entity data model.
#
# @!attribute [rw] authorized_amount
#   @return [String, nil]
#
# @!attribute [rw] currency_code
#   @return [String, nil]
#
# @!attribute [rw] is_payer_authentication_required
#   @return [Boolean, nil]
#
# @!attribute [rw] is_soft_decline
#   @return [Boolean, nil]
#
# @!attribute [rw] payer_authentication_access_token
#   @return [String, nil]
#
# @!attribute [rw] payer_authentication_merchant_data
#   @return [String, nil]
#
# @!attribute [rw] payer_authentication_url
#   @return [String, nil]
#
# @!attribute [rw] payer_authentication_window_height
#   @return [Integer, nil]
#
# @!attribute [rw] payer_authentication_window_width
#   @return [Integer, nil]
#
# @!attribute [rw] payment_request_callback_url
#   @return [String, nil]
#
# @!attribute [rw] payment_request_id
#   @return [String, nil]
#
# @!attribute [rw] request_id
#   @return [String, nil]
#
# @!attribute [rw] response_code
#   @return [String, nil]
#
# @!attribute [rw] response_type
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] three_ds_redirect_url
#   @return [String, nil]
#
# @!attribute [rw] transaction_id
#   @return [String, nil]
Card = Struct.new(
  :authorized_amount,
  :currency_code,
  :is_payer_authentication_required,
  :is_soft_decline,
  :payer_authentication_access_token,
  :payer_authentication_merchant_data,
  :payer_authentication_url,
  :payer_authentication_window_height,
  :payer_authentication_window_width,
  :payment_request_callback_url,
  :payment_request_id,
  :request_id,
  :response_code,
  :response_type,
  :status,
  :three_ds_redirect_url,
  :transaction_id,
  keyword_init: true
)

# Request payload for Card#create.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
CardCreateData = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# CardCustomerToken entity data model.
#
# @!attribute [rw] card_type
#   @return [String, nil]
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] expiry_month
#   @return [String, nil]
#
# @!attribute [rw] expiry_year
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] last_four_digit
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] masked_card_number
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] payment_request_id
#   @return [String, nil]
CardCustomerToken = Struct.new(
  :card_type,
  :customer_email_address,
  :expiry_month,
  :expiry_year,
  :id,
  :inserted,
  :last_four_digit,
  :last_updated,
  :masked_card_number,
  :merchant_id,
  :payment_request_id,
  keyword_init: true
)

# Request payload for CardCustomerToken#load.
#
# @!attribute [rw] customer_email_address
#   @return [String]
CardCustomerTokenLoadMatch = Struct.new(
  :customer_email_address,
  keyword_init: true
)

# Request payload for CardCustomerToken#list.
#
# @!attribute [rw] customer_email_address
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String]
CardCustomerTokenListMatch = Struct.new(
  :customer_email_address,
  :merchant_id,
  keyword_init: true
)

# Request payload for CardCustomerToken#remove.
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
CardCustomerTokenRemoveMatch = Struct.new(
  :customer_email_address,
  :merchant_id,
  :id,
  keyword_init: true
)

# CardPayment entity data model.
#
# @!attribute [rw] authorized_amount
#   @return [String, nil]
#
# @!attribute [rw] currency_code
#   @return [String, nil]
#
# @!attribute [rw] is_payer_authentication_required
#   @return [Boolean, nil]
#
# @!attribute [rw] is_soft_decline
#   @return [Boolean, nil]
#
# @!attribute [rw] payer_authentication_access_token
#   @return [String, nil]
#
# @!attribute [rw] payer_authentication_merchant_data
#   @return [String, nil]
#
# @!attribute [rw] payer_authentication_url
#   @return [String, nil]
#
# @!attribute [rw] payer_authentication_window_height
#   @return [Integer, nil]
#
# @!attribute [rw] payer_authentication_window_width
#   @return [Integer, nil]
#
# @!attribute [rw] payment_request_callback_url
#   @return [String, nil]
#
# @!attribute [rw] payment_request_id
#   @return [String, nil]
#
# @!attribute [rw] request_id
#   @return [String, nil]
#
# @!attribute [rw] response_code
#   @return [String, nil]
#
# @!attribute [rw] response_type
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] three_ds_redirect_url
#   @return [String, nil]
#
# @!attribute [rw] transaction_id
#   @return [String, nil]
CardPayment = Struct.new(
  :authorized_amount,
  :currency_code,
  :is_payer_authentication_required,
  :is_soft_decline,
  :payer_authentication_access_token,
  :payer_authentication_merchant_data,
  :payer_authentication_url,
  :payer_authentication_window_height,
  :payer_authentication_window_width,
  :payment_request_callback_url,
  :payment_request_id,
  :request_id,
  :response_code,
  :response_type,
  :status,
  :three_ds_redirect_url,
  :transaction_id,
  keyword_init: true
)

# Request payload for CardPayment#create.
#
# @!attribute [rw] partial_refund_amount
#   @return [Float, nil]
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
CardPaymentCreateData = Struct.new(
  :partial_refund_amount,
  :paymentrequest_id,
  keyword_init: true
)

# CardPublicKey entity data model.
#
# @!attribute [rw] jwt
#   @return [String, nil]
CardPublicKey = Struct.new(
  :jwt,
  keyword_init: true
)

# Request payload for CardPublicKey#load.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
CardPublicKeyLoadMatch = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# Consent entity data model.
#
# @!attribute [rw] authorisation_url
#   @return [String, nil]
#
# @!attribute [rw] callback_url
#   @return [String, nil]
#
# @!attribute [rw] consent_id
#   @return [String, nil]
#
# @!attribute [rw] email_address
#   @return [String, nil]
#
# @!attribute [rw] expiry_date
#   @return [String, nil]
#
# @!attribute [rw] failure_callback_url
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] institution_id
#   @return [String, nil]
#
# @!attribute [rw] is_connected_account
#   @return [Boolean, nil]
#
# @!attribute [rw] is_enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] provider
#   @return [String, nil]
#
# @!attribute [rw] success_web_hook_url
#   @return [String, nil]
Consent = Struct.new(
  :authorisation_url,
  :callback_url,
  :consent_id,
  :email_address,
  :expiry_date,
  :failure_callback_url,
  :id,
  :inserted,
  :institution_id,
  :is_connected_account,
  :is_enabled,
  :merchant_id,
  :provider,
  :success_web_hook_url,
  keyword_init: true
)

# Request payload for Consent#load.
#
# @!attribute [rw] id
#   @return [String]
ConsentLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Consent#list.
#
# @!attribute [rw] email
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String]
ConsentListMatch = Struct.new(
  :email,
  :merchant_id,
  keyword_init: true
)

# Request payload for Consent#create.
#
# @!attribute [rw] authorisation_url
#   @return [String, nil]
#
# @!attribute [rw] callback_url
#   @return [String, nil]
#
# @!attribute [rw] consent_id
#   @return [String, nil]
#
# @!attribute [rw] email_address
#   @return [String, nil]
#
# @!attribute [rw] expiry_date
#   @return [String, nil]
#
# @!attribute [rw] failure_callback_url
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] institution_id
#   @return [String, nil]
#
# @!attribute [rw] is_connected_account
#   @return [Boolean, nil]
#
# @!attribute [rw] is_enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] provider
#   @return [String, nil]
#
# @!attribute [rw] success_web_hook_url
#   @return [String, nil]
ConsentCreateData = Struct.new(
  :authorisation_url,
  :callback_url,
  :consent_id,
  :email_address,
  :expiry_date,
  :failure_callback_url,
  :id,
  :inserted,
  :institution_id,
  :is_connected_account,
  :is_enabled,
  :merchant_id,
  :provider,
  :success_web_hook_url,
  keyword_init: true
)

# Request payload for Consent#update.
#
# @!attribute [rw] id
#   @return [String]
ConsentUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Consent#remove.
#
# @!attribute [rw] id
#   @return [String]
ConsentRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Currency entity data model.
#
# @!attribute [rw] code
#   @return [String, nil]
#
# @!attribute [rw] decimal
#   @return [Integer, nil]
#
# @!attribute [rw] is_fiat
#   @return [Boolean, nil]
#
# @!attribute [rw] iso4217_alpha_code
#   @return [String, nil]
#
# @!attribute [rw] iso4217_numeric_code
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
Currency = Struct.new(
  :code,
  :decimal,
  :is_fiat,
  :iso4217_alpha_code,
  :iso4217_numeric_code,
  :symbol,
  keyword_init: true
)

# Request payload for Currency#list.
#
# @!attribute [rw] code
#   @return [String, nil]
#
# @!attribute [rw] decimal
#   @return [Integer, nil]
#
# @!attribute [rw] is_fiat
#   @return [Boolean, nil]
#
# @!attribute [rw] iso4217_alpha_code
#   @return [String, nil]
#
# @!attribute [rw] iso4217_numeric_code
#   @return [String, nil]
#
# @!attribute [rw] symbol
#   @return [String, nil]
CurrencyListMatch = Struct.new(
  :code,
  :decimal,
  :is_fiat,
  :iso4217_alpha_code,
  :iso4217_numeric_code,
  :symbol,
  keyword_init: true
)

# DirectDebitBatchSubmit entity data model.
#
# @!attribute [rw] failed_submission
#   @return [Hash, nil]
#
# @!attribute [rw] successful_submission
#   @return [Array, nil]
DirectDebitBatchSubmit = Struct.new(
  :failed_submission,
  :successful_submission,
  keyword_init: true
)

# Request payload for DirectDebitBatchSubmit#create.
#
# @!attribute [rw] failed_submission
#   @return [Hash, nil]
#
# @!attribute [rw] successful_submission
#   @return [Array, nil]
DirectDebitBatchSubmitCreateData = Struct.new(
  :failed_submission,
  :successful_submission,
  keyword_init: true
)

# FxRate entity data model.
#
# @!attribute [rw] destination_currency
#   @return [String, nil]
#
# @!attribute [rw] exchange_rate
#   @return [Float, nil]
#
# @!attribute [rw] expiry_time
#   @return [String, nil]
#
# @!attribute [rw] quote_id
#   @return [String, nil]
#
# @!attribute [rw] source_currency
#   @return [String, nil]
FxRate = Struct.new(
  :destination_currency,
  :exchange_rate,
  :expiry_time,
  :quote_id,
  :source_currency,
  keyword_init: true
)

# Request payload for FxRate#load.
#
# @!attribute [rw] destination
#   @return [String]
#
# @!attribute [rw] source
#   @return [String]
#
# @!attribute [rw] valid_for_minute
#   @return [Integer]
FxRateLoadMatch = Struct.new(
  :destination,
  :source,
  :valid_for_minute,
  keyword_init: true
)

# Request payload for FxRate#list.
#
# @!attribute [rw] destination
#   @return [String]
#
# @!attribute [rw] source
#   @return [String]
FxRateListMatch = Struct.new(
  :destination,
  :source,
  keyword_init: true
)

# IPayment entity data model.
#
# @!attribute [rw] payment_request_id
#   @return [String, nil]
#
# @!attribute [rw] response_type
#   @return [String, nil]
IPayment = Struct.new(
  :payment_request_id,
  :response_type,
  keyword_init: true
)

# Request payload for IPayment#create.
#
# @!attribute [rw] payment_request_id
#   @return [String, nil]
#
# @!attribute [rw] response_type
#   @return [String, nil]
IPaymentCreateData = Struct.new(
  :payment_request_id,
  :response_type,
  keyword_init: true
)

# Mandate entity data model.
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] address_line1
#   @return [String]
#
# @!attribute [rw] address_line2
#   @return [String, nil]
#
# @!attribute [rw] approved_at
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String]
#
# @!attribute [rw] country_code
#   @return [String]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customer_account_number
#   @return [String, nil]
#
# @!attribute [rw] customer_city
#   @return [String, nil]
#
# @!attribute [rw] customer_country_code
#   @return [String, nil]
#
# @!attribute [rw] customer_country_name
#   @return [String, nil]
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] customer_first_name
#   @return [String, nil]
#
# @!attribute [rw] customer_iban
#   @return [String, nil]
#
# @!attribute [rw] customer_last_name
#   @return [String, nil]
#
# @!attribute [rw] customer_sort_code
#   @return [String, nil]
#
# @!attribute [rw] email_address
#   @return [String]
#
# @!attribute [rw] first_name
#   @return [String]
#
# @!attribute [rw] iban
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_recurring
#   @return [Boolean, nil]
#
# @!attribute [rw] last_name
#   @return [String]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] postal_code
#   @return [String]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] sort_code
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplier_bank_account_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_customer_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_mandate_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_name
#   @return [String, nil]
#
# @!attribute [rw] supplier_status
#   @return [String, nil]
Mandate = Struct.new(
  :account_number,
  :address_line1,
  :address_line2,
  :approved_at,
  :city,
  :country_code,
  :currency,
  :customer_account_number,
  :customer_city,
  :customer_country_code,
  :customer_country_name,
  :customer_email_address,
  :customer_first_name,
  :customer_iban,
  :customer_last_name,
  :customer_sort_code,
  :email_address,
  :first_name,
  :iban,
  :id,
  :inserted,
  :is_recurring,
  :last_name,
  :last_updated,
  :merchant_id,
  :postal_code,
  :reference,
  :sort_code,
  :status,
  :supplier_bank_account_id,
  :supplier_customer_id,
  :supplier_mandate_id,
  :supplier_name,
  :supplier_status,
  keyword_init: true
)

# Request payload for Mandate#load.
#
# @!attribute [rw] id
#   @return [String]
MandateLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Mandate#create.
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] address_line1
#   @return [String]
#
# @!attribute [rw] address_line2
#   @return [String, nil]
#
# @!attribute [rw] approved_at
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String]
#
# @!attribute [rw] country_code
#   @return [String]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customer_account_number
#   @return [String, nil]
#
# @!attribute [rw] customer_city
#   @return [String, nil]
#
# @!attribute [rw] customer_country_code
#   @return [String, nil]
#
# @!attribute [rw] customer_country_name
#   @return [String, nil]
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] customer_first_name
#   @return [String, nil]
#
# @!attribute [rw] customer_iban
#   @return [String, nil]
#
# @!attribute [rw] customer_last_name
#   @return [String, nil]
#
# @!attribute [rw] customer_sort_code
#   @return [String, nil]
#
# @!attribute [rw] email_address
#   @return [String]
#
# @!attribute [rw] first_name
#   @return [String]
#
# @!attribute [rw] iban
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_recurring
#   @return [Boolean, nil]
#
# @!attribute [rw] last_name
#   @return [String]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] postal_code
#   @return [String]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] sort_code
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplier_bank_account_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_customer_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_mandate_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_name
#   @return [String, nil]
#
# @!attribute [rw] supplier_status
#   @return [String, nil]
MandateCreateData = Struct.new(
  :account_number,
  :address_line1,
  :address_line2,
  :approved_at,
  :city,
  :country_code,
  :currency,
  :customer_account_number,
  :customer_city,
  :customer_country_code,
  :customer_country_name,
  :customer_email_address,
  :customer_first_name,
  :customer_iban,
  :customer_last_name,
  :customer_sort_code,
  :email_address,
  :first_name,
  :iban,
  :id,
  :inserted,
  :is_recurring,
  :last_name,
  :last_updated,
  :merchant_id,
  :postal_code,
  :reference,
  :sort_code,
  :status,
  :supplier_bank_account_id,
  :supplier_customer_id,
  :supplier_mandate_id,
  :supplier_name,
  :supplier_status,
  keyword_init: true
)

# Merchant entity data model.
#
# @!attribute [rw] account_currency
#   @return [Array, nil]
#
# @!attribute [rw] can_have_trust_account
#   @return [Boolean, nil]
#
# @!attribute [rw] card_payment_processor
#   @return [String, nil]
#
# @!attribute [rw] company_id
#   @return [String, nil]
#
# @!attribute [rw] display_qr_on_hosted_pay
#   @return [Boolean, nil]
#
# @!attribute [rw] hosted_pay_version
#   @return [Integer, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_blocked
#   @return [Boolean, nil]
#
# @!attribute [rw] is_exited
#   @return [Boolean, nil]
#
# @!attribute [rw] is_suspended
#   @return [Boolean, nil]
#
# @!attribute [rw] jurisdiction
#   @return [String, nil]
#
# @!attribute [rw] logo_url_png
#   @return [String, nil]
#
# @!attribute [rw] logo_url_svg
#   @return [String, nil]
#
# @!attribute [rw] merchant_category_code
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] note
#   @return [String, nil]
#
# @!attribute [rw] parent_merchant
#   @return [Hash, nil]
#
# @!attribute [rw] payment_account
#   @return [Array, nil]
#
# @!attribute [rw] payment_account_limit
#   @return [Integer, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] short_name
#   @return [String, nil]
#
# @!attribute [rw] supported_payment_methods_list
#   @return [Array, nil]
#
# @!attribute [rw] suspension_reason
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [Array, nil]
#
# @!attribute [rw] time_zone_id
#   @return [String, nil]
#
# @!attribute [rw] trading_name
#   @return [String, nil]
#
# @!attribute [rw] web_hook_limit
#   @return [Integer, nil]
#
# @!attribute [rw] your_role_name
#   @return [String, nil]
Merchant = Struct.new(
  :account_currency,
  :can_have_trust_account,
  :card_payment_processor,
  :company_id,
  :display_qr_on_hosted_pay,
  :hosted_pay_version,
  :id,
  :inserted,
  :is_blocked,
  :is_exited,
  :is_suspended,
  :jurisdiction,
  :logo_url_png,
  :logo_url_svg,
  :merchant_category_code,
  :name,
  :note,
  :parent_merchant,
  :payment_account,
  :payment_account_limit,
  :reason,
  :short_name,
  :supported_payment_methods_list,
  :suspension_reason,
  :tag,
  :time_zone_id,
  :trading_name,
  :web_hook_limit,
  :your_role_name,
  keyword_init: true
)

# Request payload for Merchant#load.
#
# @!attribute [rw] id
#   @return [String]
MerchantLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Merchant#list.
#
# @!attribute [rw] account_currency
#   @return [Array, nil]
#
# @!attribute [rw] can_have_trust_account
#   @return [Boolean, nil]
#
# @!attribute [rw] card_payment_processor
#   @return [String, nil]
#
# @!attribute [rw] company_id
#   @return [String, nil]
#
# @!attribute [rw] display_qr_on_hosted_pay
#   @return [Boolean, nil]
#
# @!attribute [rw] hosted_pay_version
#   @return [Integer, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_blocked
#   @return [Boolean, nil]
#
# @!attribute [rw] is_exited
#   @return [Boolean, nil]
#
# @!attribute [rw] is_suspended
#   @return [Boolean, nil]
#
# @!attribute [rw] jurisdiction
#   @return [String, nil]
#
# @!attribute [rw] logo_url_png
#   @return [String, nil]
#
# @!attribute [rw] logo_url_svg
#   @return [String, nil]
#
# @!attribute [rw] merchant_category_code
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] note
#   @return [String, nil]
#
# @!attribute [rw] parent_merchant
#   @return [Hash, nil]
#
# @!attribute [rw] payment_account
#   @return [Array, nil]
#
# @!attribute [rw] payment_account_limit
#   @return [Integer, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] short_name
#   @return [String, nil]
#
# @!attribute [rw] supported_payment_methods_list
#   @return [Array, nil]
#
# @!attribute [rw] suspension_reason
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [Array, nil]
#
# @!attribute [rw] time_zone_id
#   @return [String, nil]
#
# @!attribute [rw] trading_name
#   @return [String, nil]
#
# @!attribute [rw] web_hook_limit
#   @return [Integer, nil]
#
# @!attribute [rw] your_role_name
#   @return [String, nil]
MerchantListMatch = Struct.new(
  :account_currency,
  :can_have_trust_account,
  :card_payment_processor,
  :company_id,
  :display_qr_on_hosted_pay,
  :hosted_pay_version,
  :id,
  :inserted,
  :is_blocked,
  :is_exited,
  :is_suspended,
  :jurisdiction,
  :logo_url_png,
  :logo_url_svg,
  :merchant_category_code,
  :name,
  :note,
  :parent_merchant,
  :payment_account,
  :payment_account_limit,
  :reason,
  :short_name,
  :supported_payment_methods_list,
  :suspension_reason,
  :tag,
  :time_zone_id,
  :trading_name,
  :web_hook_limit,
  :your_role_name,
  keyword_init: true
)

# Request payload for Merchant#update.
#
# @!attribute [rw] id
#   @return [String]
MerchantUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Merchant#remove.
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] user_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] tag_id
#   @return [String, nil]
MerchantRemoveMatch = Struct.new(
  :id,
  :user_id,
  :merchant_id,
  :tag_id,
  keyword_init: true
)

# MerchantAuthorisationSetting entity data model.
#
# @!attribute [rw] amount_lower
#   @return [Float, nil]
#
# @!attribute [rw] amount_upper
#   @return [Float, nil]
#
# @!attribute [rw] authorisation_type
#   @return [String, nil]
#
# @!attribute [rw] beneficiaries_only
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] last_editor_cant_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] number_of_authoriser
#   @return [Integer, nil]
#
# @!attribute [rw] role_setting
#   @return [Array, nil]
MerchantAuthorisationSetting = Struct.new(
  :amount_lower,
  :amount_upper,
  :authorisation_type,
  :beneficiaries_only,
  :id,
  :inserted,
  :last_editor_cant_authorise,
  :last_updated,
  :merchant_id,
  :number_of_authoriser,
  :role_setting,
  keyword_init: true
)

# Request payload for MerchantAuthorisationSetting#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
MerchantAuthorisationSettingListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# MerchantDirectDebitMandatePage entity data model.
#
# @!attribute [rw] approved_at
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customer_account_number
#   @return [String, nil]
#
# @!attribute [rw] customer_city
#   @return [String, nil]
#
# @!attribute [rw] customer_country_code
#   @return [String, nil]
#
# @!attribute [rw] customer_country_name
#   @return [String, nil]
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] customer_first_name
#   @return [String, nil]
#
# @!attribute [rw] customer_iban
#   @return [String, nil]
#
# @!attribute [rw] customer_last_name
#   @return [String, nil]
#
# @!attribute [rw] customer_sort_code
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_recurring
#   @return [Boolean, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplier_bank_account_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_customer_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_mandate_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_name
#   @return [String, nil]
#
# @!attribute [rw] supplier_status
#   @return [String, nil]
MerchantDirectDebitMandatePage = Struct.new(
  :approved_at,
  :currency,
  :customer_account_number,
  :customer_city,
  :customer_country_code,
  :customer_country_name,
  :customer_email_address,
  :customer_first_name,
  :customer_iban,
  :customer_last_name,
  :customer_sort_code,
  :id,
  :inserted,
  :is_recurring,
  :last_updated,
  :merchant_id,
  :reference,
  :status,
  :supplier_bank_account_id,
  :supplier_customer_id,
  :supplier_mandate_id,
  :supplier_name,
  :supplier_status,
  keyword_init: true
)

# Request payload for MerchantDirectDebitMandatePage#list.
#
# @!attribute [rw] approved_at
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customer_account_number
#   @return [String, nil]
#
# @!attribute [rw] customer_city
#   @return [String, nil]
#
# @!attribute [rw] customer_country_code
#   @return [String, nil]
#
# @!attribute [rw] customer_country_name
#   @return [String, nil]
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] customer_first_name
#   @return [String, nil]
#
# @!attribute [rw] customer_iban
#   @return [String, nil]
#
# @!attribute [rw] customer_last_name
#   @return [String, nil]
#
# @!attribute [rw] customer_sort_code
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_recurring
#   @return [Boolean, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplier_bank_account_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_customer_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_mandate_id
#   @return [String, nil]
#
# @!attribute [rw] supplier_name
#   @return [String, nil]
#
# @!attribute [rw] supplier_status
#   @return [String, nil]
MerchantDirectDebitMandatePageListMatch = Struct.new(
  :approved_at,
  :currency,
  :customer_account_number,
  :customer_city,
  :customer_country_code,
  :customer_country_name,
  :customer_email_address,
  :customer_first_name,
  :customer_iban,
  :customer_last_name,
  :customer_sort_code,
  :id,
  :inserted,
  :is_recurring,
  :last_updated,
  :merchant_id,
  :reference,
  :status,
  :supplier_bank_account_id,
  :supplier_customer_id,
  :supplier_mandate_id,
  :supplier_name,
  :supplier_status,
  keyword_init: true
)

# MerchantPayByBankSetting entity data model.
#
# @!attribute [rw] bank_country_code
#   @return [Array, nil]
#
# @!attribute [rw] bank_id
#   @return [String, nil]
#
# @!attribute [rw] bank_name
#   @return [String, nil]
#
# @!attribute [rw] business_institution_id
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] logo
#   @return [String, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] message_image_url
#   @return [String, nil]
#
# @!attribute [rw] order
#   @return [Integer, nil]
#
# @!attribute [rw] personal_institution_id
#   @return [String, nil]
#
# @!attribute [rw] processor
#   @return [String, nil]
#
# @!attribute [rw] warning_heading
#   @return [String, nil]
#
# @!attribute [rw] warning_message
#   @return [String, nil]
MerchantPayByBankSetting = Struct.new(
  :bank_country_code,
  :bank_id,
  :bank_name,
  :business_institution_id,
  :currency,
  :logo,
  :message,
  :message_image_url,
  :order,
  :personal_institution_id,
  :processor,
  :warning_heading,
  :warning_message,
  keyword_init: true
)

# Request payload for MerchantPayByBankSetting#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
MerchantPayByBankSettingListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# MerchantPaymentRequestTemplate entity data model.
#
# @!attribute [rw] description
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] template
#   @return [Hash]
MerchantPaymentRequestTemplate = Struct.new(
  :description,
  :id,
  :inserted,
  :last_updated,
  :merchant_id,
  :name,
  :template,
  keyword_init: true
)

# Request payload for MerchantPaymentRequestTemplate#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
MerchantPaymentRequestTemplateLoadMatch = Struct.new(
  :id,
  :paymentrequest_id,
  keyword_init: true
)

# Request payload for MerchantPaymentRequestTemplate#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
MerchantPaymentRequestTemplateListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for MerchantPaymentRequestTemplate#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
MerchantPaymentRequestTemplateUpdateData = Struct.new(
  :id,
  :paymentrequest_id,
  keyword_init: true
)

# Request payload for MerchantPaymentRequestTemplate#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
MerchantPaymentRequestTemplateRemoveMatch = Struct.new(
  :id,
  :paymentrequest_id,
  keyword_init: true
)

# MerchantToken entity data model.
#
# @!attribute [rw] authentication_method
#   @return [Array, nil]
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] expires_at
#   @return [String, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] hmac_algorithm
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] ip_address_whitelist
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] last_authorised
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String]
#
# @!attribute [rw] permission_type
#   @return [Array, nil]
#
# @!attribute [rw] request_signature_version
#   @return [Integer, nil]
#
# @!attribute [rw] shared_secret_algorithm
#   @return [String, nil]
#
# @!attribute [rw] shared_secret_base64
#   @return [String, nil]
#
# @!attribute [rw] token
#   @return [String, nil]
MerchantToken = Struct.new(
  :authentication_method,
  :authorisation,
  :authorisers_completed_count,
  :authorisers_required_count,
  :can_authorise,
  :description,
  :expires_at,
  :has_current_user_authorised,
  :hmac_algorithm,
  :id,
  :inserted,
  :ip_address_whitelist,
  :is_archived,
  :is_enabled,
  :last_authorised,
  :last_updated,
  :merchant_id,
  :nonce,
  :permission_type,
  :request_signature_version,
  :shared_secret_algorithm,
  :shared_secret_base64,
  :token,
  keyword_init: true
)

# Request payload for MerchantToken#load.
#
# @!attribute [rw] id
#   @return [String]
MerchantTokenLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for MerchantToken#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
MerchantTokenListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for MerchantToken#create.
#
# @!attribute [rw] authentication_method
#   @return [Array, nil]
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] expires_at
#   @return [String, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] hmac_algorithm
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] ip_address_whitelist
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] last_authorised
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String]
#
# @!attribute [rw] permission_type
#   @return [Array, nil]
#
# @!attribute [rw] request_signature_version
#   @return [Integer, nil]
#
# @!attribute [rw] shared_secret_algorithm
#   @return [String, nil]
#
# @!attribute [rw] shared_secret_base64
#   @return [String, nil]
#
# @!attribute [rw] token
#   @return [String, nil]
MerchantTokenCreateData = Struct.new(
  :authentication_method,
  :authorisation,
  :authorisers_completed_count,
  :authorisers_required_count,
  :can_authorise,
  :description,
  :expires_at,
  :has_current_user_authorised,
  :hmac_algorithm,
  :id,
  :inserted,
  :ip_address_whitelist,
  :is_archived,
  :is_enabled,
  :last_authorised,
  :last_updated,
  :merchant_id,
  :nonce,
  :permission_type,
  :request_signature_version,
  :shared_secret_algorithm,
  :shared_secret_base64,
  :token,
  keyword_init: true
)

# Request payload for MerchantToken#update.
#
# @!attribute [rw] id
#   @return [String]
MerchantTokenUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# Metadata entity data model.
class Metadata
end

# Request payload for Metadata#load.
class MetadataLoadMatch
end

# NoFrixionVersion entity data model.
#
# @!attribute [rw] build_version
#   @return [Integer, nil]
#
# @!attribute [rw] major_version
#   @return [Integer, nil]
#
# @!attribute [rw] minor_version
#   @return [Integer, nil]
#
# @!attribute [rw] release_name
#   @return [String, nil]
NoFrixionVersion = Struct.new(
  :build_version,
  :major_version,
  :minor_version,
  :release_name,
  keyword_init: true
)

# Request payload for NoFrixionVersion#load.
#
# @!attribute [rw] build_version
#   @return [Integer, nil]
#
# @!attribute [rw] major_version
#   @return [Integer, nil]
#
# @!attribute [rw] minor_version
#   @return [Integer, nil]
#
# @!attribute [rw] release_name
#   @return [String, nil]
NoFrixionVersionLoadMatch = Struct.new(
  :build_version,
  :major_version,
  :minor_version,
  :release_name,
  keyword_init: true
)

# OpenBanking entity data model.
class OpenBanking
end

# Request payload for OpenBanking#create.
#
# @!attribute [rw] account_id
#   @return [String]
OpenBankingCreateData = Struct.new(
  :account_id,
  keyword_init: true
)

# Request payload for OpenBanking#remove.
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] account_id
#   @return [String, nil]
OpenBankingRemoveMatch = Struct.new(
  :email,
  :merchant_id,
  :account_id,
  keyword_init: true
)

# Payeeverification entity data model.
#
# @!attribute [rw] account_name
#   @return [String]
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] iban
#   @return [String]
#
# @!attribute [rw] payee_verified_account_name
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [String, nil]
#
# @!attribute [rw] secondary_identification
#   @return [String, nil]
#
# @!attribute [rw] sort_code
#   @return [String, nil]
Payeeverification = Struct.new(
  :account_name,
  :account_number,
  :iban,
  :payee_verified_account_name,
  :result,
  :secondary_identification,
  :sort_code,
  keyword_init: true
)

# Request payload for Payeeverification#create.
#
# @!attribute [rw] account_name
#   @return [String]
#
# @!attribute [rw] account_number
#   @return [String, nil]
#
# @!attribute [rw] iban
#   @return [String]
#
# @!attribute [rw] payee_verified_account_name
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [String, nil]
#
# @!attribute [rw] secondary_identification
#   @return [String, nil]
#
# @!attribute [rw] sort_code
#   @return [String, nil]
PayeeverificationCreateData = Struct.new(
  :account_name,
  :account_number,
  :iban,
  :payee_verified_account_name,
  :result,
  :secondary_identification,
  :sort_code,
  keyword_init: true
)

# Payment entity data model.
#
# @!attribute [rw] address
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amount_pending
#   @return [Float, nil]
#
# @!attribute [rw] amount_received
#   @return [Float, nil]
#
# @!attribute [rw] amount_refunded
#   @return [Float, nil]
#
# @!attribute [rw] auto_send_receipt
#   @return [Boolean, nil]
#
# @!attribute [rw] base_origin_url
#   @return [String, nil]
#
# @!attribute [rw] callback_url
#   @return [String, nil]
#
# @!attribute [rw] card_authorize_only
#   @return [Boolean, nil]
#
# @!attribute [rw] card_create_token
#   @return [Boolean, nil]
#
# @!attribute [rw] card_create_token_mode
#   @return [String, nil]
#
# @!attribute [rw] card_ignore_cvn
#   @return [Boolean, nil]
#
# @!attribute [rw] card_no_payer_authentication
#   @return [Boolean, nil]
#
# @!attribute [rw] card_processor_merchant_id
#   @return [String, nil]
#
# @!attribute [rw] card_stripe_payment_intent_id
#   @return [String, nil]
#
# @!attribute [rw] card_stripe_payment_intent_secret
#   @return [String, nil]
#
# @!attribute [rw] card_transmit_raw_detail
#   @return [Boolean, nil]
#
# @!attribute [rw] created_by_user
#   @return [Hash]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] custom_field
#   @return [Array, nil]
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] customer_id
#   @return [String, nil]
#
# @!attribute [rw] customer_name
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination_account
#   @return [Hash, nil]
#
# @!attribute [rw] direct_debit_payment
#   @return [Hash, nil]
#
# @!attribute [rw] due_date
#   @return [String, nil]
#
# @!attribute [rw] event
#   @return [Array, nil]
#
# @!attribute [rw] failure_callback_url
#   @return [String, nil]
#
# @!attribute [rw] field_display_setting
#   @return [Array, nil]
#
# @!attribute [rw] formatted_amount
#   @return [String, nil]
#
# @!attribute [rw] hosted_pay_checkout_url
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignore_address_verification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] inserted_sortable
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] lightning_invoice
#   @return [String, nil]
#
# @!attribute [rw] lightning_invoice_expires_at
#   @return [String, nil]
#
# @!attribute [rw] merchant_direct_debit_mandate_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_token_description
#   @return [String, nil]
#
# @!attribute [rw] notification_email_address
#   @return [String, nil]
#
# @!attribute [rw] notification_role_i_d
#   @return [Array, nil]
#
# @!attribute [rw] order_id
#   @return [String, nil]
#
# @!attribute [rw] partial_payment_method
#   @return [String, nil]
#
# @!attribute [rw] partial_payment_step
#   @return [String, nil]
#
# @!attribute [rw] payment_attempt
#   @return [Array, nil]
#
# @!attribute [rw] payment_method
#   @return [Array, nil]
#
# @!attribute [rw] payment_processor
#   @return [String, nil]
#
# @!attribute [rw] payrun_id
#   @return [String, nil]
#
# @!attribute [rw] pisp_account_id
#   @return [String, nil]
#
# @!attribute [rw] priority_bank_id
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandbox_settle_delay_in_second
#   @return [Integer, nil]
#
# @!attribute [rw] shipping_address
#   @return [Hash, nil]
#
# @!attribute [rw] shipping_address_city
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_country_code
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_county
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_line1
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_line2
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_post_code
#   @return [String, nil]
#
# @!attribute [rw] shipping_email
#   @return [String, nil]
#
# @!attribute [rw] shipping_first_name
#   @return [String, nil]
#
# @!attribute [rw] shipping_last_name
#   @return [String, nil]
#
# @!attribute [rw] shipping_phone
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] success_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [Array, nil]
#
# @!attribute [rw] tag_id
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenised_card
#   @return [Array, nil]
#
# @!attribute [rw] transaction
#   @return [Array, nil]
#
# @!attribute [rw] use_hosted_payment_page
#   @return [Boolean, nil]
Payment = Struct.new(
  :address,
  :amount,
  :amount_pending,
  :amount_received,
  :amount_refunded,
  :auto_send_receipt,
  :base_origin_url,
  :callback_url,
  :card_authorize_only,
  :card_create_token,
  :card_create_token_mode,
  :card_ignore_cvn,
  :card_no_payer_authentication,
  :card_processor_merchant_id,
  :card_stripe_payment_intent_id,
  :card_stripe_payment_intent_secret,
  :card_transmit_raw_detail,
  :created_by_user,
  :currency,
  :custom_field,
  :customer_email_address,
  :customer_id,
  :customer_name,
  :description,
  :destination_account,
  :direct_debit_payment,
  :due_date,
  :event,
  :failure_callback_url,
  :field_display_setting,
  :formatted_amount,
  :hosted_pay_checkout_url,
  :id,
  :ignore_address_verification,
  :inserted,
  :inserted_sortable,
  :is_archived,
  :jwk,
  :last_updated,
  :lightning_invoice,
  :lightning_invoice_expires_at,
  :merchant_direct_debit_mandate_id,
  :merchant_id,
  :merchant_token_description,
  :notification_email_address,
  :notification_role_i_d,
  :order_id,
  :partial_payment_method,
  :partial_payment_step,
  :payment_attempt,
  :payment_method,
  :payment_processor,
  :payrun_id,
  :pisp_account_id,
  :priority_bank_id,
  :result,
  :sandbox_settle_delay_in_second,
  :shipping_address,
  :shipping_address_city,
  :shipping_address_country_code,
  :shipping_address_county,
  :shipping_address_line1,
  :shipping_address_line2,
  :shipping_address_post_code,
  :shipping_email,
  :shipping_first_name,
  :shipping_last_name,
  :shipping_phone,
  :status,
  :success_web_hook_url,
  :tag,
  :tag_id,
  :title,
  :tokenised_card,
  :transaction,
  :use_hosted_payment_page,
  keyword_init: true
)

# Request payload for Payment#load.
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] order_id
#   @return [String, nil]
PaymentLoadMatch = Struct.new(
  :id,
  :order_id,
  keyword_init: true
)

# Request payload for Payment#create.
#
# @!attribute [rw] address
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amount_pending
#   @return [Float, nil]
#
# @!attribute [rw] amount_received
#   @return [Float, nil]
#
# @!attribute [rw] amount_refunded
#   @return [Float, nil]
#
# @!attribute [rw] auto_send_receipt
#   @return [Boolean, nil]
#
# @!attribute [rw] base_origin_url
#   @return [String, nil]
#
# @!attribute [rw] callback_url
#   @return [String, nil]
#
# @!attribute [rw] card_authorize_only
#   @return [Boolean, nil]
#
# @!attribute [rw] card_create_token
#   @return [Boolean, nil]
#
# @!attribute [rw] card_create_token_mode
#   @return [String, nil]
#
# @!attribute [rw] card_ignore_cvn
#   @return [Boolean, nil]
#
# @!attribute [rw] card_no_payer_authentication
#   @return [Boolean, nil]
#
# @!attribute [rw] card_processor_merchant_id
#   @return [String, nil]
#
# @!attribute [rw] card_stripe_payment_intent_id
#   @return [String, nil]
#
# @!attribute [rw] card_stripe_payment_intent_secret
#   @return [String, nil]
#
# @!attribute [rw] card_transmit_raw_detail
#   @return [Boolean, nil]
#
# @!attribute [rw] created_by_user
#   @return [Hash]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] custom_field
#   @return [Array, nil]
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] customer_id
#   @return [String, nil]
#
# @!attribute [rw] customer_name
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination_account
#   @return [Hash, nil]
#
# @!attribute [rw] direct_debit_payment
#   @return [Hash, nil]
#
# @!attribute [rw] due_date
#   @return [String, nil]
#
# @!attribute [rw] event
#   @return [Array, nil]
#
# @!attribute [rw] failure_callback_url
#   @return [String, nil]
#
# @!attribute [rw] field_display_setting
#   @return [Array, nil]
#
# @!attribute [rw] formatted_amount
#   @return [String, nil]
#
# @!attribute [rw] hosted_pay_checkout_url
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignore_address_verification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] inserted_sortable
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] lightning_invoice
#   @return [String, nil]
#
# @!attribute [rw] lightning_invoice_expires_at
#   @return [String, nil]
#
# @!attribute [rw] merchant_direct_debit_mandate_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_token_description
#   @return [String, nil]
#
# @!attribute [rw] notification_email_address
#   @return [String, nil]
#
# @!attribute [rw] notification_role_i_d
#   @return [Array, nil]
#
# @!attribute [rw] order_id
#   @return [String, nil]
#
# @!attribute [rw] partial_payment_method
#   @return [String, nil]
#
# @!attribute [rw] partial_payment_step
#   @return [String, nil]
#
# @!attribute [rw] payment_attempt
#   @return [Array, nil]
#
# @!attribute [rw] payment_method
#   @return [Array, nil]
#
# @!attribute [rw] payment_processor
#   @return [String, nil]
#
# @!attribute [rw] payrun_id
#   @return [String, nil]
#
# @!attribute [rw] pisp_account_id
#   @return [String, nil]
#
# @!attribute [rw] priority_bank_id
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandbox_settle_delay_in_second
#   @return [Integer, nil]
#
# @!attribute [rw] shipping_address
#   @return [Hash, nil]
#
# @!attribute [rw] shipping_address_city
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_country_code
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_county
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_line1
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_line2
#   @return [String, nil]
#
# @!attribute [rw] shipping_address_post_code
#   @return [String, nil]
#
# @!attribute [rw] shipping_email
#   @return [String, nil]
#
# @!attribute [rw] shipping_first_name
#   @return [String, nil]
#
# @!attribute [rw] shipping_last_name
#   @return [String, nil]
#
# @!attribute [rw] shipping_phone
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] success_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [Array, nil]
#
# @!attribute [rw] tag_id
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenised_card
#   @return [Array, nil]
#
# @!attribute [rw] transaction
#   @return [Array, nil]
#
# @!attribute [rw] use_hosted_payment_page
#   @return [Boolean, nil]
PaymentCreateData = Struct.new(
  :address,
  :amount,
  :amount_pending,
  :amount_received,
  :amount_refunded,
  :auto_send_receipt,
  :base_origin_url,
  :callback_url,
  :card_authorize_only,
  :card_create_token,
  :card_create_token_mode,
  :card_ignore_cvn,
  :card_no_payer_authentication,
  :card_processor_merchant_id,
  :card_stripe_payment_intent_id,
  :card_stripe_payment_intent_secret,
  :card_transmit_raw_detail,
  :created_by_user,
  :currency,
  :custom_field,
  :customer_email_address,
  :customer_id,
  :customer_name,
  :description,
  :destination_account,
  :direct_debit_payment,
  :due_date,
  :event,
  :failure_callback_url,
  :field_display_setting,
  :formatted_amount,
  :hosted_pay_checkout_url,
  :id,
  :ignore_address_verification,
  :inserted,
  :inserted_sortable,
  :is_archived,
  :jwk,
  :last_updated,
  :lightning_invoice,
  :lightning_invoice_expires_at,
  :merchant_direct_debit_mandate_id,
  :merchant_id,
  :merchant_token_description,
  :notification_email_address,
  :notification_role_i_d,
  :order_id,
  :partial_payment_method,
  :partial_payment_step,
  :payment_attempt,
  :payment_method,
  :payment_processor,
  :payrun_id,
  :pisp_account_id,
  :priority_bank_id,
  :result,
  :sandbox_settle_delay_in_second,
  :shipping_address,
  :shipping_address_city,
  :shipping_address_country_code,
  :shipping_address_county,
  :shipping_address_line1,
  :shipping_address_line2,
  :shipping_address_post_code,
  :shipping_email,
  :shipping_first_name,
  :shipping_last_name,
  :shipping_phone,
  :status,
  :success_web_hook_url,
  :tag,
  :tag_id,
  :title,
  :tokenised_card,
  :transaction,
  :use_hosted_payment_page,
  keyword_init: true
)

# Request payload for Payment#update.
#
# @!attribute [rw] id
#   @return [String]
PaymentUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# PaymentAccount entity data model.
#
# @!attribute [rw] account_name
#   @return [String, nil]
#
# @!attribute [rw] account_supplier_name
#   @return [String, nil]
#
# @!attribute [rw] available_balance
#   @return [Float, nil]
#
# @!attribute [rw] available_balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] bank_name
#   @return [String, nil]
#
# @!attribute [rw] consent_id
#   @return [String, nil]
#
# @!attribute [rw] created_by
#   @return [Hash]
#
# @!attribute [rw] created_by_display_name
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] default_payment_rail
#   @return [String, nil]
#
# @!attribute [rw] display_name
#   @return [String, nil]
#
# @!attribute [rw] expiry_date
#   @return [String, nil]
#
# @!attribute [rw] external_account_icon
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_connected_account
#   @return [Boolean, nil]
#
# @!attribute [rw] is_default
#   @return [Boolean, nil]
#
# @!attribute [rw] is_trust_account
#   @return [Boolean, nil]
#
# @!attribute [rw] is_virtual
#   @return [Boolean, nil]
#
# @!attribute [rw] last_transaction
#   @return [Hash, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_name
#   @return [String, nil]
#
# @!attribute [rw] physical_account_id
#   @return [String, nil]
#
# @!attribute [rw] rule
#   @return [Array, nil]
#
# @!attribute [rw] submitted_payouts_balance
#   @return [Float, nil]
#
# @!attribute [rw] submitted_payouts_balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplier_sepa_instant_status
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_connection_status
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_last_synced_at
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_sync_last_failed_at
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_sync_last_failure_reason
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_sync_status
#   @return [String, nil]
#
# @!attribute [rw] xero_unsynchronised_transactions_count
#   @return [Integer, nil]
PaymentAccount = Struct.new(
  :account_name,
  :account_supplier_name,
  :available_balance,
  :available_balance_minor_unit,
  :balance,
  :balance_minor_unit,
  :bank_name,
  :consent_id,
  :created_by,
  :created_by_display_name,
  :currency,
  :default_payment_rail,
  :display_name,
  :expiry_date,
  :external_account_icon,
  :id,
  :identifier,
  :inserted,
  :is_archived,
  :is_connected_account,
  :is_default,
  :is_trust_account,
  :is_virtual,
  :last_transaction,
  :last_updated,
  :merchant_id,
  :merchant_name,
  :physical_account_id,
  :rule,
  :submitted_payouts_balance,
  :submitted_payouts_balance_minor_unit,
  :summary,
  :supplier_sepa_instant_status,
  :xero_bank_feed_connection_status,
  :xero_bank_feed_last_synced_at,
  :xero_bank_feed_sync_last_failed_at,
  :xero_bank_feed_sync_last_failure_reason,
  :xero_bank_feed_sync_status,
  :xero_unsynchronised_transactions_count,
  keyword_init: true
)

# Request payload for PaymentAccount#list.
#
# @!attribute [rw] account_id
#   @return [String, nil]
PaymentAccountListMatch = Struct.new(
  :account_id,
  keyword_init: true
)

# PaymentAccountMinimal entity data model.
#
# @!attribute [rw] account_name
#   @return [String, nil]
#
# @!attribute [rw] available_balance
#   @return [Float, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_connected_account
#   @return [Boolean, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] submitted_payouts_balance
#   @return [Float, nil]
PaymentAccountMinimal = Struct.new(
  :account_name,
  :available_balance,
  :balance,
  :balance_minor_unit,
  :currency,
  :id,
  :identifier,
  :is_archived,
  :is_connected_account,
  :merchant_id,
  :submitted_payouts_balance,
  keyword_init: true
)

# Request payload for PaymentAccountMinimal#list.
#
# @!attribute [rw] account_name
#   @return [String, nil]
#
# @!attribute [rw] available_balance
#   @return [Float, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_connected_account
#   @return [Boolean, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] submitted_payouts_balance
#   @return [Float, nil]
PaymentAccountMinimalListMatch = Struct.new(
  :account_name,
  :available_balance,
  :balance,
  :balance_minor_unit,
  :currency,
  :id,
  :identifier,
  :is_archived,
  :is_connected_account,
  :merchant_id,
  :submitted_payouts_balance,
  keyword_init: true
)

# PaymentInitiation entity data model.
#
# @!attribute [rw] payment_initiation_id
#   @return [String, nil]
#
# @!attribute [rw] payment_request_callback_url
#   @return [String, nil]
#
# @!attribute [rw] payment_request_id
#   @return [String, nil]
#
# @!attribute [rw] redirect_url
#   @return [String, nil]
#
# @!attribute [rw] response_type
#   @return [String, nil]
#
# @!attribute [rw] specific_error_message
#   @return [String, nil]
PaymentInitiation = Struct.new(
  :payment_initiation_id,
  :payment_request_callback_url,
  :payment_request_id,
  :redirect_url,
  :response_type,
  :specific_error_message,
  keyword_init: true
)

# Request payload for PaymentInitiation#create.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
PaymentInitiationCreateData = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# PaymentRequest entity data model.
#
# @!attribute [rw] address
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amount_pending
#   @return [Float, nil]
#
# @!attribute [rw] amount_received
#   @return [Float, nil]
#
# @!attribute [rw] amount_refunded
#   @return [Float, nil]
#
# @!attribute [rw] auto_send_receipt
#   @return [Boolean, nil]
#
# @!attribute [rw] base_origin_url
#   @return [String, nil]
#
# @!attribute [rw] callback_url
#   @return [String, nil]
#
# @!attribute [rw] card_authorize_only
#   @return [Boolean, nil]
#
# @!attribute [rw] card_create_token
#   @return [Boolean, nil]
#
# @!attribute [rw] card_create_token_mode
#   @return [String, nil]
#
# @!attribute [rw] card_ignore_cvn
#   @return [Boolean, nil]
#
# @!attribute [rw] card_processor_merchant_id
#   @return [String, nil]
#
# @!attribute [rw] card_stripe_payment_intent_id
#   @return [String, nil]
#
# @!attribute [rw] card_stripe_payment_intent_secret
#   @return [String, nil]
#
# @!attribute [rw] created_by_user
#   @return [Hash]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] custom_field
#   @return [Array, nil]
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] customer_id
#   @return [String, nil]
#
# @!attribute [rw] customer_name
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination_account
#   @return [Hash, nil]
#
# @!attribute [rw] direct_debit_payment
#   @return [Hash, nil]
#
# @!attribute [rw] do_simulate_settlement_failure
#   @return [Boolean, nil]
#
# @!attribute [rw] due_date
#   @return [String, nil]
#
# @!attribute [rw] error_description
#   @return [String, nil]
#
# @!attribute [rw] event
#   @return [Array, nil]
#
# @!attribute [rw] failed_payment_request
#   @return [Hash, nil]
#
# @!attribute [rw] failure_callback_url
#   @return [String, nil]
#
# @!attribute [rw] field_display_setting
#   @return [Array, nil]
#
# @!attribute [rw] formatted_amount
#   @return [String, nil]
#
# @!attribute [rw] hosted_pay_checkout_url
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignore_address_verification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] inserted_sortable
#   @return [String, nil]
#
# @!attribute [rw] institution
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] lightning_invoice
#   @return [String, nil]
#
# @!attribute [rw] lightning_invoice_expires_at
#   @return [String, nil]
#
# @!attribute [rw] merchant_direct_debit_mandate_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_token_description
#   @return [String, nil]
#
# @!attribute [rw] notification_email_address
#   @return [String, nil]
#
# @!attribute [rw] notification_role_i_d
#   @return [Array, nil]
#
# @!attribute [rw] order_id
#   @return [String, nil]
#
# @!attribute [rw] partial_payment_method
#   @return [String, nil]
#
# @!attribute [rw] partial_payment_step
#   @return [String, nil]
#
# @!attribute [rw] payment_attempt
#   @return [Array, nil]
#
# @!attribute [rw] payment_initiation_id
#   @return [String, nil]
#
# @!attribute [rw] payment_method
#   @return [Array, nil]
#
# @!attribute [rw] payment_processor
#   @return [String, nil]
#
# @!attribute [rw] payment_request
#   @return [Array, nil]
#
# @!attribute [rw] payrun_id
#   @return [String, nil]
#
# @!attribute [rw] pisp_account_id
#   @return [String, nil]
#
# @!attribute [rw] priority_bank_id
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandbox_settle_delay_in_second
#   @return [Integer, nil]
#
# @!attribute [rw] shipping_address
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] success_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenised_card
#   @return [Array, nil]
#
# @!attribute [rw] transaction
#   @return [Array, nil]
#
# @!attribute [rw] use_hosted_payment_page
#   @return [Boolean, nil]
PaymentRequest = Struct.new(
  :address,
  :amount,
  :amount_pending,
  :amount_received,
  :amount_refunded,
  :auto_send_receipt,
  :base_origin_url,
  :callback_url,
  :card_authorize_only,
  :card_create_token,
  :card_create_token_mode,
  :card_ignore_cvn,
  :card_processor_merchant_id,
  :card_stripe_payment_intent_id,
  :card_stripe_payment_intent_secret,
  :created_by_user,
  :currency,
  :custom_field,
  :customer_email_address,
  :customer_id,
  :customer_name,
  :description,
  :destination_account,
  :direct_debit_payment,
  :do_simulate_settlement_failure,
  :due_date,
  :error_description,
  :event,
  :failed_payment_request,
  :failure_callback_url,
  :field_display_setting,
  :formatted_amount,
  :hosted_pay_checkout_url,
  :id,
  :ignore_address_verification,
  :inserted,
  :inserted_sortable,
  :institution,
  :is_archived,
  :jwk,
  :last_updated,
  :lightning_invoice,
  :lightning_invoice_expires_at,
  :merchant_direct_debit_mandate_id,
  :merchant_id,
  :merchant_token_description,
  :notification_email_address,
  :notification_role_i_d,
  :order_id,
  :partial_payment_method,
  :partial_payment_step,
  :payment_attempt,
  :payment_initiation_id,
  :payment_method,
  :payment_processor,
  :payment_request,
  :payrun_id,
  :pisp_account_id,
  :priority_bank_id,
  :result,
  :sandbox_settle_delay_in_second,
  :shipping_address,
  :status,
  :success_web_hook_url,
  :tag,
  :title,
  :tokenised_card,
  :transaction,
  :use_hosted_payment_page,
  keyword_init: true
)

# Request payload for PaymentRequest#load.
#
# @!attribute [rw] paymentrequest_id
#   @return [String, nil]
PaymentRequestLoadMatch = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# Request payload for PaymentRequest#list.
#
# @!attribute [rw] address
#   @return [Array, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amount_pending
#   @return [Float, nil]
#
# @!attribute [rw] amount_received
#   @return [Float, nil]
#
# @!attribute [rw] amount_refunded
#   @return [Float, nil]
#
# @!attribute [rw] auto_send_receipt
#   @return [Boolean, nil]
#
# @!attribute [rw] base_origin_url
#   @return [String, nil]
#
# @!attribute [rw] callback_url
#   @return [String, nil]
#
# @!attribute [rw] card_authorize_only
#   @return [Boolean, nil]
#
# @!attribute [rw] card_create_token
#   @return [Boolean, nil]
#
# @!attribute [rw] card_create_token_mode
#   @return [String, nil]
#
# @!attribute [rw] card_ignore_cvn
#   @return [Boolean, nil]
#
# @!attribute [rw] card_processor_merchant_id
#   @return [String, nil]
#
# @!attribute [rw] card_stripe_payment_intent_id
#   @return [String, nil]
#
# @!attribute [rw] card_stripe_payment_intent_secret
#   @return [String, nil]
#
# @!attribute [rw] created_by_user
#   @return [Hash, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] custom_field
#   @return [Array, nil]
#
# @!attribute [rw] customer_email_address
#   @return [String, nil]
#
# @!attribute [rw] customer_id
#   @return [String, nil]
#
# @!attribute [rw] customer_name
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination_account
#   @return [Hash, nil]
#
# @!attribute [rw] direct_debit_payment
#   @return [Hash, nil]
#
# @!attribute [rw] do_simulate_settlement_failure
#   @return [Boolean, nil]
#
# @!attribute [rw] due_date
#   @return [String, nil]
#
# @!attribute [rw] error_description
#   @return [String, nil]
#
# @!attribute [rw] event
#   @return [Array, nil]
#
# @!attribute [rw] failed_payment_request
#   @return [Hash, nil]
#
# @!attribute [rw] failure_callback_url
#   @return [String, nil]
#
# @!attribute [rw] field_display_setting
#   @return [Array, nil]
#
# @!attribute [rw] formatted_amount
#   @return [String, nil]
#
# @!attribute [rw] hosted_pay_checkout_url
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ignore_address_verification
#   @return [Boolean, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] inserted_sortable
#   @return [String, nil]
#
# @!attribute [rw] institution
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] lightning_invoice
#   @return [String, nil]
#
# @!attribute [rw] lightning_invoice_expires_at
#   @return [String, nil]
#
# @!attribute [rw] merchant_direct_debit_mandate_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_token_description
#   @return [String, nil]
#
# @!attribute [rw] notification_email_address
#   @return [String, nil]
#
# @!attribute [rw] notification_role_i_d
#   @return [Array, nil]
#
# @!attribute [rw] order_id
#   @return [String, nil]
#
# @!attribute [rw] partial_payment_method
#   @return [String, nil]
#
# @!attribute [rw] partial_payment_step
#   @return [String, nil]
#
# @!attribute [rw] payment_attempt
#   @return [Array, nil]
#
# @!attribute [rw] payment_initiation_id
#   @return [String, nil]
#
# @!attribute [rw] payment_method
#   @return [Array, nil]
#
# @!attribute [rw] payment_processor
#   @return [String, nil]
#
# @!attribute [rw] payment_request
#   @return [Array, nil]
#
# @!attribute [rw] payrun_id
#   @return [String, nil]
#
# @!attribute [rw] pisp_account_id
#   @return [String, nil]
#
# @!attribute [rw] priority_bank_id
#   @return [String, nil]
#
# @!attribute [rw] result
#   @return [Hash, nil]
#
# @!attribute [rw] sandbox_settle_delay_in_second
#   @return [Integer, nil]
#
# @!attribute [rw] shipping_address
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] success_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [Array, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] tokenised_card
#   @return [Array, nil]
#
# @!attribute [rw] transaction
#   @return [Array, nil]
#
# @!attribute [rw] use_hosted_payment_page
#   @return [Boolean, nil]
PaymentRequestListMatch = Struct.new(
  :address,
  :amount,
  :amount_pending,
  :amount_received,
  :amount_refunded,
  :auto_send_receipt,
  :base_origin_url,
  :callback_url,
  :card_authorize_only,
  :card_create_token,
  :card_create_token_mode,
  :card_ignore_cvn,
  :card_processor_merchant_id,
  :card_stripe_payment_intent_id,
  :card_stripe_payment_intent_secret,
  :created_by_user,
  :currency,
  :custom_field,
  :customer_email_address,
  :customer_id,
  :customer_name,
  :description,
  :destination_account,
  :direct_debit_payment,
  :do_simulate_settlement_failure,
  :due_date,
  :error_description,
  :event,
  :failed_payment_request,
  :failure_callback_url,
  :field_display_setting,
  :formatted_amount,
  :hosted_pay_checkout_url,
  :id,
  :ignore_address_verification,
  :inserted,
  :inserted_sortable,
  :institution,
  :is_archived,
  :jwk,
  :last_updated,
  :lightning_invoice,
  :lightning_invoice_expires_at,
  :merchant_direct_debit_mandate_id,
  :merchant_id,
  :merchant_token_description,
  :notification_email_address,
  :notification_role_i_d,
  :order_id,
  :partial_payment_method,
  :partial_payment_step,
  :payment_attempt,
  :payment_initiation_id,
  :payment_method,
  :payment_processor,
  :payment_request,
  :payrun_id,
  :pisp_account_id,
  :priority_bank_id,
  :result,
  :sandbox_settle_delay_in_second,
  :shipping_address,
  :status,
  :success_web_hook_url,
  :tag,
  :title,
  :tokenised_card,
  :transaction,
  :use_hosted_payment_page,
  keyword_init: true
)

# Request payload for PaymentRequest#create.
#
# @!attribute [rw] paymentrequest_id
#   @return [String, nil]
PaymentRequestCreateData = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# Request payload for PaymentRequest#update.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
PaymentRequestUpdateData = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# Request payload for PaymentRequest#remove.
#
# @!attribute [rw] id
#   @return [String]
PaymentRequestRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# PaymentRequestEvent entity data model.
#
# @!attribute [rw] amount
#   @return [Float]
#
# @!attribute [rw] apple_pay_transaction_id
#   @return [String, nil]
#
# @!attribute [rw] card_authorization_response_id
#   @return [String, nil]
#
# @!attribute [rw] card_expiry_month
#   @return [Integer, nil]
#
# @!attribute [rw] card_expiry_year
#   @return [Integer, nil]
#
# @!attribute [rw] card_issuer
#   @return [String, nil]
#
# @!attribute [rw] card_issuer_country
#   @return [String, nil]
#
# @!attribute [rw] card_last_four_digit
#   @return [String, nil]
#
# @!attribute [rw] card_request_id
#   @return [String, nil]
#
# @!attribute [rw] card_scheme
#   @return [String, nil]
#
# @!attribute [rw] card_token_customer_id
#   @return [String, nil]
#
# @!attribute [rw] card_transaction_id
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] direct_debit_payment_id
#   @return [String, nil]
#
# @!attribute [rw] direct_debit_payment_reference
#   @return [String, nil]
#
# @!attribute [rw] drirect_debit_mandate_id
#   @return [String, nil]
#
# @!attribute [rw] error_message
#   @return [String, nil]
#
# @!attribute [rw] error_reason
#   @return [String, nil]
#
# @!attribute [rw] event_type
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] lightning_invoice
#   @return [String, nil]
#
# @!attribute [rw] lightning_r_hash
#   @return [String, nil]
#
# @!attribute [rw] origin_url
#   @return [String, nil]
#
# @!attribute [rw] payment_method_type
#   @return [String, nil]
#
# @!attribute [rw] payment_processor_name
#   @return [String, nil]
#
# @!attribute [rw] payment_request_id
#   @return [String, nil]
#
# @!attribute [rw] pisp_bank_status
#   @return [String, nil]
#
# @!attribute [rw] pisp_payment_initiation_id
#   @return [String, nil]
#
# @!attribute [rw] pisp_payment_institution_name
#   @return [String, nil]
#
# @!attribute [rw] pisp_payment_service_provider_id
#   @return [String, nil]
#
# @!attribute [rw] pisp_redirect_url
#   @return [String, nil]
#
# @!attribute [rw] reconciled_transaction_id
#   @return [String, nil]
#
# @!attribute [rw] refund_payout_id
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] wallet_name
#   @return [String, nil]
PaymentRequestEvent = Struct.new(
  :amount,
  :apple_pay_transaction_id,
  :card_authorization_response_id,
  :card_expiry_month,
  :card_expiry_year,
  :card_issuer,
  :card_issuer_country,
  :card_last_four_digit,
  :card_request_id,
  :card_scheme,
  :card_token_customer_id,
  :card_transaction_id,
  :currency,
  :direct_debit_payment_id,
  :direct_debit_payment_reference,
  :drirect_debit_mandate_id,
  :error_message,
  :error_reason,
  :event_type,
  :id,
  :inserted,
  :lightning_invoice,
  :lightning_r_hash,
  :origin_url,
  :payment_method_type,
  :payment_processor_name,
  :payment_request_id,
  :pisp_bank_status,
  :pisp_payment_initiation_id,
  :pisp_payment_institution_name,
  :pisp_payment_service_provider_id,
  :pisp_redirect_url,
  :reconciled_transaction_id,
  :refund_payout_id,
  :status,
  :wallet_name,
  keyword_init: true
)

# Request payload for PaymentRequestEvent#list.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
PaymentRequestEventListMatch = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# PaymentRequestMetric entity data model.
#
# @!attribute [rw] all
#   @return [Integer, nil]
#
# @!attribute [rw] authorized
#   @return [Integer, nil]
#
# @!attribute [rw] paid
#   @return [Integer, nil]
#
# @!attribute [rw] partially_paid
#   @return [Integer, nil]
#
# @!attribute [rw] total_amounts_by_currency
#   @return [Hash, nil]
#
# @!attribute [rw] unpaid
#   @return [Integer, nil]
PaymentRequestMetric = Struct.new(
  :all,
  :authorized,
  :paid,
  :partially_paid,
  :total_amounts_by_currency,
  :unpaid,
  keyword_init: true
)

# Request payload for PaymentRequestMetric#load.
#
# @!attribute [rw] all
#   @return [Integer, nil]
#
# @!attribute [rw] authorized
#   @return [Integer, nil]
#
# @!attribute [rw] paid
#   @return [Integer, nil]
#
# @!attribute [rw] partially_paid
#   @return [Integer, nil]
#
# @!attribute [rw] total_amounts_by_currency
#   @return [Hash, nil]
#
# @!attribute [rw] unpaid
#   @return [Integer, nil]
PaymentRequestMetricLoadMatch = Struct.new(
  :all,
  :authorized,
  :paid,
  :partially_paid,
  :total_amounts_by_currency,
  :unpaid,
  keyword_init: true
)

# PaymentRequestMinimal entity data model.
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amount_pending
#   @return [Float, nil]
#
# @!attribute [rw] amount_received
#   @return [Float, nil]
#
# @!attribute [rw] amount_refunded
#   @return [Float, nil]
#
# @!attribute [rw] callback_url
#   @return [String, nil]
#
# @!attribute [rw] card_stripe_payment_intent_secret
#   @return [String, nil]
#
# @!attribute [rw] country_code
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] custom_fields_to_display
#   @return [Array, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] due_date
#   @return [String, nil]
#
# @!attribute [rw] field_display_setting
#   @return [Array, nil]
#
# @!attribute [rw] google_pay_merchant_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] jwk
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_logo_url_png
#   @return [String, nil]
#
# @!attribute [rw] merchant_logo_url_svg
#   @return [String, nil]
#
# @!attribute [rw] merchant_name
#   @return [String, nil]
#
# @!attribute [rw] merchant_short_name
#   @return [String, nil]
#
# @!attribute [rw] partial_payment_method
#   @return [String, nil]
#
# @!attribute [rw] payment_attempt
#   @return [Array, nil]
#
# @!attribute [rw] payment_methods_list
#   @return [Array, nil]
#
# @!attribute [rw] payment_processor
#   @return [String, nil]
#
# @!attribute [rw] payment_processor_key
#   @return [String, nil]
#
# @!attribute [rw] pisp_error
#   @return [String, nil]
#
# @!attribute [rw] priority_bank_id
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] stripe_account_id
#   @return [String, nil]
#
# @!attribute [rw] title
#   @return [String, nil]
PaymentRequestMinimal = Struct.new(
  :amount,
  :amount_pending,
  :amount_received,
  :amount_refunded,
  :callback_url,
  :card_stripe_payment_intent_secret,
  :country_code,
  :currency,
  :custom_fields_to_display,
  :description,
  :due_date,
  :field_display_setting,
  :google_pay_merchant_id,
  :id,
  :jwk,
  :merchant_id,
  :merchant_logo_url_png,
  :merchant_logo_url_svg,
  :merchant_name,
  :merchant_short_name,
  :partial_payment_method,
  :payment_attempt,
  :payment_methods_list,
  :payment_processor,
  :payment_processor_key,
  :pisp_error,
  :priority_bank_id,
  :status,
  :stripe_account_id,
  :title,
  keyword_init: true
)

# Request payload for PaymentRequestMinimal#list.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
PaymentRequestMinimalListMatch = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# PaymentRequestResult entity data model.
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amount_pending
#   @return [Float, nil]
#
# @!attribute [rw] amount_received
#   @return [Float, nil]
#
# @!attribute [rw] amount_refunded
#   @return [Float, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] customer_id
#   @return [String, nil]
#
# @!attribute [rw] payment
#   @return [Array, nil]
#
# @!attribute [rw] payment_request_id
#   @return [String, nil]
#
# @!attribute [rw] pisp_authorization
#   @return [Array, nil]
#
# @!attribute [rw] requested_amount
#   @return [Float, nil]
#
# @!attribute [rw] result
#   @return [String, nil]
PaymentRequestResult = Struct.new(
  :amount,
  :amount_pending,
  :amount_received,
  :amount_refunded,
  :currency,
  :customer_id,
  :payment,
  :payment_request_id,
  :pisp_authorization,
  :requested_amount,
  :result,
  keyword_init: true
)

# Request payload for PaymentRequestResult#list.
#
# @!attribute [rw] paymentrequest_id
#   @return [String]
PaymentRequestResultListMatch = Struct.new(
  :paymentrequest_id,
  keyword_init: true
)

# Payout entity data model.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] allow_incomplete
#   @return [Boolean, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amount_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] approve_payout_url
#   @return [String, nil]
#
# @!attribute [rw] approver_id
#   @return [String, nil]
#
# @!attribute [rw] authentication_method
#   @return [Array, nil]
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] batch_payout_id
#   @return [String, nil]
#
# @!attribute [rw] beneficiary
#   @return [Hash]
#
# @!attribute [rw] beneficiary_id
#   @return [String, nil]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] can_process
#   @return [Boolean, nil]
#
# @!attribute [rw] can_update
#   @return [Boolean, nil]
#
# @!attribute [rw] charge_bearer
#   @return [String, nil]
#
# @!attribute [rw] created_by
#   @return [String, nil]
#
# @!attribute [rw] created_by_email_address
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] current_user_id
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] document
#   @return [Array, nil]
#
# @!attribute [rw] event
#   @return [Array, nil]
#
# @!attribute [rw] failed_payout
#   @return [Hash, nil]
#
# @!attribute [rw] formatted_amount
#   @return [String, nil]
#
# @!attribute [rw] formatted_fx_destination_amount
#   @return [String, nil]
#
# @!attribute [rw] formatted_schedule
#   @return [String, nil]
#
# @!attribute [rw] formatted_schedule_day_only
#   @return [String, nil]
#
# @!attribute [rw] formatted_source_account_available_balance
#   @return [String, nil]
#
# @!attribute [rw] fx_destination_amount
#   @return [Float, nil]
#
# @!attribute [rw] fx_destination_amount_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] fx_destination_currency
#   @return [String, nil]
#
# @!attribute [rw] fx_quote_expires_at
#   @return [String, nil]
#
# @!attribute [rw] fx_quote_id
#   @return [String, nil]
#
# @!attribute [rw] fx_rate
#   @return [Float, nil]
#
# @!attribute [rw] fx_use_destination_amount
#   @return [Boolean, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoice_id
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_failed
#   @return [Boolean, nil]
#
# @!attribute [rw] is_settled
#   @return [Boolean, nil]
#
# @!attribute [rw] is_submitted
#   @return [Boolean, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_token_description
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] payment_processor
#   @return [String, nil]
#
# @!attribute [rw] payment_rail
#   @return [String, nil]
#
# @!attribute [rw] payout
#   @return [Array, nil]
#
# @!attribute [rw] payrun_id
#   @return [String, nil]
#
# @!attribute [rw] payrun_name
#   @return [String, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] rule
#   @return [Hash, nil]
#
# @!attribute [rw] schedule_date
#   @return [String, nil]
#
# @!attribute [rw] scheduled
#   @return [Boolean, nil]
#
# @!attribute [rw] source_account_available_balance
#   @return [Float, nil]
#
# @!attribute [rw] source_account_available_balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] source_account_bic
#   @return [String, nil]
#
# @!attribute [rw] source_account_currency
#   @return [String, nil]
#
# @!attribute [rw] source_account_iban
#   @return [String, nil]
#
# @!attribute [rw] source_account_identifier
#   @return [Hash]
#
# @!attribute [rw] source_account_name
#   @return [String, nil]
#
# @!attribute [rw] source_account_number
#   @return [String, nil]
#
# @!attribute [rw] source_account_sortcode
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [Array, nil]
#
# @!attribute [rw] tag_id
#   @return [Array, nil]
#
# @!attribute [rw] their_reference
#   @return [String, nil]
#
# @!attribute [rw] topup_payrun_id
#   @return [String, nil]
#
# @!attribute [rw] transacted_amount
#   @return [Float, nil]
#
# @!attribute [rw] transacted_fx_amount
#   @return [Float, nil]
#
# @!attribute [rw] transacted_fx_rate
#   @return [Float, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] user_id
#   @return [String, nil]
#
# @!attribute [rw] your_reference
#   @return [String, nil]
Payout = Struct.new(
  :account_id,
  :allow_incomplete,
  :amount,
  :amount_minor_unit,
  :approve_payout_url,
  :approver_id,
  :authentication_method,
  :authorisation,
  :authorisers_completed_count,
  :authorisers_required_count,
  :batch_payout_id,
  :beneficiary,
  :beneficiary_id,
  :can_authorise,
  :can_process,
  :can_update,
  :charge_bearer,
  :created_by,
  :created_by_email_address,
  :currency,
  :current_user_id,
  :description,
  :destination,
  :document,
  :event,
  :failed_payout,
  :formatted_amount,
  :formatted_fx_destination_amount,
  :formatted_schedule,
  :formatted_schedule_day_only,
  :formatted_source_account_available_balance,
  :fx_destination_amount,
  :fx_destination_amount_minor_unit,
  :fx_destination_currency,
  :fx_quote_expires_at,
  :fx_quote_id,
  :fx_rate,
  :fx_use_destination_amount,
  :has_current_user_authorised,
  :id,
  :inserted,
  :invoice_id,
  :is_archived,
  :is_failed,
  :is_settled,
  :is_submitted,
  :last_updated,
  :merchant_id,
  :merchant_token_description,
  :nonce,
  :payment_processor,
  :payment_rail,
  :payout,
  :payrun_id,
  :payrun_name,
  :reason,
  :rule,
  :schedule_date,
  :scheduled,
  :source_account_available_balance,
  :source_account_available_balance_minor_unit,
  :source_account_bic,
  :source_account_currency,
  :source_account_iban,
  :source_account_identifier,
  :source_account_name,
  :source_account_number,
  :source_account_sortcode,
  :status,
  :tag,
  :tag_id,
  :their_reference,
  :topup_payrun_id,
  :transacted_amount,
  :transacted_fx_amount,
  :transacted_fx_rate,
  :type,
  :user_id,
  :your_reference,
  keyword_init: true
)

# Request payload for Payout#load.
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] destination
#   @return [String, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
PayoutLoadMatch = Struct.new(
  :amount,
  :destination,
  :source,
  :id,
  keyword_init: true
)

# Request payload for Payout#list.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
PayoutListMatch = Struct.new(
  :account_id,
  :merchant_id,
  keyword_init: true
)

# Request payload for Payout#create.
#
# @!attribute [rw] id
#   @return [String, nil]
PayoutCreateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Payout#update.
#
# @!attribute [rw] id
#   @return [String]
PayoutUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Payout#remove.
#
# @!attribute [rw] id
#   @return [String]
PayoutRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# PayoutKeysetPage entity data model.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amount_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] approve_payout_url
#   @return [String, nil]
#
# @!attribute [rw] approver_id
#   @return [String, nil]
#
# @!attribute [rw] authentication_method
#   @return [Array, nil]
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] batch_payout_id
#   @return [String, nil]
#
# @!attribute [rw] beneficiary
#   @return [Hash]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] can_process
#   @return [Boolean, nil]
#
# @!attribute [rw] can_update
#   @return [Boolean, nil]
#
# @!attribute [rw] charge_bearer
#   @return [String, nil]
#
# @!attribute [rw] created_by
#   @return [String, nil]
#
# @!attribute [rw] created_by_email_address
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] current_user_id
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] destination
#   @return [Hash, nil]
#
# @!attribute [rw] document
#   @return [Array, nil]
#
# @!attribute [rw] event
#   @return [Array, nil]
#
# @!attribute [rw] formatted_amount
#   @return [String, nil]
#
# @!attribute [rw] formatted_fx_destination_amount
#   @return [String, nil]
#
# @!attribute [rw] formatted_schedule
#   @return [String, nil]
#
# @!attribute [rw] formatted_schedule_day_only
#   @return [String, nil]
#
# @!attribute [rw] formatted_source_account_available_balance
#   @return [String, nil]
#
# @!attribute [rw] fx_destination_amount
#   @return [Float, nil]
#
# @!attribute [rw] fx_destination_amount_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] fx_destination_currency
#   @return [String, nil]
#
# @!attribute [rw] fx_quote_expires_at
#   @return [String, nil]
#
# @!attribute [rw] fx_quote_id
#   @return [String, nil]
#
# @!attribute [rw] fx_rate
#   @return [Float, nil]
#
# @!attribute [rw] fx_use_destination_amount
#   @return [Boolean, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoice_id
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_failed
#   @return [Boolean, nil]
#
# @!attribute [rw] is_settled
#   @return [Boolean, nil]
#
# @!attribute [rw] is_submitted
#   @return [Boolean, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_token_description
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] payment_processor
#   @return [String, nil]
#
# @!attribute [rw] payment_rail
#   @return [String, nil]
#
# @!attribute [rw] payrun_id
#   @return [String, nil]
#
# @!attribute [rw] payrun_name
#   @return [String, nil]
#
# @!attribute [rw] rule
#   @return [Hash, nil]
#
# @!attribute [rw] schedule_date
#   @return [String, nil]
#
# @!attribute [rw] scheduled
#   @return [Boolean, nil]
#
# @!attribute [rw] source_account_available_balance
#   @return [Float, nil]
#
# @!attribute [rw] source_account_available_balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] source_account_bic
#   @return [String, nil]
#
# @!attribute [rw] source_account_currency
#   @return [String, nil]
#
# @!attribute [rw] source_account_iban
#   @return [String, nil]
#
# @!attribute [rw] source_account_identifier
#   @return [Hash]
#
# @!attribute [rw] source_account_name
#   @return [String, nil]
#
# @!attribute [rw] source_account_number
#   @return [String, nil]
#
# @!attribute [rw] source_account_sortcode
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] tag
#   @return [Array, nil]
#
# @!attribute [rw] their_reference
#   @return [String, nil]
#
# @!attribute [rw] topup_payrun_id
#   @return [String, nil]
#
# @!attribute [rw] transacted_amount
#   @return [Float, nil]
#
# @!attribute [rw] transacted_fx_amount
#   @return [Float, nil]
#
# @!attribute [rw] transacted_fx_rate
#   @return [Float, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] user_id
#   @return [String, nil]
#
# @!attribute [rw] your_reference
#   @return [String, nil]
PayoutKeysetPage = Struct.new(
  :account_id,
  :amount,
  :amount_minor_unit,
  :approve_payout_url,
  :approver_id,
  :authentication_method,
  :authorisation,
  :authorisers_completed_count,
  :authorisers_required_count,
  :batch_payout_id,
  :beneficiary,
  :can_authorise,
  :can_process,
  :can_update,
  :charge_bearer,
  :created_by,
  :created_by_email_address,
  :currency,
  :current_user_id,
  :description,
  :destination,
  :document,
  :event,
  :formatted_amount,
  :formatted_fx_destination_amount,
  :formatted_schedule,
  :formatted_schedule_day_only,
  :formatted_source_account_available_balance,
  :fx_destination_amount,
  :fx_destination_amount_minor_unit,
  :fx_destination_currency,
  :fx_quote_expires_at,
  :fx_quote_id,
  :fx_rate,
  :fx_use_destination_amount,
  :has_current_user_authorised,
  :id,
  :inserted,
  :invoice_id,
  :is_archived,
  :is_failed,
  :is_settled,
  :is_submitted,
  :last_updated,
  :merchant_id,
  :merchant_token_description,
  :nonce,
  :payment_processor,
  :payment_rail,
  :payrun_id,
  :payrun_name,
  :rule,
  :schedule_date,
  :scheduled,
  :source_account_available_balance,
  :source_account_available_balance_minor_unit,
  :source_account_bic,
  :source_account_currency,
  :source_account_iban,
  :source_account_identifier,
  :source_account_name,
  :source_account_number,
  :source_account_sortcode,
  :status,
  :tag,
  :their_reference,
  :topup_payrun_id,
  :transacted_amount,
  :transacted_fx_amount,
  :transacted_fx_rate,
  :type,
  :user_id,
  :your_reference,
  keyword_init: true
)

# Request payload for PayoutKeysetPage#list.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
PayoutKeysetPageListMatch = Struct.new(
  :account_id,
  :merchant_id,
  keyword_init: true
)

# PayoutMetric entity data model.
#
# @!attribute [rw] all
#   @return [Float, nil]
#
# @!attribute [rw] failed
#   @return [Float, nil]
#
# @!attribute [rw] in_progress
#   @return [Float, nil]
#
# @!attribute [rw] paid
#   @return [Float, nil]
#
# @!attribute [rw] pending_approval
#   @return [Float, nil]
#
# @!attribute [rw] scheduled
#   @return [Float, nil]
#
# @!attribute [rw] total_amounts_by_currency
#   @return [Hash, nil]
PayoutMetric = Struct.new(
  :all,
  :failed,
  :in_progress,
  :paid,
  :pending_approval,
  :scheduled,
  :total_amounts_by_currency,
  keyword_init: true
)

# Request payload for PayoutMetric#load.
#
# @!attribute [rw] all
#   @return [Float, nil]
#
# @!attribute [rw] failed
#   @return [Float, nil]
#
# @!attribute [rw] in_progress
#   @return [Float, nil]
#
# @!attribute [rw] paid
#   @return [Float, nil]
#
# @!attribute [rw] pending_approval
#   @return [Float, nil]
#
# @!attribute [rw] scheduled
#   @return [Float, nil]
#
# @!attribute [rw] total_amounts_by_currency
#   @return [Hash, nil]
PayoutMetricLoadMatch = Struct.new(
  :all,
  :failed,
  :in_progress,
  :paid,
  :pending_approval,
  :scheduled,
  :total_amounts_by_currency,
  keyword_init: true
)

# Payrun entity data model.
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisation_date
#   @return [String, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] batch_payout_id
#   @return [String, nil]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] can_delete
#   @return [Boolean, nil]
#
# @!attribute [rw] can_edit
#   @return [Boolean, nil]
#
# @!attribute [rw] event
#   @return [Array, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoice
#   @return [Array, nil]
#
# @!attribute [rw] invoices_minimal
#   @return [Array, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] last_updated_by
#   @return [Hash]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] note
#   @return [String, nil]
#
# @!attribute [rw] payment
#   @return [Array, nil]
#
# @!attribute [rw] payout
#   @return [Array, nil]
#
# @!attribute [rw] payouts_count
#   @return [Integer, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] schedule_date
#   @return [String, nil]
#
# @!attribute [rw] scheduled_date
#   @return [String, nil]
#
# @!attribute [rw] source_account
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] total_eur
#   @return [Float, nil]
#
# @!attribute [rw] total_gbp
#   @return [Float, nil]
#
# @!attribute [rw] total_usd
#   @return [Float, nil]
Payrun = Struct.new(
  :authorisation,
  :authorisation_date,
  :authorisers_completed_count,
  :authorisers_required_count,
  :batch_payout_id,
  :can_authorise,
  :can_delete,
  :can_edit,
  :event,
  :has_current_user_authorised,
  :id,
  :inserted,
  :invoice,
  :invoices_minimal,
  :is_archived,
  :last_updated,
  :last_updated_by,
  :merchant_id,
  :name,
  :nonce,
  :note,
  :payment,
  :payout,
  :payouts_count,
  :reason,
  :schedule_date,
  :scheduled_date,
  :source_account,
  :status,
  :total_eur,
  :total_gbp,
  :total_usd,
  keyword_init: true
)

# Request payload for Payrun#load.
#
# @!attribute [rw] id
#   @return [String]
PayrunLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Payrun#list.
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisation_date
#   @return [String, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] batch_payout_id
#   @return [String, nil]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] can_delete
#   @return [Boolean, nil]
#
# @!attribute [rw] can_edit
#   @return [Boolean, nil]
#
# @!attribute [rw] event
#   @return [Array, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] invoice
#   @return [Array, nil]
#
# @!attribute [rw] invoices_minimal
#   @return [Array, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] last_updated_by
#   @return [Hash, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] note
#   @return [String, nil]
#
# @!attribute [rw] payment
#   @return [Array, nil]
#
# @!attribute [rw] payout
#   @return [Array, nil]
#
# @!attribute [rw] payouts_count
#   @return [Integer, nil]
#
# @!attribute [rw] reason
#   @return [String, nil]
#
# @!attribute [rw] schedule_date
#   @return [String, nil]
#
# @!attribute [rw] scheduled_date
#   @return [String, nil]
#
# @!attribute [rw] source_account
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] total_eur
#   @return [Float, nil]
#
# @!attribute [rw] total_gbp
#   @return [Float, nil]
#
# @!attribute [rw] total_usd
#   @return [Float, nil]
PayrunListMatch = Struct.new(
  :authorisation,
  :authorisation_date,
  :authorisers_completed_count,
  :authorisers_required_count,
  :batch_payout_id,
  :can_authorise,
  :can_delete,
  :can_edit,
  :event,
  :has_current_user_authorised,
  :id,
  :inserted,
  :invoice,
  :invoices_minimal,
  :is_archived,
  :last_updated,
  :last_updated_by,
  :merchant_id,
  :name,
  :nonce,
  :note,
  :payment,
  :payout,
  :payouts_count,
  :reason,
  :schedule_date,
  :scheduled_date,
  :source_account,
  :status,
  :total_eur,
  :total_gbp,
  :total_usd,
  keyword_init: true
)

# Request payload for Payrun#create.
#
# @!attribute [rw] id
#   @return [String]
PayrunCreateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Payrun#update.
#
# @!attribute [rw] id
#   @return [String]
PayrunUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Payrun#remove.
#
# @!attribute [rw] id
#   @return [String]
PayrunRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Report entity data model.
class Report
end

# Request payload for Report#update.
#
# @!attribute [rw] id
#   @return [String]
ReportUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# ReportResult entity data model.
#
# @!attribute [rw] content
#   @return [String, nil]
#
# @!attribute [rw] content_type
#   @return [String, nil]
#
# @!attribute [rw] last_completed_at
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] report_name
#   @return [String, nil]
#
# @!attribute [rw] report_type
#   @return [String, nil]
#
# @!attribute [rw] statement_number
#   @return [Integer, nil]
ReportResult = Struct.new(
  :content,
  :content_type,
  :last_completed_at,
  :merchant_id,
  :report_name,
  :report_type,
  :statement_number,
  keyword_init: true
)

# Request payload for ReportResult#load.
#
# @!attribute [rw] id
#   @return [Integer]
#
# @!attribute [rw] report_id
#   @return [String]
ReportResultLoadMatch = Struct.new(
  :id,
  :report_id,
  keyword_init: true
)

# Role entity data model.
#
# @!attribute [rw] failed_role
#   @return [Hash, nil]
#
# @!attribute [rw] role
#   @return [Array, nil]
Role = Struct.new(
  :failed_role,
  :role,
  keyword_init: true
)

# Request payload for Role#create.
#
# @!attribute [rw] merchant_id
#   @return [String]
RoleCreateData = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Rule entity data model.
#
# @!attribute [rw] account
#   @return [Hash, nil]
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] approve_url
#   @return [String, nil]
#
# @!attribute [rw] approver_id
#   @return [String, nil]
#
# @!attribute [rw] authentication_method
#   @return [Array, nil]
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] created_by
#   @return [Hash]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] end_at
#   @return [String, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_disabled
#   @return [Boolean, nil]
#
# @!attribute [rw] last_executed_at
#   @return [String, nil]
#
# @!attribute [rw] last_run_at_transaction_date
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String]
#
# @!attribute [rw] on_approved_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] on_execution_error_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] on_execution_success_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] start_at
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] sweep_action
#   @return [Hash, nil]
#
# @!attribute [rw] time_zone_id
#   @return [String, nil]
#
# @!attribute [rw] trigger_cron_expression
#   @return [String, nil]
#
# @!attribute [rw] trigger_on_pay_in
#   @return [Boolean, nil]
#
# @!attribute [rw] user_id
#   @return [String, nil]
#
# @!attribute [rw] web_hook_secret
#   @return [String, nil]
Rule = Struct.new(
  :account,
  :account_id,
  :approve_url,
  :approver_id,
  :authentication_method,
  :authorisation,
  :authorisers_completed_count,
  :authorisers_required_count,
  :can_authorise,
  :created_by,
  :description,
  :end_at,
  :has_current_user_authorised,
  :id,
  :inserted,
  :is_disabled,
  :last_executed_at,
  :last_run_at_transaction_date,
  :last_updated,
  :merchant_id,
  :name,
  :nonce,
  :on_approved_web_hook_url,
  :on_execution_error_web_hook_url,
  :on_execution_success_web_hook_url,
  :start_at,
  :status,
  :sweep_action,
  :time_zone_id,
  :trigger_cron_expression,
  :trigger_on_pay_in,
  :user_id,
  :web_hook_secret,
  keyword_init: true
)

# Request payload for Rule#load.
#
# @!attribute [rw] id
#   @return [String]
RuleLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Rule#list.
#
# @!attribute [rw] account
#   @return [Hash, nil]
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] approve_url
#   @return [String, nil]
#
# @!attribute [rw] approver_id
#   @return [String, nil]
#
# @!attribute [rw] authentication_method
#   @return [Array, nil]
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] created_by
#   @return [Hash, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] end_at
#   @return [String, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_disabled
#   @return [Boolean, nil]
#
# @!attribute [rw] last_executed_at
#   @return [String, nil]
#
# @!attribute [rw] last_run_at_transaction_date
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String, nil]
#
# @!attribute [rw] on_approved_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] on_execution_error_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] on_execution_success_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] start_at
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] sweep_action
#   @return [Hash, nil]
#
# @!attribute [rw] time_zone_id
#   @return [String, nil]
#
# @!attribute [rw] trigger_cron_expression
#   @return [String, nil]
#
# @!attribute [rw] trigger_on_pay_in
#   @return [Boolean, nil]
#
# @!attribute [rw] user_id
#   @return [String, nil]
#
# @!attribute [rw] web_hook_secret
#   @return [String, nil]
RuleListMatch = Struct.new(
  :account,
  :account_id,
  :approve_url,
  :approver_id,
  :authentication_method,
  :authorisation,
  :authorisers_completed_count,
  :authorisers_required_count,
  :can_authorise,
  :created_by,
  :description,
  :end_at,
  :has_current_user_authorised,
  :id,
  :inserted,
  :is_disabled,
  :last_executed_at,
  :last_run_at_transaction_date,
  :last_updated,
  :merchant_id,
  :name,
  :nonce,
  :on_approved_web_hook_url,
  :on_execution_error_web_hook_url,
  :on_execution_success_web_hook_url,
  :start_at,
  :status,
  :sweep_action,
  :time_zone_id,
  :trigger_cron_expression,
  :trigger_on_pay_in,
  :user_id,
  :web_hook_secret,
  keyword_init: true
)

# Request payload for Rule#create.
#
# @!attribute [rw] account
#   @return [Hash, nil]
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] approve_url
#   @return [String, nil]
#
# @!attribute [rw] approver_id
#   @return [String, nil]
#
# @!attribute [rw] authentication_method
#   @return [Array, nil]
#
# @!attribute [rw] authorisation
#   @return [Array, nil]
#
# @!attribute [rw] authorisers_completed_count
#   @return [Integer, nil]
#
# @!attribute [rw] authorisers_required_count
#   @return [Integer, nil]
#
# @!attribute [rw] can_authorise
#   @return [Boolean, nil]
#
# @!attribute [rw] created_by
#   @return [Hash]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] end_at
#   @return [String, nil]
#
# @!attribute [rw] has_current_user_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_disabled
#   @return [Boolean, nil]
#
# @!attribute [rw] last_executed_at
#   @return [String, nil]
#
# @!attribute [rw] last_run_at_transaction_date
#   @return [String, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] nonce
#   @return [String]
#
# @!attribute [rw] on_approved_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] on_execution_error_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] on_execution_success_web_hook_url
#   @return [String, nil]
#
# @!attribute [rw] start_at
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] sweep_action
#   @return [Hash, nil]
#
# @!attribute [rw] time_zone_id
#   @return [String, nil]
#
# @!attribute [rw] trigger_cron_expression
#   @return [String, nil]
#
# @!attribute [rw] trigger_on_pay_in
#   @return [Boolean, nil]
#
# @!attribute [rw] user_id
#   @return [String, nil]
#
# @!attribute [rw] web_hook_secret
#   @return [String, nil]
RuleCreateData = Struct.new(
  :account,
  :account_id,
  :approve_url,
  :approver_id,
  :authentication_method,
  :authorisation,
  :authorisers_completed_count,
  :authorisers_required_count,
  :can_authorise,
  :created_by,
  :description,
  :end_at,
  :has_current_user_authorised,
  :id,
  :inserted,
  :is_disabled,
  :last_executed_at,
  :last_run_at_transaction_date,
  :last_updated,
  :merchant_id,
  :name,
  :nonce,
  :on_approved_web_hook_url,
  :on_execution_error_web_hook_url,
  :on_execution_success_web_hook_url,
  :start_at,
  :status,
  :sweep_action,
  :time_zone_id,
  :trigger_cron_expression,
  :trigger_on_pay_in,
  :user_id,
  :web_hook_secret,
  keyword_init: true
)

# Request payload for Rule#update.
#
# @!attribute [rw] id
#   @return [String]
RuleUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Rule#remove.
#
# @!attribute [rw] id
#   @return [String]
RuleRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# RuleEvent entity data model.
#
# @!attribute [rw] error_message
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_authorise_to_enable
#   @return [Boolean, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] raw_response
#   @return [String, nil]
#
# @!attribute [rw] rule_event_type
#   @return [String, nil]
#
# @!attribute [rw] rule_id
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [Hash]
RuleEvent = Struct.new(
  :error_message,
  :id,
  :inserted,
  :is_authorise_to_enable,
  :message,
  :raw_response,
  :rule_event_type,
  :rule_id,
  :user,
  keyword_init: true
)

# Request payload for RuleEvent#list.
#
# @!attribute [rw] id
#   @return [String]
RuleEventListMatch = Struct.new(
  :id,
  keyword_init: true
)

# Tag entity data model.
#
# @!attribute [rw] colour_hex
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String]
#
# @!attribute [rw] name
#   @return [String]
Tag = Struct.new(
  :colour_hex,
  :description,
  :id,
  :merchant_id,
  :name,
  keyword_init: true
)

# Request payload for Tag#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
TagListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for Tag#create.
#
# @!attribute [rw] merchant_id
#   @return [String]
TagCreateData = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Token entity data model.
class Token
end

# Request payload for Token#create.
#
# @!attribute [rw] id
#   @return [String]
TokenCreateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Token#remove.
#
# @!attribute [rw] id
#   @return [String]
TokenRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Transaction entity data model.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] account_name
#   @return [String, nil]
#
# @!attribute [rw] account_sequence_number
#   @return [Integer, nil]
#
# @!attribute [rw] address_detail
#   @return [Hash, nil]
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] amount_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] booking_date_time
#   @return [String, nil]
#
# @!attribute [rw] charge_detail
#   @return [Hash, nil]
#
# @!attribute [rw] content
#   @return [Array, nil]
#
# @!attribute [rw] counterparty
#   @return [Hash, nil]
#
# @!attribute [rw] counterparty_summary
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] currency_exchange
#   @return [Hash, nil]
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] enrichment
#   @return [Hash, nil]
#
# @!attribute [rw] fx_amount
#   @return [Float, nil]
#
# @!attribute [rw] fx_currency
#   @return [String, nil]
#
# @!attribute [rw] fx_rate
#   @return [Float, nil]
#
# @!attribute [rw] gross_amount
#   @return [Hash]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] iso_bank_transaction_code
#   @return [Hash, nil]
#
# @!attribute [rw] merchant
#   @return [Hash, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] page_number
#   @return [Integer, nil]
#
# @!attribute [rw] page_size
#   @return [Integer, nil]
#
# @!attribute [rw] payee_detail
#   @return [Hash]
#
# @!attribute [rw] payer_detail
#   @return [Hash]
#
# @!attribute [rw] payment_request_custom_field
#   @return [Hash, nil]
#
# @!attribute [rw] payment_request_id
#   @return [String, nil]
#
# @!attribute [rw] payout_id
#   @return [String, nil]
#
# @!attribute [rw] proprietary_bank_transaction_code
#   @return [Hash, nil]
#
# @!attribute [rw] raw_reference
#   @return [String, nil]
#
# @!attribute [rw] reference
#   @return [String, nil]
#
# @!attribute [rw] rule_id
#   @return [String, nil]
#
# @!attribute [rw] statement_reference
#   @return [Array, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] supplementary_data
#   @return [Object, nil]
#
# @!attribute [rw] tag
#   @return [Array, nil]
#
# @!attribute [rw] their_reference
#   @return [String, nil]
#
# @!attribute [rw] total_page
#   @return [Integer, nil]
#
# @!attribute [rw] total_size
#   @return [Integer, nil]
#
# @!attribute [rw] transaction_amount
#   @return [Hash]
#
# @!attribute [rw] transaction_date
#   @return [String, nil]
#
# @!attribute [rw] transaction_information
#   @return [Array, nil]
#
# @!attribute [rw] transaction_mutability
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] value_date_time
#   @return [String, nil]
#
# @!attribute [rw] virtual_iban
#   @return [String, nil]
#
# @!attribute [rw] your_reference
#   @return [String, nil]
Transaction = Struct.new(
  :account_id,
  :account_name,
  :account_sequence_number,
  :address_detail,
  :amount,
  :amount_minor_unit,
  :balance,
  :balance_minor_unit,
  :booking_date_time,
  :charge_detail,
  :content,
  :counterparty,
  :counterparty_summary,
  :currency,
  :currency_exchange,
  :date,
  :description,
  :enrichment,
  :fx_amount,
  :fx_currency,
  :fx_rate,
  :gross_amount,
  :id,
  :inserted,
  :iso_bank_transaction_code,
  :merchant,
  :merchant_id,
  :page_number,
  :page_size,
  :payee_detail,
  :payer_detail,
  :payment_request_custom_field,
  :payment_request_id,
  :payout_id,
  :proprietary_bank_transaction_code,
  :raw_reference,
  :reference,
  :rule_id,
  :statement_reference,
  :status,
  :supplementary_data,
  :tag,
  :their_reference,
  :total_page,
  :total_size,
  :transaction_amount,
  :transaction_date,
  :transaction_information,
  :transaction_mutability,
  :type,
  :value_date_time,
  :virtual_iban,
  :your_reference,
  keyword_init: true
)

# Request payload for Transaction#load.
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] sequence_number
#   @return [Integer, nil]
#
# @!attribute [rw] transaction_id
#   @return [String, nil]
#
# @!attribute [rw] account_id
#   @return [String, nil]
TransactionLoadMatch = Struct.new(
  :id,
  :sequence_number,
  :transaction_id,
  :account_id,
  keyword_init: true
)

# Request payload for Transaction#list.
#
# @!attribute [rw] account_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
TransactionListMatch = Struct.new(
  :account_id,
  :id,
  :merchant_id,
  keyword_init: true
)

# Request payload for Transaction#create.
#
# @!attribute [rw] id
#   @return [String]
TransactionCreateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Transaction#remove.
#
# @!attribute [rw] id
#   @return [String]
TransactionRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# User entity data model.
#
# @!attribute [rw] client_session_timeout
#   @return [Array, nil]
#
# @!attribute [rw] email_address
#   @return [String]
#
# @!attribute [rw] first_name
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] last_name
#   @return [String]
#
# @!attribute [rw] passkey_added
#   @return [Boolean, nil]
#
# @!attribute [rw] permission
#   @return [Hash, nil]
#
# @!attribute [rw] profile
#   @return [String, nil]
#
# @!attribute [rw] roles_with_scope
#   @return [Array, nil]
#
# @!attribute [rw] two_factor_enabled
#   @return [Boolean, nil]
#
# @!attribute [rw] user_invite_id
#   @return [String, nil]
User = Struct.new(
  :client_session_timeout,
  :email_address,
  :first_name,
  :id,
  :last_name,
  :passkey_added,
  :permission,
  :profile,
  :roles_with_scope,
  :two_factor_enabled,
  :user_invite_id,
  keyword_init: true
)

# Request payload for User#list.
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
UserListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for User#update.
#
# @!attribute [rw] id
#   @return [String]
UserUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# UserInvite entity data model.
#
# @!attribute [rw] authorisation_status
#   @return [Hash, nil]
#
# @!attribute [rw] failed_user_invite
#   @return [Hash, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] initial_role_id
#   @return [String, nil]
#
# @!attribute [rw] invitee_email_address
#   @return [String, nil]
#
# @!attribute [rw] invitee_first_name
#   @return [String, nil]
#
# @!attribute [rw] invitee_last_name
#   @return [String, nil]
#
# @!attribute [rw] inviter_email_address
#   @return [String, nil]
#
# @!attribute [rw] inviter_first_name
#   @return [String, nil]
#
# @!attribute [rw] inviter_last_name
#   @return [String, nil]
#
# @!attribute [rw] is_authorised
#   @return [Boolean, nil]
#
# @!attribute [rw] is_invitee_registered
#   @return [Boolean, nil]
#
# @!attribute [rw] last_invited
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_name
#   @return [String, nil]
#
# @!attribute [rw] message
#   @return [String, nil]
#
# @!attribute [rw] registration_url
#   @return [String, nil]
#
# @!attribute [rw] send_invite_email
#   @return [Boolean, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [Hash]
#
# @!attribute [rw] user_id
#   @return [String, nil]
#
# @!attribute [rw] user_invite
#   @return [Array, nil]
UserInvite = Struct.new(
  :authorisation_status,
  :failed_user_invite,
  :id,
  :initial_role_id,
  :invitee_email_address,
  :invitee_first_name,
  :invitee_last_name,
  :inviter_email_address,
  :inviter_first_name,
  :inviter_last_name,
  :is_authorised,
  :is_invitee_registered,
  :last_invited,
  :merchant_id,
  :merchant_name,
  :message,
  :registration_url,
  :send_invite_email,
  :status,
  :user,
  :user_id,
  :user_invite,
  keyword_init: true
)

# Request payload for UserInvite#load.
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] userinvite_id
#   @return [String, nil]
UserInviteLoadMatch = Struct.new(
  :id,
  :userinvite_id,
  keyword_init: true
)

# Request payload for UserInvite#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
UserInviteListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for UserInvite#create.
#
# @!attribute [rw] id
#   @return [String, nil]
UserInviteCreateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for UserInvite#update.
#
# @!attribute [rw] id
#   @return [String]
UserInviteUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for UserInvite#remove.
#
# @!attribute [rw] id
#   @return [String]
UserInviteRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

# Virtual entity data model.
#
# @!attribute [rw] account_name
#   @return [String, nil]
#
# @!attribute [rw] account_supplier_name
#   @return [String, nil]
#
# @!attribute [rw] available_balance
#   @return [Float, nil]
#
# @!attribute [rw] available_balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] balance
#   @return [Float, nil]
#
# @!attribute [rw] balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] bank_name
#   @return [String, nil]
#
# @!attribute [rw] consent_id
#   @return [String, nil]
#
# @!attribute [rw] created_by
#   @return [Hash]
#
# @!attribute [rw] created_by_display_name
#   @return [String, nil]
#
# @!attribute [rw] currency
#   @return [String, nil]
#
# @!attribute [rw] default_payment_rail
#   @return [String, nil]
#
# @!attribute [rw] display_name
#   @return [String, nil]
#
# @!attribute [rw] expiry_date
#   @return [String, nil]
#
# @!attribute [rw] external_account_icon
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] identifier
#   @return [Hash]
#
# @!attribute [rw] inserted
#   @return [String, nil]
#
# @!attribute [rw] is_archived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_connected_account
#   @return [Boolean, nil]
#
# @!attribute [rw] is_default
#   @return [Boolean, nil]
#
# @!attribute [rw] is_trust_account
#   @return [Boolean, nil]
#
# @!attribute [rw] is_virtual
#   @return [Boolean, nil]
#
# @!attribute [rw] last_transaction
#   @return [Hash, nil]
#
# @!attribute [rw] last_updated
#   @return [String, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] merchant_name
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String]
#
# @!attribute [rw] physical_account_id
#   @return [String, nil]
#
# @!attribute [rw] rule
#   @return [Array, nil]
#
# @!attribute [rw] submitted_payouts_balance
#   @return [Float, nil]
#
# @!attribute [rw] submitted_payouts_balance_minor_unit
#   @return [Integer, nil]
#
# @!attribute [rw] summary
#   @return [String, nil]
#
# @!attribute [rw] supplier_sepa_instant_status
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_connection_status
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_last_synced_at
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_sync_last_failed_at
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_sync_last_failure_reason
#   @return [String, nil]
#
# @!attribute [rw] xero_bank_feed_sync_status
#   @return [String, nil]
#
# @!attribute [rw] xero_unsynchronised_transactions_count
#   @return [Integer, nil]
Virtual = Struct.new(
  :account_name,
  :account_supplier_name,
  :available_balance,
  :available_balance_minor_unit,
  :balance,
  :balance_minor_unit,
  :bank_name,
  :consent_id,
  :created_by,
  :created_by_display_name,
  :currency,
  :default_payment_rail,
  :display_name,
  :expiry_date,
  :external_account_icon,
  :id,
  :identifier,
  :inserted,
  :is_archived,
  :is_connected_account,
  :is_default,
  :is_trust_account,
  :is_virtual,
  :last_transaction,
  :last_updated,
  :merchant_id,
  :merchant_name,
  :name,
  :physical_account_id,
  :rule,
  :submitted_payouts_balance,
  :submitted_payouts_balance_minor_unit,
  :summary,
  :supplier_sepa_instant_status,
  :xero_bank_feed_connection_status,
  :xero_bank_feed_last_synced_at,
  :xero_bank_feed_sync_last_failed_at,
  :xero_bank_feed_sync_last_failure_reason,
  :xero_bank_feed_sync_status,
  :xero_unsynchronised_transactions_count,
  keyword_init: true
)

# Request payload for Virtual#create.
#
# @!attribute [rw] account_id
#   @return [String]
VirtualCreateData = Struct.new(
  :account_id,
  keyword_init: true
)

# Request payload for Virtual#update.
#
# @!attribute [rw] account_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
VirtualUpdateData = Struct.new(
  :account_id,
  :id,
  keyword_init: true
)

# Webhook entity data model.
#
# @!attribute [rw] destination_url
#   @return [String, nil]
#
# @!attribute [rw] email_address
#   @return [String, nil]
#
# @!attribute [rw] failed_notification_email_address
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_active
#   @return [Boolean, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] notification_method
#   @return [String, nil]
#
# @!attribute [rw] resource_type
#   @return [Array, nil]
#
# @!attribute [rw] retry
#   @return [Boolean, nil]
#
# @!attribute [rw] secret
#   @return [String, nil]
#
# @!attribute [rw] version
#   @return [Integer, nil]
Webhook = Struct.new(
  :destination_url,
  :email_address,
  :failed_notification_email_address,
  :id,
  :is_active,
  :merchant_id,
  :notification_method,
  :resource_type,
  :retry,
  :secret,
  :version,
  keyword_init: true
)

# Request payload for Webhook#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
WebhookLoadMatch = Struct.new(
  :id,
  :merchant_id,
  keyword_init: true
)

# Request payload for Webhook#list.
#
# @!attribute [rw] merchant_id
#   @return [String]
WebhookListMatch = Struct.new(
  :merchant_id,
  keyword_init: true
)

# Request payload for Webhook#create.
#
# @!attribute [rw] destination_url
#   @return [String, nil]
#
# @!attribute [rw] email_address
#   @return [String, nil]
#
# @!attribute [rw] failed_notification_email_address
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] is_active
#   @return [Boolean, nil]
#
# @!attribute [rw] merchant_id
#   @return [String, nil]
#
# @!attribute [rw] notification_method
#   @return [String, nil]
#
# @!attribute [rw] resource_type
#   @return [Array, nil]
#
# @!attribute [rw] retry
#   @return [Boolean, nil]
#
# @!attribute [rw] secret
#   @return [String, nil]
#
# @!attribute [rw] version
#   @return [Integer, nil]
WebhookCreateData = Struct.new(
  :destination_url,
  :email_address,
  :failed_notification_email_address,
  :id,
  :is_active,
  :merchant_id,
  :notification_method,
  :resource_type,
  :retry,
  :secret,
  :version,
  keyword_init: true
)

# Request payload for Webhook#update.
#
# @!attribute [rw] id
#   @return [String]
WebhookUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Webhook#remove.
#
# @!attribute [rw] id
#   @return [String]
WebhookRemoveMatch = Struct.new(
  :id,
  keyword_init: true
)

