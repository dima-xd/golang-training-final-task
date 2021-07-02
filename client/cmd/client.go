package main

import (
	"context"
	"os"

	pb "github.com/dimaxdqwerty/golang-training-final-task/proto/go_proto"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	listen = os.Getenv("listen")
)

func init() {
	if listen == "" {
		listen = "localhost:8081"
	}
}

func main() {
	conn, err := grpc.Dial(listen, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	stream, err := pb.NewEventServiceClient(conn).GetEvent(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for {
		getEventInfo(stream)
	}
}

func getEventInfo(stream go_proto.EventService_GetEventClient) {
	event, err := stream.Recv()
	if err != nil {
		log.Fatal("Can't receive stream from server")
	}
	err = stream.Send(&pb.EventRequest{})
	if err != nil {
		log.Fatal("Can't send stream to server")
	}
	log.WithFields(log.Fields{
		"Severity": event.GetEvent().Severity,
		"Facility": event.GetEvent().Facility,
	}).Info(event.GetEvent().Message)
}
