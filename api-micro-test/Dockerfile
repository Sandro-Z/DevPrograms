FROM docker.io/library/alpine:latest

RUN apk --no-cache add ca-certificates

COPY ./build/server_linux_amd64 /bin/server

EXPOSE 8080

ENTRYPOINT ["/bin/server"]
