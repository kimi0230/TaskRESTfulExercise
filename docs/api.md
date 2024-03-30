# APIs

Postman sample : [Link](task-restful-exercise.postman_collection.json)

- [APIs](#apis)
  - [0. Heartbeat](#0-heartbeat)
    - [Ping](#ping)
      - [HTTP Request](#http-request)
      - [HTTP Response](#http-response)
      - [HTTP Response Headers](#http-response-headers)
  - [1. Tasks](#1-tasks)
    - [1.1 Get Tasks](#11-get-tasks)
      - [HTTP Request](#http-request-1)
      - [HTTP Response](#http-response-1)
      - [HTTP Response Headers](#http-response-headers-1)
    - [1.2 Post Tasks](#12-post-tasks)
      - [HTTP Request](#http-request-2)
      - [HTTP Response](#http-response-2)
      - [HTTP Response Headers](#http-response-headers-2)
    - [1.3 Put Tasks](#13-put-tasks)
      - [HTTP Request](#http-request-3)
      - [HTTP Response](#http-response-3)
      - [HTTP Response Headers](#http-response-headers-3)
    - [1.4 Delete Tasks](#14-delete-tasks)
      - [HTTP Request](#http-request-4)
      - [HTTP Response](#http-response-4)
      - [HTTP Response Headers](#http-response-headers-4)


## 0. Heartbeat

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

---

## 1. Tasks

### 1.1 Get Tasks

`GET {{URL}}/api/v1/tasks`

#### HTTP Request
```sh
curl --location 'http://127.0.0.1:8080/api/v1/tasks?limit=100&keyword=&order=desc&by=created_at&page=1'
```

| Parameter | type   | Description                              | default |
|-----------|--------|------------------------------------------|---------|
| limit     | int    | (optional) The number of items per page  | 100     |
| page      | int    | (optional) The index of the current page | 1       |
| order     | string | (optional) desc / asc                    | desc    |
| by        | string | (optional) order field                   | id      |
| keyword   | string | (optional) search displayName            |         |

#### HTTP Response

```json
{
    "data": [
        {
            "_id": "660782e2d56526b8dac58a09",
            "name": "Nobody",
            "description": "Task-K-0001 description",
            "priority": 1,
            "status": 0,
            "dueDate": "2024-03-30",
            "created_at": "2024-03-30T03:11:30.837Z",
            "updated_at": "2024-03-30T03:11:30.837Z"
        },
        ...
    ],
    "code": "200-API-V1-0001",
    "error": ""
}
```

| Parameter   | type   | Description                           |
|-------------|--------|---------------------------------------|
| _id         | string | _id                                   |
| name        | string | task name                             |
| description | string | description                           |
| priority    | int    | task priority                         |
| status      | int    | 0: incomplete task; 1: completed task |
| dueDate     | string | dueDate                               |


#### HTTP Response Headers
| Header        | Description                   |
|---------------|-------------------------------|
| X-Page        | The index of the current page |
| X-Per-Page    | The number of items per page  |
| x-total       | The total number of items     |
| x-total-pages | The total number of pages     |


### 1.2 Post Tasks

`{{URL}}/api/v1/tasks`

#### HTTP Request
```sh
curl --location 'http://127.0.0.1:8080/api/v1/tasks' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Task-K-0001",
    "description": "Task-K-0001 description",
    "priority": 1,
    "dueDate": "2024-03-30"
}'
```

| Parameter | type | Description | default |
|-----------|------|-------------|---------|

#### HTTP Response

```json
{
    "data": {
        "_id": "66078543d56526b8dac58a14",
        "name": "Task-K-0001",
        "description": "Task-K-0001 description",
        "priority": 1,
        "status": 0,
        "dueDate": "2024-03-30",
        "created_at": "2024-03-30T03:21:39.638Z",
        "updated_at": "2024-03-30T03:21:39.638Z"
    },
    "code": "201-API-V1-0001",
    "error": ""
}
```

| Parameter   | type   | Description                           |
|-------------|--------|---------------------------------------|
| _id         | string | _id                                   |
| name        | string | task name                             |
| description | string | description                           |
| priority    | int    | task priority                         |
| status      | int    | 0: incomplete task; 1: completed task |
| dueDate     | string | dueDate                               |


#### HTTP Response Headers
| Header | Description |
|--------|-------------|


### 1.3 Put Tasks

PUT `{{URL}}/api/v1/tasks/{id}`

#### HTTP Request
```sh
curl --location --request PUT 'http://127.0.0.1:8080/api/v1/tasks/66078543d56526b8dac58a14' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Task-K-0001",
    "description":"Task-K-0001 description Updates",
    "priority": 3,
    "status": 0,
    "dueDate":"2024-05-06"
}'
```

| Parameter | type | Description | default |
|-----------|------|-------------|---------|

#### HTTP Response

```json
{
    "data": {
        "_id": "66078543d56526b8dac58a14",
        "name": "Task-K-0001",
        "description": "Task-K-0001 description Updates",
        "priority": 3,
        "status": 0,
        "dueDate": "2024-05-06",
        "created_at": "2024-03-30T03:21:39.638Z",
        "updated_at": "2024-03-30T03:26:57.889Z"
    },
    "code": "200-API-V1-0001",
    "error": ""
}
```

| Parameter   | type   | Description                           |
|-------------|--------|---------------------------------------|
| _id         | string | _id                                   |
| name        | string | task name                             |
| description | string | description                           |
| priority    | int    | task priority                         |
| status      | int    | 0: incomplete task; 1: completed task |
| dueDate     | string | dueDate                               |


#### HTTP Response Headers
| Header | Description |
|--------|-------------|

### 1.4 Delete Tasks

DELETE `{{URL}}/api/v1/tasks/{id}`

#### HTTP Request
```sh
curl --location --request DELETE 'http://127.0.0.1:8080/api/v1/tasks/66078543d56526b8dac58a14'
```

| Parameter | type | Description | default |
|-----------|------|-------------|---------|

#### HTTP Response

```json
{
    "data": null,
    "code": "200-API-V1-0001",
    "error": ""
}
```

| Parameter | type | Description |
|-----------|------|-------------|


#### HTTP Response Headers
| Header | Description |
|--------|-------------|