import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PaymentInitiation, PaymentInitiationCreateData } from '../NofrixionTypes';
declare class PaymentInitiationEntity extends NofrixionEntityBase<PaymentInitiation> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PaymentInitiationEntity): PaymentInitiationEntity;
    create(this: any, reqdata?: PaymentInitiationCreateData, ctrl?: Control): Promise<PaymentInitiationEntity>;
}
export { PaymentInitiationEntity };
