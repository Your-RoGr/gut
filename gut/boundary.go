package gut

import "fmt"

// BoundaryTest is a struct for testing boundary cases
type BoundaryTest[T BoundaryComparable] struct{}

// AssertLess checks if value1 is less than value2
func (bt *BoundaryTest[T]) AssertLess(value1, value2 T, name, message string) {

	boundaryTests++
	if value1.Less(value2) {
		logger.Info(fmt.Sprintf("#%d%s assert_less PASSED", boundaryTests, name))
	} else {
		failedTests++
		logger.Warning(fmt.Sprintf("#%d%s assert_less FAILED: %s. Expected %v to be less than %v", boundaryTests, name, message, value1, value2))
		failed = append(failed, fmt.Sprintf("#%d%s", boundaryTests, name))
	}
	prof.AddTest(fmt.Sprintf("#%d%s", boundaryTests, name))
}

// AssertLessOrEqual checks if value1 is less than or equal to value2
func (bt *BoundaryTest[T]) AssertLessOrEqual(value1, value2 T, name, message string) {
	boundaryTests++
	if value1.LessOrEqual(value2) {
		logger.Info(fmt.Sprintf("#%d%s assert_less_or_equal PASSED", boundaryTests, name))
	} else {
		failedTests++
		logger.Warning(fmt.Sprintf("#%d%s assert_less_or_equal FAILED: %s. Expected %v to be less than or equal to %v", boundaryTests, name, message, value1, value2))
		failed = append(failed, fmt.Sprintf("#%d%s", boundaryTests, name))
	}
	prof.AddTest(fmt.Sprintf("#%d%s", boundaryTests, name))
}

// AssertGreater checks if value1 is greater than value2
func (bt *BoundaryTest[T]) AssertGreater(value1, value2 T, name, message string) {

	boundaryTests++
	if value1.Greater(value2) {
		logger.Info(fmt.Sprintf("#%d%s assert_greater PASSED", boundaryTests, name))
	} else {
		failedTests++
		logger.Warning(fmt.Sprintf("#%d%s assert_greater FAILED: %s. Expected %v to be greater than %v", boundaryTests, name, message, value1, value2))
		failed = append(failed, fmt.Sprintf("#%d%s", boundaryTests, name))
	}
	prof.AddTest(fmt.Sprintf("#%d%s", boundaryTests, name))
}

// AssertGreaterOrEqual checks if value1 is greater than or equal to value2
func (bt *BoundaryTest[T]) AssertGreaterOrEqual(value1, value2 T, name, message string) {

	boundaryTests++
	if value1.GreaterOrEqual(value2) {
		logger.Info(fmt.Sprintf("#%d%s assert_greater_or_equal PASSED", boundaryTests, name))
	} else {
		failedTests++
		logger.Warning(fmt.Sprintf("#%d%s assert_greater_or_equal FAILED: %s. Expected %v to be greater than or equal to %v", boundaryTests, name, message, value1, value2))
		failed = append(failed, fmt.Sprintf("#%d%s", boundaryTests, name))
	}
	prof.AddTest(fmt.Sprintf("#%d%s", boundaryTests, name))
}
