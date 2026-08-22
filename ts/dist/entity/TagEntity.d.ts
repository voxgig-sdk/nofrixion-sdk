import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Tag, TagListMatch, TagCreateData } from '../NofrixionTypes';
declare class TagEntity extends NofrixionEntityBase<Tag> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: TagEntity): TagEntity;
    list(this: any, reqmatch?: TagListMatch, ctrl?: Control): Promise<TagEntity[]>;
    create(this: any, reqdata?: TagCreateData, ctrl?: Control): Promise<TagEntity>;
}
export { TagEntity };
