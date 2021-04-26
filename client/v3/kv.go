package v3

import "context"

type KV interface {
	Put(ctx context.Context, key, val string)

	Get(ctx context.Context, key string)

	Delete(ctx context.Context, key string)

	Compact(ctx context.Context, rev int64)

	Do(ctx context.Context)

	Txn(ctx context.Context)
}
