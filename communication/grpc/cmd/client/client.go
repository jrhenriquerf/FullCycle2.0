package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jrhenriquerf/fc2-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jair",
		Email: "jj@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jair",
		Email: "jj@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive stream: %v", err)
		}

		fmt.Println("Status: ", stream.Status, stream.User)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "j1",
			Name:  "Jair 1",
			Email: "jair1@gmail.com",
		},
		{
			Id:    "j2",
			Name:  "Jair 2",
			Email: "jair2@gmail.com",
		},
		{
			Id:    "j3",
			Name:  "Jair 3",
			Email: "jair3@gmail.com",
		},
		{
			Id:    "j4",
			Name:  "Jair 4",
			Email: "jair4@gmail.com",
		},
		{
			Id:    "j5",
			Name:  "Jair 5",
			Email: "jair5@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating reuest: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response %v", err)
	}

	fmt.Print(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		{
			Id:    "j1",
			Name:  "Jair 1",
			Email: "jair1@gmail.com",
		},
		{
			Id:    "j2",
			Name:  "Jair 2",
			Email: "jair2@gmail.com",
		},
		{
			Id:    "j3",
			Name:  "Jair 3",
			Email: "jair3@gmail.com",
		},
		{
			Id:    "j4",
			Name:  "Jair 4",
			Email: "jair4@gmail.com",
		},
		{
			Id:    "j5",
			Name:  "Jair 5",
			Email: "jair5@gmail.com",
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

			fmt.Printf("Receiving user %v com status %v\n", res.GetUser().GetName(), res.GetStatus())
		}
		close(wait)
	}()

	<-wait
}
