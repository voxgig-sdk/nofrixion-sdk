// Typed models for the Nofrixion SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Account {
  account_id?: string
  account_name?: string
  account_supplier_name?: string
  account_type?: string
  available_balance?: number
  available_balance_minor_unit?: number
  balance?: number
  balance_minor_unit?: number
  bank_name?: string
  consent_id?: string
  created_by: Record<string, any>
  created_by_display_name?: string
  currency?: string
  default_payment_rail?: string
  display_name?: string
  expiry_date?: string
  external_account_icon?: string
  format?: string
  from_date?: string
  id?: string
  identifier: Record<string, any>
  inserted?: string
  is_archived?: boolean
  is_connected_account?: boolean
  is_default?: boolean
  is_trust_account?: boolean
  is_virtual?: boolean
  last_transaction?: Record<string, any>
  last_updated?: string
  merchant_id?: string
  merchant_name?: string
  physical_account_id?: string
  role_i_d?: any[]
  rule?: any[]
  submitted_payouts_balance?: number
  submitted_payouts_balance_minor_unit?: number
  summary?: string
  supplier_physical_account_id?: string
  supplier_sepa_instant_status?: string
  to_date?: string
  xero_bank_feed_connection_status?: string
  xero_bank_feed_last_synced_at?: string
  xero_bank_feed_sync_last_failed_at?: string
  xero_bank_feed_sync_last_failure_reason?: string
  xero_bank_feed_sync_status?: string
  xero_unsynchronised_transactions_count?: number
}

export interface AccountLoadMatch {
  account_id?: string
  id: string
  merchant_id?: string
}

export interface AccountListMatch {
  merchant_id?: string
}

export interface AccountCreateData {
  account_id?: string
  currency?: string
}

export interface AccountUpdateData {
  account_id?: string
  amount?: number
  id?: string
}

export interface AccountRemoveMatch {
  id: string
}

export interface Beneficiary {
  approval_callback_url?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  beneficiary_event?: any[]
  can_authorise?: boolean
  can_update?: boolean
  created_by: Record<string, any>
  created_by_email_address?: string
  currency: string
  destination?: Record<string, any>
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  is_enabled?: boolean
  last_authorised?: string
  last_updated?: string
  merchant_id?: string
  name: string
  nonce?: string
  source_account?: any[]
  source_account_i_d?: any[]
  their_reference?: string
}

export interface BeneficiaryLoadMatch {
  id: string
  merchant_id?: string
}

export interface BeneficiaryCreateData {
  id?: string
}

export interface BeneficiaryUpdateData {
  id: string
}

export interface BeneficiaryRemoveMatch {
  id: string
}

export interface Cancel {
  account_id?: string
  amount?: number
  amount_minor_unit?: number
  approve_payout_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  beneficiary: Record<string, any>
  can_authorise?: boolean
  can_process?: boolean
  can_update?: boolean
  charge_bearer?: string
  created_by?: string
  created_by_email_address?: string
  currency?: string
  current_user_id?: string
  description?: string
  destination?: Record<string, any>
  document?: any[]
  event?: any[]
  formatted_amount?: string
  formatted_fx_destination_amount?: string
  formatted_schedule?: string
  formatted_schedule_day_only?: string
  formatted_source_account_available_balance?: string
  fx_destination_amount?: number
  fx_destination_amount_minor_unit?: number
  fx_destination_currency?: string
  fx_quote_expires_at?: string
  fx_quote_id?: string
  fx_rate?: number
  fx_use_destination_amount?: boolean
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice_id?: string
  is_archived?: boolean
  is_failed?: boolean
  is_settled?: boolean
  is_submitted?: boolean
  last_updated?: string
  merchant_id?: string
  merchant_token_description?: string
  nonce?: string
  payment_processor?: string
  payment_rail?: string
  payrun_id?: string
  payrun_name?: string
  rule?: Record<string, any>
  schedule_date?: string
  scheduled?: boolean
  source_account_available_balance?: number
  source_account_available_balance_minor_unit?: number
  source_account_bic?: string
  source_account_currency?: string
  source_account_iban?: string
  source_account_identifier: Record<string, any>
  source_account_name?: string
  source_account_number?: string
  source_account_sortcode?: string
  status?: string
  tag?: any[]
  their_reference?: string
  topup_payrun_id?: string
  transacted_amount?: number
  transacted_fx_amount?: number
  transacted_fx_rate?: number
  type?: string
  user_id?: string
  your_reference?: string
}

export interface CancelUpdateData {
  id: string
}

export interface Disable {
  approval_callback_url?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  beneficiary_event?: any[]
  can_authorise?: boolean
  can_update?: boolean
  created_by: Record<string, any>
  created_by_email_address?: string
  currency: string
  destination?: Record<string, any>
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  is_enabled?: boolean
  last_authorised?: string
  last_updated?: string
  merchant_id?: string
  name: string
  nonce?: string
  source_account?: any[]
  their_reference?: string
}

export interface DisableUpdateData {
  id: string
}

export interface Enable {
  approval_callback_url?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  beneficiary_event?: any[]
  can_authorise?: boolean
  can_update?: boolean
  created_by: Record<string, any>
  created_by_email_address?: string
  currency: string
  destination?: Record<string, any>
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  is_enabled?: boolean
  last_authorised?: string
  last_updated?: string
  merchant_id?: string
  name: string
  nonce?: string
  source_account?: any[]
  their_reference?: string
}

export interface EnableUpdateData {
  id: string
}

export interface Merchant {
  reason?: string
}

export interface MerchantLoadMatch {
  merchant_id: string
}

export interface MerchantUpdateData {
  id: string
}

export interface MerchantRemoveMatch {
  id?: string
  user_id?: string
  merchant_id?: string
  tag_id?: string
}

export interface Metadata {
}

export interface MetadataLoadMatch {
}

export interface NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage {
  approved_at?: string
  currency?: string
  customer_account_number?: string
  customer_city?: string
  customer_country_code?: string
  customer_country_name?: string
  customer_email_address?: string
  customer_first_name?: string
  customer_iban?: string
  customer_last_name?: string
  customer_sort_code?: string
  id?: string
  inserted?: string
  is_recurring?: boolean
  last_updated?: string
  merchant_id?: string
  reference?: string
  status?: string
  supplier_bank_account_id?: string
  supplier_customer_id?: string
  supplier_mandate_id?: string
  supplier_name?: string
  supplier_status?: string
}

export interface NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageListMatch {
  approved_at?: string
  currency?: string
  customer_account_number?: string
  customer_city?: string
  customer_country_code?: string
  customer_country_name?: string
  customer_email_address?: string
  customer_first_name?: string
  customer_iban?: string
  customer_last_name?: string
  customer_sort_code?: string
  id?: string
  inserted?: string
  is_recurring?: boolean
  last_updated?: string
  merchant_id?: string
  reference?: string
  status?: string
  supplier_bank_account_id?: string
  supplier_customer_id?: string
  supplier_mandate_id?: string
  supplier_name?: string
  supplier_status?: string
}

export interface NoFrixionBizBizModelsPagingPaymentRequestPage {
  address?: any[]
  amount?: number
  amount_pending?: number
  amount_received?: number
  amount_refunded?: number
  auto_send_receipt?: boolean
  base_origin_url?: string
  callback_url?: string
  card_authorize_only?: boolean
  card_create_token?: boolean
  card_create_token_mode?: string
  card_ignore_cvn?: boolean
  card_processor_merchant_id?: string
  card_stripe_payment_intent_id?: string
  card_stripe_payment_intent_secret?: string
  created_by_user: Record<string, any>
  currency?: string
  custom_field?: any[]
  customer_email_address?: string
  customer_id?: string
  customer_name?: string
  description?: string
  destination_account?: Record<string, any>
  direct_debit_payment?: Record<string, any>
  due_date?: string
  event?: any[]
  failure_callback_url?: string
  field_display_setting?: any[]
  formatted_amount?: string
  hosted_pay_checkout_url?: string
  id?: string
  ignore_address_verification?: boolean
  inserted?: string
  inserted_sortable?: string
  is_archived?: boolean
  jwk?: string
  last_updated?: string
  lightning_invoice?: string
  lightning_invoice_expires_at?: string
  merchant_direct_debit_mandate_id?: string
  merchant_id?: string
  merchant_token_description?: string
  notification_email_address?: string
  notification_role_i_d?: any[]
  order_id?: string
  partial_payment_method?: string
  partial_payment_step?: string
  payment_attempt?: any[]
  payment_method?: any[]
  payment_processor?: string
  payrun_id?: string
  pisp_account_id?: string
  priority_bank_id?: string
  result?: Record<string, any>
  sandbox_settle_delay_in_second?: number
  shipping_address?: Record<string, any>
  status?: string
  success_web_hook_url?: string
  tag?: any[]
  title?: string
  tokenised_card?: any[]
  transaction?: any[]
  use_hosted_payment_page?: boolean
}

export interface NoFrixionBizBizModelsPagingPaymentRequestPageListMatch {
  address?: any[]
  amount?: number
  amount_pending?: number
  amount_received?: number
  amount_refunded?: number
  auto_send_receipt?: boolean
  base_origin_url?: string
  callback_url?: string
  card_authorize_only?: boolean
  card_create_token?: boolean
  card_create_token_mode?: string
  card_ignore_cvn?: boolean
  card_processor_merchant_id?: string
  card_stripe_payment_intent_id?: string
  card_stripe_payment_intent_secret?: string
  created_by_user?: Record<string, any>
  currency?: string
  custom_field?: any[]
  customer_email_address?: string
  customer_id?: string
  customer_name?: string
  description?: string
  destination_account?: Record<string, any>
  direct_debit_payment?: Record<string, any>
  due_date?: string
  event?: any[]
  failure_callback_url?: string
  field_display_setting?: any[]
  formatted_amount?: string
  hosted_pay_checkout_url?: string
  id?: string
  ignore_address_verification?: boolean
  inserted?: string
  inserted_sortable?: string
  is_archived?: boolean
  jwk?: string
  last_updated?: string
  lightning_invoice?: string
  lightning_invoice_expires_at?: string
  merchant_direct_debit_mandate_id?: string
  merchant_id?: string
  merchant_token_description?: string
  notification_email_address?: string
  notification_role_i_d?: any[]
  order_id?: string
  partial_payment_method?: string
  partial_payment_step?: string
  payment_attempt?: any[]
  payment_method?: any[]
  payment_processor?: string
  payrun_id?: string
  pisp_account_id?: string
  priority_bank_id?: string
  result?: Record<string, any>
  sandbox_settle_delay_in_second?: number
  shipping_address?: Record<string, any>
  status?: string
  success_web_hook_url?: string
  tag?: any[]
  title?: string
  tokenised_card?: any[]
  transaction?: any[]
  use_hosted_payment_page?: boolean
}

export interface NoFrixionBizBizModelsPagingPayoutPage {
  account_id?: string
  amount?: number
  amount_minor_unit?: number
  approve_payout_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  beneficiary: Record<string, any>
  can_authorise?: boolean
  can_process?: boolean
  can_update?: boolean
  charge_bearer?: string
  created_by?: string
  created_by_email_address?: string
  currency?: string
  current_user_id?: string
  description?: string
  destination?: Record<string, any>
  document?: any[]
  event?: any[]
  formatted_amount?: string
  formatted_fx_destination_amount?: string
  formatted_schedule?: string
  formatted_schedule_day_only?: string
  formatted_source_account_available_balance?: string
  fx_destination_amount?: number
  fx_destination_amount_minor_unit?: number
  fx_destination_currency?: string
  fx_quote_expires_at?: string
  fx_quote_id?: string
  fx_rate?: number
  fx_use_destination_amount?: boolean
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice_id?: string
  is_archived?: boolean
  is_failed?: boolean
  is_settled?: boolean
  is_submitted?: boolean
  last_updated?: string
  merchant_id?: string
  merchant_token_description?: string
  nonce?: string
  payment_processor?: string
  payment_rail?: string
  payrun_id?: string
  payrun_name?: string
  rule?: Record<string, any>
  schedule_date?: string
  scheduled?: boolean
  source_account_available_balance?: number
  source_account_available_balance_minor_unit?: number
  source_account_bic?: string
  source_account_currency?: string
  source_account_iban?: string
  source_account_identifier: Record<string, any>
  source_account_name?: string
  source_account_number?: string
  source_account_sortcode?: string
  status?: string
  tag?: any[]
  their_reference?: string
  topup_payrun_id?: string
  transacted_amount?: number
  transacted_fx_amount?: number
  transacted_fx_rate?: number
  type?: string
  user_id?: string
  your_reference?: string
}

export interface NoFrixionBizBizModelsPagingPayoutPageListMatch {
  account_id?: string
  merchant_id?: string
}

export interface NoFrixionBizBizModelsPagingPayrunPage {
  authorisation?: any[]
  authorisation_date?: string
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  can_authorise?: boolean
  can_delete?: boolean
  can_edit?: boolean
  event?: any[]
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice?: any[]
  invoices_minimal?: any[]
  is_archived?: boolean
  last_updated?: string
  last_updated_by: Record<string, any>
  merchant_id?: string
  name?: string
  nonce?: string
  payment?: any[]
  payout?: any[]
  payouts_count?: number
  schedule_date?: string
  source_account?: any[]
  status?: string
  total_eur?: number
  total_gbp?: number
  total_usd?: number
}

export interface NoFrixionBizBizModelsPagingPayrunPageListMatch {
  authorisation?: any[]
  authorisation_date?: string
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  can_authorise?: boolean
  can_delete?: boolean
  can_edit?: boolean
  event?: any[]
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice?: any[]
  invoices_minimal?: any[]
  is_archived?: boolean
  last_updated?: string
  last_updated_by?: Record<string, any>
  merchant_id?: string
  name?: string
  nonce?: string
  payment?: any[]
  payout?: any[]
  payouts_count?: number
  schedule_date?: string
  source_account?: any[]
  status?: string
  total_eur?: number
  total_gbp?: number
  total_usd?: number
}

export interface NoFrixionBizBizModelsPagingRuleEventsPage {
  error_message?: string
  id?: string
  inserted?: string
  is_authorise_to_enable?: boolean
  message?: string
  raw_response?: string
  rule_event_type?: string
  rule_id?: string
  user: Record<string, any>
}

export interface NoFrixionBizBizModelsPagingRuleEventsPageListMatch {
  rule_id: string
}

export interface NoFrixionBizBizModelsPagingRulesPage {
  account?: Record<string, any>
  account_id?: string
  approve_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  can_authorise?: boolean
  created_by: Record<string, any>
  description?: string
  end_at?: string
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  is_disabled?: boolean
  last_executed_at?: string
  last_run_at_transaction_date?: string
  last_updated?: string
  merchant_id?: string
  name?: string
  nonce: string
  on_approved_web_hook_url?: string
  on_execution_error_web_hook_url?: string
  on_execution_success_web_hook_url?: string
  start_at?: string
  status?: string
  sweep_action?: Record<string, any>
  time_zone_id?: string
  trigger_cron_expression?: string
  trigger_on_pay_in?: boolean
  user_id?: string
  web_hook_secret?: string
}

export interface NoFrixionBizBizModelsPagingRulesPageListMatch {
  account?: Record<string, any>
  account_id?: string
  approve_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  can_authorise?: boolean
  created_by?: Record<string, any>
  description?: string
  end_at?: string
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  is_disabled?: boolean
  last_executed_at?: string
  last_run_at_transaction_date?: string
  last_updated?: string
  merchant_id?: string
  name?: string
  nonce?: string
  on_approved_web_hook_url?: string
  on_execution_error_web_hook_url?: string
  on_execution_success_web_hook_url?: string
  start_at?: string
  status?: string
  sweep_action?: Record<string, any>
  time_zone_id?: string
  trigger_cron_expression?: string
  trigger_on_pay_in?: boolean
  user_id?: string
  web_hook_secret?: string
}

export interface NoFrixionBizBizModelsPaymentsCardPayment {
  authorized_amount?: string
  currency_code?: string
  is_payer_authentication_required?: boolean
  is_soft_decline?: boolean
  payer_authentication_access_token?: string
  payer_authentication_merchant_data?: string
  payer_authentication_url?: string
  payer_authentication_window_height?: number
  payer_authentication_window_width?: number
  payment_request_callback_url?: string
  payment_request_id?: string
  request_id?: string
  response_code?: string
  response_type?: string
  status?: string
  three_ds_redirect_url?: string
  transaction_id?: string
}

export interface NoFrixionBizBizModelsPaymentsCardPaymentCreateData {
  partial_refund_amount?: number
  paymentrequest_id: string
}

export interface NoFrixionBizBizModelsPaymentsCardPublicKey {
  jwt?: string
}

export interface NoFrixionBizBizModelsPaymentsCardPublicKeyLoadMatch {
  paymentrequest_id: string
}

export interface NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries {
  beneficiary?: any[]
  failed_beneficiary?: Record<string, any>
}

export interface NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesCreateData {
  beneficiary?: any[]
  failed_beneficiary?: Record<string, any>
}

export interface NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment {
  failed_payment_request?: Record<string, any>
  payment_request?: any[]
}

export interface NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentCreateData {
  failed_payment_request?: Record<string, any>
  payment_request?: any[]
}

export interface NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate {
  failed_role?: Record<string, any>
  role?: any[]
}

export interface NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateCreateData {
  merchant_id: string
}

export interface NoFrixionMoneyMoovApiFeaturesUserInvitesCreate {
  failed_user_invite?: Record<string, any>
  user_invite?: any[]
}

export interface NoFrixionMoneyMoovApiFeaturesUserInvitesCreateCreateData {
  failed_user_invite?: Record<string, any>
  user_invite?: any[]
}

export interface NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant {
  amount_lower?: number
  amount_upper?: number
  authorisation_type?: string
  beneficiaries_only?: boolean
  id?: string
  inserted?: string
  last_editor_cant_authorise?: boolean
  last_updated?: string
  merchant_id?: string
  number_of_authoriser?: number
  role_setting?: any[]
}

export interface NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantListMatch {
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsBatchPayout {
  approve_url?: string
  id?: string
  payout?: any[]
}

export interface NoFrixionMoneyMoovModelsBatchPayoutLoadMatch {
  id: string
}

export interface NoFrixionMoneyMoovModelsBatchPayoutCreateData {
  approve_url?: string
  id?: string
  payout?: any[]
}

export interface NoFrixionMoneyMoovModelsBeneficiaryGroupPage {
  group_member?: any[]
  group_name: string
  id?: string
  inserted?: string
  last_updated?: string
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsBeneficiaryGroupPageListMatch {
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsBeneficiaryPage {
  approval_callback_url?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  beneficiary_event?: any[]
  can_authorise?: boolean
  can_update?: boolean
  created_by: Record<string, any>
  created_by_email_address?: string
  currency: string
  destination?: Record<string, any>
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  is_enabled?: boolean
  last_authorised?: string
  last_updated?: string
  merchant_id?: string
  name: string
  nonce?: string
  source_account?: any[]
  their_reference?: string
}

export interface NoFrixionMoneyMoovModelsBeneficiaryPageListMatch {
  merchant_id?: string
}

export interface NoFrixionMoneyMoovModelsCardCustomerToken {
  card_type?: string
  customer_email_address?: string
  expiry_month?: string
  expiry_year?: string
  id?: string
  inserted?: string
  last_four_digit?: string
  last_updated?: string
  masked_card_number?: string
  merchant_id?: string
  payment_request_id?: string
}

export interface NoFrixionMoneyMoovModelsCardCustomerTokenLoadMatch {
  customer_email_address: string
}

export interface NoFrixionMoneyMoovModelsCardCustomerTokenListMatch {
  customer_email_address: string
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsCardCustomerTokenRemoveMatch {
  customer_email_address?: string
  merchant_id?: string
  id?: string
}

export interface NoFrixionMoneyMoovModelsCurrencyCurrencyInfo {
  code?: string
  decimal?: number
  is_fiat?: boolean
  iso4217_alpha_code?: string
  iso4217_numeric_code?: string
  symbol?: string
}

export interface NoFrixionMoneyMoovModelsCurrencyCurrencyInfoListMatch {
  code?: string
  decimal?: number
  is_fiat?: boolean
  iso4217_alpha_code?: string
  iso4217_numeric_code?: string
  symbol?: string
}

export interface NoFrixionMoneyMoovModelsDirectDebitBatchSubmit {
  failed_submission?: Record<string, any>
  successful_submission?: any[]
}

export interface NoFrixionMoneyMoovModelsDirectDebitBatchSubmitCreateData {
  failed_submission?: Record<string, any>
  successful_submission?: any[]
}

export interface NoFrixionMoneyMoovModelsFxRate {
  destination_currency?: string
  exchange_rate?: number
  expiry_time?: string
  quote_id?: string
  source_currency?: string
}

export interface NoFrixionMoneyMoovModelsFxRateLoadMatch {
  destination: string
  source: string
  valid_for_minute: number
}

export interface NoFrixionMoneyMoovModelsFxRateListMatch {
  destination: string
  source: string
}

export interface NoFrixionMoneyMoovModelsIPayment {
  payment_request_id?: string
  response_type?: string
}

export interface NoFrixionMoneyMoovModelsIPaymentCreateData {
  payment_request_id?: string
  response_type?: string
}

export interface NoFrixionMoneyMoovModelsMandatesMandate {
  account_number?: string
  address_line1: string
  address_line2?: string
  approved_at?: string
  city: string
  country_code: string
  currency?: string
  customer_account_number?: string
  customer_city?: string
  customer_country_code?: string
  customer_country_name?: string
  customer_email_address?: string
  customer_first_name?: string
  customer_iban?: string
  customer_last_name?: string
  customer_sort_code?: string
  email_address: string
  first_name: string
  iban?: string
  id?: string
  inserted?: string
  is_recurring?: boolean
  last_name: string
  last_updated?: string
  merchant_id?: string
  postal_code: string
  reference?: string
  sort_code?: string
  status?: string
  supplier_bank_account_id?: string
  supplier_customer_id?: string
  supplier_mandate_id?: string
  supplier_name?: string
  supplier_status?: string
}

export interface NoFrixionMoneyMoovModelsMandatesMandateLoadMatch {
  id: string
}

export interface NoFrixionMoneyMoovModelsMandatesMandateCreateData {
  account_number?: string
  address_line1: string
  address_line2?: string
  approved_at?: string
  city: string
  country_code: string
  currency?: string
  customer_account_number?: string
  customer_city?: string
  customer_country_code?: string
  customer_country_name?: string
  customer_email_address?: string
  customer_first_name?: string
  customer_iban?: string
  customer_last_name?: string
  customer_sort_code?: string
  email_address: string
  first_name: string
  iban?: string
  id?: string
  inserted?: string
  is_recurring?: boolean
  last_name: string
  last_updated?: string
  merchant_id?: string
  postal_code: string
  reference?: string
  sort_code?: string
  status?: string
  supplier_bank_account_id?: string
  supplier_customer_id?: string
  supplier_mandate_id?: string
  supplier_name?: string
  supplier_status?: string
}

export interface NoFrixionMoneyMoovModelsMerchant {
  account_currency?: any[]
  can_have_trust_account?: boolean
  card_payment_processor?: string
  company_id?: string
  display_qr_on_hosted_pay?: boolean
  hosted_pay_version?: number
  id?: string
  inserted?: string
  is_blocked?: boolean
  is_exited?: boolean
  is_suspended?: boolean
  jurisdiction?: string
  logo_url_png?: string
  logo_url_svg?: string
  merchant_category_code?: string
  name?: string
  note?: string
  parent_merchant?: Record<string, any>
  payment_account?: any[]
  payment_account_limit?: number
  short_name?: string
  supported_payment_methods_list?: any[]
  suspension_reason?: string
  tag?: any[]
  time_zone_id?: string
  trading_name?: string
  web_hook_limit?: number
  your_role_name?: string
}

export interface NoFrixionMoneyMoovModelsMerchantLoadMatch {
  id: string
}

export interface NoFrixionMoneyMoovModelsMerchantListMatch {
  account_currency?: any[]
  can_have_trust_account?: boolean
  card_payment_processor?: string
  company_id?: string
  display_qr_on_hosted_pay?: boolean
  hosted_pay_version?: number
  id?: string
  inserted?: string
  is_blocked?: boolean
  is_exited?: boolean
  is_suspended?: boolean
  jurisdiction?: string
  logo_url_png?: string
  logo_url_svg?: string
  merchant_category_code?: string
  name?: string
  note?: string
  parent_merchant?: Record<string, any>
  payment_account?: any[]
  payment_account_limit?: number
  short_name?: string
  supported_payment_methods_list?: any[]
  suspension_reason?: string
  tag?: any[]
  time_zone_id?: string
  trading_name?: string
  web_hook_limit?: number
  your_role_name?: string
}

export interface NoFrixionMoneyMoovModelsMerchantUpdateData {
  id: string
}

export interface NoFrixionMoneyMoovModelsMerchantPage {
  account_currency?: any[]
  can_have_trust_account?: boolean
  card_payment_processor?: string
  company_id?: string
  display_qr_on_hosted_pay?: boolean
  hosted_pay_version?: number
  id?: string
  inserted?: string
  is_blocked?: boolean
  is_exited?: boolean
  is_suspended?: boolean
  jurisdiction?: string
  logo_url_png?: string
  logo_url_svg?: string
  merchant_category_code?: string
  name?: string
  note?: string
  parent_merchant?: Record<string, any>
  payment_account?: any[]
  payment_account_limit?: number
  short_name?: string
  supported_payment_methods_list?: any[]
  suspension_reason?: string
  tag?: any[]
  time_zone_id?: string
  trading_name?: string
  web_hook_limit?: number
  your_role_name?: string
}

export interface NoFrixionMoneyMoovModelsMerchantPageListMatch {
  merchant_id?: string
}

export interface NoFrixionMoneyMoovModelsMerchantPayByBankSetting {
  bank_country_code?: any[]
  bank_id?: string
  bank_name?: string
  business_institution_id?: string
  currency?: string
  logo?: string
  message?: string
  message_image_url?: string
  order?: number
  personal_institution_id?: string
  processor?: string
  warning_heading?: string
  warning_message?: string
}

export interface NoFrixionMoneyMoovModelsMerchantPayByBankSettingListMatch {
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsMerchantToken {
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  can_authorise?: boolean
  description?: string
  expires_at?: string
  has_current_user_authorised?: boolean
  hmac_algorithm?: string
  id?: string
  inserted?: string
  ip_address_whitelist?: string
  is_archived?: boolean
  is_enabled?: boolean
  last_authorised?: string
  last_updated?: string
  merchant_id?: string
  nonce: string
  permission_type?: any[]
  request_signature_version?: number
  shared_secret_algorithm?: string
  shared_secret_base64?: string
  token?: string
}

export interface NoFrixionMoneyMoovModelsMerchantTokenLoadMatch {
  id: string
}

export interface NoFrixionMoneyMoovModelsMerchantTokenCreateData {
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  can_authorise?: boolean
  description?: string
  expires_at?: string
  has_current_user_authorised?: boolean
  hmac_algorithm?: string
  id?: string
  inserted?: string
  ip_address_whitelist?: string
  is_archived?: boolean
  is_enabled?: boolean
  last_authorised?: string
  last_updated?: string
  merchant_id?: string
  nonce: string
  permission_type?: any[]
  request_signature_version?: number
  shared_secret_algorithm?: string
  shared_secret_base64?: string
  token?: string
}

export interface NoFrixionMoneyMoovModelsMerchantTokenUpdateData {
  id: string
}

export interface NoFrixionMoneyMoovModelsMerchantTokenPage {
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  can_authorise?: boolean
  description?: string
  expires_at?: string
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  ip_address_whitelist?: string
  is_archived?: boolean
  is_enabled?: boolean
  last_authorised?: string
  last_updated?: string
  merchant_id?: string
  nonce: string
  permission_type?: any[]
  request_signature_version?: number
  shared_secret_algorithm?: string
  shared_secret_base64?: string
  token?: string
}

export interface NoFrixionMoneyMoovModelsMerchantTokenPageListMatch {
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsNoFrixionVersion {
  build_version?: number
  major_version?: number
  minor_version?: number
  release_name?: string
}

export interface NoFrixionMoneyMoovModelsNoFrixionVersionLoadMatch {
  build_version?: number
  major_version?: number
  minor_version?: number
  release_name?: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingAccount {
  account_balance?: any[]
  account_identification?: any[]
  account_name?: any[]
  account_type?: string
  balance?: number
  consolidated_account_information?: Record<string, any>
  currency?: string
  description?: string
  detail?: string
  id?: string
  nickname?: string
  type?: string
  usage_type?: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingAccountLoadMatch {
  id: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingConsent {
  authorisation_url?: string
  callback_url?: string
  consent_id?: string
  email_address?: string
  expiry_date?: string
  failure_callback_url?: string
  id?: string
  inserted?: string
  institution_id?: string
  is_connected_account?: boolean
  is_enabled?: boolean
  merchant_id?: string
  provider?: string
  success_web_hook_url?: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingConsentLoadMatch {
  id: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingConsentListMatch {
  email: string
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingConsentCreateData {
  authorisation_url?: string
  callback_url?: string
  consent_id?: string
  email_address?: string
  expiry_date?: string
  failure_callback_url?: string
  id?: string
  inserted?: string
  institution_id?: string
  is_connected_account?: boolean
  is_enabled?: boolean
  merchant_id?: string
  provider?: string
  success_web_hook_url?: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingConsentUpdateData {
  id: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingConsentRemoveMatch {
  id: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingTransaction {
  address_detail?: Record<string, any>
  amount?: number
  balance?: Record<string, any>
  booking_date_time?: string
  charge_detail?: Record<string, any>
  currency?: string
  currency_exchange?: Record<string, any>
  date?: string
  description?: string
  enrichment?: Record<string, any>
  gross_amount: Record<string, any>
  id?: string
  iso_bank_transaction_code?: Record<string, any>
  merchant?: Record<string, any>
  payee_detail: Record<string, any>
  payer_detail: Record<string, any>
  proprietary_bank_transaction_code?: Record<string, any>
  reference?: string
  statement_reference?: any[]
  status?: string
  supplementary_data?: any
  transaction_amount: Record<string, any>
  transaction_information?: any[]
  transaction_mutability?: string
  value_date_time?: string
}

export interface NoFrixionMoneyMoovModelsOpenBankingTransactionListMatch {
  account_id: string
  id: string
}

export interface NoFrixionMoneyMoovModelsPayment {
  address?: any[]
  amount?: number
  amount_pending?: number
  amount_received?: number
  amount_refunded?: number
  auto_send_receipt?: boolean
  base_origin_url?: string
  callback_url?: string
  card_authorize_only?: boolean
  card_create_token?: boolean
  card_create_token_mode?: string
  card_ignore_cvn?: boolean
  card_no_payer_authentication?: boolean
  card_processor_merchant_id?: string
  card_stripe_payment_intent_id?: string
  card_stripe_payment_intent_secret?: string
  card_transmit_raw_detail?: boolean
  created_by_user: Record<string, any>
  currency?: string
  custom_field?: any[]
  customer_email_address?: string
  customer_id?: string
  customer_name?: string
  description?: string
  destination_account?: Record<string, any>
  direct_debit_payment?: Record<string, any>
  due_date?: string
  event?: any[]
  failure_callback_url?: string
  field_display_setting?: any[]
  formatted_amount?: string
  hosted_pay_checkout_url?: string
  id?: string
  ignore_address_verification?: boolean
  inserted?: string
  inserted_sortable?: string
  is_archived?: boolean
  jwk?: string
  last_updated?: string
  lightning_invoice?: string
  lightning_invoice_expires_at?: string
  merchant_direct_debit_mandate_id?: string
  merchant_id?: string
  merchant_token_description?: string
  notification_email_address?: string
  notification_role_i_d?: any[]
  order_id?: string
  partial_payment_method?: string
  partial_payment_step?: string
  payment_attempt?: any[]
  payment_method?: any[]
  payment_processor?: string
  payrun_id?: string
  pisp_account_id?: string
  priority_bank_id?: string
  result?: Record<string, any>
  sandbox_settle_delay_in_second?: number
  shipping_address?: Record<string, any>
  shipping_address_city?: string
  shipping_address_country_code?: string
  shipping_address_county?: string
  shipping_address_line1?: string
  shipping_address_line2?: string
  shipping_address_post_code?: string
  shipping_email?: string
  shipping_first_name?: string
  shipping_last_name?: string
  shipping_phone?: string
  status?: string
  success_web_hook_url?: string
  tag?: any[]
  tag_id?: any[]
  title?: string
  tokenised_card?: any[]
  transaction?: any[]
  use_hosted_payment_page?: boolean
}

export interface NoFrixionMoneyMoovModelsPaymentLoadMatch {
  id?: string
  order_id?: string
}

export interface NoFrixionMoneyMoovModelsPaymentCreateData {
  address?: any[]
  amount?: number
  amount_pending?: number
  amount_received?: number
  amount_refunded?: number
  auto_send_receipt?: boolean
  base_origin_url?: string
  callback_url?: string
  card_authorize_only?: boolean
  card_create_token?: boolean
  card_create_token_mode?: string
  card_ignore_cvn?: boolean
  card_no_payer_authentication?: boolean
  card_processor_merchant_id?: string
  card_stripe_payment_intent_id?: string
  card_stripe_payment_intent_secret?: string
  card_transmit_raw_detail?: boolean
  created_by_user: Record<string, any>
  currency?: string
  custom_field?: any[]
  customer_email_address?: string
  customer_id?: string
  customer_name?: string
  description?: string
  destination_account?: Record<string, any>
  direct_debit_payment?: Record<string, any>
  due_date?: string
  event?: any[]
  failure_callback_url?: string
  field_display_setting?: any[]
  formatted_amount?: string
  hosted_pay_checkout_url?: string
  id?: string
  ignore_address_verification?: boolean
  inserted?: string
  inserted_sortable?: string
  is_archived?: boolean
  jwk?: string
  last_updated?: string
  lightning_invoice?: string
  lightning_invoice_expires_at?: string
  merchant_direct_debit_mandate_id?: string
  merchant_id?: string
  merchant_token_description?: string
  notification_email_address?: string
  notification_role_i_d?: any[]
  order_id?: string
  partial_payment_method?: string
  partial_payment_step?: string
  payment_attempt?: any[]
  payment_method?: any[]
  payment_processor?: string
  payrun_id?: string
  pisp_account_id?: string
  priority_bank_id?: string
  result?: Record<string, any>
  sandbox_settle_delay_in_second?: number
  shipping_address?: Record<string, any>
  shipping_address_city?: string
  shipping_address_country_code?: string
  shipping_address_county?: string
  shipping_address_line1?: string
  shipping_address_line2?: string
  shipping_address_post_code?: string
  shipping_email?: string
  shipping_first_name?: string
  shipping_last_name?: string
  shipping_phone?: string
  status?: string
  success_web_hook_url?: string
  tag?: any[]
  tag_id?: any[]
  title?: string
  tokenised_card?: any[]
  transaction?: any[]
  use_hosted_payment_page?: boolean
}

export interface NoFrixionMoneyMoovModelsPaymentUpdateData {
  id: string
}

export interface NoFrixionMoneyMoovModelsPaymentAccountMinimalPage {
  account_name?: string
  available_balance?: number
  balance?: number
  balance_minor_unit?: number
  currency?: string
  id?: string
  identifier: Record<string, any>
  is_archived?: boolean
  is_connected_account?: boolean
  merchant_id?: string
  submitted_payouts_balance?: number
}

export interface NoFrixionMoneyMoovModelsPaymentAccountMinimalPageListMatch {
  account_name?: string
  available_balance?: number
  balance?: number
  balance_minor_unit?: number
  currency?: string
  id?: string
  identifier?: Record<string, any>
  is_archived?: boolean
  is_connected_account?: boolean
  merchant_id?: string
  submitted_payouts_balance?: number
}

export interface NoFrixionMoneyMoovModelsPaymentAccountPage {
  account_name?: string
  account_supplier_name?: string
  available_balance?: number
  available_balance_minor_unit?: number
  balance?: number
  balance_minor_unit?: number
  bank_name?: string
  consent_id?: string
  created_by: Record<string, any>
  created_by_display_name?: string
  currency?: string
  default_payment_rail?: string
  display_name?: string
  expiry_date?: string
  external_account_icon?: string
  id?: string
  identifier: Record<string, any>
  inserted?: string
  is_archived?: boolean
  is_connected_account?: boolean
  is_default?: boolean
  is_trust_account?: boolean
  is_virtual?: boolean
  last_transaction?: Record<string, any>
  last_updated?: string
  merchant_id?: string
  merchant_name?: string
  physical_account_id?: string
  rule?: any[]
  submitted_payouts_balance?: number
  submitted_payouts_balance_minor_unit?: number
  summary?: string
  supplier_sepa_instant_status?: string
  xero_bank_feed_connection_status?: string
  xero_bank_feed_last_synced_at?: string
  xero_bank_feed_sync_last_failed_at?: string
  xero_bank_feed_sync_last_failure_reason?: string
  xero_bank_feed_sync_status?: string
  xero_unsynchronised_transactions_count?: number
}

export interface NoFrixionMoneyMoovModelsPaymentAccountPageListMatch {
  account_id?: string
}

export interface NoFrixionMoneyMoovModelsPaymentInitiation {
  payment_initiation_id?: string
  payment_request_callback_url?: string
  payment_request_id?: string
  redirect_url?: string
  response_type?: string
  specific_error_message?: string
}

export interface NoFrixionMoneyMoovModelsPaymentInitiationCreateData {
  paymentrequest_id: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestEvent {
  amount: number
  apple_pay_transaction_id?: string
  card_authorization_response_id?: string
  card_expiry_month?: number
  card_expiry_year?: number
  card_issuer?: string
  card_issuer_country?: string
  card_last_four_digit?: string
  card_request_id?: string
  card_scheme?: string
  card_token_customer_id?: string
  card_transaction_id?: string
  currency?: string
  direct_debit_payment_id?: string
  direct_debit_payment_reference?: string
  drirect_debit_mandate_id?: string
  error_message?: string
  error_reason?: string
  event_type?: string
  id?: string
  inserted?: string
  lightning_invoice?: string
  lightning_r_hash?: string
  origin_url?: string
  payment_method_type?: string
  payment_processor_name?: string
  payment_request_id?: string
  pisp_bank_status?: string
  pisp_payment_initiation_id?: string
  pisp_payment_institution_name?: string
  pisp_payment_service_provider_id?: string
  pisp_redirect_url?: string
  reconciled_transaction_id?: string
  refund_payout_id?: string
  status?: string
  wallet_name?: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestEventListMatch {
  paymentrequest_id: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestMetric {
  all?: number
  authorized?: number
  paid?: number
  partially_paid?: number
  total_amounts_by_currency?: Record<string, any>
  unpaid?: number
}

export interface NoFrixionMoneyMoovModelsPaymentRequestMetricLoadMatch {
  all?: number
  authorized?: number
  paid?: number
  partially_paid?: number
  total_amounts_by_currency?: Record<string, any>
  unpaid?: number
}

export interface NoFrixionMoneyMoovModelsPaymentRequestMinimal {
  amount?: number
  amount_pending?: number
  amount_received?: number
  amount_refunded?: number
  callback_url?: string
  card_stripe_payment_intent_secret?: string
  country_code?: string
  currency?: string
  custom_fields_to_display?: any[]
  description?: string
  due_date?: string
  field_display_setting?: any[]
  google_pay_merchant_id?: string
  id?: string
  jwk?: string
  merchant_id?: string
  merchant_logo_url_png?: string
  merchant_logo_url_svg?: string
  merchant_name?: string
  merchant_short_name?: string
  partial_payment_method?: string
  payment_attempt?: any[]
  payment_methods_list?: any[]
  payment_processor?: string
  payment_processor_key?: string
  pisp_error?: string
  priority_bank_id?: string
  status?: string
  stripe_account_id?: string
  title?: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestMinimalListMatch {
  paymentrequest_id: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestResult {
  amount?: number
  amount_pending?: number
  amount_received?: number
  amount_refunded?: number
  currency?: string
  customer_id?: string
  payment?: any[]
  payment_request_id?: string
  pisp_authorization?: any[]
  requested_amount?: number
  result?: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestResultListMatch {
  paymentrequest_id: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment {
  description: string
  id?: string
  inserted?: string
  last_updated?: string
  merchant_id?: string
  name: string
  template: Record<string, any>
}

export interface NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentListMatch {
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 {
  description: string
  id?: string
  inserted?: string
  last_updated?: string
  merchant_id?: string
  name: string
  template: Record<string, any>
}

export interface NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2LoadMatch {
  paymentrequest_id: string
  template_id: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 {
  description: string
  id?: string
  inserted?: string
  last_updated?: string
  merchant_id?: string
  name: string
  template: Record<string, any>
}

export interface NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3UpdateData {
  paymentrequest_id: string
  template_id: string
}

export interface NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 {
}

export interface NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4RemoveMatch {
  paymentrequest_id: string
  template_id: string
}

export interface NoFrixionMoneyMoovModelsPayoutKeysetPage {
  account_id?: string
  amount?: number
  amount_minor_unit?: number
  approve_payout_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  beneficiary: Record<string, any>
  can_authorise?: boolean
  can_process?: boolean
  can_update?: boolean
  charge_bearer?: string
  created_by?: string
  created_by_email_address?: string
  currency?: string
  current_user_id?: string
  description?: string
  destination?: Record<string, any>
  document?: any[]
  event?: any[]
  formatted_amount?: string
  formatted_fx_destination_amount?: string
  formatted_schedule?: string
  formatted_schedule_day_only?: string
  formatted_source_account_available_balance?: string
  fx_destination_amount?: number
  fx_destination_amount_minor_unit?: number
  fx_destination_currency?: string
  fx_quote_expires_at?: string
  fx_quote_id?: string
  fx_rate?: number
  fx_use_destination_amount?: boolean
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice_id?: string
  is_archived?: boolean
  is_failed?: boolean
  is_settled?: boolean
  is_submitted?: boolean
  last_updated?: string
  merchant_id?: string
  merchant_token_description?: string
  nonce?: string
  payment_processor?: string
  payment_rail?: string
  payrun_id?: string
  payrun_name?: string
  rule?: Record<string, any>
  schedule_date?: string
  scheduled?: boolean
  source_account_available_balance?: number
  source_account_available_balance_minor_unit?: number
  source_account_bic?: string
  source_account_currency?: string
  source_account_iban?: string
  source_account_identifier: Record<string, any>
  source_account_name?: string
  source_account_number?: string
  source_account_sortcode?: string
  status?: string
  tag?: any[]
  their_reference?: string
  topup_payrun_id?: string
  transacted_amount?: number
  transacted_fx_amount?: number
  transacted_fx_rate?: number
  type?: string
  user_id?: string
  your_reference?: string
}

export interface NoFrixionMoneyMoovModelsPayoutKeysetPageListMatch {
  account_id?: string
  merchant_id?: string
}

export interface NoFrixionMoneyMoovModelsPayoutMetric {
  all?: number
  failed?: number
  in_progress?: number
  paid?: number
  pending_approval?: number
  scheduled?: number
  total_amounts_by_currency?: Record<string, any>
}

export interface NoFrixionMoneyMoovModelsPayoutMetricLoadMatch {
  all?: number
  failed?: number
  in_progress?: number
  paid?: number
  pending_approval?: number
  scheduled?: number
  total_amounts_by_currency?: Record<string, any>
}

export interface NoFrixionMoneyMoovModelsPayoutsPayoutsCreate {
  failed_payout?: Record<string, any>
  payout?: any[]
}

export interface NoFrixionMoneyMoovModelsPayoutsPayoutsCreateCreateData {
  failed_payout?: Record<string, any>
  payout?: any[]
}

export interface NoFrixionMoneyMoovModelsPayrun {
  authorisation?: any[]
  authorisation_date?: string
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  can_authorise?: boolean
  can_delete?: boolean
  can_edit?: boolean
  event?: any[]
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice?: any[]
  invoices_minimal?: any[]
  is_archived?: boolean
  last_updated?: string
  last_updated_by: Record<string, any>
  merchant_id?: string
  name?: string
  nonce?: string
  payment?: any[]
  payout?: any[]
  payouts_count?: number
  reason?: string
  schedule_date?: string
  scheduled_date?: string
  source_account?: any[]
  status?: string
  total_eur?: number
  total_gbp?: number
  total_usd?: number
}

export interface NoFrixionMoneyMoovModelsPayrunLoadMatch {
  id: string
}

export interface NoFrixionMoneyMoovModelsPayrunCreateData {
  id: string
}

export interface NoFrixionMoneyMoovModelsPayrunUpdateData {
  id?: string
  payrun_id?: string
}

export interface NoFrixionMoneyMoovModelsReportResult {
  content?: string
  content_type?: string
  last_completed_at?: string
  merchant_id?: string
  report_name?: string
  report_type?: string
  statement_number?: number
}

export interface NoFrixionMoneyMoovModelsReportResultLoadMatch {
  id: number
  report_id: string
}

export interface NoFrixionMoneyMoovModelsRule {
  account?: Record<string, any>
  account_id?: string
  approve_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  can_authorise?: boolean
  created_by: Record<string, any>
  description?: string
  end_at?: string
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  is_disabled?: boolean
  last_executed_at?: string
  last_run_at_transaction_date?: string
  last_updated?: string
  merchant_id?: string
  name?: string
  nonce: string
  on_approved_web_hook_url?: string
  on_execution_error_web_hook_url?: string
  on_execution_success_web_hook_url?: string
  start_at?: string
  status?: string
  sweep_action?: Record<string, any>
  time_zone_id?: string
  trigger_cron_expression?: string
  trigger_on_pay_in?: boolean
  user_id?: string
  web_hook_secret?: string
}

export interface NoFrixionMoneyMoovModelsRuleLoadMatch {
  id: string
}

export interface NoFrixionMoneyMoovModelsRuleCreateData {
  account?: Record<string, any>
  account_id?: string
  approve_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  can_authorise?: boolean
  created_by: Record<string, any>
  description?: string
  end_at?: string
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  is_disabled?: boolean
  last_executed_at?: string
  last_run_at_transaction_date?: string
  last_updated?: string
  merchant_id?: string
  name?: string
  nonce: string
  on_approved_web_hook_url?: string
  on_execution_error_web_hook_url?: string
  on_execution_success_web_hook_url?: string
  start_at?: string
  status?: string
  sweep_action?: Record<string, any>
  time_zone_id?: string
  trigger_cron_expression?: string
  trigger_on_pay_in?: boolean
  user_id?: string
  web_hook_secret?: string
}

export interface NoFrixionMoneyMoovModelsRuleUpdateData {
  id: string
}

export interface NoFrixionMoneyMoovModelsTransaction {
  account_id?: string
  account_name?: string
  account_sequence_number?: number
  amount?: number
  amount_minor_unit?: number
  balance?: number
  balance_minor_unit?: number
  counterparty?: Record<string, any>
  counterparty_summary?: string
  currency?: string
  description?: string
  fx_amount?: number
  fx_currency?: string
  fx_rate?: number
  id?: string
  inserted?: string
  merchant_id?: string
  payment_request_custom_field?: Record<string, any>
  payment_request_id?: string
  payout_id?: string
  raw_reference?: string
  rule_id?: string
  tag?: any[]
  their_reference?: string
  transaction_date?: string
  type?: string
  virtual_iban?: string
  your_reference?: string
}

export interface NoFrixionMoneyMoovModelsTransactionLoadMatch {
  account_id?: string
  id: string
}

export interface NoFrixionMoneyMoovModelsTransactionPage {
  account_id?: string
  account_name?: string
  account_sequence_number?: number
  amount?: number
  amount_minor_unit?: number
  balance?: number
  balance_minor_unit?: number
  content?: any[]
  counterparty?: Record<string, any>
  counterparty_summary?: string
  currency?: string
  description?: string
  fx_amount?: number
  fx_currency?: string
  fx_rate?: number
  id?: string
  inserted?: string
  merchant_id?: string
  page_number?: number
  page_size?: number
  payment_request_custom_field?: Record<string, any>
  payment_request_id?: string
  payout_id?: string
  raw_reference?: string
  rule_id?: string
  tag?: any[]
  their_reference?: string
  total_page?: number
  total_size?: number
  transaction_date?: string
  type?: string
  virtual_iban?: string
  your_reference?: string
}

export interface NoFrixionMoneyMoovModelsTransactionPageLoadMatch {
  account_id: string
}

export interface NoFrixionMoneyMoovModelsTransactionPageListMatch {
  account_id?: string
  merchant_id?: string
}

export interface NoFrixionMoneyMoovModelsUserInvite {
  authorisation_status?: Record<string, any>
  id?: string
  initial_role_id?: string
  invitee_email_address?: string
  invitee_first_name?: string
  invitee_last_name?: string
  inviter_email_address?: string
  inviter_first_name?: string
  inviter_last_name?: string
  is_authorised?: boolean
  is_invitee_registered?: boolean
  last_invited?: string
  merchant_id?: string
  merchant_name?: string
  message?: string
  registration_url?: string
  send_invite_email?: boolean
  status?: string
  user: Record<string, any>
  user_id?: string
}

export interface NoFrixionMoneyMoovModelsUserInviteLoadMatch {
  id?: string
  userinvite_id?: string
}

export interface NoFrixionMoneyMoovModelsUserInviteCreateData {
  authorisation_status?: Record<string, any>
  id?: string
  initial_role_id?: string
  invitee_email_address?: string
  invitee_first_name?: string
  invitee_last_name?: string
  inviter_email_address?: string
  inviter_first_name?: string
  inviter_last_name?: string
  is_authorised?: boolean
  is_invitee_registered?: boolean
  last_invited?: string
  merchant_id?: string
  merchant_name?: string
  message?: string
  registration_url?: string
  send_invite_email?: boolean
  status?: string
  user: Record<string, any>
  user_id?: string
}

export interface NoFrixionMoneyMoovModelsUserInvitePage {
  authorisation_status?: Record<string, any>
  id?: string
  initial_role_id?: string
  invitee_email_address?: string
  invitee_first_name?: string
  invitee_last_name?: string
  inviter_email_address?: string
  inviter_first_name?: string
  inviter_last_name?: string
  is_authorised?: boolean
  is_invitee_registered?: boolean
  last_invited?: string
  merchant_id?: string
  merchant_name?: string
  message?: string
  registration_url?: string
  status?: string
  user: Record<string, any>
  user_id?: string
}

export interface NoFrixionMoneyMoovModelsUserInvitePageListMatch {
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsUserPage {
  client_session_timeout?: any[]
  email_address: string
  first_name: string
  id?: string
  last_name: string
  passkey_added?: boolean
  permission?: Record<string, any>
  roles_with_scope?: any[]
  two_factor_enabled?: boolean
}

export interface NoFrixionMoneyMoovModelsUserPageListMatch {
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsWebhook {
  destination_url?: string
  email_address?: string
  failed_notification_email_address?: string
  id?: string
  is_active?: boolean
  merchant_id?: string
  notification_method?: string
  resource_type?: any[]
  retry?: boolean
  secret?: string
  version?: number
}

export interface NoFrixionMoneyMoovModelsWebhookLoadMatch {
  id: string
  merchant_id?: string
}

export interface NoFrixionMoneyMoovModelsWebhookListMatch {
  merchant_id: string
}

export interface NoFrixionMoneyMoovModelsWebhookCreateData {
  destination_url?: string
  email_address?: string
  failed_notification_email_address?: string
  id?: string
  is_active?: boolean
  merchant_id?: string
  notification_method?: string
  resource_type?: any[]
  retry?: boolean
  secret?: string
  version?: number
}

export interface NoFrixionMoneyMoovModelsWebhookUpdateData {
  id: string
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
  account_name: string
  account_number?: string
  iban: string
  payee_verified_account_name?: string
  result?: string
  secondary_identification?: string
  sort_code?: string
}

export interface PayeeverificationCreateData {
  account_name: string
  account_number?: string
  iban: string
  payee_verified_account_name?: string
  result?: string
  secondary_identification?: string
  sort_code?: string
}

export interface PaymentRequest {
  amount?: number
  do_simulate_settlement_failure?: boolean
  error_description?: string
  institution?: string
  payment_initiation_id?: string
}

export interface PaymentRequestLoadMatch {
  paymentrequest_id?: string
}

export interface PaymentRequestCreateData {
  paymentrequest_id: string
}

export interface PaymentRequestUpdateData {
  paymentrequest_id: string
}

export interface PaymentRequestRemoveMatch {
  id: string
}

export interface Payout {
  account_id?: string
  allow_incomplete?: boolean
  amount?: number
  amount_minor_unit?: number
  approve_payout_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  beneficiary: Record<string, any>
  beneficiary_id?: string
  can_authorise?: boolean
  can_process?: boolean
  can_update?: boolean
  charge_bearer?: string
  created_by?: string
  created_by_email_address?: string
  currency?: string
  current_user_id?: string
  description?: string
  destination?: Record<string, any>
  document?: any[]
  event?: any[]
  formatted_amount?: string
  formatted_fx_destination_amount?: string
  formatted_schedule?: string
  formatted_schedule_day_only?: string
  formatted_source_account_available_balance?: string
  fx_destination_amount?: number
  fx_destination_amount_minor_unit?: number
  fx_destination_currency?: string
  fx_quote_expires_at?: string
  fx_quote_id?: string
  fx_rate?: number
  fx_use_destination_amount?: boolean
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice_id?: string
  is_archived?: boolean
  is_failed?: boolean
  is_settled?: boolean
  is_submitted?: boolean
  last_updated?: string
  merchant_id?: string
  merchant_token_description?: string
  nonce?: string
  payment_processor?: string
  payment_rail?: string
  payrun_id?: string
  payrun_name?: string
  rule?: Record<string, any>
  schedule_date?: string
  scheduled?: boolean
  source_account_available_balance?: number
  source_account_available_balance_minor_unit?: number
  source_account_bic?: string
  source_account_currency?: string
  source_account_iban?: string
  source_account_identifier: Record<string, any>
  source_account_name?: string
  source_account_number?: string
  source_account_sortcode?: string
  status?: string
  tag?: any[]
  tag_id?: any[]
  their_reference?: string
  topup_payrun_id?: string
  transacted_amount?: number
  transacted_fx_amount?: number
  transacted_fx_rate?: number
  type?: string
  user_id?: string
  your_reference?: string
}

export interface PayoutLoadMatch {
  amount?: number
  destination?: string
  source?: string
  id?: string
}

export interface PayoutCreateData {
  id?: string
}

export interface PayoutUpdateData {
  id: string
}

export interface PayoutRemoveMatch {
  id: string
}

export interface Payrun {
  id?: string
  note?: string
  scheduled_date?: string
}

export interface PayrunCreateData {
  id: string
}

export interface PayrunUpdateData {
  id: string
}

export interface PayrunRemoveMatch {
  id: string
}

export interface Reject {
  account_id?: string
  amount?: number
  amount_minor_unit?: number
  approve_payout_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  beneficiary: Record<string, any>
  can_authorise?: boolean
  can_process?: boolean
  can_update?: boolean
  charge_bearer?: string
  created_by?: string
  created_by_email_address?: string
  currency?: string
  current_user_id?: string
  description?: string
  destination?: Record<string, any>
  document?: any[]
  event?: any[]
  formatted_amount?: string
  formatted_fx_destination_amount?: string
  formatted_schedule?: string
  formatted_schedule_day_only?: string
  formatted_source_account_available_balance?: string
  fx_destination_amount?: number
  fx_destination_amount_minor_unit?: number
  fx_destination_currency?: string
  fx_quote_expires_at?: string
  fx_quote_id?: string
  fx_rate?: number
  fx_use_destination_amount?: boolean
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice_id?: string
  is_archived?: boolean
  is_failed?: boolean
  is_settled?: boolean
  is_submitted?: boolean
  last_updated?: string
  merchant_id?: string
  merchant_token_description?: string
  nonce?: string
  payment_processor?: string
  payment_rail?: string
  payrun_id?: string
  payrun_name?: string
  reason?: string
  rule?: Record<string, any>
  schedule_date?: string
  scheduled?: boolean
  source_account_available_balance?: number
  source_account_available_balance_minor_unit?: number
  source_account_bic?: string
  source_account_currency?: string
  source_account_iban?: string
  source_account_identifier: Record<string, any>
  source_account_name?: string
  source_account_number?: string
  source_account_sortcode?: string
  status?: string
  tag?: any[]
  their_reference?: string
  topup_payrun_id?: string
  transacted_amount?: number
  transacted_fx_amount?: number
  transacted_fx_rate?: number
  type?: string
  user_id?: string
  your_reference?: string
}

export interface RejectUpdateData {
  id: string
}

export interface Report {
}

export interface ReportUpdateData {
  id: string
}

export interface Rule {
}

export interface RuleUpdateData {
  id: string
}

export interface RuleRemoveMatch {
  id: string
}

export interface Send {
  account_id?: string
  allow_incomplete?: boolean
  amount?: number
  amount_minor_unit?: number
  approve_payout_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  beneficiary: Record<string, any>
  beneficiary_id?: string
  can_authorise?: boolean
  can_process?: boolean
  can_update?: boolean
  charge_bearer?: string
  created_by?: string
  created_by_email_address?: string
  currency?: string
  current_user_id?: string
  description?: string
  destination?: Record<string, any>
  document?: any[]
  event?: any[]
  formatted_amount?: string
  formatted_fx_destination_amount?: string
  formatted_schedule?: string
  formatted_schedule_day_only?: string
  formatted_source_account_available_balance?: string
  fx_destination_amount?: number
  fx_destination_amount_minor_unit?: number
  fx_destination_currency?: string
  fx_quote_expires_at?: string
  fx_quote_id?: string
  fx_rate?: number
  fx_use_destination_amount?: boolean
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice_id?: string
  is_archived?: boolean
  is_failed?: boolean
  is_settled?: boolean
  is_submitted?: boolean
  last_updated?: string
  merchant_id?: string
  merchant_token_description?: string
  nonce?: string
  payment_processor?: string
  payment_rail?: string
  payrun_id?: string
  payrun_name?: string
  rule?: Record<string, any>
  schedule_date?: string
  scheduled?: boolean
  source_account_available_balance?: number
  source_account_available_balance_minor_unit?: number
  source_account_bic?: string
  source_account_currency?: string
  source_account_iban?: string
  source_account_identifier: Record<string, any>
  source_account_name?: string
  source_account_number?: string
  source_account_sortcode?: string
  status?: string
  tag?: any[]
  tag_id?: any[]
  their_reference?: string
  topup_payrun_id?: string
  transacted_amount?: number
  transacted_fx_amount?: number
  transacted_fx_rate?: number
  type?: string
  user_id?: string
  your_reference?: string
}

export interface SendCreateData {
  account_id?: string
  allow_incomplete?: boolean
  amount?: number
  amount_minor_unit?: number
  approve_payout_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  beneficiary: Record<string, any>
  beneficiary_id?: string
  can_authorise?: boolean
  can_process?: boolean
  can_update?: boolean
  charge_bearer?: string
  created_by?: string
  created_by_email_address?: string
  currency?: string
  current_user_id?: string
  description?: string
  destination?: Record<string, any>
  document?: any[]
  event?: any[]
  formatted_amount?: string
  formatted_fx_destination_amount?: string
  formatted_schedule?: string
  formatted_schedule_day_only?: string
  formatted_source_account_available_balance?: string
  fx_destination_amount?: number
  fx_destination_amount_minor_unit?: number
  fx_destination_currency?: string
  fx_quote_expires_at?: string
  fx_quote_id?: string
  fx_rate?: number
  fx_use_destination_amount?: boolean
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice_id?: string
  is_archived?: boolean
  is_failed?: boolean
  is_settled?: boolean
  is_submitted?: boolean
  last_updated?: string
  merchant_id?: string
  merchant_token_description?: string
  nonce?: string
  payment_processor?: string
  payment_rail?: string
  payrun_id?: string
  payrun_name?: string
  rule?: Record<string, any>
  schedule_date?: string
  scheduled?: boolean
  source_account_available_balance?: number
  source_account_available_balance_minor_unit?: number
  source_account_bic?: string
  source_account_currency?: string
  source_account_iban?: string
  source_account_identifier: Record<string, any>
  source_account_name?: string
  source_account_number?: string
  source_account_sortcode?: string
  status?: string
  tag?: any[]
  tag_id?: any[]
  their_reference?: string
  topup_payrun_id?: string
  transacted_amount?: number
  transacted_fx_amount?: number
  transacted_fx_rate?: number
  type?: string
  user_id?: string
  your_reference?: string
}

export interface Sendbeneficiary {
  account_id?: string
  allow_incomplete?: boolean
  amount?: number
  amount_minor_unit?: number
  approve_payout_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  beneficiary: Record<string, any>
  beneficiary_id?: string
  can_authorise?: boolean
  can_process?: boolean
  can_update?: boolean
  charge_bearer?: string
  created_by?: string
  created_by_email_address?: string
  currency?: string
  current_user_id?: string
  description?: string
  destination?: Record<string, any>
  document?: any[]
  event?: any[]
  formatted_amount?: string
  formatted_fx_destination_amount?: string
  formatted_schedule?: string
  formatted_schedule_day_only?: string
  formatted_source_account_available_balance?: string
  fx_destination_amount?: number
  fx_destination_amount_minor_unit?: number
  fx_destination_currency?: string
  fx_quote_expires_at?: string
  fx_quote_id?: string
  fx_rate?: number
  fx_use_destination_amount?: boolean
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice_id?: string
  is_archived?: boolean
  is_failed?: boolean
  is_settled?: boolean
  is_submitted?: boolean
  last_updated?: string
  merchant_id?: string
  merchant_token_description?: string
  nonce?: string
  payment_processor?: string
  payment_rail?: string
  payrun_id?: string
  payrun_name?: string
  rule?: Record<string, any>
  schedule_date?: string
  scheduled?: boolean
  source_account_available_balance?: number
  source_account_available_balance_minor_unit?: number
  source_account_bic?: string
  source_account_currency?: string
  source_account_iban?: string
  source_account_identifier: Record<string, any>
  source_account_name?: string
  source_account_number?: string
  source_account_sortcode?: string
  status?: string
  tag?: any[]
  tag_id?: any[]
  their_reference?: string
  topup_payrun_id?: string
  transacted_amount?: number
  transacted_fx_amount?: number
  transacted_fx_rate?: number
  type?: string
  user_id?: string
  your_reference?: string
}

export interface SendbeneficiaryCreateData {
  account_id?: string
  allow_incomplete?: boolean
  amount?: number
  amount_minor_unit?: number
  approve_payout_url?: string
  approver_id?: string
  authentication_method?: any[]
  authorisation?: any[]
  authorisers_completed_count?: number
  authorisers_required_count?: number
  batch_payout_id?: string
  beneficiary: Record<string, any>
  beneficiary_id?: string
  can_authorise?: boolean
  can_process?: boolean
  can_update?: boolean
  charge_bearer?: string
  created_by?: string
  created_by_email_address?: string
  currency?: string
  current_user_id?: string
  description?: string
  destination?: Record<string, any>
  document?: any[]
  event?: any[]
  formatted_amount?: string
  formatted_fx_destination_amount?: string
  formatted_schedule?: string
  formatted_schedule_day_only?: string
  formatted_source_account_available_balance?: string
  fx_destination_amount?: number
  fx_destination_amount_minor_unit?: number
  fx_destination_currency?: string
  fx_quote_expires_at?: string
  fx_quote_id?: string
  fx_rate?: number
  fx_use_destination_amount?: boolean
  has_current_user_authorised?: boolean
  id?: string
  inserted?: string
  invoice_id?: string
  is_archived?: boolean
  is_failed?: boolean
  is_settled?: boolean
  is_submitted?: boolean
  last_updated?: string
  merchant_id?: string
  merchant_token_description?: string
  nonce?: string
  payment_processor?: string
  payment_rail?: string
  payrun_id?: string
  payrun_name?: string
  rule?: Record<string, any>
  schedule_date?: string
  scheduled?: boolean
  source_account_available_balance?: number
  source_account_available_balance_minor_unit?: number
  source_account_bic?: string
  source_account_currency?: string
  source_account_iban?: string
  source_account_identifier: Record<string, any>
  source_account_name?: string
  source_account_number?: string
  source_account_sortcode?: string
  status?: string
  tag?: any[]
  tag_id?: any[]
  their_reference?: string
  topup_payrun_id?: string
  transacted_amount?: number
  transacted_fx_amount?: number
  transacted_fx_rate?: number
  type?: string
  user_id?: string
  your_reference?: string
}

export interface Tag {
  colour_hex?: string
  description?: string
  id?: string
  merchant_id: string
  name: string
}

export interface TagListMatch {
  merchant_id: string
}

export interface TagCreateData {
  merchant_id: string
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
}

export interface TransactionLoadMatch {
  sequence_number: number
  transaction_id: string
}

export interface TransactionCreateData {
  id: string
}

export interface TransactionRemoveMatch {
  id: string
}

export interface User {
  client_session_timeout?: any[]
  email_address: string
  first_name: string
  id?: string
  last_name: string
  passkey_added?: boolean
  permission?: Record<string, any>
  profile?: string
  roles_with_scope?: any[]
  two_factor_enabled?: boolean
  user_invite_id?: string
}

export interface UserListMatch {
  merchant_id?: string
}

export interface UserUpdateData {
  id: string
}

export interface UserInvite {
}

export interface UserInviteCreateData {
  id: string
}

export interface UserInviteUpdateData {
  id: string
}

export interface UserInviteRemoveMatch {
  id: string
}

export interface Virtual {
  account_name?: string
  account_supplier_name?: string
  available_balance?: number
  available_balance_minor_unit?: number
  balance?: number
  balance_minor_unit?: number
  bank_name?: string
  consent_id?: string
  created_by: Record<string, any>
  created_by_display_name?: string
  currency?: string
  default_payment_rail?: string
  display_name?: string
  expiry_date?: string
  external_account_icon?: string
  id?: string
  identifier: Record<string, any>
  inserted?: string
  is_archived?: boolean
  is_connected_account?: boolean
  is_default?: boolean
  is_trust_account?: boolean
  is_virtual?: boolean
  last_transaction?: Record<string, any>
  last_updated?: string
  merchant_id?: string
  merchant_name?: string
  name: string
  physical_account_id?: string
  rule?: any[]
  submitted_payouts_balance?: number
  submitted_payouts_balance_minor_unit?: number
  summary?: string
  supplier_sepa_instant_status?: string
  xero_bank_feed_connection_status?: string
  xero_bank_feed_last_synced_at?: string
  xero_bank_feed_sync_last_failed_at?: string
  xero_bank_feed_sync_last_failure_reason?: string
  xero_bank_feed_sync_status?: string
  xero_unsynchronised_transactions_count?: number
}

export interface VirtualCreateData {
  account_id: string
}

export interface VirtualUpdateData {
  account_id: string
  id: string
}

export interface Webhook {
}

export interface WebhookRemoveMatch {
  id: string
}

export interface Whoami {
  client_session_timeout?: any[]
  email_address: string
  first_name: string
  id?: string
  last_name: string
  passkey_added?: boolean
  permission?: Record<string, any>
  roles_with_scope?: any[]
  two_factor_enabled?: boolean
}

export interface WhoamiListMatch {
  client_session_timeout?: any[]
  email_address?: string
  first_name?: string
  id?: string
  last_name?: string
  passkey_added?: boolean
  permission?: Record<string, any>
  roles_with_scope?: any[]
  two_factor_enabled?: boolean
}

export interface Whoamitrustedapp {
  client_session_timeout?: any[]
  email_address: string
  first_name: string
  id?: string
  last_name: string
  passkey_added?: boolean
  permission?: Record<string, any>
  roles_with_scope?: any[]
  two_factor_enabled?: boolean
}

export interface WhoamitrustedappListMatch {
  client_session_timeout?: any[]
  email_address?: string
  first_name?: string
  id?: string
  last_name?: string
  passkey_added?: boolean
  permission?: Record<string, any>
  roles_with_scope?: any[]
  two_factor_enabled?: boolean
}

