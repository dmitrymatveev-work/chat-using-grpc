package main

import (
	"bufio"
	pb "chat/chat-client/grpc"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = "50051"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter server name/address: ")
	address, _ := reader.ReadString('\n')
	address = strings.Trim(address, "\r\n")

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", address, port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatClient(conn)

	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.Trim(username, "\r\n")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	r, err := c.Introduce(ctx, &pb.IntroRequest{Username: username})
	if err != nil {
		log.Fatalf("could not introduce: %v", err)
	}
	log.Printf("Response: %s", r.Message)

	stream, err := c.Connect(context.Background())
	waitc := make(chan struct{})
	go func() {
		for {
			post, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a post : %v", err)
			}
			if post.Username == username {
				continue
			}
			fmt.Printf("%s: %s\n", post.Username, strings.Trim(post.Message, "\r\n"))
		}
	}()

	{
		err := stream.Send(&pb.Post{Username: username, Message: "Entered."})
		if err != nil {
			log.Fatalf("Failed to send a post: %v", err)
		}
	}

	go func() {
		for {
			message, _ := reader.ReadString('\n')
			err := stream.Send(&pb.Post{Username: username, Message: message})
			if err != nil {
				log.Fatalf("Failed to send a post: %v", err)
			}
		}
	}()

	<-waitc

	fmt.Println("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
