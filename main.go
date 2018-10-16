package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const separator = "=============================="

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	maxA := 10 // Max random multiplicand
	maxB := 10 // Max random multiplier

	reader := bufio.NewReader(os.Stdin)

	correct := 0
	questions := 0

	fmt.Printf("\n%s\n", separator)
	start := time.Now()

	for i := 0; i < 25; i++ {
		questions = questions + 1
		a := r.Intn(maxA)
		b := r.Intn(maxB)
		p := a * b

		fmt.Printf("Q%d. %2d x %-2d = ", questions, a, b)
		if text, err := reader.ReadString('\n'); err == nil {
			n, _ := strconv.ParseInt(strings.TrimSpace(text), 10, 32)
			c := '!'

			if n == int64(p) {
				correct = correct + 1
				c = '='
			}

			fmt.Printf("    %d/%d: %2d x %-2d %c= %d\n", correct, questions, a, b, c, p)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("%s\n\n", separator)
	fmt.Printf("Your final score: %d/%d (%d%%) in %.1fs (%.1fs/question)\n", correct, questions, correct * 100 / questions, elapsed.Seconds(), elapsed.Seconds() / float64(questions))
}
