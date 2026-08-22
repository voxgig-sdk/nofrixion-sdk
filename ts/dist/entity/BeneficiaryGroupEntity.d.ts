import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { BeneficiaryGroup, BeneficiaryGroupListMatch } from '../NofrixionTypes';
declare class BeneficiaryGroupEntity extends NofrixionEntityBase<BeneficiaryGroup> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: BeneficiaryGroupEntity): BeneficiaryGroupEntity;
    list(this: any, reqmatch?: BeneficiaryGroupListMatch, ctrl?: Control): Promise<BeneficiaryGroupEntity[]>;
}
export { BeneficiaryGroupEntity };
