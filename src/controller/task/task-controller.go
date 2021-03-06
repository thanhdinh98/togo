package task

type TaskController struct {
}

func (tc *TaskController) Create(context interface{}) (interface{}, error) {
	return nil, nil
}

func NewTaskController() ITaskController {
	return &TaskController{}
}
