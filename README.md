# Test project for deployment of HTTP service in K8S


# Tech stack

Go 1.16
Makefile
Docker
Kubernetes 1.18


# Dev run

```
go run main.go
```

Test requests:
```
curl -v http://localhost:10000
```


# Build

```
make build
```

Binary execution:
```
./test-adv-tech
```

