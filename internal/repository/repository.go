package repository

import (
	"fmt"

	"github.com/Awes0meEman/shopping-time/internal/models"
)
//==========================Games==========================
var games []models.Game

func GetGameById(id int) (models.Game, error) {
	for _, game := range games {
		if game.Id == id{
			return game, nil
		}
	}
	return *&models.Game{}, fmt.Errorf("No game with Id %d was found", id)
}

func GetAllGames() []models.Game {
	return games
}

func CreateGame(g models.Game) (models.Game, error) {
	g.Id = len(games) + 1
	games = append(games, g)
	return g, nil
}

func UpdateGameById(g models.Game) (models.Game, error) {
	for i, game := range games {
		if game.Id == g.Id {
			games[i] = g
			return games[i], nil
		}
	}
	return *&models.Game{}, fmt.Errorf("No game with Id %d was found", g.Id)
}

func DeleteGame(id int) error {
	for i, game := range games {
		if game.Id == id {
			games = append(games[:i], games[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No game with Id %d was found", id)
}

//==========================Games==========================

//==========================Items==========================
var items []models.Item

func GetItemById(id int) (models.Item, error) {
	for _, item := range items {
		if item.Id == id{
			return item, nil
		}
	}
	return *&models.Item{}, fmt.Errorf("No item with Id %d was found", id)
}

func GetAllItems() []models.Item{
	return items
}

func CreateItem(i models.Item) (models.Item, error) {
	i.Id = len(items) + 1
	items = append(items, i)
	return i, nil
}

func UpdateItem(it models.Item) (models.Item, error) {
	for i, item := range items{
		if item.Id == it.Id {
			items[i] = it
			return items[i], nil
		}
	}
	return *&models.Item{}, fmt.Errorf("No item with Id %d was found", it.Id)
}

func DeleteItem(id int) error {
	for i, item := range items{
		if item.Id == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No item with Id %d was found", id)
}

//==========================Items==========================
