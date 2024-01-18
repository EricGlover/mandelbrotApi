package mandelbrot

import (
	"testing"
)

type CoordTestCase struct {
	top           float64
	bottom        float64
	left          float64
	right         float64
	canvasWidth   int
	canvasHeight  int
	maxIterations int
}

// testing that the end points of the coordinates are correct for
// row 0 [bottom left point , ... , bottom right point]
// row n [top left point , ... , top right point]
func TestCoordinates(t *testing.T) {
	test_case1 := CoordTestCase{
		top:           float64(1),
		bottom:        float64(-1),
		left:          float64(-2),
		right:         float64(1),
		canvasWidth:   10,
		canvasHeight:  10,
		maxIterations: 10,
	}
	test_case2 := CoordTestCase{
		top:           float64(1),
		bottom:        float64(-1),
		left:          float64(-2),
		right:         float64(1),
		canvasWidth:   11,
		canvasHeight:  10,
		maxIterations: 10,
	}
	test_case3 := CoordTestCase{
		top:           float64(1),
		bottom:        float64(-1),
		left:          float64(-2),
		right:         float64(1),
		canvasWidth:   825,
		canvasHeight:  577,
		maxIterations: 10,
	}
	test_case4 := CoordTestCase{
		top:           float64(1),
		bottom:        float64(-1),
		left:          float64(-2),
		right:         float64(1),
		canvasWidth:   900,
		canvasHeight:  900,
		maxIterations: 10,
	}
	test_case5 := CoordTestCase{
		top:           float64(1),
		bottom:        float64(-1),
		left:          float64(-2),
		right:         float64(1),
		canvasWidth:   900,
		canvasHeight:  500,
		maxIterations: 10,
	}
	test_case6 := CoordTestCase{
		top:           float64(1),
		bottom:        float64(-1),
		left:          float64(-2),
		right:         float64(1),
		canvasWidth:   900,
		canvasHeight:  525,
		maxIterations: 10,
	}
	test_case7 := CoordTestCase{
		top:           float64(1),
		bottom:        float64(-1),
		left:          float64(-2),
		right:         float64(1),
		canvasWidth:   900,
		canvasHeight:  527,
		maxIterations: 10,
	}
	test_case8 := CoordTestCase{
		top:           float64(1),
		bottom:        float64(-1),
		left:          float64(-2),
		right:         float64(1),
		canvasWidth:   901,
		canvasHeight:  525,
		maxIterations: 10,
	}
	runTestCase := func(t *testing.T, test_case CoordTestCase) bool {
		passed := true
		topLeft := complex(test_case.left, test_case.top)
		topRight := complex(test_case.right, test_case.top)
		bottomLeft := complex(test_case.left, test_case.bottom)
		bottomRight := complex(test_case.right, test_case.bottom)
		coordinates := Coordinates(
			test_case.canvasWidth,
			test_case.canvasHeight,
			[4]float64{test_case.top, test_case.right, test_case.bottom, test_case.left},
			test_case.maxIterations)
		// now we find our coordinates
		// row 0 [bottom left point , ... , bottom right point]
		// row n [top left point , ... , top right point]
		rowLength := len(coordinates[0])
		rowCount := len(coordinates)
		foundBottomLeft := coordinates[0][0]
		foundBottomRight := coordinates[0][rowLength-1]
		foundTopLeft := coordinates[rowCount-1][0]
		foundTopRight := coordinates[rowCount-1][rowLength-1]
		if foundBottomLeft != bottomLeft {
			t.Error("Failed for test case", test_case, "\n", "bottomLeft expected", bottomLeft, "bottomLeft found", foundBottomLeft)
			passed = false
		}
		if foundBottomRight != bottomRight {
			t.Error("Failed for test case", test_case, "\n", "bottomRight expected", bottomRight, "bottomRight found", foundBottomRight)
			passed = false
		}
		if foundTopLeft != topLeft {
			t.Error("Failed for test case", test_case, "\n", "topLeft expected", topLeft, "topLeft found", foundTopLeft)
			passed = false
		}
		if foundTopRight != topRight {
			t.Error("Failed for test case", test_case, "\n", "topRight expected", topRight, "topRight found", foundTopRight)
			passed = false
		}
		return passed
	}

	runTestCase(t, test_case1)
	runTestCase(t, test_case2)
	runTestCase(t, test_case3)
	runTestCase(t, test_case4)
	runTestCase(t, test_case5)
	runTestCase(t, test_case6)
	runTestCase(t, test_case7)
	runTestCase(t, test_case8)
}
