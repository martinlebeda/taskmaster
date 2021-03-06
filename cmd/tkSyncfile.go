// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/martinlebeda/taskmaster/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tkExportCmd represents the tkExport command
var tkSyncfileCmd = &cobra.Command{
	Use: "syncfile",
	// TODO Lebeda - check 1 argument
	Short: "sync tasks in todo.txt format",
	Long: `usage: tm task syncfile [file]

Do import next export with file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		service.TskImportTasks(args[0])
		service.TskExportTasks(args[0], []string{})
		if viper.GetString("afterchange") != "" {
			service.SysAfterChange()
		}
	},
}

func init() {
	taskCmd.AddCommand(tkSyncfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tkSyncfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tkSyncfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
