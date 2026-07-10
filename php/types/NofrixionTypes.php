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
    public ?string $account_id = null;
    public ?string $account_name = null;
    public ?string $account_supplier_name = null;
    public ?string $account_type = null;
    public ?float $available_balance = null;
    public ?int $available_balance_minor_unit = null;
    public ?float $balance = null;
    public ?int $balance_minor_unit = null;
    public ?string $bank_name = null;
    public ?string $consent_id = null;
    public array $created_by;
    public ?string $created_by_display_name = null;
    public ?string $currency = null;
    public ?string $default_payment_rail = null;
    public ?string $display_name = null;
    public ?string $expiry_date = null;
    public ?string $external_account_icon = null;
    public ?string $format = null;
    public ?string $from_date = null;
    public ?string $id = null;
    public array $identifier;
    public ?string $inserted = null;
    public ?bool $is_archived = null;
    public ?bool $is_connected_account = null;
    public ?bool $is_default = null;
    public ?bool $is_trust_account = null;
    public ?bool $is_virtual = null;
    public ?array $last_transaction = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_name = null;
    public ?string $physical_account_id = null;
    public ?array $role_i_d = null;
    public ?array $rule = null;
    public ?float $submitted_payouts_balance = null;
    public ?int $submitted_payouts_balance_minor_unit = null;
    public ?string $summary = null;
    public ?string $supplier_physical_account_id = null;
    public ?string $supplier_sepa_instant_status = null;
    public ?string $to_date = null;
    public ?string $xero_bank_feed_connection_status = null;
    public ?string $xero_bank_feed_last_synced_at = null;
    public ?string $xero_bank_feed_sync_last_failed_at = null;
    public ?string $xero_bank_feed_sync_last_failure_reason = null;
    public ?string $xero_bank_feed_sync_status = null;
    public ?int $xero_unsynchronised_transactions_count = null;
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
}

/** Request payload for Account#update. */
class AccountUpdateData
{
    public ?string $account_id = null;
    public ?float $amount = null;
    public ?string $id = null;
}

/** Request payload for Account#remove. */
class AccountRemoveMatch
{
    public string $id;
}

/** Beneficiary entity data model. */
class Beneficiary
{
    public ?string $approval_callback_url = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?array $beneficiary_event = null;
    public ?bool $can_authorise = null;
    public ?bool $can_update = null;
    public array $created_by;
    public ?string $created_by_email_address = null;
    public string $currency;
    public ?array $destination = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_enabled = null;
    public ?string $last_authorised = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $name;
    public ?string $nonce = null;
    public ?array $source_account = null;
    public ?array $source_account_i_d = null;
    public ?string $their_reference = null;
}

/** Request payload for Beneficiary#load. */
class BeneficiaryLoadMatch
{
    public string $id;
    public ?string $merchant_id = null;
}

/** Request payload for Beneficiary#create. */
class BeneficiaryCreateData
{
    public ?string $id = null;
}

/** Request payload for Beneficiary#update. */
class BeneficiaryUpdateData
{
    public string $id;
}

/** Request payload for Beneficiary#remove. */
class BeneficiaryRemoveMatch
{
    public string $id;
}

/** Cancel entity data model. */
class Cancel
{
    public ?string $account_id = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?string $approve_payout_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public array $beneficiary;
    public ?bool $can_authorise = null;
    public ?bool $can_process = null;
    public ?bool $can_update = null;
    public ?string $charge_bearer = null;
    public ?string $created_by = null;
    public ?string $created_by_email_address = null;
    public ?string $currency = null;
    public ?string $current_user_id = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $document = null;
    public ?array $event = null;
    public ?string $formatted_amount = null;
    public ?string $formatted_fx_destination_amount = null;
    public ?string $formatted_schedule = null;
    public ?string $formatted_schedule_day_only = null;
    public ?string $formatted_source_account_available_balance = null;
    public ?float $fx_destination_amount = null;
    public ?int $fx_destination_amount_minor_unit = null;
    public ?string $fx_destination_currency = null;
    public ?string $fx_quote_expires_at = null;
    public ?string $fx_quote_id = null;
    public ?float $fx_rate = null;
    public ?bool $fx_use_destination_amount = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoice_id = null;
    public ?bool $is_archived = null;
    public ?bool $is_failed = null;
    public ?bool $is_settled = null;
    public ?bool $is_submitted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $nonce = null;
    public ?string $payment_processor = null;
    public ?string $payment_rail = null;
    public ?string $payrun_id = null;
    public ?string $payrun_name = null;
    public ?array $rule = null;
    public ?string $schedule_date = null;
    public ?bool $scheduled = null;
    public ?float $source_account_available_balance = null;
    public ?int $source_account_available_balance_minor_unit = null;
    public ?string $source_account_bic = null;
    public ?string $source_account_currency = null;
    public ?string $source_account_iban = null;
    public array $source_account_identifier;
    public ?string $source_account_name = null;
    public ?string $source_account_number = null;
    public ?string $source_account_sortcode = null;
    public ?string $status = null;
    public ?array $tag = null;
    public ?string $their_reference = null;
    public ?string $topup_payrun_id = null;
    public ?float $transacted_amount = null;
    public ?float $transacted_fx_amount = null;
    public ?float $transacted_fx_rate = null;
    public ?string $type = null;
    public ?string $user_id = null;
    public ?string $your_reference = null;
}

/** Request payload for Cancel#update. */
class CancelUpdateData
{
    public string $id;
}

/** Disable entity data model. */
class Disable
{
    public ?string $approval_callback_url = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?array $beneficiary_event = null;
    public ?bool $can_authorise = null;
    public ?bool $can_update = null;
    public array $created_by;
    public ?string $created_by_email_address = null;
    public string $currency;
    public ?array $destination = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_enabled = null;
    public ?string $last_authorised = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $name;
    public ?string $nonce = null;
    public ?array $source_account = null;
    public ?string $their_reference = null;
}

/** Request payload for Disable#update. */
class DisableUpdateData
{
    public string $id;
}

/** Enable entity data model. */
class Enable
{
    public ?string $approval_callback_url = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?array $beneficiary_event = null;
    public ?bool $can_authorise = null;
    public ?bool $can_update = null;
    public array $created_by;
    public ?string $created_by_email_address = null;
    public string $currency;
    public ?array $destination = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_enabled = null;
    public ?string $last_authorised = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $name;
    public ?string $nonce = null;
    public ?array $source_account = null;
    public ?string $their_reference = null;
}

/** Request payload for Enable#update. */
class EnableUpdateData
{
    public string $id;
}

/** Merchant entity data model. */
class Merchant
{
    public ?string $reason = null;
}

/** Request payload for Merchant#load. */
class MerchantLoadMatch
{
    public string $merchant_id;
}

/** Request payload for Merchant#update. */
class MerchantUpdateData
{
    public string $id;
}

/** Request payload for Merchant#remove. */
class MerchantRemoveMatch
{
    public ?string $id = null;
    public ?string $user_id = null;
    public ?string $merchant_id = null;
    public ?string $tag_id = null;
}

/** Metadata entity data model. */
class Metadata
{
}

/** Request payload for Metadata#load. */
class MetadataLoadMatch
{
}

/** NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage entity data model. */
class NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage
{
    public ?string $approved_at = null;
    public ?string $currency = null;
    public ?string $customer_account_number = null;
    public ?string $customer_city = null;
    public ?string $customer_country_code = null;
    public ?string $customer_country_name = null;
    public ?string $customer_email_address = null;
    public ?string $customer_first_name = null;
    public ?string $customer_iban = null;
    public ?string $customer_last_name = null;
    public ?string $customer_sort_code = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_recurring = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $reference = null;
    public ?string $status = null;
    public ?string $supplier_bank_account_id = null;
    public ?string $supplier_customer_id = null;
    public ?string $supplier_mandate_id = null;
    public ?string $supplier_name = null;
    public ?string $supplier_status = null;
}

/** Request payload for NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage#list. */
class NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageListMatch
{
    public ?string $approved_at = null;
    public ?string $currency = null;
    public ?string $customer_account_number = null;
    public ?string $customer_city = null;
    public ?string $customer_country_code = null;
    public ?string $customer_country_name = null;
    public ?string $customer_email_address = null;
    public ?string $customer_first_name = null;
    public ?string $customer_iban = null;
    public ?string $customer_last_name = null;
    public ?string $customer_sort_code = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_recurring = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $reference = null;
    public ?string $status = null;
    public ?string $supplier_bank_account_id = null;
    public ?string $supplier_customer_id = null;
    public ?string $supplier_mandate_id = null;
    public ?string $supplier_name = null;
    public ?string $supplier_status = null;
}

/** NoFrixionBizBizModelsPagingPaymentRequestPage entity data model. */
class NoFrixionBizBizModelsPagingPaymentRequestPage
{
    public ?array $address = null;
    public ?float $amount = null;
    public ?float $amount_pending = null;
    public ?float $amount_received = null;
    public ?float $amount_refunded = null;
    public ?bool $auto_send_receipt = null;
    public ?string $base_origin_url = null;
    public ?string $callback_url = null;
    public ?bool $card_authorize_only = null;
    public ?bool $card_create_token = null;
    public ?string $card_create_token_mode = null;
    public ?bool $card_ignore_cvn = null;
    public ?string $card_processor_merchant_id = null;
    public ?string $card_stripe_payment_intent_id = null;
    public ?string $card_stripe_payment_intent_secret = null;
    public array $created_by_user;
    public ?string $currency = null;
    public ?array $custom_field = null;
    public ?string $customer_email_address = null;
    public ?string $customer_id = null;
    public ?string $customer_name = null;
    public ?string $description = null;
    public ?array $destination_account = null;
    public ?array $direct_debit_payment = null;
    public ?string $due_date = null;
    public ?array $event = null;
    public ?string $failure_callback_url = null;
    public ?array $field_display_setting = null;
    public ?string $formatted_amount = null;
    public ?string $hosted_pay_checkout_url = null;
    public ?string $id = null;
    public ?bool $ignore_address_verification = null;
    public ?string $inserted = null;
    public ?string $inserted_sortable = null;
    public ?bool $is_archived = null;
    public ?string $jwk = null;
    public ?string $last_updated = null;
    public ?string $lightning_invoice = null;
    public ?string $lightning_invoice_expires_at = null;
    public ?string $merchant_direct_debit_mandate_id = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $notification_email_address = null;
    public ?array $notification_role_i_d = null;
    public ?string $order_id = null;
    public ?string $partial_payment_method = null;
    public ?string $partial_payment_step = null;
    public ?array $payment_attempt = null;
    public ?array $payment_method = null;
    public ?string $payment_processor = null;
    public ?string $payrun_id = null;
    public ?string $pisp_account_id = null;
    public ?string $priority_bank_id = null;
    public ?array $result = null;
    public ?int $sandbox_settle_delay_in_second = null;
    public ?array $shipping_address = null;
    public ?string $status = null;
    public ?string $success_web_hook_url = null;
    public ?array $tag = null;
    public ?string $title = null;
    public ?array $tokenised_card = null;
    public ?array $transaction = null;
    public ?bool $use_hosted_payment_page = null;
}

/** Request payload for NoFrixionBizBizModelsPagingPaymentRequestPage#list. */
class NoFrixionBizBizModelsPagingPaymentRequestPageListMatch
{
    public ?array $address = null;
    public ?float $amount = null;
    public ?float $amount_pending = null;
    public ?float $amount_received = null;
    public ?float $amount_refunded = null;
    public ?bool $auto_send_receipt = null;
    public ?string $base_origin_url = null;
    public ?string $callback_url = null;
    public ?bool $card_authorize_only = null;
    public ?bool $card_create_token = null;
    public ?string $card_create_token_mode = null;
    public ?bool $card_ignore_cvn = null;
    public ?string $card_processor_merchant_id = null;
    public ?string $card_stripe_payment_intent_id = null;
    public ?string $card_stripe_payment_intent_secret = null;
    public ?array $created_by_user = null;
    public ?string $currency = null;
    public ?array $custom_field = null;
    public ?string $customer_email_address = null;
    public ?string $customer_id = null;
    public ?string $customer_name = null;
    public ?string $description = null;
    public ?array $destination_account = null;
    public ?array $direct_debit_payment = null;
    public ?string $due_date = null;
    public ?array $event = null;
    public ?string $failure_callback_url = null;
    public ?array $field_display_setting = null;
    public ?string $formatted_amount = null;
    public ?string $hosted_pay_checkout_url = null;
    public ?string $id = null;
    public ?bool $ignore_address_verification = null;
    public ?string $inserted = null;
    public ?string $inserted_sortable = null;
    public ?bool $is_archived = null;
    public ?string $jwk = null;
    public ?string $last_updated = null;
    public ?string $lightning_invoice = null;
    public ?string $lightning_invoice_expires_at = null;
    public ?string $merchant_direct_debit_mandate_id = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $notification_email_address = null;
    public ?array $notification_role_i_d = null;
    public ?string $order_id = null;
    public ?string $partial_payment_method = null;
    public ?string $partial_payment_step = null;
    public ?array $payment_attempt = null;
    public ?array $payment_method = null;
    public ?string $payment_processor = null;
    public ?string $payrun_id = null;
    public ?string $pisp_account_id = null;
    public ?string $priority_bank_id = null;
    public ?array $result = null;
    public ?int $sandbox_settle_delay_in_second = null;
    public ?array $shipping_address = null;
    public ?string $status = null;
    public ?string $success_web_hook_url = null;
    public ?array $tag = null;
    public ?string $title = null;
    public ?array $tokenised_card = null;
    public ?array $transaction = null;
    public ?bool $use_hosted_payment_page = null;
}

/** NoFrixionBizBizModelsPagingPayoutPage entity data model. */
class NoFrixionBizBizModelsPagingPayoutPage
{
    public ?string $account_id = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?string $approve_payout_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public array $beneficiary;
    public ?bool $can_authorise = null;
    public ?bool $can_process = null;
    public ?bool $can_update = null;
    public ?string $charge_bearer = null;
    public ?string $created_by = null;
    public ?string $created_by_email_address = null;
    public ?string $currency = null;
    public ?string $current_user_id = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $document = null;
    public ?array $event = null;
    public ?string $formatted_amount = null;
    public ?string $formatted_fx_destination_amount = null;
    public ?string $formatted_schedule = null;
    public ?string $formatted_schedule_day_only = null;
    public ?string $formatted_source_account_available_balance = null;
    public ?float $fx_destination_amount = null;
    public ?int $fx_destination_amount_minor_unit = null;
    public ?string $fx_destination_currency = null;
    public ?string $fx_quote_expires_at = null;
    public ?string $fx_quote_id = null;
    public ?float $fx_rate = null;
    public ?bool $fx_use_destination_amount = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoice_id = null;
    public ?bool $is_archived = null;
    public ?bool $is_failed = null;
    public ?bool $is_settled = null;
    public ?bool $is_submitted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $nonce = null;
    public ?string $payment_processor = null;
    public ?string $payment_rail = null;
    public ?string $payrun_id = null;
    public ?string $payrun_name = null;
    public ?array $rule = null;
    public ?string $schedule_date = null;
    public ?bool $scheduled = null;
    public ?float $source_account_available_balance = null;
    public ?int $source_account_available_balance_minor_unit = null;
    public ?string $source_account_bic = null;
    public ?string $source_account_currency = null;
    public ?string $source_account_iban = null;
    public array $source_account_identifier;
    public ?string $source_account_name = null;
    public ?string $source_account_number = null;
    public ?string $source_account_sortcode = null;
    public ?string $status = null;
    public ?array $tag = null;
    public ?string $their_reference = null;
    public ?string $topup_payrun_id = null;
    public ?float $transacted_amount = null;
    public ?float $transacted_fx_amount = null;
    public ?float $transacted_fx_rate = null;
    public ?string $type = null;
    public ?string $user_id = null;
    public ?string $your_reference = null;
}

/** Request payload for NoFrixionBizBizModelsPagingPayoutPage#list. */
class NoFrixionBizBizModelsPagingPayoutPageListMatch
{
    public ?string $account_id = null;
    public ?string $merchant_id = null;
}

/** NoFrixionBizBizModelsPagingPayrunPage entity data model. */
class NoFrixionBizBizModelsPagingPayrunPage
{
    public ?array $authorisation = null;
    public ?string $authorisation_date = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public ?bool $can_authorise = null;
    public ?bool $can_delete = null;
    public ?bool $can_edit = null;
    public ?array $event = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?array $invoice = null;
    public ?array $invoices_minimal = null;
    public ?bool $is_archived = null;
    public ?string $last_updated = null;
    public array $last_updated_by;
    public ?string $merchant_id = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?array $payment = null;
    public ?array $payout = null;
    public ?int $payouts_count = null;
    public ?string $schedule_date = null;
    public ?array $source_account = null;
    public ?string $status = null;
    public ?float $total_eur = null;
    public ?float $total_gbp = null;
    public ?float $total_usd = null;
}

/** Request payload for NoFrixionBizBizModelsPagingPayrunPage#list. */
class NoFrixionBizBizModelsPagingPayrunPageListMatch
{
    public ?array $authorisation = null;
    public ?string $authorisation_date = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public ?bool $can_authorise = null;
    public ?bool $can_delete = null;
    public ?bool $can_edit = null;
    public ?array $event = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?array $invoice = null;
    public ?array $invoices_minimal = null;
    public ?bool $is_archived = null;
    public ?string $last_updated = null;
    public ?array $last_updated_by = null;
    public ?string $merchant_id = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?array $payment = null;
    public ?array $payout = null;
    public ?int $payouts_count = null;
    public ?string $schedule_date = null;
    public ?array $source_account = null;
    public ?string $status = null;
    public ?float $total_eur = null;
    public ?float $total_gbp = null;
    public ?float $total_usd = null;
}

/** NoFrixionBizBizModelsPagingRuleEventsPage entity data model. */
class NoFrixionBizBizModelsPagingRuleEventsPage
{
    public ?string $error_message = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_authorise_to_enable = null;
    public ?string $message = null;
    public ?string $raw_response = null;
    public ?string $rule_event_type = null;
    public ?string $rule_id = null;
    public array $user;
}

/** Request payload for NoFrixionBizBizModelsPagingRuleEventsPage#list. */
class NoFrixionBizBizModelsPagingRuleEventsPageListMatch
{
    public string $rule_id;
}

/** NoFrixionBizBizModelsPagingRulesPage entity data model. */
class NoFrixionBizBizModelsPagingRulesPage
{
    public ?array $account = null;
    public ?string $account_id = null;
    public ?string $approve_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?bool $can_authorise = null;
    public array $created_by;
    public ?string $description = null;
    public ?string $end_at = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_disabled = null;
    public ?string $last_executed_at = null;
    public ?string $last_run_at_transaction_date = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $name = null;
    public string $nonce;
    public ?string $on_approved_web_hook_url = null;
    public ?string $on_execution_error_web_hook_url = null;
    public ?string $on_execution_success_web_hook_url = null;
    public ?string $start_at = null;
    public ?string $status = null;
    public ?array $sweep_action = null;
    public ?string $time_zone_id = null;
    public ?string $trigger_cron_expression = null;
    public ?bool $trigger_on_pay_in = null;
    public ?string $user_id = null;
    public ?string $web_hook_secret = null;
}

/** Request payload for NoFrixionBizBizModelsPagingRulesPage#list. */
class NoFrixionBizBizModelsPagingRulesPageListMatch
{
    public ?array $account = null;
    public ?string $account_id = null;
    public ?string $approve_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?bool $can_authorise = null;
    public ?array $created_by = null;
    public ?string $description = null;
    public ?string $end_at = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_disabled = null;
    public ?string $last_executed_at = null;
    public ?string $last_run_at_transaction_date = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?string $on_approved_web_hook_url = null;
    public ?string $on_execution_error_web_hook_url = null;
    public ?string $on_execution_success_web_hook_url = null;
    public ?string $start_at = null;
    public ?string $status = null;
    public ?array $sweep_action = null;
    public ?string $time_zone_id = null;
    public ?string $trigger_cron_expression = null;
    public ?bool $trigger_on_pay_in = null;
    public ?string $user_id = null;
    public ?string $web_hook_secret = null;
}

/** NoFrixionBizBizModelsPaymentsCardPayment entity data model. */
class NoFrixionBizBizModelsPaymentsCardPayment
{
    public ?string $authorized_amount = null;
    public ?string $currency_code = null;
    public ?bool $is_payer_authentication_required = null;
    public ?bool $is_soft_decline = null;
    public ?string $payer_authentication_access_token = null;
    public ?string $payer_authentication_merchant_data = null;
    public ?string $payer_authentication_url = null;
    public ?int $payer_authentication_window_height = null;
    public ?int $payer_authentication_window_width = null;
    public ?string $payment_request_callback_url = null;
    public ?string $payment_request_id = null;
    public ?string $request_id = null;
    public ?string $response_code = null;
    public ?string $response_type = null;
    public ?string $status = null;
    public ?string $three_ds_redirect_url = null;
    public ?string $transaction_id = null;
}

/** Request payload for NoFrixionBizBizModelsPaymentsCardPayment#create. */
class NoFrixionBizBizModelsPaymentsCardPaymentCreateData
{
    public ?float $partial_refund_amount = null;
    public string $paymentrequest_id;
}

/** NoFrixionBizBizModelsPaymentsCardPublicKey entity data model. */
class NoFrixionBizBizModelsPaymentsCardPublicKey
{
    public ?string $jwt = null;
}

/** Request payload for NoFrixionBizBizModelsPaymentsCardPublicKey#load. */
class NoFrixionBizBizModelsPaymentsCardPublicKeyLoadMatch
{
    public string $paymentrequest_id;
}

/** NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries entity data model. */
class NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries
{
    public ?array $beneficiary = null;
    public ?array $failed_beneficiary = null;
}

/** Request payload for NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries#create. */
class NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesCreateData
{
    public ?array $beneficiary = null;
    public ?array $failed_beneficiary = null;
}

/** NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment entity data model. */
class NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment
{
    public ?array $failed_payment_request = null;
    public ?array $payment_request = null;
}

/** Request payload for NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment#create. */
class NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentCreateData
{
    public ?array $failed_payment_request = null;
    public ?array $payment_request = null;
}

/** NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate entity data model. */
class NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate
{
    public ?array $failed_role = null;
    public ?array $role = null;
}

/** Request payload for NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate#create. */
class NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateCreateData
{
    public string $merchant_id;
}

/** NoFrixionMoneyMoovApiFeaturesUserInvitesCreate entity data model. */
class NoFrixionMoneyMoovApiFeaturesUserInvitesCreate
{
    public ?array $failed_user_invite = null;
    public ?array $user_invite = null;
}

/** Request payload for NoFrixionMoneyMoovApiFeaturesUserInvitesCreate#create. */
class NoFrixionMoneyMoovApiFeaturesUserInvitesCreateCreateData
{
    public ?array $failed_user_invite = null;
    public ?array $user_invite = null;
}

/** NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant entity data model. */
class NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant
{
    public ?float $amount_lower = null;
    public ?float $amount_upper = null;
    public ?string $authorisation_type = null;
    public ?bool $beneficiaries_only = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $last_editor_cant_authorise = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?int $number_of_authoriser = null;
    public ?array $role_setting = null;
}

/** Request payload for NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant#list. */
class NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantListMatch
{
    public string $merchant_id;
}

/** NoFrixionMoneyMoovModelsBatchPayout entity data model. */
class NoFrixionMoneyMoovModelsBatchPayout
{
    public ?string $approve_url = null;
    public ?string $id = null;
    public ?array $payout = null;
}

/** Request payload for NoFrixionMoneyMoovModelsBatchPayout#load. */
class NoFrixionMoneyMoovModelsBatchPayoutLoadMatch
{
    public string $id;
}

/** Request payload for NoFrixionMoneyMoovModelsBatchPayout#create. */
class NoFrixionMoneyMoovModelsBatchPayoutCreateData
{
    public ?string $approve_url = null;
    public ?string $id = null;
    public ?array $payout = null;
}

/** NoFrixionMoneyMoovModelsBeneficiaryGroupPage entity data model. */
class NoFrixionMoneyMoovModelsBeneficiaryGroupPage
{
    public ?array $group_member = null;
    public string $group_name;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $last_updated = null;
    public string $merchant_id;
}

/** Request payload for NoFrixionMoneyMoovModelsBeneficiaryGroupPage#list. */
class NoFrixionMoneyMoovModelsBeneficiaryGroupPageListMatch
{
    public string $merchant_id;
}

/** NoFrixionMoneyMoovModelsBeneficiaryPage entity data model. */
class NoFrixionMoneyMoovModelsBeneficiaryPage
{
    public ?string $approval_callback_url = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?array $beneficiary_event = null;
    public ?bool $can_authorise = null;
    public ?bool $can_update = null;
    public array $created_by;
    public ?string $created_by_email_address = null;
    public string $currency;
    public ?array $destination = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_enabled = null;
    public ?string $last_authorised = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $name;
    public ?string $nonce = null;
    public ?array $source_account = null;
    public ?string $their_reference = null;
}

/** Request payload for NoFrixionMoneyMoovModelsBeneficiaryPage#list. */
class NoFrixionMoneyMoovModelsBeneficiaryPageListMatch
{
    public ?string $merchant_id = null;
}

/** NoFrixionMoneyMoovModelsCardCustomerToken entity data model. */
class NoFrixionMoneyMoovModelsCardCustomerToken
{
    public ?string $card_type = null;
    public ?string $customer_email_address = null;
    public ?string $expiry_month = null;
    public ?string $expiry_year = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $last_four_digit = null;
    public ?string $last_updated = null;
    public ?string $masked_card_number = null;
    public ?string $merchant_id = null;
    public ?string $payment_request_id = null;
}

/** Request payload for NoFrixionMoneyMoovModelsCardCustomerToken#load. */
class NoFrixionMoneyMoovModelsCardCustomerTokenLoadMatch
{
    public string $customer_email_address;
}

/** Request payload for NoFrixionMoneyMoovModelsCardCustomerToken#list. */
class NoFrixionMoneyMoovModelsCardCustomerTokenListMatch
{
    public string $customer_email_address;
    public string $merchant_id;
}

/** Request payload for NoFrixionMoneyMoovModelsCardCustomerToken#remove. */
class NoFrixionMoneyMoovModelsCardCustomerTokenRemoveMatch
{
    public ?string $customer_email_address = null;
    public ?string $merchant_id = null;
    public ?string $id = null;
}

/** NoFrixionMoneyMoovModelsCurrencyCurrencyInfo entity data model. */
class NoFrixionMoneyMoovModelsCurrencyCurrencyInfo
{
    public ?string $code = null;
    public ?int $decimal = null;
    public ?bool $is_fiat = null;
    public ?string $iso4217_alpha_code = null;
    public ?string $iso4217_numeric_code = null;
    public ?string $symbol = null;
}

/** Request payload for NoFrixionMoneyMoovModelsCurrencyCurrencyInfo#list. */
class NoFrixionMoneyMoovModelsCurrencyCurrencyInfoListMatch
{
    public ?string $code = null;
    public ?int $decimal = null;
    public ?bool $is_fiat = null;
    public ?string $iso4217_alpha_code = null;
    public ?string $iso4217_numeric_code = null;
    public ?string $symbol = null;
}

/** NoFrixionMoneyMoovModelsDirectDebitBatchSubmit entity data model. */
class NoFrixionMoneyMoovModelsDirectDebitBatchSubmit
{
    public ?array $failed_submission = null;
    public ?array $successful_submission = null;
}

/** Request payload for NoFrixionMoneyMoovModelsDirectDebitBatchSubmit#create. */
class NoFrixionMoneyMoovModelsDirectDebitBatchSubmitCreateData
{
    public ?array $failed_submission = null;
    public ?array $successful_submission = null;
}

/** NoFrixionMoneyMoovModelsFxRate entity data model. */
class NoFrixionMoneyMoovModelsFxRate
{
    public ?string $destination_currency = null;
    public ?float $exchange_rate = null;
    public ?string $expiry_time = null;
    public ?string $quote_id = null;
    public ?string $source_currency = null;
}

/** Request payload for NoFrixionMoneyMoovModelsFxRate#load. */
class NoFrixionMoneyMoovModelsFxRateLoadMatch
{
    public string $destination;
    public string $source;
    public int $valid_for_minute;
}

/** Request payload for NoFrixionMoneyMoovModelsFxRate#list. */
class NoFrixionMoneyMoovModelsFxRateListMatch
{
    public string $destination;
    public string $source;
}

/** NoFrixionMoneyMoovModelsIPayment entity data model. */
class NoFrixionMoneyMoovModelsIPayment
{
    public ?string $payment_request_id = null;
    public ?string $response_type = null;
}

/** Request payload for NoFrixionMoneyMoovModelsIPayment#create. */
class NoFrixionMoneyMoovModelsIPaymentCreateData
{
    public ?string $payment_request_id = null;
    public ?string $response_type = null;
}

/** NoFrixionMoneyMoovModelsMandatesMandate entity data model. */
class NoFrixionMoneyMoovModelsMandatesMandate
{
    public ?string $account_number = null;
    public string $address_line1;
    public ?string $address_line2 = null;
    public ?string $approved_at = null;
    public string $city;
    public string $country_code;
    public ?string $currency = null;
    public ?string $customer_account_number = null;
    public ?string $customer_city = null;
    public ?string $customer_country_code = null;
    public ?string $customer_country_name = null;
    public ?string $customer_email_address = null;
    public ?string $customer_first_name = null;
    public ?string $customer_iban = null;
    public ?string $customer_last_name = null;
    public ?string $customer_sort_code = null;
    public string $email_address;
    public string $first_name;
    public ?string $iban = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_recurring = null;
    public string $last_name;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $postal_code;
    public ?string $reference = null;
    public ?string $sort_code = null;
    public ?string $status = null;
    public ?string $supplier_bank_account_id = null;
    public ?string $supplier_customer_id = null;
    public ?string $supplier_mandate_id = null;
    public ?string $supplier_name = null;
    public ?string $supplier_status = null;
}

/** Request payload for NoFrixionMoneyMoovModelsMandatesMandate#load. */
class NoFrixionMoneyMoovModelsMandatesMandateLoadMatch
{
    public string $id;
}

/** Request payload for NoFrixionMoneyMoovModelsMandatesMandate#create. */
class NoFrixionMoneyMoovModelsMandatesMandateCreateData
{
    public ?string $account_number = null;
    public string $address_line1;
    public ?string $address_line2 = null;
    public ?string $approved_at = null;
    public string $city;
    public string $country_code;
    public ?string $currency = null;
    public ?string $customer_account_number = null;
    public ?string $customer_city = null;
    public ?string $customer_country_code = null;
    public ?string $customer_country_name = null;
    public ?string $customer_email_address = null;
    public ?string $customer_first_name = null;
    public ?string $customer_iban = null;
    public ?string $customer_last_name = null;
    public ?string $customer_sort_code = null;
    public string $email_address;
    public string $first_name;
    public ?string $iban = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_recurring = null;
    public string $last_name;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $postal_code;
    public ?string $reference = null;
    public ?string $sort_code = null;
    public ?string $status = null;
    public ?string $supplier_bank_account_id = null;
    public ?string $supplier_customer_id = null;
    public ?string $supplier_mandate_id = null;
    public ?string $supplier_name = null;
    public ?string $supplier_status = null;
}

/** NoFrixionMoneyMoovModelsMerchant entity data model. */
class NoFrixionMoneyMoovModelsMerchant
{
    public ?array $account_currency = null;
    public ?bool $can_have_trust_account = null;
    public ?string $card_payment_processor = null;
    public ?string $company_id = null;
    public ?bool $display_qr_on_hosted_pay = null;
    public ?int $hosted_pay_version = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_blocked = null;
    public ?bool $is_exited = null;
    public ?bool $is_suspended = null;
    public ?string $jurisdiction = null;
    public ?string $logo_url_png = null;
    public ?string $logo_url_svg = null;
    public ?string $merchant_category_code = null;
    public ?string $name = null;
    public ?string $note = null;
    public ?array $parent_merchant = null;
    public ?array $payment_account = null;
    public ?int $payment_account_limit = null;
    public ?string $short_name = null;
    public ?array $supported_payment_methods_list = null;
    public ?string $suspension_reason = null;
    public ?array $tag = null;
    public ?string $time_zone_id = null;
    public ?string $trading_name = null;
    public ?int $web_hook_limit = null;
    public ?string $your_role_name = null;
}

/** Request payload for NoFrixionMoneyMoovModelsMerchant#load. */
class NoFrixionMoneyMoovModelsMerchantLoadMatch
{
    public string $id;
}

/** Request payload for NoFrixionMoneyMoovModelsMerchant#list. */
class NoFrixionMoneyMoovModelsMerchantListMatch
{
    public ?array $account_currency = null;
    public ?bool $can_have_trust_account = null;
    public ?string $card_payment_processor = null;
    public ?string $company_id = null;
    public ?bool $display_qr_on_hosted_pay = null;
    public ?int $hosted_pay_version = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_blocked = null;
    public ?bool $is_exited = null;
    public ?bool $is_suspended = null;
    public ?string $jurisdiction = null;
    public ?string $logo_url_png = null;
    public ?string $logo_url_svg = null;
    public ?string $merchant_category_code = null;
    public ?string $name = null;
    public ?string $note = null;
    public ?array $parent_merchant = null;
    public ?array $payment_account = null;
    public ?int $payment_account_limit = null;
    public ?string $short_name = null;
    public ?array $supported_payment_methods_list = null;
    public ?string $suspension_reason = null;
    public ?array $tag = null;
    public ?string $time_zone_id = null;
    public ?string $trading_name = null;
    public ?int $web_hook_limit = null;
    public ?string $your_role_name = null;
}

/** Request payload for NoFrixionMoneyMoovModelsMerchant#update. */
class NoFrixionMoneyMoovModelsMerchantUpdateData
{
    public string $id;
}

/** NoFrixionMoneyMoovModelsMerchantPage entity data model. */
class NoFrixionMoneyMoovModelsMerchantPage
{
    public ?array $account_currency = null;
    public ?bool $can_have_trust_account = null;
    public ?string $card_payment_processor = null;
    public ?string $company_id = null;
    public ?bool $display_qr_on_hosted_pay = null;
    public ?int $hosted_pay_version = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_blocked = null;
    public ?bool $is_exited = null;
    public ?bool $is_suspended = null;
    public ?string $jurisdiction = null;
    public ?string $logo_url_png = null;
    public ?string $logo_url_svg = null;
    public ?string $merchant_category_code = null;
    public ?string $name = null;
    public ?string $note = null;
    public ?array $parent_merchant = null;
    public ?array $payment_account = null;
    public ?int $payment_account_limit = null;
    public ?string $short_name = null;
    public ?array $supported_payment_methods_list = null;
    public ?string $suspension_reason = null;
    public ?array $tag = null;
    public ?string $time_zone_id = null;
    public ?string $trading_name = null;
    public ?int $web_hook_limit = null;
    public ?string $your_role_name = null;
}

/** Request payload for NoFrixionMoneyMoovModelsMerchantPage#list. */
class NoFrixionMoneyMoovModelsMerchantPageListMatch
{
    public ?string $merchant_id = null;
}

/** NoFrixionMoneyMoovModelsMerchantPayByBankSetting entity data model. */
class NoFrixionMoneyMoovModelsMerchantPayByBankSetting
{
    public ?array $bank_country_code = null;
    public ?string $bank_id = null;
    public ?string $bank_name = null;
    public ?string $business_institution_id = null;
    public ?string $currency = null;
    public ?string $logo = null;
    public ?string $message = null;
    public ?string $message_image_url = null;
    public ?int $order = null;
    public ?string $personal_institution_id = null;
    public ?string $processor = null;
    public ?string $warning_heading = null;
    public ?string $warning_message = null;
}

/** Request payload for NoFrixionMoneyMoovModelsMerchantPayByBankSetting#list. */
class NoFrixionMoneyMoovModelsMerchantPayByBankSettingListMatch
{
    public string $merchant_id;
}

/** NoFrixionMoneyMoovModelsMerchantToken entity data model. */
class NoFrixionMoneyMoovModelsMerchantToken
{
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?bool $can_authorise = null;
    public ?string $description = null;
    public ?string $expires_at = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $hmac_algorithm = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $ip_address_whitelist = null;
    public ?bool $is_archived = null;
    public ?bool $is_enabled = null;
    public ?string $last_authorised = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $nonce;
    public ?array $permission_type = null;
    public ?int $request_signature_version = null;
    public ?string $shared_secret_algorithm = null;
    public ?string $shared_secret_base64 = null;
    public ?string $token = null;
}

/** Request payload for NoFrixionMoneyMoovModelsMerchantToken#load. */
class NoFrixionMoneyMoovModelsMerchantTokenLoadMatch
{
    public string $id;
}

/** Request payload for NoFrixionMoneyMoovModelsMerchantToken#create. */
class NoFrixionMoneyMoovModelsMerchantTokenCreateData
{
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?bool $can_authorise = null;
    public ?string $description = null;
    public ?string $expires_at = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $hmac_algorithm = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $ip_address_whitelist = null;
    public ?bool $is_archived = null;
    public ?bool $is_enabled = null;
    public ?string $last_authorised = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $nonce;
    public ?array $permission_type = null;
    public ?int $request_signature_version = null;
    public ?string $shared_secret_algorithm = null;
    public ?string $shared_secret_base64 = null;
    public ?string $token = null;
}

/** Request payload for NoFrixionMoneyMoovModelsMerchantToken#update. */
class NoFrixionMoneyMoovModelsMerchantTokenUpdateData
{
    public string $id;
}

/** NoFrixionMoneyMoovModelsMerchantTokenPage entity data model. */
class NoFrixionMoneyMoovModelsMerchantTokenPage
{
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?bool $can_authorise = null;
    public ?string $description = null;
    public ?string $expires_at = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $ip_address_whitelist = null;
    public ?bool $is_archived = null;
    public ?bool $is_enabled = null;
    public ?string $last_authorised = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $nonce;
    public ?array $permission_type = null;
    public ?int $request_signature_version = null;
    public ?string $shared_secret_algorithm = null;
    public ?string $shared_secret_base64 = null;
    public ?string $token = null;
}

/** Request payload for NoFrixionMoneyMoovModelsMerchantTokenPage#list. */
class NoFrixionMoneyMoovModelsMerchantTokenPageListMatch
{
    public string $merchant_id;
}

/** NoFrixionMoneyMoovModelsNoFrixionVersion entity data model. */
class NoFrixionMoneyMoovModelsNoFrixionVersion
{
    public ?int $build_version = null;
    public ?int $major_version = null;
    public ?int $minor_version = null;
    public ?string $release_name = null;
}

/** Request payload for NoFrixionMoneyMoovModelsNoFrixionVersion#load. */
class NoFrixionMoneyMoovModelsNoFrixionVersionLoadMatch
{
    public ?int $build_version = null;
    public ?int $major_version = null;
    public ?int $minor_version = null;
    public ?string $release_name = null;
}

/** NoFrixionMoneyMoovModelsOpenBankingAccount entity data model. */
class NoFrixionMoneyMoovModelsOpenBankingAccount
{
    public ?array $account_balance = null;
    public ?array $account_identification = null;
    public ?array $account_name = null;
    public ?string $account_type = null;
    public ?float $balance = null;
    public ?array $consolidated_account_information = null;
    public ?string $currency = null;
    public ?string $description = null;
    public ?string $detail = null;
    public ?string $id = null;
    public ?string $nickname = null;
    public ?string $type = null;
    public ?string $usage_type = null;
}

/** Request payload for NoFrixionMoneyMoovModelsOpenBankingAccount#load. */
class NoFrixionMoneyMoovModelsOpenBankingAccountLoadMatch
{
    public string $id;
}

/** NoFrixionMoneyMoovModelsOpenBankingConsent entity data model. */
class NoFrixionMoneyMoovModelsOpenBankingConsent
{
    public ?string $authorisation_url = null;
    public ?string $callback_url = null;
    public ?string $consent_id = null;
    public ?string $email_address = null;
    public ?string $expiry_date = null;
    public ?string $failure_callback_url = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $institution_id = null;
    public ?bool $is_connected_account = null;
    public ?bool $is_enabled = null;
    public ?string $merchant_id = null;
    public ?string $provider = null;
    public ?string $success_web_hook_url = null;
}

/** Request payload for NoFrixionMoneyMoovModelsOpenBankingConsent#load. */
class NoFrixionMoneyMoovModelsOpenBankingConsentLoadMatch
{
    public string $id;
}

/** Request payload for NoFrixionMoneyMoovModelsOpenBankingConsent#list. */
class NoFrixionMoneyMoovModelsOpenBankingConsentListMatch
{
    public string $email;
    public string $merchant_id;
}

/** Request payload for NoFrixionMoneyMoovModelsOpenBankingConsent#create. */
class NoFrixionMoneyMoovModelsOpenBankingConsentCreateData
{
    public ?string $authorisation_url = null;
    public ?string $callback_url = null;
    public ?string $consent_id = null;
    public ?string $email_address = null;
    public ?string $expiry_date = null;
    public ?string $failure_callback_url = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $institution_id = null;
    public ?bool $is_connected_account = null;
    public ?bool $is_enabled = null;
    public ?string $merchant_id = null;
    public ?string $provider = null;
    public ?string $success_web_hook_url = null;
}

/** Request payload for NoFrixionMoneyMoovModelsOpenBankingConsent#update. */
class NoFrixionMoneyMoovModelsOpenBankingConsentUpdateData
{
    public string $id;
}

/** Request payload for NoFrixionMoneyMoovModelsOpenBankingConsent#remove. */
class NoFrixionMoneyMoovModelsOpenBankingConsentRemoveMatch
{
    public string $id;
}

/** NoFrixionMoneyMoovModelsOpenBankingTransaction entity data model. */
class NoFrixionMoneyMoovModelsOpenBankingTransaction
{
    public ?array $address_detail = null;
    public ?float $amount = null;
    public ?array $balance = null;
    public ?string $booking_date_time = null;
    public ?array $charge_detail = null;
    public ?string $currency = null;
    public ?array $currency_exchange = null;
    public ?string $date = null;
    public ?string $description = null;
    public ?array $enrichment = null;
    public array $gross_amount;
    public ?string $id = null;
    public ?array $iso_bank_transaction_code = null;
    public ?array $merchant = null;
    public array $payee_detail;
    public array $payer_detail;
    public ?array $proprietary_bank_transaction_code = null;
    public ?string $reference = null;
    public ?array $statement_reference = null;
    public ?string $status = null;
    public mixed $supplementary_data = null;
    public array $transaction_amount;
    public ?array $transaction_information = null;
    public ?string $transaction_mutability = null;
    public ?string $value_date_time = null;
}

/** Request payload for NoFrixionMoneyMoovModelsOpenBankingTransaction#list. */
class NoFrixionMoneyMoovModelsOpenBankingTransactionListMatch
{
    public string $account_id;
    public string $id;
}

/** NoFrixionMoneyMoovModelsPayment entity data model. */
class NoFrixionMoneyMoovModelsPayment
{
    public ?array $address = null;
    public ?float $amount = null;
    public ?float $amount_pending = null;
    public ?float $amount_received = null;
    public ?float $amount_refunded = null;
    public ?bool $auto_send_receipt = null;
    public ?string $base_origin_url = null;
    public ?string $callback_url = null;
    public ?bool $card_authorize_only = null;
    public ?bool $card_create_token = null;
    public ?string $card_create_token_mode = null;
    public ?bool $card_ignore_cvn = null;
    public ?bool $card_no_payer_authentication = null;
    public ?string $card_processor_merchant_id = null;
    public ?string $card_stripe_payment_intent_id = null;
    public ?string $card_stripe_payment_intent_secret = null;
    public ?bool $card_transmit_raw_detail = null;
    public array $created_by_user;
    public ?string $currency = null;
    public ?array $custom_field = null;
    public ?string $customer_email_address = null;
    public ?string $customer_id = null;
    public ?string $customer_name = null;
    public ?string $description = null;
    public ?array $destination_account = null;
    public ?array $direct_debit_payment = null;
    public ?string $due_date = null;
    public ?array $event = null;
    public ?string $failure_callback_url = null;
    public ?array $field_display_setting = null;
    public ?string $formatted_amount = null;
    public ?string $hosted_pay_checkout_url = null;
    public ?string $id = null;
    public ?bool $ignore_address_verification = null;
    public ?string $inserted = null;
    public ?string $inserted_sortable = null;
    public ?bool $is_archived = null;
    public ?string $jwk = null;
    public ?string $last_updated = null;
    public ?string $lightning_invoice = null;
    public ?string $lightning_invoice_expires_at = null;
    public ?string $merchant_direct_debit_mandate_id = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $notification_email_address = null;
    public ?array $notification_role_i_d = null;
    public ?string $order_id = null;
    public ?string $partial_payment_method = null;
    public ?string $partial_payment_step = null;
    public ?array $payment_attempt = null;
    public ?array $payment_method = null;
    public ?string $payment_processor = null;
    public ?string $payrun_id = null;
    public ?string $pisp_account_id = null;
    public ?string $priority_bank_id = null;
    public ?array $result = null;
    public ?int $sandbox_settle_delay_in_second = null;
    public ?array $shipping_address = null;
    public ?string $shipping_address_city = null;
    public ?string $shipping_address_country_code = null;
    public ?string $shipping_address_county = null;
    public ?string $shipping_address_line1 = null;
    public ?string $shipping_address_line2 = null;
    public ?string $shipping_address_post_code = null;
    public ?string $shipping_email = null;
    public ?string $shipping_first_name = null;
    public ?string $shipping_last_name = null;
    public ?string $shipping_phone = null;
    public ?string $status = null;
    public ?string $success_web_hook_url = null;
    public ?array $tag = null;
    public ?array $tag_id = null;
    public ?string $title = null;
    public ?array $tokenised_card = null;
    public ?array $transaction = null;
    public ?bool $use_hosted_payment_page = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPayment#load. */
class NoFrixionMoneyMoovModelsPaymentLoadMatch
{
    public ?string $id = null;
    public ?string $order_id = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPayment#create. */
class NoFrixionMoneyMoovModelsPaymentCreateData
{
    public ?array $address = null;
    public ?float $amount = null;
    public ?float $amount_pending = null;
    public ?float $amount_received = null;
    public ?float $amount_refunded = null;
    public ?bool $auto_send_receipt = null;
    public ?string $base_origin_url = null;
    public ?string $callback_url = null;
    public ?bool $card_authorize_only = null;
    public ?bool $card_create_token = null;
    public ?string $card_create_token_mode = null;
    public ?bool $card_ignore_cvn = null;
    public ?bool $card_no_payer_authentication = null;
    public ?string $card_processor_merchant_id = null;
    public ?string $card_stripe_payment_intent_id = null;
    public ?string $card_stripe_payment_intent_secret = null;
    public ?bool $card_transmit_raw_detail = null;
    public array $created_by_user;
    public ?string $currency = null;
    public ?array $custom_field = null;
    public ?string $customer_email_address = null;
    public ?string $customer_id = null;
    public ?string $customer_name = null;
    public ?string $description = null;
    public ?array $destination_account = null;
    public ?array $direct_debit_payment = null;
    public ?string $due_date = null;
    public ?array $event = null;
    public ?string $failure_callback_url = null;
    public ?array $field_display_setting = null;
    public ?string $formatted_amount = null;
    public ?string $hosted_pay_checkout_url = null;
    public ?string $id = null;
    public ?bool $ignore_address_verification = null;
    public ?string $inserted = null;
    public ?string $inserted_sortable = null;
    public ?bool $is_archived = null;
    public ?string $jwk = null;
    public ?string $last_updated = null;
    public ?string $lightning_invoice = null;
    public ?string $lightning_invoice_expires_at = null;
    public ?string $merchant_direct_debit_mandate_id = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $notification_email_address = null;
    public ?array $notification_role_i_d = null;
    public ?string $order_id = null;
    public ?string $partial_payment_method = null;
    public ?string $partial_payment_step = null;
    public ?array $payment_attempt = null;
    public ?array $payment_method = null;
    public ?string $payment_processor = null;
    public ?string $payrun_id = null;
    public ?string $pisp_account_id = null;
    public ?string $priority_bank_id = null;
    public ?array $result = null;
    public ?int $sandbox_settle_delay_in_second = null;
    public ?array $shipping_address = null;
    public ?string $shipping_address_city = null;
    public ?string $shipping_address_country_code = null;
    public ?string $shipping_address_county = null;
    public ?string $shipping_address_line1 = null;
    public ?string $shipping_address_line2 = null;
    public ?string $shipping_address_post_code = null;
    public ?string $shipping_email = null;
    public ?string $shipping_first_name = null;
    public ?string $shipping_last_name = null;
    public ?string $shipping_phone = null;
    public ?string $status = null;
    public ?string $success_web_hook_url = null;
    public ?array $tag = null;
    public ?array $tag_id = null;
    public ?string $title = null;
    public ?array $tokenised_card = null;
    public ?array $transaction = null;
    public ?bool $use_hosted_payment_page = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPayment#update. */
class NoFrixionMoneyMoovModelsPaymentUpdateData
{
    public string $id;
}

/** NoFrixionMoneyMoovModelsPaymentAccountMinimalPage entity data model. */
class NoFrixionMoneyMoovModelsPaymentAccountMinimalPage
{
    public ?string $account_name = null;
    public ?float $available_balance = null;
    public ?float $balance = null;
    public ?int $balance_minor_unit = null;
    public ?string $currency = null;
    public ?string $id = null;
    public array $identifier;
    public ?bool $is_archived = null;
    public ?bool $is_connected_account = null;
    public ?string $merchant_id = null;
    public ?float $submitted_payouts_balance = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentAccountMinimalPage#list. */
class NoFrixionMoneyMoovModelsPaymentAccountMinimalPageListMatch
{
    public ?string $account_name = null;
    public ?float $available_balance = null;
    public ?float $balance = null;
    public ?int $balance_minor_unit = null;
    public ?string $currency = null;
    public ?string $id = null;
    public ?array $identifier = null;
    public ?bool $is_archived = null;
    public ?bool $is_connected_account = null;
    public ?string $merchant_id = null;
    public ?float $submitted_payouts_balance = null;
}

/** NoFrixionMoneyMoovModelsPaymentAccountPage entity data model. */
class NoFrixionMoneyMoovModelsPaymentAccountPage
{
    public ?string $account_name = null;
    public ?string $account_supplier_name = null;
    public ?float $available_balance = null;
    public ?int $available_balance_minor_unit = null;
    public ?float $balance = null;
    public ?int $balance_minor_unit = null;
    public ?string $bank_name = null;
    public ?string $consent_id = null;
    public array $created_by;
    public ?string $created_by_display_name = null;
    public ?string $currency = null;
    public ?string $default_payment_rail = null;
    public ?string $display_name = null;
    public ?string $expiry_date = null;
    public ?string $external_account_icon = null;
    public ?string $id = null;
    public array $identifier;
    public ?string $inserted = null;
    public ?bool $is_archived = null;
    public ?bool $is_connected_account = null;
    public ?bool $is_default = null;
    public ?bool $is_trust_account = null;
    public ?bool $is_virtual = null;
    public ?array $last_transaction = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_name = null;
    public ?string $physical_account_id = null;
    public ?array $rule = null;
    public ?float $submitted_payouts_balance = null;
    public ?int $submitted_payouts_balance_minor_unit = null;
    public ?string $summary = null;
    public ?string $supplier_sepa_instant_status = null;
    public ?string $xero_bank_feed_connection_status = null;
    public ?string $xero_bank_feed_last_synced_at = null;
    public ?string $xero_bank_feed_sync_last_failed_at = null;
    public ?string $xero_bank_feed_sync_last_failure_reason = null;
    public ?string $xero_bank_feed_sync_status = null;
    public ?int $xero_unsynchronised_transactions_count = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentAccountPage#list. */
class NoFrixionMoneyMoovModelsPaymentAccountPageListMatch
{
    public ?string $account_id = null;
}

/** NoFrixionMoneyMoovModelsPaymentInitiation entity data model. */
class NoFrixionMoneyMoovModelsPaymentInitiation
{
    public ?string $payment_initiation_id = null;
    public ?string $payment_request_callback_url = null;
    public ?string $payment_request_id = null;
    public ?string $redirect_url = null;
    public ?string $response_type = null;
    public ?string $specific_error_message = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentInitiation#create. */
class NoFrixionMoneyMoovModelsPaymentInitiationCreateData
{
    public string $paymentrequest_id;
}

/** NoFrixionMoneyMoovModelsPaymentRequestEvent entity data model. */
class NoFrixionMoneyMoovModelsPaymentRequestEvent
{
    public float $amount;
    public ?string $apple_pay_transaction_id = null;
    public ?string $card_authorization_response_id = null;
    public ?int $card_expiry_month = null;
    public ?int $card_expiry_year = null;
    public ?string $card_issuer = null;
    public ?string $card_issuer_country = null;
    public ?string $card_last_four_digit = null;
    public ?string $card_request_id = null;
    public ?string $card_scheme = null;
    public ?string $card_token_customer_id = null;
    public ?string $card_transaction_id = null;
    public ?string $currency = null;
    public ?string $direct_debit_payment_id = null;
    public ?string $direct_debit_payment_reference = null;
    public ?string $drirect_debit_mandate_id = null;
    public ?string $error_message = null;
    public ?string $error_reason = null;
    public ?string $event_type = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $lightning_invoice = null;
    public ?string $lightning_r_hash = null;
    public ?string $origin_url = null;
    public ?string $payment_method_type = null;
    public ?string $payment_processor_name = null;
    public ?string $payment_request_id = null;
    public ?string $pisp_bank_status = null;
    public ?string $pisp_payment_initiation_id = null;
    public ?string $pisp_payment_institution_name = null;
    public ?string $pisp_payment_service_provider_id = null;
    public ?string $pisp_redirect_url = null;
    public ?string $reconciled_transaction_id = null;
    public ?string $refund_payout_id = null;
    public ?string $status = null;
    public ?string $wallet_name = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentRequestEvent#list. */
class NoFrixionMoneyMoovModelsPaymentRequestEventListMatch
{
    public string $paymentrequest_id;
}

/** NoFrixionMoneyMoovModelsPaymentRequestMetric entity data model. */
class NoFrixionMoneyMoovModelsPaymentRequestMetric
{
    public ?int $all = null;
    public ?int $authorized = null;
    public ?int $paid = null;
    public ?int $partially_paid = null;
    public ?array $total_amounts_by_currency = null;
    public ?int $unpaid = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentRequestMetric#load. */
class NoFrixionMoneyMoovModelsPaymentRequestMetricLoadMatch
{
    public ?int $all = null;
    public ?int $authorized = null;
    public ?int $paid = null;
    public ?int $partially_paid = null;
    public ?array $total_amounts_by_currency = null;
    public ?int $unpaid = null;
}

/** NoFrixionMoneyMoovModelsPaymentRequestMinimal entity data model. */
class NoFrixionMoneyMoovModelsPaymentRequestMinimal
{
    public ?float $amount = null;
    public ?float $amount_pending = null;
    public ?float $amount_received = null;
    public ?float $amount_refunded = null;
    public ?string $callback_url = null;
    public ?string $card_stripe_payment_intent_secret = null;
    public ?string $country_code = null;
    public ?string $currency = null;
    public ?array $custom_fields_to_display = null;
    public ?string $description = null;
    public ?string $due_date = null;
    public ?array $field_display_setting = null;
    public ?string $google_pay_merchant_id = null;
    public ?string $id = null;
    public ?string $jwk = null;
    public ?string $merchant_id = null;
    public ?string $merchant_logo_url_png = null;
    public ?string $merchant_logo_url_svg = null;
    public ?string $merchant_name = null;
    public ?string $merchant_short_name = null;
    public ?string $partial_payment_method = null;
    public ?array $payment_attempt = null;
    public ?array $payment_methods_list = null;
    public ?string $payment_processor = null;
    public ?string $payment_processor_key = null;
    public ?string $pisp_error = null;
    public ?string $priority_bank_id = null;
    public ?string $status = null;
    public ?string $stripe_account_id = null;
    public ?string $title = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentRequestMinimal#list. */
class NoFrixionMoneyMoovModelsPaymentRequestMinimalListMatch
{
    public string $paymentrequest_id;
}

/** NoFrixionMoneyMoovModelsPaymentRequestResult entity data model. */
class NoFrixionMoneyMoovModelsPaymentRequestResult
{
    public ?float $amount = null;
    public ?float $amount_pending = null;
    public ?float $amount_received = null;
    public ?float $amount_refunded = null;
    public ?string $currency = null;
    public ?string $customer_id = null;
    public ?array $payment = null;
    public ?string $payment_request_id = null;
    public ?array $pisp_authorization = null;
    public ?float $requested_amount = null;
    public ?string $result = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentRequestResult#list. */
class NoFrixionMoneyMoovModelsPaymentRequestResultListMatch
{
    public string $paymentrequest_id;
}

/** NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment entity data model. */
class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment
{
    public string $description;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $name;
    public array $template;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment#list. */
class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentListMatch
{
    public string $merchant_id;
}

/** NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2 entity data model. */
class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2
{
    public string $description;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $name;
    public array $template;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2#load. */
class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2LoadMatch
{
    public string $paymentrequest_id;
    public string $template_id;
}

/** NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3 entity data model. */
class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3
{
    public string $description;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $name;
    public array $template;
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3#update. */
class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3UpdateData
{
    public string $paymentrequest_id;
    public string $template_id;
}

/** NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4 entity data model. */
class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4
{
}

/** Request payload for NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4#remove. */
class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4RemoveMatch
{
    public string $paymentrequest_id;
    public string $template_id;
}

/** NoFrixionMoneyMoovModelsPayoutKeysetPage entity data model. */
class NoFrixionMoneyMoovModelsPayoutKeysetPage
{
    public ?string $account_id = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?string $approve_payout_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public array $beneficiary;
    public ?bool $can_authorise = null;
    public ?bool $can_process = null;
    public ?bool $can_update = null;
    public ?string $charge_bearer = null;
    public ?string $created_by = null;
    public ?string $created_by_email_address = null;
    public ?string $currency = null;
    public ?string $current_user_id = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $document = null;
    public ?array $event = null;
    public ?string $formatted_amount = null;
    public ?string $formatted_fx_destination_amount = null;
    public ?string $formatted_schedule = null;
    public ?string $formatted_schedule_day_only = null;
    public ?string $formatted_source_account_available_balance = null;
    public ?float $fx_destination_amount = null;
    public ?int $fx_destination_amount_minor_unit = null;
    public ?string $fx_destination_currency = null;
    public ?string $fx_quote_expires_at = null;
    public ?string $fx_quote_id = null;
    public ?float $fx_rate = null;
    public ?bool $fx_use_destination_amount = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoice_id = null;
    public ?bool $is_archived = null;
    public ?bool $is_failed = null;
    public ?bool $is_settled = null;
    public ?bool $is_submitted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $nonce = null;
    public ?string $payment_processor = null;
    public ?string $payment_rail = null;
    public ?string $payrun_id = null;
    public ?string $payrun_name = null;
    public ?array $rule = null;
    public ?string $schedule_date = null;
    public ?bool $scheduled = null;
    public ?float $source_account_available_balance = null;
    public ?int $source_account_available_balance_minor_unit = null;
    public ?string $source_account_bic = null;
    public ?string $source_account_currency = null;
    public ?string $source_account_iban = null;
    public array $source_account_identifier;
    public ?string $source_account_name = null;
    public ?string $source_account_number = null;
    public ?string $source_account_sortcode = null;
    public ?string $status = null;
    public ?array $tag = null;
    public ?string $their_reference = null;
    public ?string $topup_payrun_id = null;
    public ?float $transacted_amount = null;
    public ?float $transacted_fx_amount = null;
    public ?float $transacted_fx_rate = null;
    public ?string $type = null;
    public ?string $user_id = null;
    public ?string $your_reference = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPayoutKeysetPage#list. */
class NoFrixionMoneyMoovModelsPayoutKeysetPageListMatch
{
    public ?string $account_id = null;
    public ?string $merchant_id = null;
}

/** NoFrixionMoneyMoovModelsPayoutMetric entity data model. */
class NoFrixionMoneyMoovModelsPayoutMetric
{
    public ?float $all = null;
    public ?float $failed = null;
    public ?float $in_progress = null;
    public ?float $paid = null;
    public ?float $pending_approval = null;
    public ?float $scheduled = null;
    public ?array $total_amounts_by_currency = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPayoutMetric#load. */
class NoFrixionMoneyMoovModelsPayoutMetricLoadMatch
{
    public ?float $all = null;
    public ?float $failed = null;
    public ?float $in_progress = null;
    public ?float $paid = null;
    public ?float $pending_approval = null;
    public ?float $scheduled = null;
    public ?array $total_amounts_by_currency = null;
}

/** NoFrixionMoneyMoovModelsPayoutsPayoutsCreate entity data model. */
class NoFrixionMoneyMoovModelsPayoutsPayoutsCreate
{
    public ?array $failed_payout = null;
    public ?array $payout = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPayoutsPayoutsCreate#create. */
class NoFrixionMoneyMoovModelsPayoutsPayoutsCreateCreateData
{
    public ?array $failed_payout = null;
    public ?array $payout = null;
}

/** NoFrixionMoneyMoovModelsPayrun entity data model. */
class NoFrixionMoneyMoovModelsPayrun
{
    public ?array $authorisation = null;
    public ?string $authorisation_date = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public ?bool $can_authorise = null;
    public ?bool $can_delete = null;
    public ?bool $can_edit = null;
    public ?array $event = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?array $invoice = null;
    public ?array $invoices_minimal = null;
    public ?bool $is_archived = null;
    public ?string $last_updated = null;
    public array $last_updated_by;
    public ?string $merchant_id = null;
    public ?string $name = null;
    public ?string $nonce = null;
    public ?array $payment = null;
    public ?array $payout = null;
    public ?int $payouts_count = null;
    public ?string $reason = null;
    public ?string $schedule_date = null;
    public ?string $scheduled_date = null;
    public ?array $source_account = null;
    public ?string $status = null;
    public ?float $total_eur = null;
    public ?float $total_gbp = null;
    public ?float $total_usd = null;
}

/** Request payload for NoFrixionMoneyMoovModelsPayrun#load. */
class NoFrixionMoneyMoovModelsPayrunLoadMatch
{
    public string $id;
}

/** Request payload for NoFrixionMoneyMoovModelsPayrun#create. */
class NoFrixionMoneyMoovModelsPayrunCreateData
{
    public string $id;
}

/** Request payload for NoFrixionMoneyMoovModelsPayrun#update. */
class NoFrixionMoneyMoovModelsPayrunUpdateData
{
    public ?string $id = null;
    public ?string $payrun_id = null;
}

/** NoFrixionMoneyMoovModelsReportResult entity data model. */
class NoFrixionMoneyMoovModelsReportResult
{
    public ?string $content = null;
    public ?string $content_type = null;
    public ?string $last_completed_at = null;
    public ?string $merchant_id = null;
    public ?string $report_name = null;
    public ?string $report_type = null;
    public ?int $statement_number = null;
}

/** Request payload for NoFrixionMoneyMoovModelsReportResult#load. */
class NoFrixionMoneyMoovModelsReportResultLoadMatch
{
    public int $id;
    public string $report_id;
}

/** NoFrixionMoneyMoovModelsRule entity data model. */
class NoFrixionMoneyMoovModelsRule
{
    public ?array $account = null;
    public ?string $account_id = null;
    public ?string $approve_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?bool $can_authorise = null;
    public array $created_by;
    public ?string $description = null;
    public ?string $end_at = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_disabled = null;
    public ?string $last_executed_at = null;
    public ?string $last_run_at_transaction_date = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $name = null;
    public string $nonce;
    public ?string $on_approved_web_hook_url = null;
    public ?string $on_execution_error_web_hook_url = null;
    public ?string $on_execution_success_web_hook_url = null;
    public ?string $start_at = null;
    public ?string $status = null;
    public ?array $sweep_action = null;
    public ?string $time_zone_id = null;
    public ?string $trigger_cron_expression = null;
    public ?bool $trigger_on_pay_in = null;
    public ?string $user_id = null;
    public ?string $web_hook_secret = null;
}

/** Request payload for NoFrixionMoneyMoovModelsRule#load. */
class NoFrixionMoneyMoovModelsRuleLoadMatch
{
    public string $id;
}

/** Request payload for NoFrixionMoneyMoovModelsRule#create. */
class NoFrixionMoneyMoovModelsRuleCreateData
{
    public ?array $account = null;
    public ?string $account_id = null;
    public ?string $approve_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?bool $can_authorise = null;
    public array $created_by;
    public ?string $description = null;
    public ?string $end_at = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?bool $is_disabled = null;
    public ?string $last_executed_at = null;
    public ?string $last_run_at_transaction_date = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $name = null;
    public string $nonce;
    public ?string $on_approved_web_hook_url = null;
    public ?string $on_execution_error_web_hook_url = null;
    public ?string $on_execution_success_web_hook_url = null;
    public ?string $start_at = null;
    public ?string $status = null;
    public ?array $sweep_action = null;
    public ?string $time_zone_id = null;
    public ?string $trigger_cron_expression = null;
    public ?bool $trigger_on_pay_in = null;
    public ?string $user_id = null;
    public ?string $web_hook_secret = null;
}

/** Request payload for NoFrixionMoneyMoovModelsRule#update. */
class NoFrixionMoneyMoovModelsRuleUpdateData
{
    public string $id;
}

/** NoFrixionMoneyMoovModelsTransaction entity data model. */
class NoFrixionMoneyMoovModelsTransaction
{
    public ?string $account_id = null;
    public ?string $account_name = null;
    public ?int $account_sequence_number = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?float $balance = null;
    public ?int $balance_minor_unit = null;
    public ?array $counterparty = null;
    public ?string $counterparty_summary = null;
    public ?string $currency = null;
    public ?string $description = null;
    public ?float $fx_amount = null;
    public ?string $fx_currency = null;
    public ?float $fx_rate = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $merchant_id = null;
    public ?array $payment_request_custom_field = null;
    public ?string $payment_request_id = null;
    public ?string $payout_id = null;
    public ?string $raw_reference = null;
    public ?string $rule_id = null;
    public ?array $tag = null;
    public ?string $their_reference = null;
    public ?string $transaction_date = null;
    public ?string $type = null;
    public ?string $virtual_iban = null;
    public ?string $your_reference = null;
}

/** Request payload for NoFrixionMoneyMoovModelsTransaction#load. */
class NoFrixionMoneyMoovModelsTransactionLoadMatch
{
    public ?string $account_id = null;
    public string $id;
}

/** NoFrixionMoneyMoovModelsTransactionPage entity data model. */
class NoFrixionMoneyMoovModelsTransactionPage
{
    public ?string $account_id = null;
    public ?string $account_name = null;
    public ?int $account_sequence_number = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?float $balance = null;
    public ?int $balance_minor_unit = null;
    public ?array $content = null;
    public ?array $counterparty = null;
    public ?string $counterparty_summary = null;
    public ?string $currency = null;
    public ?string $description = null;
    public ?float $fx_amount = null;
    public ?string $fx_currency = null;
    public ?float $fx_rate = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $merchant_id = null;
    public ?int $page_number = null;
    public ?int $page_size = null;
    public ?array $payment_request_custom_field = null;
    public ?string $payment_request_id = null;
    public ?string $payout_id = null;
    public ?string $raw_reference = null;
    public ?string $rule_id = null;
    public ?array $tag = null;
    public ?string $their_reference = null;
    public ?int $total_page = null;
    public ?int $total_size = null;
    public ?string $transaction_date = null;
    public ?string $type = null;
    public ?string $virtual_iban = null;
    public ?string $your_reference = null;
}

/** Request payload for NoFrixionMoneyMoovModelsTransactionPage#load. */
class NoFrixionMoneyMoovModelsTransactionPageLoadMatch
{
    public string $account_id;
}

/** Request payload for NoFrixionMoneyMoovModelsTransactionPage#list. */
class NoFrixionMoneyMoovModelsTransactionPageListMatch
{
    public ?string $account_id = null;
    public ?string $merchant_id = null;
}

/** NoFrixionMoneyMoovModelsUserInvite entity data model. */
class NoFrixionMoneyMoovModelsUserInvite
{
    public ?array $authorisation_status = null;
    public ?string $id = null;
    public ?string $initial_role_id = null;
    public ?string $invitee_email_address = null;
    public ?string $invitee_first_name = null;
    public ?string $invitee_last_name = null;
    public ?string $inviter_email_address = null;
    public ?string $inviter_first_name = null;
    public ?string $inviter_last_name = null;
    public ?bool $is_authorised = null;
    public ?bool $is_invitee_registered = null;
    public ?string $last_invited = null;
    public ?string $merchant_id = null;
    public ?string $merchant_name = null;
    public ?string $message = null;
    public ?string $registration_url = null;
    public ?bool $send_invite_email = null;
    public ?string $status = null;
    public array $user;
    public ?string $user_id = null;
}

/** Request payload for NoFrixionMoneyMoovModelsUserInvite#load. */
class NoFrixionMoneyMoovModelsUserInviteLoadMatch
{
    public ?string $id = null;
    public ?string $userinvite_id = null;
}

/** Request payload for NoFrixionMoneyMoovModelsUserInvite#create. */
class NoFrixionMoneyMoovModelsUserInviteCreateData
{
    public ?array $authorisation_status = null;
    public ?string $id = null;
    public ?string $initial_role_id = null;
    public ?string $invitee_email_address = null;
    public ?string $invitee_first_name = null;
    public ?string $invitee_last_name = null;
    public ?string $inviter_email_address = null;
    public ?string $inviter_first_name = null;
    public ?string $inviter_last_name = null;
    public ?bool $is_authorised = null;
    public ?bool $is_invitee_registered = null;
    public ?string $last_invited = null;
    public ?string $merchant_id = null;
    public ?string $merchant_name = null;
    public ?string $message = null;
    public ?string $registration_url = null;
    public ?bool $send_invite_email = null;
    public ?string $status = null;
    public array $user;
    public ?string $user_id = null;
}

/** NoFrixionMoneyMoovModelsUserInvitePage entity data model. */
class NoFrixionMoneyMoovModelsUserInvitePage
{
    public ?array $authorisation_status = null;
    public ?string $id = null;
    public ?string $initial_role_id = null;
    public ?string $invitee_email_address = null;
    public ?string $invitee_first_name = null;
    public ?string $invitee_last_name = null;
    public ?string $inviter_email_address = null;
    public ?string $inviter_first_name = null;
    public ?string $inviter_last_name = null;
    public ?bool $is_authorised = null;
    public ?bool $is_invitee_registered = null;
    public ?string $last_invited = null;
    public ?string $merchant_id = null;
    public ?string $merchant_name = null;
    public ?string $message = null;
    public ?string $registration_url = null;
    public ?string $status = null;
    public array $user;
    public ?string $user_id = null;
}

/** Request payload for NoFrixionMoneyMoovModelsUserInvitePage#list. */
class NoFrixionMoneyMoovModelsUserInvitePageListMatch
{
    public string $merchant_id;
}

/** NoFrixionMoneyMoovModelsUserPage entity data model. */
class NoFrixionMoneyMoovModelsUserPage
{
    public ?array $client_session_timeout = null;
    public string $email_address;
    public string $first_name;
    public ?string $id = null;
    public string $last_name;
    public ?bool $passkey_added = null;
    public ?array $permission = null;
    public ?array $roles_with_scope = null;
    public ?bool $two_factor_enabled = null;
}

/** Request payload for NoFrixionMoneyMoovModelsUserPage#list. */
class NoFrixionMoneyMoovModelsUserPageListMatch
{
    public string $merchant_id;
}

/** NoFrixionMoneyMoovModelsWebhook entity data model. */
class NoFrixionMoneyMoovModelsWebhook
{
    public ?string $destination_url = null;
    public ?string $email_address = null;
    public ?string $failed_notification_email_address = null;
    public ?string $id = null;
    public ?bool $is_active = null;
    public ?string $merchant_id = null;
    public ?string $notification_method = null;
    public ?array $resource_type = null;
    public ?bool $retry = null;
    public ?string $secret = null;
    public ?int $version = null;
}

/** Request payload for NoFrixionMoneyMoovModelsWebhook#load. */
class NoFrixionMoneyMoovModelsWebhookLoadMatch
{
    public string $id;
    public ?string $merchant_id = null;
}

/** Request payload for NoFrixionMoneyMoovModelsWebhook#list. */
class NoFrixionMoneyMoovModelsWebhookListMatch
{
    public string $merchant_id;
}

/** Request payload for NoFrixionMoneyMoovModelsWebhook#create. */
class NoFrixionMoneyMoovModelsWebhookCreateData
{
    public ?string $destination_url = null;
    public ?string $email_address = null;
    public ?string $failed_notification_email_address = null;
    public ?string $id = null;
    public ?bool $is_active = null;
    public ?string $merchant_id = null;
    public ?string $notification_method = null;
    public ?array $resource_type = null;
    public ?bool $retry = null;
    public ?string $secret = null;
    public ?int $version = null;
}

/** Request payload for NoFrixionMoneyMoovModelsWebhook#update. */
class NoFrixionMoneyMoovModelsWebhookUpdateData
{
    public string $id;
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
    public string $account_name;
    public ?string $account_number = null;
    public string $iban;
    public ?string $payee_verified_account_name = null;
    public ?string $result = null;
    public ?string $secondary_identification = null;
    public ?string $sort_code = null;
}

/** Request payload for Payeeverification#create. */
class PayeeverificationCreateData
{
    public string $account_name;
    public ?string $account_number = null;
    public string $iban;
    public ?string $payee_verified_account_name = null;
    public ?string $result = null;
    public ?string $secondary_identification = null;
    public ?string $sort_code = null;
}

/** PaymentRequest entity data model. */
class PaymentRequest
{
    public ?float $amount = null;
    public ?bool $do_simulate_settlement_failure = null;
    public ?string $error_description = null;
    public ?string $institution = null;
    public ?string $payment_initiation_id = null;
}

/** Request payload for PaymentRequest#load. */
class PaymentRequestLoadMatch
{
    public ?string $paymentrequest_id = null;
}

/** Request payload for PaymentRequest#create. */
class PaymentRequestCreateData
{
    public string $paymentrequest_id;
}

/** Request payload for PaymentRequest#update. */
class PaymentRequestUpdateData
{
    public string $paymentrequest_id;
}

/** Request payload for PaymentRequest#remove. */
class PaymentRequestRemoveMatch
{
    public string $id;
}

/** Payout entity data model. */
class Payout
{
    public ?string $account_id = null;
    public ?bool $allow_incomplete = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?string $approve_payout_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public array $beneficiary;
    public ?string $beneficiary_id = null;
    public ?bool $can_authorise = null;
    public ?bool $can_process = null;
    public ?bool $can_update = null;
    public ?string $charge_bearer = null;
    public ?string $created_by = null;
    public ?string $created_by_email_address = null;
    public ?string $currency = null;
    public ?string $current_user_id = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $document = null;
    public ?array $event = null;
    public ?string $formatted_amount = null;
    public ?string $formatted_fx_destination_amount = null;
    public ?string $formatted_schedule = null;
    public ?string $formatted_schedule_day_only = null;
    public ?string $formatted_source_account_available_balance = null;
    public ?float $fx_destination_amount = null;
    public ?int $fx_destination_amount_minor_unit = null;
    public ?string $fx_destination_currency = null;
    public ?string $fx_quote_expires_at = null;
    public ?string $fx_quote_id = null;
    public ?float $fx_rate = null;
    public ?bool $fx_use_destination_amount = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoice_id = null;
    public ?bool $is_archived = null;
    public ?bool $is_failed = null;
    public ?bool $is_settled = null;
    public ?bool $is_submitted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $nonce = null;
    public ?string $payment_processor = null;
    public ?string $payment_rail = null;
    public ?string $payrun_id = null;
    public ?string $payrun_name = null;
    public ?array $rule = null;
    public ?string $schedule_date = null;
    public ?bool $scheduled = null;
    public ?float $source_account_available_balance = null;
    public ?int $source_account_available_balance_minor_unit = null;
    public ?string $source_account_bic = null;
    public ?string $source_account_currency = null;
    public ?string $source_account_iban = null;
    public array $source_account_identifier;
    public ?string $source_account_name = null;
    public ?string $source_account_number = null;
    public ?string $source_account_sortcode = null;
    public ?string $status = null;
    public ?array $tag = null;
    public ?array $tag_id = null;
    public ?string $their_reference = null;
    public ?string $topup_payrun_id = null;
    public ?float $transacted_amount = null;
    public ?float $transacted_fx_amount = null;
    public ?float $transacted_fx_rate = null;
    public ?string $type = null;
    public ?string $user_id = null;
    public ?string $your_reference = null;
}

/** Request payload for Payout#load. */
class PayoutLoadMatch
{
    public ?float $amount = null;
    public ?string $destination = null;
    public ?string $source = null;
    public ?string $id = null;
}

/** Request payload for Payout#create. */
class PayoutCreateData
{
    public ?string $id = null;
}

/** Request payload for Payout#update. */
class PayoutUpdateData
{
    public string $id;
}

/** Request payload for Payout#remove. */
class PayoutRemoveMatch
{
    public string $id;
}

/** Payrun entity data model. */
class Payrun
{
    public ?string $id = null;
    public ?string $note = null;
    public ?string $scheduled_date = null;
}

/** Request payload for Payrun#create. */
class PayrunCreateData
{
    public string $id;
}

/** Request payload for Payrun#update. */
class PayrunUpdateData
{
    public string $id;
}

/** Request payload for Payrun#remove. */
class PayrunRemoveMatch
{
    public string $id;
}

/** Reject entity data model. */
class Reject
{
    public ?string $account_id = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?string $approve_payout_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public array $beneficiary;
    public ?bool $can_authorise = null;
    public ?bool $can_process = null;
    public ?bool $can_update = null;
    public ?string $charge_bearer = null;
    public ?string $created_by = null;
    public ?string $created_by_email_address = null;
    public ?string $currency = null;
    public ?string $current_user_id = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $document = null;
    public ?array $event = null;
    public ?string $formatted_amount = null;
    public ?string $formatted_fx_destination_amount = null;
    public ?string $formatted_schedule = null;
    public ?string $formatted_schedule_day_only = null;
    public ?string $formatted_source_account_available_balance = null;
    public ?float $fx_destination_amount = null;
    public ?int $fx_destination_amount_minor_unit = null;
    public ?string $fx_destination_currency = null;
    public ?string $fx_quote_expires_at = null;
    public ?string $fx_quote_id = null;
    public ?float $fx_rate = null;
    public ?bool $fx_use_destination_amount = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoice_id = null;
    public ?bool $is_archived = null;
    public ?bool $is_failed = null;
    public ?bool $is_settled = null;
    public ?bool $is_submitted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $nonce = null;
    public ?string $payment_processor = null;
    public ?string $payment_rail = null;
    public ?string $payrun_id = null;
    public ?string $payrun_name = null;
    public ?string $reason = null;
    public ?array $rule = null;
    public ?string $schedule_date = null;
    public ?bool $scheduled = null;
    public ?float $source_account_available_balance = null;
    public ?int $source_account_available_balance_minor_unit = null;
    public ?string $source_account_bic = null;
    public ?string $source_account_currency = null;
    public ?string $source_account_iban = null;
    public array $source_account_identifier;
    public ?string $source_account_name = null;
    public ?string $source_account_number = null;
    public ?string $source_account_sortcode = null;
    public ?string $status = null;
    public ?array $tag = null;
    public ?string $their_reference = null;
    public ?string $topup_payrun_id = null;
    public ?float $transacted_amount = null;
    public ?float $transacted_fx_amount = null;
    public ?float $transacted_fx_rate = null;
    public ?string $type = null;
    public ?string $user_id = null;
    public ?string $your_reference = null;
}

/** Request payload for Reject#update. */
class RejectUpdateData
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

/** Rule entity data model. */
class Rule
{
}

/** Request payload for Rule#update. */
class RuleUpdateData
{
    public string $id;
}

/** Request payload for Rule#remove. */
class RuleRemoveMatch
{
    public string $id;
}

/** Send entity data model. */
class Send
{
    public ?string $account_id = null;
    public ?bool $allow_incomplete = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?string $approve_payout_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public array $beneficiary;
    public ?string $beneficiary_id = null;
    public ?bool $can_authorise = null;
    public ?bool $can_process = null;
    public ?bool $can_update = null;
    public ?string $charge_bearer = null;
    public ?string $created_by = null;
    public ?string $created_by_email_address = null;
    public ?string $currency = null;
    public ?string $current_user_id = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $document = null;
    public ?array $event = null;
    public ?string $formatted_amount = null;
    public ?string $formatted_fx_destination_amount = null;
    public ?string $formatted_schedule = null;
    public ?string $formatted_schedule_day_only = null;
    public ?string $formatted_source_account_available_balance = null;
    public ?float $fx_destination_amount = null;
    public ?int $fx_destination_amount_minor_unit = null;
    public ?string $fx_destination_currency = null;
    public ?string $fx_quote_expires_at = null;
    public ?string $fx_quote_id = null;
    public ?float $fx_rate = null;
    public ?bool $fx_use_destination_amount = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoice_id = null;
    public ?bool $is_archived = null;
    public ?bool $is_failed = null;
    public ?bool $is_settled = null;
    public ?bool $is_submitted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $nonce = null;
    public ?string $payment_processor = null;
    public ?string $payment_rail = null;
    public ?string $payrun_id = null;
    public ?string $payrun_name = null;
    public ?array $rule = null;
    public ?string $schedule_date = null;
    public ?bool $scheduled = null;
    public ?float $source_account_available_balance = null;
    public ?int $source_account_available_balance_minor_unit = null;
    public ?string $source_account_bic = null;
    public ?string $source_account_currency = null;
    public ?string $source_account_iban = null;
    public array $source_account_identifier;
    public ?string $source_account_name = null;
    public ?string $source_account_number = null;
    public ?string $source_account_sortcode = null;
    public ?string $status = null;
    public ?array $tag = null;
    public ?array $tag_id = null;
    public ?string $their_reference = null;
    public ?string $topup_payrun_id = null;
    public ?float $transacted_amount = null;
    public ?float $transacted_fx_amount = null;
    public ?float $transacted_fx_rate = null;
    public ?string $type = null;
    public ?string $user_id = null;
    public ?string $your_reference = null;
}

/** Request payload for Send#create. */
class SendCreateData
{
    public ?string $account_id = null;
    public ?bool $allow_incomplete = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?string $approve_payout_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public array $beneficiary;
    public ?string $beneficiary_id = null;
    public ?bool $can_authorise = null;
    public ?bool $can_process = null;
    public ?bool $can_update = null;
    public ?string $charge_bearer = null;
    public ?string $created_by = null;
    public ?string $created_by_email_address = null;
    public ?string $currency = null;
    public ?string $current_user_id = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $document = null;
    public ?array $event = null;
    public ?string $formatted_amount = null;
    public ?string $formatted_fx_destination_amount = null;
    public ?string $formatted_schedule = null;
    public ?string $formatted_schedule_day_only = null;
    public ?string $formatted_source_account_available_balance = null;
    public ?float $fx_destination_amount = null;
    public ?int $fx_destination_amount_minor_unit = null;
    public ?string $fx_destination_currency = null;
    public ?string $fx_quote_expires_at = null;
    public ?string $fx_quote_id = null;
    public ?float $fx_rate = null;
    public ?bool $fx_use_destination_amount = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoice_id = null;
    public ?bool $is_archived = null;
    public ?bool $is_failed = null;
    public ?bool $is_settled = null;
    public ?bool $is_submitted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $nonce = null;
    public ?string $payment_processor = null;
    public ?string $payment_rail = null;
    public ?string $payrun_id = null;
    public ?string $payrun_name = null;
    public ?array $rule = null;
    public ?string $schedule_date = null;
    public ?bool $scheduled = null;
    public ?float $source_account_available_balance = null;
    public ?int $source_account_available_balance_minor_unit = null;
    public ?string $source_account_bic = null;
    public ?string $source_account_currency = null;
    public ?string $source_account_iban = null;
    public array $source_account_identifier;
    public ?string $source_account_name = null;
    public ?string $source_account_number = null;
    public ?string $source_account_sortcode = null;
    public ?string $status = null;
    public ?array $tag = null;
    public ?array $tag_id = null;
    public ?string $their_reference = null;
    public ?string $topup_payrun_id = null;
    public ?float $transacted_amount = null;
    public ?float $transacted_fx_amount = null;
    public ?float $transacted_fx_rate = null;
    public ?string $type = null;
    public ?string $user_id = null;
    public ?string $your_reference = null;
}

/** Sendbeneficiary entity data model. */
class Sendbeneficiary
{
    public ?string $account_id = null;
    public ?bool $allow_incomplete = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?string $approve_payout_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public array $beneficiary;
    public ?string $beneficiary_id = null;
    public ?bool $can_authorise = null;
    public ?bool $can_process = null;
    public ?bool $can_update = null;
    public ?string $charge_bearer = null;
    public ?string $created_by = null;
    public ?string $created_by_email_address = null;
    public ?string $currency = null;
    public ?string $current_user_id = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $document = null;
    public ?array $event = null;
    public ?string $formatted_amount = null;
    public ?string $formatted_fx_destination_amount = null;
    public ?string $formatted_schedule = null;
    public ?string $formatted_schedule_day_only = null;
    public ?string $formatted_source_account_available_balance = null;
    public ?float $fx_destination_amount = null;
    public ?int $fx_destination_amount_minor_unit = null;
    public ?string $fx_destination_currency = null;
    public ?string $fx_quote_expires_at = null;
    public ?string $fx_quote_id = null;
    public ?float $fx_rate = null;
    public ?bool $fx_use_destination_amount = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoice_id = null;
    public ?bool $is_archived = null;
    public ?bool $is_failed = null;
    public ?bool $is_settled = null;
    public ?bool $is_submitted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $nonce = null;
    public ?string $payment_processor = null;
    public ?string $payment_rail = null;
    public ?string $payrun_id = null;
    public ?string $payrun_name = null;
    public ?array $rule = null;
    public ?string $schedule_date = null;
    public ?bool $scheduled = null;
    public ?float $source_account_available_balance = null;
    public ?int $source_account_available_balance_minor_unit = null;
    public ?string $source_account_bic = null;
    public ?string $source_account_currency = null;
    public ?string $source_account_iban = null;
    public array $source_account_identifier;
    public ?string $source_account_name = null;
    public ?string $source_account_number = null;
    public ?string $source_account_sortcode = null;
    public ?string $status = null;
    public ?array $tag = null;
    public ?array $tag_id = null;
    public ?string $their_reference = null;
    public ?string $topup_payrun_id = null;
    public ?float $transacted_amount = null;
    public ?float $transacted_fx_amount = null;
    public ?float $transacted_fx_rate = null;
    public ?string $type = null;
    public ?string $user_id = null;
    public ?string $your_reference = null;
}

/** Request payload for Sendbeneficiary#create. */
class SendbeneficiaryCreateData
{
    public ?string $account_id = null;
    public ?bool $allow_incomplete = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?string $approve_payout_url = null;
    public ?string $approver_id = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?string $batch_payout_id = null;
    public array $beneficiary;
    public ?string $beneficiary_id = null;
    public ?bool $can_authorise = null;
    public ?bool $can_process = null;
    public ?bool $can_update = null;
    public ?string $charge_bearer = null;
    public ?string $created_by = null;
    public ?string $created_by_email_address = null;
    public ?string $currency = null;
    public ?string $current_user_id = null;
    public ?string $description = null;
    public ?array $destination = null;
    public ?array $document = null;
    public ?array $event = null;
    public ?string $formatted_amount = null;
    public ?string $formatted_fx_destination_amount = null;
    public ?string $formatted_schedule = null;
    public ?string $formatted_schedule_day_only = null;
    public ?string $formatted_source_account_available_balance = null;
    public ?float $fx_destination_amount = null;
    public ?int $fx_destination_amount_minor_unit = null;
    public ?string $fx_destination_currency = null;
    public ?string $fx_quote_expires_at = null;
    public ?string $fx_quote_id = null;
    public ?float $fx_rate = null;
    public ?bool $fx_use_destination_amount = null;
    public ?bool $has_current_user_authorised = null;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $invoice_id = null;
    public ?bool $is_archived = null;
    public ?bool $is_failed = null;
    public ?bool $is_settled = null;
    public ?bool $is_submitted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_token_description = null;
    public ?string $nonce = null;
    public ?string $payment_processor = null;
    public ?string $payment_rail = null;
    public ?string $payrun_id = null;
    public ?string $payrun_name = null;
    public ?array $rule = null;
    public ?string $schedule_date = null;
    public ?bool $scheduled = null;
    public ?float $source_account_available_balance = null;
    public ?int $source_account_available_balance_minor_unit = null;
    public ?string $source_account_bic = null;
    public ?string $source_account_currency = null;
    public ?string $source_account_iban = null;
    public array $source_account_identifier;
    public ?string $source_account_name = null;
    public ?string $source_account_number = null;
    public ?string $source_account_sortcode = null;
    public ?string $status = null;
    public ?array $tag = null;
    public ?array $tag_id = null;
    public ?string $their_reference = null;
    public ?string $topup_payrun_id = null;
    public ?float $transacted_amount = null;
    public ?float $transacted_fx_amount = null;
    public ?float $transacted_fx_rate = null;
    public ?string $type = null;
    public ?string $user_id = null;
    public ?string $your_reference = null;
}

/** Tag entity data model. */
class Tag
{
    public ?string $colour_hex = null;
    public ?string $description = null;
    public ?string $id = null;
    public string $merchant_id;
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
}

/** Request payload for Transaction#load. */
class TransactionLoadMatch
{
    public int $sequence_number;
    public string $transaction_id;
}

/** Request payload for Transaction#create. */
class TransactionCreateData
{
    public string $id;
}

/** Request payload for Transaction#remove. */
class TransactionRemoveMatch
{
    public string $id;
}

/** User entity data model. */
class User
{
    public ?array $client_session_timeout = null;
    public string $email_address;
    public string $first_name;
    public ?string $id = null;
    public string $last_name;
    public ?bool $passkey_added = null;
    public ?array $permission = null;
    public ?string $profile = null;
    public ?array $roles_with_scope = null;
    public ?bool $two_factor_enabled = null;
    public ?string $user_invite_id = null;
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
}

/** UserInvite entity data model. */
class UserInvite
{
}

/** Request payload for UserInvite#create. */
class UserInviteCreateData
{
    public string $id;
}

/** Request payload for UserInvite#update. */
class UserInviteUpdateData
{
    public string $id;
}

/** Request payload for UserInvite#remove. */
class UserInviteRemoveMatch
{
    public string $id;
}

/** Virtual entity data model. */
class Virtual
{
    public ?string $account_name = null;
    public ?string $account_supplier_name = null;
    public ?float $available_balance = null;
    public ?int $available_balance_minor_unit = null;
    public ?float $balance = null;
    public ?int $balance_minor_unit = null;
    public ?string $bank_name = null;
    public ?string $consent_id = null;
    public array $created_by;
    public ?string $created_by_display_name = null;
    public ?string $currency = null;
    public ?string $default_payment_rail = null;
    public ?string $display_name = null;
    public ?string $expiry_date = null;
    public ?string $external_account_icon = null;
    public ?string $id = null;
    public array $identifier;
    public ?string $inserted = null;
    public ?bool $is_archived = null;
    public ?bool $is_connected_account = null;
    public ?bool $is_default = null;
    public ?bool $is_trust_account = null;
    public ?bool $is_virtual = null;
    public ?array $last_transaction = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public ?string $merchant_name = null;
    public string $name;
    public ?string $physical_account_id = null;
    public ?array $rule = null;
    public ?float $submitted_payouts_balance = null;
    public ?int $submitted_payouts_balance_minor_unit = null;
    public ?string $summary = null;
    public ?string $supplier_sepa_instant_status = null;
    public ?string $xero_bank_feed_connection_status = null;
    public ?string $xero_bank_feed_last_synced_at = null;
    public ?string $xero_bank_feed_sync_last_failed_at = null;
    public ?string $xero_bank_feed_sync_last_failure_reason = null;
    public ?string $xero_bank_feed_sync_status = null;
    public ?int $xero_unsynchronised_transactions_count = null;
}

/** Request payload for Virtual#create. */
class VirtualCreateData
{
    public string $account_id;
}

/** Request payload for Virtual#update. */
class VirtualUpdateData
{
    public string $account_id;
    public string $id;
}

/** Webhook entity data model. */
class Webhook
{
}

/** Request payload for Webhook#remove. */
class WebhookRemoveMatch
{
    public string $id;
}

/** Whoami entity data model. */
class Whoami
{
    public ?array $client_session_timeout = null;
    public string $email_address;
    public string $first_name;
    public ?string $id = null;
    public string $last_name;
    public ?bool $passkey_added = null;
    public ?array $permission = null;
    public ?array $roles_with_scope = null;
    public ?bool $two_factor_enabled = null;
}

/** Request payload for Whoami#list. */
class WhoamiListMatch
{
    public ?array $client_session_timeout = null;
    public ?string $email_address = null;
    public ?string $first_name = null;
    public ?string $id = null;
    public ?string $last_name = null;
    public ?bool $passkey_added = null;
    public ?array $permission = null;
    public ?array $roles_with_scope = null;
    public ?bool $two_factor_enabled = null;
}

/** Whoamitrustedapp entity data model. */
class Whoamitrustedapp
{
    public ?array $client_session_timeout = null;
    public string $email_address;
    public string $first_name;
    public ?string $id = null;
    public string $last_name;
    public ?bool $passkey_added = null;
    public ?array $permission = null;
    public ?array $roles_with_scope = null;
    public ?bool $two_factor_enabled = null;
}

/** Request payload for Whoamitrustedapp#list. */
class WhoamitrustedappListMatch
{
    public ?array $client_session_timeout = null;
    public ?string $email_address = null;
    public ?string $first_name = null;
    public ?string $id = null;
    public ?string $last_name = null;
    public ?bool $passkey_added = null;
    public ?array $permission = null;
    public ?array $roles_with_scope = null;
    public ?bool $two_factor_enabled = null;
}

