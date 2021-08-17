#!/bin/bash
docker build -t session_server .
docker run -it -p 4101:4101 -p 4102:4102 -d session_server
