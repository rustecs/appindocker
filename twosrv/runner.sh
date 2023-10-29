#!/bin/bash

docker network create --driver bridge twosrv-net

#run two server at brige
docker run -dit --name twosrv-memcache --network twosrv-net memcached
docker run -dit --name tsrw --network twosrv-net -p 8080:8080 twosrvapp 
