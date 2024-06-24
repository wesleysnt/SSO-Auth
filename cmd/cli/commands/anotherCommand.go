package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var AnotherCommand = &cobra.Command{
	Use:   "subcommand2",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("lalalal")
		os.Exit(0)

		return nil
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mta/addCustomerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mta/addCustomerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
