import { CustomsWindowEntityBase } from '../CustomsWindowEntityBase';
import type { CustomsWindowSDK } from '../CustomsWindowSDK';
import type { Control } from '../types';
import type { File, FileCreateData } from '../CustomsWindowTypes';
declare class FileEntity extends CustomsWindowEntityBase<File> {
    constructor(client: CustomsWindowSDK, entopts: any);
    make(this: FileEntity): FileEntity;
    create(this: any, reqdata?: FileCreateData, ctrl?: Control): Promise<FileEntity>;
}
export { FileEntity };
