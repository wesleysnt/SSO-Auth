package commands

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var migrateResetCmd = &cobra.Command{
	Use:   "migrate:reset",
	Short: "Rollback all migrations",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()

		if err != nil {
			fmt.Println(err)
			return
		}

		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			color.Redf("error down: %v \n", err)
			return
		}

		color.Greenln("Database reset")
	},
}

func init() {
	ssoCommand.AddCommand(migrateResetCmd)
}
