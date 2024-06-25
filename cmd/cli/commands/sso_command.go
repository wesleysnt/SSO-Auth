package commands

import (
	"github.com/spf13/cobra"
)

var ssoCommand = &cobra.Command{
	Use:   "sso",
	Short: "sso related commands",
}

func init() {
	rootCmd.AddCommand(ssoCommand)
}
