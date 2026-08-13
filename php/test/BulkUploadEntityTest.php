<?php
declare(strict_types=1);

// BulkUpload entity test

require_once __DIR__ . '/../customswindow_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class BulkUploadEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = CustomsWindowSDK::test(null, null);
        $ent = $testsdk->BulkUpload(null);
        $this->assertNotNull($ent);
    }

    // Feature #4: the entity stream(action, ...) method runs the op pipeline
    // and yields result items. With the streaming feature active it yields the
    // feature's incremental output; otherwise it falls back to the materialised
    // list so stream always yields.
    public function test_stream(): void
    {
        $seed = [
            "entity" => [
                "bulk_upload" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = CustomsWindowSDK::test($seed, null);
        $seen = iterator_to_array($base->BulkUpload(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = CustomsWindowConfig::make_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = CustomsWindowSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->BulkUpload(null)->stream("list", null, null) as $item) {
                if (is_array($item) && array_is_list($item)) {
                    foreach ($item as $sub) {
                        $got[] = $sub;
                    }
                } else {
                    $got[] = $item;
                }
            }
            $this->assertCount(3, $got);
        }
    }

    public function test_basic_flow(): void
    {
        $setup = bulk_upload_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "update", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "bulk_upload." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $bulk_upload_ref01_ent = $client->BulkUpload(null);
        $bulk_upload_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.bulk_upload"), "bulk_upload_ref01"));

        $bulk_upload_ref01_data_result = $bulk_upload_ref01_ent->create($bulk_upload_ref01_data, null);
        $bulk_upload_ref01_data = Helpers::to_map(is_object($bulk_upload_ref01_data_result) && method_exists($bulk_upload_ref01_data_result, 'data_get') ? $bulk_upload_ref01_data_result->data_get() : $bulk_upload_ref01_data_result);
        $this->assertNotNull($bulk_upload_ref01_data);
        $this->assertNotNull($bulk_upload_ref01_data["id"]);

        // LIST
        $bulk_upload_ref01_match = [];

        $bulk_upload_ref01_list_result = $bulk_upload_ref01_ent->list($bulk_upload_ref01_match, null);
        $this->assertIsArray($bulk_upload_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($bulk_upload_ref01_list_result),
            ["id" => $bulk_upload_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // UPDATE
        $bulk_upload_ref01_data_up0_up = [
            "id" => $bulk_upload_ref01_data["id"],
        ];

        $bulk_upload_ref01_markdef_up0_name = "active_transport_nationality";
        $bulk_upload_ref01_markdef_up0_value = "Mark01-bulk_upload_ref01_" . $setup["now"];
        $bulk_upload_ref01_data_up0_up[$bulk_upload_ref01_markdef_up0_name] = $bulk_upload_ref01_markdef_up0_value;

        $bulk_upload_ref01_resdata_up0_result = $bulk_upload_ref01_ent->update($bulk_upload_ref01_data_up0_up, null);
        $bulk_upload_ref01_resdata_up0 = Helpers::to_map(is_object($bulk_upload_ref01_resdata_up0_result) && method_exists($bulk_upload_ref01_resdata_up0_result, 'data_get') ? $bulk_upload_ref01_resdata_up0_result->data_get() : $bulk_upload_ref01_resdata_up0_result);
        $this->assertNotNull($bulk_upload_ref01_resdata_up0);
        $this->assertEquals($bulk_upload_ref01_resdata_up0["id"], $bulk_upload_ref01_data_up0_up["id"]);
        $this->assertEquals($bulk_upload_ref01_resdata_up0[$bulk_upload_ref01_markdef_up0_name], $bulk_upload_ref01_markdef_up0_value);

        // LOAD
        $bulk_upload_ref01_match_dt0 = [
            "id" => $bulk_upload_ref01_data["id"],
        ];
        $bulk_upload_ref01_data_dt0_loaded = $bulk_upload_ref01_ent->load($bulk_upload_ref01_match_dt0, null);
        $bulk_upload_ref01_data_dt0_load_result = Helpers::to_map(is_object($bulk_upload_ref01_data_dt0_loaded) && method_exists($bulk_upload_ref01_data_dt0_loaded, 'data_get') ? $bulk_upload_ref01_data_dt0_loaded->data_get() : $bulk_upload_ref01_data_dt0_loaded);
        $this->assertNotNull($bulk_upload_ref01_data_dt0_load_result);
        $this->assertEquals($bulk_upload_ref01_data_dt0_load_result["id"], $bulk_upload_ref01_data["id"]);

        // REMOVE
        $bulk_upload_ref01_match_rm0 = [
            "id" => $bulk_upload_ref01_data["id"],
        ];
        $bulk_upload_ref01_ent->remove($bulk_upload_ref01_match_rm0, null);

        // LIST
        $bulk_upload_ref01_match_rt0 = [];

        $bulk_upload_ref01_list_rt0_result = $bulk_upload_ref01_ent->list($bulk_upload_ref01_match_rt0, null);
        $this->assertIsArray($bulk_upload_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($bulk_upload_ref01_list_rt0_result),
            ["id" => $bulk_upload_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function bulk_upload_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/bulk_upload/BulkUploadTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = CustomsWindowSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["bulk_upload01", "bulk_upload02", "bulk_upload03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID" => $idmap,
        "CUSTOMS_WINDOW_TEST_LIVE" => "FALSE",
        "CUSTOMS_WINDOW_TEST_EXPLAIN" => "FALSE",
        "CUSTOMS_WINDOW_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["CUSTOMS_WINDOW_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["CUSTOMS_WINDOW_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new CustomsWindowSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["CUSTOMS_WINDOW_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["CUSTOMS_WINDOW_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
