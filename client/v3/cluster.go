package v3

import "context"

type Cluster interface {
	MemberList(ctx context.Context)

	MemberAdd(ctx context.Context)

	MemberAddAsLearner(ctx context.Context)
}
