import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Token, TokenCreateData, TokenRemoveMatch } from '../NofrixionTypes';
declare class TokenEntity extends NofrixionEntityBase<Token> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: TokenEntity): TokenEntity;
    create(this: any, reqdata?: TokenCreateData, ctrl?: Control): Promise<TokenEntity>;
    remove(this: any, reqmatch?: TokenRemoveMatch, ctrl?: Control): Promise<TokenEntity>;
}
export { TokenEntity };
