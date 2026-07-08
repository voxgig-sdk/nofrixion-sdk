# Nofrixion SDK exists test

require "minitest/autorun"
require_relative "../Nofrixion_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = NofrixionSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
