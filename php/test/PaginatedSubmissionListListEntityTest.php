<?php
declare(strict_types=1);

// PaginatedSubmissionListList entity test

require_once __DIR__ . '/../customswindow_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class PaginatedSubmissionListListEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = CustomsWindowSDK::test(null, null);
        $ent = $testsdk->PaginatedSubmissionListList(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = paginated_submission_list_list_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach ([] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "paginated_submission_list_list." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set CUSTOMSWINDOW_TEST_PAGINATED_SUBMISSION_LIST_LIST_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $paginated_submission_list_list_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.paginated_submission_list_list")));
        $paginated_submission_list_list_ref01_data = null;
        if (count($paginated_submission_list_list_ref01_data_raw) > 0) {
            $paginated_submission_list_list_ref01_data = Helpers::to_map($paginated_submission_list_list_ref01_data_raw[0][1]);
        }

    }
}

function paginated_submission_list_list_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/paginated_submission_list_list/PaginatedSubmissionListListTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = CustomsWindowSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["paginated_submission_list_list01", "paginated_submission_list_list02", "paginated_submission_list_list03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("CUSTOMSWINDOW_TEST_PAGINATED_SUBMISSION_LIST_LIST_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "CUSTOMSWINDOW_TEST_PAGINATED_SUBMISSION_LIST_LIST_ENTID" => $idmap,
        "CUSTOMSWINDOW_TEST_LIVE" => "FALSE",
        "CUSTOMSWINDOW_TEST_EXPLAIN" => "FALSE",
        "CUSTOMSWINDOW_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["CUSTOMSWINDOW_TEST_PAGINATED_SUBMISSION_LIST_LIST_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["CUSTOMSWINDOW_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["CUSTOMSWINDOW_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new CustomsWindowSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["CUSTOMSWINDOW_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["CUSTOMSWINDOW_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
