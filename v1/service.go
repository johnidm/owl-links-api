package main

import (
	"encoding/json"
	"fmt"
	"github.com/johnidm/owl-links-api/database"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"flag"
	"log"
)

const (
	APIKey = "4B5B8A4948F8AA4FF918A353B5CAE"
)

var db = database.OpenDB()

func main() {

	port := flag.String("port", "8000", "HTTP Port")
	flag.Parse()

	defer db.Close()

	router := httprouter.New()

	router.GET("/", RunProject)

	router.GET("/links", GetLinks)
	router.GET("/link/:id", GetLink)
	router.DELETE("/link/:id", DeleteLink)
	router.PUT("/link/", PutLink)
	router.POST("/link", PostLink)

	log.Println("Stating Server on ", *port)

	log.Fatal(http.ListenAndServe(":" + *port, router))

}

func RunProject(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Write([]byte("<h2><font color=\"green\">Owl Link API is running!</font></h2>"))
}

func GetLinks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if !APIKeyIsValid(w, r) {
		return
	}

	links, err := database.GetLinks(db)
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(GetLinks) - Erro ao executar a função database.GetLinks",
			500)
		return
	}

	js, err := json.Marshal(links)
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(GetLinks) - Erro ao fazer o marshal dos links",
			500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(js)
}

func PutLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if !APIKeyIsValid(w, r) {
		return
	}

	var link database.Link

	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(PutLink) - Erro ao ler o conteúdo da requisição",
			500)
		return
	}

	if err := json.Unmarshal(body, &link); err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(PutLink) - Erro ao fazer unmarshal do conteúdo da requisição",
			422)
		return
	}

	err = database.UpdateLink(db, &link)
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(PutLink) - Erro ao executar a função database.PutLink",
			500)
		return
	}
}

func PostLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if !APIKeyIsValid(w, r) {
		return
	}

	var link database.Link

	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(PostLink) - Erro ao ler o conteúdo da requisição",
			500)
		return
	}

	if err := json.Unmarshal(body, &link); err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(PostLink) - Erro ao fazer unmarshal do conteúdo da requisição",
			422)
		return
	}

	err = database.CreateLink(db, &link)
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(PostLink) - Erro ao executar a função database.PostLink",
			500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

}

func GetLink(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if !APIKeyIsValid(w, r) {
		return
	}

	id, err := strconv.ParseInt(p.ByName("id"), 0, 64)
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"Informe um \"id\" correto para excluir o registro",
			500)
		return
	}

	link, err := database.GetLink(db, id)
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(GetLink) - Erro ao executar a função database.GetLink",
			500)

		return
	}

	js, err := json.Marshal(link)
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(GetLink) - Erro ao fazer marshal do conteúdo da requisição",
			422)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(js)

	fmt.Fprintf(w, "Get, %s!\n")
}

func DeleteLink(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	if !APIKeyIsValid(w, r) {
		return
	}

	id, err := strconv.ParseInt(p.ByName("id"), 0, 64)
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"Informe um \"id\" correto para excluir o registro",
			500)
		return
	}

	err = database.DeleteLink(db, id)
	if err != nil {
		DefineReturnErrorHttpRequest(
			w,
			err.Error(),
			"(DeleteLink) - Erro ao executar a função database.DeleteLink",
			500)
		return
	}

}

func APIKeyIsValid(w http.ResponseWriter, r *http.Request) bool {

	if (r.Header.Get("api-key")) != APIKey {
		DefineReturnErrorHttpRequest(
			w,
			"API Key inválida ou não informada",
			"Verifique o valor da chave api-key do cabeçalho da requisição",
			http.StatusUnauthorized)
		return false
	}

	return true
}

func DefineReturnErrorHttpRequest(w http.ResponseWriter, e string, es string, s int) {

	js, _ := json.Marshal(map[string]string{"error": e, "sugestion": es})

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(s)
	w.Write(js)

}
