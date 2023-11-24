package profiler

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	mem          runtime.MemStats
	profilerFile *os.File
	testsFile    *os.File
	profMtx      sync.Mutex
)

// Profiler is a struct for profiling
type Profiler struct {
	t            time.Time
	profilerFile string
	testsFile    string
	profiling    bool
}

// NewProfiler creates a new Profiler instance
func NewProfiler() *Profiler {

	profilerFile, err := os.Create("profiler\\.profiler.csv")

	if err != nil {
		panic(err)
	}

	_, err = profilerFile.WriteString("time,RAM\n")
	if err != nil {
		panic(err)
	}

	testsFile, err = os.Create("profiler\\.tests.csv")

	if err != nil {
		panic(err)
	}

	_, err = testsFile.WriteString("time,labels\n")
	if err != nil {
		panic(err)
	}

	return &Profiler{
		t:            time.Now(),
		profilerFile: "profiler\\.profiler.csv",
		testsFile:    "profiler\\.tests.csv",
		profiling:    true,
	}
}

// AddTest adds a test for writing to CSV
func (p *Profiler) AddTest(test string) {
	profMtx.Lock()
	defer profMtx.Unlock()

	p.strToCSV(testsFile, p.testsFile, time.Since(p.t).Seconds(), test)
}

func (p *Profiler) ProfilerThreadFunction() {
	for {

		profMtx.Lock()
		p.strToCSV(profilerFile, p.profilerFile, time.Since(p.t).Seconds(), float64(getMemoryUsage()))

		if !p.profiling {
			if profilerFile != nil {
				err := profilerFile.Close()

				if err != nil {
					fmt.Println(err)
				}
			}

			if testsFile != nil {
				err := testsFile.Close()

				if err != nil {
					fmt.Println(err)
				}
			}

			break
		}
		profMtx.Unlock()
	}
}

func (p *Profiler) strToCSV(file *os.File, filename string, time float64, data interface{}) {
	if file == nil {
		var err error
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
	}

	_, err := fmt.Fprintf(file, "%f,%v\n", time, data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// getMemoryUsage returns the current memory usage in bytes
func getMemoryUsage() uint64 {
	runtime.ReadMemStats(&mem)
	return mem.Alloc
}

func (p *Profiler) Close() {
	p.profiling = false
}
