#!/bin/sh

systemctl daemon-reload
systemctl enable prometheus-junos-exporter
systemctl start prometheus-junos-exporter