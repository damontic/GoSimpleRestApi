# Go Simple Rest API

Used for trainings in CI/CD and DevOps.

## To run the application

```shell
go run simpleRest.go
```

## To format the code

```shell
go fmt simpleRest.go
```

## To check style

```shell
golint
```

## To check suspicious constructs

```shell
go vet
```

## To install

Make sure that the $GOPATH/bin directory is in your $PATH.

```shell
go build -x github.com/damontic/GoSimpleRestApi
```

## This specifies a software pipeline

The app is compatible with Jenkins pipelines. It defines its own Jenkinsfile.

## This app can be monitored using Prometheus

The app uses the prmetheus go client to expose a metrics endpoint.

