#!/bin/bash

if [ $EUID -ne 0 ]; then
  echo "Run as root."
fi

if ! grep -q "irc.demo.server" /etc/hosts; then
  echo "127.0.0.1 irc.demo.server" >> /etc/hosts
fi

docker rm -f ircd
docker run -d --net host --name ircd inspircd/inspircd-docker
