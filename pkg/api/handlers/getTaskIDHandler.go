package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"proj/pkg/api/service"
)

func GetIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		log.Println("Попытка получить данные по пустому ID")
		service.SendErrorJSON(w, "Поле id не может быть пустым", http.StatusBadRequest)
		return
	}

	res, err := service.GetTaskID(id)
	if err != nil {
		log.Println(err.Error())
		service.SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	byteResp, err := json.Marshal(res)
	if err != nil {
		service.SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	service.SendOkJSONBytes(w, http.StatusOK, byteResp)

}
