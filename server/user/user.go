package user2

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/icub3d/gorca"
	"net/http"
)

func MakeMuxer(profix string) http.Handler {
	var m *mux.Router
	if profix == "" {
		m = mux.NewRouter()
	} else {
		m = mux.NewRouter().PathPrefix(profix).Subrouter()
	}

	m.HandleFunc("/", GetUser).Methods("GET")
	m.HandleFunc("/{path:*}", gorca.NotFoundFunc)
	return m
}

type UserInfo struct {
	Email     string
	LogoutURL string
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	cat := make(map[string]string)
	cat["name"] = "Tom"
	cat["age"] = "six"
	json, _ := json.Marshal(cat)
	w.Write(json)
}
