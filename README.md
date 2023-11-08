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
$ kubectl get pods -A
NAMESPACE       NAME                                     READY   STATUS             RESTARTS          AGE
alwaysbespoke   coba-5d46789454-xhpfh                    1/1     Running            1 (21s ago)       66s
```

