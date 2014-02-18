// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

import (
	"fmt"
)

type language interface {
	isValidHeader(line string) bool
}

var languages = []language{ja{}}

func isValidNonAlphabeticalScriptHeader(line string) bool {
	for _, v := range languages {
		if fmt.Sprintf("%T", v) == "doc."+commentLang {
			if v.isValidHeader(line) {
				return true
			}
			return false
		}
	}
	return false
}
