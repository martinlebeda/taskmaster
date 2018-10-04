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

package service

import (
	"bufio"
	"github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/tools"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func TskExportTasks(filename string, ids []string) {
	tasks := TskGetList(time.Now().Add(24*time.Hour), false, "", []string{}, []string{}, ids)
	TskExportOutput(tasks, filename)
}

func TskImportTasks(filename string) {
	tasks := TskParseImport(filename)
	TskDoImport(tasks)
}

func TskExportOutput(tasks []model.Task, exportfileName string) {
	f, err := os.Create(exportfileName)
	tools.CheckErr(err)

	defer f.Close()

	for _, task := range tasks {
		// fmt.Println(task.Desc + " id:" + strconv.Itoa(task.Id)) TODO Lebeda - when verbose
		f.WriteString(task.Desc + " id:" + strconv.Itoa(task.Id) + "\n")
		tools.CheckErr(err)
	}
	// TODO Lebeda - on verbose counting
}

func TskParseImport(importFile string) []model.Task {
	file, err := os.Open(importFile)
	tools.CheckErr(err)
	defer file.Close()

	var result []model.Task

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		task := parseTask(text)

		result = append(result, task)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func TskDoImport(tasks []model.Task) {
	for _, task := range tasks {
		if task.Status != "X" {
			task.Status = ""
		}
		if task.Id == 0 {
			TskAdd(task)
		} else {
			TskUpdate(task, []string{strconv.Itoa(task.Id)})
		}
	}
	TskRepairDoneDate()
	// TODO Lebeda - on verbose counting
}

func parseTask(text string) model.Task {
	var task model.Task

	// status
	if strings.HasPrefix(text, "x ") {
		task.Status = "X"
	} else {
		task.Status = "N"
	}

	// id
	re := regexp.MustCompile(" id:\\d+")
	findString := re.FindString(text)
	trimSpace := strings.TrimSpace(findString)
	idString := strings.Replace(trimSpace, "id:", "", 1)
	//fmt.Println(idString)
	if idString != "" {
		id, err := strconv.Atoi(idString)
		task.Id = id
		tools.CheckErr(err)
	}

	// desc
	desc := re.ReplaceAllString(text, "")
	if strings.HasPrefix(text, "x ") {
		reDone := regexp.MustCompile("^x \\d{4}-\\d{2}-\\d{2} ")
		desc = reDone.ReplaceAllString(desc, "")
	}
	task.Desc = strings.TrimSpace(desc)
	//fmt.Println(text)
	return task
}
