# Party entity test

import json
import os
import time

import pytest

from customswindow_sdk.utility.voxgig_struct import voxgig_struct as vs
from customswindow_sdk import CustomsWindowSDK
from customswindow_sdk.core import helpers
from customswindow_sdk.config import shared_config
from customswindow_sdk.feature.base_feature import CustomsWindowBaseFeature

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class _FailHook(CustomsWindowBaseFeature):
    def __init__(self):
        super().__init__()
        self.name = "failhook"
        self.unexpected = 0

    def init(self, ctx, options):
        pass

    def PreSpec(self, ctx):
        raise RuntimeError("party hook failed")

    def PreUnexpected(self, ctx):
        self.unexpected += 1



# main.kit.test.live.strict is true (the default is true): a live
# request that fails, or a live test missing an input it needs,
# fails the test.
# An account with no record for a test to read skips it either way.
LIVE_STRICT = True


class TestPartyEntity:

    def test_should_create_instance(self):
        testsdk = CustomsWindowSDK.test(None, None)
        ent = testsdk.Party(None)
        assert ent is not None

    def test_should_stream(self):
        # Feature #4: the entity stream(action, ...) method runs the op
        # pipeline and yields result items. With the streaming feature active
        # it yields the feature's incremental output; otherwise it falls back
        # to the materialised list so stream always yields.
        seed = {
            "entity": {
                "party": {
                    "s1": {"id": "s1"},
                    "s2": {"id": "s2"},
                    "s3": {"id": "s3"},
                }
            }
        }

        # Fallback: streaming inactive -> yields the materialised list items.
        base = CustomsWindowSDK.test(seed, None)
        seen = list(base.Party(None).stream("list", None, None))
        assert len(seen) == 3

        # Inbound: streaming active -> yields each item from the feature.
        from customswindow_sdk.config import shared_config
        cfg = shared_config()
        if isinstance(cfg.get("feature"), dict) and "streaming" in cfg["feature"]:
            sdk = CustomsWindowSDK.test(
                seed, {"feature": {"streaming": {"active": True}}})
            got = []
            for item in sdk.Party(None).stream("list", None, None):
                if isinstance(item, list):
                    got.extend(item)
                else:
                    got.append(item)
            assert len(got) == 3

    def test_should_report_a_failed_stream(self):
        offline = {"net": {"offline": True}}
        with pytest.raises(Exception, match="offline"):
            list(CustomsWindowSDK.test(offline, None).Party(None).stream("list", None, None))

        quiet = {"ctrl": {"throw": False}}
        list(CustomsWindowSDK.test(offline, None).Party(None).stream("list", None, quiet))

        if "rbac" in (shared_config().get("feature") or {}):
            denied = CustomsWindowSDK.test(
                None, {"feature": {"rbac": {"active": True, "deny": True}}})
            with pytest.raises(Exception) as err:
                list(denied.Party(None).stream("list", None, None))
            assert "rbac_denied" == getattr(err.value, "code", None)

    def test_should_leave_the_callers_ctrl(self):
        explain = {}
        ctrl = {"explain": explain}
        list(CustomsWindowSDK.test(None, None).Party(None).stream("list", None, {"ctrl": ctrl}))
        assert ["explain"] == list(ctrl.keys())
        assert explain is ctrl["explain"] and 0 < len(explain)

    def test_should_fire_pre_unexpected(self):
        hook = _FailHook()
        client = CustomsWindowSDK({"feature": {"test": {"active": True}}, "extend": [hook]})
        with pytest.raises(Exception, match="hook failed"):
            client.Party(None).list(None, None)
        assert 0 < hook.unexpected

        fired = hook.unexpected
        assert client.Party(None).list(None, {"throw": False}) is None
        assert fired < hook.unexpected

    def test_should_refuse_an_invalid_request(self):
        if "validate" not in (shared_config().get("feature") or {}):
            pytest.skip("feature not present in this SDK: validate")
        client = CustomsWindowSDK.test(
            None, {"feature": {"validate": {"active": True}}})
        with pytest.raises(Exception) as err:
            client.Party(None).list({"cursor": 1}, None)
        assert "validate_failed" == getattr(err.value, "code", None)

    def test_should_run_basic_flow(self):
        setup = _party_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "list", "update", "load", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "party." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        client = setup["client"]

        # CREATE
        party_ref01_ent = client.Party(None)
        party_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.party"), "party_ref01"))

        party_ref01_data = helpers.to_map(runner.entity_data(party_ref01_ent.create(party_ref01_data, None)))
        assert party_ref01_data is not None
        assert party_ref01_data["id"] is not None

        # LIST
        party_ref01_match = {}

        party_ref01_list_result = party_ref01_ent.list(party_ref01_match, None)
        assert isinstance(party_ref01_list_result, list)

        found_item = vs.select(
            runner.entity_list_to_data(party_ref01_list_result),
            {"id": party_ref01_data["id"]})
        assert not vs.isempty(found_item)

        # UPDATE
        party_ref01_data_up0_up = {
            "id": party_ref01_data["id"],
        }

        party_ref01_markdef_up0_name = "bank_details"
        party_ref01_markdef_up0_value = "Mark01-party_ref01_" + str(setup["now"])
        party_ref01_data_up0_up[party_ref01_markdef_up0_name] = party_ref01_markdef_up0_value

        party_ref01_resdata_up0 = helpers.to_map(runner.entity_data(party_ref01_ent.update(party_ref01_data_up0_up, None)))
        assert party_ref01_resdata_up0 is not None
        assert party_ref01_resdata_up0["id"] == party_ref01_data_up0_up["id"]
        assert party_ref01_resdata_up0[party_ref01_markdef_up0_name] == party_ref01_markdef_up0_value

        # LOAD
        party_ref01_match_dt0 = {
            "id": party_ref01_data["id"],
        }
        party_ref01_data_dt0_loaded = party_ref01_ent.load(party_ref01_match_dt0, None)
        party_ref01_data_dt0_load_result = helpers.to_map(runner.entity_data(party_ref01_data_dt0_loaded))
        assert party_ref01_data_dt0_load_result is not None
        assert party_ref01_data_dt0_load_result["id"] == party_ref01_data["id"]

        # REMOVE
        party_ref01_match_rm0 = {
            "id": party_ref01_data["id"],
        }
        party_ref01_ent.remove(party_ref01_match_rm0, None)

        # LIST
        party_ref01_match_rt0 = {}

        party_ref01_list_rt0_result = party_ref01_ent.list(party_ref01_match_rt0, None)
        assert isinstance(party_ref01_list_rt0_result, list)

        not_found_item = vs.select(
            runner.entity_list_to_data(party_ref01_list_rt0_result),
            {"id": party_ref01_data["id"]})
        assert vs.isempty(not_found_item)



def _party_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/party/PartyTestData.json")
    with open(entity_data_file, "r", encoding="utf-8") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = CustomsWindowSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["party01", "party02", "party03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Whether *_ENTID supplied the idmap, read before env_override consumes
    # it: without it, the ids a live flow binds are the fixture's synthetic ones.
    _entid_env_raw = os.environ.get(
        "CUSTOMS_WINDOW_TEST_PARTY_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "CUSTOMS_WINDOW_TEST_PARTY_ENTID": idmap,
        "CUSTOMS_WINDOW_TEST_LIVE": "FALSE",
        "CUSTOMS_WINDOW_TEST_EXPLAIN": "FALSE",
        "CUSTOMS_WINDOW_APIKEY": "",
    })

    idmap_resolved = helpers.to_map(
        env.get("CUSTOMS_WINDOW_TEST_PARTY_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

    if env.get("CUSTOMS_WINDOW_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            # FIRST, so the generated fields below win: sdk-test-control.json's
            # test.client.options adds to the live client, it does not
            # redirect it.
            runner.live_client_options(),
            {
                "apikey": env.get("CUSTOMS_WINDOW_APIKEY"),
            },
            extra or {},
        ])
        client = CustomsWindowSDK(helpers.to_map(merged_opts))

    _live = env.get("CUSTOMS_WINDOW_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("CUSTOMS_WINDOW_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
