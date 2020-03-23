package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 创建一个计算函数运行时间的函数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

// 定义一个可变长参数函数
func Sum(ops ...int) int {
	ret := 0
	// 传过来的参数会放入一个数组
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

// 延迟执行函数,延迟执行函数;一般用于释放资源与释放锁;它是函数中最后执行的且一定会执行
func TestDefer(t *testing.T) {
	// 调用次函数不会马上执行,而是在执行下面函数后,在函数返回前才执行此函数
	defer Clear()
	fmt.Println("Start")
	// 抛出一个程序异常中断,
	panic("err")
}
