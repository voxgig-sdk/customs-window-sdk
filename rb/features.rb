# CustomsWindow SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/debug_feature'
require_relative 'feature/idempotency_feature'
require_relative 'feature/metrics_feature'
require_relative 'feature/paging_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module CustomsWindowFeatures
  def self.make_feature(name)
    case name
    when "base"
      CustomsWindowBaseFeature.new
    when "debug"
      CustomsWindowDebugFeature.new
    when "idempotency"
      CustomsWindowIdempotencyFeature.new
    when "metrics"
      CustomsWindowMetricsFeature.new
    when "paging"
      CustomsWindowPagingFeature.new
    when "ratelimit"
      CustomsWindowRatelimitFeature.new
    when "retry"
      CustomsWindowRetryFeature.new
    when "test"
      CustomsWindowTestFeature.new
    when "timeout"
      CustomsWindowTimeoutFeature.new
    else
      CustomsWindowBaseFeature.new
    end
  end
end
