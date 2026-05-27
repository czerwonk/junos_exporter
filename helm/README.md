# Helm chart for helm v3

## Installing the chart

The chart is published to the GitHub Container Registry as an OCI artifact on every release.

```shell
helm install junos-exporter oci://ghcr.io/czerwonk/charts/junos-exporter --version <version>
```


### Local installation from source

```shell
cd helm
helm install junosexporter ./junosexporter
```

## Authentication

### SSH key authentication
Add your SSH key and `configyml` to `values.yml`.

`sshkey` is a base64-encoded SSH private key you want to use for authentication.

You can generate an `sshkey` with `cat $HOME/.ssh/id_rsa | base64 -w0 && echo`:
```yaml
sshkey: QWRkIHlvdXIgb3duIGlkX3JzYSBoZXJl
```

It is also possible to use the existing-secret pattern (e.g. with ExternalSecrets operator),
the secret with the SSH key should be mounted via `extraVolumes` and `extraVolumeMounts`.

### Password authentication
To use password authentication the following `values.yaml` configuration could
be used with a `junos-exporter-ssh` secret object storing SSH secrets:

``` yaml
extraArgs:
- "-ssh.targets=$(JUNOS_EXPORTER_SSH_TARGETS)"
- "-ssh.user=$(JUNOS_EXPORTER_SSH_USER)"
- "-ssh.password=$(JUNOS_EXPORTER_SSH_PASSWORD)"

extraEnv:
- name: JUNOS_EXPORTER_SSH_TARGETS
  valueFrom:
    secretKeyRef:
      name: junos-exporter-ssh
      key: targets
- name: JUNOS_EXPORTER_SSH_USER
  valueFrom:
    secretKeyRef:
      name: junos-exporter-ssh
      key: username
- name: JUNOS_EXPORTER_SSH_PASSWORD
  valueFrom:
    secretKeyRef:
      name: junos-exporter-ssh
      key: password
```

``` yaml
apiVersion: v1
kind: Secret
type: Opaque
data:
  password: BASE64_ENCODED_SSH_PASSWORD
  targets: BASE64_ENCODED_SSH_TARGETS
  username: BASE64_ENCODED_SSH_USERNAME
```

## Devices configuration
Add your devices to the devices in `configyml` in `values.yaml`

## Handling configuration/authorization changes
To force reload of the exporter pods upon `configyml` or `sshkey` configuration changes,
enable the `rollOutJunosExporterPods` option in `values.yaml`.

If Reloader controller is installed in the cluster, for `extraEnv` passwords or `extraVolumes` keys
the `annotations` map in `values.yaml` can be used to specify a policy to handle the updates:

```yaml
annotations:
  reloader.stakater.com/auto: "true"
```

## Versioning

The chart version tracks the application version — `chart version 0.15.1` packages
`junos_exporter v0.15.1`. Chart-only fixes (template changes, new values) are released as
patch bumps independent of the application version.

## Contributing to the chart

Any pull request that modifies chart files (`templates/`, `values.yaml`, etc.) **must** also
bump the `version` field in `Chart.yaml`. Without a `version` bump the Helm Chart Publish
workflow will not trigger and no new chart artifact will be published.

To verify the chart is valid before opening a PR:

```shell
helm lint helm/junosexporter
```
