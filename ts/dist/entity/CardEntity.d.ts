import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Card, CardCreateData } from '../NofrixionTypes';
declare class CardEntity extends NofrixionEntityBase<Card> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: CardEntity): CardEntity;
    create(this: any, reqdata?: CardCreateData, ctrl?: Control): Promise<CardEntity>;
}
export { CardEntity };
