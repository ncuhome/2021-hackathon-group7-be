#!/bin/bash 
#pyf

cp /root/secret/.env ./.env
docker-compose build
docker-compose up -d