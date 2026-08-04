package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/customs-window-sdk/go"
	"github.com/voxgig-sdk/customs-window-sdk/go/core"

	vs "github.com/voxgig-sdk/customs-window-sdk/go/utility/struct"
)

func TestPaginatedBulkUploadListListEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.PaginatedBulkUploadListList(nil)
		if ent == nil {
			t.Fatal("expected non-nil PaginatedBulkUploadListListEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := paginated_bulk_upload_list_listBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "paginated_bulk_upload_list_list." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set CUSTOMSWINDOW_TEST_PAGINATED_BULK_UPLOAD_LIST_LIST_ENTID JSON to run live")
			return
		}
		// Bootstrap entity data from existing test data (no create step in flow).
		paginatedBulkUploadListListRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.paginated_bulk_upload_list_list", setup.data)))
		var paginatedBulkUploadListListRef01Data map[string]any
		if len(paginatedBulkUploadListListRef01DataRaw) > 0 {
			paginatedBulkUploadListListRef01Data = core.ToMapAny(paginatedBulkUploadListListRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = paginatedBulkUploadListListRef01Data

	})
}

func paginated_bulk_upload_list_listBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "paginated_bulk_upload_list_list", "PaginatedBulkUploadListListTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read paginated_bulk_upload_list_list test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse paginated_bulk_upload_list_list test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"paginated_bulk_upload_list_list01", "paginated_bulk_upload_list_list02", "paginated_bulk_upload_list_list03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("CUSTOMSWINDOW_TEST_PAGINATED_BULK_UPLOAD_LIST_LIST_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CUSTOMSWINDOW_TEST_PAGINATED_BULK_UPLOAD_LIST_LIST_ENTID": idmap,
		"CUSTOMSWINDOW_TEST_LIVE":      "FALSE",
		"CUSTOMSWINDOW_TEST_EXPLAIN":   "FALSE",
		"CUSTOMSWINDOW_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["CUSTOMSWINDOW_TEST_PAGINATED_BULK_UPLOAD_LIST_LIST_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["CUSTOMSWINDOW_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["CUSTOMSWINDOW_APIKEY"],
			},
			extra,
		})
		client = sdk.NewCustomsWindowSDK(core.ToMapAny(mergedOpts))
	}

	live := env["CUSTOMSWINDOW_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["CUSTOMSWINDOW_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
