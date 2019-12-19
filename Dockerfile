FROM golang:1.13.5 as builder
ENV DATA_DIRECTORY /go/src/github.com/suchaimi/finance
WORKDIR $DATA_DIRECTORY
ARG APP_VERSION
ARG CGO_ENABLED=0
COPY . .
RUN go build -ldflags="-X github.com/suchaimi/finance/internal/config.Version=$APP_VERSION" github.com/suchaimi/finance/cmd/server

FROM alpine:3.10
ENV DATA_DIRECTORY /go/src/github.com/suchaimi/finance
RUN apk add --update --no-cache \
		ca-certificates
COPY --from=builder ${DATA_DIRECTORY}/server /finance
ENTRYPOINT ["/finance"]