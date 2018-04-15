# ShadowFighter

Customiable shadowboxing drills for different martial arts.

## Usage

First build the binary with:

    $ go build

Now you can run `./shadowfighter [flags] < [file.txt]`

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

* Boxing:

        ./shadowfighter -jab -combo < boxing.txt

* Muay Thai: 

        ./shadowfighter -delay=3s -combo -wildcard < thai-boxing.txt

  For wildcard it will say "whatever" last. So do whatever you want (e.g. Spinning Back Fist).

* FMA (Filipino Martial Arts, such as Kali/Arnis/Eskrima.)

        ./shadowfighter -combo < fma.txt

  Calls out angles of attack 1-7. Add more if you want.

  Tip: Alternatively think of blocking against those angles coming at you.

* Tai Chi:

        ./shadowfighter -delay=7s --speed=200 -voice=Victoria < taichi.txt
