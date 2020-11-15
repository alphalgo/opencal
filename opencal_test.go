package opencal

import (
	"testing"

	"github.com/i0Ek3/asrt"
)

func TestCal(t *testing.T) {
	o := &Ocal{}
	a := 1
	b := "0"
	direction := -1

	o.Cal(a, b, direction)
	got := o.D(a, b)
	want := ""

	asrt.Asrt(t, got, want)
}
