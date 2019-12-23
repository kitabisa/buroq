package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/kitabisa/buroq/config"
	"github.com/kitabisa/buroq/internal/app/appcontext"
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
		logger := log.NewLogger("buroq-migrate")
		mSource := getMigrateSource()

		if c.GetBool("mysql.is_enable") && c.GetBool("mysql.is_migration_enable") {
			doMigrate(appCtx, logger, mSource, appcontext.DBDialectMysql, migrate.Up)
		}

		if c.GetBool("postgre.is_enable") && c.GetBool("postgre.is_migration_enable") {
			doMigrate(appCtx, logger, mSource, appcontext.DBDialectPostgres, migrate.Up)
		}
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migratedown",
	Short: "Migrate Up DB Sekawan",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config()
		appCtx := appcontext.NewAppContext(c)
		logger := log.NewLogger("buroq-migrate")
		mSource := getMigrateSource()

		if c.GetBool("mysql.is_enable") && c.GetBool("mysql.is_migration_enable") {
			doMigrate(appCtx, logger, mSource, appcontext.DBDialectMysql, migrate.Down)
		}

		if c.GetBool("postgre.is_enable") && c.GetBool("postgre.is_migration_enable") {
			doMigrate(appCtx, logger, mSource, appcontext.DBDialectPostgres, migrate.Down)
		}
	},
}

var migrateNewCmd = &cobra.Command{
	Use:   "migratenew [migration name]",
	Short: "Create new migration file",
	Long:  `Create new migration file on folder migrations/sql with timestamp as prefix`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.NewLogger("buroq-migrate")
		mDir := "migrations/sql/"

		createMigrationFile(logger, mDir, args[0])
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrateNewCmd)
}

func getMigrateSource() migrate.FileMigrationSource {
	source := migrate.FileMigrationSource{
		Dir: "migrations/sql",
	}

	return source
}

func doMigrate(appCtx *appcontext.AppContext, logger *log.Logger, mSource migrate.FileMigrationSource, dbDialect string, direction migrate.MigrationDirection) error {
	db, err := appCtx.GetDBInstance(dbDialect)
	if err != nil {
		logger.AddMessage(log.FatalLevel, fmt.Sprintf("Error connection to DB | %v", err)).Print()
		return err
	}

	defer db.Db.Close()

	total, err := migrate.Exec(db.Db, dbDialect, mSource, direction)
	if err != nil {
		logger.AddMessage(log.ErrorLevel, fmt.Sprintf("Fail migration | %v", err)).Print()
		return err
	}

	logger.AddMessage(log.InfoLevel, fmt.Sprintf("Migrate Success, total migrated: %d", total)).Print()
	return nil
}

func createMigrationFile(logger *log.Logger, mDir string, mName string) error {
	var migrationContent = `-- +migrate Up
 		-- SQL in section 'Up' is executed when this migration is applied
 		-- [your SQL script here]
		 
		 -- +migrate Down
 		-- SQL section 'Down' is executed when this migration is rolled back
 		-- [your SQL script here]
 	`
	filename := fmt.Sprintf("%d_%s.sql", time.Now().Unix(), mName)
	filepath := fmt.Sprintf("%s%s", mDir, filename)

	f, err := os.Create(filepath)
	if err != nil {
		logger.AddMessage(log.ErrorLevel, fmt.Sprintf("Error create migration file | %v", err)).Print()
		return err
	}
	defer f.Close()

	f.WriteString(migrationContent)
	f.Sync()

	logger.AddMessage(log.InfoLevel, fmt.Sprintf("New migration file has been created: %s)", filepath)).Print()
	return nil
}
