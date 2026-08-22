import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { User, UserListMatch, UserUpdateData } from '../NofrixionTypes';
declare class UserEntity extends NofrixionEntityBase<User> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: UserEntity): UserEntity;
    list(this: any, reqmatch?: UserListMatch, ctrl?: Control): Promise<UserEntity[]>;
    update(this: any, reqdata?: UserUpdateData, ctrl?: Control): Promise<UserEntity>;
}
export { UserEntity };
