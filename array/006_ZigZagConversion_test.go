package array

import (
	"testing"
)

func TestConvert(t *testing.T) {
	want := "AB"
	s := "AB"
	numRows := 1
	if got := convert(s, numRows); got != want {
		t.Errorf("convert(%s, %d)=%s, want %s", s, numRows, got, want)
	}

	want = "PAHNAPLSIIGYIR"
	s = "PAYPALISHIRING"
	numRows = 3
	if got := convert(s, numRows); got != want {
		t.Errorf("convert(%s, %d)=%s, want %s", s, numRows, got, want)
	}

	want = "PINALSIGYAHRPI"
	s = "PAYPALISHIRING"
	numRows = 4
	if got := convert(s, numRows); got != want {
		t.Errorf("convert(%s, %d)=%s, want %s", s, numRows, got, want)
	}

	want = ""
	s = ""
	numRows = 4
	if got := convert(s, numRows); got != want {
		t.Errorf("convert(%s, %d)=%s, want %s", s, numRows, got, want)
	}
}
