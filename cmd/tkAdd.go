// Copyright © 2018 Martin Lebeda <martin.lebeda@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/service"
	"github.com/spf13/cobra"
)

var taskOpt model.Task

// tkAddCmd represents the tkAdd command
var tkAddCmd = &cobra.Command{
	Use:   "add",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	// TODO Lebeda - add long description
	//Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskOpt.Desc = args[0]
		service.TskAdd(taskOpt)

	},
}

func init() {
	taskCmd.AddCommand(tkAddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tkAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tkAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	tkAddCmd.Flags().Int64VarP(&taskOpt.ParentId.Int64, "parent", "t", 0, "parent task id")
	tkAddCmd.Flags().StringVarP(&taskOpt.Prio.String, "prio", "p", "", "task priority")
	tkAddCmd.Flags().StringVarP(&taskOpt.Url.String, "url", "u", "", "url for this task (ie. sources on internet)")
	tkAddCmd.Flags().StringVarP(&taskOpt.Note.String, "note", "n", "", "path to file with note")
	tkAddCmd.Flags().StringVarP(&taskOpt.Script.String, "script", "s", "", "path to file with script")
}

// TODO Lebeda - add if not exists
