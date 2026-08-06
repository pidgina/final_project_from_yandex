package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"proj/pkg/api/service"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")

	tasks, err := service.GetListTask(service.Limit, search)
	if err != nil {
		log.Println(err.Error())
		service.SendErrorJSON(w, "Не удалось получить задачи", http.StatusInternalServerError)
		return
	}

	jsonbyte, err := json.Marshal(tasks)
	if err != nil {
		log.Println(err.Error())
		service.SendErrorJSON(w, "Не удалось подготовить тело ответа.", http.StatusInternalServerError)
		return
	}

	service.SendOkJSONBytes(w, http.StatusOK, jsonbyte)

}
