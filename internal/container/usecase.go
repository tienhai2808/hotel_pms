package container

import (
	authUC "github.com/InstayPMS/backend/internal/application/usecase/auth"
	fileUC "github.com/InstayPMS/backend/internal/application/usecase/file"
)

func (c *Container) initUseCases() {
	c.fileUC = fileUC.NewFileUseCase(c.cfg, c.stor, c.log)
	c.authUC = authUC.NewAuthUseCase(c.cfg.JWT, c.log, c.idGen, c.jwtPro, c.cachePro, c.userRepo, c.tokenRepo)
}
