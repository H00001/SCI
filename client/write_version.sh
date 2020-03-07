#!/usr/bin/env bash
LIB="client"
FILE="version.json"

go build
printf '{"version":"%s","md5":"%s","sha1":"%s","sha256":"%s"}' "$(cat version)" "$(md5sum ${LIB}|cut -d" " -f1)" "$(sha1sum ${LIB}|cut -d" " -f1)" "$(sha256sum ${LIB}|cut -d" " -f1)" > $FILE
mv client SCI_client_linux
scp SCI_client_linux root@123.56.223.150:/var/www/physics/static
scp version.json root@123.56.223.150:/var/www/physics/static