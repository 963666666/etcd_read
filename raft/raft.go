package raft

type Config struct {
	ID uint64
}

type raft struct {
	id uint64

	Term uint64
	Vote uint64
}
