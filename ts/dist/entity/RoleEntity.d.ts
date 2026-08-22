import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Role, RoleCreateData } from '../NofrixionTypes';
declare class RoleEntity extends NofrixionEntityBase<Role> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: RoleEntity): RoleEntity;
    create(this: any, reqdata?: RoleCreateData, ctrl?: Control): Promise<RoleEntity>;
}
export { RoleEntity };
