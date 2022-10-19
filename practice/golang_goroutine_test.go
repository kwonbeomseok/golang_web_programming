package practice

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T) {
	t.Run("goroutine으로 값 출력하기", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			i := i
			go func() {
				fmt.Println(i)
			}()
		}
		fmt.Println("go routing 끝")
	})

	t.Run("goroutine 끝날때까지 기다리기", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(10)
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println(i)
				wg.Done()
			}
		}()
		wg.Wait()
		fmt.Println("go routine end")
	})
}
