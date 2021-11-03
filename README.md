# Kafman
Kafman is a simple HTTP API for publishing Kafka message & console log for kafka's topic consumer

# Prerequisites
- Go 1.15
- Kafka

## Usage

- Update config file config.json as your kafka configuration
- run with 

```
go run main.go
```
- for local purpuse, you can run with docker compose :
```
docker-compose up
```
- Sample Request for Publishing message to kafka, you can set message with string or json object
```
curl --location --request POST 'localhost:3000/kafman/v1/publish/mytopic' \
--header 'Content-Type: text/plain' \
--data-raw 'ini topic mytopic'
```

```
curl --location --request POST 'localhost:3000/kafman/v1/publish/mytopic' \
--header 'Content-Type: application/json' \
--data-raw '{
    "topic": "yourtopic",
    "message": {
        "key": "is a key",
        "value": "is a value",
        "other": "what is this? :D"
    }
}'
```

## Contributor
- [mauwahid](https://www.linkedin.com/in/mauwahid/)

## License
- [MIT](https://choosealicense.com/licenses/mit/)
