package clusterClient

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func ReadAndWriteFromCluster() {

	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"localhost:7000",
			"localhost:7001",
			"localhost:7002",
			"localhost:7003",
			"localhost:7004",
			"localhost:7005",
		},
	})

	_, err := clusterClient.Ping(context.TODO()).Result()
	if err != nil {
		log.Panic(err)
	}

	err = clusterClient.Set(context.TODO(), "test", "test-val", time.Second*10).Err()
	if err != nil {
		log.Panic(err)
	}

	val, err := clusterClient.Get(context.TODO(), "test").Result()
	if err != nil {
		log.Panic(err)
	}

	log.Println(val)
}
