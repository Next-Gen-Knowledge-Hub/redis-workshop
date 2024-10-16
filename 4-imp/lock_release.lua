local key = KEYS[1]
local val = ARGV[1]

local resault = redis.call("GET", key)

-- check if we could release the lock or not
if resault == val then
    -- delete the key
    local delResault = redis.call("DEL", key)
    if delResault == 1 then
        return true
    end

    -- we couldn't remove the key !!
    error("you aren't lock owner")
end

return false
