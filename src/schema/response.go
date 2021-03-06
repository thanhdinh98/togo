package schema

type LoginResponse struct {
}

type RegisterResponse struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type CreateTaskByOwnerResponse struct {
}

type DeleteTaskByOwnerResponse struct {
}

type GetOwnerTaskResponse struct {
}
