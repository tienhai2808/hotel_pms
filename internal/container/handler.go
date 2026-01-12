package container

import "github.com/InstaySystem/is_v2-be/internal/infrastructure/api/http/handler"

func (c *Container) initHandlers() {
	c.FileHdl = handler.NewFileHandler(c.fileUC)

	c.AuthHdl = handler.NewAuthHandler(c.cfg, c.authUC)

	c.UserHdl = handler.NewUserHandler(c.userUC)

	c.DepartmentHdl = handler.NewDepartmentHandler(c.departmentUC)
}
