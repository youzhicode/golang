package main

import (
	"fmt"
	"log"
	"encoding/json"
)

type Movie struct {
	Title 	string
	Year 	int `json:"released"`
	Color	bool `json:"color,omitempty"`
	Actors  []string
}


func main() {
	var movies = []Movie{
		{
			Title:"Casablanca", Year: 1942,Color:false,
			Actors:[]string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title: "Cool Hand Luke", Year: 1967, Color:true,
			Actors:[]string{"Paul Newman"},
		},
		{
			Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
		},
	}
	
	data, err := json.Marshal(movies)
	if err != nil {
//		fmt.Fprintf(os.Stderr, "Have error %v\n", err)
		log.Fatal("Json marshaling failed: %s\n", err)
	}
	fmt.Printf("%s\n", data)
	
	
	var titles []struct{ Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed:%s\n", err)
	}
	fmt.Println(titles)
	
}

