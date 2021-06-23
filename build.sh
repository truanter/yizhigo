#!/usr/bin/env sh

TAG=`date +"%Y%m%d_%H%M%S"`
docker build -f dockers/yizhigo.docker -t zhcherish/yizhigo:"$TAG" .