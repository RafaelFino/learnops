FROM ubuntu:latest

RUN apt update
RUN apt install -y curl
RUN curl https://raw.githubusercontent.com/RafaelFino/learnops/master/scripts/debian-root-prepare.sh | bash

WORKDIR ~

CMD [ "zsh" ]

