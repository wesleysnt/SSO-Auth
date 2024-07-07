package commands

import (
	"fmt"
	"os"
	"sso-auth/cmd/cli/commands/stubs"
	"time"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var makeMigrationCmd = &cobra.Command{
	Use:   "make:migration <file_name>",
	Short: "Create Migration File",
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]

		if fileName == "" {
			cmd.Help()
			return
		}
		time := time.Now().Format("20060102150405")

		CreateUp(fileName, time)
		CreateDown(fileName, time)

		color.Greenln("Migration file created successfully")
	},
}

func init() {
	ssoCommand.AddCommand(makeMigrationCmd)
}

func CreateUp(fileName, time string) {

	fileName1 := fmt.Sprintf("app/database/migrations/%v_%v.up.sql", time, fileName)

	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.WriteString(stubs.PostgresqlStubs{}.CreateUp())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateDown(fileName, time string) {

	fileName1 := fmt.Sprintf("app/database/migrations/%v_%v.down.sql", time, fileName)

	f, err := os.Create(fileName1)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.WriteString(stubs.PostgresqlStubs{}.CreateDown())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
