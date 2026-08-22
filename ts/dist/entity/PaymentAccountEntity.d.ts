import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PaymentAccount, PaymentAccountListMatch } from '../NofrixionTypes';
declare class PaymentAccountEntity extends NofrixionEntityBase<PaymentAccount> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PaymentAccountEntity): PaymentAccountEntity;
    list(this: any, reqmatch?: PaymentAccountListMatch, ctrl?: Control): Promise<PaymentAccountEntity[]>;
}
export { PaymentAccountEntity };
