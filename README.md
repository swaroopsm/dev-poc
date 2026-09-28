# dev-poc

POC for a declarative dev environment based on [mise](https://mise.jdx.dev) and [colima](https://github.com/abiosoft/colima).

Two toy servers (Rust and Go) run on the host, colima runs a containerd VM with Caddy reverse-proxying `/rust` and `/go` to them.

## Tasks

Run `mise tasks` for the full list. Common ones:

- `mise run bootstrap` — start colima and bring up the containers (Caddy)
- `mise run rust:dev` / `mise run go:dev` — run a server, restarting on file changes
