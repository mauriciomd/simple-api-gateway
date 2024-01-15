# API Gateway

This project implements an API Gateway responsible for reading a configuration file (config.yml) and forwarding the request to the internal services. All request headers and body will be sent to the internal service. Also, the gateway will preserve the service response content, returning all headers and body provided by the called service.

## Implementation
The repository has three folders:
```
- .
-   /gateway
-   /payments
-   /shippings
```
The folder Gateway implements a simple API Gateway using Go. The service reads a config file that contains definitions for all internal endpoints and plugins. Plugins are executed before a request is sent to the internal service and it could be an API Logging or a JWT verification, for example.

Finally, both payments and shippings folders contain simple NodeJs webserver, exposing a GET and POST endpoint simulating internal services.

## How to run
To run the project, just use
```
docker-compose up
```
All requests can be done using the address ```http://localhost:8080```.
