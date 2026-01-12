package container

import (
	"log"

	"github.com/InstaySystem/is_v2-be/internal/application/port"
	authUC "github.com/InstaySystem/is_v2-be/internal/application/usecase/auth"
	departmentUC "github.com/InstaySystem/is_v2-be/internal/application/usecase/department"
	fileUC "github.com/InstaySystem/is_v2-be/internal/application/usecase/file"
	userUC "github.com/InstaySystem/is_v2-be/internal/application/usecase/user"
	"github.com/InstaySystem/is_v2-be/internal/domain/repository"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/api/http/handler"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/api/http/middleware"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/config"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/initialization"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake/v2"
	"go.uber.org/zap"
)

type Container struct {
	cfg            *config.Config
	Log            *zap.Logger
	db             *initialization.Database
	cache          *redis.Client
	mq             *initialization.MQ
	stor           *minio.Client
	idGen          *sonyflake.Sonyflake
	jwtPro         port.JWTProvider
	MQPro          port.MessageQueueProvider
	cachePro       port.CacheProvider
	SMTPPro        port.SMTPProvider
	userRepo       repository.UserRepository
	tokenRepo      repository.TokenRepository
	departmentRepo repository.DepartmentRepository
	fileUC         fileUC.FileUseCase
	authUC         authUC.AuthUseCase
	userUC         userUC.UserUseCase
	departmentUC   departmentUC.DepartmentUseCase
	FileHdl        *handler.FileHandler
	AuthHdl        *handler.AuthHandler
	UserHdl        *handler.UserHandler
	DepartmentHdl  *handler.DepartmentHandler
	CtxMid         *middleware.ContextMiddleware
	AuthMid        *middleware.AuthMiddleware
}

func NewContainer(cfg *config.Config) (*Container, error) {
	c := &Container{
		cfg: cfg,
	}

	if err := c.initInfrastructure(); err != nil {
		return nil, err
	}

	c.initUseCases()

	c.initHandlers()

	return c, nil
}

func (c *Container) Cleanup() {
	if c.db != nil {
		c.db.Close()
	}
	if c.mq != nil {
		c.mq.Close()
	}

	log.Println("Container cleaned successfully")
}
