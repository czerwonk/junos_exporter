# Helm chart for helm v3

## How to use

### Authentication

#### SSH key authentication
Add your ssh-keyfile and config.yml to values.yml

> sshkey is the base64 encoded id_rsa you want to use for authentication  
> generate sshkey with `cat $HOME/.ssh/id_rsa | base64 -w0 && echo`

`sshkey: QWRkIHlvdXIgb3duIGlkX3JzYSBoZXJl`

#### Password authentication
To use password authentication the following values.yaml configuration could
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

### Devices configuration
Add your devices to the devices in configyml in values.yml

### Installation
> cd helm  
> helm install junosexporter ./junosexporter 
