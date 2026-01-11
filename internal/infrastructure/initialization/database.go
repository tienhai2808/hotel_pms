package initialization

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/InstayPMS/backend/internal/domain/model"
	"github.com/InstayPMS/backend/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Gorm *gorm.DB
	sql  *sql.DB
}

func InitDatabase(cfg config.PostgreSQLConfig) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.DBName,
		cfg.User,
		cfg.Password,
		cfg.SSLMode,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "[DB] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	if err := runAutoMigrations(pg); err != nil {
		return nil, err
	}

	sql, err := pg.DB()
	if err != nil {
		return nil, err
	}

	return &Database{
		pg,
		sql,
	}, nil
}

func (d *Database) Close() {
	_ = d.sql.Close()
}

var allModels = []any{
	&model.Outlet{},
	&model.Department{},
	&model.User{},
	&model.Token{},
}

func runAutoMigrations(db *gorm.DB) error {
	oldOption := db.Config.DisableForeignKeyConstraintWhenMigrating
	db.Config.DisableForeignKeyConstraintWhenMigrating = true

	if err := db.AutoMigrate(allModels...); err != nil {
		db.Config.DisableForeignKeyConstraintWhenMigrating = oldOption
		return err
	}

	db.Config.DisableForeignKeyConstraintWhenMigrating = oldOption
	
	return db.AutoMigrate(allModels...)
}
