#!/bin/bash

sudo docker stop string-service
sudo docker rm string-service

sudo docker build -t string-service-image .
sudo docker run --name string-service-server -d -p "8080:8080" string-service-image
