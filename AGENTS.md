# Repository Guidelines

## Project Structure & Module Organization

`go.mod` declares module `go-novel` and Go 1.19.

- `api.go`, `admin.go`, and `bqgapi.go` are separate server entry points; other root files include collectors and maintenance tools.
- `routers/` registers Gin endpoints; `app/controller/` handles requests; `app/service/` contains business logic; `app/models/` defines GORM models and request/response types.
- `middleware/`, `db/`, `config/`, and `global/` provide shared infrastructure; `pkg/` and `utils/` hold helpers.
- `public/` holds static assets, Casbin policies, and deployment scripts.
- `gofound-1.4.1/` is a separate search-engine module with its own tests, documentation, and Vue admin assets.

## Build, Test, and Development Commands

From the repository root:

- `go mod download`: download dependencies.
- `go build -o novel-api api.go`: build the API executable.
- `go build -o novel-admin admin.go`: build the administration server.
- `go run api.go` or `go run admin.go`: start the selected server locally.
- `go test ./app/... -cover`: test application packages and report coverage.
- `(cd gofound-1.4.1 && go test ./searcher/pagination)`: run an isolated search-engine test package.

Select individual root entry files: multiple `main` functions prevent treating the root as one buildable package, including with `go test ./...`. Deployment scripts contain production paths and process-management commands; inspect them before use.

## Coding Style & Naming Conventions

Format changed Go files with `gofmt -w <file.go>`; use tab indentation. Follow nearby naming patterns such as `book_route.go`, `book_service`, and `mc_book.go`. Keep request handling in controllers and business logic in services. Preserve existing JSON/form tags and database mappings. No repository-wide lint configuration was found.

## Testing Guidelines

Use Go's `testing` package, `*_test.go` files, and `TestXxx` functions. Existing tests reside in the search-engine module; no coverage threshold is configured. Add focused regression tests alongside changed packages. Use temporary directories for storage tests; some existing search tests reference developer-specific absolute paths. Tests importing `config` require `config.yml` in their working directory.

## Commit & Pull Request Guidelines

History contains `first commit` and numeric subjects; no descriptive convention is established. Use concise imperative subjects, for example `Fix chapter cache invalidation`. PRs should explain the change, link relevant issues, list validation results and blockers, and identify configuration or schema changes. Include screenshots for admin UI changes.

## Security & Configuration

Create an ignored local `config.yml` using `config.yml.dev` as a reference. Verify environment-dependent database and Redis settings before starting; workers may also require NSQ. Keep credentials, logs, queue data, database dumps, and generated binaries out of new commits.
