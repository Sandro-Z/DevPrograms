package dao

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"git.ana/xjtuana/api-micro-mail/config"
	"git.ana/xjtuana/api-micro-mail/model"
)

type DAO struct {
	cfg *config.Config
	db  *gorm.DB
}

func New(cfg *config.Config) *DAO {
	d := &DAO{cfg: cfg}
	d.init()
	return d
}

func (d *DAO) init() {
	db, err := gorm.Open(mysql.Open(d.cfg.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to open database connection, error: " + err.Error())
	}
	d.db = db.Debug()
	d.db.Migrator().AutoMigrate(&model.AuditLog{})
	d.db.Migrator().AutoMigrate(&model.Mail{})
	d.db.Migrator().AutoMigrate(&model.Template{})
}

func (d *DAO) Close() error {
	if d.db != nil {
		db, err := d.db.DB()
		if err != nil {
			return err
		}
		return db.Close()
	}
	return nil
}

func (d *DAO) Ping(ctx context.Context) error {
	if d.db != nil {
		db, err := d.db.DB()
		if err != nil {
			return err
		}
		return db.PingContext(ctx)
	}
	return nil
}
