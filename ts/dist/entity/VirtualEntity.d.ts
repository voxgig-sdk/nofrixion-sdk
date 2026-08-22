import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Virtual, VirtualCreateData, VirtualUpdateData } from '../NofrixionTypes';
declare class VirtualEntity extends NofrixionEntityBase<Virtual> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: VirtualEntity): VirtualEntity;
    create(this: any, reqdata?: VirtualCreateData, ctrl?: Control): Promise<VirtualEntity>;
    update(this: any, reqdata?: VirtualUpdateData, ctrl?: Control): Promise<VirtualEntity>;
}
export { VirtualEntity };
