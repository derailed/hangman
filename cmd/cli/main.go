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
	"github.com/derailed/hangman2/internal/hangman"
)

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

	ga, tally, err := cli.NewGame(*svcURL)
	if err != nil {
		panic(err)
	}

	for {
		cli.Display(tally)
		ga, tally, err = cli.Guess(*svcURL, ga, prompt())
		if err != nil {
			panic(err)
		}
		allDone(*svcURL, ga, tally)
	}
}

func allDone(url string, ga game.Game, t hangman.Tally) {
	if t.Status == game.Won {
		fmt.Println("Noace! You've won!!")
		os.Exit(0)
	}
	if t.Status == game.Lost {
		fmt.Printf("Rats! You've just lost... It was %s\n", ga.Letters)
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
