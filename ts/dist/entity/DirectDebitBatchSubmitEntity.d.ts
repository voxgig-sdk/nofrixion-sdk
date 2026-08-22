import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { DirectDebitBatchSubmit, DirectDebitBatchSubmitCreateData } from '../NofrixionTypes';
declare class DirectDebitBatchSubmitEntity extends NofrixionEntityBase<DirectDebitBatchSubmit> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: DirectDebitBatchSubmitEntity): DirectDebitBatchSubmitEntity;
    create(this: any, reqdata?: DirectDebitBatchSubmitCreateData, ctrl?: Control): Promise<DirectDebitBatchSubmitEntity>;
}
export { DirectDebitBatchSubmitEntity };
