# Build stage backend
FROM golang:1.12-alpine AS build-backend
WORKDIR /build/
COPY . .
RUN apk add --no-cache git ca-certificates
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
COPY --from=0 /build/goban .
COPY --from=1 /build/dist/ ./ui/
COPY config.yml .
ENTRYPOINT [ "./goban", "server" ]
