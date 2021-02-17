package main

import (
    "context"
    "fmt"
    "github.com/rgeorgel/fc2-grpc/pb"
    "google.golang.org/grpc"
    "log"
    "time"
)

func main() {
    connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Could not connect to the gRPC server: %v", err)
    }

    defer connection.Close()

    client := pb.NewUserServiceClient(connection)
    AddUsers(client)
}

func AddUsers(client pb.UserServiceClient) {
    reqs := []*pb.User{
        &pb.User{
            Id: "1",
            Name: "Ricardo 1",
            Email: "ricardo@ricardo1.com",
        },
        &pb.User{
            Id: "2",
            Name: "Ricardo 2",
            Email: "ricardo@ricardo2.com",
        },
        &pb.User{
            Id: "3",
            Name: "Ricardo 3",
            Email: "ricardo@ricardo3.com",
        },
        &pb.User{
            Id: "4",
            Name: "Ricardo 4",
            Email: "ricardo@ricardo4.com",
        },
        &pb.User{
            Id: "5",
            Name: "Ricardo 5",
            Email: "ricardo@ricardo5.com",
        },
        &pb.User{
            Id: "6",
            Name: "Ricardo 6",
            Email: "ricardo@ricardo6.com",
        },
        &pb.User{
            Id: "7",
            Name: "Ricardo 7",
            Email: "ricardo@ricardo7.com",
        },
    }

    stream, err := client.AddUsers(context.Background())
    if err != nil {
        log.Fatalf("Error creating request: %v", err)
    }

    for _, req := range reqs {
        stream.Send(req)
        time.Sleep(time.Second * 3)
    }

    res, err := stream.CloseAndRecv()
    if err != nil {
        log.Fatalf("Error receiving response: %v", err)
    }

    fmt.Println(res)
}
