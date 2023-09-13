# MTA Hosts Service

## Introduction
The MTA Hosts Service is a software solution designed to help identify inefficient servers hosting mail transfer agents (MTAs) with few active IP addresses. This README provides an overview of the service, its components, and how to use it.

## Table of Contents
- Features
- Prerequisites
- Getting Started
- API Endpoint
- Sample Data
- Testing

## Features:
- Retrieve hostnames with less than or equal to X active IP addresses.
- Configurable threshold (X) using an environment variable (default is set to 1).

## Prerequisites
Before using the MTA Hosts Service, make sure you have the following prerequisites installed:

- Go programming language (https://golang.org/dl/)
- Git (https://git-scm.com/downloads)

## Getting Started
To get started with the MTA Hosts Service, follow these steps:

1. Clone the repository:
``` git clone <repository-url>```,
```  cd Mta-Hosts```
2. Run the server:
``` go run main.go```

The service should now be running and accessible at the specified port.

## API Endpoint
The service exposes a HTTP/REST endpoint to retrieve hostnames with less than or equal to X active IP addresses. You can make a GET request to the following endpoint:
- GET /hostnames?threshold=X: Retrieve hostnames with active IP addresses less than or equal to the specified threshold (X).

Example:
```curl http://localhost:8090/hostnames?threshold=2```

## Sample Data
The MTA Hosts Service includes sample data provided by a mock service. This sample data simulates IP configuration information for mail transfer agents. You can replace this data with your own data source as needed.

IP (String) IP address. 

Hostname (String) Hostname of the underlying server

Active (Boolean) Defines whether an IP
address is actively used or not.

## Testing
The project includes unit tests for the service, handler, and store components. To run the tests, use the following command:
```go test ./...```