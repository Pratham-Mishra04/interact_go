package subscribers

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Pratham-Mishra04/interact/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var ctx = context.Background()

func ImpressionsDumpSub(client *redis.Client, db *gorm.DB) { //TODO2 add logger here for errors
	if client == nil {
		fmt.Println("Error Subscribing to Redis Expiration Event for Impressions Dump: ", "redis client is nil")
		return
	}
	if db == nil {
		fmt.Println("Error Subscribing to Redis Expiration Event for Impressions Dump: ", "gorm db is nil")
		return
	}

	RedisExpirationSub := client.Subscribe(ctx, "__keyevent@0__:expired")

	// Wait for confirmation that subscription is created before publishing anything
	_, err := RedisExpirationSub.Receive(ctx)
	if err != nil {
		fmt.Println("Error Subscribing to Redis Expiration Event for Impressions Dump: ", err)
	} else {
		if RedisExpirationSub == nil {
			fmt.Println("Error Subscribing to Redis Expiration Event for Impressions Dump: ", "redis subscriber is nil")
			return
		}

		fmt.Println("Subscribed to Redis Expiration Event for Impressions Dump!")

		ch := RedisExpirationSub.Channel()

		for msg := range ch {
			modelName := extractModelDataFromKey(msg.Payload, 1)
			modelID := extractModelDataFromKey(msg.Payload, 2)

			if modelName == "" || modelID == "" {
				continue
			}

			model := getModelFromStr(modelName)

			if err := db.Model(model).Where("id = ?", modelID).UpdateColumn("Impressions", gorm.Expr("Impressions + ?", rand.Intn(5)+3)).Error; err != nil {
				fmt.Printf("\n Error dumping impressions of key:%s, error:%e", msg.Payload, err)
			} else {
				fmt.Printf("\nKey %s dumped\n", msg.Payload)
			}
		}
	}
}

func extractModelDataFromKey(key string, index int) string {
	parts := strings.Split(key, "_")
	if len(parts) >= 3 {
		if parts[0] != "impressions" {
			return ""
		}
		return parts[index]
	}

	return ""
}

func getModelFromStr(modelName string) interface{} {
	switch modelName {
	case "Post":
		return models.Post{}
	case "Project":
		return models.Project{}
	case "Event":
		return models.Event{}
	case "Opening":
		return models.Opening{}
	case "User":
		return models.User{}
	}

	return nil
}
