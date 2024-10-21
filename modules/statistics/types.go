package statistics

import (
	"github.com/kirakashin/gamification/types"
)

type FireEventReq struct {
	// SendTime   int64           `json:"time"`
	ActivityID string `json:"activityID"`
	StreamUUID string `json:"streamName"`
	// UniqueUUID string          `json:"uniqueID"`
	ViewerID string          `json:"viewerId"`
	Type     types.EventType `json:"event"`
	Data     interface{}     `json:"data"`
}
