package enum

type EntityStatus string

const (
	EntityStatusPending   EntityStatus = "PENDING"
	EntityStatusCompleted EntityStatus = "COMPLETED"
)
