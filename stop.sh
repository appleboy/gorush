#!/bin/bash

CONTAINER=js-gorush

# Stop a docker container
docker stop $CONTAINER > /dev/null 2>&1;