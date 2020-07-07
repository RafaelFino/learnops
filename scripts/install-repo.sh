#!/bin/bash
cd

#update
sudo yum update -y
sudo yum install -y git vim zsh wget unzip jq

#install docker
#update
sudo yum update -y
sudo yum install -y git vim zsh wget unzip jq telnet

#get learn ops git repo and build
git clone https://github.com/RafaelFino/learnops.git
cd learnops
make

#create example data
bin/data-creator /home/ec2-user/learnops/etc/output.txt
bin/get-currency /home/ec2-user/learnops/etc/currency.json
cd

#create amd perm simple-server log
sudo touch /var/log/simple-server
sudo chmod +wr /var/log/simple-server

#install simple-server daemon
sudo cp learnops/daemon/simple-server.service /lib/systemd/system/ 
sudo cp learnops/daemon/db-api.service /lib/systemd/system/ 
sudo systemctl daemon-reload 
sudo systemctl enable simple-server
sudo systemctl enable db-api
sudo systemctl start simple-server 
sudo systemctl start db-api
sudo systemctl status simple-server
sudo systemctl status db-api

echo 'Server test:'
curl localhost

echo 'Done!'