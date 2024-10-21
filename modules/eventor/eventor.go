package eventor

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kirakashin/gamification/cache"
)

var chatToActivity = cache.InitCacheBucket(time.Hour)
var hallToActivity = cache.InitCacheBucket(time.Hour)

const (
	// PING_PATH = "/health"
	GET_ACTIVITY_BY_CHAT_PATH       = "/activity/chat/%s?token=%s"
	GET_ACTIVITY_ID_BY_HALL_ID_PATH = "/hall/%d?token=%s"
	GET_ACTIVITY_PATH               = "/activity/%d?token=%s"
)

type EventorService struct {
	URL   string
	Token string
}

func InitService(url, token string) (es EventorService, err error) {
	es.URL = url
	es.Token = token

	// err = es.Ping()
	// if err != nil {
	// 	return EventorService{}, err
	// }

	return
}

// func (es *EventorService) Ping() error {
// 	fullURL, _ := url.JoinPath(es.URL, PING_PATH)

// 	resp, err := http.Get(fullURL)
// 	if err != nil {
// 		return err
// 	}

// 	if resp.StatusCode != 200 {
// 		return fmt.Errorf("UNAUTHORIZED")
// 	}

// 	return nil
// }

func (es *EventorService) TranslateChatToActivity(chatRoomUUID string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	isCached, data := chatToActivity.Get(ctx, chatRoomUUID)
	if isCached {
		dataP := data.(map[string]string)
		return dataP["activityID"], dataP["streamUUID"], nil
	}

	activity, err := es.GetActivityByChat(chatRoomUUID)
	if err != nil {
		return "", "", err
	}

	chatToActivity.Put(chatRoomUUID, map[string]string{"activityID": fmt.Sprint(activity.ActivityID), "streamUUID": activity.Halls[0].Stream.UUIDStream})

	return fmt.Sprint(activity.ActivityID), activity.Halls[0].Stream.UUIDStream, nil
}

func (es *EventorService) GetActivityByChat(chatRoomUUID string) (*Activity, error) {
	fullURL := fmt.Sprintf(es.URL+GET_ACTIVITY_BY_CHAT_PATH, chatRoomUUID, es.Token)

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("INTERNAL ERROR")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var activity struct {
		Activity Activity `json:"activity"`
	}

	err = json.Unmarshal(b, &activity)
	if err != nil {
		return nil, err
	}

	return &activity.Activity, nil
}

func (es *EventorService) TranslateHallIDToActivity(hallID uint) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	isCached, data := hallToActivity.Get(ctx, fmt.Sprint(hallID))
	if isCached {
		dataP := data.(map[string]string)
		return dataP["activityID"], dataP["streamUUID"], nil
	}

	activityID, err := es.GetActivityIDByHallID(hallID)
	if err != nil {
		return "", "", err
	}

	activity, err := es.GetActivityByActivityID(*activityID)
	if err != nil {
		return "", "", err
	}

	hallToActivity.Put(fmt.Sprint(hallID), map[string]string{"activityID": fmt.Sprint(activity.ActivityID), "streamUUID": activity.Halls[0].Stream.UUIDStream})

	return fmt.Sprint(activity.ActivityID), activity.Halls[0].Stream.UUIDStream, nil
}

func (es *EventorService) GetActivityIDByHallID(hallID uint) (*uint, error) {
	fullURL := fmt.Sprintf(es.URL+GET_ACTIVITY_ID_BY_HALL_ID_PATH, hallID, es.Token)

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("INTERNAL ERROR")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var hall struct {
		Hall Hall `json:"hall"`
	}

	err = json.Unmarshal(b, &hall)
	if err != nil {
		return nil, err
	}

	return &hall.Hall.ActivityID, nil
}

func (es *EventorService) GetActivityByActivityID(activityID uint) (*Activity, error) {
	fullURL := fmt.Sprintf(es.URL+GET_ACTIVITY_PATH, activityID, es.Token)

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("INTERNAL ERROR")
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var activity struct {
		Activity Activity `json:"activity"`
	}

	err = json.Unmarshal(b, &activity)
	if err != nil {
		return nil, err
	}

	return &activity.Activity, nil
}
