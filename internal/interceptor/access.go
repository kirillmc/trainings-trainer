package interceptor

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/kirillmc/trainings-server/internal/client/rpc"
	"github.com/kirillmc/trainings-server/pkg/training_v1"
)

type Interceptor struct {
	Client rpc.AccessClient
}

func (i *Interceptor) PolicyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	err := i.Client.Check(metadata.NewOutgoingContext(ctx, md), info.FullMethod)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func (i *Interceptor) IsUserExistInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	if info.FullMethod == "/training_v1.TrainingV1/CreateTrainingProgram" {
		var userId int64

		newReq, ok := req.(*training_v1.CreateTrainingProgramRequest)
		if ok {
			userId = newReq.Info.UserId
		}

		err := i.Client.CheckIsExistUser(metadata.NewOutgoingContext(ctx, md), userId)
		if err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}
