import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Consent, ConsentLoadMatch, ConsentListMatch, ConsentCreateData, ConsentUpdateData, ConsentRemoveMatch } from '../NofrixionTypes';
declare class ConsentEntity extends NofrixionEntityBase<Consent> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: ConsentEntity): ConsentEntity;
    load(this: any, reqmatch?: ConsentLoadMatch, ctrl?: Control): Promise<ConsentEntity>;
    list(this: any, reqmatch?: ConsentListMatch, ctrl?: Control): Promise<ConsentEntity[]>;
    create(this: any, reqdata?: ConsentCreateData, ctrl?: Control): Promise<ConsentEntity>;
    update(this: any, reqdata?: ConsentUpdateData, ctrl?: Control): Promise<ConsentEntity>;
    remove(this: any, reqmatch?: ConsentRemoveMatch, ctrl?: Control): Promise<ConsentEntity>;
}
export { ConsentEntity };
