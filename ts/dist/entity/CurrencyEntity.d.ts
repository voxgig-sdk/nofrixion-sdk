import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Currency, CurrencyListMatch } from '../NofrixionTypes';
declare class CurrencyEntity extends NofrixionEntityBase<Currency> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: CurrencyEntity): CurrencyEntity;
    list(this: any, reqmatch?: CurrencyListMatch, ctrl?: Control): Promise<CurrencyEntity[]>;
}
export { CurrencyEntity };
