package main

import (
	"cache-study/handler"
	pb "cache-study/proto"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
)

var (
	service = "go.micro.srv.cache"
	version = "latest"
)

func main() {

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterCacheHandler(srv.Server(), handler.NewCache())

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
