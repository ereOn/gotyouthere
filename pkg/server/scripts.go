package server

const (
	LockScript = `
local presence_key = KEYS[1];
local lock_key = KEYS[2];
local waits_key = KEYS[3];
local unblocked_waits_key = KEYS[4];

local id = ARGV[1];
local presence = ARGV[2];

local current_presence = tonumber(redis.call('HGET', presence_key, id) or "0");

if presence == current_presence then
	redis.call('RPUSH', unblocked_waits_key, "x");
	return current_presence;
end

local locked = tonumber(redis.call('HGET', lock_key, id) or "0");

if locked > 0 then
	redis.call('RPUSH', waits_key, "x");
	return current_presence
else
	redis.call('HSET', lock_key, id, "1");
	redis.call('DEL', unblocked_waits_key, waits_key);
	redis.call('RPUSH', unblocked_waits_key, "x");
	return current_presence;
end
`
	UnlockScript = `
local presence_key = KEYS[1];
local lock_key = KEYS[2];
local waits_key = KEYS[3];
local unblocked_waits_key = KEYS[4];

local id = ARGV[1];
local presence = ARGV[2];

local locked = tonumber(redis.call('HGET', lock_key, id) or "0");

if locked > 0 then
	local waits_count = redis.call('LLEN', waits_key);
	redis.call('DEL', waits_key)

	for i=1,waits_count do
		redis.call('RPUSH', unblocked_waits_key, "x");
	end

	redis.call('HDEL', lock_key, id);
end

return ""
`
)
