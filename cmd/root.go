package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "golosina",
	Short: "A brief description of your command",
	Long: `
Welcome to the Golosina Cmd!
	`,
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
