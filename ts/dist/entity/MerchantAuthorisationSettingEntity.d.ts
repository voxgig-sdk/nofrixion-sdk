import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { MerchantAuthorisationSetting, MerchantAuthorisationSettingListMatch } from '../NofrixionTypes';
declare class MerchantAuthorisationSettingEntity extends NofrixionEntityBase<MerchantAuthorisationSetting> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: MerchantAuthorisationSettingEntity): MerchantAuthorisationSettingEntity;
    list(this: any, reqmatch?: MerchantAuthorisationSettingListMatch, ctrl?: Control): Promise<MerchantAuthorisationSettingEntity[]>;
}
export { MerchantAuthorisationSettingEntity };
