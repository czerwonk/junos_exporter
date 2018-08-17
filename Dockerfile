FROM golang:alpine as builder
RUN go get github.com/czerwonk/junos_exporter


FROM alpine:latest

ENV SSH_KEYFILE "/ssh-keyfile"
ENV CONFIG_FILE "/config.yml"
ENV ALARM_FILTER ""

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/bin/junos_exporter .

CMD junos_exporter -ssh.keyfile=$SSH_KEYFILE -config.file=$CONFIG_FILE -alarms.filter=$ALARM_FILTER

EXPOSE 9326
