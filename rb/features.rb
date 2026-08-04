# CustomsWindow SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module CustomsWindowFeatures
  def self.make_feature(name)
    case name
    when "base"
      CustomsWindowBaseFeature.new
    when "test"
      CustomsWindowTestFeature.new
    else
      CustomsWindowBaseFeature.new
    end
  end
end
