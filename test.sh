#!/bin/bash

PORT=8088

curl \
	-XGET \
	-H "Accept: application/json" \
 	"localhost:$PORT/stats/test" | python -mjson.tool