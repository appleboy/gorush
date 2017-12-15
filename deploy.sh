#! bin/bash

# the file for config.core.cert_path
if [ ! -z "$CORE__CERT_PATHFILE" ]; then echo $CORE__CERT_PATHFILE | base64 -d > cert.pem; fi
# the file for config.core.key_path
if [ ! -z "$CORE__KEY_PATHFILE" ]; then echo $CORE__KEY_PATHFILE | base64 -d > key.pem; fi
# ios key file
if [ ! -z "$IOS__KEY_FILE" ]; then echo $IOS__KEY_FILE | base64 -d > ios_key.pem; fi
if [ ! -z "$CONFIG_YML" ]; then echo $CONFIG_YML | base64 -d > config.yml; fi

/bin/gorush -c config.yml
