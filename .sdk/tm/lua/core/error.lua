-- Nofrixion SDK error

local NofrixionError = {}
NofrixionError.__index = NofrixionError


function NofrixionError.new(code, msg, ctx)
  local self = setmetatable({}, NofrixionError)
  self.is_sdk_error = true
  self.sdk = "Nofrixion"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function NofrixionError:error()
  return self.msg
end


function NofrixionError:__tostring()
  return self.msg
end


return NofrixionError
