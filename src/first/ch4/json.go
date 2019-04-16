package ch4

import (
	"encoding/json"
)

type Movie struct {
	Title  string
	yr     int  `json:"released"` // if it is unexported, then it is not marshaled
	// Yr int `json:"released"` // try this instead to see that "Yr" gets exported as "released" in JSON
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", yr: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", yr: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", yr: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func Marshal() ([]byte, error) {
	return json.Marshal(movies)
}