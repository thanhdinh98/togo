package task

type TaskWorkflow struct {
}

func NewTaskWorkflow() ITaskWorkflow {
	return &TaskWorkflow{}
}
