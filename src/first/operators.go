package main

import "fmt"

/* Go binary operators
<--------associate to left--------------
 H   *    /    %    <<    >>    &    &^
 |   +    -    |    ^
 |   ==   !=   <    <=    >     >=
 |   &&
 L   ||
*/
func main() {
	r := 1 + 2*3
	fmt.Printf("1 + 2 * 3 = %d because * is at a higher precedence\n", r)
	r = 10 * 2 / 5
	fmt.Printf("10 * 2 / 5 = %d because at the same level * and / associate to the left (operations are grouped from left)\n", r)
	fmt.Printf("%s\n", "This may not be a surprise in the above example, but")
	r = 12 % 5 * 2
	fmt.Printf("12 %% 5 * 2 = %d gives the wrong answer to the question \"What is the remainder when the double of 5 divides 12?\"\n", r)
	r = 12 % (5 * 2)
	fmt.Printf(", whereas 12 %% (5 * 2) = %d gives the right one\n", r)
}
