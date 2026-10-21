# SubmissionDetail entity test

require "minitest/autorun"
require "json"
require_relative "../CustomsWindow_sdk"
require_relative "runner"

class SubmissionDetailEntityTest < Minitest::Test
  # main.kit.test.live.strict is true (the default is true): a live
  # request that fails, or a live test missing an input it needs,
  # fails the test.
  # An account with no record for a test to read skips it either way.
  LIVE_STRICT = true

  def test_create_instance
    testsdk = CustomsWindowSDK.test(nil, nil)
    ent = testsdk.SubmissionDetail(nil)
    assert !ent.nil?
  end

  def test_validate
    cfg = CustomsWindowConfig.shared_config
    unless cfg["feature"].is_a?(Hash) && cfg["feature"].key?("validate")
      skip("feature not present in this SDK: validate")
    end
    client = CustomsWindowSDK.test(nil, { "feature" => { "validate" => { "active" => true } } })
    err = assert_raises(StandardError) do
      client.SubmissionDetail(nil).create({ "additional_information" => "x", "company" => 1, "company_member" => "x", "submission" => "x" }, nil)
    end
    assert_equal "validate_failed", err.code
  end

  def test_basic_flow
    setup = submission_detail_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "submission_detail." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    client = setup[:client]

    # CREATE
    submission_detail_ref01_ent = client.SubmissionDetail(nil)
    submission_detail_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.submission_detail"), "submission_detail_ref01"))

    submission_detail_ref01_data_result = submission_detail_ref01_ent.create(submission_detail_ref01_data, nil)
    submission_detail_ref01_data = Helpers.to_map(submission_detail_ref01_data_result.respond_to?(:data_get) ? submission_detail_ref01_data_result.data_get : submission_detail_ref01_data_result)
    assert !submission_detail_ref01_data.nil?
    assert !submission_detail_ref01_data["id"].nil?

  end
end

def submission_detail_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "submission_detail", "SubmissionDetailTestData.json")
  entity_data_source = File.read(entity_data_file, encoding: "UTF-8")
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = CustomsWindowSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["submission_detail01", "submission_detail02", "submission_detail03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Whether *_ENTID supplied the idmap, read before env_override consumes
  # it: without it, the ids a live flow binds are the fixture's synthetic ones.
  entid_env_raw = ENV["CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID" => idmap,
    "CUSTOMS_WINDOW_TEST_LIVE" => "FALSE",
    "CUSTOMS_WINDOW_TEST_EXPLAIN" => "FALSE",
    "CUSTOMS_WINDOW_APIKEY" => "",
  })

  idmap_resolved = Helpers.to_map(
    env["CUSTOMS_WINDOW_TEST_SUBMISSION_DETAIL_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["CUSTOMS_WINDOW_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      # FIRST, so the generated fields below win: sdk-test-control.json's
      # test.client.options adds to the live client, it does not redirect it.
      Runner.live_client_options,
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
