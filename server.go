package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mandelbrotAPI/mandelbrot"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

/* routes

now defunct due to cors and modifications
/point body: [{real: float64, imaginary: float64}] -> [bool...]

specify a canvas
/img
params:
  specifications of the canvas size
canvasWidth: int (in pixels)
canvasHeight: int (in pixels)
  specifications of the coordinates of the complex plane
planeCoordinates: [Float64....] (top, right, bottom, left)
  other specs
maxIterations: int
*/

// type complexPoint struct {
//   real, imaginary float64
// }

func readBody(r *http.Request) ([]byte, error) {
	b := []byte{}
	for {
		c := make([]byte, 8)
		read, err := r.Body.Read(c)
		for i := 0; i < read; i++ {
			b = append(b, c[i])
		}
		if err == io.EOF {
			break
		}

	}
	return b, nil
}

func point(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("request: ", r)
	//
	// //read our r.Body and parse the json into points
	// var points interface{}
	// b, _ := readBody(r)
	// fmt.Println("body: ", string(b))
	// err := json.Unmarshal(b, &points)
	// if err != nil {
	// 	fmt.Println("error in json")
	// }
	// //convert points into []complex
	// nums := make([]complex128, len(points.([]interface{})))
	// //...turned rather complex when interface{}'s came into the picture
	// for i, point := range points.([]interface{}) {
	// 	p := point.(map[string]interface{})
	// 	var real, imaginary float64
	// 	//refactor this later
	// 	switch p["real"].(type) {
	// 	case float64:
	// 		real = p["real"].(float64)
	// 	case string:
	// 		real, _ = strconv.ParseFloat(p["real"].(string), 64)
	// 	default:
	// 		fmt.Println("json parsing is all fucked")
	// 	}
	// 	switch p["imaginary"].(type) {
	// 	case float64:
	// 		imaginary = p["imaginary"].(float64)
	// 	case string:
	// 		imaginary, _ = strconv.ParseFloat(p["imaginary"].(string), 64)
	// 	default:
	// 		fmt.Println("json parsing is all fucked")
	// 	}
	// 	nums[i] = complex(real, imaginary)
	// }
	// fmt.Println("numbers: ", nums)
	// //check these numbers against our mandelbrot setup
	// answer := make([]bool, len(nums))
	// for i := 0; i < len(nums); i++ {
	// 	answer[i] = mandelbrot.IsMandelbrot(nums[i])
	// }
	// fmt.Println("answer: ", answer)
	// b, err = json.Marshal(answer)
	// w.Write(b)
}

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Fprintf(w, "Hello world")

	//r.Body implements ioReader
	str, _ := readBody(r)
	fmt.Println(str)
	//parse some json here and we're good to go
}
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("running root")
	fmt.Fprintf(w, "Hello everyone and welcome to my snazzy mandelbrot set api")
}

//do it with a get request
/*canvasWidth: int (in pixels)
canvasHeight: int (in pixels)
  specifications of the coordinates of the complex plane
planeCoordinates: [Float64....] (top, right, bottom, left)
  other specs
maxIterations: int
*/
type params struct {
	canvasWidth, canvasHeight int
	planeCoordinates          [4]float64
	maxIterations             int
}

// parse some url query values, and set some default values
func (p *params) set(q url.Values) {
	w, ok := q["canvasWidth"]
	if !ok {
		p.canvasWidth = 10
	} else {
		i, err := strconv.Atoi(w[0])
		if err != nil {
			fmt.Println(err)
		}
		p.canvasWidth = i
	}
	h, ok := q["canvasHeight"]
	if !ok {
		p.canvasHeight = 5
	} else {
		i, err := strconv.Atoi(h[0])
		if err != nil {
			fmt.Println(err)
		}
		p.canvasHeight = i
	}
	coord, ok := q["planeCoordinates"]
	if !ok {
		p.planeCoordinates = [4]float64{1, 1, -1, -2}
	} else {
		arr := [4]float64{}
		_, err := fmt.Sscanf(coord[0], "%f,%f,%f,%f", &arr[0], &arr[1], &arr[2], &arr[3])
		if err != nil {
			fmt.Println(err)
		}
		p.planeCoordinates = arr
	}
	iter, ok := q["maxIterations"]
	if !ok {
		p.maxIterations = 80
	} else {
		i, err := strconv.Atoi(iter[0])
		if err != nil {
			fmt.Println(err)
		}
		p.maxIterations = i
	}
}
func coordinates(w http.ResponseWriter, r *http.Request) {
	//fuck CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//parse the query params
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("running coordinates")
	p := params{}
	p.set(r.Form)
	fmt.Println("p = ", p)
	//ship it to mandelbrot
	answer := mandelbrot.Coordinates(p.canvasWidth, p.canvasHeight, p.planeCoordinates, p.maxIterations)
	fmt.Print(answer)
	//write our answer as a json response
	j, _ := json.Marshal(answer)
	fmt.Print(j)
	w.Write(j)
}
func img(w http.ResponseWriter, r *http.Request) {
	//fuck CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//parse the query params
	r.ParseForm()
	// fmt.Println(r.Form)
	p := params{}
	p.set(r.Form)
	fmt.Println("p = ", p)
	//ship it to mandelbrot
	answer := mandelbrot.Img(p.canvasWidth, p.canvasHeight, p.planeCoordinates, p.maxIterations)
	//write our answer as a json response
	j, _ := json.Marshal(answer)
	w.Write(j)
}

func imgTest(w http.ResponseWriter, r *http.Request) {
	//fuck CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	canvasWidth := 5
	canvasHeight := 5
	planeCoordinates := [4]float64{1, 1, 0, 0}
	maxIterations := 100
	//ship it to mandelbrot
	answer := mandelbrot.Img(canvasWidth, canvasHeight, planeCoordinates, maxIterations)
	fmt.Println(answer)
	//write our answer as a json response
	j, _ := json.Marshal(answer)
	w.Write(j)
}

const (
	production = false
	devPort    = 8080
	apiVersion = 1
)

func main() {
	// fmt.Println(mandelbrot.IsMandelbrot(1.00))
	// https://mandelbrot-api.herokuapp.com/api/img?canvasWidth=907&canvasHeight=578&maxIterations=100&planeCoordinates=1,1,-1,-2

	//routing
	//root
	// http.HandleFunc("/", root)
	//points router
	http.HandleFunc("/api/points", point)
	//api router
	http.HandleFunc("/test", api)
	//query testing
	http.HandleFunc("/api/img", img)
	//query testing
	http.HandleFunc("/api/imgTest", imgTest)
	//coordinates
	http.HandleFunc("/api/coordinates", coordinates)

	//server
	//if error log and exit
	if !production {
		fmt.Printf("Running on port :%d", devPort)
		portStr := fmt.Sprintf(":%d", devPort)
		log.Fatal(http.ListenAndServe(portStr, nil))
	} else {
		port := os.Getenv("PORT")
		if port == "" {
			log.Fatal("$PORT must be set")
		}
		fmt.Println("Running production")
		portStr := ":" + port
		log.Fatal(http.ListenAndServe(portStr, nil))
	}

}
