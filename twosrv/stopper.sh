#!/bin/bash

# docker container stop alpine1 alpine2 alpine3 alpine4
docker container stop tsrw twosrv-memcache

# docker container rm alpine1 alpine2 alpine3 alpine4
docker container rm tsrw twosrv-memcache

docker network rm twosrv-net
