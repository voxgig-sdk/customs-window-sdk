import { CustomsWindowEntityBase } from '../CustomsWindowEntityBase';
import type { CustomsWindowSDK } from '../CustomsWindowSDK';
import type { Control } from '../types';
import type { Submission, SubmissionLoadMatch, SubmissionListMatch, SubmissionCreateData, SubmissionUpdateData, SubmissionRemoveMatch } from '../CustomsWindowTypes';
declare class SubmissionEntity extends CustomsWindowEntityBase<Submission> {
    constructor(client: CustomsWindowSDK, entopts: any);
    make(this: SubmissionEntity): SubmissionEntity;
    load(this: any, reqmatch?: SubmissionLoadMatch, ctrl?: Control): Promise<SubmissionEntity>;
    list(this: any, reqmatch?: SubmissionListMatch, ctrl?: Control): Promise<SubmissionEntity[]>;
    create(this: any, reqdata?: SubmissionCreateData, ctrl?: Control): Promise<SubmissionEntity>;
    update(this: any, reqdata?: SubmissionUpdateData, ctrl?: Control): Promise<SubmissionEntity>;
    remove(this: any, reqmatch?: SubmissionRemoveMatch, ctrl?: Control): Promise<SubmissionEntity>;
}
export { SubmissionEntity };
