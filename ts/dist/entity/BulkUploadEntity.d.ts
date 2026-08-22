import { CustomsWindowEntityBase } from '../CustomsWindowEntityBase';
import type { CustomsWindowSDK } from '../CustomsWindowSDK';
import type { Control } from '../types';
import type { BulkUpload, BulkUploadLoadMatch, BulkUploadListMatch, BulkUploadCreateData, BulkUploadUpdateData, BulkUploadRemoveMatch } from '../CustomsWindowTypes';
declare class BulkUploadEntity extends CustomsWindowEntityBase<BulkUpload> {
    constructor(client: CustomsWindowSDK, entopts: any);
    make(this: BulkUploadEntity): BulkUploadEntity;
    load(this: any, reqmatch?: BulkUploadLoadMatch, ctrl?: Control): Promise<BulkUploadEntity>;
    list(this: any, reqmatch?: BulkUploadListMatch, ctrl?: Control): Promise<BulkUploadEntity[]>;
    create(this: any, reqdata?: BulkUploadCreateData, ctrl?: Control): Promise<BulkUploadEntity>;
    update(this: any, reqdata?: BulkUploadUpdateData, ctrl?: Control): Promise<BulkUploadEntity>;
    remove(this: any, reqmatch?: BulkUploadRemoveMatch, ctrl?: Control): Promise<BulkUploadEntity>;
}
export { BulkUploadEntity };
