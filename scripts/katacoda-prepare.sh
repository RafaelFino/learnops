#!/bin/bash
cd

#update
apt update -y
apt install -y git vim zsh wget unzip jq telnet curl terminator

#install vim
git clone --depth=1 https://github.com/amix/vimrc.git ~/.vim_runtime
sh ~/.vim_runtime/install_awesome_vimrc.sh

#install zsh/vim fonts
git clone https://github.com/powerline/fonts.git 
./fonts/install.sh

#install exa (new LS)
wget https://github.com/ogham/exa/releases/download/v0.9.0/exa-linux-x86_64-0.9.0.zip
unzip exa-linux-x86_64-0.9.0.zip
mv exa-linux-x86_64 /usr/local/bin/exa
rm exa-linux-x86_64-0.9.0.zip

#install oh-my-zsh
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

sed -i 's/ZSH_THEME=\"robbyrussell\"/ZSH_THEME=\"avit\"/g' .zshrc
echo 'alias ls="exa -hHBmgaFl --git"' >> ~/.zshrc
echo set nu >> ~/.vim_runtime/my_configs.vim
sed -i 's/ec2-user:\/bin\/bash/ec2-user:\/usr\/bin\/zsh/g' /etc/passwd
echo zsh >> ~/.bashrc

#install go
wget https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz
rm go1.14.4.linux-amd64.tar.gz
echo PATH=$PATH:/usr/local/go/bin >> ~/.zshrc
export PATH=$PATH:/usr/local/go/bin

go version

echo 'Done!'

zsh 
