package rpc

import "context"

type AccessClient interface {
	Check(ctx context.Context, endpoint string) error
	CheckIsExistUser(ctx context.Context, userId int64) error
}
