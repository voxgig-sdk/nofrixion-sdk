import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { CardPayment, CardPaymentCreateData } from '../NofrixionTypes';
declare class CardPaymentEntity extends NofrixionEntityBase<CardPayment> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: CardPaymentEntity): CardPaymentEntity;
    create(this: any, reqdata?: CardPaymentCreateData, ctrl?: Control): Promise<CardPaymentEntity>;
}
export { CardPaymentEntity };
