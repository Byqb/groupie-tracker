package main

import (
	"fmt"
	"net/http"

	function "GT/function"
)

func main() {
	function.UnmarshalData()
	http.HandleFunc("/", function.Handler)
	fmt.Println("Server started on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

// Fetch data from the API
// fetchArtistData()
// fetchLocationData()
// fetchDateData()
// fetchRelationData()

/*
func fetchArtistData() {
	// Fetch artists data
	// ... (same as before)
}

func fetchLocationData() {
	// Fetch location data
	// ... (same as before)
}

func fetchDateData() {
	// Fetch date data
	// ... (same as before)
}

func fetchRelationData() {
	// Fetch relation data
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error fetching relation data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading relation data:", err)
		return
	}

	err = json.Unmarshal(body, &relations)
	if err != nil {
		fmt.Println("Error unmarshaling relation data:", err)
		return
	}
}


	data := struct {
		Artists   []Artist
		Locations []Location
		Dates     []Date
		Relations []Relation
	}{
		Artists:   artists,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	// Render the artist details page template with artist, location, date, and relation data
	// ...
}




package GT

type PageDataArtice struct {
	All                    Artist
	MergeDatesAndLocations ResponseData3
}

type ErrorPage struct{
	Message  string
}

type PageData struct {
	All []Artist
	LocFLT []string
	NumberOfMember []string
}

type Artist struct {
	ID           int         `json:"id"`
	Image        string      `json:"image"`
	Name         string      `json:"name"`
	Members      []string    `json:"members"`
	CreationDate interface{} `json:"creationDate"`
	FirstAlbum   string      `json:"firstAlbum"`
}

type ResponseData struct {
	Index []Index `json:"index"`
}

type Index struct {
	ID      int      `json:"id"`
	TheData []string `json:"locations"`
}

type ResponseData2 struct {
	Index2 []Index2 `json:"index"`
}

type Index2 struct {
	ID      int      `json:"id"`
	TheData []string `json:"Dates"`
}

type ResponseData3 struct {
	Index3 []Index3
}

type Index3 struct {
	Location string
	TheData  string
}

type ResponseData4 struct {
	Index4 []Index4 `json:"index"`
}

type Index4 struct {
	ID      int         `json:"id"`
	TheData interface{} `json:"datesLocations"`
}

var artists []Artist
var TheLocations ResponseData
var TheDates ResponseData2
var TheRelations ResponseData4
var flag bool
var LoctionList []string

*/
