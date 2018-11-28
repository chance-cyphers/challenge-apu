#!/usr/bin/env bash
set -ex

git pull

docker build -t skyfrog28/challenge:latest .
docker push skyfrog28/challenge:latest

kubectl set image deployment challenge-api challenge-api=skyfrog28/challenge:latest