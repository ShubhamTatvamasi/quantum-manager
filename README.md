# quantum-manager

Initialize the project with kubebuilder:
```bash
kubebuilder init \
  --domain quantum-manager.io \
  --repo github.com/ShubhamTatvamasi/quantum-manager
```

Create a new API:
```bash
kubebuilder create api \
  --version v1 \
  --kind KeyRequest \
  --resource \
  --controller
```

Generate the manifests:
```
make manifests
```

Install CRDs into the Kubernetes cluster using kubectl apply:
```
make install
make uninstall
```

Regenerate code and run against the Kubernetes cluster configured by `~/.kube/config`:
```
make run
```

Create a KeyRequest Custom Resource:
```bash
kubectl apply -k config/samples/
kubectl delete -k config/samples/
```

Export the docker image:
```bash
export IMG=shubhamtatvamasi/quantum-manager:v0.1.2
```

Build the docker image:
```bash
make docker-build docker-push
```

Create a deployment:
```bash
make deploy
make undeploy
```
