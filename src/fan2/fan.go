package fan2

// Rob Pike: Fan1 is good for most cases, but I want to be able to use the option mechanism to set temporary values,
// which means it would be nice if the Option method could return the previous state. That's easy: just save it
// in an empty interface value that is returned by the Option method and the underlying function type.
// That value flows through the code as below. Note that the previous state is only of the option, _not_ of the thing
// i.e. the fan (Fan2) that we are operating on.
import (
	"fmt"
	"time"
)

const (
	Table   string = "table"
	Ceiling string = "ceiling"
)

type Fan2 struct {
	speed int    // 1 through 5
	kind  string // "table", "ceiling"
	color string
}

type option func(fan *Fan2) interface{} // compare with fan1.option which denotes a function that returns nothing

// Option, as for Fan1, sets the options specified.
// It returns the previous value of the last argument.
func (f *Fan2) Option(opts ...option) (previous interface{}) {
	for _, opt := range opts {
		previous = opt(f)
	}
	return previous
}

// Now to provide an option, we define in package fan1 a function with the appropriate name and signature.
// Let's say we want to control speed by setting an integer value stored in a field of a Fan2.
// We provide the speed option by writing a function with the obvious name and have it return an option,
// which means a closure; inside that closure we set the field (Me: encapsulation?) :

//Speed returns a function (closure) that accepts a Fan2, sets its speed to the given speed and returns previous speed
func Speed(s int) option {
	return func(fan *Fan2) interface{} {
		prevSpeed := fan.speed
		fan.speed = s
		return prevSpeed
	}
}

//Kind returns a function (closure) that accepts a Fan2 and sets its kind to the given kind and returns previous kind
func Kind(k string) option {
	return func(fan *Fan2) interface{} {
		prevKind := fan.kind
		fan.kind = k
		return prevKind
	}
}

//Color returns a function (closure) that accepts a Fan2 and sets its color to the given color
func Color(c string) option {
	return func(fan *Fan2) interface{} {
		prevColor := fan.color
		fan.color = c
		return prevColor
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

func (f *Fan2) Run() {
	fmt.Printf("Running ... %v\n", f)
	time.Sleep(100 * time.Millisecond)
}
