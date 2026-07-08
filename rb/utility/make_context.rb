# Nofrixion SDK utility: make_context
require_relative '../core/context'
module NofrixionUtilities
  MakeContext = ->(ctxmap, basectx) {
    NofrixionContext.new(ctxmap, basectx)
  }
end
