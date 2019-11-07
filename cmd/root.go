package cmd

import (
	"fmt"
	"os"

	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/repository"
	"github.com/kitabisa/go-bootstrap/internal/app/server"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
	"github.com/kitabisa/go-bootstrap/internal/pkg/appcontext"
	plog "github.com/kitabisa/perkakas/v2/log"
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
	logger := plog.NewLogger("go-bootstrap")

	app := appcontext.NewAppContext(cfg)
	dbMysql, err := app.GetDBInstance(appcontext.DBTypeMysql)
	if err != nil {
		logger.AddMessage(plog.FatalLevel, fmt.Sprintf("Failed to start | %v", err))
		logger.Print()
		return
	}

	dbPostgre, err := app.GetDBInstance(appcontext.DBTypePostgre)
	if err != nil {
		logger.AddMessage(plog.FatalLevel, fmt.Sprintf("Failed to start | %v", err))
		logger.Print()
		return
	}

	cache := app.GetCachePool()
	cacheConn, err := cache.Dial()
	if err != nil {
		logger.AddMessage(plog.FatalLevel, fmt.Sprintf("Failed to start | %v", err))
		logger.Print()
		return
	}
	defer cacheConn.Close()

	repo := wiringRepository(repository.Option{
		DbMysql:   dbMysql,
		DbPostgre: dbPostgre,
		CachePool: cache,
	})

	service := wiringService(service.Option{
		DbMysql:   dbMysql,
		DbPostgre: dbPostgre,
		CachePool: cache,
		Repo:      repo,
	})

	server := server.NewServer(cfg, service, dbMysql, dbPostgre, cache, logger)
	server.StartApp()
}

func wiringRepository(repoOption repository.Option) *repository.Repository {
	repo := repository.NewRepository()

	// TODO: wiring up all your repos here

	return repo
}

func wiringService(serviceOption service.Option) *service.Service {
	svc := service.NewService()

	// wiring up all services
	hc := service.NewHealthCheck(serviceOption)
	svc.HealthCheck = hc

	return svc
}
