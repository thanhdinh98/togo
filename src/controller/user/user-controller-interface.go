package user

import "gtodo/src/schema"

type IUserController interface {
	Login() (*schema.LoginResponse, error)
	Register() (*schema.RegisterResponse, error)
	CreateTaskByOwner() (*schema.CreateTaskByOwnerResponse, error)
	DeleteTaskByOwner() (*schema.DeleteTaskByOwnerResponse, error)
}
