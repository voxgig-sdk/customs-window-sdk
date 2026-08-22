import { CustomsWindowEntityBase } from '../CustomsWindowEntityBase';
import type { CustomsWindowSDK } from '../CustomsWindowSDK';
import type { PaginatedBulkUploadListList } from '../CustomsWindowTypes';
declare class PaginatedBulkUploadListListEntity extends CustomsWindowEntityBase<PaginatedBulkUploadListList> {
    constructor(client: CustomsWindowSDK, entopts: any);
    make(this: PaginatedBulkUploadListListEntity): PaginatedBulkUploadListListEntity;
}
export { PaginatedBulkUploadListListEntity };
