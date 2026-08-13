package sdktest

import (
	"encoding/json"
	"fmt"
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

func TestBulkUploadEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.BulkUpload(nil)
		if ent == nil {
			t.Fatal("expected non-nil BulkUploadEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"bulk_upload": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.BulkUpload(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.MakeConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.BulkUpload(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := bulk_uploadBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "bulk_upload." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		bulkUploadRef01Ent := client.BulkUpload(nil)
		bulkUploadRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "bulk_upload"}, setup.data), "bulk_upload_ref01"))

		bulkUploadRef01DataResult, err := bulkUploadRef01Ent.Create(bulkUploadRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		bulkUploadRef01Data = core.ToMapAny(entityData(bulkUploadRef01DataResult))
		if bulkUploadRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if bulkUploadRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		bulkUploadRef01Match := map[string]any{}

		bulkUploadRef01ListResult, err := bulkUploadRef01Ent.List(bulkUploadRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		bulkUploadRef01List, bulkUploadRef01ListOk := bulkUploadRef01ListResult.([]any)
		if !bulkUploadRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", bulkUploadRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(bulkUploadRef01List), map[string]any{"id": bulkUploadRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		bulkUploadRef01DataUp0Up := map[string]any{
			"id": bulkUploadRef01Data["id"],
		}

		bulkUploadRef01MarkdefUp0Name := "active_transport_nationality"
		bulkUploadRef01MarkdefUp0Value := fmt.Sprintf("Mark01-bulk_upload_ref01_%d", setup.now)
		bulkUploadRef01DataUp0Up[bulkUploadRef01MarkdefUp0Name] = bulkUploadRef01MarkdefUp0Value

		bulkUploadRef01ResdataUp0Result, err := bulkUploadRef01Ent.Update(bulkUploadRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		bulkUploadRef01ResdataUp0 := core.ToMapAny(entityData(bulkUploadRef01ResdataUp0Result))
		if bulkUploadRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if bulkUploadRef01ResdataUp0["id"] != bulkUploadRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if bulkUploadRef01ResdataUp0[bulkUploadRef01MarkdefUp0Name] != bulkUploadRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", bulkUploadRef01MarkdefUp0Name, bulkUploadRef01ResdataUp0[bulkUploadRef01MarkdefUp0Name])
		}

		// LOAD
		bulkUploadRef01MatchDt0 := map[string]any{
			"id": bulkUploadRef01Data["id"],
		}
		bulkUploadRef01DataDt0Loaded, err := bulkUploadRef01Ent.Load(bulkUploadRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		bulkUploadRef01DataDt0LoadResult := core.ToMapAny(entityData(bulkUploadRef01DataDt0Loaded))
		if bulkUploadRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if bulkUploadRef01DataDt0LoadResult["id"] != bulkUploadRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		bulkUploadRef01MatchRm0 := map[string]any{
			"id": bulkUploadRef01Data["id"],
		}
		_, err = bulkUploadRef01Ent.Remove(bulkUploadRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		bulkUploadRef01MatchRt0 := map[string]any{}

		bulkUploadRef01ListRt0Result, err := bulkUploadRef01Ent.List(bulkUploadRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		bulkUploadRef01ListRt0, bulkUploadRef01ListRt0Ok := bulkUploadRef01ListRt0Result.([]any)
		if !bulkUploadRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", bulkUploadRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(bulkUploadRef01ListRt0), map[string]any{"id": bulkUploadRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func bulk_uploadBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "bulk_upload", "BulkUploadTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read bulk_upload test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse bulk_upload test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"bulk_upload01", "bulk_upload02", "bulk_upload03"},
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
	entidEnvRaw := os.Getenv("CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID": idmap,
		"CUSTOMS_WINDOW_TEST_LIVE":      "FALSE",
		"CUSTOMS_WINDOW_TEST_EXPLAIN":   "FALSE",
		"CUSTOMS_WINDOW_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["CUSTOMS_WINDOW_TEST_BULK_UPLOAD_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["CUSTOMS_WINDOW_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["CUSTOMS_WINDOW_APIKEY"],
			},
			extra,
		})
		client = sdk.NewCustomsWindowSDK(core.ToMapAny(mergedOpts))
	}

	live := env["CUSTOMS_WINDOW_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["CUSTOMS_WINDOW_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
