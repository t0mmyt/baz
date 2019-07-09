package main

import (
	"context"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"net"
	"t0mmyt/baz/messaging"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	LISTEN   = kingpin.Arg("listen", "ip:port or :port to listen on").Envar("LISTEN").Required().String()
	DEBUG    = kingpin.Flag("debug", "Debug logging").Short('d').Bool()
	KEYWORDS = []string{"echo"}
)

type server struct {
	logger *logrus.Logger
}

func (s *server) HandleMessage(ctx context.Context, req *messaging.Request) (*messaging.Response, error) {
	log := s.logger.WithField("txn", uuidFromBytes(req.Txn))
	log.Debug(fmt.Sprintf("Received: %s", req.Message))
	var response string
	response = req.Message
	log.Debug(fmt.Sprintf("Sending: %s", response))
	return &messaging.Response{
		Txn:     req.Txn,
		Message: response,
	}, nil
}

func (s *server) GetKeywords(ctx context.Context, req *messaging.KeywordsRequest) (*messaging.KeywordsResponse, error) {
	return &messaging.KeywordsResponse{
		Keyword: KEYWORDS,
	}, nil
}

func main() {
	kingpin.Parse()
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{})
	if *DEBUG {
		logger.SetLevel(logrus.DebugLevel)
	}

	listen, err := net.Listen("tcp", *LISTEN)
	if err != nil {
		logger.Fatal(err)
	}

	srv := grpc.NewServer()
	handler := &server{logger: logger}
	messaging.RegisterKeywordsServer(srv, handler)
	messaging.RegisterMessageServer(srv, handler)
	logger.Infof("Listening on %s", *LISTEN)
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
