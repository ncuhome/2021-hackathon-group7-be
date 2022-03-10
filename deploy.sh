#!/bin/bash
#pyf

swag init &&
./build.sh &&
rsync -az ./ root@101.132.170.112:/root/project-be/tudo/ &&
ssh root@101.132.170.112 "cd /root/project-be/tudo && sudo docker-compose build && sudo docker-compose up -d"