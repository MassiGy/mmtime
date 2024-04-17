package types

import "time"

type Task struct {
	Name       string
	Running    bool
	LaunchedAt time.Time
	UsedFor    time.Duration
}
