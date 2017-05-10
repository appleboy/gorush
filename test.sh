#!/bin/bash

PORT=10421
VERSION="/push/v1"

curl \
	-XGET \
	-H "Accept: application/json" \
 	"localhost:$PORT$VERSION/stats/test" | python -mjson.tool