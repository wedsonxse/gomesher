# Gomesher

A go program to explore the golang language. It uses STOMP, HTTP communication and Go routines.

# Configuration

Add the informations to the .env file and all should be fine. In the original example, the university API (https://github.com/Hipo/university-domains-list) is being used to increment data trough the user input.

# Execution

Run this command to start the execution:

```sh
go run main.go
```

# How it works

The program will ask periodically for a user input. This input will be used to perform a HTTP request to the universities API. The results obtained from this request will be published to a queue in a stomp server. The messages are then consumed by our service that will be watching that queue as well. The producer and Consumer perform parallel tasks using go routines.
