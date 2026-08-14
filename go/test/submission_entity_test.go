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

func TestSubmissionEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Submission(nil)
		if ent == nil {
			t.Fatal("expected non-nil SubmissionEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"submission": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.Submission(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.SharedConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.Submission(nil).Stream("list", nil, nil) {
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
		setup := submissionBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "submission." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		submissionRef01Ent := client.Submission(nil)
		submissionRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "submission"}, setup.data), "submission_ref01"))

		submissionRef01DataResult, err := submissionRef01Ent.Create(submissionRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		submissionRef01Data = core.ToMapAny(entityData(submissionRef01DataResult))
		if submissionRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if submissionRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		submissionRef01Match := map[string]any{}

		submissionRef01ListResult, err := submissionRef01Ent.List(submissionRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		submissionRef01List, submissionRef01ListOk := submissionRef01ListResult.([]any)
		if !submissionRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", submissionRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(submissionRef01List), map[string]any{"id": submissionRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		submissionRef01DataUp0Up := map[string]any{
			"id": submissionRef01Data["id"],
		}

		submissionRef01MarkdefUp0Name := "amendment_reason"
		submissionRef01MarkdefUp0Value := fmt.Sprintf("Mark01-submission_ref01_%d", setup.now)
		submissionRef01DataUp0Up[submissionRef01MarkdefUp0Name] = submissionRef01MarkdefUp0Value

		submissionRef01ResdataUp0Result, err := submissionRef01Ent.Update(submissionRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		submissionRef01ResdataUp0 := core.ToMapAny(entityData(submissionRef01ResdataUp0Result))
		if submissionRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if submissionRef01ResdataUp0["id"] != submissionRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if submissionRef01ResdataUp0[submissionRef01MarkdefUp0Name] != submissionRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", submissionRef01MarkdefUp0Name, submissionRef01ResdataUp0[submissionRef01MarkdefUp0Name])
		}

		// LOAD
		submissionRef01MatchDt0 := map[string]any{
			"id": submissionRef01Data["id"],
		}
		submissionRef01DataDt0Loaded, err := submissionRef01Ent.Load(submissionRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		submissionRef01DataDt0LoadResult := core.ToMapAny(entityData(submissionRef01DataDt0Loaded))
		if submissionRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if submissionRef01DataDt0LoadResult["id"] != submissionRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		submissionRef01MatchRm0 := map[string]any{
			"id": submissionRef01Data["id"],
		}
		_, err = submissionRef01Ent.Remove(submissionRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		submissionRef01MatchRt0 := map[string]any{}

		submissionRef01ListRt0Result, err := submissionRef01Ent.List(submissionRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		submissionRef01ListRt0, submissionRef01ListRt0Ok := submissionRef01ListRt0Result.([]any)
		if !submissionRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", submissionRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(submissionRef01ListRt0), map[string]any{"id": submissionRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func submissionBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "submission", "SubmissionTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read submission test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse submission test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"submission01", "submission02", "submission03"},
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
	entidEnvRaw := os.Getenv("CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID": idmap,
		"CUSTOMS_WINDOW_TEST_LIVE":      "FALSE",
		"CUSTOMS_WINDOW_TEST_EXPLAIN":   "FALSE",
		"CUSTOMS_WINDOW_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID"])
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
