package gamification

import (
	"github.com/kirakashin/gamification/modules/eventor"
	"github.com/kirakashin/gamification/modules/statistics"
	"github.com/kirakashin/gamification/types"
)

type GameConnection struct {
	StatisticsService statistics.StatisticsService
	EventorService    eventor.EventorService
}

func InitConnection(statisticsURL, eventorURL, eventorToken string) (c *GameConnection, err error) {
	c = &GameConnection{}
	c.StatisticsService, _ = statistics.InitService(statisticsURL)
	c.EventorService, _ = eventor.InitService(eventorURL, eventorToken)

	return
}

func (c *GameConnection) SendMessageEvent(viewerID, chatRoomUUID string, payload interface{}) (err error) {
	activityID, streamUUID, err := c.EventorService.TranslateChatToActivity(chatRoomUUID)
	if err != nil {
		return
	}

	err = c.StatisticsService.FireEvent(viewerID, activityID, streamUUID, types.EVENT_TYPE_MESSAGE, payload)
	if err != nil {
		return
	}

	return
}

func (c *GameConnection) SendRejectMessageEvent(viewerID, chatRoomUUID string, payload interface{}) (err error) {
	activityID, streamUUID, err := c.EventorService.TranslateChatToActivity(chatRoomUUID)
	if err != nil {
		return
	}

	err = c.StatisticsService.FireEvent(viewerID, activityID, streamUUID, types.EVENT_TYPE_MESSAGE_REJECTED, payload)
	if err != nil {
		return
	}

	return
}

func (c *GameConnection) SendLikeMessageEvent(viewerID, chatRoomUUID string, payload interface{}) (err error) {
	activityID, streamUUID, err := c.EventorService.TranslateChatToActivity(chatRoomUUID)
	if err != nil {
		return
	}

	err = c.StatisticsService.FireEvent(viewerID, activityID, streamUUID, types.EVENT_TYPE_LIKE_MESSAGE, payload)
	if err != nil {
		return
	}

	return
}

func (c *GameConnection) SendDislikeMessageEvent(viewerID, chatRoomUUID string, payload interface{}) (err error) {
	activityID, streamUUID, err := c.EventorService.TranslateChatToActivity(chatRoomUUID)
	if err != nil {
		return
	}

	err = c.StatisticsService.FireEvent(viewerID, activityID, streamUUID, types.EVENT_TYPE_DISLIKE_MESSAGE, payload)
	if err != nil {
		return
	}

	return
}

func (c *GameConnection) SendAnswerQuestionEvent(viewerID, chatRoomUUID string, payload interface{}) (err error) {
	activityID, streamUUID, err := c.EventorService.TranslateChatToActivity(chatRoomUUID)
	if err != nil {
		return
	}

	err = c.StatisticsService.FireEvent(viewerID, activityID, streamUUID, types.EVENT_TYPE_ANSWER_QUESTION, payload)
	if err != nil {
		return
	}

	return
}

func (c *GameConnection) SendAnswerQuestionRejectedEvent(viewerID, chatRoomUUID string, payload interface{}) (err error) {
	activityID, streamUUID, err := c.EventorService.TranslateChatToActivity(chatRoomUUID)
	if err != nil {
		return
	}

	err = c.StatisticsService.FireEvent(viewerID, activityID, streamUUID, types.EVENT_TYPE_ANSWER_QUESTION_REJECTED, payload)
	if err != nil {
		return
	}

	return
}

func (c *GameConnection) SendPollVoteEvent(viewerID string, hallID uint, payload interface{}) (err error) {
	activityID, streamUUID, err := c.EventorService.TranslateHallIDToActivity(hallID)
	if err != nil {
		return
	}

	err = c.StatisticsService.FireEvent(viewerID, activityID, streamUUID, types.EVENT_TYPE_POLL_VOTE, payload)
	if err != nil {
		return
	}

	return
}

func (c *GameConnection) SendCorrectPollVoteEvent(viewerID string, hallID uint, payload interface{}) (err error) {
	activityID, streamUUID, err := c.EventorService.TranslateHallIDToActivity(hallID)
	if err != nil {
		return
	}

	err = c.StatisticsService.FireEvent(viewerID, activityID, streamUUID, types.EVENT_TYPE_CORRECT_POLL_VOTE, payload)
	if err != nil {
		return
	}

	return
}
