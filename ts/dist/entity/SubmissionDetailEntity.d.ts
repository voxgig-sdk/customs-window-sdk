import { CustomsWindowEntityBase } from '../CustomsWindowEntityBase';
import type { CustomsWindowSDK } from '../CustomsWindowSDK';
import type { Control } from '../types';
import type { SubmissionDetail, SubmissionDetailCreateData } from '../CustomsWindowTypes';
declare class SubmissionDetailEntity extends CustomsWindowEntityBase<SubmissionDetail> {
    constructor(client: CustomsWindowSDK, entopts: any);
    make(this: SubmissionDetailEntity): SubmissionDetailEntity;
    create(this: any, reqdata?: SubmissionDetailCreateData, ctrl?: Control): Promise<SubmissionDetailEntity>;
}
export { SubmissionDetailEntity };
