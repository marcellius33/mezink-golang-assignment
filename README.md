# Mezink Golang Assignment

## Things need to be installed to run
1. Docker
2. Docker compose
3. Go

## To run project
1. run `docker compose up --build` to run database and application

## To test the endpoint
```
curl --location 'localhost:8080/records' \
--header 'Content-Type: application/json' \
--data '{
    "startDate": "2023-01-01",
    "endDate": "2024-05-01",
    "minCount": 180,
    "maxCount": 224
}'
```