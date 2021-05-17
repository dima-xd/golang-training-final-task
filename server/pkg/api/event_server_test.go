package api

import (
	"log"
	"os/exec"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestLoggerCommandServer(t *testing.T) {
	assert := assert.New(t)
	cmdLogger := exec.Command("/bin/sh", "-c", "logger --tcp -n 0.0.0.0 --port 1514 --rfc5424 \"Test\"")
	err := cmdLogger.Run()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("/bin/sh", "-c", "docker logs -n 1 syslog-server")
	actual, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	expected := []byte("level=info Facility=1 Message=Test Severity=5\n")
	assert.Equal(expected, actual[28:])
}
