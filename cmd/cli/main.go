package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/derailed/hangman2/internal/cli"
	"github.com/derailed/hangman2/internal/game"
)

type result struct {
	Game  *game.Game  `json:"game"`
	Tally *game.Tally `json:"tally"`
}

func main() {
	svcURL := flag.String("url", "localhost:9096", "Hangman service url")

	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println()
		os.Exit(0)
	}()

	id, tally, err := cli.NewGame(*svcURL)
	if err != nil {
		panic(err)
	}

	for {
		cli.Display(tally)
		tally, err = cli.Guess(*svcURL, id, prompt())
		if err != nil {
			panic(err)
		}
		allDone(*svcURL, id, tally)
	}
}

func allDone(url string, id int, t game.Tally) {
	if t.Status == game.Won {
		fmt.Println("Noace! You've won!!")
		cli.EndGame(url, id)
		os.Exit(0)
	}
	if t.Status == game.Lost {
		fmt.Printf("Rats! You've just lost...\n")
		cli.EndGame(url, id)
		os.Exit(0)
	}
}

func prompt() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%10s? ", "Your guess")
	char, _, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}
	return string(char)
}
