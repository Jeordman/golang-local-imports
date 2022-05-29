package api;

import (
	"log"
	"net/http"
	"local-imports-proj/api/controllers"
)

func StartApi() {
	http.HandleFunc("/", articleController.RootCtrl)
	http.HandleFunc("/articles", articleController.AllArticles)
	log.Fatal(http.ListenAndServe(":8000", nil)) 
}
