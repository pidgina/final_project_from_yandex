package handlers

import (
	"errors"
	"net/http"

	"proj/pkg/api/service"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		service.SendErrorJSON(w, "Поле id не может быть пустым", http.StatusBadRequest)
		return
	}

	byteResp, err := service.DeleteTask(id)
	if err != nil {
		if errors.Is(err, service.NoneID) {
			service.SendErrorJSON(w, err.Error(), http.StatusBadRequest)
			return
		}
		service.SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	service.SendOkJSONBytes(w, http.StatusOK, byteResp)

}
