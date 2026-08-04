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

func TestFileEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.File(nil)
		if ent == nil {
			t.Fatal("expected non-nil FileEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := fileBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "file." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CUSTOMSWINDOW_TEST_FILE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		fileRef01Ent := client.File(nil)
		fileRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "file"}, setup.data), "file_ref01"))

		fileRef01DataResult, err := fileRef01Ent.Create(fileRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		fileRef01Data = core.ToMapAny(fileRef01DataResult)
		if fileRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if fileRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

	})
}

func fileBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "file", "FileTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read file test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse file test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"file01", "file02", "file03"},
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
	entidEnvRaw := os.Getenv("CUSTOMSWINDOW_TEST_FILE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CUSTOMSWINDOW_TEST_FILE_ENTID": idmap,
		"CUSTOMSWINDOW_TEST_LIVE":      "FALSE",
		"CUSTOMSWINDOW_TEST_EXPLAIN":   "FALSE",
		"CUSTOMSWINDOW_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["CUSTOMSWINDOW_TEST_FILE_ENTID"])
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
