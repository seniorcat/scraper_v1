package cmd

import (
	"database/sql"
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

var (
	migrationsPath string
	name           string

	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "migrate database",
	}

	createArgsValidator = func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("migrate create need one arg - migration name (string)")
		}

		if args[0] == "" {
			return fmt.Errorf("migration name can't be empty")
		}

		return nil
	}

	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "apply all migrations",
		Long:  "Migrate the DB to the most recent version available",
		RunE:  migrateUpCmdHandler,
	}

	migrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "rollback all migrations",
		Long:  "Roll back all migrations",
		RunE:  migrateDownCmdHandler,
	}

	migrateDownByOneCmd = &cobra.Command{
		Use:   "down-by-one",
		Short: "rollback one transaction",
		Long:  "Roll back the version by 1",
		RunE:  migrateDownByOneCmdHandler,
	}

	migrateCreateCmd = &cobra.Command{
		Use:   "create [migration_name]",
		Short: "create migration",
		Long:  "Creates new migration file with the current timestamp",
		Args:  createArgsValidator,
		RunE:  migrateCreateCmdHandler,
	}
)

// Command init function.
func init() {
	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
	migrateCmd.AddCommand(migrateDownByOneCmd)
	migrateCmd.AddCommand(migrateCreateCmd)
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.PersistentFlags().StringVarP(&migrationsPath, "migrationsPath", "m", "migrations", `path to migration files`)
	migrateCmd.PersistentFlags().StringVarP(&name, "name", "n", "migrate1", `name of migration file`)
}

func migrateUpCmdHandler(*cobra.Command, []string) (err error) {
	var db *sql.DB
	var embedMigrations embed.FS
	// setup database

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, migrationsPath); err != nil {
		panic(err)
	}
	return nil
}

func migrateDownCmdHandler(*cobra.Command, []string) (err error) {
	var db *sql.DB
	var embedMigrations embed.FS
	// setup database

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, migrationsPath); err != nil {
		panic(err)
	}
	return nil
}

func migrateDownByOneCmdHandler(*cobra.Command, []string) (err error) {
	var db *sql.DB
	var embedMigrations embed.FS
	// setup database

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, migrationsPath); err != nil {
		panic(err)
	}
	return nil
}

func migrateCreateCmdHandler(_ *cobra.Command, args []string) error {
	if args[0] == "" {
		return fmt.Errorf("no name provided")
	}

	var err error
	if err = os.MkdirAll(migrationsPath, 0755); err != nil {
		return err
	}

	var sqlMigrationTemplate = template.Must(
		template.New("goose.sql-migration").
			Parse(
				`-- +goose Up 
-- SQL in this section is executed when the migration is applied. 
	  
-- +goose Down 
-- SQL in this section is executed when the migration is rolled back. 
`))

	if err = goose.CreateWithTemplate(nil, migrationsPath, sqlMigrationTemplate, args[0], "sql"); err != nil {
		return err
	}
	return nil
}
