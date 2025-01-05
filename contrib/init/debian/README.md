# Run gorush in Debian/Ubuntu

## Installation

Put `gorush` binary into `/usr/bin` folder.

```sh
cp gorush /usr/bin/
chmod +x /usr/bin/gorush
```

put `gorush` init script into `/etc/rc.d`

```sh
cp contrib/init/debian/gorush /etc.rc.d/
```

install and remove System-V style init script links

```sh
update-rc.d gorush start 20 2 3 4 5 . stop 80 0 1 6 .
```

## Start service

create gorush configuration file.

```sh
mkdir -p /etc/gorush
cp config/testdata/config.yml /etc/gorush/
```

start gorush service.

```sh
/etc/init.d/gorush start
```
