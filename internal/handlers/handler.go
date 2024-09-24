package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Awes0meEman/shopping-time/internal/models"
	"github.com/Awes0meEman/shopping-time/internal/repository"
	"github.com/gorilla/mux"
)

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var game models.Game
	_ = json.NewDecoder(r.Body).Decode(&game)
	_, err := repository.CreateGame(game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(game)
}

func GetAllGames(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(repository.GetAllGames())
}

func GetGameById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Game ID", http.StatusBadRequest)
		return
	}

	game, err := repository.GetGameById(gameId)

	if err != nil {
		http.NotFound(w, r)
		return
	} else {
		json.NewEncoder(w).Encode(game)
	}
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Game ID", http.StatusBadRequest)
		return
	}

	var updatedGame models.Game
	_ = json.NewDecoder(r.Body).Decode(&updatedGame)
	updatedGame.Id = gameId

	game, err := repository.UpdateGameById(updatedGame)

	if err != nil {
		http.NotFound(w, r)
	} else {
		json.NewEncoder(w).Encode(game)
	}
}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Game ID", http.StatusBadRequest)
		return
	}

	err = repository.DeleteGame(gameId)

	if err != nil {
		http.NotFound(w, r)
	} else {
		fmt.Fprintf(w, "Game with Id %d is deleted", gameId)
	}
}
