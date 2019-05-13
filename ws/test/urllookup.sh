#!/bin/bash
if [ -z $1 ]
then
  echo "usage : " $0 "start|stop"
  exit
fi

if [ $1 = "start" ]
then
    docker-compose -f docker/urllookup.yml up -d
fi
if [ $1 = "stop" ]
then
   docker-compose -f docker/urllookup.yml down
fi

