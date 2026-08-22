import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { RuleEvent, RuleEventListMatch } from '../NofrixionTypes';
declare class RuleEventEntity extends NofrixionEntityBase<RuleEvent> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: RuleEventEntity): RuleEventEntity;
    list(this: any, reqmatch?: RuleEventListMatch, ctrl?: Control): Promise<RuleEventEntity[]>;
}
export { RuleEventEntity };
