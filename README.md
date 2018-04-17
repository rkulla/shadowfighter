# ShadowFighter

Audio Shadowboxing drills for different martial arts.

This program will use MacOS's speech command to call out random lines from a text file
that you can use to drill different martial arts strikes/blocks.

It's designed to be customizable: Simply tweak the provided text files or add your own.

## Requirements

Currently requires a Mac and the free [say](https://support.apple.com/kb/PH25378?viewlocale=en_US&locale=en_US
) command-line utility.

## Usage

    $ ./shadowfighter [flags] < [drills/file.txt]

If you want to change the source code, edit it, then rebuild the shadowfighter binary with:
 
    $ go build -o shadowfighter cmd/shadowfighter/main.go

Flags:

    -combo
	    Sometimes do combinations
    -delay duration
        Delay seconds. (default 2s)
    -jab
        Always jab first. Great for boxing
	-speed
        Voice speed. (default "300")
	-voice string
	    Voice (default "Alex")
    -wildcard
        Sometimes finish with whatever move you want

If you don't specify a filename, it will read from STDIN so you can dynamically input commands instead. For example, if you wanted to just create a random workout:

    $ ./shadowfighter
    jump
    squat
    lunge
    ^D
That last line `^D` means type `Ctrl+D` on your keyboard to specify you're done inputting commands.

There are some default files to get you started. Feel free to edit them how you like on your machine:

* Boxing:

    `./shadowfighter -jab -combo < drills/boxing.txt`

* Muay Thai: 

    `./shadowfighter -delay=3s -combo -wildcard < drills/thai-boxing.txt`

  For wildcard it will say "anything" last. So do whatever you want (e.g. Spinning Back Fist).

* FMA (Filipino Martial Arts, such as Kali/Arnis/Eskrima):

    `./shadowfighter -combo < drills/fma.txt`

  Calls out angles of attack 1-7. Add more if you want.

  Tip: Alternatively think of blocking against those angles coming at you.

* Tai Chi:

    `./shadowfighter -delay=7s --speed=200 -voice=Victoria < drills/taichi.txt`

