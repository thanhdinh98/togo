package user

import (
	"gtodo/src/domain/user"
	"gtodo/src/schema"
)

type UserController struct {
	workflow user.IUserWorkflow
}

func (this *UserController) Login() (*schema.LoginResponse, error) {
	return nil, nil
}

func (this *UserController) Register() (*schema.RegisterResponse, error) {
	return nil, nil
}

func (this *UserController) CreateTaskByOwner() (*schema.CreateTaskByOwnerResponse, error) {
	return nil, nil
}

func (this *UserController) DeleteTaskByOwner() (*schema.DeleteTaskByOwnerResponse, error) {
	return nil, nil
}

func NewUserController() IUserController {
	return &UserController{
		user.NewUserWorkflow(),
	}
}
