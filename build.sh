#/usr/bin/env bash

APPS=$(ls exec)
for APP in $APPS; do
    go build exec/$APP/*.go
done


