#!/usr/bin/env sh

base_url="https://us-central1-cloud-functions-tests-45711.cloudfunctions.net/ServeHTTP"

curl -d '{"name":"YOUR NAME"}' -H "Content-Type: application/json" "${base_url}/hello"
