/*
Copyright © 2022 Farye Nwede <farye@aeekay.com>

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
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dbConfigFilename string

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate [config]",
	Short: "A way to migrate a database based on the yaml passed in",
	Long: `A way to migrate a database based on the yaml passed in.

descartes db migrate <YAML>

This allows you to migrate your database with the parameter YAML.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("migrate called")
		// retrieve the argument and use it as a filename
		dbConfigFilename = args[0]

		// validate the path to see if it exists
		if _, err := os.Stat(dbConfigFilename); errors.Is(err, os.ErrNotExist) {
			fmt.Printf("the file %s does not exist. exiting.", dbConfigFilename)
		}

		// if all is well, run the migration
	},
}

func init() {
	dbCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().StringVar(&dbConfigFilename, "filename", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
