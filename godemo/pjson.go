package main

import (
	"encoding/json"
	"fmt"
)
type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string
	Description string
	Dimen Dimensions
}

func main() {
	birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks", "dimensions":{"height":24,"width":10}},
	{"species":"eagle","description":"bird of prey", "dimensions":{"height":34,"width":18}}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("Birds : %+v", birds)
	//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]
}
