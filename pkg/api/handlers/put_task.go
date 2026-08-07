package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"proj/pkg/api/service"
	"proj/pkg/db"
)

func PutTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task

	byteBody, err := io.ReadAll(r.Body)
	if err != nil {
		SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(byteBody, &task)
	if err != nil {
		SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if task.Title == "" {
		SendErrorJSON(w, "Не указан заголовок задачи", http.StatusBadRequest)
		return
	}

	err = service.CheckDate(&task)
	if err != nil {
		SendErrorJSON(w, "Ошибка при проверке параметров", http.StatusBadRequest)
		return
	}

	err = db.UpdateTask(task)
	if err != nil {
		SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SendOkJSONBytes(w, http.StatusOK, []byte("{}"))

}
