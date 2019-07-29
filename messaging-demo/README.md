# Messaging demo

Example of how to connect to a remote RabbitMQ messaging broker and send a message to a queue

## Pre-requisites

  -  Install [go dep](https://golang.github.io/dep/)
  -  Fetch dependencies by running
  ```
  dep ensure
  ```

## Running

Provide the credentials for the RabbitMQ messaging broker:

```
go run main.go 
```