package v3

import (
	"context"
)

type LeaseID int64

type Lease interface {
	Grant(ctx context.Context, ttl int64)

	Revoke(ctx context.Context, id LeaseID)

	TimeToLive(ctx context.Context, id LeaseID)

	Leases(ctx context.Context)

	KeepAlive(ctx context.Context, id LeaseID)

	KeepAliveOnce(ctx context.Context, id LeaseID)

	Close() error
}
