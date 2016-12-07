#!/bin/sh
#

set -e

sed -i"" -e "s/localhost/redis/g" config/config.yml
sed -i"" -e "s/localhost/redis/g" config/config.go
sed -i"" -e "s/localhost/redis/g" config/config_test.go
sed -i"" -e "s/localhost/redis/g" gorush/status_test.go
sed -i"" -e "s/localhost/redis/g" storage/redis/redis_test.go

echo "install package and testing code coverage."
make dep_install && coverage all
