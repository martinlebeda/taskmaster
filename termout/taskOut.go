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

func TskListTasks(tasks []model.Task) {

	w := tabwriter.NewWriter(os.Stdout, 5, 2, 1, ' ', 0)
	for _, task := range tasks {
		fmt.Fprintf(w, "%d\t  %s\t%s   \t%s\t %s\t %s\n",
			task.Id, task.Prio.String, task.Desc, task.Url.String, task.Note.String, task.Script.String)
		//}
	}
	w.Flush()

	if isVerbose() {
		fmt.Println("\nCount of tasks: ", len(tasks))
	}
}
