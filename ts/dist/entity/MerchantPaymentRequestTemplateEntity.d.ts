import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { MerchantPaymentRequestTemplate, MerchantPaymentRequestTemplateLoadMatch, MerchantPaymentRequestTemplateListMatch, MerchantPaymentRequestTemplateUpdateData, MerchantPaymentRequestTemplateRemoveMatch } from '../NofrixionTypes';
declare class MerchantPaymentRequestTemplateEntity extends NofrixionEntityBase<MerchantPaymentRequestTemplate> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: MerchantPaymentRequestTemplateEntity): MerchantPaymentRequestTemplateEntity;
    load(this: any, reqmatch?: MerchantPaymentRequestTemplateLoadMatch, ctrl?: Control): Promise<MerchantPaymentRequestTemplateEntity>;
    list(this: any, reqmatch?: MerchantPaymentRequestTemplateListMatch, ctrl?: Control): Promise<MerchantPaymentRequestTemplateEntity[]>;
    update(this: any, reqdata?: MerchantPaymentRequestTemplateUpdateData, ctrl?: Control): Promise<MerchantPaymentRequestTemplateEntity>;
    remove(this: any, reqmatch?: MerchantPaymentRequestTemplateRemoveMatch, ctrl?: Control): Promise<MerchantPaymentRequestTemplateEntity>;
}
export { MerchantPaymentRequestTemplateEntity };
