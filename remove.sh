#!/bin/bash

DEPLOY_ACCOUNT="jaraxasoftware"
EXECUTABLE="gorush"
CONTAINER=js-gorush

# Stop a docker container
docker stop $CONTAINER > /dev/null 2>&1;

# Remove a docker container
docker rm $CONTAINER > /dev/null 2>&1;

# Remove docker images
docker rmi $DEPLOY_ACCOUNT/$EXECUTABLE:latest > /dev/null 2>&1;
docker rmi centurylink/ca-certs:latest > /dev/null 2>&1;