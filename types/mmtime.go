package types

import "time"

type Task struct {
	Name       string
	LaunchedAt time.Time
	UsedFor    time.Duration
}



