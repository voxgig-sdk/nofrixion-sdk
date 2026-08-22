import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PayoutKeysetPage, PayoutKeysetPageListMatch } from '../NofrixionTypes';
declare class PayoutKeysetPageEntity extends NofrixionEntityBase<PayoutKeysetPage> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PayoutKeysetPageEntity): PayoutKeysetPageEntity;
    list(this: any, reqmatch?: PayoutKeysetPageListMatch, ctrl?: Control): Promise<PayoutKeysetPageEntity[]>;
}
export { PayoutKeysetPageEntity };
