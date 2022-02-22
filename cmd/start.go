/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aeekayy/descartes/pkg/config"
	"github.com/aeekayy/descartes/pkg/http"
	"github.com/aeekayy/descartes/pkg/cron"
)

var (
	serverPort string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a http server",
	Long: `Starts a http server. Supply the port 
that will run the web server.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")

		// start the cron job
		cronConfig, err := config.NewCronConfig()
		if err != nil {
			fmt.Printf("error starting the cron: %s", err)
			return
		}
		cronManager := cron.New(cronConfig)
		cronManager.Start()

		appConfig, err := config.NewAppConfig(serverPort)
		if err != nil {
			fmt.Printf("error starting the server: %s", err)
			return
		}
		app, err := http.NewServer(appConfig)
		port := fmt.Sprintf(":%s", serverPort)
		app.Listen(port)
	},
}

func init() {
	serverCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().StringVarP(&serverPort, "port", "p", "8080", "The port where the server should run.")
}
