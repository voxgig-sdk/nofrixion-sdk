import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { FxRate, FxRateLoadMatch, FxRateListMatch } from '../NofrixionTypes';
declare class FxRateEntity extends NofrixionEntityBase<FxRate> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: FxRateEntity): FxRateEntity;
    load(this: any, reqmatch?: FxRateLoadMatch, ctrl?: Control): Promise<FxRateEntity>;
    list(this: any, reqmatch?: FxRateListMatch, ctrl?: Control): Promise<FxRateEntity[]>;
}
export { FxRateEntity };
