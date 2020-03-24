package share_mem

import (
	"sync"
	"testing"
	"time"
)

// 简单的共享内存并发,使用简单的协程，但是并不安全
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		// 协程中自增
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)

}

// 完善共享内存并发机制
func TestCounterThreadSafe(t *testing.T) {
	// 创建一个锁
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		// 创建协程
		go func() {
			// 异常释放锁
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)

}
