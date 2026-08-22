import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PaymentRequestMinimal, PaymentRequestMinimalListMatch } from '../NofrixionTypes';
declare class PaymentRequestMinimalEntity extends NofrixionEntityBase<PaymentRequestMinimal> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PaymentRequestMinimalEntity): PaymentRequestMinimalEntity;
    list(this: any, reqmatch?: PaymentRequestMinimalListMatch, ctrl?: Control): Promise<PaymentRequestMinimalEntity[]>;
}
export { PaymentRequestMinimalEntity };
