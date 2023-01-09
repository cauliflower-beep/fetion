package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
)

var ctx = context.Background()

func TestClient(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// set
	err := rdb.Set(ctx, "龙珠: 赛亚人", "贝吉塔", 0).Err()
	if err != nil {
		panic(err)
	}

	// get
	//val, err := rdb.Get(ctx, "龙珠").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//val2, err := rdb.Get(ctx, "龙珠gt").Result()
	//if err == redis.Nil {
	//	fmt.Println("龙珠gt does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
}
