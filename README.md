# REST API Implementation

Go server and RESTful API template. Uses only Go's standard library plus `github.com/pkg/errors`.

## Run

```bash
make run       # go run main.go
make build     # compiles to ./main
```

Server listens on `http://localhost:3000`.

## Endpoints

| Method | Path | Response |
|--------|------|----------|
| GET | `/` | `200 OK` (empty body) |
| GET | `/record` | `403 {"message": "forbidden"}` |

## Structure

```
api/        # route registration and handlers
routing/    # custom router (exact URL matching, method enforcement)
server/     # alternate ListenAndServe wrapper (unused by main.go)
utils/
  errors/   # JSON error types (400, 403, 404)
  logger/   # leveled logging (INFO/WARN/ERROR)
```

## Adding Routes

Register in `api/routes.go` using a handler that returns an error:

```go
apiRouter.RegisterRoute("GET", "/path", func(w http.ResponseWriter, r *http.Request) error {
    return nil // or return errors.NewNotFoundError(...)
})
```

Returned errors are automatically serialized to JSON with the correct HTTP status code.
