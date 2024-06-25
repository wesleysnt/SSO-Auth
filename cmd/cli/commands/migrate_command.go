package commands

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var migrateCommand = &cobra.Command{
	Use:   "db:migrate",
	Short: "Run migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()

		if err != nil {
			color.Yellowln(err)
			return
		}

		if m == nil {
			color.Yellowln("Please fill database config first")
			return
		}

		if err = m.Up(); err != nil && err != migrate.ErrNoChange {
			color.Redln("Migration failed:", err.Error())
			return
		}

		color.Greenln("Migration success")
	},
}

func init() {
	ssoCommand.AddCommand(migrateCommand)
}
