package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gencli to the latest version",
	Long:  `This command will help you to update gencli to the latest version.`,
	Run: func(cmd *cobra.Command, args []string) {
		update()
	},
}

func update() {
	cmd := exec.Command("go", "install", "github.com/Pradumnasaraf/gencli@latest")
	_, err := cmd.Output()

	CheckNilError(err)

	fmt.Printf("CLI updated successfully to the latest version (If any).\n")
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
