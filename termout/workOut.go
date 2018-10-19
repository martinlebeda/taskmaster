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
	"github.com/martinlebeda/taskmaster/model"
	"github.com/martinlebeda/taskmaster/tools"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func WrkListWork(works []model.WorkList, showTable bool) {

	w := tabwriter.NewWriter(os.Stdout, 5, 2, 1, ' ', 0)
	for _, work := range works {

		format := "2006-01-02 15:04"

		stopFmt := ""
		duration := time.Now().Sub(work.Start)

		if work.Stop.After(tools.GetZeroTime()) {
			duration = work.Stop.Sub(work.Start)
			stopFmt = work.Stop.Format(format)
		}

		durationFmt := strings.TrimRight(duration.Round(time.Minute).String(), "0s")

		if showTable {
			fmt.Fprintf(w, "%d\t%s\t - %s  \t%s  \t %s\n", work.Id, work.Start.Format(format), stopFmt, durationFmt, work.Desc)
		} else {
			format = "15:04"
			fmt.Fprintf(w, "%s (%s) %s", work.Start.Format(format), durationFmt, work.Desc)
		}
	}
	w.Flush()

	if isVerbose() {
		fmt.Println("\nCount of timers: ", len(works)) // TODO Lebeda - zajistit součet odpracovaného času
	}
}

func WrkSumWork(sums []model.WorkSum) {

	w := tabwriter.NewWriter(os.Stdout, 5, 2, 1, ' ', 0)
	for _, workSum := range sums {
		duration, _ := time.ParseDuration(strconv.Itoa(workSum.Seconds) + "s")
		fmt.Fprintf(w, "%s\t    %s\n", workSum.Desc, duration)
	}
	w.Flush()
}
