FROM golang:1.27rc2-alpine3.24@sha256:dcbb18cc5fa1082364dc6aa95224b6b55429d09cbb9631a053d8064c1c367300 AS builder
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
