FROM golang

ENV SSH_KEYFILE "/ssh-keyfile"
ENV ROUTES_ENABLED true
ENV BGP_ENABLED true
ENV OSPF_ENABLED true
ENV ALARM_FILTER ""

RUN apt-get install -y git && \
    go get github.com/czerwonk/junos_exporter

CMD junos_exporter -ssh.targets $TARGETS -ssh.keyfile $SSH_KEYFILE -routes.enabled=$ROUTES_ENABLED -bgp.enabled=$BGP_ENABLED -ospf.enabled=$OSPF_ENABLED -alarms.filter=$ALARM_FILTER
EXPOSE 9326
