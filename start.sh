echo $EEN_IOS_PROD_CERT | base64 -d > ios_prod.p12
./bin/gorush -c config/config.yml -k $EEN_ANDROID_API_KEY
