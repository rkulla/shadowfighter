package main

import (
	"bufio"
	cryptorand "crypto/rand"
	"flag"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	voiceCmd       = "say"
	voiceFlag      = "-v"
	voiceSpeedFlag = "-r"

	jab      = "jab"
	wildcard = "anything"
)

type myFlags struct {
	delay                     *time.Duration
	voice, speed              *string
	combo, jabFirst, wildcard *bool
}

func newMyFlags() *myFlags {
	mf := myFlags{
		delay:    flag.Duration("delay", 2*time.Second, "Delay seconds."),
		voice:    flag.String("voice", "Alex", "Voice"),
		speed:    flag.String("speed", "300", "Voice speed."),
		combo:    flag.Bool("combo", false, "Sometimes do combinations"),
		jabFirst: flag.Bool("jab", false, "Always jab first. Great for boxing"),
		wildcard: flag.Bool("wildcard", false, "Sometimes finish with whatever move you want"),
	}
	flag.Parse()
	return &mf
}

func (mf *myFlags) Ready() {
	initialWait := 5*time.Second - *mf.delay
	mf.SayMove("Get ready")
	if initialWait >= 0*time.Second {
		time.Sleep(initialWait)
	}
}

func (mf *myFlags) GetRandomMoveCount() int {
	var sum big.Int
	randCount, err := cryptorand.Int(cryptorand.Reader, big.NewInt(2))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't randomize moves", err)
	}
	return int(sum.Add(randCount, big.NewInt(1)).Int64())
}

func (mf *myFlags) SayMove(move string) {
	cmd := exec.Command(voiceCmd, voiceFlag, *mf.voice, voiceSpeedFlag, *mf.speed, move)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	moves, err := getInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	mf := newMyFlags()
	mf.Ready()
	moveCount := 1

	for {
		time.Sleep(*mf.delay)

		if *mf.jabFirst {
			mf.SayMove(jab)
		}

		if *mf.combo {
			moveCount = mf.GetRandomMoveCount()
		}

		i := 0
		for i < moveCount {
			mf.SayMove(getPseudoRandomMove(moves))
			i++
		}

		if *mf.wildcard && moveCount%2 == 0 {
			mf.SayMove(wildcard)
		}
	}
}

func getInput() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getPseudoRandomMove(moves []string) string {
	rand.Shuffle(len(moves), func(i, j int) {
		moves[i], moves[j] = moves[j], moves[i]
	})
	return moves[0]
}
