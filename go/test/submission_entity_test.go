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

// main.kit.test.live.strict is true (the default is true): a live
// request that fails, or a live test missing an input it needs,
// fails the test.
// An account with no record for a test to read skips it either way.
const submissionEntityLiveStrict = true


type submissionFailHook struct {
	sdk.BaseFeature
	unexpected int
}

func (f *submissionFailHook) PreSpec(ctx *sdk.Context) {
	panic("submission hook failed")
}

func (f *submissionFailHook) PreUnexpected(ctx *sdk.Context) {
	f.unexpected++
}

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
		for si := range base.Submission(nil).Stream("list", nil, nil) {
			if si.Err != nil {
				t.Fatalf("stream failed: %v", si.Err)
			}
			seen = append(seen, si.Item)
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
			for si := range streamSdk.Submission(nil).Stream("list", nil, nil) {
				if si.Err != nil {
					t.Fatalf("stream failed: %v", si.Err)
				}
				if sub, ok := si.Item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, si.Item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("stream-error", func(t *testing.T) {
		offline := map[string]any{"net": map[string]any{"offline": true}}
		var streamerr error
		for si := range sdk.TestSDK(offline, nil).Submission(nil).Stream("list", nil, nil) {
			if si.Err != nil {
				streamerr = si.Err
			}
		}
		if nil == streamerr || !strings.Contains(streamerr.Error(), "offline") {
			t.Fatalf("expected the transport failure as a stream value, got %v", streamerr)
		}

		quiet := map[string]any{"ctrl": map[string]any{"throw": false}}
		for si := range sdk.TestSDK(offline, nil).Submission(nil).Stream("list", nil, quiet) {
			if si.Err != nil {
				t.Fatalf("throw false: expected no error value, got %v", si.Err)
			}
		}

		if fhHasFeature("rbac") {
			denied := sdk.TestSDK(nil, map[string]any{
				"feature": map[string]any{"rbac": map[string]any{"active": true, "deny": true}},
			})
			var denyerr error
			for si := range denied.Submission(nil).Stream("list", nil, nil) {
				if si.Err != nil {
					denyerr = si.Err
				}
			}
			if sdkerr, ok := denyerr.(*core.CustomsWindowError); !ok || "rbac_denied" != sdkerr.Code {
				t.Fatalf("expected the rbac denial as a stream value, got %v", denyerr)
			}
		}
	})

	t.Run("stream-ctrl", func(t *testing.T) {
		explain := map[string]any{}
		ctrl := map[string]any{"explain": explain}
		for range sdk.TestSDK(nil, nil).Submission(nil).Stream("list", nil, map[string]any{"ctrl": ctrl}) {
		}
		if _, has := ctrl["stream"]; has || 1 != len(ctrl) {
			t.Fatalf("the stream changed the caller's ctrl")
		}
		if 0 == len(explain) {
			t.Fatalf("the caller's explain record was not filled")
		}
	})

	t.Run("unexpected", func(t *testing.T) {
		hook := &submissionFailHook{
			BaseFeature: sdk.BaseFeature{Version: "0.0.1", Name: "failhook", Active: true}}
		client := sdk.TestSDK(nil, map[string]any{"extend": []any{hook}})

		_, err := client.Submission(nil).List(nil, nil)
		if nil == err || !strings.Contains(err.Error(), "hook failed") {
			t.Fatalf("expected the hook's failure, got %v", err)
		}
		if 0 == hook.unexpected {
			t.Fatalf("PreUnexpected did not fire")
		}

		fired := hook.unexpected
		if _, err := client.Submission(nil).List(nil, map[string]any{"throw": false}); nil != err {
			t.Fatalf("throw false: expected no error, got %v", err)
		}
		if fired == hook.unexpected {
			t.Fatalf("throw false: PreUnexpected did not fire")
		}
	})

	t.Run("validate", func(t *testing.T) {
		if !fhHasFeature("validate") {
			t.Skip("feature not present in this SDK: validate")
		}
		client := sdk.TestSDK(nil, map[string]any{
			"feature": map[string]any{"validate": map[string]any{"active": true}},
		})
		_, err := client.Submission(nil).List(map[string]any{"cursor": 1}, nil)
		if sdkerr, ok := err.(*core.CustomsWindowError); !ok || "validate_failed" != sdkerr.Code {
			t.Fatalf("expected validate_failed, got %v", err)
		}
	})

	t.Run("basic", func(tt *testing.T) {
		var t testing.TB = tt
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
		client := setup.client
		_ = client

		// CREATE
		submissionRef01Ent := client.Submission(nil)
		submissionRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath(setup.data, []any{"new", "submission"}), "submission_ref01"))

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
	idmap, _ := vs.Transform(
		[]any{"submission01", "submission02", "submission03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Whether *_ENTID supplied the idmap, read before envOverride consumes it:
	// without it, the ids a live flow binds are the fixture's synthetic ones.
	entidEnvRaw := os.Getenv("CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID": idmap,
		"CUSTOMS_WINDOW_TEST_LIVE":      "FALSE",
		"CUSTOMS_WINDOW_TEST_EXPLAIN":   "FALSE",
		"CUSTOMS_WINDOW_APIKEY":         "",
	})

	idmapResolved := core.ToMapAny(env["CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["CUSTOMS_WINDOW_TEST_LIVE"] == "TRUE" {
		// An empty map, not a nil one: Merge returns nil when its last entry
		// is nil, and BasicSetup is normally called with no extras - so a
		// bare nil silently discarded the apikey and server values below.
		extraOpts := extra
		if extraOpts == nil {
			extraOpts = map[string]any{}
		}

		mergedOpts := vs.Merge([]any{
			// liveClientOptions() FIRST, so the generated fields below win:
			// sdk-test-control.json's test.client.options adds to the live
			// client, it does not redirect it.
			liveClientOptions(),
			map[string]any{
				"apikey": env["CUSTOMS_WINDOW_APIKEY"],
			},
			extraOpts,
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
