// $2 can buy 1 bottle of beer.
// 4 bottle caps can be exchanged for 1 bottle beer.
// 2 empty bottles can be exchanged for 1 bottle of beer.
// How many bottles of beer can you get for $10?
// From: https://math.stackexchange.com/questions/1512083/the-number-of-bottles-of-beer-one-can-buy-with-10-after-exchanging-bottles-and
package puzzle

const cost = 2 // a bottle of beer costs $2
func BottlesOfBeer(empty int, caps int) int {

	if empty < 2 && caps < 4 {
		return 0
	}
	fb := 0
	// every two empty bottles yield a full bottle
	fb += empty / 2
	e2 := empty % 2
	// every four caps yield a full bottle
	fb += caps / 4
	c2 := caps % 4
	return fb + BottlesOfBeer(e2, c2)
}