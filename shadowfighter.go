// Package shadowfighter provides shadowboxing functions.
package shadowfighter

import (
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

// Constants for the speech command
const (
	voiceCmd       = "say"
	voiceFlag      = "-v"
	voiceSpeedFlag = "-r"
)

// Struct to hold cli flag values
type myFlags struct {
	Delay                     *time.Duration
	Voice, Speed              *string
	Combo, JabFirst, Wildcard *bool
}

// NewMyFlags is a constructor to return myFlags instance of parsed cli flags
func NewMyFlags() *myFlags {
	mf := myFlags{
		Delay:    flag.Duration("delay", 2*time.Second, "Delay seconds."),
		Voice:    flag.String("voice", "Alex", "Voice"),
		Speed:    flag.String("speed", "300", "Voice speed."),
		Combo:    flag.Bool("combo", false, "Sometimes do combinations"),
		JabFirst: flag.Bool("jab", false, "Always jab first. Great for boxing"),
		Wildcard: flag.Bool("wildcard", false, "Sometimes finish with whatever move you want"),
	}
	flag.Parse()
	return &mf
}

// Ready method waits for you for at least 5 seconds
func (mf *myFlags) Ready() {
	initialWait := 5*time.Second - *mf.Delay
	mf.SayMove("Get ready")
	if initialWait >= 0*time.Second {
		time.Sleep(initialWait)
	}
}

// GetRandomMoveCount returns 0 or 1 randomly
func (mf *myFlags) GetRandomMoveCount() int {
	var sum big.Int
	randCount, err := cryptorand.Int(cryptorand.Reader, big.NewInt(2))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't randomize moves", err)
	}
	return int(sum.Add(randCount, big.NewInt(1)).Int64())
}

// SayMove executes your Mac OS's voice command
func (mf *myFlags) SayMove(move string) {
	cmd := exec.Command(voiceCmd, voiceFlag, *mf.Voice, voiceSpeedFlag, *mf.Speed, move)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// GetPseudoRandomMove returns a single pseudo-random move from a slice
func GetPseudoRandomMove(moves []string) string {
	rand.Shuffle(len(moves), func(i, j int) {
		moves[i], moves[j] = moves[j], moves[i]
	})
	return moves[0]
}
