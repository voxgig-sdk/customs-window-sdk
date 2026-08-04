-- CustomsWindow SDK error

local CustomsWindowError = {}
CustomsWindowError.__index = CustomsWindowError


function CustomsWindowError.new(code, msg, ctx)
  local self = setmetatable({}, CustomsWindowError)
  self.is_sdk_error = true
  self.sdk = "CustomsWindow"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function CustomsWindowError:error()
  return self.msg
end


function CustomsWindowError:__tostring()
  return self.msg
end


return CustomsWindowError
