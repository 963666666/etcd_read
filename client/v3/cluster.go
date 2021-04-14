package v3

import "context"

type Cluster interface {
	MemberList(ctx context.Context)

	MemberAdd(ctx context.Context)

	MemberAddAsLearner(ctx context.Context)

	MemberRemove(ctx context.Context)

	MemberUpdate(ctx context.Context)

	MemberPromote(ctx context.Context)
}
