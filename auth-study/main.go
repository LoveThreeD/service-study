package main

import (
	"auth-study/handler"
	pb "auth-study/proto"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("auth-study"),
	)

	// Register handler
	pb.RegisterAuthStudyHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
