// Copyright 2016 Vastri. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package token

import "fmt"

// Position describes an arbitrary source position.
type Position struct {
	Filename string
	Line     int
	Column   int
}

// IsValid returns true is the position is valid.
func (pos *Position) IsValid() bool { return pos.Line > 0 }

// String returns a string corresponding to the position pos.
func (pos Position) String() string {
	s := pos.Filename
	if pos.IsValid() {
		if s != "" {
			s += ":"
		}
		s += fmt.Sprintf("%d:%d", pos.Line, pos.Column)
	}
	if s == "" {
		s = "-"
	}
	return s
}
