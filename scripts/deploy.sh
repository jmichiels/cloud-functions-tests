#!/usr/bin/env sh

gcloud functions deploy ServeHTTP --runtime go113 --trigger-http --allow-unauthenticated