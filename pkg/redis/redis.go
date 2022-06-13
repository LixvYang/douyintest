package redis

import (
	"context"
	"fmt"
	"tiktok/utils"
	"time"

	"github.com/go-redis/redis/v8"
)



var (
	ctx = context.Background()
 	redisClient *redis.Client
)

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     utils.RedisHost + ":" + utils.RedisPort,
		Password: utils.RedisPassword, // no password set
		DB:       0,        // use default DB
	})
}

func RCGet(key string) *redis.StringCmd {
	return redisClient.Get(ctx, key)
}
func RCExists(key string) bool {
	return redisClient.Exists(ctx, key).Val() != 0
}
func RCSet(key string, value interface{}, expiration time.Duration) {
	if RCExists(key) {
		redisClient.Expire(ctx, key, expiration)
		return
	}
	redisClient.Set(ctx, key, value, expiration)
}
func RCIncrement(key string) {
	redisClient.Incr(ctx, key)
}

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func RCSAdd(key string, members interface{}) {
	redisClient.SAdd(ctx, key, members)
}

func RCSRem(key string, members interface{}) {
	redisClient.SRem(ctx, key, members)
}

func RCSmembers(key string) *redis.StringSliceCmd {
	return redisClient.SMembers(ctx, key)
}
