#!/usr/bin/env bash

TAG=`cat VERSION`

docker build -t mylxsw/wizard-enhance .

docker tag mylxsw/wizard-enhance mylxsw/wizard-enhance:$TAG
docker tag mylxsw/wizard-enhance:$TAG mylxsw/wizard-enhance:latest
docker push mylxsw/wizard-enhance:$TAG
docker push mylxsw/wizard-enhance:latest

