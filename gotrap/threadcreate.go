package main

import (
	"bufio"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

// curl http://localhost:8080/debug/pprof/threadcreate

func main() {
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 100; i++ {
		go func(id int) {
			for {
				data, _, _ := reader.ReadLine()
				println(id, string(data))
			}
		}(i)
		time.Sleep(1 * time.Second)
	}
	println("create finish!")
	time.Sleep(5 * time.Second)
}
