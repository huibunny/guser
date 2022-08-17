# TEST

## deregister service

```bash

curl -X PUT http://dog.ap:8500/v1/agent/service/deregister/user_172.16.12.8:8820

```

## login 

```bash

curl -X POST -d '{"username":"alice", "password": "123456"}' "http://localhost:9090/api/user/v1/user/login"

```