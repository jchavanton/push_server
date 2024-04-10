#!/bin/bash
DIR_PREFIX=`pwd`
CONTAINER=push_server
VERSION="1.0.2"
IMAGE=${CONTAINER}:${VERSION}
docker stop ${CONTAINER}
docker rm ${CONTAINER}
docker run -d --net=host \
              --name=${CONTAINER} \
              -v ${DIR_PREFIX}/cert:/go/cert \
              --env-file ${CONTAINER}.env \
              ${IMAGE} #\
#              tail -f /dev/null
