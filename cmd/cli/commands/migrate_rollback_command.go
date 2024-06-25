package commands

import (
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var migrateRollbackCmd = &cobra.Command{
	Use:   "migrate:rollback [step]",
	Short: "Rollback the last migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()

		if err != nil {
			color.Yellowln("Please fill database config first")
			return
		}

		step := "-" + args[0]

		convStep, err := strconv.Atoi(step)

		if err != nil {
			color.Redf("error step: %v \n", err)
			return
		}

		err = m.Steps(convStep)
		if err != nil && err != migrate.ErrNoChange {
			color.Redf("error process step: %v \n", err)
			return
		}

		color.Greenln("Rolled back the last migration")
	},
}

func init() {
	ssoCommand.AddCommand(migrateRollbackCmd)
}
