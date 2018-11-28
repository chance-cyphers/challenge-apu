#!/usr/bin/env bash
set -x

# Be up to date
git pull

# Bump version
version=`cat VERSION`
major="$(cut -d'.' -f1 <<< $version)"
minor="$(cut -d'.' -f2 <<< $version)"
newVersion="$major.$((minor+1))"
echo ${newVersion} > VERSION

# Commit version bump
git add -A
git commit -m "Version $version"
git push

# Build/store image
docker build -t skyfrog28/challenge:${newVersion} .
docker push skyfrog28/challenge:${newVersion}

# Put it up there
kubectl set image deployment challenge-api challenge-api=skyfrog28/challenge:${newVersion}