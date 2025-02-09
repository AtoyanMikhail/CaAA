package main

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func min(values ...int) int {
	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func isPowerOfTwoMinusOne(n int) bool {
    return (n+1)&n == 0
}

func getExponent(n int) int {
    e := 0
    for n > 1 {
        n >>= 1
        e++
    }
    return e
}