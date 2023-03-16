package test

// import (
// 	"fmt"
// 	"sync"
// 	"sync/atomic"
// 	"time"

// 	"golang.org/x/sync/singleflight"
// )

// // import (
// // 	"fmt"
// // 	"strconv"
// // 	"strings"
// // 	"time"
// // )

// // func ParseDuration(d string) (time.Duration, error) {
// // 	d = strings.TrimSpace(d)
// // 	dr, err := time.ParseDuration(d)
// // 	if err == nil {
// // 		return dr, nil
// // 	}
// // 	if strings.Contains(d, "d") {
// // 		index := strings.Index(d, "d")

// // 		hour, _ := strconv.Atoi(d[:index])
// // 		dr = time.Hour * 24 * time.Duration(hour)
// // 		ndr, err := time.ParseDuration(d[index+1:])
// // 		if err != nil {
// // 			return dr, nil
// // 		}
// // 		return dr + ndr, nil
// // 	}

// // 	dv, err := strconv.ParseInt(d, 10, 64)
// // 	return time.Duration(dv), err
// // }

// // func main() {
// // 	a, _ := ParseDuration("2y")
// // 	fmt.Printf("a: %v\n", a)
// // }

// var count int32

// func main() {
// 	total := 1000
// 	sg := &singleflight.Group{}

// 	var wg sync.WaitGroup
// 	wg.Add(total)

// 	key := "key"
// 	for i := 0; i < total; i++ {
// 		go func() {
// 			defer wg.Done()
// 			sg.Do(key, func() (interface{}, error) {
// 				res, err := getData(key)
// 				return res, err
// 			})
// 			// getData(key)
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Printf("total num is %v\n", count)
// }

// func getData(key string) (interface{}, error) {
// 	atomic.AddInt32(&count, 1)
// 	time.Sleep(100 * time.Millisecond)
// 	return "result", nil
// }
