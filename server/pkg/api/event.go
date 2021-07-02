package api

import (
	"fmt"

	pb "github.com/dimaxdqwerty/golang-training-final-task/proto/go_proto"
	log "github.com/sirupsen/logrus"

	"gopkg.in/mcuadros/go-syslog.v2"
)

type EventServer struct {
	syslogServer *syslog.Server
	channel      *syslog.LogPartsChannel
}

func (e EventServer) GetEvent(stream pb.EventService_GetEventServer) error {
	log.Println("Client connected...")
	eventChannel := make(chan *pb.Event)
	syslogChannel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(syslogChannel)
	e.syslogServer.SetHandler(handler)

	go printLogInfo(syslogChannel, eventChannel)

	go recvStream(stream, handler, e)

	for {
		err := stream.Send(&pb.EventResponse{Event: <-eventChannel})
		if err != nil {
			log.Fatal("Can't send stream to client")
		}
	}
}

func NewEventServer(syslogServer *syslog.Server, channel *syslog.LogPartsChannel) *EventServer {
	return &EventServer{syslogServer: syslogServer, channel: channel}
}

func printLogInfo(channel syslog.LogPartsChannel, eventChannel chan *go_proto.Event) {
	for logParts := range channel {
		event := pb.Event{
			Message:  fmt.Sprintf("%v", logParts["message"]),
			Severity: fmt.Sprintf("%v", logParts["severity"]),
			Facility: fmt.Sprintf("%v", logParts["facility"]),
		}
		eventChannel <- &event
	}
}

func recvStream(stream go_proto.EventService_GetEventServer, handler *syslog.ChannelHandler, e EventServer) {
	_, err := stream.Recv()
	if err != nil {
		log.Println("Client disconnected...")
		handler.SetChannel(*e.channel)
		return
	}
}
