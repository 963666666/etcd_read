package v3

import "context"

type Watcher interface {
	Watch(ctx context.Context, key string)

	RequestProgress(ctx context.Context) error

	Close() error
}