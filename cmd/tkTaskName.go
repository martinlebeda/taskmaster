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
	"fmt"
	"github.com/martinlebeda/taskmaster/service"
	"github.com/martinlebeda/taskmaster/tools"
	"github.com/spf13/cobra"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// tkTaskNameCmd represents the tkTaskName command
var tkTaskNameCmd = &cobra.Command{
	Use: "taskname",
	// TODO Lebeda - check only 1 arg
	Short: "printout task name without context (@...), projects (+...), urls and priority",
	Long:  `usage: tm tk taskname ID`,
	Run: func(cmd *cobra.Command, args []string) {
		normalize, err := cmd.Flags().GetBool("normalize")
		tools.CheckErr(err)

		fs, err := cmd.Flags().GetBool("fs")
		tools.CheckErr(err)

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			tools.CheckErr(err)
			task := service.TskGetById(id)
			name := task.Desc

			name = regexp.MustCompile(`^\(.\) `).ReplaceAllString(name, "")
			name = regexp.MustCompile(` \+[^ ]*`).ReplaceAllString(name, "")
			name = regexp.MustCompile(` @[^ ]*`).ReplaceAllString(name, "")
			name = regexp.MustCompile(` http://[^ ]*`).ReplaceAllString(name, "")
			name = regexp.MustCompile(` https://[^ ]*`).ReplaceAllString(name, "")

			if normalize {
				isMn := func(r rune) bool {
					return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
				}
				t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
				name, _, _ = transform.String(t, name)
			}

			if fs {

				/*
				   illegal characters:
				     < (less than)
				     > (greater than)
				     : (colon - sometimes works, but is actually NTFS Alternate Data Streams)
				     " (double quote)
				     / (forward slash)
				     \ (backslash)
				     | (vertical bar or pipe)
				     ? (question mark)
				     * (asterisk)
				*/
				name = regexp.MustCompile(`[<>:"/\\|?*']`).ReplaceAllString(name, " ")
				name = regexp.MustCompile(`\s+`).ReplaceAllString(name, " ")
			}

			fmt.Println(strings.TrimSpace(name))
		}
	},
}

func init() {
	taskCmd.AddCommand(tkTaskNameCmd)

	tkTaskNameCmd.Flags().BoolP("normalize", "n", false, "normalize name to ascii")
	tkTaskNameCmd.Flags().BoolP("fs", "f", false, "normalize name to filesystem friendly")
}
