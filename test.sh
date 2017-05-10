#!/bin/bash

PORT=8088

curl \
	-XGET \
	-H "Accept: application/json" \
 	"localhost:$PORT/stat/go" | python -mjson.tool