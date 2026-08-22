import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { CardPublicKey, CardPublicKeyLoadMatch } from '../NofrixionTypes';
declare class CardPublicKeyEntity extends NofrixionEntityBase<CardPublicKey> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: CardPublicKeyEntity): CardPublicKeyEntity;
    load(this: any, reqmatch?: CardPublicKeyLoadMatch, ctrl?: Control): Promise<CardPublicKeyEntity>;
}
export { CardPublicKeyEntity };
