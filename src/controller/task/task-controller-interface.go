package task

type ITaskController interface {
	Create(contect interface{}) (interface{}, error)
}
