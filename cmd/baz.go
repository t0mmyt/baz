package main

import (
	"fmt"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"t0mmyt/baz/servicediscovery"
)

var (
	TOKEN = kingpin.Arg("token", "Slack token").Envar("TOKEN").Required().String()
)

func main() {
	kingpin.Parse()
	fmt.Printf("Token: '%s'\n", *TOKEN)
	api := slack.New(*TOKEN)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	m := servicediscovery.NewMessageHandler(rtm)
	m.AddKeyword("ping", "pingservice")

MainEvtLoop:
	for {
		select {
		case evt := <-rtm.IncomingEvents:
			switch et := evt.Data.(type) {
			case *slack.MessageEvent:
				text := et.Text
				_ = m.HandleMessage(text, et.Channel)

			case *slack.InvalidAuthEvent:
				log.Errorln("Auth Error")
				break MainEvtLoop
			default:
				if evt.Type != "latency_report" {
					fmt.Printf("Event: %s\n", evt.Type)
				}
			}
		}
	}
	fmt.Println("Fin.")
}
