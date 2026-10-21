FROM --platform=$BUILDPLATFORM golang:1.27.2-alpine3.24@sha256:f92b6ef800e499660581efdabdf25d9d817a9d124eaf900924f0504e7e27e12d AS builder

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

# passed by czerwonk/go-ci/.github/workflows/docker-publish.yml; empty for local builds,
# which then keep the version compiled into main.go
ARG VERSION=""

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
COPY internal ./internal
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
  GOARM=${TARGETVARIANT#v} \
  go build -trimpath \
  -ldflags="-s -w ${VERSION:+-X main.version=$VERSION}" \
  -o /out/junos_exporter .


FROM gcr.io/distroless/static-debian13:nonroot@sha256:516757106de35205efb70b91ba03f6f1298424a0678c603255cf2852a7000087

# evaluated by junos_exporter itself when started without arguments (see docker_env.go)
ENV SSH_KEYFILE=""
ENV CONFIG_FILE="/config.yml"
ENV ALARM_FILTER=""
ENV CMD_FLAGS=""

WORKDIR /app
COPY --from=builder /out/junos_exporter /app/junos_exporter

EXPOSE 9326
CMD ["/app/junos_exporter"]
