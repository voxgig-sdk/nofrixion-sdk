import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Batch, BatchLoadMatch, BatchCreateData } from '../NofrixionTypes';
declare class BatchEntity extends NofrixionEntityBase<Batch> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: BatchEntity): BatchEntity;
    load(this: any, reqmatch?: BatchLoadMatch, ctrl?: Control): Promise<BatchEntity>;
    create(this: any, reqdata?: BatchCreateData, ctrl?: Control): Promise<BatchEntity>;
}
export { BatchEntity };
