package embed

type Config struct {
	Name   string `json:"name"`
	Dir    string `json:"dir"`
	WalDir string `json:"wal-dir"`

	SnapshotCount uint64 `json:"snapshot-count"`

	SnapShotCatchUpEntries uint64
}
