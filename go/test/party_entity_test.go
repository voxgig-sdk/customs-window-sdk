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

func TestPartyEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Party(nil)
		if ent == nil {
			t.Fatal("expected non-nil PartyEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"party": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.Party(nil).Stream("list", nil, nil) {
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
			for item := range streamSdk.Party(nil).Stream("list", nil, nil) {
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
		setup := partyBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "party." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set CUSTOMSWINDOW_TEST_PARTY_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		partyRef01Ent := client.Party(nil)
		partyRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "party"}, setup.data), "party_ref01"))

		partyRef01DataResult, err := partyRef01Ent.Create(partyRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		partyRef01Data = core.ToMapAny(partyRef01DataResult)
		if partyRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if partyRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		partyRef01Match := map[string]any{}

		partyRef01ListResult, err := partyRef01Ent.List(partyRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		partyRef01List, partyRef01ListOk := partyRef01ListResult.([]any)
		if !partyRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", partyRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(partyRef01List), map[string]any{"id": partyRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		partyRef01DataUp0Up := map[string]any{
			"id": partyRef01Data["id"],
		}

		partyRef01MarkdefUp0Name := "bank_detail"
		partyRef01MarkdefUp0Value := fmt.Sprintf("Mark01-party_ref01_%d", setup.now)
		partyRef01DataUp0Up[partyRef01MarkdefUp0Name] = partyRef01MarkdefUp0Value

		partyRef01ResdataUp0Result, err := partyRef01Ent.Update(partyRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		partyRef01ResdataUp0 := core.ToMapAny(partyRef01ResdataUp0Result)
		if partyRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if partyRef01ResdataUp0["id"] != partyRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if partyRef01ResdataUp0[partyRef01MarkdefUp0Name] != partyRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", partyRef01MarkdefUp0Name, partyRef01ResdataUp0[partyRef01MarkdefUp0Name])
		}

		// LOAD
		partyRef01MatchDt0 := map[string]any{
			"id": partyRef01Data["id"],
		}
		partyRef01DataDt0Loaded, err := partyRef01Ent.Load(partyRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		partyRef01DataDt0LoadResult := core.ToMapAny(partyRef01DataDt0Loaded)
		if partyRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if partyRef01DataDt0LoadResult["id"] != partyRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		partyRef01MatchRm0 := map[string]any{
			"id": partyRef01Data["id"],
		}
		_, err = partyRef01Ent.Remove(partyRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		partyRef01MatchRt0 := map[string]any{}

		partyRef01ListRt0Result, err := partyRef01Ent.List(partyRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		partyRef01ListRt0, partyRef01ListRt0Ok := partyRef01ListRt0Result.([]any)
		if !partyRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", partyRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(partyRef01ListRt0), map[string]any{"id": partyRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func partyBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "party", "PartyTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read party test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse party test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"party01", "party02", "party03"},
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
	entidEnvRaw := os.Getenv("CUSTOMSWINDOW_TEST_PARTY_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CUSTOMSWINDOW_TEST_PARTY_ENTID": idmap,
		"CUSTOMSWINDOW_TEST_LIVE":      "FALSE",
		"CUSTOMSWINDOW_TEST_EXPLAIN":   "FALSE",
		"CUSTOMSWINDOW_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["CUSTOMSWINDOW_TEST_PARTY_ENTID"])
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
