import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { PaymentRequest, PaymentRequestLoadMatch, PaymentRequestListMatch, PaymentRequestCreateData, PaymentRequestUpdateData, PaymentRequestRemoveMatch } from '../NofrixionTypes';
declare class PaymentRequestEntity extends NofrixionEntityBase<PaymentRequest> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: PaymentRequestEntity): PaymentRequestEntity;
    load(this: any, reqmatch?: PaymentRequestLoadMatch, ctrl?: Control): Promise<PaymentRequestEntity>;
    list(this: any, reqmatch?: PaymentRequestListMatch, ctrl?: Control): Promise<PaymentRequestEntity[]>;
    create(this: any, reqdata?: PaymentRequestCreateData, ctrl?: Control): Promise<PaymentRequestEntity>;
    update(this: any, reqdata?: PaymentRequestUpdateData, ctrl?: Control): Promise<PaymentRequestEntity>;
    remove(this: any, reqmatch?: PaymentRequestRemoveMatch, ctrl?: Control): Promise<PaymentRequestEntity>;
}
export { PaymentRequestEntity };
