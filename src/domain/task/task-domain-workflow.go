package task

import (
	"gtodo/src"
	"gtodo/src/entity/task"
	"gtodo/src/entity/user"
	"gtodo/src/schema"
	"time"

	"github.com/google/uuid"

	taskRepo "gtodo/src/infra/repository/task"
	userRepo "gtodo/src/infra/repository/user"
)

type TaskWorkflow struct {
	repository     task.ITaskRepository
	userRepository user.IUserRepository
}

func (this *TaskWorkflow) AddTaskByOwner(context src.IContextService, data *schema.AddTaskRequest) (*schema.AddTaskResponse, error) {
	tokenData := context.GetTokenData()

	if _, err := this.userRepository.FindOne(user.User{ID: tokenData.UserId}); err != nil {
		return nil, err
	}

	task := &task.Task{
		Id:          uuid.NewString(),
		Content:     data.Content,
		UserId:      tokenData.UserId,
		CreatedDate: time.Now(),
	}

	createdTask, err := this.repository.Create(task)
	if err != nil {
		return nil, err
	}

	return &schema.AddTaskResponse{
		TaskId: createdTask.Id,
	}, nil
}

func NewTaskWorkflow() ITaskWorkflow {
	return &TaskWorkflow{
		taskRepo.NewTaskRepository(),
		userRepo.NewUserRepository(),
	}
}
