package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func GetTime(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	dateFormat := "2006-01-02 15:04:05 -0700 MST"
	jsonData := make(map[string]string)
	if !vars.Has("tz") {
		jsonData["current_time"] = time.Now().Format(dateFormat)
		_ = json.NewEncoder(w).Encode(jsonData)
		return
	}
	for _, tz := range strings.Split(vars.Get("tz"), ",") {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(fmt.Sprintf("invalid timezone %v", tz))
			return
		}
		jsonData[tz] = time.Now().In(loc).Format(dateFormat)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}
