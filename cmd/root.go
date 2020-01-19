package cmd

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/graphql-go/graphql"
	"github.com/kitabisa/buroq/config"
	"github.com/kitabisa/buroq/internal/app/appcontext"
	"github.com/kitabisa/buroq/internal/app/commons"
	"github.com/kitabisa/buroq/internal/app/graphql/mutation"
	"github.com/kitabisa/buroq/internal/app/graphql/query"
	"github.com/kitabisa/buroq/internal/app/graphql/resolver"
	"github.com/kitabisa/buroq/internal/app/repository"
	"github.com/kitabisa/buroq/internal/app/server"
	"github.com/kitabisa/buroq/internal/app/service"
	"github.com/kitabisa/perkakas/v2/log"
	"github.com/kitabisa/perkakas/v2/metrics/influx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/gorp.v2"
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
	var err error

	var dbMysql *gorp.DbMap
	if cfg.GetBool("mysql.is_enabled") {
		dbMysql, err = app.GetDBInstance(appcontext.DBDialectMysql)
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB MySQL | %v", err)
			return
		}
	}

	var dbPostgre *gorp.DbMap
	if cfg.GetBool("postgre.is_enabled") {
		dbPostgre, err = app.GetDBInstance(appcontext.DBDialectPostgres)
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Postgre | %v", err)
			return
		}
	}

	var cache *redis.Pool
	if cfg.GetBool("cache.is_enabled") {
		cache = app.GetCachePool()
		cacheConn, err := cache.Dial()
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Cache | %v", err)
			return
		}
		defer cacheConn.Close()
	}

	var influx *influx.Client
	if cfg.GetBool("influx.is_enabled") {
		influx, err = app.GetInfluxDBClient()
		if err != nil {
			logrus.Fatalf("Failed to start, error connect to DB Influx | %v", err)
			return
		}
	}

	opt := commons.Options{
		Config:    cfg,
		DbMysql:   dbMysql,
		DbPostgre: dbPostgre,
		CachePool: cache,
		Influx:    influx,
		Logger:    logger,
	}

	repo := wiringRepository(repository.Option{
		Options: opt,
	})

	service := wiringService(service.Option{
		Options:    opt,
		Repository: repo,
	})

	// Init graphql: load all schema and connect to services
	resolver.Init(
		resolver.WithServices(service),
	)
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(query.GetQuerySchema()),
		Mutation: graphql.NewObject(mutation.GetMutationSchema()),
	})
	if err != nil {
		logrus.Fatalf("Failed to create schema for the graphql: %s", err)
		return
	}

	server := server.NewServer(opt, service)

	// run app
	server.StartApp()
}

func wiringRepository(repoOption repository.Option) *repository.Repository {
	// wiring up all your repos here
	cacheRepo := repository.NewCacheRepository(repoOption)

	repo := repository.Repository{
		Cache: cacheRepo,
	}

	return &repo
}

func wiringService(serviceOption service.Option) *service.Services {
	// wiring up all services
	hc := service.NewHealthCheck(serviceOption)

	svc := service.Services{
		HealthCheck: hc,
	}

	return &svc
}
