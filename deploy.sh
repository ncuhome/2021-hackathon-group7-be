#!/bin/bash
#pyf

swag init &&
./build.sh &&
rsync -az ./ root@nspyf.top:/root/project-be/tudo/ &&
ssh root@nspyf.top "cd /root/project-be/tudo && sudo docker-compose build && sudo docker-compose up -d"