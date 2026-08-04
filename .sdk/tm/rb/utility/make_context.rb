# CustomsWindow SDK utility: make_context
require_relative '../core/context'
module CustomsWindowUtilities
  MakeContext = ->(ctxmap, basectx) {
    CustomsWindowContext.new(ctxmap, basectx)
  }
end
