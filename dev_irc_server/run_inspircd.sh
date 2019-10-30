#!/bin/bash

docker rm -f ircd
docker run -d --name ircd -p 6667:6667 inspircd/inspircd-docker
