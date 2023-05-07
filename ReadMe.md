# Simple To Do List

## Description
Simple crud todo list with unit and integration testing for pikpo challenge

## Stack
- Golang
- gRPC
- PostgreSQL

## How To Start
1. Setup `.env` by following `.env.example`
2. Open terminal
3. Make sure you in the right path
4. Type `make run` for start or `make test` for start the testing

## Activity API
- if you use postman, import the proto file from proto folder
### Add Activity
* Request Body
```
{
    "days": "friday",
    "description": "testing"
}
```

### Get One Activity
* Request Body
```
{
    "id: 1
}
```

### Get All Activity
* Request Body
```
{

}
```

### Update Activity
* Request Body
```
{
    "id": 8,
    "description": "do the challenge"
}
```

### Delete Activity
* Request Body
```
{
    "id": 8
}
```
