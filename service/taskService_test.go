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

func TestRemovePrioFromDesc(t *testing.T) {

	if RemovePrioFromDesc("xxx xxx") != "xxx xxx" {
		t.Error("remove from desc without priority")
	}

	if RemovePrioFromDesc("xxx (A) xxx") != "xxx (A) xxx" {
		t.Error("remove from desc without priority")
	}

	desc := RemovePrioFromDesc("(A) xxx xxx")
	if desc != "xxx xxx" {
		t.Error("remove from desc with A priority: ", desc)
	}

	//no shift
	//timeResult := getDateTime("", "17:10", "2018-07-11")
	//
	//if timeResult.Format("2006-01-02 15:04") != "2018-07-11 17:10" {
	//	t.Error("Expected 2018-07-11 17:10, got ", timeResult.Format("2006-01-02 15:04"))
	//}
	//
	//with shift
	//timeResult = getDateTime("-1h", "17:10", "2018-07-11")
	//
	//if timeResult.Format("2006-01-02 15:04") != "2018-07-11 16:10" {
	//	t.Error("Expected 2018-07-11 16:10, got ", timeResult.Format("2006-01-02 15:04"))
	//}

}
