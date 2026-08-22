import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { MerchantToken, MerchantTokenLoadMatch, MerchantTokenListMatch, MerchantTokenCreateData, MerchantTokenUpdateData } from '../NofrixionTypes';
declare class MerchantTokenEntity extends NofrixionEntityBase<MerchantToken> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: MerchantTokenEntity): MerchantTokenEntity;
    load(this: any, reqmatch?: MerchantTokenLoadMatch, ctrl?: Control): Promise<MerchantTokenEntity>;
    list(this: any, reqmatch?: MerchantTokenListMatch, ctrl?: Control): Promise<MerchantTokenEntity[]>;
    create(this: any, reqdata?: MerchantTokenCreateData, ctrl?: Control): Promise<MerchantTokenEntity>;
    update(this: any, reqdata?: MerchantTokenUpdateData, ctrl?: Control): Promise<MerchantTokenEntity>;
}
export { MerchantTokenEntity };
