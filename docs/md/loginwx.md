Login system by weixin.


## request

| field | type | required | description |
|:------|:------------|:---:|:----|
| code | string | Y | weixin auth code. |



## response

| field | type | description |
|:------|:----|:------------|
| errcode | int | error code: 200 - success, 400 - bad request, 404 - not found, 500 - internal server error.|
| token | string | token string. |

## example

### request

```bash

curl "http://localhost:8820/v1/user/loginwx" \
  -i \
  -X 'POST' \
  -d '{"code":"fsafsdfasfsa"}' 

```

#### response

```json

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 19 Sep 2022 15:16:49 GMT
Content-Length: 74

{
    "errcode": 41002,
    "token": "appid missing, rid: 632887e1-53ee34d6-6e3015e3"
}

```

