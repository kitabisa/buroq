package cmd

import (
	"fmt"

	"github.com/kitabisa/go-bootstrap/config"
	"github.com/kitabisa/go-bootstrap/internal/pkg/appcontext"
	"github.com/kitabisa/perkakas/v2/log"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Up DB Sekawan",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config()
		appCtx := appcontext.NewAppContext(c)
		logger := log.NewLogger("go-bootstrap-migrate")
		mSource := getMigrateSource()

		doMigrate(appCtx, logger, mSource, appcontext.DBDialectMysql, migrate.Up)
		doMigrate(appCtx, logger, mSource, appcontext.DBDialectPostgres, migrate.Up)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migratedown",
	Short: "Migrate Up DB Sekawan",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config()
		appCtx := appcontext.NewAppContext(c)
		logger := log.NewLogger("go-bootstrap-migrate")
		mSource := getMigrateSource()

		doMigrate(appCtx, logger, mSource, appcontext.DBDialectMysql, migrate.Down)
		doMigrate(appCtx, logger, mSource, appcontext.DBDialectPostgres, migrate.Down)
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
}

func getMigrateSource() migrate.FileMigrationSource {
	source := migrate.FileMigrationSource{
		Dir: "migrations",
	}

	return source
}

func doMigrate(appCtx *appcontext.AppContext, logger *log.Logger, mSource migrate.FileMigrationSource, dbDialect string, direction migrate.MigrationDirection) error {
	db, err := appCtx.GetDBInstance(dbDialect)
	if err != nil {
		logger.AddMessage(log.FatalLevel, fmt.Sprintf("Error connection to DB | %v", err))
		logger.Print()
		return err
	}

	defer db.Db.Close()

	total, err := migrate.Exec(db.Db, dbDialect, mSource, direction)
	if err != nil {
		logger.AddMessage(log.ErrorLevel, fmt.Sprintf("Fail migration | %v", err))
		logger.Print()
		return err
	}

	logger.AddMessage(log.InfoLevel, fmt.Sprintf("Migrate Success, total migrated: %d", total))
	logger.Print()
	return nil
}
