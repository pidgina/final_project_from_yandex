package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"proj/pkg/db"
)

func GetIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		log.Println("Попытка получить данные по пустому ID")
		SendErrorJSON(w, "Поле id не может быть пустым", http.StatusBadRequest)
		return
	}

	res, err := db.GetTaskID(id)
	if err != nil {
		log.Println(err.Error())
		SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	byteResp, err := json.Marshal(res)
	if err != nil {
		SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendOkJSONBytes(w, http.StatusOK, byteResp)

}
