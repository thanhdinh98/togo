package task

import (
	"gtodo/src"
	"gtodo/src/schema"
)

type ITaskWorkflow interface {
	AddTaskByOwner(context src.IContextService, data *schema.AddTaskRequest) (*schema.AddTaskResponse, error)
}
