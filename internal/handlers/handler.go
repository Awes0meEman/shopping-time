package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/Awes0meEman/shopping-time/internal/models"
)
var games []models.Game

func CreateGame(w http.ResponseWriter, r *http.Request) {
	var game models.Game
	_ = json.NewDecoder(r.Body).Decode(&game)
	game.Id = len(games) + 1
	games = append(games, game)
	json.NewEncoder(w).Encode(game)
}

func GetAllGames(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(games)
}

func GetGameById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Game ID", http.StatusBadRequest)
		return
	}

	for _, game := range games {
		if game.Id == gameId {
			json.NewEncoder(w).Encode(game)
			return
		}
	}

	http.NotFound(w, r)
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
}
