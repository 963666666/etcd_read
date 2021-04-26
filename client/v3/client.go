package v3

import (
	"context"
	"etcd_read/client/v3/credentials"
	"etcd_read/client/v3/internal/resolver"
	"etcd_read/pkg/logutil"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	grpccredentials "google.golang.org/grpc/credentials"
	"sync"
)

type Client struct {
	Cluster
	KV
	Lease
	Watcher
	Auth
	Maintenance

	conn *grpc.ClientConn

	cfg      Config
	creds    grpccredentials.TransportCredentials
	resolver *resolver.EtcdManualResolver
	mu       *sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc

	Username        string
	Password        string
	authTokenBundle credentials.Bundle

	callOpts []grpc.CallOption

	lgMu *sync.RWMutex
	lg   *zap.Logger
}

func newClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		cfg = &Config{}
	}

	var creds grpccredentials.TransportCredentials
	if cfg.TLS != nil {
		creds = credentials.NewBundle(credentials.Config{TLSConfig: cfg.TLS}).TransportCredentials()
	}

	baseCtx := context.TODO()
	if cfg.Context != nil {
		baseCtx = cfg.Context
	}

	ctx, cancel := context.WithCancel(baseCtx)
	client := &Client{
		conn:  nil,
		cfg:   *cfg,
		creds: creds,
		ctx:   ctx,
		cancel: cancel,
		mu: new(sync.RWMutex),
		callOpts: defaultCallOpts,
		lgMu: new(sync.RWMutex),
	}

	lcfg := logutil.DefaultZapLoggerConfig
	if cfg.LogConfig != nil {
		lcfg = *cfg.LogConfig
	}
	var err error
	client.lg, err = lcfg.Build()
	if err != nil {
		return nil, err
	}

	if cfg.Username != "" && cfg.Password != "" {
		client.Username = cfg.Username
		client.Password = cfg.Password
	}
	if cfg.MaxCallSendMsgSize > 0 || cfg.MaxCallRecvMsgSize > 0 {
		if cfg.MaxCallRecvMsgSize > 0 && cfg.MaxCallRecvMsgSize > cfg.MaxCallSendMsgSize {
			return nil, fmt.Errorf("gRPC message recv limit (%d bytes) must be greater than send limit (%d bytes)", cfg.MaxCallRecvMsgSize, cfg.MaxCallSendMsgSize)
		}
		callOpts := []grpc.CallOption{
			defaultFailFast,
			defaultMaxCallSendMsgSize,
			defaultMaxCallRecvMsgSize,
		}
		if cfg.MaxCallSendMsgSize > 0 {
			callOpts[1] = grpc.MaxCallSendMsgSize(cfg.MaxCallSendMsgSize)
		}
		if cfg.MaxCallRecvMsgSize > 0 {
			callOpts[2] = grpc.MaxCallRecvMsgSize(cfg.MaxCallRecvMsgSize)
		}
		client.callOpts = callOpts
	}

	client.resolver = resolver.New(cfg.Endpoints...)

	if len(cfg.Endpoints) < 1 {
		client.cancel()
		return nil, fmt.Errorf("at least one Endpoints is required in client config")
	}

	conn, err := client.dail


	return client, nil
}
