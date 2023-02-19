package handler

import (
	"net/http"
	"test-go-api/repository"
	"test-go-api/response"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {

	user, err := repository.GetDataUser()
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	dataJson, errJson := response.MapResponseListUser(user)

	if errJson != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(errJson.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, _ = w.Write(dataJson)
}
