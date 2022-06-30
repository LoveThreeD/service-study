package main

import (
	"context"
	"log"
	pb "wrapper-study/assets/proto"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/server"
)

type Greeter struct{}

func (r Greeter) Hello(_ context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

// logWrapper is a handler wrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[wrapper] server request: %v", req.Endpoint())
		err := fn(ctx, req, rsp)
		return err
	}
}

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("greeter"),
		micro.WrapHandler(logWrapper),
	)

	server.Init()
	pb.RegisterGreeterHandler(srv.Server(), new(Greeter))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
