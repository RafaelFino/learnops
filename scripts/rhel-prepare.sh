#!/bin/bash
cd

#update
sudo yum update -y
sudo yum install -y git vim zsh wget unzip jq

#install docker
sudo amazon-linux-extras install docker

#install oh-my-zsh
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" &

#install vim
git clone --depth=1 https://github.com/amix/vimrc.git ~/.vim_runtime
sh ~/.vim_runtime/install_awesome_vimrc.sh

#install zsh/vim fonts
git clone https://github.com/powerline/fonts.git 
./fonts/install.sh

#install exa (new LS)
wget https://github.com/ogham/exa/releases/download/v0.9.0/exa-linux-x86_64-0.9.0.zip
unzip exa-linux-x86_64-0.9.0.zip
sudo mv exa-linux-x86_64 /usr/local/bin/exa
rm exa-linux-x86_64-0.9.0.zip

#install go
wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
rm go1.14.4.linux-amd64.tar.gz
echo PATH=$PATH:/usr/local/go/bin >> ~/.zshrc
export PATH=$PATH:/usr/local/go/bin

sed -i 's/ZSH_THEME=\"robbyrussell\"/ZSH_THEME=\"agnoster\"/g' .zshrc
echo 'alias ls="exa -hHBmgaFl --git"' >> ~/.zshrc
echo set nu >> ~/.vim_runtime/my_configs.vim
sudo sed -i 's/ec2-user:\/bin\/bash/ec2-user:\/usr\/bin\/zsh/g' /etc/passwd

echo zsh >> ~/.bashrc

sudo service docker start
sudo usermod -a -G docker ec2-user

docker --version

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
sudo systemctl daemon-reload 
sudo systemctl enable simple-server
sudo systemctl start simple-server 
sudo systemctl status simple-server

echo 'Server test:'
curl localhost

echo 'Done!'

zsh