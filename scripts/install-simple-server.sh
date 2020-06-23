#!/bin/bash
cp dademon/simple-server.service /lib/systemd/systemd/simple-server.service
systemctl daemon-reload
systemctl enable
systemctl start simple-service