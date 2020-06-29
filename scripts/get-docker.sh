#!/bin/bash
cd

#update
sudo yum update -y
sudo yum install -y git vim zsh wget unzip jq

#install docker
sudo amazon-linux-extras install docker

sudo service docker start
sudo usermod -a -G docker ec2-user

docker --version

echo 'Done!'

exit