package main

import (
    "context"
    "fmt"
    "github.com/rgeorgel/fc2-grpc/pb"
    "google.golang.org/grpc"
    "io"
    "log"
)

func main() {
    connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect to the gRPC server: %v", err)
    }

    defer connection.Close()

    client := pb.NewUserServiceClient(connection)
    AddUserVerbose(client)
}

func AddUserVerbose(client pb.UserServiceClient) {
    req := &pb.User{
        Id: "0",
        Name: "Ricardo",
        Email: "rickygeorgel@yahoo.com.br",
    }

    res, err := client.AddUserVerbose(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not make gRPC request: %v", err)
    }

    for {
        stream, err := res.Recv()
        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatalf("Could not receive the response: %v", err)
        }

        fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
    }
}
