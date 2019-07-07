package servicediscovery

import (
	"context"
	"fmt"
	"github.com/nlopes/slack"
	"gopkg.in/jdkato/prose.v2"
	"strings"
)

type MessageHandler struct {
	ctx      context.Context
	rtm      *slack.RTM
	keywords map[string]string
}

func NewMessageHandler(rtm *slack.RTM) *MessageHandler {
	return &MessageHandler{
		keywords: make(map[string]string),
		rtm: rtm,
	}
}

func (m MessageHandler) AddKeyword(key, value string) {
	m.keywords[key] = value
}

func (m MessageHandler) HandleMessage(msg, channel string) error {
	var response string
	switch msg {
	case "ping":
		response = "pong"
	default:
		response = "Fuck knows?"
	}
	m.rtm.SendMessage(m.rtm.NewOutgoingMessage(response, channel))
	return nil
}

func HandleMessage(msg string) error {
	doc, err := prose.NewDocument(msg, prose.WithExtraction(false))
	if err != nil {
		return err
	}

	tokens := doc.Tokens()
	texts := make([]string, len(tokens))
	for i := range texts {
		texts[i] = tokens[i].Text
		fmt.Printf("%s", tokens[i].Label)
	}
	fmt.Println()
	fmt.Println(strings.Join(texts, " "))

	return nil
}
