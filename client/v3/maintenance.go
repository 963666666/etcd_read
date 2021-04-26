package v3

import "context"

type Maintenance interface {
	AlarmList(ctx context.Context)

	AlarmDisarm(ctx context.Context)

	Defragment(ctx context.Context)

	Status(ctx context.Context)

	HashKV(ctx context.Context)

	Snapshot(ctx context.Context)

	MoveLeader(ctx context.Context)
}
