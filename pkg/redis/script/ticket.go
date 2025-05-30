package script

// script decrease ticket item lua in redis
const DecreaseTicketItemLua = `
	local cTicket = tonumber(redis.call("GET", KEYS[1]));
	if not cTicket then
		return 0;
	end;
	if cTicket < tonumber(ARGV[1]) then
		return 0;
	end;
	cTicket = cTicket - tonumber(ARGV[1]);
	redis.call("SET", KEYS[1], cTicket);
	redis.call("EXPIRE", KEYS[1], 60 * 60 * 8) -- set expire time 8 hour
	return 1;
`