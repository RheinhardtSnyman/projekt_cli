package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintMD5(t *testing.T) {
	in := strings.NewReader("golang")
	out := &bytes.Buffer{}
	printMD5(in, out)
	want := "21cc28409729565fc1a4d2dd92db269f"
	got := out.String()
	if got != want {
		t.Errorf("Want " + want + " but got " + got)
	}
}
