import { NofrixionEntityBase } from '../NofrixionEntityBase';
import type { NofrixionSDK } from '../NofrixionSDK';
import type { Control } from '../types';
import type { Transaction, TransactionLoadMatch, TransactionListMatch, TransactionCreateData, TransactionRemoveMatch } from '../NofrixionTypes';
declare class TransactionEntity extends NofrixionEntityBase<Transaction> {
    constructor(client: NofrixionSDK, entopts: any);
    make(this: TransactionEntity): TransactionEntity;
    load(this: any, reqmatch?: TransactionLoadMatch, ctrl?: Control): Promise<TransactionEntity>;
    list(this: any, reqmatch?: TransactionListMatch, ctrl?: Control): Promise<TransactionEntity[]>;
    create(this: any, reqdata?: TransactionCreateData, ctrl?: Control): Promise<TransactionEntity>;
    remove(this: any, reqmatch?: TransactionRemoveMatch, ctrl?: Control): Promise<TransactionEntity>;
}
export { TransactionEntity };
