#!/bin/bash

#PLATFORM_IP="$( ipconfig getifaddr en0 )"
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEPLOY_ACCOUNT="jaraxasoftware"
EXECUTABLE="gorush"
CONTAINER=js-gorush
PORT=10421

# Remove previous container
docker rm $CONTAINER > /dev/null 2>&1

docker run -ti -d --name $CONTAINER --restart always \
	-p $PORT:8088 \
	-v $DIR/config:/config:ro \
	$DEPLOY_ACCOUNT/$EXECUTABLE:latest /gorush -c /config/config.yml