package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"proj/pkg/db"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	tasks, err := db.GetListTask(db.Limit, search)
	if err != nil {
		log.Println(err.Error())
		SendErrorJSON(w, "Не удалось получить задачи", http.StatusInternalServerError)
		return
	}

	jsonbyte, err := json.Marshal(tasks)
	if err != nil {
		log.Println(err.Error())
		SendErrorJSON(w, "Не удалось подготовить тело ответа.", http.StatusInternalServerError)
		return
	}

	SendOkJSONBytes(w, http.StatusOK, jsonbyte)

}
