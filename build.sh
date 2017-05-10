#!/bin/bash

IMAGE_FILE="jaraxasoftware_gorush"

# Load a docker image from a .tar.gz file
gunzip < $IMAGE_FILE.tar.gz | docker load
