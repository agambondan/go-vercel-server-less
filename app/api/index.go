package api

import (
	"encoding/json"
	"net/http"

	"github.com/agambondan/go-vercel-server-less/app/model"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&model.Response{
		Code:   200,
		Status: "success",
		Data:   map[string]interface{}{"username": "agam", "password": "agam"},
	})
}
