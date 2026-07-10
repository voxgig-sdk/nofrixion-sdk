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
    public ?array $account_balance = null;
    public ?string $account_id = null;
    public ?array $account_identification = null;
    public ?string $account_name = null;
    public ?string $account_supplier_name = null;
    public ?string $account_type = null;
    public ?float $available_balance = null;
    public ?int $available_balance_minor_unit = null;
    public ?float $balance = null;
    public ?int $balance_minor_unit = null;
    public ?string $bank_name = null;
    public ?string $consent_id = null;
    public ?array $consolidated_account_information = null;
    public array $created_by;
    public ?string $created_by_display_name = null;
    public ?string $currency = null;
    public ?string $default_payment_rail = null;
    public ?string $description = null;
    public ?string $detail = null;
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
    public ?string $nickname = null;
    public ?string $physical_account_id = null;
    public ?array $role_i_d = null;
    public ?array $rule = null;
    public ?float $submitted_payouts_balance = null;
    public ?int $submitted_payouts_balance_minor_unit = null;
    public ?string $summary = null;
    public ?string $supplier_physical_account_id = null;
    public ?string $supplier_sepa_instant_status = null;
    public ?string $to_date = null;
    public ?string $type = null;
    public ?string $usage_type = null;
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

/** Batch entity data model. */
class Batch
{
    public ?string $approve_url = null;
    public ?string $id = null;
    public ?array $payout = null;
}

/** Request payload for Batch#load. */
class BatchLoadMatch
{
    public string $id;
}

/** Request payload for Batch#create. */
class BatchCreateData
{
    public ?string $approve_url = null;
    public ?string $id = null;
    public ?array $payout = null;
}

/** Beneficiary entity data model. */
class Beneficiary
{
    public ?string $approval_callback_url = null;
    public ?array $authentication_method = null;
    public ?array $authorisation = null;
    public ?int $authorisers_completed_count = null;
    public ?int $authorisers_required_count = null;
    public ?array $beneficiary = null;
    public ?array $beneficiary_event = null;
    public ?bool $can_authorise = null;
    public ?bool $can_update = null;
    public array $created_by;
    public ?string $created_by_email_address = null;
    public string $currency;
    public ?array $destination = null;
    public ?array $failed_beneficiary = null;
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

/** Request payload for Beneficiary#list. */
class BeneficiaryListMatch
{
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

/** BeneficiaryGroup entity data model. */
class BeneficiaryGroup
{
    public ?array $group_member = null;
    public string $group_name;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $last_updated = null;
    public string $merchant_id;
}

/** Request payload for BeneficiaryGroup#list. */
class BeneficiaryGroupListMatch
{
    public string $merchant_id;
}

/** Card entity data model. */
class Card
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

/** Request payload for Card#create. */
class CardCreateData
{
    public string $paymentrequest_id;
}

/** CardCustomerToken entity data model. */
class CardCustomerToken
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

/** Request payload for CardPayment#create. */
class CardPaymentCreateData
{
    public ?float $partial_refund_amount = null;
    public string $paymentrequest_id;
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

/** Request payload for Consent#update. */
class ConsentUpdateData
{
    public string $id;
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
    public ?int $decimal = null;
    public ?bool $is_fiat = null;
    public ?string $iso4217_alpha_code = null;
    public ?string $iso4217_numeric_code = null;
    public ?string $symbol = null;
}

/** Request payload for Currency#list. */
class CurrencyListMatch
{
    public ?string $code = null;
    public ?int $decimal = null;
    public ?bool $is_fiat = null;
    public ?string $iso4217_alpha_code = null;
    public ?string $iso4217_numeric_code = null;
    public ?string $symbol = null;
}

/** DirectDebitBatchSubmit entity data model. */
class DirectDebitBatchSubmit
{
    public ?array $failed_submission = null;
    public ?array $successful_submission = null;
}

/** Request payload for DirectDebitBatchSubmit#create. */
class DirectDebitBatchSubmitCreateData
{
    public ?array $failed_submission = null;
    public ?array $successful_submission = null;
}

/** FxRate entity data model. */
class FxRate
{
    public ?string $destination_currency = null;
    public ?float $exchange_rate = null;
    public ?string $expiry_time = null;
    public ?string $quote_id = null;
    public ?string $source_currency = null;
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
    public ?string $payment_request_id = null;
    public ?string $response_type = null;
}

/** Request payload for IPayment#create. */
class IPaymentCreateData
{
    public ?string $payment_request_id = null;
    public ?string $response_type = null;
}

/** Mandate entity data model. */
class Mandate
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

/** Request payload for Mandate#load. */
class MandateLoadMatch
{
    public string $id;
}

/** Request payload for Mandate#create. */
class MandateCreateData
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

/** Merchant entity data model. */
class Merchant
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
    public ?string $reason = null;
    public ?string $short_name = null;
    public ?array $supported_payment_methods_list = null;
    public ?string $suspension_reason = null;
    public ?array $tag = null;
    public ?string $time_zone_id = null;
    public ?string $trading_name = null;
    public ?int $web_hook_limit = null;
    public ?string $your_role_name = null;
}

/** Request payload for Merchant#load. */
class MerchantLoadMatch
{
    public string $id;
}

/** Request payload for Merchant#list. */
class MerchantListMatch
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
    public ?string $reason = null;
    public ?string $short_name = null;
    public ?array $supported_payment_methods_list = null;
    public ?string $suspension_reason = null;
    public ?array $tag = null;
    public ?string $time_zone_id = null;
    public ?string $trading_name = null;
    public ?int $web_hook_limit = null;
    public ?string $your_role_name = null;
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

/** MerchantAuthorisationSetting entity data model. */
class MerchantAuthorisationSetting
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

/** Request payload for MerchantAuthorisationSetting#list. */
class MerchantAuthorisationSettingListMatch
{
    public string $merchant_id;
}

/** MerchantDirectDebitMandatePage entity data model. */
class MerchantDirectDebitMandatePage
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

/** Request payload for MerchantDirectDebitMandatePage#list. */
class MerchantDirectDebitMandatePageListMatch
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

/** MerchantPayByBankSetting entity data model. */
class MerchantPayByBankSetting
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

/** Request payload for MerchantPayByBankSetting#list. */
class MerchantPayByBankSettingListMatch
{
    public string $merchant_id;
}

/** MerchantPaymentRequestTemplate entity data model. */
class MerchantPaymentRequestTemplate
{
    public string $description;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?string $last_updated = null;
    public ?string $merchant_id = null;
    public string $name;
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

/** Request payload for MerchantToken#update. */
class MerchantTokenUpdateData
{
    public string $id;
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
    public ?int $build_version = null;
    public ?int $major_version = null;
    public ?int $minor_version = null;
    public ?string $release_name = null;
}

/** Request payload for NoFrixionVersion#load. */
class NoFrixionVersionLoadMatch
{
    public ?int $build_version = null;
    public ?int $major_version = null;
    public ?int $minor_version = null;
    public ?string $release_name = null;
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

/** Payment entity data model. */
class Payment
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

/** Request payload for Payment#load. */
class PaymentLoadMatch
{
    public ?string $id = null;
    public ?string $order_id = null;
}

/** Request payload for Payment#create. */
class PaymentCreateData
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

/** Request payload for Payment#update. */
class PaymentUpdateData
{
    public string $id;
}

/** PaymentAccount entity data model. */
class PaymentAccount
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

/** Request payload for PaymentAccount#list. */
class PaymentAccountListMatch
{
    public ?string $account_id = null;
}

/** PaymentAccountMinimal entity data model. */
class PaymentAccountMinimal
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

/** Request payload for PaymentAccountMinimal#list. */
class PaymentAccountMinimalListMatch
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

/** PaymentInitiation entity data model. */
class PaymentInitiation
{
    public ?string $payment_initiation_id = null;
    public ?string $payment_request_callback_url = null;
    public ?string $payment_request_id = null;
    public ?string $redirect_url = null;
    public ?string $response_type = null;
    public ?string $specific_error_message = null;
}

/** Request payload for PaymentInitiation#create. */
class PaymentInitiationCreateData
{
    public string $paymentrequest_id;
}

/** PaymentRequest entity data model. */
class PaymentRequest
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
    public ?bool $do_simulate_settlement_failure = null;
    public ?string $due_date = null;
    public ?string $error_description = null;
    public ?array $event = null;
    public ?array $failed_payment_request = null;
    public ?string $failure_callback_url = null;
    public ?array $field_display_setting = null;
    public ?string $formatted_amount = null;
    public ?string $hosted_pay_checkout_url = null;
    public ?string $id = null;
    public ?bool $ignore_address_verification = null;
    public ?string $inserted = null;
    public ?string $inserted_sortable = null;
    public ?string $institution = null;
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
    public ?string $payment_initiation_id = null;
    public ?array $payment_method = null;
    public ?string $payment_processor = null;
    public ?array $payment_request = null;
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

/** Request payload for PaymentRequest#load. */
class PaymentRequestLoadMatch
{
    public ?string $paymentrequest_id = null;
}

/** Request payload for PaymentRequest#list. */
class PaymentRequestListMatch
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
    public ?bool $do_simulate_settlement_failure = null;
    public ?string $due_date = null;
    public ?string $error_description = null;
    public ?array $event = null;
    public ?array $failed_payment_request = null;
    public ?string $failure_callback_url = null;
    public ?array $field_display_setting = null;
    public ?string $formatted_amount = null;
    public ?string $hosted_pay_checkout_url = null;
    public ?string $id = null;
    public ?bool $ignore_address_verification = null;
    public ?string $inserted = null;
    public ?string $inserted_sortable = null;
    public ?string $institution = null;
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
    public ?string $payment_initiation_id = null;
    public ?array $payment_method = null;
    public ?string $payment_processor = null;
    public ?array $payment_request = null;
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

/** Request payload for PaymentRequest#create. */
class PaymentRequestCreateData
{
    public ?string $paymentrequest_id = null;
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

/** PaymentRequestEvent entity data model. */
class PaymentRequestEvent
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

/** Request payload for PaymentRequestEvent#list. */
class PaymentRequestEventListMatch
{
    public string $paymentrequest_id;
}

/** PaymentRequestMetric entity data model. */
class PaymentRequestMetric
{
    public ?int $all = null;
    public ?int $authorized = null;
    public ?int $paid = null;
    public ?int $partially_paid = null;
    public ?array $total_amounts_by_currency = null;
    public ?int $unpaid = null;
}

/** Request payload for PaymentRequestMetric#load. */
class PaymentRequestMetricLoadMatch
{
    public ?int $all = null;
    public ?int $authorized = null;
    public ?int $paid = null;
    public ?int $partially_paid = null;
    public ?array $total_amounts_by_currency = null;
    public ?int $unpaid = null;
}

/** PaymentRequestMinimal entity data model. */
class PaymentRequestMinimal
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

/** Request payload for PaymentRequestMinimal#list. */
class PaymentRequestMinimalListMatch
{
    public string $paymentrequest_id;
}

/** PaymentRequestResult entity data model. */
class PaymentRequestResult
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

/** Request payload for PaymentRequestResult#list. */
class PaymentRequestResultListMatch
{
    public string $paymentrequest_id;
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
    public ?array $failed_payout = null;
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
    public ?array $payout = null;
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

/** PayoutKeysetPage entity data model. */
class PayoutKeysetPage
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

/** Request payload for PayoutKeysetPage#list. */
class PayoutKeysetPageListMatch
{
    public ?string $account_id = null;
    public ?string $merchant_id = null;
}

/** PayoutMetric entity data model. */
class PayoutMetric
{
    public ?float $all = null;
    public ?float $failed = null;
    public ?float $in_progress = null;
    public ?float $paid = null;
    public ?float $pending_approval = null;
    public ?float $scheduled = null;
    public ?array $total_amounts_by_currency = null;
}

/** Request payload for PayoutMetric#load. */
class PayoutMetricLoadMatch
{
    public ?float $all = null;
    public ?float $failed = null;
    public ?float $in_progress = null;
    public ?float $paid = null;
    public ?float $pending_approval = null;
    public ?float $scheduled = null;
    public ?array $total_amounts_by_currency = null;
}

/** Payrun entity data model. */
class Payrun
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
    public ?string $note = null;
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

/** Request payload for Payrun#load. */
class PayrunLoadMatch
{
    public string $id;
}

/** Request payload for Payrun#list. */
class PayrunListMatch
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
    public ?string $note = null;
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
    public ?string $content = null;
    public ?string $content_type = null;
    public ?string $last_completed_at = null;
    public ?string $merchant_id = null;
    public ?string $report_name = null;
    public ?string $report_type = null;
    public ?int $statement_number = null;
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
    public ?array $failed_role = null;
    public ?array $role = null;
}

/** Request payload for Role#create. */
class RoleCreateData
{
    public string $merchant_id;
}

/** Rule entity data model. */
class Rule
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

/** Request payload for Rule#load. */
class RuleLoadMatch
{
    public string $id;
}

/** Request payload for Rule#list. */
class RuleListMatch
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

/** Request payload for Rule#create. */
class RuleCreateData
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

/** RuleEvent entity data model. */
class RuleEvent
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

/** Request payload for RuleEvent#list. */
class RuleEventListMatch
{
    public string $id;
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
    public ?string $account_id = null;
    public ?string $account_name = null;
    public ?int $account_sequence_number = null;
    public ?array $address_detail = null;
    public ?float $amount = null;
    public ?int $amount_minor_unit = null;
    public ?float $balance = null;
    public ?int $balance_minor_unit = null;
    public ?string $booking_date_time = null;
    public ?array $charge_detail = null;
    public ?array $content = null;
    public ?array $counterparty = null;
    public ?string $counterparty_summary = null;
    public ?string $currency = null;
    public ?array $currency_exchange = null;
    public ?string $date = null;
    public ?string $description = null;
    public ?array $enrichment = null;
    public ?float $fx_amount = null;
    public ?string $fx_currency = null;
    public ?float $fx_rate = null;
    public array $gross_amount;
    public ?string $id = null;
    public ?string $inserted = null;
    public ?array $iso_bank_transaction_code = null;
    public ?array $merchant = null;
    public ?string $merchant_id = null;
    public ?int $page_number = null;
    public ?int $page_size = null;
    public array $payee_detail;
    public array $payer_detail;
    public ?array $payment_request_custom_field = null;
    public ?string $payment_request_id = null;
    public ?string $payout_id = null;
    public ?array $proprietary_bank_transaction_code = null;
    public ?string $raw_reference = null;
    public ?string $reference = null;
    public ?string $rule_id = null;
    public ?array $statement_reference = null;
    public ?string $status = null;
    public mixed $supplementary_data = null;
    public ?array $tag = null;
    public ?string $their_reference = null;
    public ?int $total_page = null;
    public ?int $total_size = null;
    public array $transaction_amount;
    public ?string $transaction_date = null;
    public ?array $transaction_information = null;
    public ?string $transaction_mutability = null;
    public ?string $type = null;
    public ?string $value_date_time = null;
    public ?string $virtual_iban = null;
    public ?string $your_reference = null;
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
    public ?array $authorisation_status = null;
    public ?array $failed_user_invite = null;
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
    public ?array $user_invite = null;
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

/** Request payload for Webhook#update. */
class WebhookUpdateData
{
    public string $id;
}

/** Request payload for Webhook#remove. */
class WebhookRemoveMatch
{
    public string $id;
}

