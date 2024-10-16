local key = KEYS[1]

local resault = tonumber(redis.call("GET", key) or "0")

if resault <= 0 then
    return false
else
    redis.call("DECR", key)
    return true
end
