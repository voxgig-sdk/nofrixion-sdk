import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Payrun, PayrunLoadMatch, PayrunListMatch, PayrunCreateData, PayrunUpdateData, PayrunRemoveMatch } from '../NofrixionTypes';
declare class PayrunEntity extends NofrixionEntityBase<Payrun> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PayrunEntity): PayrunEntity;
    load(this: any, reqmatch?: PayrunLoadMatch, ctrl?: Control): Promise<PayrunEntity>;
    list(this: any, reqmatch?: PayrunListMatch, ctrl?: Control): Promise<PayrunEntity[]>;
    create(this: any, reqdata?: PayrunCreateData, ctrl?: Control): Promise<PayrunEntity>;
    update(this: any, reqdata?: PayrunUpdateData, ctrl?: Control): Promise<PayrunEntity>;
    remove(this: any, reqmatch?: PayrunRemoveMatch, ctrl?: Control): Promise<PayrunEntity>;
}
export { PayrunEntity };
