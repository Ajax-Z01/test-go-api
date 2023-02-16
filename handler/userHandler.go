package handler

import (
	"net/http"
	"test-go-api/repository"
	"test-go-api/response"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" /* http.MethodGet */ {
		w.WriteHeader(400)
		w.Write([]byte("Method not allowed."))

		return
	}

	user, err := repository.GetDataUser()

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))

		return
	}

	dataJson, errJson := response.MapResponseListUser(user)

	if errJson != nil {
		w.WriteHeader(500)
		w.Write([]byte(errJson.Error()))

		return
	}

	w.Write(dataJson)
}
