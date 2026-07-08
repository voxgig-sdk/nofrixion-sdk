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
    created_by: dict
    identifier: dict


class Account(AccountRequired, total=False):
    account_id: str
    account_name: str
    account_supplier_name: str
    account_type: str
    available_balance: float
    available_balance_minor_unit: int
    balance: float
    balance_minor_unit: int
    bank_name: str
    consent_id: str
    created_by_display_name: str
    currency: str
    default_payment_rail: str
    display_name: str
    expiry_date: str
    external_account_icon: str
    format: str
    from_date: str
    id: str
    inserted: str
    is_archived: bool
    is_connected_account: bool
    is_default: bool
    is_trust_account: bool
    is_virtual: bool
    last_transaction: dict
    last_updated: str
    merchant_id: str
    merchant_name: str
    physical_account_id: str
    role_i_d: list
    rule: list
    submitted_payouts_balance: float
    submitted_payouts_balance_minor_unit: int
    summary: str
    supplier_physical_account_id: str
    supplier_sepa_instant_status: str
    to_date: str
    xero_bank_feed_connection_status: str
    xero_bank_feed_last_synced_at: str
    xero_bank_feed_sync_last_failed_at: str
    xero_bank_feed_sync_last_failure_reason: str
    xero_bank_feed_sync_status: str
    xero_unsynchronised_transactions_count: int


class AccountLoadMatchRequired(TypedDict):
    id: str


class AccountLoadMatch(AccountLoadMatchRequired, total=False):
    account_id: str
    merchant_id: str


class AccountListMatch(TypedDict, total=False):
    merchant_id: str


class AccountCreateData(TypedDict, total=False):
    account_id: str
    currency: str


class AccountUpdateData(TypedDict, total=False):
    account_id: str
    amount: float
    id: str


class AccountRemoveMatch(TypedDict):
    id: str


class BeneficiaryRequired(TypedDict):
    created_by: dict
    currency: str
    name: str


class Beneficiary(BeneficiaryRequired, total=False):
    approval_callback_url: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    beneficiary_event: list
    can_authorise: bool
    can_update: bool
    created_by_email_address: str
    destination: dict
    has_current_user_authorised: bool
    id: str
    inserted: str
    is_enabled: bool
    last_authorised: str
    last_updated: str
    merchant_id: str
    nonce: str
    source_account: list
    source_account_i_d: list
    their_reference: str


class BeneficiaryLoadMatchRequired(TypedDict):
    id: str


class BeneficiaryLoadMatch(BeneficiaryLoadMatchRequired, total=False):
    merchant_id: str


class BeneficiaryCreateData(TypedDict, total=False):
    id: str


class BeneficiaryUpdateData(TypedDict):
    id: str


class BeneficiaryRemoveMatch(TypedDict):
    id: str


class CancelRequired(TypedDict):
    beneficiary: dict
    source_account_identifier: dict


class Cancel(CancelRequired, total=False):
    account_id: str
    amount: float
    amount_minor_unit: int
    approve_payout_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    can_authorise: bool
    can_process: bool
    can_update: bool
    charge_bearer: str
    created_by: str
    created_by_email_address: str
    currency: str
    current_user_id: str
    description: str
    destination: dict
    document: list
    event: list
    formatted_amount: str
    formatted_fx_destination_amount: str
    formatted_schedule: str
    formatted_schedule_day_only: str
    formatted_source_account_available_balance: str
    fx_destination_amount: float
    fx_destination_amount_minor_unit: int
    fx_destination_currency: str
    fx_quote_expires_at: str
    fx_quote_id: str
    fx_rate: float
    fx_use_destination_amount: bool
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice_id: str
    is_archived: bool
    is_failed: bool
    is_settled: bool
    is_submitted: bool
    last_updated: str
    merchant_id: str
    merchant_token_description: str
    nonce: str
    payment_processor: str
    payment_rail: str
    payrun_id: str
    payrun_name: str
    rule: dict
    schedule_date: str
    scheduled: bool
    source_account_available_balance: float
    source_account_available_balance_minor_unit: int
    source_account_bic: str
    source_account_currency: str
    source_account_iban: str
    source_account_name: str
    source_account_number: str
    source_account_sortcode: str
    status: str
    tag: list
    their_reference: str
    topup_payrun_id: str
    transacted_amount: float
    transacted_fx_amount: float
    transacted_fx_rate: float
    type: str
    user_id: str
    your_reference: str


class CancelUpdateData(TypedDict):
    id: str


class DisableRequired(TypedDict):
    created_by: dict
    currency: str
    name: str


class Disable(DisableRequired, total=False):
    approval_callback_url: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    beneficiary_event: list
    can_authorise: bool
    can_update: bool
    created_by_email_address: str
    destination: dict
    has_current_user_authorised: bool
    id: str
    inserted: str
    is_enabled: bool
    last_authorised: str
    last_updated: str
    merchant_id: str
    nonce: str
    source_account: list
    their_reference: str


class DisableUpdateData(TypedDict):
    id: str


class EnableRequired(TypedDict):
    created_by: dict
    currency: str
    name: str


class Enable(EnableRequired, total=False):
    approval_callback_url: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    beneficiary_event: list
    can_authorise: bool
    can_update: bool
    created_by_email_address: str
    destination: dict
    has_current_user_authorised: bool
    id: str
    inserted: str
    is_enabled: bool
    last_authorised: str
    last_updated: str
    merchant_id: str
    nonce: str
    source_account: list
    their_reference: str


class EnableUpdateData(TypedDict):
    id: str


class Merchant(TypedDict, total=False):
    reason: str


class MerchantLoadMatch(TypedDict):
    merchant_id: str


class MerchantUpdateData(TypedDict):
    id: str


class MerchantRemoveMatch(TypedDict, total=False):
    id: str
    user_id: str
    merchant_id: str
    tag_id: str


class Metadata(TypedDict):
    pass


class MetadataLoadMatch(TypedDict):
    pass


class NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePage(TypedDict, total=False):
    approved_at: str
    currency: str
    customer_account_number: str
    customer_city: str
    customer_country_code: str
    customer_country_name: str
    customer_email_address: str
    customer_first_name: str
    customer_iban: str
    customer_last_name: str
    customer_sort_code: str
    id: str
    inserted: str
    is_recurring: bool
    last_updated: str
    merchant_id: str
    reference: str
    status: str
    supplier_bank_account_id: str
    supplier_customer_id: str
    supplier_mandate_id: str
    supplier_name: str
    supplier_status: str


class NoFrixionBizBizModelsPagingMerchantDirectDebitMandatePageListMatch(TypedDict, total=False):
    approved_at: str
    currency: str
    customer_account_number: str
    customer_city: str
    customer_country_code: str
    customer_country_name: str
    customer_email_address: str
    customer_first_name: str
    customer_iban: str
    customer_last_name: str
    customer_sort_code: str
    id: str
    inserted: str
    is_recurring: bool
    last_updated: str
    merchant_id: str
    reference: str
    status: str
    supplier_bank_account_id: str
    supplier_customer_id: str
    supplier_mandate_id: str
    supplier_name: str
    supplier_status: str


class NoFrixionBizBizModelsPagingPaymentRequestPageRequired(TypedDict):
    created_by_user: dict


class NoFrixionBizBizModelsPagingPaymentRequestPage(NoFrixionBizBizModelsPagingPaymentRequestPageRequired, total=False):
    address: list
    amount: float
    amount_pending: float
    amount_received: float
    amount_refunded: float
    auto_send_receipt: bool
    base_origin_url: str
    callback_url: str
    card_authorize_only: bool
    card_create_token: bool
    card_create_token_mode: str
    card_ignore_cvn: bool
    card_processor_merchant_id: str
    card_stripe_payment_intent_id: str
    card_stripe_payment_intent_secret: str
    currency: str
    custom_field: list
    customer_email_address: str
    customer_id: str
    customer_name: str
    description: str
    destination_account: dict
    direct_debit_payment: dict
    due_date: str
    event: list
    failure_callback_url: str
    field_display_setting: list
    formatted_amount: str
    hosted_pay_checkout_url: str
    id: str
    ignore_address_verification: bool
    inserted: str
    inserted_sortable: str
    is_archived: bool
    jwk: str
    last_updated: str
    lightning_invoice: str
    lightning_invoice_expires_at: str
    merchant_direct_debit_mandate_id: str
    merchant_id: str
    merchant_token_description: str
    notification_email_address: str
    notification_role_i_d: list
    order_id: str
    partial_payment_method: str
    partial_payment_step: str
    payment_attempt: list
    payment_method: list
    payment_processor: str
    payrun_id: str
    pisp_account_id: str
    priority_bank_id: str
    result: dict
    sandbox_settle_delay_in_second: int
    shipping_address: dict
    status: str
    success_web_hook_url: str
    tag: list
    title: str
    tokenised_card: list
    transaction: list
    use_hosted_payment_page: bool


class NoFrixionBizBizModelsPagingPaymentRequestPageListMatch(TypedDict, total=False):
    address: list
    amount: float
    amount_pending: float
    amount_received: float
    amount_refunded: float
    auto_send_receipt: bool
    base_origin_url: str
    callback_url: str
    card_authorize_only: bool
    card_create_token: bool
    card_create_token_mode: str
    card_ignore_cvn: bool
    card_processor_merchant_id: str
    card_stripe_payment_intent_id: str
    card_stripe_payment_intent_secret: str
    created_by_user: dict
    currency: str
    custom_field: list
    customer_email_address: str
    customer_id: str
    customer_name: str
    description: str
    destination_account: dict
    direct_debit_payment: dict
    due_date: str
    event: list
    failure_callback_url: str
    field_display_setting: list
    formatted_amount: str
    hosted_pay_checkout_url: str
    id: str
    ignore_address_verification: bool
    inserted: str
    inserted_sortable: str
    is_archived: bool
    jwk: str
    last_updated: str
    lightning_invoice: str
    lightning_invoice_expires_at: str
    merchant_direct_debit_mandate_id: str
    merchant_id: str
    merchant_token_description: str
    notification_email_address: str
    notification_role_i_d: list
    order_id: str
    partial_payment_method: str
    partial_payment_step: str
    payment_attempt: list
    payment_method: list
    payment_processor: str
    payrun_id: str
    pisp_account_id: str
    priority_bank_id: str
    result: dict
    sandbox_settle_delay_in_second: int
    shipping_address: dict
    status: str
    success_web_hook_url: str
    tag: list
    title: str
    tokenised_card: list
    transaction: list
    use_hosted_payment_page: bool


class NoFrixionBizBizModelsPagingPayoutPageRequired(TypedDict):
    beneficiary: dict
    source_account_identifier: dict


class NoFrixionBizBizModelsPagingPayoutPage(NoFrixionBizBizModelsPagingPayoutPageRequired, total=False):
    account_id: str
    amount: float
    amount_minor_unit: int
    approve_payout_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    can_authorise: bool
    can_process: bool
    can_update: bool
    charge_bearer: str
    created_by: str
    created_by_email_address: str
    currency: str
    current_user_id: str
    description: str
    destination: dict
    document: list
    event: list
    formatted_amount: str
    formatted_fx_destination_amount: str
    formatted_schedule: str
    formatted_schedule_day_only: str
    formatted_source_account_available_balance: str
    fx_destination_amount: float
    fx_destination_amount_minor_unit: int
    fx_destination_currency: str
    fx_quote_expires_at: str
    fx_quote_id: str
    fx_rate: float
    fx_use_destination_amount: bool
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice_id: str
    is_archived: bool
    is_failed: bool
    is_settled: bool
    is_submitted: bool
    last_updated: str
    merchant_id: str
    merchant_token_description: str
    nonce: str
    payment_processor: str
    payment_rail: str
    payrun_id: str
    payrun_name: str
    rule: dict
    schedule_date: str
    scheduled: bool
    source_account_available_balance: float
    source_account_available_balance_minor_unit: int
    source_account_bic: str
    source_account_currency: str
    source_account_iban: str
    source_account_name: str
    source_account_number: str
    source_account_sortcode: str
    status: str
    tag: list
    their_reference: str
    topup_payrun_id: str
    transacted_amount: float
    transacted_fx_amount: float
    transacted_fx_rate: float
    type: str
    user_id: str
    your_reference: str


class NoFrixionBizBizModelsPagingPayoutPageListMatch(TypedDict, total=False):
    account_id: str
    merchant_id: str


class NoFrixionBizBizModelsPagingPayrunPageRequired(TypedDict):
    last_updated_by: dict


class NoFrixionBizBizModelsPagingPayrunPage(NoFrixionBizBizModelsPagingPayrunPageRequired, total=False):
    authorisation: list
    authorisation_date: str
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    can_authorise: bool
    can_delete: bool
    can_edit: bool
    event: list
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice: list
    invoices_minimal: list
    is_archived: bool
    last_updated: str
    merchant_id: str
    name: str
    nonce: str
    payment: list
    payout: list
    payouts_count: int
    schedule_date: str
    source_account: list
    status: str
    total_eur: float
    total_gbp: float
    total_usd: float


class NoFrixionBizBizModelsPagingPayrunPageListMatch(TypedDict, total=False):
    authorisation: list
    authorisation_date: str
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    can_authorise: bool
    can_delete: bool
    can_edit: bool
    event: list
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice: list
    invoices_minimal: list
    is_archived: bool
    last_updated: str
    last_updated_by: dict
    merchant_id: str
    name: str
    nonce: str
    payment: list
    payout: list
    payouts_count: int
    schedule_date: str
    source_account: list
    status: str
    total_eur: float
    total_gbp: float
    total_usd: float


class NoFrixionBizBizModelsPagingRuleEventsPageRequired(TypedDict):
    user: dict


class NoFrixionBizBizModelsPagingRuleEventsPage(NoFrixionBizBizModelsPagingRuleEventsPageRequired, total=False):
    error_message: str
    id: str
    inserted: str
    is_authorise_to_enable: bool
    message: str
    raw_response: str
    rule_event_type: str
    rule_id: str


class NoFrixionBizBizModelsPagingRuleEventsPageListMatch(TypedDict):
    rule_id: str


class NoFrixionBizBizModelsPagingRulesPageRequired(TypedDict):
    created_by: dict
    nonce: str


class NoFrixionBizBizModelsPagingRulesPage(NoFrixionBizBizModelsPagingRulesPageRequired, total=False):
    account: dict
    account_id: str
    approve_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    can_authorise: bool
    description: str
    end_at: str
    has_current_user_authorised: bool
    id: str
    inserted: str
    is_disabled: bool
    last_executed_at: str
    last_run_at_transaction_date: str
    last_updated: str
    merchant_id: str
    name: str
    on_approved_web_hook_url: str
    on_execution_error_web_hook_url: str
    on_execution_success_web_hook_url: str
    start_at: str
    status: str
    sweep_action: dict
    time_zone_id: str
    trigger_cron_expression: str
    trigger_on_pay_in: bool
    user_id: str
    web_hook_secret: str


class NoFrixionBizBizModelsPagingRulesPageListMatch(TypedDict, total=False):
    account: dict
    account_id: str
    approve_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    can_authorise: bool
    created_by: dict
    description: str
    end_at: str
    has_current_user_authorised: bool
    id: str
    inserted: str
    is_disabled: bool
    last_executed_at: str
    last_run_at_transaction_date: str
    last_updated: str
    merchant_id: str
    name: str
    nonce: str
    on_approved_web_hook_url: str
    on_execution_error_web_hook_url: str
    on_execution_success_web_hook_url: str
    start_at: str
    status: str
    sweep_action: dict
    time_zone_id: str
    trigger_cron_expression: str
    trigger_on_pay_in: bool
    user_id: str
    web_hook_secret: str


class NoFrixionBizBizModelsPaymentsCardPayment(TypedDict, total=False):
    authorized_amount: str
    currency_code: str
    is_payer_authentication_required: bool
    is_soft_decline: bool
    payer_authentication_access_token: str
    payer_authentication_merchant_data: str
    payer_authentication_url: str
    payer_authentication_window_height: int
    payer_authentication_window_width: int
    payment_request_callback_url: str
    payment_request_id: str
    request_id: str
    response_code: str
    response_type: str
    status: str
    three_ds_redirect_url: str
    transaction_id: str


class NoFrixionBizBizModelsPaymentsCardPaymentCreateDataRequired(TypedDict):
    paymentrequest_id: str


class NoFrixionBizBizModelsPaymentsCardPaymentCreateData(NoFrixionBizBizModelsPaymentsCardPaymentCreateDataRequired, total=False):
    partial_refund_amount: float


class NoFrixionBizBizModelsPaymentsCardPublicKey(TypedDict, total=False):
    jwt: str


class NoFrixionBizBizModelsPaymentsCardPublicKeyLoadMatch(TypedDict):
    paymentrequest_id: str


class NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiaries(TypedDict, total=False):
    beneficiary: list
    failed_beneficiary: dict


class NoFrixionMoneyMoovApiFeaturesBeneficiariesBeneficiariesCreateData(TypedDict, total=False):
    beneficiary: list
    failed_beneficiary: dict


class NoFrixionMoneyMoovApiFeaturesPaymentRequestsPayment(TypedDict, total=False):
    failed_payment_request: dict
    payment_request: list


class NoFrixionMoneyMoovApiFeaturesPaymentRequestsPaymentCreateData(TypedDict, total=False):
    failed_payment_request: dict
    payment_request: list


class NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreate(TypedDict, total=False):
    failed_role: dict
    role: list


class NoFrixionMoneyMoovApiFeaturesPermissionsRolesCreateCreateData(TypedDict):
    merchant_id: str


class NoFrixionMoneyMoovApiFeaturesUserInvitesCreate(TypedDict, total=False):
    failed_user_invite: dict
    user_invite: list


class NoFrixionMoneyMoovApiFeaturesUserInvitesCreateCreateData(TypedDict, total=False):
    failed_user_invite: dict
    user_invite: list


class NoFrixionMoneyMoovModelsAuthorisationSettingsMerchant(TypedDict, total=False):
    amount_lower: float
    amount_upper: float
    authorisation_type: str
    beneficiaries_only: bool
    id: str
    inserted: str
    last_editor_cant_authorise: bool
    last_updated: str
    merchant_id: str
    number_of_authoriser: int
    role_setting: list


class NoFrixionMoneyMoovModelsAuthorisationSettingsMerchantListMatch(TypedDict):
    merchant_id: str


class NoFrixionMoneyMoovModelsBatchPayout(TypedDict, total=False):
    approve_url: str
    id: str
    payout: list


class NoFrixionMoneyMoovModelsBatchPayoutLoadMatch(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsBatchPayoutCreateData(TypedDict, total=False):
    approve_url: str
    id: str
    payout: list


class NoFrixionMoneyMoovModelsBeneficiaryGroupPageRequired(TypedDict):
    group_name: str
    merchant_id: str


class NoFrixionMoneyMoovModelsBeneficiaryGroupPage(NoFrixionMoneyMoovModelsBeneficiaryGroupPageRequired, total=False):
    group_member: list
    id: str
    inserted: str
    last_updated: str


class NoFrixionMoneyMoovModelsBeneficiaryGroupPageListMatch(TypedDict):
    merchant_id: str


class NoFrixionMoneyMoovModelsBeneficiaryPageRequired(TypedDict):
    created_by: dict
    currency: str
    name: str


class NoFrixionMoneyMoovModelsBeneficiaryPage(NoFrixionMoneyMoovModelsBeneficiaryPageRequired, total=False):
    approval_callback_url: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    beneficiary_event: list
    can_authorise: bool
    can_update: bool
    created_by_email_address: str
    destination: dict
    has_current_user_authorised: bool
    id: str
    inserted: str
    is_enabled: bool
    last_authorised: str
    last_updated: str
    merchant_id: str
    nonce: str
    source_account: list
    their_reference: str


class NoFrixionMoneyMoovModelsBeneficiaryPageListMatch(TypedDict, total=False):
    merchant_id: str


class NoFrixionMoneyMoovModelsCardCustomerToken(TypedDict, total=False):
    card_type: str
    customer_email_address: str
    expiry_month: str
    expiry_year: str
    id: str
    inserted: str
    last_four_digit: str
    last_updated: str
    masked_card_number: str
    merchant_id: str
    payment_request_id: str


class NoFrixionMoneyMoovModelsCardCustomerTokenLoadMatch(TypedDict):
    customer_email_address: str


class NoFrixionMoneyMoovModelsCardCustomerTokenListMatch(TypedDict):
    customer_email_address: str
    merchant_id: str


class NoFrixionMoneyMoovModelsCardCustomerTokenRemoveMatch(TypedDict, total=False):
    customer_email_address: str
    merchant_id: str
    id: str


class NoFrixionMoneyMoovModelsCurrencyCurrencyInfo(TypedDict, total=False):
    code: str
    decimal: int
    is_fiat: bool
    iso4217_alpha_code: str
    iso4217_numeric_code: str
    symbol: str


class NoFrixionMoneyMoovModelsCurrencyCurrencyInfoListMatch(TypedDict, total=False):
    code: str
    decimal: int
    is_fiat: bool
    iso4217_alpha_code: str
    iso4217_numeric_code: str
    symbol: str


class NoFrixionMoneyMoovModelsDirectDebitBatchSubmit(TypedDict, total=False):
    failed_submission: dict
    successful_submission: list


class NoFrixionMoneyMoovModelsDirectDebitBatchSubmitCreateData(TypedDict, total=False):
    failed_submission: dict
    successful_submission: list


class NoFrixionMoneyMoovModelsFxRate(TypedDict, total=False):
    destination_currency: str
    exchange_rate: float
    expiry_time: str
    quote_id: str
    source_currency: str


class NoFrixionMoneyMoovModelsFxRateLoadMatch(TypedDict):
    destination: str
    source: str
    valid_for_minute: int


class NoFrixionMoneyMoovModelsFxRateListMatch(TypedDict):
    destination: str
    source: str


class NoFrixionMoneyMoovModelsIPayment(TypedDict, total=False):
    payment_request_id: str
    response_type: str


class NoFrixionMoneyMoovModelsIPaymentCreateData(TypedDict, total=False):
    payment_request_id: str
    response_type: str


class NoFrixionMoneyMoovModelsMandatesMandateRequired(TypedDict):
    address_line1: str
    city: str
    country_code: str
    email_address: str
    first_name: str
    last_name: str
    postal_code: str


class NoFrixionMoneyMoovModelsMandatesMandate(NoFrixionMoneyMoovModelsMandatesMandateRequired, total=False):
    account_number: str
    address_line2: str
    approved_at: str
    currency: str
    customer_account_number: str
    customer_city: str
    customer_country_code: str
    customer_country_name: str
    customer_email_address: str
    customer_first_name: str
    customer_iban: str
    customer_last_name: str
    customer_sort_code: str
    iban: str
    id: str
    inserted: str
    is_recurring: bool
    last_updated: str
    merchant_id: str
    reference: str
    sort_code: str
    status: str
    supplier_bank_account_id: str
    supplier_customer_id: str
    supplier_mandate_id: str
    supplier_name: str
    supplier_status: str


class NoFrixionMoneyMoovModelsMandatesMandateLoadMatch(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsMandatesMandateCreateDataRequired(TypedDict):
    address_line1: str
    city: str
    country_code: str
    email_address: str
    first_name: str
    last_name: str
    postal_code: str


class NoFrixionMoneyMoovModelsMandatesMandateCreateData(NoFrixionMoneyMoovModelsMandatesMandateCreateDataRequired, total=False):
    account_number: str
    address_line2: str
    approved_at: str
    currency: str
    customer_account_number: str
    customer_city: str
    customer_country_code: str
    customer_country_name: str
    customer_email_address: str
    customer_first_name: str
    customer_iban: str
    customer_last_name: str
    customer_sort_code: str
    iban: str
    id: str
    inserted: str
    is_recurring: bool
    last_updated: str
    merchant_id: str
    reference: str
    sort_code: str
    status: str
    supplier_bank_account_id: str
    supplier_customer_id: str
    supplier_mandate_id: str
    supplier_name: str
    supplier_status: str


class NoFrixionMoneyMoovModelsMerchant(TypedDict, total=False):
    account_currency: list
    can_have_trust_account: bool
    card_payment_processor: str
    company_id: str
    display_qr_on_hosted_pay: bool
    hosted_pay_version: int
    id: str
    inserted: str
    is_blocked: bool
    is_exited: bool
    is_suspended: bool
    jurisdiction: str
    logo_url_png: str
    logo_url_svg: str
    merchant_category_code: str
    name: str
    note: str
    parent_merchant: dict
    payment_account: list
    payment_account_limit: int
    short_name: str
    supported_payment_methods_list: list
    suspension_reason: str
    tag: list
    time_zone_id: str
    trading_name: str
    web_hook_limit: int
    your_role_name: str


class NoFrixionMoneyMoovModelsMerchantLoadMatch(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsMerchantListMatch(TypedDict, total=False):
    account_currency: list
    can_have_trust_account: bool
    card_payment_processor: str
    company_id: str
    display_qr_on_hosted_pay: bool
    hosted_pay_version: int
    id: str
    inserted: str
    is_blocked: bool
    is_exited: bool
    is_suspended: bool
    jurisdiction: str
    logo_url_png: str
    logo_url_svg: str
    merchant_category_code: str
    name: str
    note: str
    parent_merchant: dict
    payment_account: list
    payment_account_limit: int
    short_name: str
    supported_payment_methods_list: list
    suspension_reason: str
    tag: list
    time_zone_id: str
    trading_name: str
    web_hook_limit: int
    your_role_name: str


class NoFrixionMoneyMoovModelsMerchantUpdateData(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsMerchantPage(TypedDict, total=False):
    account_currency: list
    can_have_trust_account: bool
    card_payment_processor: str
    company_id: str
    display_qr_on_hosted_pay: bool
    hosted_pay_version: int
    id: str
    inserted: str
    is_blocked: bool
    is_exited: bool
    is_suspended: bool
    jurisdiction: str
    logo_url_png: str
    logo_url_svg: str
    merchant_category_code: str
    name: str
    note: str
    parent_merchant: dict
    payment_account: list
    payment_account_limit: int
    short_name: str
    supported_payment_methods_list: list
    suspension_reason: str
    tag: list
    time_zone_id: str
    trading_name: str
    web_hook_limit: int
    your_role_name: str


class NoFrixionMoneyMoovModelsMerchantPageListMatch(TypedDict, total=False):
    merchant_id: str


class NoFrixionMoneyMoovModelsMerchantPayByBankSetting(TypedDict, total=False):
    bank_country_code: list
    bank_id: str
    bank_name: str
    business_institution_id: str
    currency: str
    logo: str
    message: str
    message_image_url: str
    order: int
    personal_institution_id: str
    processor: str
    warning_heading: str
    warning_message: str


class NoFrixionMoneyMoovModelsMerchantPayByBankSettingListMatch(TypedDict):
    merchant_id: str


class NoFrixionMoneyMoovModelsMerchantTokenRequired(TypedDict):
    nonce: str


class NoFrixionMoneyMoovModelsMerchantToken(NoFrixionMoneyMoovModelsMerchantTokenRequired, total=False):
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    can_authorise: bool
    description: str
    expires_at: str
    has_current_user_authorised: bool
    hmac_algorithm: str
    id: str
    inserted: str
    ip_address_whitelist: str
    is_archived: bool
    is_enabled: bool
    last_authorised: str
    last_updated: str
    merchant_id: str
    permission_type: list
    request_signature_version: int
    shared_secret_algorithm: str
    shared_secret_base64: str
    token: str


class NoFrixionMoneyMoovModelsMerchantTokenLoadMatch(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsMerchantTokenCreateDataRequired(TypedDict):
    nonce: str


class NoFrixionMoneyMoovModelsMerchantTokenCreateData(NoFrixionMoneyMoovModelsMerchantTokenCreateDataRequired, total=False):
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    can_authorise: bool
    description: str
    expires_at: str
    has_current_user_authorised: bool
    hmac_algorithm: str
    id: str
    inserted: str
    ip_address_whitelist: str
    is_archived: bool
    is_enabled: bool
    last_authorised: str
    last_updated: str
    merchant_id: str
    permission_type: list
    request_signature_version: int
    shared_secret_algorithm: str
    shared_secret_base64: str
    token: str


class NoFrixionMoneyMoovModelsMerchantTokenUpdateData(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsMerchantTokenPageRequired(TypedDict):
    nonce: str


class NoFrixionMoneyMoovModelsMerchantTokenPage(NoFrixionMoneyMoovModelsMerchantTokenPageRequired, total=False):
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    can_authorise: bool
    description: str
    expires_at: str
    has_current_user_authorised: bool
    id: str
    inserted: str
    ip_address_whitelist: str
    is_archived: bool
    is_enabled: bool
    last_authorised: str
    last_updated: str
    merchant_id: str
    permission_type: list
    request_signature_version: int
    shared_secret_algorithm: str
    shared_secret_base64: str
    token: str


class NoFrixionMoneyMoovModelsMerchantTokenPageListMatch(TypedDict):
    merchant_id: str


class NoFrixionMoneyMoovModelsNoFrixionVersion(TypedDict, total=False):
    build_version: int
    major_version: int
    minor_version: int
    release_name: str


class NoFrixionMoneyMoovModelsNoFrixionVersionLoadMatch(TypedDict, total=False):
    build_version: int
    major_version: int
    minor_version: int
    release_name: str


class NoFrixionMoneyMoovModelsOpenBankingAccount(TypedDict, total=False):
    account_balance: list
    account_identification: list
    account_name: list
    account_type: str
    balance: float
    consolidated_account_information: dict
    currency: str
    description: str
    detail: str
    id: str
    nickname: str
    type: str
    usage_type: str


class NoFrixionMoneyMoovModelsOpenBankingAccountLoadMatch(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsOpenBankingConsent(TypedDict, total=False):
    authorisation_url: str
    callback_url: str
    consent_id: str
    email_address: str
    expiry_date: str
    failure_callback_url: str
    id: str
    inserted: str
    institution_id: str
    is_connected_account: bool
    is_enabled: bool
    merchant_id: str
    provider: str
    success_web_hook_url: str


class NoFrixionMoneyMoovModelsOpenBankingConsentLoadMatch(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsOpenBankingConsentListMatch(TypedDict):
    email: str
    merchant_id: str


class NoFrixionMoneyMoovModelsOpenBankingConsentCreateData(TypedDict, total=False):
    authorisation_url: str
    callback_url: str
    consent_id: str
    email_address: str
    expiry_date: str
    failure_callback_url: str
    id: str
    inserted: str
    institution_id: str
    is_connected_account: bool
    is_enabled: bool
    merchant_id: str
    provider: str
    success_web_hook_url: str


class NoFrixionMoneyMoovModelsOpenBankingConsentUpdateData(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsOpenBankingConsentRemoveMatch(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsOpenBankingTransactionRequired(TypedDict):
    gross_amount: dict
    payee_detail: dict
    payer_detail: dict
    transaction_amount: dict


class NoFrixionMoneyMoovModelsOpenBankingTransaction(NoFrixionMoneyMoovModelsOpenBankingTransactionRequired, total=False):
    address_detail: dict
    amount: float
    balance: dict
    booking_date_time: str
    charge_detail: dict
    currency: str
    currency_exchange: dict
    date: str
    description: str
    enrichment: dict
    id: str
    iso_bank_transaction_code: dict
    merchant: dict
    proprietary_bank_transaction_code: dict
    reference: str
    statement_reference: list
    status: str
    supplementary_data: Any
    transaction_information: list
    transaction_mutability: str
    value_date_time: str


class NoFrixionMoneyMoovModelsOpenBankingTransactionListMatch(TypedDict):
    account_id: str
    id: str


class NoFrixionMoneyMoovModelsPaymentRequired(TypedDict):
    created_by_user: dict


class NoFrixionMoneyMoovModelsPayment(NoFrixionMoneyMoovModelsPaymentRequired, total=False):
    address: list
    amount: float
    amount_pending: float
    amount_received: float
    amount_refunded: float
    auto_send_receipt: bool
    base_origin_url: str
    callback_url: str
    card_authorize_only: bool
    card_create_token: bool
    card_create_token_mode: str
    card_ignore_cvn: bool
    card_no_payer_authentication: bool
    card_processor_merchant_id: str
    card_stripe_payment_intent_id: str
    card_stripe_payment_intent_secret: str
    card_transmit_raw_detail: bool
    currency: str
    custom_field: list
    customer_email_address: str
    customer_id: str
    customer_name: str
    description: str
    destination_account: dict
    direct_debit_payment: dict
    due_date: str
    event: list
    failure_callback_url: str
    field_display_setting: list
    formatted_amount: str
    hosted_pay_checkout_url: str
    id: str
    ignore_address_verification: bool
    inserted: str
    inserted_sortable: str
    is_archived: bool
    jwk: str
    last_updated: str
    lightning_invoice: str
    lightning_invoice_expires_at: str
    merchant_direct_debit_mandate_id: str
    merchant_id: str
    merchant_token_description: str
    notification_email_address: str
    notification_role_i_d: list
    order_id: str
    partial_payment_method: str
    partial_payment_step: str
    payment_attempt: list
    payment_method: list
    payment_processor: str
    payrun_id: str
    pisp_account_id: str
    priority_bank_id: str
    result: dict
    sandbox_settle_delay_in_second: int
    shipping_address: dict
    shipping_address_city: str
    shipping_address_country_code: str
    shipping_address_county: str
    shipping_address_line1: str
    shipping_address_line2: str
    shipping_address_post_code: str
    shipping_email: str
    shipping_first_name: str
    shipping_last_name: str
    shipping_phone: str
    status: str
    success_web_hook_url: str
    tag: list
    tag_id: list
    title: str
    tokenised_card: list
    transaction: list
    use_hosted_payment_page: bool


class NoFrixionMoneyMoovModelsPaymentLoadMatch(TypedDict, total=False):
    id: str
    order_id: str


class NoFrixionMoneyMoovModelsPaymentCreateDataRequired(TypedDict):
    created_by_user: dict


class NoFrixionMoneyMoovModelsPaymentCreateData(NoFrixionMoneyMoovModelsPaymentCreateDataRequired, total=False):
    address: list
    amount: float
    amount_pending: float
    amount_received: float
    amount_refunded: float
    auto_send_receipt: bool
    base_origin_url: str
    callback_url: str
    card_authorize_only: bool
    card_create_token: bool
    card_create_token_mode: str
    card_ignore_cvn: bool
    card_no_payer_authentication: bool
    card_processor_merchant_id: str
    card_stripe_payment_intent_id: str
    card_stripe_payment_intent_secret: str
    card_transmit_raw_detail: bool
    currency: str
    custom_field: list
    customer_email_address: str
    customer_id: str
    customer_name: str
    description: str
    destination_account: dict
    direct_debit_payment: dict
    due_date: str
    event: list
    failure_callback_url: str
    field_display_setting: list
    formatted_amount: str
    hosted_pay_checkout_url: str
    id: str
    ignore_address_verification: bool
    inserted: str
    inserted_sortable: str
    is_archived: bool
    jwk: str
    last_updated: str
    lightning_invoice: str
    lightning_invoice_expires_at: str
    merchant_direct_debit_mandate_id: str
    merchant_id: str
    merchant_token_description: str
    notification_email_address: str
    notification_role_i_d: list
    order_id: str
    partial_payment_method: str
    partial_payment_step: str
    payment_attempt: list
    payment_method: list
    payment_processor: str
    payrun_id: str
    pisp_account_id: str
    priority_bank_id: str
    result: dict
    sandbox_settle_delay_in_second: int
    shipping_address: dict
    shipping_address_city: str
    shipping_address_country_code: str
    shipping_address_county: str
    shipping_address_line1: str
    shipping_address_line2: str
    shipping_address_post_code: str
    shipping_email: str
    shipping_first_name: str
    shipping_last_name: str
    shipping_phone: str
    status: str
    success_web_hook_url: str
    tag: list
    tag_id: list
    title: str
    tokenised_card: list
    transaction: list
    use_hosted_payment_page: bool


class NoFrixionMoneyMoovModelsPaymentUpdateData(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsPaymentAccountMinimalPageRequired(TypedDict):
    identifier: dict


class NoFrixionMoneyMoovModelsPaymentAccountMinimalPage(NoFrixionMoneyMoovModelsPaymentAccountMinimalPageRequired, total=False):
    account_name: str
    available_balance: float
    balance: float
    balance_minor_unit: int
    currency: str
    id: str
    is_archived: bool
    is_connected_account: bool
    merchant_id: str
    submitted_payouts_balance: float


class NoFrixionMoneyMoovModelsPaymentAccountMinimalPageListMatch(TypedDict, total=False):
    account_name: str
    available_balance: float
    balance: float
    balance_minor_unit: int
    currency: str
    id: str
    identifier: dict
    is_archived: bool
    is_connected_account: bool
    merchant_id: str
    submitted_payouts_balance: float


class NoFrixionMoneyMoovModelsPaymentAccountPageRequired(TypedDict):
    created_by: dict
    identifier: dict


class NoFrixionMoneyMoovModelsPaymentAccountPage(NoFrixionMoneyMoovModelsPaymentAccountPageRequired, total=False):
    account_name: str
    account_supplier_name: str
    available_balance: float
    available_balance_minor_unit: int
    balance: float
    balance_minor_unit: int
    bank_name: str
    consent_id: str
    created_by_display_name: str
    currency: str
    default_payment_rail: str
    display_name: str
    expiry_date: str
    external_account_icon: str
    id: str
    inserted: str
    is_archived: bool
    is_connected_account: bool
    is_default: bool
    is_trust_account: bool
    is_virtual: bool
    last_transaction: dict
    last_updated: str
    merchant_id: str
    merchant_name: str
    physical_account_id: str
    rule: list
    submitted_payouts_balance: float
    submitted_payouts_balance_minor_unit: int
    summary: str
    supplier_sepa_instant_status: str
    xero_bank_feed_connection_status: str
    xero_bank_feed_last_synced_at: str
    xero_bank_feed_sync_last_failed_at: str
    xero_bank_feed_sync_last_failure_reason: str
    xero_bank_feed_sync_status: str
    xero_unsynchronised_transactions_count: int


class NoFrixionMoneyMoovModelsPaymentAccountPageListMatch(TypedDict, total=False):
    account_id: str


class NoFrixionMoneyMoovModelsPaymentInitiation(TypedDict, total=False):
    payment_initiation_id: str
    payment_request_callback_url: str
    payment_request_id: str
    redirect_url: str
    response_type: str
    specific_error_message: str


class NoFrixionMoneyMoovModelsPaymentInitiationCreateData(TypedDict):
    paymentrequest_id: str


class NoFrixionMoneyMoovModelsPaymentRequestEventRequired(TypedDict):
    amount: float


class NoFrixionMoneyMoovModelsPaymentRequestEvent(NoFrixionMoneyMoovModelsPaymentRequestEventRequired, total=False):
    apple_pay_transaction_id: str
    card_authorization_response_id: str
    card_expiry_month: int
    card_expiry_year: int
    card_issuer: str
    card_issuer_country: str
    card_last_four_digit: str
    card_request_id: str
    card_scheme: str
    card_token_customer_id: str
    card_transaction_id: str
    currency: str
    direct_debit_payment_id: str
    direct_debit_payment_reference: str
    drirect_debit_mandate_id: str
    error_message: str
    error_reason: str
    event_type: str
    id: str
    inserted: str
    lightning_invoice: str
    lightning_r_hash: str
    origin_url: str
    payment_method_type: str
    payment_processor_name: str
    payment_request_id: str
    pisp_bank_status: str
    pisp_payment_initiation_id: str
    pisp_payment_institution_name: str
    pisp_payment_service_provider_id: str
    pisp_redirect_url: str
    reconciled_transaction_id: str
    refund_payout_id: str
    status: str
    wallet_name: str


class NoFrixionMoneyMoovModelsPaymentRequestEventListMatch(TypedDict):
    paymentrequest_id: str


class NoFrixionMoneyMoovModelsPaymentRequestMetric(TypedDict, total=False):
    all: int
    authorized: int
    paid: int
    partially_paid: int
    total_amounts_by_currency: dict
    unpaid: int


class NoFrixionMoneyMoovModelsPaymentRequestMetricLoadMatch(TypedDict, total=False):
    all: int
    authorized: int
    paid: int
    partially_paid: int
    total_amounts_by_currency: dict
    unpaid: int


class NoFrixionMoneyMoovModelsPaymentRequestMinimal(TypedDict, total=False):
    amount: float
    amount_pending: float
    amount_received: float
    amount_refunded: float
    callback_url: str
    card_stripe_payment_intent_secret: str
    country_code: str
    currency: str
    custom_fields_to_display: list
    description: str
    due_date: str
    field_display_setting: list
    google_pay_merchant_id: str
    id: str
    jwk: str
    merchant_id: str
    merchant_logo_url_png: str
    merchant_logo_url_svg: str
    merchant_name: str
    merchant_short_name: str
    partial_payment_method: str
    payment_attempt: list
    payment_methods_list: list
    payment_processor: str
    payment_processor_key: str
    pisp_error: str
    priority_bank_id: str
    status: str
    stripe_account_id: str
    title: str


class NoFrixionMoneyMoovModelsPaymentRequestMinimalListMatch(TypedDict):
    paymentrequest_id: str


class NoFrixionMoneyMoovModelsPaymentRequestResult(TypedDict, total=False):
    amount: float
    amount_pending: float
    amount_received: float
    amount_refunded: float
    currency: str
    customer_id: str
    payment: list
    payment_request_id: str
    pisp_authorization: list
    requested_amount: float
    result: str


class NoFrixionMoneyMoovModelsPaymentRequestResultListMatch(TypedDict):
    paymentrequest_id: str


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentRequired(TypedDict):
    description: str
    name: str
    template: dict


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment(NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentRequired, total=False):
    id: str
    inserted: str
    last_updated: str
    merchant_id: str


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPaymentListMatch(TypedDict):
    merchant_id: str


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Required(TypedDict):
    description: str
    name: str
    template: dict


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2(NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2Required, total=False):
    id: str
    inserted: str
    last_updated: str
    merchant_id: str


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment2LoadMatch(TypedDict):
    paymentrequest_id: str
    template_id: str


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Required(TypedDict):
    description: str
    name: str
    template: dict


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3(NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3Required, total=False):
    id: str
    inserted: str
    last_updated: str
    merchant_id: str


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment3UpdateData(TypedDict):
    paymentrequest_id: str
    template_id: str


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4(TypedDict):
    pass


class NoFrixionMoneyMoovModelsPaymentRequestsMerchantPayment4RemoveMatch(TypedDict):
    paymentrequest_id: str
    template_id: str


class NoFrixionMoneyMoovModelsPayoutKeysetPageRequired(TypedDict):
    beneficiary: dict
    source_account_identifier: dict


class NoFrixionMoneyMoovModelsPayoutKeysetPage(NoFrixionMoneyMoovModelsPayoutKeysetPageRequired, total=False):
    account_id: str
    amount: float
    amount_minor_unit: int
    approve_payout_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    can_authorise: bool
    can_process: bool
    can_update: bool
    charge_bearer: str
    created_by: str
    created_by_email_address: str
    currency: str
    current_user_id: str
    description: str
    destination: dict
    document: list
    event: list
    formatted_amount: str
    formatted_fx_destination_amount: str
    formatted_schedule: str
    formatted_schedule_day_only: str
    formatted_source_account_available_balance: str
    fx_destination_amount: float
    fx_destination_amount_minor_unit: int
    fx_destination_currency: str
    fx_quote_expires_at: str
    fx_quote_id: str
    fx_rate: float
    fx_use_destination_amount: bool
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice_id: str
    is_archived: bool
    is_failed: bool
    is_settled: bool
    is_submitted: bool
    last_updated: str
    merchant_id: str
    merchant_token_description: str
    nonce: str
    payment_processor: str
    payment_rail: str
    payrun_id: str
    payrun_name: str
    rule: dict
    schedule_date: str
    scheduled: bool
    source_account_available_balance: float
    source_account_available_balance_minor_unit: int
    source_account_bic: str
    source_account_currency: str
    source_account_iban: str
    source_account_name: str
    source_account_number: str
    source_account_sortcode: str
    status: str
    tag: list
    their_reference: str
    topup_payrun_id: str
    transacted_amount: float
    transacted_fx_amount: float
    transacted_fx_rate: float
    type: str
    user_id: str
    your_reference: str


class NoFrixionMoneyMoovModelsPayoutKeysetPageListMatch(TypedDict, total=False):
    account_id: str
    merchant_id: str


class NoFrixionMoneyMoovModelsPayoutMetric(TypedDict, total=False):
    all: float
    failed: float
    in_progress: float
    paid: float
    pending_approval: float
    scheduled: float
    total_amounts_by_currency: dict


class NoFrixionMoneyMoovModelsPayoutMetricLoadMatch(TypedDict, total=False):
    all: float
    failed: float
    in_progress: float
    paid: float
    pending_approval: float
    scheduled: float
    total_amounts_by_currency: dict


class NoFrixionMoneyMoovModelsPayoutsPayoutsCreate(TypedDict, total=False):
    failed_payout: dict
    payout: list


class NoFrixionMoneyMoovModelsPayoutsPayoutsCreateCreateData(TypedDict, total=False):
    failed_payout: dict
    payout: list


class NoFrixionMoneyMoovModelsPayrunRequired(TypedDict):
    last_updated_by: dict


class NoFrixionMoneyMoovModelsPayrun(NoFrixionMoneyMoovModelsPayrunRequired, total=False):
    authorisation: list
    authorisation_date: str
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    can_authorise: bool
    can_delete: bool
    can_edit: bool
    event: list
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice: list
    invoices_minimal: list
    is_archived: bool
    last_updated: str
    merchant_id: str
    name: str
    nonce: str
    payment: list
    payout: list
    payouts_count: int
    reason: str
    schedule_date: str
    scheduled_date: str
    source_account: list
    status: str
    total_eur: float
    total_gbp: float
    total_usd: float


class NoFrixionMoneyMoovModelsPayrunLoadMatch(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsPayrunCreateData(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsPayrunUpdateData(TypedDict, total=False):
    id: str
    payrun_id: str


class NoFrixionMoneyMoovModelsReportResult(TypedDict, total=False):
    content: str
    content_type: str
    last_completed_at: str
    merchant_id: str
    report_name: str
    report_type: str
    statement_number: int


class NoFrixionMoneyMoovModelsReportResultLoadMatch(TypedDict):
    id: int
    report_id: str


class NoFrixionMoneyMoovModelsRuleRequired(TypedDict):
    created_by: dict
    nonce: str


class NoFrixionMoneyMoovModelsRule(NoFrixionMoneyMoovModelsRuleRequired, total=False):
    account: dict
    account_id: str
    approve_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    can_authorise: bool
    description: str
    end_at: str
    has_current_user_authorised: bool
    id: str
    inserted: str
    is_disabled: bool
    last_executed_at: str
    last_run_at_transaction_date: str
    last_updated: str
    merchant_id: str
    name: str
    on_approved_web_hook_url: str
    on_execution_error_web_hook_url: str
    on_execution_success_web_hook_url: str
    start_at: str
    status: str
    sweep_action: dict
    time_zone_id: str
    trigger_cron_expression: str
    trigger_on_pay_in: bool
    user_id: str
    web_hook_secret: str


class NoFrixionMoneyMoovModelsRuleLoadMatch(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsRuleCreateDataRequired(TypedDict):
    created_by: dict
    nonce: str


class NoFrixionMoneyMoovModelsRuleCreateData(NoFrixionMoneyMoovModelsRuleCreateDataRequired, total=False):
    account: dict
    account_id: str
    approve_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    can_authorise: bool
    description: str
    end_at: str
    has_current_user_authorised: bool
    id: str
    inserted: str
    is_disabled: bool
    last_executed_at: str
    last_run_at_transaction_date: str
    last_updated: str
    merchant_id: str
    name: str
    on_approved_web_hook_url: str
    on_execution_error_web_hook_url: str
    on_execution_success_web_hook_url: str
    start_at: str
    status: str
    sweep_action: dict
    time_zone_id: str
    trigger_cron_expression: str
    trigger_on_pay_in: bool
    user_id: str
    web_hook_secret: str


class NoFrixionMoneyMoovModelsRuleUpdateData(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsTransaction(TypedDict, total=False):
    account_id: str
    account_name: str
    account_sequence_number: int
    amount: float
    amount_minor_unit: int
    balance: float
    balance_minor_unit: int
    counterparty: dict
    counterparty_summary: str
    currency: str
    description: str
    fx_amount: float
    fx_currency: str
    fx_rate: float
    id: str
    inserted: str
    merchant_id: str
    payment_request_custom_field: dict
    payment_request_id: str
    payout_id: str
    raw_reference: str
    rule_id: str
    tag: list
    their_reference: str
    transaction_date: str
    type: str
    virtual_iban: str
    your_reference: str


class NoFrixionMoneyMoovModelsTransactionLoadMatchRequired(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsTransactionLoadMatch(NoFrixionMoneyMoovModelsTransactionLoadMatchRequired, total=False):
    account_id: str


class NoFrixionMoneyMoovModelsTransactionPage(TypedDict, total=False):
    account_id: str
    account_name: str
    account_sequence_number: int
    amount: float
    amount_minor_unit: int
    balance: float
    balance_minor_unit: int
    content: list
    counterparty: dict
    counterparty_summary: str
    currency: str
    description: str
    fx_amount: float
    fx_currency: str
    fx_rate: float
    id: str
    inserted: str
    merchant_id: str
    page_number: int
    page_size: int
    payment_request_custom_field: dict
    payment_request_id: str
    payout_id: str
    raw_reference: str
    rule_id: str
    tag: list
    their_reference: str
    total_page: int
    total_size: int
    transaction_date: str
    type: str
    virtual_iban: str
    your_reference: str


class NoFrixionMoneyMoovModelsTransactionPageLoadMatch(TypedDict):
    account_id: str


class NoFrixionMoneyMoovModelsTransactionPageListMatch(TypedDict, total=False):
    account_id: str
    merchant_id: str


class NoFrixionMoneyMoovModelsUserInviteRequired(TypedDict):
    user: dict


class NoFrixionMoneyMoovModelsUserInvite(NoFrixionMoneyMoovModelsUserInviteRequired, total=False):
    authorisation_status: dict
    id: str
    initial_role_id: str
    invitee_email_address: str
    invitee_first_name: str
    invitee_last_name: str
    inviter_email_address: str
    inviter_first_name: str
    inviter_last_name: str
    is_authorised: bool
    is_invitee_registered: bool
    last_invited: str
    merchant_id: str
    merchant_name: str
    message: str
    registration_url: str
    send_invite_email: bool
    status: str
    user_id: str


class NoFrixionMoneyMoovModelsUserInviteLoadMatch(TypedDict, total=False):
    id: str
    userinvite_id: str


class NoFrixionMoneyMoovModelsUserInviteCreateDataRequired(TypedDict):
    user: dict


class NoFrixionMoneyMoovModelsUserInviteCreateData(NoFrixionMoneyMoovModelsUserInviteCreateDataRequired, total=False):
    authorisation_status: dict
    id: str
    initial_role_id: str
    invitee_email_address: str
    invitee_first_name: str
    invitee_last_name: str
    inviter_email_address: str
    inviter_first_name: str
    inviter_last_name: str
    is_authorised: bool
    is_invitee_registered: bool
    last_invited: str
    merchant_id: str
    merchant_name: str
    message: str
    registration_url: str
    send_invite_email: bool
    status: str
    user_id: str


class NoFrixionMoneyMoovModelsUserInvitePageRequired(TypedDict):
    user: dict


class NoFrixionMoneyMoovModelsUserInvitePage(NoFrixionMoneyMoovModelsUserInvitePageRequired, total=False):
    authorisation_status: dict
    id: str
    initial_role_id: str
    invitee_email_address: str
    invitee_first_name: str
    invitee_last_name: str
    inviter_email_address: str
    inviter_first_name: str
    inviter_last_name: str
    is_authorised: bool
    is_invitee_registered: bool
    last_invited: str
    merchant_id: str
    merchant_name: str
    message: str
    registration_url: str
    status: str
    user_id: str


class NoFrixionMoneyMoovModelsUserInvitePageListMatch(TypedDict):
    merchant_id: str


class NoFrixionMoneyMoovModelsUserPageRequired(TypedDict):
    email_address: str
    first_name: str
    last_name: str


class NoFrixionMoneyMoovModelsUserPage(NoFrixionMoneyMoovModelsUserPageRequired, total=False):
    client_session_timeout: list
    id: str
    passkey_added: bool
    permission: dict
    roles_with_scope: list
    two_factor_enabled: bool


class NoFrixionMoneyMoovModelsUserPageListMatch(TypedDict):
    merchant_id: str


class NoFrixionMoneyMoovModelsWebhook(TypedDict, total=False):
    destination_url: str
    email_address: str
    failed_notification_email_address: str
    id: str
    is_active: bool
    merchant_id: str
    notification_method: str
    resource_type: list
    retry: bool
    secret: str
    version: int


class NoFrixionMoneyMoovModelsWebhookLoadMatchRequired(TypedDict):
    id: str


class NoFrixionMoneyMoovModelsWebhookLoadMatch(NoFrixionMoneyMoovModelsWebhookLoadMatchRequired, total=False):
    merchant_id: str


class NoFrixionMoneyMoovModelsWebhookListMatch(TypedDict):
    merchant_id: str


class NoFrixionMoneyMoovModelsWebhookCreateData(TypedDict, total=False):
    destination_url: str
    email_address: str
    failed_notification_email_address: str
    id: str
    is_active: bool
    merchant_id: str
    notification_method: str
    resource_type: list
    retry: bool
    secret: str
    version: int


class NoFrixionMoneyMoovModelsWebhookUpdateData(TypedDict):
    id: str


class OpenBanking(TypedDict):
    pass


class OpenBankingCreateData(TypedDict):
    account_id: str


class OpenBankingRemoveMatch(TypedDict, total=False):
    email: str
    merchant_id: str
    account_id: str


class PayeeverificationRequired(TypedDict):
    account_name: str
    iban: str


class Payeeverification(PayeeverificationRequired, total=False):
    account_number: str
    payee_verified_account_name: str
    result: str
    secondary_identification: str
    sort_code: str


class PayeeverificationCreateDataRequired(TypedDict):
    account_name: str
    iban: str


class PayeeverificationCreateData(PayeeverificationCreateDataRequired, total=False):
    account_number: str
    payee_verified_account_name: str
    result: str
    secondary_identification: str
    sort_code: str


class PaymentRequest(TypedDict, total=False):
    amount: float
    do_simulate_settlement_failure: bool
    error_description: str
    institution: str
    payment_initiation_id: str


class PaymentRequestLoadMatch(TypedDict, total=False):
    paymentrequest_id: str


class PaymentRequestCreateData(TypedDict):
    paymentrequest_id: str


class PaymentRequestUpdateData(TypedDict):
    paymentrequest_id: str


class PaymentRequestRemoveMatch(TypedDict):
    id: str


class PayoutRequired(TypedDict):
    beneficiary: dict
    source_account_identifier: dict


class Payout(PayoutRequired, total=False):
    account_id: str
    allow_incomplete: bool
    amount: float
    amount_minor_unit: int
    approve_payout_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    beneficiary_id: str
    can_authorise: bool
    can_process: bool
    can_update: bool
    charge_bearer: str
    created_by: str
    created_by_email_address: str
    currency: str
    current_user_id: str
    description: str
    destination: dict
    document: list
    event: list
    formatted_amount: str
    formatted_fx_destination_amount: str
    formatted_schedule: str
    formatted_schedule_day_only: str
    formatted_source_account_available_balance: str
    fx_destination_amount: float
    fx_destination_amount_minor_unit: int
    fx_destination_currency: str
    fx_quote_expires_at: str
    fx_quote_id: str
    fx_rate: float
    fx_use_destination_amount: bool
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice_id: str
    is_archived: bool
    is_failed: bool
    is_settled: bool
    is_submitted: bool
    last_updated: str
    merchant_id: str
    merchant_token_description: str
    nonce: str
    payment_processor: str
    payment_rail: str
    payrun_id: str
    payrun_name: str
    rule: dict
    schedule_date: str
    scheduled: bool
    source_account_available_balance: float
    source_account_available_balance_minor_unit: int
    source_account_bic: str
    source_account_currency: str
    source_account_iban: str
    source_account_name: str
    source_account_number: str
    source_account_sortcode: str
    status: str
    tag: list
    tag_id: list
    their_reference: str
    topup_payrun_id: str
    transacted_amount: float
    transacted_fx_amount: float
    transacted_fx_rate: float
    type: str
    user_id: str
    your_reference: str


class PayoutLoadMatch(TypedDict, total=False):
    amount: float
    destination: str
    source: str
    id: str


class PayoutCreateData(TypedDict, total=False):
    id: str


class PayoutUpdateData(TypedDict):
    id: str


class PayoutRemoveMatch(TypedDict):
    id: str


class Payrun(TypedDict, total=False):
    id: str
    note: str
    scheduled_date: str


class PayrunCreateData(TypedDict):
    id: str


class PayrunUpdateData(TypedDict):
    id: str


class PayrunRemoveMatch(TypedDict):
    id: str


class RejectRequired(TypedDict):
    beneficiary: dict
    source_account_identifier: dict


class Reject(RejectRequired, total=False):
    account_id: str
    amount: float
    amount_minor_unit: int
    approve_payout_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    can_authorise: bool
    can_process: bool
    can_update: bool
    charge_bearer: str
    created_by: str
    created_by_email_address: str
    currency: str
    current_user_id: str
    description: str
    destination: dict
    document: list
    event: list
    formatted_amount: str
    formatted_fx_destination_amount: str
    formatted_schedule: str
    formatted_schedule_day_only: str
    formatted_source_account_available_balance: str
    fx_destination_amount: float
    fx_destination_amount_minor_unit: int
    fx_destination_currency: str
    fx_quote_expires_at: str
    fx_quote_id: str
    fx_rate: float
    fx_use_destination_amount: bool
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice_id: str
    is_archived: bool
    is_failed: bool
    is_settled: bool
    is_submitted: bool
    last_updated: str
    merchant_id: str
    merchant_token_description: str
    nonce: str
    payment_processor: str
    payment_rail: str
    payrun_id: str
    payrun_name: str
    reason: str
    rule: dict
    schedule_date: str
    scheduled: bool
    source_account_available_balance: float
    source_account_available_balance_minor_unit: int
    source_account_bic: str
    source_account_currency: str
    source_account_iban: str
    source_account_name: str
    source_account_number: str
    source_account_sortcode: str
    status: str
    tag: list
    their_reference: str
    topup_payrun_id: str
    transacted_amount: float
    transacted_fx_amount: float
    transacted_fx_rate: float
    type: str
    user_id: str
    your_reference: str


class RejectUpdateData(TypedDict):
    id: str


class Report(TypedDict):
    pass


class ReportUpdateData(TypedDict):
    id: str


class Rule(TypedDict):
    pass


class RuleUpdateData(TypedDict):
    id: str


class RuleRemoveMatch(TypedDict):
    id: str


class SendRequired(TypedDict):
    beneficiary: dict
    source_account_identifier: dict


class Send(SendRequired, total=False):
    account_id: str
    allow_incomplete: bool
    amount: float
    amount_minor_unit: int
    approve_payout_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    beneficiary_id: str
    can_authorise: bool
    can_process: bool
    can_update: bool
    charge_bearer: str
    created_by: str
    created_by_email_address: str
    currency: str
    current_user_id: str
    description: str
    destination: dict
    document: list
    event: list
    formatted_amount: str
    formatted_fx_destination_amount: str
    formatted_schedule: str
    formatted_schedule_day_only: str
    formatted_source_account_available_balance: str
    fx_destination_amount: float
    fx_destination_amount_minor_unit: int
    fx_destination_currency: str
    fx_quote_expires_at: str
    fx_quote_id: str
    fx_rate: float
    fx_use_destination_amount: bool
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice_id: str
    is_archived: bool
    is_failed: bool
    is_settled: bool
    is_submitted: bool
    last_updated: str
    merchant_id: str
    merchant_token_description: str
    nonce: str
    payment_processor: str
    payment_rail: str
    payrun_id: str
    payrun_name: str
    rule: dict
    schedule_date: str
    scheduled: bool
    source_account_available_balance: float
    source_account_available_balance_minor_unit: int
    source_account_bic: str
    source_account_currency: str
    source_account_iban: str
    source_account_name: str
    source_account_number: str
    source_account_sortcode: str
    status: str
    tag: list
    tag_id: list
    their_reference: str
    topup_payrun_id: str
    transacted_amount: float
    transacted_fx_amount: float
    transacted_fx_rate: float
    type: str
    user_id: str
    your_reference: str


class SendCreateDataRequired(TypedDict):
    beneficiary: dict
    source_account_identifier: dict


class SendCreateData(SendCreateDataRequired, total=False):
    account_id: str
    allow_incomplete: bool
    amount: float
    amount_minor_unit: int
    approve_payout_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    beneficiary_id: str
    can_authorise: bool
    can_process: bool
    can_update: bool
    charge_bearer: str
    created_by: str
    created_by_email_address: str
    currency: str
    current_user_id: str
    description: str
    destination: dict
    document: list
    event: list
    formatted_amount: str
    formatted_fx_destination_amount: str
    formatted_schedule: str
    formatted_schedule_day_only: str
    formatted_source_account_available_balance: str
    fx_destination_amount: float
    fx_destination_amount_minor_unit: int
    fx_destination_currency: str
    fx_quote_expires_at: str
    fx_quote_id: str
    fx_rate: float
    fx_use_destination_amount: bool
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice_id: str
    is_archived: bool
    is_failed: bool
    is_settled: bool
    is_submitted: bool
    last_updated: str
    merchant_id: str
    merchant_token_description: str
    nonce: str
    payment_processor: str
    payment_rail: str
    payrun_id: str
    payrun_name: str
    rule: dict
    schedule_date: str
    scheduled: bool
    source_account_available_balance: float
    source_account_available_balance_minor_unit: int
    source_account_bic: str
    source_account_currency: str
    source_account_iban: str
    source_account_name: str
    source_account_number: str
    source_account_sortcode: str
    status: str
    tag: list
    tag_id: list
    their_reference: str
    topup_payrun_id: str
    transacted_amount: float
    transacted_fx_amount: float
    transacted_fx_rate: float
    type: str
    user_id: str
    your_reference: str


class SendbeneficiaryRequired(TypedDict):
    beneficiary: dict
    source_account_identifier: dict


class Sendbeneficiary(SendbeneficiaryRequired, total=False):
    account_id: str
    allow_incomplete: bool
    amount: float
    amount_minor_unit: int
    approve_payout_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    beneficiary_id: str
    can_authorise: bool
    can_process: bool
    can_update: bool
    charge_bearer: str
    created_by: str
    created_by_email_address: str
    currency: str
    current_user_id: str
    description: str
    destination: dict
    document: list
    event: list
    formatted_amount: str
    formatted_fx_destination_amount: str
    formatted_schedule: str
    formatted_schedule_day_only: str
    formatted_source_account_available_balance: str
    fx_destination_amount: float
    fx_destination_amount_minor_unit: int
    fx_destination_currency: str
    fx_quote_expires_at: str
    fx_quote_id: str
    fx_rate: float
    fx_use_destination_amount: bool
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice_id: str
    is_archived: bool
    is_failed: bool
    is_settled: bool
    is_submitted: bool
    last_updated: str
    merchant_id: str
    merchant_token_description: str
    nonce: str
    payment_processor: str
    payment_rail: str
    payrun_id: str
    payrun_name: str
    rule: dict
    schedule_date: str
    scheduled: bool
    source_account_available_balance: float
    source_account_available_balance_minor_unit: int
    source_account_bic: str
    source_account_currency: str
    source_account_iban: str
    source_account_name: str
    source_account_number: str
    source_account_sortcode: str
    status: str
    tag: list
    tag_id: list
    their_reference: str
    topup_payrun_id: str
    transacted_amount: float
    transacted_fx_amount: float
    transacted_fx_rate: float
    type: str
    user_id: str
    your_reference: str


class SendbeneficiaryCreateDataRequired(TypedDict):
    beneficiary: dict
    source_account_identifier: dict


class SendbeneficiaryCreateData(SendbeneficiaryCreateDataRequired, total=False):
    account_id: str
    allow_incomplete: bool
    amount: float
    amount_minor_unit: int
    approve_payout_url: str
    approver_id: str
    authentication_method: list
    authorisation: list
    authorisers_completed_count: int
    authorisers_required_count: int
    batch_payout_id: str
    beneficiary_id: str
    can_authorise: bool
    can_process: bool
    can_update: bool
    charge_bearer: str
    created_by: str
    created_by_email_address: str
    currency: str
    current_user_id: str
    description: str
    destination: dict
    document: list
    event: list
    formatted_amount: str
    formatted_fx_destination_amount: str
    formatted_schedule: str
    formatted_schedule_day_only: str
    formatted_source_account_available_balance: str
    fx_destination_amount: float
    fx_destination_amount_minor_unit: int
    fx_destination_currency: str
    fx_quote_expires_at: str
    fx_quote_id: str
    fx_rate: float
    fx_use_destination_amount: bool
    has_current_user_authorised: bool
    id: str
    inserted: str
    invoice_id: str
    is_archived: bool
    is_failed: bool
    is_settled: bool
    is_submitted: bool
    last_updated: str
    merchant_id: str
    merchant_token_description: str
    nonce: str
    payment_processor: str
    payment_rail: str
    payrun_id: str
    payrun_name: str
    rule: dict
    schedule_date: str
    scheduled: bool
    source_account_available_balance: float
    source_account_available_balance_minor_unit: int
    source_account_bic: str
    source_account_currency: str
    source_account_iban: str
    source_account_name: str
    source_account_number: str
    source_account_sortcode: str
    status: str
    tag: list
    tag_id: list
    their_reference: str
    topup_payrun_id: str
    transacted_amount: float
    transacted_fx_amount: float
    transacted_fx_rate: float
    type: str
    user_id: str
    your_reference: str


class TagRequired(TypedDict):
    merchant_id: str
    name: str


class Tag(TagRequired, total=False):
    colour_hex: str
    description: str
    id: str


class TagListMatch(TypedDict):
    merchant_id: str


class TagCreateData(TypedDict):
    merchant_id: str


class Token(TypedDict):
    pass


class TokenCreateData(TypedDict):
    id: str


class TokenRemoveMatch(TypedDict):
    id: str


class Transaction(TypedDict):
    pass


class TransactionLoadMatch(TypedDict):
    sequence_number: int
    transaction_id: str


class TransactionCreateData(TypedDict):
    id: str


class TransactionRemoveMatch(TypedDict):
    id: str


class UserRequired(TypedDict):
    email_address: str
    first_name: str
    last_name: str


class User(UserRequired, total=False):
    client_session_timeout: list
    id: str
    passkey_added: bool
    permission: dict
    profile: str
    roles_with_scope: list
    two_factor_enabled: bool
    user_invite_id: str


class UserListMatch(TypedDict, total=False):
    merchant_id: str


class UserUpdateData(TypedDict):
    id: str


class UserInvite(TypedDict):
    pass


class UserInviteCreateData(TypedDict):
    id: str


class UserInviteUpdateData(TypedDict):
    id: str


class UserInviteRemoveMatch(TypedDict):
    id: str


class VirtualRequired(TypedDict):
    created_by: dict
    identifier: dict
    name: str


class Virtual(VirtualRequired, total=False):
    account_name: str
    account_supplier_name: str
    available_balance: float
    available_balance_minor_unit: int
    balance: float
    balance_minor_unit: int
    bank_name: str
    consent_id: str
    created_by_display_name: str
    currency: str
    default_payment_rail: str
    display_name: str
    expiry_date: str
    external_account_icon: str
    id: str
    inserted: str
    is_archived: bool
    is_connected_account: bool
    is_default: bool
    is_trust_account: bool
    is_virtual: bool
    last_transaction: dict
    last_updated: str
    merchant_id: str
    merchant_name: str
    physical_account_id: str
    rule: list
    submitted_payouts_balance: float
    submitted_payouts_balance_minor_unit: int
    summary: str
    supplier_sepa_instant_status: str
    xero_bank_feed_connection_status: str
    xero_bank_feed_last_synced_at: str
    xero_bank_feed_sync_last_failed_at: str
    xero_bank_feed_sync_last_failure_reason: str
    xero_bank_feed_sync_status: str
    xero_unsynchronised_transactions_count: int


class VirtualCreateData(TypedDict):
    account_id: str


class VirtualUpdateData(TypedDict):
    account_id: str
    id: str


class Webhook(TypedDict):
    pass


class WebhookRemoveMatch(TypedDict):
    id: str


class WhoamiRequired(TypedDict):
    email_address: str
    first_name: str
    last_name: str


class Whoami(WhoamiRequired, total=False):
    client_session_timeout: list
    id: str
    passkey_added: bool
    permission: dict
    roles_with_scope: list
    two_factor_enabled: bool


class WhoamiListMatch(TypedDict, total=False):
    client_session_timeout: list
    email_address: str
    first_name: str
    id: str
    last_name: str
    passkey_added: bool
    permission: dict
    roles_with_scope: list
    two_factor_enabled: bool


class WhoamitrustedappRequired(TypedDict):
    email_address: str
    first_name: str
    last_name: str


class Whoamitrustedapp(WhoamitrustedappRequired, total=False):
    client_session_timeout: list
    id: str
    passkey_added: bool
    permission: dict
    roles_with_scope: list
    two_factor_enabled: bool


class WhoamitrustedappListMatch(TypedDict, total=False):
    client_session_timeout: list
    email_address: str
    first_name: str
    id: str
    last_name: str
    passkey_added: bool
    permission: dict
    roles_with_scope: list
    two_factor_enabled: bool
