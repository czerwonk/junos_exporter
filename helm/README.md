# Helm chart for helm v3

## How to use
Add your ssh-keyfile and config.yml to values.yml

> sshkey is the base64 encoded id_rsa you want to use for authentication  
> generate sshkey with `cat $HOME/.ssh/id_rsa | base64 -w0 && echo`

`sshkey: QWRkIHlvdXIgb3duIGlkX3JzYSBoZXJl`

Add your devices to the devices in configyml in values.yml

> cd helm  
> helm install junosexporter ./junosexporter 
