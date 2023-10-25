# Golang DDD Boilerplate

This is the sample boilerplate of Golang DDD (Domain, Driven Design)

## Table of Contents

- [Prerequisites](#Prerequisites)
- [Installation](#installation)
- [Run Application](#running-application)
- [Migration](#migration)
- [Swagger](#swagger)
- [Test](#test)

## Prerequisites

List any prerequisites or dependencies your project requires to run. For example:

- Download and install [Go 1.19 or higher](https://go.dev/doc/install)
- Install [Golang Migrate](https://github.com/golang-migrate/migrate) 
- Download and Install [Posgresql](https://www.postgresql.org/download/)
- Install [Swaggo](https://github.com/swaggo/swag)
- Download and Install [Docker (if needed)](https://www.docker.com/products/docker-desktop/)

## Installation

Provide step-by-step instructions on how to install or set up your project. For example:

1. Clone the repository
   ```bash
   git clone ssh://git@gitlab.banksinarmas.com:1022/digital-banking/backend/personal-banking.git && cd sample-ddd-boilerplate

2. Install golang package
    ```bash
   go mod tidy

## Running Application
1. Running using docker-compose
    ```bash
   docker compose up -d
   
2. Running golang application
    ```bash
   go run cmd/main.go 

2. Running golang application using makefile
    ```bash
   make start-server

## Migration

1. create new migration
    ```bash
    migrate create -ext sql -dir migration -seq migration-name

2. apply migration bash
    ```bash
   make run-migration-up

3. apply migration manual
    ```bash
    make migrateup
you can see this [repository](https://github.com/golang-migrate/migrate) to get complete documentation

## Swagger
1. generate swagger documentation
    ```bash
   swag init -g main.go 
you can see this [repository](https://github.com/swaggo/swag) to get complete documentation

## Running Test
1. Running test
    ```bash
   make test
2. Running coverage test
    ```bash
   make test-coverage