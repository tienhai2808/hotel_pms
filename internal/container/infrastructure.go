package container

import (
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/middleware"
	"github.com/InstayPMS/backend/internal/infrastructure/config"
	"github.com/InstayPMS/backend/internal/infrastructure/initialization"
	"github.com/InstayPMS/backend/internal/infrastructure/persistence/orm"
	"github.com/InstayPMS/backend/internal/infrastructure/provider/jwt"
	"github.com/InstayPMS/backend/internal/infrastructure/provider/redis"
)

func (c *Container) initInfrastructure(cfg *config.Config) error {
	log, err := initialization.InitZap(cfg.Log)
	if err != nil {
		return err
	}
	c.log = log

	db, err := initialization.InitDatabase(cfg.PostgreSQL)
	if err != nil {
		return err
	}
	c.db = db

	rdb, err := initialization.InitRedis(cfg.Redis)
	if err != nil {
		return err
	}
	c.cache = rdb

	stor, err := initialization.InitMinIO(cfg.MinIO)
	if err != nil {
		return err
	}
	c.stor = stor

	idGen, err := initialization.InitSnowFlake()
	if err != nil {
		return err
	}
	c.idGen = idGen

	c.jwtPro = jwt.NewJWTProvider(cfg.JWT.SecretKey)

	c.cachePro = redis.NewCacheProvider(c.cache)

	c.userRepo = orm.NewUserRepository(c.db.Gorm)

	c.tokenRepo = orm.NewTokenRepository(c.db.Gorm)

	c.CtxMid = middleware.NewContextMiddleware(log)

	return nil
}
