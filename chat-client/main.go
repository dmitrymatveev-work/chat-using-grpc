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

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewChatClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.Trim(username, "\r\n")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
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
			fmt.Printf("%s: %s\n", post.Username, post.Message)
		}
	}()

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
