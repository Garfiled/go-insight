# Go Monitor

## Monitor Data

* goroutine num
* memory allocation
* gc
* fd

## Use
```go
func goProcMonitor(dur time.Duration) {
    ticker := time.Tick(dur)
    for range ticker {
        stats := gomonitor.GetRuntimeStats()
        fmt.Printf("gomonitor:%+v\n", stats)
    }
}
```
## gomonitor.go
```go
package gomonitor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type TGoProcStats struct {
	Goroutines int

	Mem_Allocated uint64
	Mem_Objects   uint64
	Mem_Mallocs   uint64
	Mem_Heap      uint64
	Mem_Stack     uint64

	Gc_Num   uint32
	Gc_Pause uint64
	Gc_Next  uint64

	Cgo int64
	Fds int
}

var (
	lastNumGc      uint32
	lastPauseTotal uint64
	lastCgoCall    int64
	lastMallocs    uint64
)

func GetRuntimeStats() *TGoProcStats {
	var (
		mem   runtime.MemStats
		stats TGoProcStats
	)

	runtime.ReadMemStats(&mem)

	stats.Mem_Mallocs = mem.Mallocs - lastMallocs
	lastMallocs = mem.Mallocs
	stats.Mem_Allocated = mem.Alloc
	stats.Mem_Objects = mem.HeapObjects
	stats.Mem_Heap = mem.HeapAlloc
	stats.Mem_Stack = mem.StackInuse

	stats.Gc_Num = mem.NumGC - lastNumGc
	lastNumGc = mem.NumGC
	stats.Gc_Pause = mem.PauseTotalNs - lastPauseTotal
	lastPauseTotal = mem.PauseTotalNs
	stats.Gc_Next = mem.NextGC

	stats.Goroutines = runtime.NumGoroutine()
	temp := runtime.NumCgoCall()
	stats.Cgo = temp - lastCgoCall
	lastCgoCall = temp
	stats.Fds = openFileCnt()

	return &stats
}

func openFileCnt() int {
	for i := 0; i < 2; i++ {
		out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -p %v", os.Getpid())).Output()
		if err == nil {
			return bytes.Count(out, []byte("\n"))
		}
	}
	return 0
}
```
