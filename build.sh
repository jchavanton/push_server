#!/bin/bash
CURRENT_COMMIT=$1
CONTAINER="push_server"
VERSION="1.0.2"


docker build . -f Dockerfile -t ${CONTAINER}:${VERSION}
