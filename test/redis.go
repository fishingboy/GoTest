package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "mykey", "value", 30*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("mykey", val)

	val2, err := rdb.Get(ctx, "productStat").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// 檢查剩餘的過期時間
	ttl, err := rdb.TTL(ctx, "mykey").Result()
	if err != nil {
		fmt.Println("Failed to get TTL:", err)
		return
	}

	for {
		time.Sleep(10 * time.Second)
		// 延長過期時間，增加 20 秒
		err = rdb.Expire(ctx, "mykey", 30*time.Second).Err()
		if err != nil {
			fmt.Println("Failed to extend expiration:", err)
			return
		}

		// 再次檢查剩餘的過期時間
		ttl, err = rdb.TTL(ctx, "mykey").Result()
		if err != nil {
			fmt.Println("Failed to get TTL:", err)
			return
		}
		fmt.Printf("Extended TTL: %v\n", ttl)
	}

	// Output: key value
	// key2 does not exist
}
