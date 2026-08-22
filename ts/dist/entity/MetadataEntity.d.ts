import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Metadata, MetadataLoadMatch } from '../NofrixionTypes';
declare class MetadataEntity extends NofrixionEntityBase<Metadata> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: MetadataEntity): MetadataEntity;
    load(this: any, reqmatch?: MetadataLoadMatch, ctrl?: Control): Promise<MetadataEntity>;
}
export { MetadataEntity };
