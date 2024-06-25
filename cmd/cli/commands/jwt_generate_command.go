package commands

import (
	"fmt"
	"os"
	"sso-auth/app/utils"
	"strings"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var jwtGenerateCommand = &cobra.Command{
	Use:   "generate:jwt",
	Short: "Generate a jwt secret",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		configFile := fmt.Sprintf("%s/.env", pwd)

		data, err := os.ReadFile(configFile)

		if err != nil {
			color.Redf("Error read file .env: %v \n", err)
			return
		}

		// Convert file data to string
		content := string(data)

		// Find the line and replace its value
		isFound := false
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			if strings.HasPrefix(line, "JWT_SECRET=") {
				lines[i] = fmt.Sprintf("JWT_SECRET=%s", utils.Random(12))
				isFound = true
				break
			}
		}

		if !isFound {
			color.Redln("Not found JWT_SECRET into .env")
			return
		}

		// Join the lines back into a single string
		updatedContent := strings.Join(lines, "\n")

		// Write the updated content back to the file
		err = os.WriteFile(configFile, []byte(updatedContent), 0644)
		if err != nil {
			color.Redf("Error writing updated config file: %v \n", err)
			return
		}

		color.Greenln("successfully generate jwt secret")
	},
}

func init() {
	ssoCommand.AddCommand(jwtGenerateCommand)
}
