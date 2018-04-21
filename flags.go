package shadowfighter

import (
	"flag"
	"time"
)

// Struct to hold cli flag values
type Flags struct {
	Delay                     *time.Duration
	Voice, Speed              *string
	Combo, JabFirst, Wildcard *bool
}

// NewFlags is a constructor to return Flags instance of parsed cli flags
func NewFlags() *Flags {
	f := Flags{
		Delay:    flag.Duration("delay", 2*time.Second, "Delay seconds."),
		Voice:    flag.String("voice", "Alex", "Voice"),
		Speed:    flag.String("speed", "300", "Voice speed."),
		Combo:    flag.Bool("combo", false, "Sometimes do combinations"),
		JabFirst: flag.Bool("jab", false, "Always jab first. Great for boxing"),
		Wildcard: flag.Bool("wildcard", false, "Sometimes finish with whatever move you want"),
	}
	flag.Parse()
	return &f
}
