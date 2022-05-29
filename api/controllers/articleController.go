package articleController;

import (
	"fmt"
	"net/http"
	"encoding/json"
	"local-imports-proj/api/db"
)

func RootCtrl(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func AllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(db.Articles)
}
