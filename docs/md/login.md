Login system by username and password.


## request

| field | type | required | description |
|:------|:------------|:---:|:----|
| username | string | Y | user's name. |
| password | string | Y | user's password. |



## response

| field | type | description |
|:------|:----|:------------|
| errcode | int | error code: 200 - success, 400 - bad request, 404 - not found, 500 - internal server error.|
| token | string | token string. |

## example

### request

```bash

curl "http://localhost:8820/v1/user/login" \
  -i \
  -X 'POST' \
  -d '{"username":"alice", "password": "123456"}' 

```

#### response

```json

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 19 Sep 2022 15:10:53 GMT
Content-Length: 205


{
    "errcode": 0,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVfdGltZSI6MTY2MzY4NjY1MywidXNlcl9pZCI6IjEwMzM3ZmM3LWE2ZjEtNDM0My05NTYxLWRmNTZkZjkwMTFiMCJ9.slCyGdF5MTYQCD26b-hENPvlQxrjxRClvG-J9-4LZGY"
}

```

