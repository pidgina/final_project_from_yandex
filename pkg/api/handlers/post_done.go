package handlers

import (
	"net/http"

	"proj/pkg/api/service"
)

func DonePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		SendErrorJSON(w, "Поле id не может быть пустым", http.StatusBadRequest)
		return
	}

	err := service.PostDone(id)
	if err != nil {
		SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SendOkJSONBytes(w, http.StatusOK, []byte("{}"))

}
