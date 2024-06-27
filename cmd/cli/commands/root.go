package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
<<<<<<< HEAD:cmd/cli/commands/root.go
	Use:   "sso",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
=======
	Run: func(command *cobra.Command, args []string) {
		fmt.Print("")
	},
>>>>>>> 325f9fc (.):cmd/cli/root.go
}

// Use: "wekekeke",
// 	Short: "A brief description of your application",
// 	Long: `A longer description that spans multiple lines and likely contains
// examples and usage of using your application. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// Uncomment the following line if your bare application
// has an action associated with it:
// Run: func(cmd *cobra.Command, args []string) { },

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mta-data-import.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
<<<<<<< HEAD:cmd/cli/commands/root.go
<<<<<<< HEAD
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
=======
	cobra.OnInitialize()
>>>>>>> 336b74a (add many functions into commands)
=======
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommand()
>>>>>>> 325f9fc (.):cmd/cli/root.go
}
