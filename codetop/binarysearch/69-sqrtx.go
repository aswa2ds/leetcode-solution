package binarysearch

func mySqrt(x int) int {
	x0, C, x1 := float64(x*x), float64(x), float64(x)
	for x0-x1 > 10e-6 {
		x0 = x1
		x1 = (x0*x0 + C) / (2 * x0)
	}
	return int(x1)
}
