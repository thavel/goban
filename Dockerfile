# Build stage backend
FROM golang:1.10.3-alpine AS build
WORKDIR /go/src/github.com/thavel/goban/
COPY . .
RUN apk add --no-cache dep git ca-certificates
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o goban .

# Exec stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/thavel/goban/goban .
COPY config.yml .
COPY rbac.conf .
ENTRYPOINT [ "./goban", "server" ]
