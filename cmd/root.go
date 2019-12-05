package cmd

import (
	"fmt"
	"os"

	"github.com/kitabisa/buroq/config"
	"github.com/kitabisa/buroq/internal/app/appcontext"
	"github.com/kitabisa/buroq/internal/app/repository"
	"github.com/kitabisa/buroq/internal/app/server"
	"github.com/kitabisa/buroq/internal/app/service"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{ cookiecutter.app_name }}",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}

func start() {
	cfg := config.Config()
	logger := log.NewLogger("buroq")

	app := appcontext.NewAppContext(cfg)
	dbMysql, err := app.GetDBInstance(appcontext.DBDialectMysql)
	if err != nil {
		logrus.Fatalf("Failed to start, error connect to DB MySQL | %v", err)
		return
	}

	dbPostgre, err := app.GetDBInstance(appcontext.DBDialectPostgres)
	if err != nil {
		logrus.Fatalf("Failed to start, error connect to DB Postgre | %v", err)
		return
	}

	cache := app.GetCachePool()
	cacheConn, err := cache.Dial()
	if err != nil {
		logrus.Fatalf("Failed to start, error connect to DB Cache | %v", err)
		return
	}
	defer cacheConn.Close()

	influx, err := app.GetInfluxDBClient()
	if err != nil {
		logrus.Fatalf("Failed to start, error connect to DB Influx | %v", err)
		return
	}

	repo := wiringRepository(repository.Option{
		DbMysql:   dbMysql,
		DbPostgre: dbPostgre,
		CachePool: cache,
		Influx:    influx,
		Logger:    logger,
	})

	service := wiringService(service.Option{
		DbMysql:   dbMysql,
		DbPostgre: dbPostgre,
		CachePool: cache,
		Influx:    influx,
		Logger:    logger,
		Repo:      repo,
	})

	server := server.NewServer(cfg, service, dbMysql, dbPostgre, cache, logger)

	// run app
	server.StartApp()
}

func wiringRepository(repoOption repository.Option) *repository.Repository {
	// wiring up all your repos here
	cacheRepo := repository.NewCacheRepository(repoOption.CachePool)

	repo := repository.Repository{
		Cache: cacheRepo,
	}

	return &repo
}

func wiringService(serviceOption service.Option) *service.Service {
	// wiring up all services
	hc := service.NewHealthCheck(serviceOption)

	svc := service.Service{
		HealthCheck: hc,
	}

	return &svc
}
