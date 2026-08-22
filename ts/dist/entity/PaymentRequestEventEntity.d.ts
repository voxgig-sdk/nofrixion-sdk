import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PaymentRequestEvent, PaymentRequestEventListMatch } from '../NofrixionTypes';
declare class PaymentRequestEventEntity extends NofrixionEntityBase<PaymentRequestEvent> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PaymentRequestEventEntity): PaymentRequestEventEntity;
    list(this: any, reqmatch?: PaymentRequestEventListMatch, ctrl?: Control): Promise<PaymentRequestEventEntity[]>;
}
export { PaymentRequestEventEntity };
