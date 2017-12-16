package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/derailed/hangman2/internal/game"
)

type response struct {
	Game  game.Game  `json:"game"`
	Tally game.Tally `json:"tally"`
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println()
		os.Exit(0)
	}()

	game, tally, err := newGame()
	if err != nil {
		panic(err)
	}

	for {
		display(tally)
		letter := prompt()
		game, tally, err = guess(game, letter)
		if err != nil {
			panic(err)
		}
		allDone(game, tally)
	}
}

func url(path string) string {
	return fmt.Sprintf("http://%s/game/v1/%s", os.Getenv("GATEWAY_URL"), path)
}

func newGame() (*game.Game, *game.Tally, error) {
	req, err := http.NewRequest("GET", url("new_game"), nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	clt := http.DefaultClient
	resp, err := clt.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, nil, errors.New("game toast")
	}

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, nil, err
	}

	return &res.Game, &res.Tally, nil
}

func guess(g *game.Game, letter rune) (*game.Game, *game.Tally, error) {
	guess := struct {
		Letter string     `json:"guess"`
		Game   *game.Game `json:"game"`
	}{
		string(letter), g,
	}
	bodyJSON, err := json.Marshal(guess)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("POST", url("guess"), bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	clt := http.DefaultClient
	resp, err := clt.Do(req)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, errors.New("game toast")
	}

	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, nil, err
	}

	return &res.Game, &res.Tally, nil
}

func allDone(g *game.Game, t *game.Tally) {
	if t.Status == game.Won {
		fmt.Println("Noace! You've won!!")
		os.Exit(0)
	}
	if t.Status == game.Lost {
		fmt.Printf("Rats! You've just lost... [%s]\n", g.Letters)
		os.Exit(0)
	}
}

func display(tally *game.Tally) {
	fmt.Printf("%10s: %d\n", "Turns", tally.TurnsLeft)
	fmt.Printf("%10s: %s\n", "Status", statusToH(tally.Status))
	fmt.Printf("%10s: %s\n", "Letters", tally.Letters)
}

func statusToH(s game.Status) string {
	switch s {
	case game.Guessed:
		return "Already Guessed"
	case game.Won:
		return "Won"
	case game.Lost:
		return "Lost"
	case game.Good:
		return "Good"
	case game.Bad:
		return "Bad"
	default:
		return fmt.Sprintf("Initialized")
	}
}

func prompt() rune {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%10s: ", "Your guess")
	char, _, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}

	return char
}
