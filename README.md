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
  --group keyrequest \
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
```

Regenerate code and run against the Kubernetes cluster configured by `~/.kube/config`:
```
make run
```

Create a KeyRequest Custom Resource:
```bash
kubectl apply -f config/samples/keyrequest_v1_keyrequest.yaml
kubectl delete -f config/samples/keyrequest_v1_keyrequest.yaml
```

Build the docker image:
```bash
docker build -t shubhamtatvamasi/quantum-manager:v0.1.0 .
docker push shubhamtatvamasi/quantum-manager:v0.1.0
```

Create a deployment:
```bash
kubectl create deployment quantum-manager --image=shubhamtatvamasi/quantum-manager:v0.1.0
kubectl delete deployment quantum-manager
```
