FROM golang as builder
ADD . /go/junos_exporter/
WORKDIR /go/junos_exporter
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/junos_exporter

FROM alpine
ENV SSH_KEYFILE ""
ENV CONFIG_FILE "/config.yml"
ENV ALARM_FILTER ""
ENV CMD_FLAGS ""
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/bin/junos_exporter .
CMD ./junos_exporter -ssh.keyfile=$SSH_KEYFILE -config.file=$CONFIG_FILE -alarms.filter=$ALARM_FILTER $CMD_FLAGS
EXPOSE 9326
