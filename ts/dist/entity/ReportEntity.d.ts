import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Report, ReportUpdateData } from '../NofrixionTypes';
declare class ReportEntity extends NofrixionEntityBase<Report> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: ReportEntity): ReportEntity;
    update(this: any, reqdata?: ReportUpdateData, ctrl?: Control): Promise<ReportEntity>;
}
export { ReportEntity };
