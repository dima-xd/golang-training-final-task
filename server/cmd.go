package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/mcuadros/go-syslog.v2"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	err := server.ListenUDP("0.0.0.0:1514")
	if err != nil {
		log.Fatal(err)
	}
	err = server.Boot()
	if err != nil {
		log.Fatal(err)
	}

	go printLogInfo(channel)

	server.Wait()
}

func printLogInfo(channel syslog.LogPartsChannel) {
	for logParts := range channel {
		log.WithFields(log.Fields{
			"Facility": logParts["facility"],
			"Severity": logParts["severity"],
		}).Info(logParts["message"])
	}
}
