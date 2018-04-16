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

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	words, err := getInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	mf := getMyFlags()

	for {
		time.Sleep(*mf.delay)

		randCount, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(2))
		var sum big.Int
		var cmdCount int
		if *mf.combo {
			cmdCount = int(sum.Add(randCount, big.NewInt(1)).Int64())
		} else {
			cmdCount = 1
		}

		if *mf.jabFirst {
			doCommand(*mf.voice, *mf.speed, "jab")
		}

		i := 0
		for i < cmdCount {
			doCommand(*mf.voice, *mf.speed, getPseudoRandomWord(words))
			i++
		}

		if *mf.wildcard && cmdCount%2 == 0 {
			doCommand(*mf.voice, *mf.speed, "whatever")
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

func getMyFlags() myFlags {
	mf := myFlags{
		delay:    flag.Duration("delay", 2*time.Second, "Delay seconds."),
		voice:    flag.String("voice", "Alex", "Voice"),
		speed:    flag.String("speed", "300", "Voice speed."),
		combo:    flag.Bool("combo", false, "Sometimes do combinations"),
		jabFirst: flag.Bool("jab", false, "Always jab first. Great for boxing"),
		wildcard: flag.Bool("wildcard", false, "Sometimes finish with whatever move you want"),
	}
	flag.Parse()
	return mf
}

func doCommand(voice, voiceSpeed, word string) {
	cmd := exec.Command(voiceCmd, voiceFlag, voice, voiceSpeedFlag, voiceSpeed, word)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getPseudoRandomWord(words []string) string {
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	return words[0]
}
