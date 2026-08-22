import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Payment, PaymentLoadMatch, PaymentCreateData, PaymentUpdateData } from '../NofrixionTypes';
declare class PaymentEntity extends NofrixionEntityBase<Payment> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PaymentEntity): PaymentEntity;
    load(this: any, reqmatch?: PaymentLoadMatch, ctrl?: Control): Promise<PaymentEntity>;
    create(this: any, reqdata?: PaymentCreateData, ctrl?: Control): Promise<PaymentEntity>;
    update(this: any, reqdata?: PaymentUpdateData, ctrl?: Control): Promise<PaymentEntity>;
}
export { PaymentEntity };
