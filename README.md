# Spyglass

Spyglass is a micro-api that allows exposing domain names hosted in the Kubernetes cluster.

## How to use it
To expose a domain, include the annotation `spyglass/name` on the ingress manifest on the annotation section:
```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: test-ingress-enabled
  annotations:
    spyglass/name: "apiName"
```

A path for the discovery endpoint can also be included:
```yaml
  name: test-ingress-enabled
  annotations:
    spyglass/name: "apiName"
    spyglass/endpoint: "discovery"
```

The result of doing a request to `/` will be a json with the data corresponding to the Ingress domain.

```json
{
    "apiName": {
        "url": "test-enabled.local.com",
        "discovery": ""
    }
}
```

## Run in Kubernetes

### Using Helm
```bash
cd ./helm/spyglass
helm install spyglass .
```

### Using kubectl
```bash
kubectl create clusterrolebinding spyglass --clusterrole=view --serviceaccount=default:default
kubectl run --rm -i spyglass --image=fpaganetto/spyglass
kubectl port-forward deployment/spyglass 8090:8090
curl localhost:8090
```

### Add Ingress example
```
kubectl create -f ./example/ingress_example.yam
```
