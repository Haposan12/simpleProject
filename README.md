# Simple Project Dans Multi Pro Test

## Installation

### Database
- Create MySQL Connection
- Create Database -> `test`

### On Local machine
- Clone this repository
- Open the source code, open terminal and run <code> go mod tidy </code>
- Run <code>go run main.go</code>

## Postman Collection

https://www.getpostman.com/collections/0da3167c654c041a1878

## Available API's
baseUrl: localhost:1234

 - POST {{baseUrl}}/login
 - GET {{baseUrl}}/job/list
 - GET {{baseUrl}}/job/detail/{id}

## Authorization
Bearer Token Authorization