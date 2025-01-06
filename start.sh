#!/bin/bash
echo $EEN_IOS_PROD_CERT | base64 -d > ios_prod.p12
echo $EEN_ANDROID_PROD_KEY | base64 -d > android_prod.p12

if [ -z $IOS_PRODUCTION_MODE ];
then
 	./bin/gorush -c config/config.yml --production
else
	./bin/gorush -c config/config.yml
fi
