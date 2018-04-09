FROM golang

ENV SSH_KEYFILE "/ssh-keyfile"
ENV CONFIG_FILE="/config.yml"
ENV ALARM_FILTER ""

RUN apt-get install -y git && \
    go get github.com/czerwonk/junos_exporter

CMD junos_exporter -ssh.keyfile=$SSH_KEYFILE -config.file=$CONFIG_FILE -alarms.filter=$ALARM_FILTER
EXPOSE 9326
