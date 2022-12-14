package main

import (
	utils "assignment-3/appUtils"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func AutoUpdateData() {
	for {
		body := utils.GenerateRandom()
		utils.JsonDump(body, utils.JSONFile)
		time.Sleep(utils.SleepTime * time.Second)
	}
}

func DisplayHTML(w http.ResponseWriter, r *http.Request) {
	body := utils.JsonLoad(utils.JSONFile)
	page, err := template.ParseFiles(utils.IndexHTML)
	if err != nil {
		log.Fatalf("error parse from [ %s ]", utils.IndexHTML)
	}
	data := utils.ParseBody(body)
	page.Execute(w, data)
}

func main() {
	go AutoUpdateData()
	http.HandleFunc("/", DisplayHTML)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
