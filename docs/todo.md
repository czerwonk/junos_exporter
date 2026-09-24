# TODO

## Open
- [ ] (medium) Helm chart: `-config.file=/config/config.yml` is only rendered when `extraArgs` is set, so `configyml` alone is ignored
- [ ] (low) Helm chart: enable `runAsNonRoot: true` and `readOnlyRootFilesystem: true` by default

## Done
- [x] Switch runtime image from alpine to `gcr.io/distroless/static-debian13:nonroot`
- [x] Multi-arch Docker build (amd64, arm64, arm/v7) with cross-compilation, `-trimpath -ldflags="-s -w"`

## Decisions
- Distroless has no shell, so the legacy env vars (`SSH_KEYFILE`, `CONFIG_FILE`, `ALARM_FILTER`, `CMD_FLAGS`) are mapped to flags by the binary itself (`docker_env.go`), only when started without arguments to mirror Docker's CMD override semantics.
- `CMD` in exec form instead of `ENTRYPOINT`, so existing overrides like `./junos_exporter -flag ...` keep working.
- Known incompatibility: the image runs as UID 65532 instead of root, mounted files must be readable by that user.
