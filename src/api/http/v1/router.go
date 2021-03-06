package v1

import (
	"fmt"
	"gtodo/src/schema"

	"net/http"

	"github.com/labstack/echo/v4"
)

type RouterV1 struct {
	framework *echo.Echo
	// taskController task.ITaskController
	// userController user.IUserController
}

func (r1 *RouterV1) LoadAPI() {
	group := r1.framework.Group("/api")

	group.POST("/register", func(c echo.Context) error {
		request := new(schema.RegisterRequest)

		if err := c.Bind(request); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Printf("%v/n", request)

		// data, err := r1.userController.Register()
		// if err != nil {
		// 	r1.framework.Logger.Fatal(err)
		// }

		return c.JSON(http.StatusOK, "Hello World")
	})

}

func NewRouterV1(framework *echo.Echo) *RouterV1 {
	return &RouterV1{
		framework: framework,
		// taskController: task.NewTaskController(),
		// userController: user.NewUserController(),
	}
}
