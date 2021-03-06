package user

import "gtodo/src/schema"

type IUserWorkflow interface {
	Login(data *schema.LoginRequest) (*schema.LoginResponse, error)
	Register(data *schema.RegisterRequest) (*schema.RegisterResponse, error)
	CreateTaskByOwner(data *schema.CreateTaskByOwnerRequest) (*schema.CreateTaskByOwnerResponse, error)
	DeleteTaskByOwner(data *schema.DeleteTaskByOwnerRequest) (*schema.DeleteTaskByOwnerResponse, error)
}
