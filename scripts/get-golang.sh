#!/bin/bash
cd

#update
sudo yum update -y
sudo yum install -y git vim zsh wget unzip jq
sudo yum groupinstall -y "Development Tools"

#install go
wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
rm go1.14.4.linux-amd64.tar.gz
echo PATH=$PATH:/usr/local/go/bin >> ~/.zshrc
export PATH=$PATH:/usr/local/go/bin

go version

echo 'Done!'