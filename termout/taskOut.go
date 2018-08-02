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

package termout

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/martinlebeda/taskmaster/model"
	"github.com/ryanuber/columnize"
	"strconv"
	"strings"
)

func TskListTasks(tasks []model.Task) {

	d := color.New(color.Bold)   // TODO Lebeda - tools funkce
	x := color.New(color.FgCyan) // TODO Lebeda - tools funkce

	// build output
	var output []string
	for _, task := range tasks {
		statuFmt := "" + task.Status + ""
		if statuFmt == "N" {
			statuFmt = ""
		}

		prioFmt := formatPrio(task)

		out := fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s | %s | %s | %s",
			strconv.Itoa(task.Id),
			statuFmt,
			task.Estimate.String,
			prioFmt,
			task.Category.String,
			task.Code.String,
			task.Desc,
			task.Url.String,
			task.Note.String,
			task.Script.String)
		output = append(output, out)

	}

	// columize
	outFmt := strings.Split(columnize.SimpleFormat(output), "\n")

	// printout
	for i, task := range tasks {
		if task.Prio.String == "A" || task.Prio.String == "B" {
			d.Println(outFmt[i])
		} else if task.Status == "X" {
			x.Println(outFmt[i])
		} else {
			fmt.Println(outFmt[i])
		}
	}

	if isVerbose() {
		fmt.Println("\nCount of tasks: ", len(tasks))
	}
}

func TskShowWork(task model.Task) {
	prioFmt := formatPrio(task)

	trimSpace := strings.TrimSpace(prioFmt + " " + task.Code.String + " " + task.Desc)
	replace := strings.Replace(trimSpace, "  ", " ", -1)
	fmt.Println(task.Id, "-", replace)
}

func formatPrio(task model.Task) string {
	prioFmt := "(" + task.Prio.String + ")"
	if prioFmt == "()" {
		prioFmt = ""
	}
	return prioFmt
}
