package main

import "testing"

func TestPerimeter(t *testing.T) {
	expected := 40.0
	actual := Perimeter(Rectangle{10.0, 10.0})
	if expected != actual {
		t.Errorf("Expected : %.2f, but Actual : %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {
	// Helper function
	// Asking for a shape to be passed in
	checkArea := func (t *testing.T, shape Shape, expected float64) {
		// To resolve tracking down problem
		t.Helper()
		actual := shape.Area()
		if expected != actual {
			t.Errorf("Expected : %.2f, but Actual : %.2f", expected, actual)
		}
	}

	t.Run("Should enable to calculate Rectangle type", func(t *testing.T) {
		rect := Rectangle{12, 6}
		expected := 72.0
		checkArea(t, rect, expected)
	})

	t.Run("Should enable to calculate Circle type", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})

	// Refactoring : Table Driven Test(Useful when you want to build a list of test cases that
	// can be tested in same manner.
	areaTestCases := [] struct{
		name string
		shape Shape
		expected float64
	}{
		{"Test Rectangle",Rectangle{12,6}, 72.0},
		{"Test Circle",Circle{10},314.1592653589793},
		{"Test Triangle",Triangle{12, 6}, 36.1},
	}
	for _, testCase := range areaTestCases{
		// Should Test speaks us to more clearly, not a seq of operation.
		// Make our test output more helpful by wrapping each case in a t.Run()
		// so that it will print the test case if the test fails.
		t.Run(testCase.name, func(t *testing.T) {
			expected := testCase.expected
			actual := testCase.shape.Area()
			if expected != actual {
				t.Errorf("%#v Expected : %.2f, but Actual : %.2f", testCase.shape, expected, actual)
			}
		})
	}
}
