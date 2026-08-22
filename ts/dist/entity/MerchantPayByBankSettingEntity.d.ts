import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { MerchantPayByBankSetting, MerchantPayByBankSettingListMatch } from '../NofrixionTypes';
declare class MerchantPayByBankSettingEntity extends NofrixionEntityBase<MerchantPayByBankSetting> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: MerchantPayByBankSettingEntity): MerchantPayByBankSettingEntity;
    list(this: any, reqmatch?: MerchantPayByBankSettingListMatch, ctrl?: Control): Promise<MerchantPayByBankSettingEntity[]>;
}
export { MerchantPayByBankSettingEntity };
