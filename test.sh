#!/bin/bash

PORT=8088

curl \
	-XGET \
	-H "Accept: application/json" \
 	"localhost:$PORT/api/stat/go" | python -mjson.tool