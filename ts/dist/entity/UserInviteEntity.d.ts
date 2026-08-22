import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { UserInvite, UserInviteLoadMatch, UserInviteListMatch, UserInviteCreateData, UserInviteUpdateData, UserInviteRemoveMatch } from '../NofrixionTypes';
declare class UserInviteEntity extends NofrixionEntityBase<UserInvite> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: UserInviteEntity): UserInviteEntity;
    load(this: any, reqmatch?: UserInviteLoadMatch, ctrl?: Control): Promise<UserInviteEntity>;
    list(this: any, reqmatch?: UserInviteListMatch, ctrl?: Control): Promise<UserInviteEntity[]>;
    create(this: any, reqdata?: UserInviteCreateData, ctrl?: Control): Promise<UserInviteEntity>;
    update(this: any, reqdata?: UserInviteUpdateData, ctrl?: Control): Promise<UserInviteEntity>;
    remove(this: any, reqmatch?: UserInviteRemoveMatch, ctrl?: Control): Promise<UserInviteEntity>;
}
export { UserInviteEntity };
