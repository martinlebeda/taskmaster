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
	"os"
	"text/tabwriter"
)

func WrkListWork(works []model.WorkList) {

	w := tabwriter.NewWriter(os.Stdout, 5, 2, 1, ' ', 0)
	for _, work := range works {

		duration := work.Stop.Sub(work.Start)

		//// check for today and format by this
		format := "2006-01-02 15:04"
		//t := time.Now()
		//roundedToday := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		//roundedGoal := time.Date(distance.Goal.Year(), distance.Goal.Month(), distance.Goal.Day(), 0, 0, 0, 0, t.Location())
		//if roundedToday == roundedGoal {
		//	format = "15:04"
		//}

		// output
		//if cndOut {
		//	fmt.Fprintf(w, "%s - %v - %s\n", duration.String(), distance.Goal.Format(format), distance.Note)
		//} else {
		fmt.Fprintf(w, "%d\t%s\t - %s  - \t%s -  \t%s\t %s\t %s\n", work.Rowid,
			work.Start.Format(format), work.Stop.Format(format), duration,
			work.Category, work.Code, work.Desc)
		//}
	}
	w.Flush()

	if isVerbose() {
		fmt.Println("\nCount of timers: ", len(works))
	}
}
