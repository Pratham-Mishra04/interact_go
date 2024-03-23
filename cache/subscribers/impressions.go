package subscribers

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Pratham-Mishra04/interact/helpers"
	"github.com/Pratham-Mishra04/interact/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"errors"
)

var ctx = context.Background()

func ImpressionsDumpSub(client *redis.Client, db *gorm.DB) { 
	if client == nil {
		go helpers.LogServerError("Error Subscribing to Redis Expiration Event for Impressions Dump: ", errors.New("redis client is nil"), "")
		return
	}
	if db == nil {
		go helpers.LogServerError("Error Subscribing to Redis Expiration Event for Impressions Dump: ", errors.New("gorm db is nil"), "")
		return
	}

	RedisExpirationSub := client.Subscribe(ctx, "__keyevent@0__:expired")

	// Wait for confirmation that subscription is created before publishing anything
	_, err := RedisExpirationSub.Receive(ctx)
	if err != nil {
		go helpers.LogServerError("Error Subscribing to Redis Expiration Event for Impressions Dump: ", err, "")
	} else {
		if RedisExpirationSub == nil {
			go helpers.LogServerError("Error Subscribing to Redis Expiration Event for Impressions Dump: ", errors.New("redis subscriber is nil"), "")
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
				helpers.LogDatabaseError("Error updating impressions of key:"+msg.Payload, err, "")
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
