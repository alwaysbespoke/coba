# coba
Coba is a configurable SIP server written in Go.

# Local Development
1. Install Kubernetes locally (eg [Docker Desktop](https://www.docker.com/products/docker-desktop/)).

2. Run the installation script.
```bash
$ make up-local
```

3. Verify installation.
```bash
$ kubectl get pods -n alwaysbespoke
NAME                          READY   STATUS             RESTARTS      AGE
sbc-service-d96855f56-hd2lp   1/1     Running            0             13m
sip-server-65d4f7d9fc-lw9md   1/1     Running            0             15m
```

