package series

import "fmt"

// 这是一个初始化方法，在main函数执行前执行
func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

func Square(n int) int {
	return n * n
}

// 小写的方法名是无法在包外被访问
func GetFibonacciSerie(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
