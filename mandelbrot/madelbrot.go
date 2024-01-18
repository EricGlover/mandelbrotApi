package mandelbrot

import (
	"math"
	"math/cmplx"
	"sync"
)

// things to remember, for the complex plane the x -axis is real, the y-axis is imaginary
// complex( real, imaginary ) => (real + imaginary i )

func GetBoundaries(canvasWidth int, canvasHeight int, planeCoordinates [4]float64) struct {
	top    float64
	bottom float64
	right  float64
	left   float64
	xRatio float64
	yRatio float64
} {
	top := planeCoordinates[0]
	right := planeCoordinates[1]
	bottom := planeCoordinates[2]
	left := planeCoordinates[3]
	planeWidth := right - left
	planeHeight := top - bottom
	ret :=
		struct {
			top    float64
			bottom float64
			right  float64
			left   float64
			xRatio float64
			yRatio float64
		}{
			top:    top,
			bottom: bottom,
			right:  right,
			left:   left,
			xRatio: planeWidth / float64(canvasWidth-1),
			yRatio: planeHeight / float64(canvasHeight-1),
		}
	return ret
}

func Coordinates(canvasWidth int, canvasHeight int, planeCoordinates [4]float64, maxIterations int) [][]complex128 {
	//make our slice
	pixels := make([][]complex128, canvasHeight)
	for i := range pixels {
		pixels[i] = make([]complex128, canvasWidth)
	}
	//iterate over each pixel checking the setMembership and iteration escape depth
	////set up some facts about our complex plane
	b := GetBoundaries(canvasWidth, canvasHeight, planeCoordinates)
	// the plane coordinates given should be included (so (left, bottom) && (right, top))
	// rows should be ordered like this
	// row 0 [bottom left point , ... , bottom right point]

	//using waitgroups for goroutine control
	var wg sync.WaitGroup
	wg.Add(len(pixels))
	for y := 0; y < len(pixels); y++ {
		go func(y1 int) {
			defer wg.Done()
			for x := 0; x < len(pixels[y1]); x++ {
				//convert pixels coords to complex plane coords
				c := complex((float64(x)*b.xRatio + b.left), (float64(y1)*b.yRatio + b.bottom))
				pixels[y1][x] = c
			}
		}(y)
	}
	wg.Wait()
	return pixels
}

func MakeImageFromCoordinates(coordinates [][]complex128, maxIterations int) [][]int {
	// make our new array
	escape := make([][]int, len(coordinates))
	rowLength := len(coordinates[0])
	for i := range escape {
		escape[i] = make([]int, rowLength)
	}

	// run our function over coordinates
	var wg sync.WaitGroup
	wg.Add(len(coordinates))
	for y := 0; y < len(coordinates); y++ {
		go func(y1 int) {
			defer wg.Done()
			for x := 0; x < len(coordinates[y1]); x++ {
				//check the setMembership
				escapeIter := escapeIteration(coordinates[y1][x], maxIterations)
				escape[y1][x] = escapeIter
			}
		}(y)
	}
	wg.Wait()
	return escape
}

func Img(canvasWidth int, canvasHeight int, planeCoordinates [4]float64, maxIterations int) [][]int {
	coordinates := Coordinates(canvasWidth, canvasHeight, planeCoordinates, maxIterations)
	return MakeImageFromCoordinates(coordinates, maxIterations)
}

// escapeIteration returns the iteration the c escapes set memebership
// 0 <= escape <= max, with 0 being the flag for "hasn't escaped yet"
// consider changing the flag thing
func escapeIteration(c complex128, max int) int {
	//set default max value
	if max == 0 {
		max = 80
	}
	distance := float64(2 * 2)
	z := 0 + 0i
	for i := 0; i < max; i++ {
		z = z*z + c
		r, theta := cmplx.Polar(z)
		real := math.Cos(theta) * r
		imaginary := math.Sin(theta) * r
		if real*real+imaginary*imaginary > distance {
			return i + 1
		}
	}
	return 0 //flag value
}

// IsMandelbrot tests whether or not c is in the set
func IsMandelbrot(c complex128) bool {
	maxIterations := 80
	distance := float64(2 * 2)
	z := 0 + 0i
	for i := 0; i < maxIterations; i++ {
		z = z*z + c
		r, theta := cmplx.Polar(z)
		real := math.Cos(theta) * r
		imaginary := math.Sin(theta) * r
		if real*real+imaginary*imaginary > distance {
			return false
		}
	}
	return true
}
