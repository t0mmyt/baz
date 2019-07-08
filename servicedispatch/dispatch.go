package servicedispatch

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/jdkato/prose.v2"
	"strings"
	"t0mmyt/baz/messaging"
	"time"
)

type MessageHandler struct {
	logger   *logrus.Logger
	rtm      *slack.RTM
	services map[string]string
}

func NewMessageHandler(rtm *slack.RTM, logger *logrus.Logger) *MessageHandler {
	return &MessageHandler{
		logger:   logger,
		services: make(map[string]string),
		rtm:      rtm,
	}
}

func (m MessageHandler) AddService(host string) error {
	txn, err := randomUuid()
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("connecting to host %s: %s", host, err)
	}
	defer conn.Close()

	client := messaging.NewKeywordsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req, err := client.GetKeywords(ctx, &messaging.KeywordsRequest{Txn: *txn})
	if err != nil {
		return fmt.Errorf("getting messaging from host %s: %s", host, err)
	}
	m.logger.Infof("Got: [%s] from %s", strings.Join(req.Keyword, ", "), host)
	for _, v := range req.Keyword {
		m.services[v] = host
	}
	return nil
}

func (m MessageHandler) HandleMessage(msg, channel string) error {
	var response string

	tokens, err := tokenise(msg)
	if err != nil {
		return err
	}

	switch (*tokens)[0] {
	case "ping":
		response = "pong"

	case "messaging":
		keys := make([]string, len(m.services))
		i := 0
		for k := range m.services {
			keys[i] = k
			i++
		}
		response = strings.Join(keys, ", ")

	default:
		tokens, err := tokenise(msg)
		if err != nil {
			return err
		}

		if responseFromMap, ok := m.services[(*tokens)[0]]; ok {
			var err error
			response, err = m.sendMessage(responseFromMap, msg)
			if err != nil {
				return err
			}
		} else {
			response = "Fuck knows?"
		}
	}
	m.logger.Infof("Sent: %s", response)
	m.rtm.SendMessage(m.rtm.NewOutgoingMessage(response, channel))
	return nil
}

func (m *MessageHandler) sendMessage(host, msg string) (string, error) {
	txn, err := randomUuid()
	if err != nil {
		return "", err
	}

	var response string
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return "", nil
	}
	defer conn.Close()

	client := messaging.NewMessageClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.HandleMessage(ctx, &messaging.Request{
		Txn:     *txn,
		Message: msg,
	})
	if err != nil {
		return response, err
	}
	return res.Message, nil
}

func tokenise(msg string) (*[]string, error) {
	doc, err := prose.NewDocument(msg, prose.WithExtraction(false))
	if err != nil {
		return nil, err
	}

	tokens := doc.Tokens()
	texts := make([]string, len(tokens))

	for i := range texts {
		texts[i] = strings.ToLower(tokens[i].Text)
	}

	return &texts, nil
}

func randomUuid() (*[]byte, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	txn, err := u.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return &txn, nil
}