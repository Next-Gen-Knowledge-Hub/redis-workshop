local key = KEYS[1]
local count = tonumber(ARGV[1]) -- semaphor count
local ttl = tonumber(ARGV[2]) -- SECONDS

local resault = tonumber(redis.call("GET", key) or "0")

if resault < count then
    redis.call("INCR", key)
    redis.call("EXPIRE", key, ttl)
    return true
else
    return false
end
