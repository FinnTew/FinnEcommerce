package util

import "strconv"

type RedisKey struct{}

func NewRedisKey() *RedisKey {
	return &RedisKey{}
}

func (r *RedisKey) GetUserTokenKey(userId uint32) string {
	return "user_token:" + strconv.Itoa(int(userId))
}
