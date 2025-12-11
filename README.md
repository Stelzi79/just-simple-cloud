# Just Simple Cloud

Just Simple Cloud (jsc) is a small, opinionated Go-based CLI and library to help test, inspect and deploy simple cloud environments on VPS or self-hosted hardware.

Key goals
- Minimal, readable command-line tooling for common cloud tasks
- Export useful debug/status information and small deployment helpers
- Keep integrations lightweight and easy to adapt

Repository layout
- `jsc/` — main Go module and entrypoint (`main.go`)
- `cmd/` — CLI command definitions and subcommands (e.g. `status`)
- `Makefile` — convenient targets for build/run/test/format
- `openapi: '3.0.yml` — API specifications (where applicable)
- `enums/`, `globals/`, `libs/`, `types/` — internal packages and helpers

Quickstart

Prerequisites
- Go 1.20+ installed
- `make` available in your PATH

Build (use Make)

The repository includes a `Makefile` with common targets. Build the project using:

```bash
make build
```

The default build artifact is written to `.out/bin/jsc` (see `Makefile`).

Run (use Make)

Show CLI help:

```bash
make run ARGS="--help"
```

Run the `status` subcommand:

```bash
make run ARGS="status"
```

Other useful targets
- `make force-build` — clean + build
- `make install` — `go install` the module
- `make test` — run `go test ./...`
- `make format` — run `go fmt ./...`
- `make tidy` — run `go mod tidy` inside `jsc/`
- `make help-message` — print help about available Make targets

Notes for developers
- The CLI entrypoint is `jsc/main.go` and commands are wired under `cmd/`.
- You can still run during development with `go run ./jsc`, but prefer `make` targets for reproducible builds and CI.

Contributing
- Fork the repo, create a feature branch and open a pull request. Include tests for new behavior when possible.

License
- See the `LICENSE` file at the repository root for license terms.

