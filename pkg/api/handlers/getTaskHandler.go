package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"proj/pkg/db"
)

type tasksResponse struct {
	Tasks []db.Task `json:"tasks"`
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	tasks, err := db.GetListTask(db.Limit, search)
	if err != nil {
		log.Println(err.Error())
		SendErrorJSON(w, "Не удалось получить задачи", http.StatusInternalServerError)
		return
	}

	jsonbyte, err := json.Marshal(tasksResponse{
		Tasks: tasks,
	})

	SendOkJSONBytes(w, http.StatusOK, jsonbyte)

}
