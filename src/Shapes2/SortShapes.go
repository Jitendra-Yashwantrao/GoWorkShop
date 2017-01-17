package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

//Rectangle Shape rectangle
type Rectangle struct {
	Breadth float64 `json:"breadth"`
	Color   string  `json:"color"`
	Length  float64 `json:"length"`
}

//Circle Shape circle
type Circle struct {
	Color  string  `json:"color"`
	Radius float64 `json:"radius"`
}

//Shape
type Shape struct {
	Circles    []Circle    `json:"circles"`
	Rectangles []Rectangle `json:"rectangles"`
}

// SortRectangleByArea
type SortRectangleByAreaAndColor []Rectangle

//Len
func (a SortRectangleByAreaAndColor) Len() int {
	return len(a)
}

//Swap
func (a SortRectangleByAreaAndColor) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//SortRectangleByAreaAndColor

//Less
func (a SortRectangleByAreaAndColor) Less(i, j int) bool {

	if a[i].Length*a[i].Breadth == a[j].Length*a[j].Breadth {
		return a[i].Color < a[j].Color
	}
	return a[i].Length*a[i].Breadth < a[j].Length*a[j].Breadth
}

func (p Rectangle) String() string {
	return fmt.Sprintf("Color %s: Length %f Breadth %f  \n", p.Color, p.Length, p.Breadth)
}

//SortCircleByArea sds
type SortCircleByAreaAndColor []Circle

func (a SortCircleByAreaAndColor) Len() int {

	return len(a)
}
func (a SortCircleByAreaAndColor) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

//SortCircleByAreaAndColor
func (a SortCircleByAreaAndColor) Less(i, j int) bool {
	if a[i].Radius*a[i].Radius == a[j].Radius*a[j].Radius {
		return a[i].Color < a[j].Color
	}

	return a[i].Radius*a[i].Radius < a[j].Radius*a[j].Radius
}

func (p Circle) String() string {
	return fmt.Sprintf("Color %s: Radius %f \n", p.Color, p.Radius)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter File Path: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	//fmt.Println(text + "|")

	//file, e := ioutil.ReadFile("./Json/InputShapes.json")
	//fmt.Println("./Json/" + text + "|")

	file, e := ioutil.ReadFile("./Json/" + text)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	fmt.Println("Input File Content")
	fmt.Printf("%s\n", string(file))

	var jsontype Shape
	err := json.Unmarshal(file, &jsontype)

	if err != nil {
		fmt.Printf("Error in marshalling json: %v\n", err)
		panic(err)
	}

	rectangles := jsontype.Rectangles
	circles := jsontype.Circles

	sort.Sort(SortRectangleByAreaAndColor(rectangles))
	fmt.Println("Rectangles sorted by  area and color")
	fmt.Println(rectangles)

	sort.Sort(SortCircleByAreaAndColor(circles))
	fmt.Println("Circles sorted by area and color")
	fmt.Println(circles)
}
