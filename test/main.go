package main

import (
	"flag"
	"fmt"

	"github.com/kirakashin/gamification"
)

func main() {
	statisticsURL := flag.String("statisticsURL", "", "statisticsURL")
	eventorURL := flag.String("eventorURL", "", "eventorURL")
	eventorToken := flag.String("eventorToken", "", "eventorToken")
	viewerID := flag.String("viewerID", "", "viewerID")
	chatUUID := flag.String("chatUUID", "", "chatUUID")

	flag.Parse()

	fmt.Printf("statisticsURL: %v\n", *statisticsURL)
	fmt.Printf("eventorURL: %v\n", *eventorURL)
	fmt.Printf("eventorToken: %v\n", *eventorToken)
	fmt.Printf("viewerID: %v\n", *viewerID)
	fmt.Printf("chatUUID: %v\n", *chatUUID)

	var err error

	conn, _ := gamification.InitConnection(*statisticsURL, *eventorURL, *eventorToken)

	err = conn.SendMessageEvent(*viewerID, *chatUUID, map[string]string{"message": "test"})
	fmt.Printf("err: %v\n", err)

	err = conn.SendRejectMessageEvent(*viewerID, *chatUUID, map[string]string{"message": "test"})
	fmt.Printf("err: %v\n", err)

	err = conn.SendLikeMessageEvent(*viewerID, *chatUUID, map[string]string{"message": "test"})
	fmt.Printf("err: %v\n", err)

	err = conn.SendDislikeMessageEvent(*viewerID, *chatUUID, map[string]string{"message": "test"})
	fmt.Printf("err: %v\n", err)

	err = conn.SendAnswerQuestionEvent(*viewerID, *chatUUID, map[string]string{"message": "test"})
	fmt.Printf("err: %v\n", err)

	err = conn.SendAnswerQuestionRejectedEvent(*viewerID, *chatUUID, map[string]string{"message": "test"})
	fmt.Printf("err: %v\n", err)
}
