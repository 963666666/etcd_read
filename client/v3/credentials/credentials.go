package credentials

import (
	"crypto/tls"
	grpccredentials "google.golang.org/grpc/credentials"
)

type Config struct {
	TLSConfig *tls.Config
}

type Bundle interface {
	grpccredentials.Bundle
	UpdateAuthToken(token string)
}

func NewBundle(cfg Config) Bundle {
	;return nil
}
