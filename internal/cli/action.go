package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/derailed/hangman2/internal/game"
	"github.com/derailed/hangman2/internal/hangman"
	"github.com/derailed/hangman2/internal/svc"
)

func svcURL(url, action string) string {
	return fmt.Sprintf("http://%s/hangman/v1/%s", url, action)
}

func NewGame(URL string) (int, game.Tally, error) {
	var resp hangman.NewGameResponse

	err := svc.Call("POST", svcURL(URL, "new_game"), nil, &resp)
	return resp.ID, resp.Tally, err
}

func EndGame(URL string, id int) error {
	var res hangman.EndGameResponse
	index := strconv.Itoa(id)

	req := hangman.EndGameRequest{ID: index}
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return svc.Call("POST", svcURL(URL, "end_game"), bytes.NewReader([]byte(payload)), &res)
}

func Guess(URL string, id int, letter string) (game.Tally, error) {
	var res game.GuessResponse

	index := strconv.Itoa(id)
	fmt.Println("ID", index)

	req := hangman.GuessRequest{ID: index, Guess: letter}
	payload, err := json.Marshal(req)
	if err != nil {
		return res.Tally, err
	}
	err = svc.Call("POST", svcURL(URL, "guess"), bytes.NewReader([]byte(payload)), &res)
	return res.Tally, err
}
