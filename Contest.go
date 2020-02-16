package main

type ProductOfNumbers struct {
	prefixProd []int
	lastZero   int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{[]int{1}, 0}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.prefixProd = append(this.prefixProd, 1)
		this.lastZero = len(this.prefixProd) - 1
	} else {
		this.prefixProd = append(this.prefixProd, this.prefixProd[len(this.prefixProd)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	start := len(this.prefixProd) - 1 - k
	if this.lastZero > start {
		return 0
	}
	return this.prefixProd[len(this.prefixProd)-1] / this.prefixProd[start]
}
func main() {
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func twodarray(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	return res
}
