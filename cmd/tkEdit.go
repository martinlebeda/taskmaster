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
	"github.com/martinlebeda/taskmaster/service"
	"github.com/spf13/cobra"
)

// tkEditCmd represents the tkEdit command
var tkEditCmd = &cobra.Command{
	Use:   "tkEdit",
	Short: "A brief description of your command",
	// TODO Lebeda - add long description
	//Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		service.TskUpdate(taskOpt, args)
	},
}

func init() {
	taskCmd.AddCommand(tkEditCmd)

	tkEditCmd.Flags().StringVarP(&taskOpt.Prio.String, "prio", "p", "", "task priority")
	tkEditCmd.Flags().StringVarP(&taskOpt.Code.String, "code", "c", "", "task code (project)")
	tkEditCmd.Flags().StringVarP(&taskOpt.Category.String, "category", "g", "", "task category (list)")
	tkEditCmd.Flags().StringVarP(&taskOpt.Url.String, "url", "u", "", "url for this task (ie. sources on internet)")
	tkEditCmd.Flags().StringVarP(&taskOpt.Note.String, "note", "n", "", "path to file with note")
	tkEditCmd.Flags().StringVarP(&taskOpt.Script.String, "script", "s", "", "path to file with script")
	tkEditCmd.Flags().StringVarP(&taskOpt.Desc, "desc", "d", "", "task description")
}
