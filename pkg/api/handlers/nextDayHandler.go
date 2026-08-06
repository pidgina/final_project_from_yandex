package handlers

import (
	"errors"
	"net/http"
	"time"

	"proj/pkg/api/service"
)

func NextDayHandler(w http.ResponseWriter, r *http.Request) {
	now := r.FormValue("now")
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")

	time, err := time.Parse(service.LayoutDate, now)
	if err != nil {
		SendErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := service.NextDate(time, date, repeat)
	if err != nil {
		if errors.Is(err, service.ErrYear) || errors.Is(err, service.ErrMounth) || errors.Is(err, service.ErrWeek) || errors.Is(err, service.ErrDay) || errors.Is(err, service.ErrRepeat) {
			SendErrorJSON(w, err.Error(), http.StatusBadRequest)
		} else {
			SendErrorJSON(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Write([]byte(res))

}
