# Just Simple Cloud

Just Simple Cloud (jsc) is a small, opinionated Go-based CLI and library to help test, inspect and deploy simple cloud environments on VPS or self-hosted hardware.

Key goals
- Minimal, readable command-line tooling for common cloud tasks
- Export useful debug/status information and small deployment helpers
- Keep integrations lightweight and easy to adapt

Repository layout
- [jsc](jsc) — main Go module and entrypoint ([jsc/main.go](jsc/main.go))
- [jsc/cmd](jsc/cmd) — CLI command definitions and subcommands (e.g. `status`)
- [Makefile](Makefile) — convenient targets for build/run/test/format
- [openapi: '3.0.yml](openapi:%20'3.0.yml) — API specifications (where applicable)
- [jsc/enums](jsc/enums), [jsc/globals](jsc/globals), [jsc/libs](jsc/libs), [jsc/types](jsc/types) — internal packages and helpers

Quickstart

Prerequisites
- Go 1.25.3 installed
- `make` available in your PATH

Build (use Make)

The repository includes a [Makefile](Makefile) with common targets. Build the project using:

```bash
make build
```

The default build artifact is written to `.out/bin/jsc` (see the Makefile).

Run (use Make)

Show CLI help:

```bash
make run ARGS="--help"
```

Run the `status` subcommand using the built-in sample stack (Make sets `BASE_PATH=.testStack` by default):

```bash
make run ARGS="status"
```

Configuration basics
- jsc expects a base directory that contains a cloud file and a stack file. At runtime `BASE_PATH` defaults to the current working directory unless you override it (the Make targets point to the sample under [.testStack](.testStack)).
- Required files inside `BASE_PATH`:
	- Cloud file (defaults to `.cloud`, override with `CLOUD_FILE_NAME`). A placeholder lives at [.testStack/.cloud](.testStack/.cloud).
	- Stack file (defaults to `.stack`, override with `STACK_FILE_NAME`). It must be of type `.rootstack` and can describe stacks and infra. Example from [.testStack/.stack](.testStack/.stack):

```
type: .rootstack
stacks:
	stackOne:
		dependsOn:
			- stackTwo
		type: folder
	stackTwo:
		src: stackTwo/.stack
infra:
	traefik:
		type: gitRepo
		src: https://github.com/Stelzi79/just-simple-cloud/traefik/.stack
```

CLI commands
- `jsc status` — currently the primary command; initializes the environment, reads the cloud/stack files, and prints a simple status summary.
- Flags: `--debug` to emit debug details, `--env` to include environment variables in debug output.

Other useful targets
- `make force-build` — clean + build
- `make install` — install the module
- `make test` — run `go test ./...` inside `jsc/`
- `make format` — run `go fmt ./...`
- `make tidy` — run `go mod tidy` inside `jsc/`
- `make help-message` — print help about available Make targets

Notes for developers
- CLI entrypoint: [jsc/main.go](jsc/main.go); commands are wired under [jsc/cmd](jsc/cmd).
- You can run locally without Make via `BASE_PATH=/your/stack go run ./jsc status --debug`, but prefer the Make targets for reproducible builds.

Contributing
- Fork the repo, create a feature branch, and open a pull request. Include tests for new behavior when possible.

License
- See the [LICENSE](LICENSE) file at the repository root for license terms.

