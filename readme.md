# cloud-functions-tests

## Usage

Install the `demo-server` and `demo-client`:

```
go install github.com/jmichiels/cloud-functions-tests/cmd/demo-server
```

```
go install github.com/jmichiels/cloud-functions-tests/cmd/demo-client
```

In one terminal, run the server:

```
demo-server
```

In another terminal, use the client:

```
demo-client --help
```

## Structure:

The structure follows those recommendations: https://github.com/golang-standards/project-layout

- `api/protos` contains the `.protos` (Protobuf definitions) files used to generate the Grpc clients and server stubs.

- `cmd` contains the code for the executables.

- `internal` contains code only accessible from within this project. Structured in logical 4 parts:
    - **Domain** (`domain` package): implements the business logic,
    - **Repository** (`repository.go`): persists and retrieves domain entities,
    - **Service** (`service.go`):
        - Uses the repository to fetch the required entities,
        - Uses the domain to manipulate those entities,
        - Persists the updated entities when done.
    - **API**:
        - Converts the requests/responses from the client from/to valid domain objects,
        - Use the service with those domain objects.
    
- `pkg` contains code that can be imported from outside this project (at the moment, only the generated protobuf and grpc code)

