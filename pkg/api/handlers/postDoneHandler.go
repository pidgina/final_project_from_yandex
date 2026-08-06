package handlers

import (
	"net/http"

	"proj/pkg/api/service"
)

func DonePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		service.SendErrorJSON(w, "Поле id не может быть пустым", http.StatusBadRequest)
		return
	}

	byteResp, err := service.PostDone(id)
	if err != nil {
		service.SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	service.SendOkJSONBytes(w, http.StatusOK, byteResp)

}
