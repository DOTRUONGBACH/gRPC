package main

import (
	"context"
	"ent-grpc-example/ent"
	"ent-grpc-example/ent/proto/entpb"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=ok password=postgres sslmode=disable ")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	svc := entpb.NewUserService(client)
	server := grpc.NewServer()
	entpb.RegisterUserServiceServer(server, svc)
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening :%s", err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}

}
