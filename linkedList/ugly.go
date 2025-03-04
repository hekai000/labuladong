package linkedlist

func nthUglyNumber(n int) int {
	p2, p3, p5 := 1, 1, 1

	product2, product3, product5 := 1, 1, 1

	p := 1
	ugly := make([]int, n+1)

	for p <= n {

		min := min(min(product2, product3), product5)
		ugly[p] = min
		p++

		if min == product2 {
			product2 = 2 * ugly[p2]
			p2++
		}
		if min == product3 {
			product3 = 3 * ugly[p3]
			p3++
		}
		if min == product5 {
			product5 = 5 * ugly[p5]
			p5++
		}

	}
	return ugly[n]
}
