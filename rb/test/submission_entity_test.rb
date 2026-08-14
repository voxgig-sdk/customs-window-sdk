# Submission entity test

require "minitest/autorun"
require "json"
require_relative "../CustomsWindow_sdk"
require_relative "runner"

class SubmissionEntityTest < Minitest::Test
  def test_create_instance
    testsdk = CustomsWindowSDK.test(nil, nil)
    ent = testsdk.Submission(nil)
    assert !ent.nil?
  end

  # Feature #4: the entity stream(action, ...) method runs the op pipeline and
  # returns an Enumerator over result items. With the streaming feature active
  # it yields the feature's incremental output; otherwise it falls back to the
  # materialised list so stream always yields.
  def test_stream
    seed = {
      "entity" => {
        "submission" => {
          "s1" => { "id" => "s1" },
          "s2" => { "id" => "s2" },
          "s3" => { "id" => "s3" },
        },
      },
    }

    # Fallback: streaming inactive -> yields the materialised list items.
    base = CustomsWindowSDK.test(seed, nil)
    seen = base.Submission(nil).stream("list", nil, nil).to_a
    assert_equal 3, seen.length

    # Inbound: streaming active -> yields each item from the feature.
    cfg = CustomsWindowConfig.shared_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("streaming")
      sdk = CustomsWindowSDK.test(seed, { "feature" => { "streaming" => { "active" => true } } })
      got = []
      sdk.Submission(nil).stream("list", nil, nil).each do |item|
        if item.is_a?(Array)
          got.concat(item)
        else
          got << item
        end
      end
      assert_equal 3, got.length
    end
  end

  def test_basic_flow
    setup = submission_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "update", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "submission." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    submission_ref01_ent = client.Submission(nil)
    submission_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.submission"), "submission_ref01"))

    submission_ref01_data_result = submission_ref01_ent.create(submission_ref01_data, nil)
    submission_ref01_data = Helpers.to_map(submission_ref01_data_result.respond_to?(:data_get) ? submission_ref01_data_result.data_get : submission_ref01_data_result)
    assert !submission_ref01_data.nil?
    assert !submission_ref01_data["id"].nil?

    # LIST
    submission_ref01_match = {}

    submission_ref01_list_result = submission_ref01_ent.list(submission_ref01_match, nil)
    assert submission_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(submission_ref01_list_result),
      { "id" => submission_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # UPDATE
    submission_ref01_data_up0_up = {
      "id" => submission_ref01_data["id"],
    }

    submission_ref01_markdef_up0_name = "amendment_reason"
    submission_ref01_markdef_up0_value = "Mark01-submission_ref01_#{setup[:now]}"
    submission_ref01_data_up0_up[submission_ref01_markdef_up0_name] = submission_ref01_markdef_up0_value

    submission_ref01_resdata_up0_result = submission_ref01_ent.update(submission_ref01_data_up0_up, nil)
    submission_ref01_resdata_up0 = Helpers.to_map(submission_ref01_resdata_up0_result.respond_to?(:data_get) ? submission_ref01_resdata_up0_result.data_get : submission_ref01_resdata_up0_result)
    assert !submission_ref01_resdata_up0.nil?
    assert_equal submission_ref01_resdata_up0["id"], submission_ref01_data_up0_up["id"]
    assert_equal submission_ref01_resdata_up0[submission_ref01_markdef_up0_name], submission_ref01_markdef_up0_value

    # LOAD
    submission_ref01_match_dt0 = {
      "id" => submission_ref01_data["id"],
    }
    submission_ref01_data_dt0_loaded = submission_ref01_ent.load(submission_ref01_match_dt0, nil)
    submission_ref01_data_dt0_load_result = Helpers.to_map(submission_ref01_data_dt0_loaded.respond_to?(:data_get) ? submission_ref01_data_dt0_loaded.data_get : submission_ref01_data_dt0_loaded)
    assert !submission_ref01_data_dt0_load_result.nil?
    assert_equal submission_ref01_data_dt0_load_result["id"], submission_ref01_data["id"]

    # REMOVE
    submission_ref01_match_rm0 = {
      "id" => submission_ref01_data["id"],
    }
    submission_ref01_ent.remove(submission_ref01_match_rm0, nil)

    # LIST
    submission_ref01_match_rt0 = {}

    submission_ref01_list_rt0_result = submission_ref01_ent.list(submission_ref01_match_rt0, nil)
    assert submission_ref01_list_rt0_result.is_a?(Array)

    not_found_item = Vs.select(
      Runner.entity_list_to_data(submission_ref01_list_rt0_result),
      { "id" => submission_ref01_data["id"] })
    assert Vs.isempty(not_found_item)

  end
end

def submission_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "submission", "SubmissionTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = CustomsWindowSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["submission01", "submission02", "submission03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID" => idmap,
    "CUSTOMS_WINDOW_TEST_LIVE" => "FALSE",
    "CUSTOMS_WINDOW_TEST_EXPLAIN" => "FALSE",
    "CUSTOMS_WINDOW_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["CUSTOMS_WINDOW_TEST_SUBMISSION_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["CUSTOMS_WINDOW_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
        "apikey" => env["CUSTOMS_WINDOW_APIKEY"],
      },
      extra || {},
    ])
    client = CustomsWindowSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["CUSTOMS_WINDOW_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["CUSTOMS_WINDOW_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
