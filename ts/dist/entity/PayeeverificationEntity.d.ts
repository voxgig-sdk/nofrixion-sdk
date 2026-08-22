import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Payeeverification, PayeeverificationCreateData } from '../NofrixionTypes';
declare class PayeeverificationEntity extends NofrixionEntityBase<Payeeverification> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PayeeverificationEntity): PayeeverificationEntity;
    create(this: any, reqdata?: PayeeverificationCreateData, ctrl?: Control): Promise<PayeeverificationEntity>;
}
export { PayeeverificationEntity };
