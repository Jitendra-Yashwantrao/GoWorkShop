package mymodels

import (
	"math/rand"
	"sort"
	"testing"
)

var customColor = []string{"Red", "Green", "Blue"}

func GetCircles(count int) []Circle {

	cirlces := []Circle{}

	for i := 0; i < count; i++ {

		cirlces = append(cirlces, Circle{
			Radius: float64(rand.Intn(10)),
			Color:  customColor[rand.Intn(len(customColor))],
		})
	}
	return cirlces
}

func GetRectangles(count int) []Rectangle {

	rectangles := []Rectangle{}

	for i := 0; i < count; i++ {

		rectangles = append(rectangles, Rectangle{
			Length:  float64(rand.Intn(10)),
			Breadth: float64(rand.Intn(10)),
			Color:   customColor[rand.Intn(len(customColor))],
		})
	}
	return rectangles
}

func TestSortRectangle1(t *testing.T) {

	rectangles := GetRectangles(5)
	sortedRectangles := SortRectangleByAreaAndColor(rectangles)
	t.Log(rectangles)

	sort.Sort(sortedRectangles)

	t.Log(sortedRectangles)
	//sortedRectangles = append(sortedRectangles, Rectangle{Length: 0.0, Breadth: 1.0, Color: "Red"})
	if !sort.IsSorted(sortedRectangles) {
		t.Error("Sorting failed")
	}
	//fmt.Println("sorted output")

	//log.Println(rectangles)

}

func TestSortCircles1(t *testing.T) {

	circles := GetCircles(5)
	sortedCircles := SortCircleByAreaAndColor(circles)
	t.Log(circles)

	sort.Sort(sortedCircles)

	t.Log(sortedCircles)
	//sortedRectangles = append(sortedRectangles, Rectangle{Length: 0.0, Breadth: 1.0, Color: "Red"})
	if !sort.IsSorted(sortedCircles) {
		t.Error("Sorting failed")
	}
	//fmt.Println("sorted output")

	//log.Println(rectangles)

}
