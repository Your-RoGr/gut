package gut

import (
	"fmt"
	"gut/gogger"
	"os"
	"os/exec"
	"time"
)

// TestRunner is a struct for running tests
type TestRunner struct {
	dataSize DataSize
	tests    []func()
}

// NewTestRunner creates a new TestRunner instance
func NewTestRunner(dataSize DataSize) *TestRunner {

	logger.SetUseConsoleLog(true)
	logger.SetUseFileLog(true)
	logger.SetLogLevelConsole(gogger.DEBUG)

	err := logger.SetLogFormat("[%timestamp%] [%level%] %message%")
	if err != nil {
		fmt.Println(err)
	}

	go prof.ProfilerThreadFunction()

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	}

	go executeCommand("python", []string{currentDir + "\\profiler\\_visualization_profiler.py", "kb"})

	return &TestRunner{
		dataSize: dataSize,
		tests:    make([]func(), 0),
	}
}

// AddTests adds new tests to the TestRunner
func (tr *TestRunner) AddTests(tests ...func()) {
	tr.tests = append(tr.tests, tests...)
}

// RunTests runs all the tests
func (tr *TestRunner) RunTests() {
	defer logger.Close()
	defer prof.Close()

	failedTests = 0

	prof.AddTest("Start of tests")
	time.Sleep(time.Millisecond)

	for _, test := range tr.tests {
		test()
	}

	time.Sleep(time.Millisecond)
	prof.AddTest("End of tests")
	prof.AddTest("")

	totalTests := equalityTests + exceptionTests + boundaryTests + performanceTests + concurrencyTests
	logger.Info(fmt.Sprintf("Total tests passed: %d/%d", totalTests-failedTests, totalTests))

	if totalTests-failedTests != totalTests {
		var str string
		for _, val := range failed {
			str += "\n" + val
		}
		logger.Warning("The following tests failed:" + str)
	}
}

// SetUs sets the delay value for the profiler in microseconds
func SetUs(microseconds int) {
	prof.SetUs(microseconds)
}

func executeCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}
