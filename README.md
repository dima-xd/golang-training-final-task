# golang-training-final-task

Go application which reads syslog and print it to terminal.

# How to use

#### via golang:
1. run application
2. use logger in terminal, e.g:
`logger --udp -n 0.0.0.0 --port 1514 --rfc5424 "Test"`

### via docker:
While you are in project root:
1. build go application:
`CGO_ENABLED=0 go build -o ./docker/server ./server`
2. build Dockerfile:
`docker build -t server ./docker/server`
3. run docker image:
`docker run -p 1514:1514/udp server`
4. use logger in terminal, e.g:
`logger --udp -n 0.0.0.0 --port 1514 --rfc5424 "Test"`
