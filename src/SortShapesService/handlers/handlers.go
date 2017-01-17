package handlers

import (
	"SortShapesService/mymodels"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
)

//Index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to sorting shapes!")
}

func SortShapes(w http.ResponseWriter, r *http.Request) {
	var inputShapes mymodels.Shape
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &inputShapes); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	rectangles := inputShapes.Rectangles
	circles := inputShapes.Circles

	sort.Sort(mymodels.SortRectangleByAreaAndColor(rectangles))
	// fmt.Println("Rectangles sorted by  area and color")
	// fmt.Println(rectangles)

	sort.Sort(mymodels.SortCircleByAreaAndColor(circles))
	// fmt.Println("Circles sorted by area and color")
	// fmt.Println(circles)
	var sortedShapes mymodels.Shape

	sortedShapes.Circles = circles
	sortedShapes.Rectangles = rectangles
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(sortedShapes); err != nil {
		panic(err)
	}
}
