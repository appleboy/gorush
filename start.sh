echo $IOS_PROD_CERT | base64 -d > ios_prod.p12
echo $ANDROID_API_KEY | base64 -d > android_api_key.txt
./bin/gorush -c config/config.yml
