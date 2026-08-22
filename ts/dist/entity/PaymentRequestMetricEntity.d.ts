import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PaymentRequestMetric, PaymentRequestMetricLoadMatch } from '../NofrixionTypes';
declare class PaymentRequestMetricEntity extends NofrixionEntityBase<PaymentRequestMetric> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PaymentRequestMetricEntity): PaymentRequestMetricEntity;
    load(this: any, reqmatch?: PaymentRequestMetricLoadMatch, ctrl?: Control): Promise<PaymentRequestMetricEntity>;
}
export { PaymentRequestMetricEntity };
