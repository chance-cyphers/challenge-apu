set -ex

git pull

docker build -t skyfrog28/challenge:latest .
docker push skyfrog28/challenge:latest