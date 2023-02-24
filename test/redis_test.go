package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

// Background()返回上下文
var ctx = context.Background()

// NewClient()返回一个客户端到Redis服务器
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestRedisSet(t *testing.T) {
	// 上下文 key value 过期时间
	rdb.Set(ctx, "name", "mmc", time.Second*10)
}

func TestRedisGet(t *testing.T) {
	v, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(v)
}
