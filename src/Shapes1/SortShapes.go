package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	var jsontype Shape
	//json.Unmarshal(file, &jsontype)
	err := json.Unmarshal([]byte(text), &jsontype)
	if err != nil {
		panic(err)
	}

	rectangles := jsontype.Rectangles
	circles := jsontype.Circles

	fmt.Println("Input rectangles")
	fmt.Println(rectangles)

	sort.Sort(SortRectangleByAreaAndColor(rectangles))
	fmt.Println("Rectangles sorted by  area and color")
	fmt.Println(rectangles)

	fmt.Println("Input cricles")
	fmt.Println(circles)

	sort.Sort(SortCircleByAreaAndColor(circles))
	fmt.Println("Circles sorted by area and color")
	fmt.Println(circles)
}
