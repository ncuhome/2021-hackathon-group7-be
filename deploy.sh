#!/bin/bash
#pyf

./build.sh &&
rsync -az ./ root@nspyf.top:/root/nspyf-be/blog/ &&
ssh root@nspyf.top "cd /root/nspyf-be/blog && sudo docker-compose build && sudo docker-compose up -d"