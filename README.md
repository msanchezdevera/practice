# How yo run locally

Clone the project and follow these instructions:

## Requisits
- Install Go 1.15

## Execution

Set the corresponding environment:

`export ENVIRONMENT=local`

Execute the comand:

`go run main.go`

You should see the following message:

```

```

# Tests

## How to execute tests

On the root folder execute the command:
```
go test ./...
```

## Coverage

To test code coverage, execute the following commands:
```
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

# Examples

### Create a transaction
```
curl -X POST \
  http://localhost:8080/transactions \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "type": "credit",
    "amount": 222
}'
```