package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"t0mmyt/baz/messaging"
)

const (
	PORT = ":8080"
)

var (
	KEYWORDS = []string{"echo"}
)

type server struct {
	logger *logrus.Logger
}

func (s *server) HandleMessage(ctx context.Context, req *messaging.Request) (*messaging.Response, error) {
	s.logger.WithField("txn", uuidFromBytes(req.Txn)).Debug(req.Message)
	return &messaging.Response{
		Txn:     req.Txn,
		Message: req.Message,
	}, nil
}

func (s *server) GetKeywords(ctx context.Context, req *messaging.KeywordsRequest) (*messaging.KeywordsResponse, error) {
	return &messaging.KeywordsResponse{
		Keyword: KEYWORDS,
	}, nil
}

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		logger.Fatal(err)
	}

	srv := grpc.NewServer()
	handler := &server{logger: logger}
	messaging.RegisterKeywordsServer(srv, handler)
	messaging.RegisterMessageServer(srv, handler)
	logger.Infof("Listening on %s", PORT)
	if err := srv.Serve(listen); err != nil {
		logger.Fatal(err)
	}
}

func uuidFromBytes(b []byte) string {
	u, err := uuid.FromBytes(b)
	if err != nil {
		return "none"
	}
	return u.String()
}