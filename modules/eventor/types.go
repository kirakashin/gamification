package eventor

type Activity struct {
	ActivityID uint64 `json:"activityID"`
	Halls      []struct {
		Stream struct {
			UUIDStream string `json:"streamUUID"`
		} `json:"stream"`
	} `json:"halls"`
}
