import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Mandate, MandateLoadMatch, MandateCreateData } from '../NofrixionTypes';
declare class MandateEntity extends NofrixionEntityBase<Mandate> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: MandateEntity): MandateEntity;
    load(this: any, reqmatch?: MandateLoadMatch, ctrl?: Control): Promise<MandateEntity>;
    create(this: any, reqdata?: MandateCreateData, ctrl?: Control): Promise<MandateEntity>;
}
export { MandateEntity };
