package handlers

import (
	"errors"
	"net/http"

	"proj/pkg/db"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		SendErrorJSON(w, "Поле id не может быть пустым", http.StatusBadRequest)
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		if errors.Is(err, db.NoneID) {
			SendErrorJSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendOkJSONBytes(w, http.StatusOK, []byte("{}"))
}
