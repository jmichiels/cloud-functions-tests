#!/usr/bin/env sh

base_url="https://us-central1-cloud-functions-tests-45711.cloudfunctions.net/ServeHTTP/api/v1"

curl -v -d '{"firstName"="Jacques","lastName"="Michiels","age"="29"}' -H "Content-Type: application/json" "${base_url}/clients/"
