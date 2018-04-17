package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rkulla/shadowfighter"
)

const (
	jab      = "jab"
	wildcard = "anything"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	moves, err := getMoves()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	mf := shadowfighter.NewMyFlags()
	mf.Ready()
	moveCount := 1

	for {
		time.Sleep(*mf.Delay)

		if *mf.JabFirst {
			mf.SayMove(jab)
		}

		if *mf.Combo {
			moveCount = mf.GetRandomMoveCount()
		}

		i := 0
		for i < moveCount {
			mf.SayMove(shadowfighter.GetPseudoRandomMove(moves))
			i++
		}

		if *mf.Wildcard && moveCount%2 == 0 {
			mf.SayMove(wildcard)
		}
	}
}

// getMoves gets moves from lines of STDIN input
func getMoves() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
