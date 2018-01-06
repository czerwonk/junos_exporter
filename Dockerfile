FROM golang

RUN apt-get install -y git && \
    go get github.com/czerwonk/junos_exporter

CMD junos_exporter -ssh.targets $TARGETS -ssh.keyfile /ssh-keyfile
EXPOSE 9326
