package schedule

import "time"

type jobType string

const (
	All     jobType = "Garbage/Recycling/Greenbin"
	Partial jobType = "Recycling/Greenbin"
)

type job struct {
	date             time.Time
	responsibilities jobType
}
