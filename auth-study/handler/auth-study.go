package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	authstudy "auth-study/proto"
)

type AuthStudy struct{}

// Return a new handler
func New() *AuthStudy {
	return &AuthStudy{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *AuthStudy) Call(ctx context.Context, req *authstudy.Request, rsp *authstudy.Response) error {
	log.Info("Received AuthStudy.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *AuthStudy) Stream(ctx context.Context, req *authstudy.StreamingRequest, stream authstudy.AuthStudy_StreamStream) error {
	log.Infof("Received AuthStudy.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&authstudy.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *AuthStudy) PingPong(ctx context.Context, stream authstudy.AuthStudy_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&authstudy.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
