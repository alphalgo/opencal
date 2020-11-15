package opencal

import (
	"reflect"
)

// Ocal defines opencal struct
type Ocal struct {
	op        string
	direction int
}

// SetOp sets the op value
func (o *Ocal) SetOp(direction int) (op string) {
	if direction == -1 {
		o.op = "/"
		o.direction = direction
	} else if direction == 0 {
		o.op = "|"
		o.direction = direction
	} else if direction == 1 {
		o.op = "-"
		o.direction = direction
	}
	return
}

// Cal runs the calculation
func (o *Ocal) Cal(a, b interface{}, direction int) {
	op := o.SetOp(direction)

	switch op {
	case "/":
		o.D(a, b)
	case "|":
		o.M(a, b)
	case "-":
		o.S(a, b)
	default:
	}
}

// isSameType checks if type of a and b is same
func isSameType(a, b interface{}) bool {
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		return true
	}
	return false
}

// D calculates the value between a and b
func (o *Ocal) D(a, b interface{}) (v interface{}) {
	if isSameType(a, b) {
		return o.d(a, b)
	}
	return o.Dcal(a, b)
}

// d calculates the same type value
func (o *Ocal) d(a, b interface{}) (v interface{}) {
	switch vl := reflect.TypeOf(a); vl.Kind() {
	case reflect.Int:
		v = nil
	default:
	}
	return v
}

func (o *Ocal) Dcal(a, b interface{}) (v interface{}) {
	return nil
}

// M calculates the value between a and b
func (o *Ocal) M(a, b interface{}) (v interface{}) {
	return nil
}

// S calculates the value between a and b
func (o *Ocal) S(a, b interface{}) (v interface{}) {
	return nil
}
