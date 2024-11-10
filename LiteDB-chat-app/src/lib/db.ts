import { Redis } from "@upstash/redis";

export const redis = new Redis({
	url: process.env.LITEDB_REST_URL,
	token: process.env.LITEDB_REST_TOKEN,
});
