import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { OpenBanking, OpenBankingCreateData, OpenBankingRemoveMatch } from '../NofrixionTypes';
declare class OpenBankingEntity extends NofrixionEntityBase<OpenBanking> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: OpenBankingEntity): OpenBankingEntity;
    create(this: any, reqdata?: OpenBankingCreateData, ctrl?: Control): Promise<OpenBankingEntity>;
    remove(this: any, reqmatch?: OpenBankingRemoveMatch, ctrl?: Control): Promise<OpenBankingEntity>;
}
export { OpenBankingEntity };
