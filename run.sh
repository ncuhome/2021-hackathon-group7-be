#!/bin/bash
#pyf

swag init &&
./build.sh &&
sudo docker-compose build &&
sudo docker-compose up