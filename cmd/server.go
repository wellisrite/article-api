package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"article-api/delivery/rest"
	vldtr "article-api/libs/validator"
	"article-api/repository/psql"
	"article-api/service"

	dbase "article-api/libs/technologies"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var ProjectFolder = flag.String("folder", "./", "absolute path of project folder")

// GitCommit : for show commit number
var GitCommit string

func serv() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	e := echo.New()
	cfg := loadConfig()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Success Run Article API Version %v %v !!!", os.Getenv("ARTICLE_API_APP_VERSION"), os.Getenv("ARTICLE_API_APP_ENV")))
	})

	e.Validator = &vldtr.CustomValidator{Validator: validator.New()}
	log.Println("[" + appName + "] Successfully initializing validator ...")

	db, pg, err := dbase.InitDB(cfg.DB)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// auto migrate DB
	autoMigrateDB(db)

	dbV2, err := gorm.Open(postgres.New(postgres.Config{Conn: pg}))
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("[" + appName + "] Database successfully initialized ...")

	redis := dbase.InitRedis(cfg.Redis)
	log.Println("[" + appName + "] Redis successfully initialized ...")

	psqlRepos := psql.NewPsqlRepositories(db, dbV2)

	svc := service.New(e.Validator, psqlRepos, redis)
	healthRoute := e.Group("") // non api-key required
	rest.New(svc, cfg.App.Env).Route(healthRoute)

	log.Println("["+appName+"] (build commit hash: "+GitCommit+") running on port:", cfg.Webserver.ListenAddress)
	if cfg.Webserver.ListenAddress != "" {
		e.HidePort = true
		e.HideBanner = true
	}

	s := &http.Server{
		Addr:         cfg.Webserver.ListenAddress,
		ReadTimeout:  cfg.Webserver.ReadTimeout,
		WriteTimeout: cfg.Webserver.WriteTimeout,
	}

	go func() {
		sigchan := make(chan os.Signal, 1)

		signal.Notify(sigchan, os.Interrupt)

		<-sigchan

		db.Close()
		fmt.Println("database has been closed")

		redis.Close()
		fmt.Println("redis has been closed")

		os.Exit(0)
	}()

	err = e.StartServer(s)
	if err != nil {
		e.Logger.Fatal(err)
	}
}
