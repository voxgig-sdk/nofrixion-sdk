import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { IPayment, IPaymentCreateData } from '../NofrixionTypes';
declare class IPaymentEntity extends NofrixionEntityBase<IPayment> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: IPaymentEntity): IPaymentEntity;
    create(this: any, reqdata?: IPaymentCreateData, ctrl?: Control): Promise<IPaymentEntity>;
}
export { IPaymentEntity };
