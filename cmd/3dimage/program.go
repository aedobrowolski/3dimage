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
	randomP  = flag.Bool("random", false, "create a new random pattern for each line")
	periodP  = flag.Int("width", 20, "pattern width to use")
	invertP  = flag.Bool("invert", false, "invert the depth for cross-eyed viewing")
	nopadP   = flag.Bool("nopad", false, "do not pad all lines to match max line")
)

func usage() {
	base := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "Usage: %s [-charset chars] [-random] [-invert] [-width int] shapepath\n", base)
	fmt.Fprintf(os.Stderr, "shapepath is a text file with a rectangular array of digits\n")
	flag.PrintDefaults()
}

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
	pattern := shuffleText(*charsetP, *periodP)
	// Paint the shape with the pattern
	lines := strings.Split(shape, "\n")
	if !*nopadP {
		padLines(lines)
	}
	for _, line := range lines {
		if *randomP {
			pattern = shuffleText(*charsetP, *periodP)
		}
		paintedLine := paint(line, pattern)
		// fmt.Println(line)
		fmt.Println(paintedLine)
	}
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
	used := make([]bool, period)
	for _, r := range line {
		// Calculate the character to append by looking back in the buffer.
		// The period of repetition depends on the plane being painted.
		if p := getPeriod(period, r); p > 0 {
			j := unused(used, len(buf)-p)
			r = buf[j]
			used[j] = true // buf[j] now used
		}
		buf = append(buf, r)
		used = append(used, false)
	}
	return string(buf)
}

func getPeriod(period int, r rune) int {
	var i int
	switch r {
	case '0', ' ':
		i = 0
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		i = int(r - '0')
	default:
		return 0 // special case to keep existing rune
	}
	if *invertP {
		i = 9 - i
	}
	return period - i
}

// unused returns the index j if not used or the first unused index instead
func unused(used []bool, j int) int {
	if used[j] {
		for j = 0; used[j]; j++ { // find first unused
		}
	}
	return j
}

func padLines(lines []string) {
	var max int
	for _, s := range lines {
		if len(s) > max {
			max = len(s)
		}
	}
	pad := strings.Repeat(" ", max)
	for i, l := range lines {
		if len(l) < max {
			lines[i] = (l + pad)[0:max]
		}
	}
}
