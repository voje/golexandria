#!/bin/bash

docker rm -f ircd
docker run -d --net host --name ircd inspircd/inspircd-docker
