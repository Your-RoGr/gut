package gut

import (
	"gut/gogger"
	"gut/profiler"
)

var (
	failedTests      int
	equalityTests    int
	exceptionTests   int
	boundaryTests    int
	performanceTests int
	concurrencyTests int
	failed           []string
	logger, _        = gogger.NewGogger("gut.log", "logs", 1000000, 1)
	prof             = profiler.NewProfiler()
)
