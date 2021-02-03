#!/usr/bin/env sh

base_url="https://us-central1-cloud-functions-tests-45711.cloudfunctions.net/ServeHTTP"

curl "${base_url}/ping"
