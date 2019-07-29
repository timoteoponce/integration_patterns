# Messaging hello

Example of how to connect to a remote RabbitMQ messaging broker 
and create a messaging channel between a sender and receiver.

## Pre-requisites

  -  Install [go dep](https://golang.github.io/dep/)
  -  Fetch dependencies by running
  ```
  dep ensure
  ```

## Running

Provide the credentials for the RabbitMQ messaging broker when starting, to run this example you need 
to execute two applications:

```
go run receiver.go 
```

To run the messages receiver, and then start the sender:
```
go run sender.go 
```
