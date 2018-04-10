# roundrobin
Golang round-robin

```
package main

import (
	"fmt"
	"sync"
	"test2/roundrobin"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	source := []interface{}{
		1,
		2,
		3,
	}

	b := roundrobin.New(source)
	wc := sync.WaitGroup{}

	for i := 0; i < 9; i++ {
		wc.Add(1)
		go func(blr *roundrobin.Balancer) {
			v, _ := blr.Pick()
			fmt.Printf("%v\n", v.(int))
			wc.Done()
		}(b)
	}

	wc.Wait()
}
```
