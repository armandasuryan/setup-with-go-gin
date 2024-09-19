package rds

import (
	"context"
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/go-redis/redis/v8"
)

func RedisConnect(host string, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})
	ctx := context.TODO()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("error redis:", err)

	} else {
		connectFigure := figure.NewColorFigure("Redis connect", "", "red", true)
		connectFigure.Print()
	}
	return client

}
