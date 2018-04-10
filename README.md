# roundrobin
Golang round-robin

```javascript
package main

import (
	"fmt"
	"sync"

	"github.com/sbabiv/roundrobin"
)

func main() {
	source := []interface{}{1, 2, 3}
	b := roundrobin.New(source)
	wc := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wc.Add(1)
		go func() {
			v, _ := b.Pick()
			fmt.Printf("%v\n", v.(int))
			wc.Done()
		}()
	}

	wc.Wait()
}
```
