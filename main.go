package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers = []rune("1234567890")
	special = []rune("+#*'-,.;:_<>^°!¹²³¼½¬{[]}\"§$%&/()=")
	umlauts = []rune("öäüÖÄÜ")

	noLetters bool
	noSpecial bool
	noUmlauts bool
	noNumbers bool

	length int64
)

func parseFlags() {
	flag.Int64Var(&length, "l", 32, "Password length")
	flag.BoolVar(&noSpecial, "no-special", false, "No special characters")
	flag.BoolVar(&noUmlauts, "no-umlauts", false, "No german umlauts")
	flag.BoolVar(&noNumbers, "no-numbers", false, "No numbers")
	flag.BoolVar(&noLetters, "no-letters", false, "No letters")
	flag.Parse()
}

func main() {
	seed, err := os.Open("/dev/urandom")
	if err != nil {
		cancel("can't open /dev/urandom as seed for number generation.")
	}
	defer seed.Close()

	parseFlags()

	var chars []rune

	if !noLetters {
		chars = append(chars, letters...)
	}
	if !noNumbers {
		chars = append(chars, numbers...)
	}
	if !noSpecial {
		chars = append(chars, special...)
	}
	if !noUmlauts {
		chars = append(chars, umlauts...)
	}

	if len(chars) == 0 {
		cancel("no possible characters for password generation available.")
	}

	lenChars := big.NewInt(int64(len(chars)))
	b := make([]rune, length)
	for i := range b {
		ri, err := rand.Int(seed, lenChars)
		if err != nil {
			cancel("unable to get random number")
		}
		b[i] = chars[int(ri.Int64())]
	}

	fmt.Println(string(b))
}

func cancel(msg string) {
	fmt.Printf("error: %s\n", msg)
	os.Exit(1)
}
