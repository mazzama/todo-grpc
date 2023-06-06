package constant

import "fmt"

type ItemStatus string

const (
	ItemStatusTodo       ItemStatus = "TODO"
	ItemStatusInProgress            = "IN_PROGRESS"
	ItemStatusDone                  = "DONE"
	ItemStatusInvalid               = "INVALID"
)

func (i ItemStatus) ToString() string {
	return fmt.Sprintf("%v", i)
}

var MapItemStatus = map[string]ItemStatus{
	"TODO":        ItemStatusTodo,
	"IN_PROGRESS": ItemStatusInProgress,
	"DONE":        ItemStatusDone,
	"INVALID":     ItemStatusInvalid,
}
