import { CustomsWindowEntityBase } from '../CustomsWindowEntityBase';
import type { CustomsWindowSDK } from '../CustomsWindowSDK';
import type { PaginatedSubmissionListList } from '../CustomsWindowTypes';
declare class PaginatedSubmissionListListEntity extends CustomsWindowEntityBase<PaginatedSubmissionListList> {
    constructor(client: CustomsWindowSDK, entopts: any);
    make(this: PaginatedSubmissionListListEntity): PaginatedSubmissionListListEntity;
}
export { PaginatedSubmissionListListEntity };
