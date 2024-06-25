package commands

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var migrateRefreshCmd = &cobra.Command{
	Use:   "migrate:refresh",
	Short: "Rollback all migrations and re-run them",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()

		if err != nil {
			fmt.Println(err)
			return
		}

		err = m.Drop()
		if err != nil {
			color.Redf("error drop: %v \n", err)
			return
		}

		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			color.Redf("error up: %v \n", err)
			return
		}

		color.Greenln("Database refreshed")
	},
}

func init() {
	ssoCommand.AddCommand(migrateRefreshCmd)
}
