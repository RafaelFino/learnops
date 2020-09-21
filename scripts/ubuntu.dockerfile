FROM ubuntu:latest

WORKDIR /root

#update system
RUN apt update && \
    apt install -y curl git zsh vim wget unzip jq telnet 

#install vim
RUN git clone --depth=1 https://github.com/amix/vimrc.git ~/.vim_runtime && \
    sh ~/.vim_runtime/install_awesome_vimrc.sh && \
    echo set nu >> ~/.vim_runtime/my_configs.vim

#install fonts
RUN git clone https://github.com/powerline/fonts.git && \
    ./fonts/install.sh

#install oh-my-zsh
RUN sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" && \
    sed -i 's/ZSH_THEME=\"robbyrussell\"/ZSH_THEME=\"avit\"/g' .zshrc

#install exa
RUN wget https://github.com/ogham/exa/releases/download/v0.9.0/exa-linux-x86_64-0.9.0.zip && \
    unzip exa-linux-x86_64-0.9.0.zip && \
    mv exa-linux-x86_64 /usr/local/bin/exa && \
    rm exa-linux-x86_64-0.9.0.zip && \
    echo 'alias ls="exa -hHBmgaFl --git"' >> ~/.zshrc

#setup timezone
ENV TZ=America/Sao_Paulo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone && \
    apt install -y tzdata
