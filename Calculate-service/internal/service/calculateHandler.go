package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func isValidRequestBody(data *map[string]string) bool {
	if len(*data) != 1 {
		return false
	}

	expression, exists := (*data)["expression"]
	if !exists || expression == "" {
		return false
	}

	operators := []string{"+", "-", "*", "/"}
	for _, op := range operators {
		if strings.Contains(expression, op) {
			return true
		}
	}
	return false
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data map[string]string

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if !isValidRequestBody(&data) {
		http.Error(w, "Expression is not valid", http.StatusUnprocessableEntity)
		return
	}

	expression := data["expression"]
	result, err1 := Calc(expression)
	if err1 != nil {
		http.Error(w, "No format expression", http.StatusBadRequest)
		return
	}

	response := map[string]string{"result": fmt.Sprint(result)}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
