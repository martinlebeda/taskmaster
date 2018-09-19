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
	"testing"
)

// standard with ID
func TestParseTaskWithID(t *testing.T) {
	task := parseTask("(A) normální task @upřesnit id:123 +režie")
	if task.Id != 123 {
		t.Error("Bad value ID: ", task.Id)
	}
	if task.Desc != "(A) normální task @upřesnit +režie" {
		t.Error("Bad value desc: ", task.Desc)
	}
	if task.Status != "N" {
		t.Error("Bad value status: ", task.Status)
	}
}

// standard without ID
func TestParseTaskWithoutID(t *testing.T) {
	task := parseTask("(A) normální task @upřesnit +režie")
	if task.Id != 0 {
		t.Error("Bad value ID: ", task.Id)
	}
	if task.Desc != "(A) normální task @upřesnit +režie" {
		t.Error("Bad value desc: ", task.Desc)
	}
	if task.Status != "N" {
		t.Error("Bad value status: ", task.Status)
	}

}

// Done task
func TestParseTaskDone(t *testing.T) {
	task := parseTask("x 2018-09-18 normální task id:123 @upřesnit +režie")
	if task.Id != 123 {
		t.Error("Bad value ID: ", task.Id)
	}
	if task.Desc != "normální task @upřesnit +režie" {
		t.Error("Bad value desc: ", task.Desc)
	}
	if task.Status != "X" {
		t.Error("Bad value status: ", task.Status)
	}
}

// done without ID
func TestParseTaskDoneWitoutID(t *testing.T) {
	task := parseTask("x 2018-09-18 normální task @upřesnit +režie")
	if task.Id != 0 {
		t.Error("Bad value ID: ", task.Id)
	}
	if task.Desc != "normální task @upřesnit +režie" {
		t.Error("Bad value desc: ", task.Desc)
	}
	if task.Status != "X" {
		t.Error("Bad value status: ", task.Status)
	}
}
