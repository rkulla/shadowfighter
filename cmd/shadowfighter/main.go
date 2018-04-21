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

	flags := shadowfighter.NewFlags()
	shadowfighter.Ready(flags)
	moveCount := 1

	for {
		time.Sleep(*flags.Delay)

		if *flags.JabFirst {
			shadowfighter.SayMove(jab, flags)
		}

		if *flags.Combo {
			moveCount = shadowfighter.GetRandomMoveCount()
		}

		i := 0
		for i < moveCount {
			shadowfighter.SayMove(shadowfighter.GetPseudoRandomMove(moves), flags)
			i++
		}

		if *flags.Wildcard && moveCount%2 == 0 {
			shadowfighter.SayMove(wildcard, flags)
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
