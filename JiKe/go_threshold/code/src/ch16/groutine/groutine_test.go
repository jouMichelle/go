package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	// 创建一个协程使用
	for i := 0; i < 10; i++ {
		go func(i int) {
			//time.Sleep(time.Second * 1)
			fmt.Println(i)
		}(i)
	}
	// 程序等待时间
	time.Sleep(time.Millisecond * 50)
}
