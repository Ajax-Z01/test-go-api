package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"test-go-api/config"
	"test-go-api/entity"
	"test-go-api/repository"
	"test-go-api/response"

	"github.com/gorilla/mux"
)

func GetDataUser(w http.ResponseWriter, r *http.Request) {
	db, err := config.MySQLConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	listUser, err := userRepository.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	byteJson, err := response.MapResponseListUser(listUser)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteJson)
}

func GetDataByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userIDStr := mux.Vars(r)["id"]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.MySQLConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	user, err := repo.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.MySQLConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	err = repo.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	byteJson, err := response.MapResponseUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(byteJson)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.MySQLConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user.ID = userID

	repo := repository.NewUserRepository(db)
	if err := repo.UpdateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	byteJson, err := response.MapResponseUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(byteJson)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := config.MySQLConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	err = repo.DeleteUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
