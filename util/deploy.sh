#!/usr/bin/env bash
set -e

gcloud endpoints services deploy api_descriptor.pb api_config.yaml
kubectl apply -f deploy.yaml