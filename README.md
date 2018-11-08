# Inventory Management API Server

This is the RESTful API Server for Inventory Management.

## Quickstart

### Docker
```bash
docker build -t 404busters/inventory-management-apiserver:latest .
docker run 404busters/inventory-management-apiserver:latest
```

### Build from source
```bash
go build gitlab.com/404busters/inventory-management/apiserver/cmd/apiserver
```

## Development

This application would follow guideline from [The Twelve-Factor App](https://12factor.net/).

### Environment

Golang: 1.11.x
OS: Mac or Linux (POSIX)

### Install Golang

For Mac,
```bash
brew install go
```

### Fetch Dependencies

For regular, go would automatic fetch packages based on `go.mod` and `go.lock`. 
For prefetch for IDE or other purpose, just run `go mod download`

### Add Library
Just directly run `go get [package]`
