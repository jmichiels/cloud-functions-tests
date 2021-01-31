# cloud-functions-tests

A simple example using the [Gin web framework](https://github.com/gin-gonic/gin) in
a [Google Cloud Function](https://cloud.google.com/functions) to host an HTTP service with multiple routes.

`cloud_functions.go` contains the definition of the cloud function (it apparently needs to be at the root of the
project, in the same directory as the `go.mod` file). It can be deployed to Google Cloud using:

```bash
./scripts/deploy.sh
```

`internal/http_handler` contains the Gin framework specific code, with the routes and handlers definitions.

`cmd` contains some executables (following the typical Go project layout
described [here](https://github.com/golang-standards/project-layout)) using the deployed service. Try it out using:

```
go run cmd/ping/main.go
```
```
go run cmd/hello/main.go "<your-name>"
```

## References

- Google Cloud documentation to create a Cloud Function using the Go
  runtime: https://cloud.google.com/functions/docs/first-go

