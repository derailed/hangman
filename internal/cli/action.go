package cli

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/derailed/hangman/internal/game"
	"github.com/derailed/hangman/internal/hangman"
	"github.com/derailed/hangman/internal/svc"
)

func svcURL(url, action string) string {
	return fmt.Sprintf("http://%s/hangman/v1/%s", url, action)
}

// NewGame starts a new game
func NewGame(URL string) (game.Game, hangman.Tally, error) {
	var resp hangman.Response

	err := svc.Call("POST", svcURL(URL, "new_game"), nil, &resp, nil)
	return resp.Game, resp.Tally, err
}

// Guess a letter
func Guess(URL string, ga game.Game, letter string) (game.Game, hangman.Tally, error) {
	var res hangman.Response

	req := game.GuessRequest{Game: ga, Guess: letter}
	payload, err := json.Marshal(req)
	if err != nil {
		return res.Game, res.Tally, err
	}
	err = svc.Call("POST", svcURL(URL, "guess"), bytes.NewReader([]byte(payload)), &res, nil)
	return res.Game, res.Tally, err
}
