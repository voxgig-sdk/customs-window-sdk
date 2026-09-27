"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('FileEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when CUSTOMS_WINDOW_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('CUSTOMS_WINDOW_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.CustomsWindowSDK.test();
        const ent = testsdk.File();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.CUSTOMS_WINDOW_TEST_LIVE;
        for (const op of ['create']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'file.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "company": { "a": true, "h": "Company", "n": "company", "r": true, "sh": "cuid-format identifier for this entity.", "t": "`$STRING`", "key$": "company", "index$": 0 }, "created_at": { "a": true, "fo": "date-time", "h": "Created At", "n": "created_at", "r": true, "ro": true, "t": "`$STRING`", "key$": "created_at", "index$": 1 }, "extension": { "a": true, "h": "Extension", "n": "extension", "r": true, "ro": true, "t": "`$STRING`", "key$": "extension", "index$": 2 }, "file": { "a": true, "fo": "uri", "h": "File", "n": "file", "r": true, "t": "`$STRING`", "key$": "file", "index$": 3 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "cuid-format identifier for this entity.", "t": "`$STRING`", "key$": "id", "index$": 4 }, "name": { "a": true, "h": "Name", "n": "name", "r": true, "ro": true, "t": "`$STRING`", "key$": "name", "index$": 5 }, "public": { "a": true, "h": "Public", "n": "public", "op": { "create": { "req": false, "type": "`$BOOLEAN`" } }, "r": true, "t": "`$BOOLEAN`", "key$": "public", "index$": 6 }, "size": { "a": true, "h": "Size", "n": "size", "r": false, "t": "`$INTEGER`", "key$": "size", "index$": 7 }, "updated_at": { "a": true, "fo": "date-time", "h": "Updated At", "n": "updated_at", "r": true, "ro": true, "t": "`$STRING`", "key$": "updated_at", "index$": 8 }, "url": { "a": true, "h": "Url", "n": "url", "r": true, "ro": true, "t": "`$STRING`", "key$": "url", "index$": 9 } }, "id": { "field": "id", "name": "id" }, "name": "file", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /files", "source": "openapi3", "version": 2 }, "g": {}, "k": "http", "m": "POST", "o": "/files", "q": {}, "r": {}, "s": [{ "lit": "files" }], "t": { "req": { "file": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "file", "name__orig": "file", "Name": "File", "name_": "file", "name-": "file", "NAME": "FILE", "index$": 1 }, { "active": true, "entity": "file", "key$": "BasicFileFlow", "kind": "basic", "name": "BasicFileFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "file_ref01" }, "m": {}, "o": "create", "s": [], "v": [], "index$": 0 }] }, 'File', { "POST /files": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "id": { "type": "string", "description": "cuid-format identifier for this entity.", "maxLength": 30, "key$": "id" }, "file": { "type": "string", "format": "uri", "key$": "file" }, "company": { "type": "string", "description": "cuid-format identifier for this entity.", "key$": "company" }, "public": { "type": "boolean", "default": true, "key$": "public" }, "url": { "type": "string", "readOnly": true, "key$": "url" }, "name": { "type": "string", "readOnly": true, "key$": "name" }, "extension": { "type": "string", "readOnly": true, "key$": "extension" }, "size": { "type": "integer", "maximum": 2147483647, "minimum": -2147483648, "nullable": true, "key$": "size" }, "updated_at": { "type": "string", "format": "date-time", "readOnly": true, "key$": "updated_at" }, "created_at": { "type": "string", "format": "date-time", "readOnly": true, "key$": "created_at" } }, "required": ["company", "created_at", "extension", "file", "name", "updated_at", "url"], "title": "File", "x-ref": "#/components/schemas/FileCreate", "index$": 1 } }, "application/x-www-form-urlencoded": { "schema": { "type": "object", "properties": { "id": { "type": "string", "description": "cuid-format identifier for this entity.", "maxLength": 30, "key$": "id" }, "file": { "type": "string", "format": "uri", "key$": "file" }, "company": { "type": "string", "description": "cuid-format identifier for this entity.", "key$": "company" }, "public": { "type": "boolean", "default": true, "key$": "public" }, "url": { "type": "string", "readOnly": true, "key$": "url" }, "name": { "type": "string", "readOnly": true, "key$": "name" }, "extension": { "type": "string", "readOnly": true, "key$": "extension" }, "size": { "type": "integer", "maximum": 2147483647, "minimum": -2147483648, "nullable": true, "key$": "size" }, "updated_at": { "type": "string", "format": "date-time", "readOnly": true, "key$": "updated_at" }, "created_at": { "type": "string", "format": "date-time", "readOnly": true, "key$": "created_at" } }, "required": ["company", "created_at", "extension", "file", "name", "updated_at", "url"], "title": "File", "x-ref": "#/components/schemas/FileCreate" } }, "multipart/form-data": { "schema": { "type": "object", "properties": { "id": { "type": "string", "description": "cuid-format identifier for this entity.", "maxLength": 30, "key$": "id" }, "file": { "type": "string", "format": "uri", "key$": "file" }, "company": { "type": "string", "description": "cuid-format identifier for this entity.", "key$": "company" }, "public": { "type": "boolean", "default": true, "key$": "public" }, "url": { "type": "string", "readOnly": true, "key$": "url" }, "name": { "type": "string", "readOnly": true, "key$": "name" }, "extension": { "type": "string", "readOnly": true, "key$": "extension" }, "size": { "type": "integer", "maximum": 2147483647, "minimum": -2147483648, "nullable": true, "key$": "size" }, "updated_at": { "type": "string", "format": "date-time", "readOnly": true, "key$": "updated_at" }, "created_at": { "type": "string", "format": "date-time", "readOnly": true, "key$": "created_at" } }, "required": ["company", "created_at", "extension", "file", "name", "updated_at", "url"], "title": "File", "x-ref": "#/components/schemas/FileCreate" } }, "*/*": { "schema": { "type": "object", "properties": { "id": { "type": "string", "description": "cuid-format identifier for this entity.", "maxLength": 30, "key$": "id" }, "file": { "type": "string", "format": "uri", "key$": "file" }, "company": { "type": "string", "description": "cuid-format identifier for this entity.", "key$": "company" }, "public": { "type": "boolean", "default": true, "key$": "public" }, "url": { "type": "string", "readOnly": true, "key$": "url" }, "name": { "type": "string", "readOnly": true, "key$": "name" }, "extension": { "type": "string", "readOnly": true, "key$": "extension" }, "size": { "type": "integer", "maximum": 2147483647, "minimum": -2147483648, "nullable": true, "key$": "size" }, "updated_at": { "type": "string", "format": "date-time", "readOnly": true, "key$": "updated_at" }, "created_at": { "type": "string", "format": "date-time", "readOnly": true, "key$": "created_at" } }, "required": ["company", "created_at", "extension", "file", "name", "updated_at", "url"], "title": "File", "x-ref": "#/components/schemas/FileCreate" } } }, "required": true }, "parameters": [] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const file_ref01_ent = client.File();
        let file_ref01_data = setup.data.new.file['file_ref01'];
        file_ref01_data = (await file_ref01_ent.create(file_ref01_data)).data();
        (0, node_assert_1.default)(null != file_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/file/FileTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.CustomsWindowSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['file01', 'file02', 'file03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'CUSTOMS_WINDOW_TEST_FILE_ENTID': idmap,
        'CUSTOMS_WINDOW_TEST_LIVE': 'FALSE',
        'CUSTOMS_WINDOW_TEST_EXPLAIN': 'FALSE',
        'CUSTOMS_WINDOW_APIKEY': '',
    });
    idmap = env['CUSTOMS_WINDOW_TEST_FILE_ENTID'];
    const live = 'TRUE' === env.CUSTOMS_WINDOW_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['CUSTOMS_WINDOW_TEST_FILE_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.CustomsWindowSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.CUSTOMS_WINDOW_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.CUSTOMS_WINDOW_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=FileEntity.test.js.map