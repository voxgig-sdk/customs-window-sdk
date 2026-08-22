import { CustomsWindowEntityBase } from '../CustomsWindowEntityBase';
import type { CustomsWindowSDK } from '../CustomsWindowSDK';
import type { PaginatedPartyListList } from '../CustomsWindowTypes';
declare class PaginatedPartyListListEntity extends CustomsWindowEntityBase<PaginatedPartyListList> {
    constructor(client: CustomsWindowSDK, entopts: any);
    make(this: PaginatedPartyListListEntity): PaginatedPartyListListEntity;
}
export { PaginatedPartyListListEntity };
