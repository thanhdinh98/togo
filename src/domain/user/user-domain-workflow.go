package user

import (
	"gtodo/src/entity/user"
	useRepo "gtodo/src/infra/repository/user"
	"gtodo/src/schema"
)

type UserWorkflow struct {
	repository user.IUserRepository
}

func (this *UserWorkflow) Login(data *schema.LoginRequest) (*schema.LoginResponse, error) {
	return nil, nil
}

func (this *UserWorkflow) Register(data *schema.RegisterRequest) (*schema.RegisterResponse, error) {
	return nil, nil
}

func (this *UserWorkflow) CreateTaskByOwner(data *schema.CreateTaskByOwnerRequest) (*schema.CreateTaskByOwnerResponse, error) {
	return nil, nil
}

func (this *UserWorkflow) DeleteTaskByOwner(data *schema.DeleteTaskByOwnerRequest) (*schema.DeleteTaskByOwnerResponse, error) {
	return nil, nil
}

func NewUserWorkflow() IUserWorkflow {
	return &UserWorkflow{
		useRepo.NewUserRepository(),
	}
}
