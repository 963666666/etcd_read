package v3

import (
	"context"
	"crypto/tls"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

type Config struct {
	Endpoints []string `json:"endpoints"`

	AutoSyncInterval time.Duration `json:"auto-sync-interval"`

	DialTimeout time.Duration `json:"dial-timeout"`

	DialKeepAliveTime time.Duration `json:"dial-keep-alive-time"`

	MaxCallSendMsgSize int

	MaxCallRecvMsgSize int

	TLS *tls.Config

	Username string `json:"username"`

	Password string `json:"password"`

	RejectOldCluster bool `json:"reject-old-cluster"`

	DialOptions []grpc.DialOption

	Context context.Context

	LogConfig *zap.Config

	PermitWithoutStream bool `json:"permit-without-stream"`
}
