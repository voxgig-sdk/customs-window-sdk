# CustomsWindow SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

CustomsWindowUtility.registrar = ->(u) {
  u.clean = CustomsWindowUtilities::Clean
  u.done = CustomsWindowUtilities::Done
  u.make_error = CustomsWindowUtilities::MakeError
  u.feature_add = CustomsWindowUtilities::FeatureAdd
  u.feature_hook = CustomsWindowUtilities::FeatureHook
  u.feature_init = CustomsWindowUtilities::FeatureInit
  u.fetcher = CustomsWindowUtilities::Fetcher
  u.make_fetch_def = CustomsWindowUtilities::MakeFetchDef
  u.make_context = CustomsWindowUtilities::MakeContext
  u.make_options = CustomsWindowUtilities::MakeOptions
  u.make_request = CustomsWindowUtilities::MakeRequest
  u.make_response = CustomsWindowUtilities::MakeResponse
  u.make_result = CustomsWindowUtilities::MakeResult
  u.make_point = CustomsWindowUtilities::MakePoint
  u.make_spec = CustomsWindowUtilities::MakeSpec
  u.make_url = CustomsWindowUtilities::MakeUrl
  u.param = CustomsWindowUtilities::Param
  u.prepare_auth = CustomsWindowUtilities::PrepareAuth
  u.prepare_body = CustomsWindowUtilities::PrepareBody
  u.prepare_headers = CustomsWindowUtilities::PrepareHeaders
  u.prepare_method = CustomsWindowUtilities::PrepareMethod
  u.prepare_params = CustomsWindowUtilities::PrepareParams
  u.prepare_path = CustomsWindowUtilities::PreparePath
  u.prepare_query = CustomsWindowUtilities::PrepareQuery
  u.graphql_body = CustomsWindowUtilities::GraphqlBody
  u.graphql_errors = CustomsWindowUtilities::GraphqlErrors
  u.result_basic = CustomsWindowUtilities::ResultBasic
  u.result_body = CustomsWindowUtilities::ResultBody
  u.result_headers = CustomsWindowUtilities::ResultHeaders
  u.transform_request = CustomsWindowUtilities::TransformRequest
  u.transform_response = CustomsWindowUtilities::TransformResponse
}
