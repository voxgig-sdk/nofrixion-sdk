import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PaymentRequestResult, PaymentRequestResultListMatch } from '../NofrixionTypes';
declare class PaymentRequestResultEntity extends NofrixionEntityBase<PaymentRequestResult> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PaymentRequestResultEntity): PaymentRequestResultEntity;
    list(this: any, reqmatch?: PaymentRequestResultListMatch, ctrl?: Control): Promise<PaymentRequestResultEntity[]>;
}
export { PaymentRequestResultEntity };
