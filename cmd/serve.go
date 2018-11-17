package cmd

import (
	"github.com/eaperezc/golosina/framework"
	"github.com/eaperezc/golosina/routes"
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server",
	Long:  `This is how you start listening for http requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := framework.New()

		// Prepare routes
		routes.WebRoutes(app.Router)
		routes.APIRoutes(app.Router)

		app.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
