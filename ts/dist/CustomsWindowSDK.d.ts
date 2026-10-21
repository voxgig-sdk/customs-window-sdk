import { BulkUploadEntity } from './entity/BulkUploadEntity';
import { FileEntity } from './entity/FileEntity';
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
import * as sekreto from './feature/secrets/sekreto';
declare const stdutil: Utility;
type DirectResult = {
    ok: false;
    err: any;
    status?: undefined;
    headers?: undefined;
    data?: undefined;
} | {
    ok: boolean;
    status: number;
    headers: any;
    data: any;
    err?: any;
};
declare class CustomsWindowSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    _secrets?: any;
    constructor(options?: any);
    options(): any;
    utility(): any;
    secrets(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<DirectResult>;
    _rawRequest(fetchargs?: any): Promise<DirectResult>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    BulkUpload(entopts?: Record<string, any>): BulkUploadEntity;
    File(entopts?: Record<string, any>): FileEntity;
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
export { stdutil, config, sekreto, BaseFeature, CustomsWindowEntityBase, CustomsWindowSDK, SDK, };
export type { DirectResult };
