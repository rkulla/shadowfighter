// Package shadowfighter provides shadowboxing functions.
package shadowfighter

import (
	cryptorand "crypto/rand"
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

// Ready method waits for you for at least 5 seconds
func Ready(f *Flags) {
	initialWait := 5*time.Second - *f.Delay
	SayMove("Get ready", f)
	if initialWait >= 0*time.Second {
		time.Sleep(initialWait)
	}
}

// SayMove executes your Mac OS's voice command
func SayMove(move string, f *Flags) {
	cmd := exec.Command(voiceCmd, voiceFlag, *f.Voice, voiceSpeedFlag, *f.Speed, move)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// GetRandomMoveCount returns int 1 or 2 randomly
func GetRandomMoveCount() int {
	var sum big.Int
	randCount, err := cryptorand.Int(cryptorand.Reader, big.NewInt(2))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't randomize moves", err)
	}
	return int(sum.Add(randCount, big.NewInt(1)).Int64())
}

// GetPseudoRandomMove returns a single pseudo-random move from a slice
func GetPseudoRandomMove(moves []string) string {
	rand.Shuffle(len(moves), func(i, j int) {
		moves[i], moves[j] = moves[j], moves[i]
	})
	return moves[0]
}
