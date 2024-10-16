local key = KEYS[1]
local val = ARGV[1]
local ttl = tonumber(ARGV[2]) -- SECONDS

local resault = redis.call("SETNX", key, val)

-- if we could set the value then we will return true
if resault == 1 then
    redis.call("EXPIRE", key, ttl)
    return true
end

-- the key already exists and we could not set the key
return false
