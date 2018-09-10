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
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

func TmrListDistance(distances []model.TimerDistance, cndOut bool) {

	d := color.New(color.Bold)
	if viper.GetBool("color") {
		color.NoColor = false // disables colorized output
	}

	// build output
	output := []string{}
	for _, distance := range distances {
		duration, _ := time.ParseDuration(strconv.Itoa(distance.Distance) + "s")

		// check for today and format by this
		format := "2006-01-02 15:04"
		t := time.Now()
		roundedToday := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		roundedGoal := time.Date(distance.Goal.Year(), distance.Goal.Month(), distance.Goal.Day(), 0, 0, 0, 0, t.Location())
		if roundedToday == roundedGoal {
			format = "15:04"
		}

		// output
		if cndOut {
			fmt.Println(duration.String(), "-", distance.Goal.Format(format), "-", distance.Note)
		} else {
			out := fmt.Sprintf("%d | %s | %s | %s | %s",
				distance.Rowid,
				duration.String(),
				distance.Goal.Format(format),
				distance.Tag,
				distance.Note)

			output = append(output, out)
		}
	}

	// columize
	outFmt := strings.Split(columnize.SimpleFormat(output), "\n")

	// printout
	for i, distance := range distances {
		if distance.Distance < 0 {
			d.Println(outFmt[i])
		} else {
			fmt.Println(outFmt[i])
		}
	}

	if isVerbose() {
		fmt.Println("\nCount of timers: ", len(distances))
	}
}
