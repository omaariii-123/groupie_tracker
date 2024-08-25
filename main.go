package main

import (
		"encoding/json"
		"net/http"
	//	"fmt"
		"html/template"
		//"io"
	)

type Data struct {
	Artists   []Artist 
	Locations Location
	Dates     Date
//	Relations  []Relation
}

type Artist struct {
	Id           int      `json:id`
	Image		 string	  `json:image`
	Name         string   `json:name`
	Members      []string `json:members`
	CreationDate int      `json:creationDate`
	FirstAlbum   string   `json:firstAlbum`
}
type Location struct {
		Index []interface{} `json:index`
}
type Date struct {
	Index []interface{} `json:index`
}

type artists1 struct {
	Id	int				`json:name`
	Name string			`json:name`
	Image string		`json:image`
}


func mainfunc(w http.ResponseWriter, r *http.Request) {
	var artist []artists1 
	
	jsonResp, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	
	decoder := json.NewDecoder(jsonResp.Body)
	decoder.Decode(&artist)
	tmpl, _ := template.ParseFiles("1index.html")
	tmpl.Execute(w, artist)

}

func funct(w http.ResponseWriter, r *http.Request) {

	var artist []Artist
	var location Location
	var date Date
	var data Data
	
	jsonRespA, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	jsonRespL, _ := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	jsonRespD, _ := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	//jsonRespR, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	//tmp, _ := io.ReadAll(jsonResp.Body)
	//fmt.Println(string(tmp)
	//	defer jsonResp.Body.Close() 
	//json.Unmarshal(tmp, &artist)
	decoder := json.NewDecoder(jsonRespA.Body)
	decoder.Decode(&artist)
	decoder = json.NewDecoder(jsonRespL.Body)
	decoder.Decode(&location)
	decoder = json.NewDecoder(jsonRespD.Body)
	decoder.Decode(&date)
	//decoder := json.NewDecoder(jsonRespA.Body)
	//decoder.Decode(&artist)
	//fmt.Fprintln(w, artist[0].Image)
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		println(err)
	}
	data.Artists = artist
	data.Locations = location
	data.Dates = date
	tmpl.Execute(w, data)

}

func main(){
	http.HandleFunc("/", mainfunc)
	http.HandleFunc("/{id}", funct)
	http.ListenAndServe(":8080", nil)
}
