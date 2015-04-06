package utils

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	APIKey = "4B5B8A4948F8AA4FF918A353B5CAE"
)

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

func DefineReturnRequestIdInvalid(w http.ResponseWriter, e error) {
	DefineReturnErrorHttpRequest(
		w,
		e.Error(),
		"É necessário um 'id' válido no parâmetro da requisição",
		500)
}

func DefineReturnRequestFailExecFunc(w http.ResponseWriter, e error, f string) {
	DefineReturnErrorHttpRequest(
		w,
		e.Error(),
		"Erro ao executar a função "+f,
		500)
}

func DefineReturnRequestError(w http.ResponseWriter, e error, m string) {
	DefineReturnErrorHttpRequest(
		w,
		e.Error(),
		m,
		500)
}

func WriteJson(w http.ResponseWriter, j []byte) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.Write(j)
	w.WriteHeader(http.StatusOK)

}

func DefaultHeader(w http.ResponseWriter) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, api-key")
}

func TestRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	DefaultHeader(w)

	w.Write([]byte(r.Method + " executed..."))
}
