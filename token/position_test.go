package token

import "testing"

func TestIsValid(t *testing.T) {
	var pos Position
	if pos.IsValid() {
		t.Errorf("A position with line 0 should not be valid.")
	}

	pos.Line = 1
	if !pos.IsValid() {
		t.Errorf("A position with line 1 should be valid.")
	}

	pos.Line = -1
	if pos.IsValid() {
		t.Errorf("A position with line -1 should not be valid.")
	}
}

func TestString(t *testing.T) {
	var pos Position
	if pos.String() != "-" {
		t.Errorf("Position string: got '%s'; want '-'.", pos.String())
	}

	pos.Line = 1
	if pos.String() != "1:0" {
		t.Errorf("Position string: got '%s'; want '1:0'.", pos.String())
	}

	pos.Column = 2
	if pos.String() != "1:2" {
		t.Errorf("Position string: got '%s'; want '1:2'.", pos.String())
	}

	pos.Filename = "file"
	if pos.String() != "file:1:2" {
		t.Errorf("Position string: got '%s'; want 'file:1:2'.", pos.String())
	}

	pos.Line = 0
	if pos.String() != "file" {
		t.Errorf("Position string: got '%s'; want 'file'.", pos.String())
	}
}
