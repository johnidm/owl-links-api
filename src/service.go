package main

import (
	"./database"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"

	"net/http"
)

var db = database.OpenDB()

func main() {

	router := httprouter.New()

	router.GET("/links", GetLinks)
	router.GET("/link/:id", GetLink)
	router.DELETE("/link/:id", DeleteLink)
	router.PUT("/link/", PutLink)
	router.POST("/link", PostLink)

	defer db.Close()

	http.ListenAndServe(":8000", router)

	/*
		err := CreateLink(db, &Link{1, "www.johni.com", "Web Site", "Meu web site"})
		if err != nil {
			panic(err)
		}
	*/
	/*
				err := database.UpdateLink(db, &database.Link{1, "www.douglas.com", "E-Commerce", "Meu E-Commerce"})
				if err != nil {
					panic(err)
				}
		dbdb
	*/
	/*
		err := DeleteLink(db, 1)
		if err != nil {
			panic(err)
		}
	*/
	/*


		fmt.Printf("Total de linhas %d\n", len(links))
		if len(links) != 0 {
			fmt.Println(links[0].Url)
		}
	*/

	//link, err := GetLink(db, 1)
	//fmt.Printf("URL", link.Url)

}

func GetLinks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	links, err := database.GetLinks(db)
	if err != nil {
		panic(err)
	}

	fmt.Println(links)
	fmt.Fprint(w, "Lista!\n")
	js, _ := json.Marshal(links)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func PutLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Update!\n")
}

func PostLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Save!\n")
}

func GetLink(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Get, %s!\n", p.ByName("id"))
}

func DeleteLink(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "Delete, %s!\n", p.ByName("id"))
}
