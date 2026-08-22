import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Rule, RuleLoadMatch, RuleListMatch, RuleCreateData, RuleUpdateData, RuleRemoveMatch } from '../NofrixionTypes';
declare class RuleEntity extends NofrixionEntityBase<Rule> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: RuleEntity): RuleEntity;
    load(this: any, reqmatch?: RuleLoadMatch, ctrl?: Control): Promise<RuleEntity>;
    list(this: any, reqmatch?: RuleListMatch, ctrl?: Control): Promise<RuleEntity[]>;
    create(this: any, reqdata?: RuleCreateData, ctrl?: Control): Promise<RuleEntity>;
    update(this: any, reqdata?: RuleUpdateData, ctrl?: Control): Promise<RuleEntity>;
    remove(this: any, reqmatch?: RuleRemoveMatch, ctrl?: Control): Promise<RuleEntity>;
}
export { RuleEntity };
