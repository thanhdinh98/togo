package schema

type HeaderRequest struct {
	AccessToken string `header:"access_token" validate:"required"`
}

type LoginRequest struct {
}

type RegisterRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateTaskByOwnerRequest struct {
}

type DeleteTaskByOwnerRequest struct {
}

type GetOwnerTaskRequest struct {
}
