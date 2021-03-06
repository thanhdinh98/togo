package src

type IWebServerSetup interface {
	LoadRouterV1() IWebServer
}

type IWebServer interface {
	Start()
}

type IContextService interface {
}

type IErrorFactory interface {
	UnauthorizedError(errorCode string, data interface{}) error
	NotFoundError(errorCode string, data interface{}) error
	InternalServerError(errorCode string, data interface{}) error
	ForbiddenError(errorCode string, data interface{}) error
	BadRequestError(errorCode string, data interface{}) error
}

const (
	ENTITY_NOT_EXISTS_ERROR = "ENTITY_NOT_EXISTS_ERROR"
	MAPPER_NOT_EXSITS_ERROR = "MAPPER_NOT_EXSITS_ERROR"

	CREATE_TASK_ERROR   = "CREATE_TASK_ERROR"
	FIND_ONE_TASK_ERROR = "FIND_ONE_TASK_ERROR"
	FIND_TASK_ERROR     = "FIND_TASK_ERROR"

	CREATE_USER_ERROR   = "CREATE_USER_ERROR"
	FIND_USER_ERROR     = "FIND_USER_ERORR"
	FIND_ONE_USER_ERROR = "FIND_ONE_USER_ERROR"
)
