package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
	"time"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})
var cbk = context.Background()

func TestGoRedis(t *testing.T) {

}

func TestRedisSet(t *testing.T) {
	rdb.Set(cbk, "email", "aaa", time.Second*10)
}

func TestRedisGet(t *testing.T) {
	result, err := rdb.Get(cbk, "email").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("结果为:", result)
}
