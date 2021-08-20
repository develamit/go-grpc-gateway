#!/bin/bash
USER=`whoami`
IMAGE="session_server:v0.1"

mkdir -p /home/${USER}/logs
mkdir -p /home/${USER}/state

# build the image
docker build -t ${IMAGE} .

# run as a detached container
# docker run -it -p 4101:4101 -p 4102:4102 -d session_server
docker run -it -p 4101:4101 -p 4102:4102 -v /home/${USER}/logs:/app/logs -v /home/${USER}/state:/app/state -d ${IMAGE}
