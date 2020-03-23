package fib

import (
	"testing"
)

// 测试⽅方法名以 Test 开头：func TestXXX(t *testing.T) {…}
// 实现 Fibonacci 数列列
func TestFibList(t *testing.T) {
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

// 经典，进行两个值交换
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
