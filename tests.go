package main

import (
	"fmt"
	"gut/gut"
)

func main() {
	testRunner := gut.NewTestRunner(gut.KB)

	testRunner.AddTests(
		func() {
			et := gut.EqualityTest[int]{}
			et.AssertEquals(1, 2, "Equals", "Значения должны быть равны")
			et.AssertNotEquals(1, 2, "Equals", "Значения должны быть не равны")
			et.AssertEquals(1, 1, "Equals", "Значения должны быть равны")
			et.AssertNotEquals(1, 1, "Equals", "Значения должны быть не равны")
		},
		func() {
			bt := gut.BoundaryTest[gut.BoundaryInt]{}
			bt.AssertGreater(1, 2, "Boundary", "Значения должны быть равны")
			bt.AssertGreaterOrEqual(1, 2, "Boundary", "Значения должны быть равны")
			bt.AssertLess(1, 2, "Boundary", "Значения должны быть равны")
			bt.AssertLessOrEqual(1, 2, "Boundary", "Значения должны быть равны")
			bt.AssertGreater(2, 1, "Boundary", "Значения должны быть равны")
			bt.AssertGreaterOrEqual(2, 1, "Boundary", "Значения должны быть равны")
			bt.AssertLess(2, 1, "Boundary", "Значения должны быть равны")
			bt.AssertLessOrEqual(2, 1, "Boundary", "Значения должны быть равны")
		})

	testRunner.RunTests()

	_, _ = fmt.Scanf("%s")
}
