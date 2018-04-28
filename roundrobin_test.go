package roundrobin

import (
	"sync"
	"testing"
)

func TestRoundrobin(t *testing.T) {
	m := make(map[int]int, 5)
	mu := new(sync.Mutex)
	wg := sync.WaitGroup{}
	b := New([]interface{}{
		1,
		2,
		3,
		4,
		5,
	})

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			item, _ := b.Pick()
			val := item.(int)
			mu.Lock()
			m[val]++
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	for _, v := range m {
		if v != 200 {
			t.FailNow()
		}
	}
}
