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
)

type myFlags struct {
	delay                     *time.Duration
	voice, speed              *string
	combo, jabFirst, wildcard *bool
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	moves, err := getInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	mf := getMyFlags()
	mf.GetReady()

	for {
		time.Sleep(*mf.delay)
		moveCount := mf.GetMoveCount()

		if *mf.jabFirst {
			mf.DoCommand("jab")
		}

		i := 0
		for i < moveCount {
			mf.DoCommand(getPseudoRandomMove(moves))
			i++
		}

		if *mf.wildcard && moveCount%2 == 0 {
			mf.DoCommand("whatever")
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

func (mf *myFlags) GetReady() {
	initialWait := 5*time.Second - *mf.delay
	mf.DoCommand("Get ready")
	if initialWait >= 0*time.Second {
		time.Sleep(initialWait)
	}
}

func getMyFlags() *myFlags {
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

func (mf *myFlags) GetMoveCount() int {
	randCount, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(2))
	var sum big.Int
	var moveCount int

	// If combos are enabled, we'll only randomly call them out
	if *mf.combo {
		moveCount = int(sum.Add(randCount, big.NewInt(1)).Int64())
	} else {
		moveCount = 1
	}

	return moveCount
}

func (mf *myFlags) DoCommand(move string) {
	cmd := exec.Command(voiceCmd, voiceFlag, *mf.voice, voiceSpeedFlag, *mf.speed, move)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getPseudoRandomMove(moves []string) string {
	rand.Shuffle(len(moves), func(i, j int) {
		moves[i], moves[j] = moves[j], moves[i]
	})
	return moves[0]
}
