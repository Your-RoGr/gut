package gut

import "fmt"

// EqualityTest is a struct for testing equality of values
type EqualityTest[T comparable] struct{}

// AssertEquals checks equality of values
func (et *EqualityTest[T]) AssertEquals(actual, expected T, name, message string) {
	equalityTests++
	if actual == expected {
		logger.Info(fmt.Sprintf("#%d%s EqualityTest PASSED", equalityTests, name))
	} else {
		failedTests++
		logger.Warning(fmt.Sprintf("#%d%s EqualityTest FAILED: %s. Expected: %v, Actual: %v", equalityTests, name, message, expected, actual))
		failed = append(failed, fmt.Sprintf("#%d%s", failedTests, name))
	}
	prof.AddTest(fmt.Sprintf("#%d%s", equalityTests, name))
}

// AssertNotEquals checks non-equality of values
func (et *EqualityTest[T]) AssertNotEquals(actual, expected T, name, message string) {

	equalityTests++
	if actual != expected {
		logger.Info(fmt.Sprintf("#%d%s No EqualityTest PASSED", equalityTests, name))
	} else {
		failedTests++
		logger.Warning(fmt.Sprintf("#%d%s No EqualityTest FAILED: %s. Expected: %v, Actual: %v", equalityTests, name, message, expected, actual))
		failed = append(failed, fmt.Sprintf("#%d%s", failedTests, name))
	}
	prof.AddTest(fmt.Sprintf("#%d%s", equalityTests, name))
}
