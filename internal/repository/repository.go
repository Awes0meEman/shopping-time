package repository

import (
	"github.com/Awes0meEman/shopping-time/internal/models"
	"fmt"
)

var games []models.Game

func GetGameById(id int) (models.Game, error) {
	for _, game := range games {
		if game.Id == id{
			return game, nil
		}
	}
	return *&models.Game{}, fmt.Errorf("No game with Id {} was found")
}
