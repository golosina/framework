package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refreshes the cmd with new commands",
	Long:  `Run this after you add a new command to the cmd package.`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := exec.Command("bash", "-c", "go build -o golosina").CombinedOutput()
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}
