# Nofrixion SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module NofrixionFeatures
  def self.make_feature(name)
    case name
    when "base"
      NofrixionBaseFeature.new
    when "test"
      NofrixionTestFeature.new
    else
      NofrixionBaseFeature.new
    end
  end
end
