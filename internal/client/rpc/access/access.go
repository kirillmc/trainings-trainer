package access

import (
	"context"

	descAccess "github.com/kirillmc/trainings-auth/pkg/access_v1"
	"github.com/kirillmc/trainings-server/internal/client/rpc"
)

type accessClient struct {
	client descAccess.AccessV1Client
}

var _ rpc.AccessClient = (*accessClient)(nil)

func NewAccessClient(client descAccess.AccessV1Client) rpc.AccessClient {
	return &accessClient{
		client: client,
	}
}

func (c *accessClient) Check(ctx context.Context, endpoint string) error {
	_, err := c.client.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: endpoint,
	})

	return err
}

func (c *accessClient) CheckIsExistUser(ctx context.Context, userId int64) error {
	_, err := c.client.CheckIsExistUser(ctx, &descAccess.CheckIsExistUserRequest{
		UserId: userId,
	})

	return err
}
