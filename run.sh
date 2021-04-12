#!/bin/bash 
#pyf

cp /root/secret/hackathon.env ./.env
docker-compose build
docker-compose up -d