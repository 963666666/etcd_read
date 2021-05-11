package v3

import (
	"context"
	"google.golang.org/grpc"
	"sync"
)

type Txn interface {
	If(cs ...Cmp) Txn

	Then(ops ...Op) Txn

	Else(ops ...Op) Txn

	Commit() (*TxnResponse, error)
}

type txn struct {
	kv *kv
	ctx context.Context

	mu sync.Mutex
	cif bool
	cthen bool
	celse bool

	isWrite bool

	cmps []*pb.Compare

	sus []*pb.RequestOp
	fas []*pb.RequestOp

	callOpts []grpc.CallOption
}

func (txn *txn) If(cs ...Cmp) Txn {
	
}
