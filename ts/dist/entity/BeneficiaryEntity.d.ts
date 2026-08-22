import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Beneficiary, BeneficiaryLoadMatch, BeneficiaryListMatch, BeneficiaryCreateData, BeneficiaryUpdateData, BeneficiaryRemoveMatch } from '../NofrixionTypes';
declare class BeneficiaryEntity extends NofrixionEntityBase<Beneficiary> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: BeneficiaryEntity): BeneficiaryEntity;
    load(this: any, reqmatch?: BeneficiaryLoadMatch, ctrl?: Control): Promise<BeneficiaryEntity>;
    list(this: any, reqmatch?: BeneficiaryListMatch, ctrl?: Control): Promise<BeneficiaryEntity[]>;
    create(this: any, reqdata?: BeneficiaryCreateData, ctrl?: Control): Promise<BeneficiaryEntity>;
    update(this: any, reqdata?: BeneficiaryUpdateData, ctrl?: Control): Promise<BeneficiaryEntity>;
    remove(this: any, reqmatch?: BeneficiaryRemoveMatch, ctrl?: Control): Promise<BeneficiaryEntity>;
}
export { BeneficiaryEntity };
