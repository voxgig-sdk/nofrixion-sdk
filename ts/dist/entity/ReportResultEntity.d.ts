import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { ReportResult, ReportResultLoadMatch } from '../NofrixionTypes';
declare class ReportResultEntity extends NofrixionEntityBase<ReportResult> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: ReportResultEntity): ReportResultEntity;
    load(this: any, reqmatch?: ReportResultLoadMatch, ctrl?: Control): Promise<ReportResultEntity>;
}
export { ReportResultEntity };
