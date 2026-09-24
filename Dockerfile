FROM golang:1.27.1-alpine3.24@sha256:cf6fca6641884b8433441b2b0652976f975e1d0fdd26d177eaaf8596087f3125 AS builder
ADD . /go/junos_exporter/
WORKDIR /go/junos_exporter
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/junos_exporter

FROM gcr.io/distroless/static-debian13:nonroot@sha256:2293b36c7c9082bf4115aab724b4d2cddec82c8eba39bf27ac0517e159acf150
# evaluated by junos_exporter itself when started without arguments (see docker_env.go)
ENV SSH_KEYFILE=""
ENV CONFIG_FILE="/config.yml"
ENV ALARM_FILTER=""
ENV CMD_FLAGS=""
WORKDIR /app
COPY --from=builder /go/bin/junos_exporter .
CMD ["/app/junos_exporter"]
EXPOSE 9326
