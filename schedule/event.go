package schedule

type EventType string

const (
	All     EventType = "Garbage/Recycling/Greenbin"
	Partial EventType = "Recycling/Greenbin"
)

type event struct {
	date             string
	responsibilities EventType
}
