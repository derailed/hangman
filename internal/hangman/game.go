package hangman

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/derailed/hangman2/internal/game"
	"github.com/derailed/hangman2/internal/svc"
)

type result struct {
	Game  game.Game  `json:"game"`
	Tally game.Tally `json:"tally"`
}

func gameURL(url, path string) string {
	return fmt.Sprintf("%s/game/v1/%s", url, path)
}

func newGame(uri, word string) (game.Game, game.Tally, error) {
	var res game.NewGameResponse

	req := game.NewGameRequest{Word: word}
	payload, err := json.Marshal(req)
	if err != nil {
		return res.Game, res.Tally, err
	}
	err = svc.Call("POST", gameURL(uri, "new_game"), bytes.NewReader([]byte(payload)), &res)
	fmt.Println("RES", res)
	return res.Game, res.Tally, err
}

func guess(uri string, g game.Game, letter string) (game.Game, game.Tally, error) {
	var res game.GuessResponse

	req := game.GuessRequest{Game: g, Guess: letter}
	payload, err := json.Marshal(req)
	if err != nil {
		return res.Game, res.Tally, err
	}
	fmt.Println("PAYLOAD", string(payload))
	err = svc.Call("POST", gameURL(uri, "guess"), bytes.NewReader([]byte(payload)), &res)
	return res.Game, res.Tally, err
}
