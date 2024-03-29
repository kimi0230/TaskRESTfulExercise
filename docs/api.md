# APIs

## 1. Heartbeat

### Ping

#### HTTP Request
```sh
curl --location 'http://127.0.0.1:8080/ping'
```

#### HTTP Response

```json
{
    "data": "pong",
    "code": "200-API-V1-0001",
    "error": ""
}
```

| Parameter | type   | Description |
|-----------|--------|-------------|
| data      | string | message     |


#### HTTP Response Headers