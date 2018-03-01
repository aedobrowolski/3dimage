package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	charsetP = flag.String("charset", "AbCdEfGhIjKlMnOpQR.,=+!@#^", "set of characters to paint shape with")
	randomP  = flag.Bool("random", false, "create random pattern for each line")
	periodP  = flag.Int("width", 20, "pattern width to use")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	// Get the shape
	rand.Seed(time.Now().Unix())
	shape, err := getShape()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		flag.Usage()
	}
	// Get the charset, randomize, truncate
	pattern := makeRandomText(*charsetP, *periodP)
	// Paint the shape with the pattern
	lines := strings.Split(shape, "\n")
	for _, line := range lines {
		if *randomP {
			pattern = makeRandomText(*charsetP, *periodP)
		}
		paintedLine := paint(line, pattern)
		// fmt.Println(line)
		fmt.Println(paintedLine)
	}
}

func usage() {
	base := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage: %s [-charset chars] [-random] [-width int] shapepath\n", base)
	fmt.Fprintf(os.Stderr, "where shapepath is a text file with a rectangular shape painted with 0's, 1's, 2's, etc\n")
	flag.PrintDefaults()
}

func getShape() (string, error) {
	if flag.NArg() != 1 || flag.Arg(0) == "" {
		return "", fmt.Errorf("shapepath is required")
	}
	b, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func paint(line, pattern string) string {
	if len(line) == 0 {
		return ""
	}
	period := len(pattern)
	buf := []rune(pattern)
	for i, r := range line {
		// calculate the character to append by looking back in the buffer
		// the amount to look back depends on the plane being painted
		// 0 - very far, 1 - far, 2 - close, 3 - very close
		var p int
		switch r {
		case '0', ' ':
			p = period
		case '1':
			p = period - 1
		case '2':
			p = period - 2
		case '3':
			p = period - 3
		}
		if p > 0 {
			if len(buf) >= p {
				r = buf[len(buf)-p]
			} else {
				r = rune(pattern[i])
			}
		}
		buf = append(buf, r)
	}
	return string(buf)
}

// makeRandomText shuffles INPUT returning the first GOAL characters
func makeRandomText(input string, goal int) string {
	if len(input) == 0 {
		panic("Invalid argument")
	}

	// Extend the input if needed
	for len(input) < goal {
		input = input + input
	}

	// Shuffle the bytes using the Fisher-Yates algorithm
	a := []byte(input)
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}

	return string(a[:goal])
}
