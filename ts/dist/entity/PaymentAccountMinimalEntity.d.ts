import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PaymentAccountMinimal, PaymentAccountMinimalListMatch } from '../NofrixionTypes';
declare class PaymentAccountMinimalEntity extends NofrixionEntityBase<PaymentAccountMinimal> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PaymentAccountMinimalEntity): PaymentAccountMinimalEntity;
    list(this: any, reqmatch?: PaymentAccountMinimalListMatch, ctrl?: Control): Promise<PaymentAccountMinimalEntity[]>;
}
export { PaymentAccountMinimalEntity };
