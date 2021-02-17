package main

import (
    "context"
    "fmt"
    "github.com/rgeorgel/fc2-grpc/pb"
    "google.golang.org/grpc"
    "log"
)

func main() {
    connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect to the gRPC server: %v", err)
    }

    defer connection.Close()

    client := pb.NewUserServiceClient(connection)
    AddUser(client)
}

func AddUser(client pb.UserServiceClient) {
    req := &pb.User{
        Id: "0",
        Name: "Ricardo",
        Email: "ricardo@ricardo.com",
    }

    res, err := client.AddUser(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not make gRPC request: %v", err)
    }

    fmt.Println(res)
}
