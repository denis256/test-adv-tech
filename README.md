# Test project for deployment of HTTP service in K8S


# Tech stack

Go 1.16
Makefile
Docker
Kubernetes 1.18
Helm 3


# Dev run

```
go run main.go
```

Test requests:
```
curl -v http://localhost:10000
curl -v http://localhost:10000/request
```


# Build

```
make build
```

Binary execution:
```
./test-adv-tech
```

# Container 

To build container with application run
```
make container
```

It will build build container tagged from current git commit hash

Example local execution

```
docker run -p 10000:10000 denis256/test-adv-tech:61d8fd5
```

# Kubernetes deployment

To deploy application in K8S, use helm charts, directory `helm`

Worklflows

Basic installation
```
kubectl create ns adv-tech

cd helm
helm install adv-tech-app -n adv-tech -f test-adv-tech-values.yaml test-adv-tech
```

Check application status
```
helm -n adv-tech ls 
kubectl -n  adv-tech  get pods 
```

Access application through node port:
```
export NODE_PORT=$(kubectl get --namespace adv-tech -o jsonpath="{.spec.ports[0].nodePort}" services adv-tech-app-test-adv-tech)
export NODE_IP=$(kubectl get nodes --namespace adv-tech -o jsonpath="{.items[0].status.addresses[0].address}")

curl -v http://$NODE_IP:$NODE_PORT/request

```

Uninstall app
```
helm -n adv-tech  uninstall adv-tech-app
```

# Future work

* CI/CD pipeline for auotmated building
* Basic tests and collection of test coverage
* Automated build and deployment with tools like ansible/terraform

# License

Only for reference
