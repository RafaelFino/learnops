!#/bin/bash
cd

wget https://github.com/nats-io/nats-streaming-server/releases/download/v0.18.0/nats-streaming-server-v0.18.0-linux-amd64.zip nats-streaming
unzip nats-streaming-server-v0.18.0-linux-amd64.zip
rm nats-streaming-server-v0.18.0-linux-amd64.zip

sudo mkdir /opt/nats
sudo mv nats-streaming-server-v0.18.0-linux-amd64/nats-streaming-server /opt/nats/nats-streaming-server

sudo mkdir /var/log/nats
sudo touch /var/log/nats/stream.log
sudo chmod 777 /var/log/nats -R

echo '
[Unit]
Description=Nats Streaming Server

[Service]
Type=simple 
Restart=always 
RestartSec=1 
StartLimitInterval=0 
RemainAfterExit=yes
ExecStart=/opt/nats/nats-streaming-server -l /var/log/nats/stream.log -m 8000
WorkingDirectory=/opt/nats

[Install]
WantedBy=multi-user.target
' > ~/nats-streaming.service

sudo systemctl daemon-reload 
sudo systemctl enable nats-streaming
sudo systemctl start nats-streaming
sudo systemctl status nats-streaming