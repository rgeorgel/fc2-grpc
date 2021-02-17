package main

import (
    "context"
    "fmt"
    "github.com/rgeorgel/fc2-grpc/pb"
    "google.golang.org/grpc"
    "io"
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
    AddUsersStreamBoth(client)
}

func AddUsersStreamBoth(client pb.UserServiceClient) {
    stream, err := client.AddUserStreamBoth(context.Background())
    if err != nil {
        log.Fatalf("Error creating request: %v", err)
    }

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

    wait := make(chan int)

    go func() {
      for _, req := range reqs {
          fmt.Println("Sending user: ", req.Name)
          stream.Send(req)
          time.Sleep(time.Second * 2)
      }
      stream.CloseSend()
    }()

    go func() {
        for {
            res, err := stream.Recv()
            if err == io.EOF {
                break
            }
            if err != nil {
                log.Fatalf("Error receiving data: %v", err)
                break
            }
            fmt.Printf(
                "Receiving user %v with status %v \n",
                res.GetUser().GetName(),
                res.GetStatus(),
            )
        }
        close(wait)
    }()

    <-wait
}
