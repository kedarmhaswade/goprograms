// Fortunately, the language helps us here. If the receiver p is a variable of type Point (on which method is defined)
// but the method requires a *Point receiver, we can use this shorthand:
// p.ScaleBy(2)
package ch6

type String string // we define a new type String so that we can define methods on it

// Here, to illustrate the point however, we define a method that requires a *String receiver

//Length returns the length of the string on which it is called
//For nil references it returns 0
func (pstr *String) Length() int {
	if pstr == nil {
		return 0
	}
	return len(*pstr)
}