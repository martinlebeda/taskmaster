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
	"fmt"
	"github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/tools"
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func SysNotifyDistance(distance model.TimerDistance) {
	format := tools.BaseDateTimeFormat
	t := time.Now()
	roundedToday := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	roundedGoal := time.Date(distance.Goal.Year(), distance.Goal.Month(), distance.Goal.Day(), 0, 0, 0, 0, t.Location())
	if roundedToday == roundedGoal {
		format = "15:04"
	}

	duration, _ := time.ParseDuration(strconv.Itoa(distance.Distance) + "s")

	msg := fmt.Sprintf("'%s - %v - %s'", duration.String(), distance.Goal.Format(format), distance.Note)

	notifycmd := viper.GetString("notifycmd")
	cmd := exec.Command(notifycmd, msg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
		log.Printf("combined out:\n%s\n", string(out))
	}
}

func SysAfterChange() {
	afterChangeCmd := viper.GetString("afterchange")
	fields := strings.Fields(afterChangeCmd)
	cmd := exec.Command(fields[0], fields[1:]...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Command finished with error: %v", err)
		log.Printf("combined out:\n%s\n", string(out))
	}
}
