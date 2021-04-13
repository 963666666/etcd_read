package v3

import (
	"context"
)

type Cluster interface {

	MemberList(ctx context.Context)
}