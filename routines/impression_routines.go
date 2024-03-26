package routines

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/Pratham-Mishra04/interact/cache"
	"github.com/Pratham-Mishra04/interact/helpers"
	"github.com/Pratham-Mishra04/interact/initializers"
	"github.com/Pratham-Mishra04/interact/models"
	"gorm.io/gorm"
)

type IncrementFunc func(id string, ch chan<- uint)

func IncrementImpressions(items interface{}, getModelID func(interface{}) string, incrementDB IncrementFunc, modelType interface{}) {
	workerCount := 5
	var itemIDs []string

	// Use reflection to get the underlying type of items
	itemsValue := reflect.ValueOf(items)
	if itemsValue.Kind() != reflect.Slice {
		fmt.Printf("Unsupported type: %T\n", items)
		return
	}

	for i := 0; i < itemsValue.Len(); i++ {
		item := itemsValue.Index(i).Interface()
		key := getModelID(item)
		suffix := getModelTypeStr(item)
		key = suffix + "_" + key

		impressionCount, err := cache.GetImpression(key)
		if err != nil {
			return
		} else if impressionCount >= 9 {
			itemIDs = append(itemIDs, key)
			checkForNotification(item, modelType, impressionCount)
			go cache.ResetImpression(key)
		} else {
			go cache.IncrementImpression(key)
		}
	}

	incrementImpressionsConcurrently(itemIDs, workerCount, incrementDB, modelType)
}

func incrementImpressionsConcurrently(ids []string, workerCount int, incrementDB IncrementFunc, modelType interface{}) {
	done := make(chan uint, workerCount)
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			incrementDB(id, done)
		}(id)
	}
	go func() {
		wg.Wait()
		close(done)
	}()
}

func incrementDBImpressions(modelType interface{}, itemID string, ch chan<- uint) {
	result := initializers.DB.Model(modelType).Where("id = ?", itemID).UpdateColumn("Impressions", gorm.Expr("Impressions + ?", 10))
	if result.Error != nil {
		typeName := reflect.TypeOf(modelType).Elem().Name()
		helpers.LogDatabaseError(fmt.Sprintf("Error updating %sImpressionCount", typeName), result.Error, "impression_routines")
		return
	}
	ch <- 1
}

// Posts
func IncrementPostImpression(posts []models.Post) {
	IncrementImpressions(posts, func(item interface{}) string {
		return item.(models.Post).ID.String()
	}, func(postID string, ch chan<- uint) {
		incrementDBImpressions(&models.Post{}, postID, ch)
	}, &models.Post{})
}

// Projects
func IncrementProjectImpression(projects []models.Project) {
	IncrementImpressions(projects, func(item interface{}) string {
		return item.(models.Project).ID.String()
	}, func(projectID string, ch chan<- uint) {
		incrementDBImpressions(&models.Project{}, projectID, ch)
	}, &models.Project{})
}

// Events
func IncrementEventImpression(events []models.Event) {
	IncrementImpressions(events, func(item interface{}) string {
		return item.(models.Event).ID.String()
	}, func(eventID string, ch chan<- uint) {
		incrementDBImpressions(&models.Event{}, eventID, ch)
	}, &models.Event{})
}

// Openings
func IncrementOpeningImpression(openings []models.Opening) {
	IncrementImpressions(openings, func(item interface{}) string {
		return item.(models.Opening).ID.String()
	}, func(openingID string, ch chan<- uint) {
		incrementDBImpressions(&models.Opening{}, openingID, ch)
	}, &models.Opening{})
}

// Users
func IncrementUserImpression(users []models.User) {
	IncrementImpressions(users, func(item interface{}) string {
		return item.(models.User).ID.String()
	}, func(userID string, ch chan<- uint) {
		incrementDBImpressions(&models.User{}, userID, ch)
	}, &models.User{})
}

func checkForNotification(item interface{}, modelType interface{}, cacheImpressionCount int) {
	switch modelType.(type) {
	case *models.Post:
		post := item.(models.Post)
		go sendImpressionNotification(post.UserID, post.UserID, &post.ID, nil, nil, post.Impressions+cacheImpressionCount+1)
	case *models.Project:
		project := item.(models.Project)
		go sendImpressionNotification(project.UserID, project.UserID, nil, &project.ID, nil, project.Impressions+cacheImpressionCount+1)
	case *models.Event:
		event := item.(models.Event)
		go sendImpressionNotification(event.Organization.UserID, event.Organization.UserID, nil, nil, &event.ID, event.Impressions+cacheImpressionCount+1)
	}
}

func getModelTypeStr(item interface{}) string {
    t := reflect.TypeOf(item)
    if t.Kind() == reflect.Ptr {
        t = t.Elem()
    }
    return t.Name()
}