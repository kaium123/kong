package conn

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/api_gaurd/pkg/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Postgres driver for sql
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB holds the database instace
type DB struct{ *gorm.DB }

// Ping tests if db connection is alive
func (db DB) Ping() error {
	d, err := db.DB.DB() // returns *sql.DB
	if err != nil {
		return err
	}
	return d.Ping()
}

// defaultConfig is the default database instance
var defaultDB DB

// ConnectDB establish database connection
func ConnectDB(cfg *config.Database) error {
	host := cfg.Host
	if cfg.Port != 0 {
		host = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	}
	uri := url.URL{
		Scheme: "postgres",
		Host:   host,
		Path:   cfg.Name,
		User:   url.UserPassword(cfg.Username, cfg.Password),
	}
	if cfg.Options != nil {
		val := url.Values(cfg.Options)
		uri.RawQuery = val.Encode()
	}

	sqlDB, err := sql.Open("postgres", uri.String())
	if err != nil {
		return err
	}

	// set connection pool information
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConn)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	sqlDB.SetConnMaxLifetime(cfg.MaxConnLifetime)
	sqlDB.SetConnMaxIdleTime(cfg.MaxIdleConnLifetime)

	sqlLogger := logger.New(
		//log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logrus.New(),
		logger.Config{
			SlowThreshold:             time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: false,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,            // Disable color
		},
	)

	gormCfg := &gorm.Config{}

	gormCfg.Logger = sqlLogger
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), gormCfg)

	// assign connection instance
	defaultDB.DB = gormDB

	return err
}

// DefaultDB returns default db
func DefaultDB() *DB {
	return &defaultDB
}

// ConnectDefaultDB connect with default configurations
func ConnectDefaultDB() error {
	cfg := config.DB()
	err := ConnectDB(&cfg)
	// run a background process to ping and establish connection
	go func() {
		for {
			if err := defaultDB.Ping(); err != nil {
				fmt.Println("db: ping error:", err)
				if err := ConnectDB(&cfg); err != nil {
					fmt.Println("db:failed to reconnect:", err)
				}
			}
			time.Sleep(cfg.PingInterval)
		}
	}()

	return err
}
