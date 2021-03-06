package main

import (
	"auth-study/handler"
	pb "auth-study/proto"

	"github.com/asim/go-micro/plugins/auth/jwt/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
)

func main() {
	auth := jwt.NewAuth()

	// Create service
	srv := micro.NewService(
		micro.Name("auth"),
		micro.Auth(auth),
	)

	// Register handler
	pb.RegisterAuthStudyHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
