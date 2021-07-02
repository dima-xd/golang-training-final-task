package main

import (
	"net"
	"os"

	pb "github.com/dimaxdqwerty/golang-training-final-task/proto/go_proto"
	"github.com/dimaxdqwerty/golang-training-final-task/server/pkg/api"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/mcuadros/go-syslog.v2"
)

var (
	listen = os.Getenv("listen")
)

func init() {
	if listen == "" {
		listen = "127.0.0.1:1514"
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	syslogServer := syslog.NewServer()
	syslogServer.SetFormat(syslog.RFC5424)
	syslogServer.SetHandler(handler)
	err = syslogServer.ListenTCP(listen)
	if err != nil {
		log.Fatal(err)
	}
	err = syslogServer.Boot()
	if err != nil {
		log.Fatal(err)
	}

	go printLogInfo(channel)

	pb.RegisterEventServiceServer(server, api.NewEventServer(syslogServer, &channel))

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}

	syslogServer.Wait()
}

func printLogInfo(channel syslog.LogPartsChannel) {
	for logParts := range channel {
		log.WithFields(log.Fields{
			"Facility": logParts["facility"],
			"Severity": logParts["severity"],
		}).Info(logParts["message"])
	}
}
