package statistics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kirakashin/gamification/types"
)

const (
	// PING_PATH       = "/ping"
	FIRE_EVENT_PATH = "/events/stream/%s/fire"
)

type StatisticsService struct {
	URL string
}

func InitService(url string) (ss StatisticsService, err error) {
	ss.URL = url

	// err = ss.Ping()
	// if err != nil {
	// 	return StatisticsService{}, err
	// }

	return
}

// func (ss *StatisticsService) Ping() error {
// 	fullURL, _ := url.JoinPath(ss.URL, PING_PATH)

// 	resp, err := http.Get(fullURL)
// 	if err != nil {
// 		return err
// 	}

// 	if resp.StatusCode != 200 {
// 		return fmt.Errorf("UNAUTHORIZED")
// 	}

// 	return nil
// }

func (ss *StatisticsService) FireEvent(viewerID, activityID, streamUUID string, eventType types.EventType, payload interface{}) error {
	fullURL, _ := url.JoinPath(ss.URL, FIRE_EVENT_PATH)

	fullURL = fmt.Sprintf(fullURL, streamUUID)

	req := FireEventReq{
		ActivityID: activityID,
		StreamUUID: streamUUID,
		UniqueUUID: streamUUID,
		ViewerID:   viewerID,
		Type:       eventType,
		Data:       payload,
	}

	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	resp, err := http.Post(fullURL, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("INTERNAL ERROR")
	}

	return nil
}
