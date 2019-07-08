package main

import (
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"t0mmyt/baz/servicedispatch"
)

var (
	TOKEN = kingpin.Arg("token", "Slack token").Envar("TOKEN").Required().String()
	DEBUG = kingpin.Arg("debug", "Debug logging").Default("false").Bool()
)

func main() {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})
	kingpin.Parse()
	if *DEBUG {
		logger.SetLevel(log.DebugLevel)
	} else {
		logger.SetLevel(log.InfoLevel)
	}

	api := slack.New(*TOKEN)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	m := servicedispatch.NewMessageHandler(rtm, logger)
	err := m.AddService("localhost:8080")
	if err != nil {
		log.Errorf("adding services: %s", err)
	}

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

			case *slack.HelloEvent:
				logger.Info("Connected")

			default:
				if evt.Type != "latency_report" {
					logger.Debugf("Event: %s\n", evt.Type)
				}
			}
		}
	}
	logger.Info("Fin.")
}
