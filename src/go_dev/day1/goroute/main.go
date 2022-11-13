package main

import (
	"time"
)

// 在同一个包下，文件存在依赖关系时，应该把依赖的文件都传入加载参数里面
//
//	eg:  go run main.go goroute.go
func main() {

	for i := 0; i < 100; i++ {
		go test_goroute(i)
	}

	time.Sleep(time.Second)
}
