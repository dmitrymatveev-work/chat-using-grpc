package main

import (
	"bufio"
	pb "chat/chat-client/grpc"
	"fmt"
	"log"
	"os"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Introduce(ctx, &pb.IntroRequest{Username: username})
	if err != nil {
		log.Fatalf("could not introduce: %v", err)
	}
	log.Printf("Response: %s", r.Message)

	fmt.Println("Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
