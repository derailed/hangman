package hangman

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/derailed/hangman/internal/game"
	"github.com/derailed/hangman/internal/svc"
)

func gameURL(url, path string) string {
	return fmt.Sprintf("%s/game/v1/%s", url, path)
}

func newGame(uri, word string) (game.Game, error) {
	var res game.Response

	req := game.NewGameRequest{Word: word}
	payload, err := json.Marshal(req)
	if err != nil {
		return game.Game{}, err
	}
	err = svc.Call("POST", gameURL(uri, "new_game"), bytes.NewReader([]byte(payload)), &res, nil)
	return res.Game, err
}

func guess(uri string, g game.Game, letter string) (game.Game, error) {
	var res game.Response

	req := game.GuessRequest{Game: g, Guess: letter}
	payload, err := json.Marshal(req)
	if err != nil {
		return res.Game, err
	}
	err = svc.Call("POST", gameURL(uri, "guess"), bytes.NewReader([]byte(payload)), &res, nil)
	return res.Game, err
}
