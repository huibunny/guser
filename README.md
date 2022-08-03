# go-clean-template

## build

```bash

#  create docs.go at  docs/docs.go
#  create swagger.json at  docs/swagger.json
#  create swagger.yaml at  docs/swagger.yaml
$ swag init --parseDependency --parseInternal -g cmd/app/main.go
# $ swag init -g internal/controller/http/v1/router.go

# run build.sh and it will output guser file.
$ ./build.sh

```

## run

```bash

$ ./goclean -consul localhost:8500 -name hello -listen :9090

```

## function

* remove http port from config file and use the one from cmd arguments.
* support consul(register/deregister/kv)

