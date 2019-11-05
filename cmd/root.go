package cmd

import (
	"fmt"
	"os"

	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/app/repository"
	"github.com/kitabisa/go-bootstrap/internal/app/server"
	"github.com/kitabisa/go-bootstrap/internal/app/service"
	"github.com/kitabisa/go-bootstrap/internal/pkg/appcontext"
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
	// TODO:
	cfg := config.Config()

	app := appcontext.NewAppContext(cfg)
	dbMysql := app.GetDBInstance(appcontext.DBTypeMysql)
	dbPostgre := app.GetDBInstance(appcontext.DBTypePostgre)
	cache := app.GetCachePool()

	repo := wiringRepository(repository.RepositoryOption{
		DbMysql:   dbMysql,
		DbPostgre: dbPostgre,
		CachePool: cache,
	})

	service := wiringService(service.ServiceOption{
		Repo:      repo,
		CachePool: cache,
	})

	server := server.NewServer(cfg, service)
	server.StartApp()
}

func wiringRepository(repoOption repository.RepositoryOption) *repository.Repository {
	repo := repository.NewRepository()

	// TODO: wiring up all your repos here

	return repo
}

func wiringService(serviceOption service.ServiceOption) *service.Service {
	service := service.NewService()

	// TODO: wiring up all your services here

	return service
}
