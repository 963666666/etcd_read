package concurrency

import v3 "etcd_read/client/v3"

type STM interface {
	Get(key ...string) string

	Put(key, val string, opts ...v3.OpOption)

	Rev(key string) int64

	Del(key string)

	commit() v3.TxnResponse
	reset()
}

func mkSTM(c *v3.Client, opts *stmOptions) STM {


	return nil
}
