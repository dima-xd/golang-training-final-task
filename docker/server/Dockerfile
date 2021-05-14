FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app
COPY server /app
ENTRYPOINT ["./server"]
EXPOSE 1514/udp