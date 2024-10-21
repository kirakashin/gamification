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
	hallID := flag.Int("hallID", 0, "hallID")

	flag.Parse()

	fmt.Printf("statisticsURL: %v\n", *statisticsURL)
	fmt.Printf("eventorURL: %v\n", *eventorURL)
	fmt.Printf("eventorToken: %v\n", *eventorToken)
	fmt.Printf("viewerID: %v\n", *viewerID)
	fmt.Printf("hallID: %v\n", *hallID)

	var err error

	conn, _ := gamification.InitConnection(*statisticsURL, *eventorURL, *eventorToken)

	err = conn.SendPollVoteEvent(*viewerID, uint(*hallID), map[string]string{"message": "test SendPollVoteEvent"})
	fmt.Printf("err: %v\n", err)

	err = conn.SendCorrectPollVoteEvent(*viewerID, uint(*hallID), map[string]string{"message": "test SendCorrectPollVoteEvent"})
	fmt.Printf("err: %v\n", err)

}
