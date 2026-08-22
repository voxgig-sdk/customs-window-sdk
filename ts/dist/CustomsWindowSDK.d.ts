import { BulkUploadEntity } from './entity/BulkUploadEntity';
import { FileEntity } from './entity/FileEntity';
import { PaginatedBulkUploadListListEntity } from './entity/PaginatedBulkUploadListListEntity';
import { PaginatedPartyListListEntity } from './entity/PaginatedPartyListListEntity';
import { PaginatedSubmissionListListEntity } from './entity/PaginatedSubmissionListListEntity';
import { PartyEntity } from './entity/PartyEntity';
import { SubmissionEntity } from './entity/SubmissionEntity';
import { SubmissionDetailEntity } from './entity/SubmissionDetailEntity';
export type * from './CustomsWindowTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { CustomsWindowEntityBase } from './CustomsWindowEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class CustomsWindowSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    BulkUpload(entopts?: Record<string, any>): BulkUploadEntity;
    File(entopts?: Record<string, any>): FileEntity;
    PaginatedBulkUploadListList(entopts?: Record<string, any>): PaginatedBulkUploadListListEntity;
    PaginatedPartyListList(entopts?: Record<string, any>): PaginatedPartyListListEntity;
    PaginatedSubmissionListList(entopts?: Record<string, any>): PaginatedSubmissionListListEntity;
    Party(entopts?: Record<string, any>): PartyEntity;
    Submission(entopts?: Record<string, any>): SubmissionEntity;
    SubmissionDetail(entopts?: Record<string, any>): SubmissionDetailEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): CustomsWindowSDK;
    tester(testopts?: any, sdkopts?: any): CustomsWindowSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof CustomsWindowSDK;
export { stdutil, config, BaseFeature, CustomsWindowEntityBase, CustomsWindowSDK, SDK, };
