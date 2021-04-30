package etcdmain

import "etcd_read/server/embed"

type config struct {
	ec embed.Config
	configFile string
	printVersion bool
}

func newConfig() *config {
	return nil
}