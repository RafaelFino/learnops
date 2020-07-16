#!/bin/bash
cd
sudo apt update -y && sudo apt install -y git vim zsh wget unzip jq telnet
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" & 
git clone --depth=1 https://github.com/amix/vimrc.git ~/.vim_runtime
./fonts/install.sh
wget https://github.com/ogham/exa/releases/download/v0.9.0/exa-linux-x86_64-0.9.0.zip
unzip exa-linux-x86_64-0.9.0.zip
sudo mv exa-linux-x86_64 /usr/local/bin/exa 
rm exa-linux-x86_64-0.9.0.zip 
sed -i 's/ZSH_THEME=\"robbyrussell\"/ZSH_THEME=\"avit\"/g' .zshrc 
echo 'alias ls="exa -hHBmgaFl --git"' >> ~/.zshrc 
echo set nu >> ~/.vim_runtime/my_configs.vim 
echo zsh >> ~/.bashrc  
git clone https://github.com/RafaelFino/learnops.git
cd learnops
make
zsh
echo 'Done!'