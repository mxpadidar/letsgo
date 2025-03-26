# Let's Go

`letsgo` is a starter project for building web applications with Go. It provides a good starting point and a basic structure for your web application.

## Features

- A simple web server with routing and middleware setup using go standard library.
- PostgreSQL database connection and migration management using `sqlx` and `goose`.
- Makefile for simple tasks like running container, migrations, and environment setup.

## What You Need Before You Start

The project uses a postgres container managed via `Podman`. If you prefer using a local or docker-based postgres instance, update the makefile accordingly.

- [Go](https://go.dev/doc/install) (version `1.22` or higher)
- [Podman](https://podman.io/docs/installation) (for running the PostgreSQL container)
- [Goose](https://pressly.github.io/goose/installation/) (for managing database migrations)

## Installation

```bash
# clone the repository and navigate into it
git clone https://github.com/mxpadidar/letsgo.git && cd letsgo

# install dependencies
go mod tidy

# move .env.example to .env
# change .env variables if needed
mv .env.example .env

# start postgresql container
# you can use a local or docker-based postgresql instance
make pgstart

# apply migrations
make migrate-up
```
