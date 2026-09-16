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
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
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
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "company", "req": true, "short": "cuid-format identifier for this entity.", "type": "`$STRING`", "index$": 0 }, { "active": true, "format": "date-time", "name": "created_at", "readOnly": true, "req": true, "type": "`$STRING`", "index$": 1 }, { "active": true, "name": "extension", "readOnly": true, "req": true, "type": "`$STRING`", "index$": 2 }, { "active": true, "format": "uri", "name": "file", "req": true, "type": "`$STRING`", "index$": 3 }, { "active": true, "name": "id", "req": false, "short": "cuid-format identifier for this entity.", "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "name", "readOnly": true, "req": true, "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "public", "op": { "create": { "req": false, "type": "`$BOOLEAN`" } }, "req": true, "type": "`$BOOLEAN`", "index$": 6 }, { "active": true, "name": "size", "req": false, "type": "`$INTEGER`", "index$": 7 }, { "active": true, "format": "date-time", "name": "updated_at", "readOnly": true, "req": true, "type": "`$STRING`", "index$": 8 }, { "active": true, "name": "url", "readOnly": true, "req": true, "type": "`$STRING`", "index$": 9 }], "id": { "field": "id", "name": "id" }, "name": "file", "op": { "create": { "input": "data", "name": "create", "points": [{ "active": true, "args": {}, "contract": { "id": "POST /files", "json": "{\"operationId\":\"Create a File\",\"parameters\":[],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"*/*\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"readOnly\":true,\"type\":\"string\"},\"file\":{\"format\":\"uri\",\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"readOnly\":true,\"type\":\"string\"},\"public\":{\"default\":true,\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"file\",\"name\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}},\"application/json\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"readOnly\":true,\"type\":\"string\"},\"file\":{\"format\":\"uri\",\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"readOnly\":true,\"type\":\"string\"},\"public\":{\"default\":true,\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"file\",\"name\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}},\"application/x-www-form-urlencoded\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"readOnly\":true,\"type\":\"string\"},\"file\":{\"format\":\"uri\",\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"readOnly\":true,\"type\":\"string\"},\"public\":{\"default\":true,\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"file\",\"name\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}},\"multipart/form-data\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"readOnly\":true,\"type\":\"string\"},\"file\":{\"format\":\"uri\",\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"readOnly\":true,\"type\":\"string\"},\"public\":{\"default\":true,\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"file\",\"name\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"properties\":{\"company\":{\"description\":\"cuid-format identifier for this entity.\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"extension\":{\"maxLength\":255,\"type\":\"string\"},\"id\":{\"description\":\"cuid-format identifier for this entity.\",\"maxLength\":30,\"type\":\"string\"},\"name\":{\"maxLength\":255,\"type\":\"string\"},\"public\":{\"type\":\"boolean\"},\"size\":{\"maximum\":2147483647,\"minimum\":-2147483648,\"nullable\":true,\"type\":\"integer\"},\"updated_at\":{\"format\":\"date-time\",\"readOnly\":true,\"type\":\"string\"},\"url\":{\"readOnly\":true,\"type\":\"string\"}},\"required\":[\"company\",\"created_at\",\"extension\",\"name\",\"public\",\"updated_at\",\"url\"],\"title\":\"File\",\"type\":\"object\"}}},\"description\":\"\"}},\"security\":[{\"Secret Key\":[]}],\"securitySchemes\":{\"Secret Key\":{\"description\":\"Basic authentication with required secret-key\",\"in\":\"header\",\"name\":\"Basic Auth\",\"type\":\"apiKey\"}},\"securitySource\":\"operation\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "POST", "orig": "/files", "segments": [{ "lit": "files" }], "select": {}, "transform": { "req": { "file": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "create" } }, "relations": { "ancestors": [] }, "key$": "file", "name__orig": "file", "Name": "File", "name_": "file", "name-": "file", "NAME": "FILE", "index$": 1 }, { "active": true, "entity": "file", "key$": "BasicFileFlow", "kind": "basic", "name": "BasicFileFlow", "param": {}, "step": [{ "active": true, "data": {}, "input": { "ref": "file_ref01" }, "match": {}, "op": "create", "spec": [], "valid": [], "index$": 0 }] }, 'File');
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