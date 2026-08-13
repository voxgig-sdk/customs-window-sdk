<?php
declare(strict_types=1);

// Submission entity test

require_once __DIR__ . '/../customswindow_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class SubmissionEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = CustomsWindowSDK::test(null, null);
        $ent = $testsdk->Submission(null);
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
                "submission" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = CustomsWindowSDK::test($seed, null);
        $seen = iterator_to_array($base->Submission(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = CustomsWindowConfig::make_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = CustomsWindowSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->Submission(null)->stream("list", null, null) as $item) {
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
        $setup = submission_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "update", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "submission." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $submission_ref01_ent = $client->Submission(null);
        $submission_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.submission"), "submission_ref01"));

        $submission_ref01_data_result = $submission_ref01_ent->create($submission_ref01_data, null);
        $submission_ref01_data = Helpers::to_map(is_object($submission_ref01_data_result) && method_exists($submission_ref01_data_result, 'data_get') ? $submission_ref01_data_result->data_get() : $submission_ref01_data_result);
        $this->assertNotNull($submission_ref01_data);
        $this->assertNotNull($submission_ref01_data["id"]);

        // LIST
        $submission_ref01_match = [];

        $submission_ref01_list_result = $submission_ref01_ent->list($submission_ref01_match, null);
        $this->assertIsArray($submission_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($submission_ref01_list_result),
            ["id" => $submission_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // UPDATE
        $submission_ref01_data_up0_up = [
            "id" => $submission_ref01_data["id"],
        ];

        $submission_ref01_markdef_up0_name = "amendment_reason";
        $submission_ref01_markdef_up0_value = "Mark01-submission_ref01_" . $setup["now"];
        $submission_ref01_data_up0_up[$submission_ref01_markdef_up0_name] = $submission_ref01_markdef_up0_value;

        $submission_ref01_resdata_up0_result = $submission_ref01_ent->update($submission_ref01_data_up0_up, null);
        $submission_ref01_resdata_up0 = Helpers::to_map(is_object($submission_ref01_resdata_up0_result) && method_exists($submission_ref01_resdata_up0_result, 'data_get') ? $submission_ref01_resdata_up0_result->data_get() : $submission_ref01_resdata_up0_result);
        $this->assertNotNull($submission_ref01_resdata_up0);
        $this->assertEquals($submission_ref01_resdata_up0["id"], $submission_ref01_data_up0_up["id"]);
        $this->assertEquals($submission_ref01_resdata_up0[$submission_ref01_markdef_up0_name], $submission_ref01_markdef_up0_value);

        // LOAD
        $submission_ref01_match_dt0 = [
            "id" => $submission_ref01_data["id"],
        ];
        $submission_ref01_data_dt0_loaded = $submission_ref01_ent->load($submission_ref01_match_dt0, null);
        $submission_ref01_data_dt0_load_result = Helpers::to_map(is_object($submission_ref01_data_dt0_loaded) && method_exists($submission_ref01_data_dt0_loaded, 'data_get') ? $submission_ref01_data_dt0_loaded->data_get() : $submission_ref01_data_dt0_loaded);
        $this->assertNotNull($submission_ref01_data_dt0_load_result);
        $this->assertEquals($submission_ref01_data_dt0_load_result["id"], $submission_ref01_data["id"]);

        // REMOVE
        $submission_ref01_match_rm0 = [
            "id" => $submission_ref01_data["id"],
        ];
        $submission_ref01_ent->remove($submission_ref01_match_rm0, null);

        // LIST
        $submission_ref01_match_rt0 = [];

        $submission_ref01_list_rt0_result = $submission_ref01_ent->list($submission_ref01_match_rt0, null);
        $this->assertIsArray($submission_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($submission_ref01_list_rt0_result),
            ["id" => $submission_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function submission_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/submission/SubmissionTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = CustomsWindowSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["submission01", "submission02", "submission03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID" => $idmap,
        "CUSTOMS_WINDOW_TEST_LIVE" => "FALSE",
        "CUSTOMS_WINDOW_TEST_EXPLAIN" => "FALSE",
        "CUSTOMS_WINDOW_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID"]);
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
