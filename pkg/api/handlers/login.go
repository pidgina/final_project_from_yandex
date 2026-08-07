package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"proj/pkg/api/service"
)

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	var pass service.PasswordType

	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Попытка передачи некорректного запроса в body")
		SendErrorJSON(w, "Передано некорректное значение в запросе", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bodyByte, &pass)
	if err != nil {
		SendErrorJSON(w, "Ошибка десериализации JSON", http.StatusInternalServerError)
		return
	}

	resp, err := service.LoginAuther(pass)
	if err != nil {
		if errors.Is(err, service.InvalidPassword) {
			SendErrorJSON(w, err.Error(), http.StatusUnauthorized)
			return
		} else {
			SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}

	byteResp, err := json.Marshal(resp)
	if err != nil {
		SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendOkJSONBytes(w, http.StatusOK, byteResp)

}
