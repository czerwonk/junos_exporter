FROM golang

RUN apt-get install -y git && \
    go get github.com/czerwonk/junos_exporter

CMD junos_exporter -snmp.targets $TARGETS -snmp.community $COMMUNITY
EXPOSE 9326
