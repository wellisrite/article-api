package main

import (
	"log"
	"time"

	"article-api/domain/article"
	db "article-api/libs/technologies"

	"github.com/jinzhu/gorm"
	godotenv "github.com/joho/godotenv"
	envconfig "github.com/kelseyhightower/envconfig"

	_ "github.com/lib/pq"
	// _ "github.com/go-sql-driver/mysql"
)

const (
	appName    = "ARTICLE_API"
	appVersion = "0.0.0.alpha1"
	calldepth  = 4
)

type (
	config struct {
		DB        db.DBConfig    `envconfig:"DB"`
		Webserver wb             `envconfig:"WEBSERVER"`
		Redis     db.RedisConfig `envconfig:"REDIS"`
		App       appConfig      `envconfig:"APP"`
	}

	wb struct {
		ListenAddress   string        `envconfig:"LISTEN_ADDRESS"`
		MaxConnections  int           `envconfig:"MAX_CONNECTION"`
		ReadTimeout     time.Duration `envconfig:"READ_TIMEOUT"`
		WriteTimeout    time.Duration `envconfig:"WRITE_TIMEOUT"`
		GracefulTimeout time.Duration `envconfig:"GRACEFUL_TIMEOUT"`
	}
	appConfig struct {
		Env     string `envconfig:"ENV"`
		Version string `envconfig:"VERSION"`
		Secret  string `envconfig:"SECRET"`
	}
)

func load(filenames ...string) error {
	return godotenv.Load(filenames...)
}

// Parse parses environment variables.
func parse(prefix string, out interface{}) error {
	return envconfig.Process(prefix, out)
}

// LoadAndParse loads and parses environment variables.
func LoadAndParse(prefix string, out interface{}, filenames ...string) error {
	if err := load(filenames...); err != nil {
		return err
	}

	if err := parse(prefix, out); err != nil {
		return err
	}

	return nil
}

func loadConfig() *config {
	var cfg config

	err := LoadAndParse(appName, &cfg)
	if err != nil {
		panic("Failed to load environment configuration. err: " + err.Error())
	}
	log.Println("[" + appName + "] env: " + cfg.App.Env + " successfully loaded")

	return &cfg
}

/*
Function to auto migrate DB
*/
func autoMigrateDB(db *gorm.DB) {
	db.AutoMigrate(&article.Article{})
}
