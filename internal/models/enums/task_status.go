package enums

type TaskStatus string

const (
	StatusAvailable  TaskStatus = "available"
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"
)
