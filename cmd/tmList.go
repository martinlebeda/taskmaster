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
	"github.com/spf13/cobra"
    "github.com/martinlebeda/taskmaster/service"
    "github.com/martinlebeda/taskmaster/termout"
)

// tmListCmd represents the tmList command
var tmListCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"ls"},
	Short: "list timer",
	Long: `List of timer records.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
        timerDistances := service.TmrGetDistance()
        termout.TmrListDistance(timerDistances)
	},
}

func init() {
	timerCmd.AddCommand(tmListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tmListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tmListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
