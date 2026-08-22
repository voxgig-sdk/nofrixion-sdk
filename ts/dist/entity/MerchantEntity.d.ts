import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Merchant, MerchantLoadMatch, MerchantListMatch, MerchantUpdateData, MerchantRemoveMatch } from '../NofrixionTypes';
declare class MerchantEntity extends NofrixionEntityBase<Merchant> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: MerchantEntity): MerchantEntity;
    load(this: any, reqmatch?: MerchantLoadMatch, ctrl?: Control): Promise<MerchantEntity>;
    list(this: any, reqmatch?: MerchantListMatch, ctrl?: Control): Promise<MerchantEntity[]>;
    update(this: any, reqdata?: MerchantUpdateData, ctrl?: Control): Promise<MerchantEntity>;
    remove(this: any, reqmatch?: MerchantRemoveMatch, ctrl?: Control): Promise<MerchantEntity>;
}
export { MerchantEntity };
