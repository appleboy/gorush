#!/bin/bash
echo $EEN_IOS_PROD_CERT | base64 -d > ios_prod.p12

if [ -z $IOS_PRODUCTION_MODE ];
then
 	./bin/gorush -c config/config.yml -k $EEN_ANDROID_API_KEY --production
else
	./bin/gorush -c config/config.yml -k $EEN_ANDROID_API_KEY
fi
