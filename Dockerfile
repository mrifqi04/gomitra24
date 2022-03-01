#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
RUN mkdir /build
WORKDIR /build
COPY . .
RUN go get -d -v ./...
RUN cd /build && go build -o gomitra24

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /build/gomitra24 /gomitra24
ENTRYPOINT /gomitra24
LABEL Name=gomitra24 Version=0.0.1
EXPOSE 8001
