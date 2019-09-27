package fan1

import "fmt"

const (
	Table   string = "table"
	Ceiling string = "ceiling"
)

type Fan2 struct {
	speed int    // 1 through 5
	kind  string // "table", "ceiling"
	color string
}

// First, we define an option type. It is a function that takes one argument, the Foo (Fan2) we are operating on.
type option func(fan *Fan2)

// The idea is that an option is implemented as a function we call to set the state of that option.
// That may seem odd, but there's a method in the madness.

// Given the option type, we next define an Option method on *Fan2 that applies the options it's passed
// by calling them as functions. That method is defined in the same package, say fan2, in which Fan2 is defined.

// Option sets the options specified.
func (f *Fan2) Option(opts ...option) {
	for _, opt := range opts {
		opt(f)
	}
}

// Now to provide an option, we define in package fan1 a function with the appropriate name and signature.
// Let's say we want to control speed by setting an integer value stored in a field of a Fan2.
// We provide the speed option by writing a function with the obvious name and have it return an option,
// which means a closure; inside that closure we set the field (Me: encapsulation?) :

//Speed returns a function (closure) that accepts a Fan2 and sets its speed to the given speed
func Speed(s int) option {
	return func(fan *Fan2) {
		fan.speed = s
	}
}

//Kind returns a function (closure) that accepts a Fan2 and sets its kind to the given kind
func Kind(k string) option {
	return func(fan *Fan2) {
		fan.kind = k
	}
}

//Color returns a function (closure) that accepts a Fan2 and sets its color to the given color
func Color(c string) option {
	return func(fan *Fan2) {
		fan.color = c
	}
}
// Why return a closure instead of just doing the setting?
// Because we don't want the user to have to write the closure and we want the
// Option method to be nice to use. (Plus there's more to come....)

func New() *Fan2 {
	return &Fan2{kind: Ceiling}
}
func (f *Fan2) String() string {
	return fmt.Sprintf("Fan with speed: %v, kind: %v, color: %v", f.speed, f.kind, f.color)
}