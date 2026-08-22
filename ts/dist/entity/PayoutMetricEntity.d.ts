import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PayoutMetric, PayoutMetricLoadMatch } from '../NofrixionTypes';
declare class PayoutMetricEntity extends NofrixionEntityBase<PayoutMetric> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PayoutMetricEntity): PayoutMetricEntity;
    load(this: any, reqmatch?: PayoutMetricLoadMatch, ctrl?: Control): Promise<PayoutMetricEntity>;
}
export { PayoutMetricEntity };
