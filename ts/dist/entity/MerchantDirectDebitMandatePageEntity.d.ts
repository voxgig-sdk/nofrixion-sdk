import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { MerchantDirectDebitMandatePage, MerchantDirectDebitMandatePageListMatch } from '../NofrixionTypes';
declare class MerchantDirectDebitMandatePageEntity extends NofrixionEntityBase<MerchantDirectDebitMandatePage> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: MerchantDirectDebitMandatePageEntity): MerchantDirectDebitMandatePageEntity;
    list(this: any, reqmatch?: MerchantDirectDebitMandatePageListMatch, ctrl?: Control): Promise<MerchantDirectDebitMandatePageEntity[]>;
}
export { MerchantDirectDebitMandatePageEntity };
