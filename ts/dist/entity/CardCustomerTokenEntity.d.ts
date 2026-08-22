import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { CardCustomerToken, CardCustomerTokenLoadMatch, CardCustomerTokenListMatch, CardCustomerTokenRemoveMatch } from '../NofrixionTypes';
declare class CardCustomerTokenEntity extends NofrixionEntityBase<CardCustomerToken> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: CardCustomerTokenEntity): CardCustomerTokenEntity;
    load(this: any, reqmatch?: CardCustomerTokenLoadMatch, ctrl?: Control): Promise<CardCustomerTokenEntity>;
    list(this: any, reqmatch?: CardCustomerTokenListMatch, ctrl?: Control): Promise<CardCustomerTokenEntity[]>;
    remove(this: any, reqmatch?: CardCustomerTokenRemoveMatch, ctrl?: Control): Promise<CardCustomerTokenEntity>;
}
export { CardCustomerTokenEntity };
