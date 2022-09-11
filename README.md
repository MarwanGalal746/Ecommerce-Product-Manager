<div align="center">
    <h1><strong>Ecommerce-Product-Manager</strong></h1>
</div>


##  Table of contents
- [**Summary**](#summary)
- [**Getting started**](#getting-started)
- [**Class Diagram**](#class-diagram)
- [**ERD Diagram**](#erd-diagram)
- [**Prerequisites**](#prerequisites)
- [**Packages**](#packages)
- [**Environment variables file**](environment-variables-file)
- [**Running**](#running)
- [**Postman collection**](#postman-collection)
---
## Summary

One of the main challenges of building an ecommerce platform is to keep an accurate
list of products and their stocks up to date. It's a system that allows us to manage products for an
hypothetical ecommerce platform.

## Getting Started

Clone the repository.<br />
Follow the instructions to complete the installation.

## Class Diagram

![](diagrams/class%20diagram/classDiagram.png)

## ERD Diagram

![](diagrams/ERD%20diagram/ERDDiagram.png)

## Prerequisites

- [Golang](https://golang.org/dl/)
- [Postgresql](https://www.postgresql.org/download/)

## Packages

- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/go-gorm/gorm)

## Environment variables file

List of environment variables which you should configure in your OS

```bash
ECOM_LOG_LEVEL=		           # log level
ECOM_WRITE_LOGS_TO_CONSOLE=    # boolean variable indicates writing variables to console
ECOM_PG_DB_HOST=               # PostgreSQL DB host
ECOM_PG_DB_PORT=               # PostgreSQL DB port
ECOM_PG_DB_USER=               # PostgreSQL DB user
ECOM_PG_DB_PASSWORD=           # PostgreSQL DB password
ECOM_PG_DB_NAME=               # PostgreSQL DB name
```

## Running

- In your cloned directory.
- open your terminal and run:

```bash
go run main.go
```

The server will start at:

- Local: http://localhost:8888

## Postman collection

you will find the postman collection [here](postman%20collection/ecom.postman_collection.json) to get how to send an HTTP request in right way to the system.

You can get what is the right structure of JSON file to send requests and recieving responses from the postman collection after importing it in the [Postman](https://www.postman.com/).