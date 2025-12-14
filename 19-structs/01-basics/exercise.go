package basics

import "fmt"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

// newGame returns a new game struct.
func newGame(id int, name string, price int, genre string) game {
	return game{
		item:  item{id: id, name: name, price: price},
		genre: genre,
	}
}

// String stringifies an item.
func (i item) String() string {
	return fmt.Sprintf("%v: %v costs %v", i.id, i.name, i.price)
}

// String stringifies a game.
func (g game) String() string {
	return fmt.Sprintf("Game %v of genre %v", g.item.String(), g.genre)
}

// newGameList creates a game store.
func newGameList() []game {
	games := []game{}

	games = append(games, newGame(1, "god of war", 50, "action adventure"))
	games = append(games, newGame(2, "x-com 2", 30, "strategy"))
	games = append(games, newGame(4, "warcraft", 40, "strategy"))

	return games
}

// queryById returns the game in the specified store with the given id or returns a "no such game" error.
func queryById(games []game, id int) (game, error) {
	for _, element := range games {
		if element.id == id {
			return element, nil
		}
	}
	return game{}, fmt.Errorf("no such game")
}

// listNameByPrice returns the name of the game(s) with price equal or smaller than a given price.
func listNameByPrice(games []game, price int) []string {
	game_list := []string{}
	for _, element := range games {
		if element.price <= price {
			game_list = append(game_list, element.name)
		}
	}
	return game_list
}
