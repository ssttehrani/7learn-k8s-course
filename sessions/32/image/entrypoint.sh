#!/bin/sh

trap "exit" SIGINT

mkdir -p /var/contents

while :
do
    echo "Index file is last updated at $(date)" > /var/contents/index.html
    echo "Sleeping for ${INTERVAL} seconds..."
    sleep ${INTERVAL}
done
