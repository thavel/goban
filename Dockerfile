# Build stage backend
FROM golang:1.10.3-alpine AS build-backend
WORKDIR /go/src/github.com/thavel/goban/
COPY . .
RUN apk add --no-cache dep git ca-certificates
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o goban .

# Build stage frontend
FROM node:10.11.0-alpine AS build-frontend
WORKDIR /build/
COPY ui/ .
RUN apk add --no-cache yarn
RUN yarn install
RUN yarn build

# Exec stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/thavel/goban/goban .
COPY --from=1 /build/dist/ ./ui/
COPY config.yml .
COPY rbac.conf .
ENTRYPOINT [ "./goban", "server" ]
