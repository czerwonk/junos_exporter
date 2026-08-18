FROM golang:1.26.6-alpine3.24@sha256:3889b425f035be855a72fb4755265311293b6d414521f0a519d819df32222d83 AS builder
ADD . /go/junos_exporter/
WORKDIR /go/junos_exporter
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/junos_exporter

FROM alpine:3.24.1@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b
ENV SSH_KEYFILE=""
ENV CONFIG_FILE="/config.yml"
ENV ALARM_FILTER=""
ENV CMD_FLAGS=""
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/bin/junos_exporter .
CMD ./junos_exporter -ssh.keyfile=$SSH_KEYFILE -config.file=$CONFIG_FILE -alarms.filter=$ALARM_FILTER $CMD_FLAGS
EXPOSE 9326
