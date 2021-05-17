# golang-training-final-task

Go application which creates grpc server and syslog server. If client is connected then client writes logs to client terminal. If client is not connected then server writes logs to server terminal. Grpc-server uses `8081` port. Syslog-server uses `1514` port.

## How to use

1. Up server and client via docker-compose while you're in project root:
`docker-compose up`
2. To send logs via logger:
`logger --tcp -n 0.0.0.0 --port 1514 --rfc5424 "<message here>"`

## How to run test

1. To test if client writes logs. While you're in project root:
`go test server/pkg/api/event_client_test.go`
2. To test if server writes logs. While you're in project root and client is not connected (to disconnect client `docker stop syslog-client` and to connect `docker start syslog-client`):
`go test server/pkg/api/event_server_test.go`
