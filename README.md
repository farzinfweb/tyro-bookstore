
# Tyro Book Store

This is a toy project to learn some golang and software architecture concepts


## Compile proto

To compile proto file, run below command in your terminal (you need to have proto compiler installed on your system):

```bash
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/bookstore.proto
```

## Install dependencies

```bash
  go mod tidy
```

## Run

Server (this would run the server at localhost:7080):

```bash
  go run server/main.go
```

Client (this would run the server at localhost:1323):

```bash
  go run client/main.go
```

## Generate fake data

Run the following command to generate 100 books and store them in the database:

```bash
  go run scripts/book_generator.go
```