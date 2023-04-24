# syntax=docker/dockerfile:1

FROM golang:1.20-alpine as builder
# HINT: Needed by 'hub' during compile time.
RUN apk add git
RUN go install github.com/kelseyhightower/hub-credential-helper@0.0.1
RUN go install github.com/github/hub@v2.11.2

FROM alpine:latest
RUN apk --no-cache add git
COPY --from=builder /go/bin/hub-credential-helper /usr/local/bin/hub-credential-helper
COPY --from=builder /go/bin/hub /usr/local/bin/hub
# HINT: needed by 'hub'.
RUN mkdir /root/.config
ENTRYPOINT ["/usr/local/bin/hub"]
