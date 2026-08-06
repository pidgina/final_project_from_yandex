package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"proj/pkg/api/service"
	"proj/pkg/db"
)

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {

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
		SendErrorJSON(w, "Ошибка при проверке параметров.", http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(task)
	if err != nil {
		SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type response struct {
		ID string `json:"id"`
	}
	var resp response

	resp.ID = strconv.Itoa(int(id))

	jsByte, err := json.Marshal(resp)
	if err != nil {
		SendErrorJSON(w, "Ошибка подготовки ответа", http.StatusInternalServerError)
		return
	}

	SendOkJSONBytes(w, http.StatusCreated, jsByte)

}
