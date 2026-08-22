import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { NoFrixionVersion, NoFrixionVersionLoadMatch } from '../NofrixionTypes';
declare class NoFrixionVersionEntity extends NofrixionEntityBase<NoFrixionVersion> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: NoFrixionVersionEntity): NoFrixionVersionEntity;
    load(this: any, reqmatch?: NoFrixionVersionLoadMatch, ctrl?: Control): Promise<NoFrixionVersionEntity>;
}
export { NoFrixionVersionEntity };
