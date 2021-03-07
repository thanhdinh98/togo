package task

import (
	"gtodo/src"
	"gtodo/src/schema"
)

type ITaskController interface {
	AddTaskByOwner(context src.IContextService, data *schema.AddTaskRequest) (*schema.AddTaskResponse, error)
}
