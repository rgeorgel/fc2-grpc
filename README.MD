# Studies of gRPC with Go

On this repo was implemented 4 ways of communication between the client and the server, using gRPC

### Implemented methods
- Single request
- Server stream
- Client stream
- Server and client stream


## Environment setup

On your terminal, execute the command `docker-compose up -d` to start the container.

After it executes `docker-compose exec grpc bash` to access the container, and in this 
container execute `go run cmd/server/server.go` to start the server.

Open a new terminal and execute `docker-compose exec grpc bash` again to access the 
container. Inside of this container, execute one of the commands to start the client.

### Clients option:
**Single request**

`go run cmd/client/client.go`

**Server stream**

`go run cmd/client/client-verbose.go`

**Client stream**

`go run cmd/client/client-stream.go`

**Server and client stream**

`go run cmd/client/client-stream-both.go`
