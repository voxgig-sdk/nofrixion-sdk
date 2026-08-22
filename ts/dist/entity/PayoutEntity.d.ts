import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Payout, PayoutLoadMatch, PayoutListMatch, PayoutCreateData, PayoutUpdateData, PayoutRemoveMatch } from '../NofrixionTypes';
declare class PayoutEntity extends NofrixionEntityBase<Payout> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PayoutEntity): PayoutEntity;
    load(this: any, reqmatch?: PayoutLoadMatch, ctrl?: Control): Promise<PayoutEntity>;
    list(this: any, reqmatch?: PayoutListMatch, ctrl?: Control): Promise<PayoutEntity[]>;
    create(this: any, reqdata?: PayoutCreateData, ctrl?: Control): Promise<PayoutEntity>;
    update(this: any, reqdata?: PayoutUpdateData, ctrl?: Control): Promise<PayoutEntity>;
    remove(this: any, reqmatch?: PayoutRemoveMatch, ctrl?: Control): Promise<PayoutEntity>;
}
export { PayoutEntity };
