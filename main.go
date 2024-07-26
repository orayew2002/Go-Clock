package main

import (
	"fmt"
	"time"
)

const (
	space = 10
	dotes = 11
)

var digits = [...][5]string{
	0:     {"████", "█  █", "█  █", "█  █", "████"},
	1:     {"  █ ", " ██ ", "  █ ", "  █ ", "████"},
	2:     {"████", "   █", "████", "█   ", "████"},
	3:     {"████", "   █", "████", "   █", "████"},
	4:     {"█  █", "█  █", "████", "   █", "   █"},
	5:     {"████", "█   ", "████", "   █", "████"},
	6:     {"████", "█   ", "████", "█  █", "████"},
	7:     {"████", "   █", "  █ ", "  █ ", "  █ "},
	8:     {"████", "█  █", "████", "█  █", "████"},
	9:     {"████", "█  █", "████", "   █", "████"},
	space: {"    ", "    ", "    ", "    ", "    "},
	dotes: {"    ", " ██ ", "    ", " ██ ", "    "},
}

func main() {
	screenClean()

	for {
		moveTopLeft()

		hour, minute, second := time.Now().Hour(), time.Now().Minute(), time.Now().Second()
		printTime(hour, minute, second)

		time.Sleep(time.Second)
	}
}

func printTime(hour, minute, second int) {
	hourDigits := [2]int{hour / 10, hour % 10}
	minuteDigits := [2]int{minute / 10, minute % 10}
	secondDigits := [2]int{second / 10, second % 10}

	for i := 0; i < 5; i++ {
		printRow(hourDigits, minuteDigits, secondDigits, i)
	}
	fmt.Println()
}

func printRow(hourDigits, minuteDigits, secondDigits [2]int, row int) {
	fmt.Print(digits[hourDigits[0]][row], digits[space][row], digits[hourDigits[1]][row])
	fmt.Print(digits[space][row], digits[dotes][row], digits[space][row])
	fmt.Print(digits[minuteDigits[0]][row], digits[space][row], digits[minuteDigits[1]][row])
	fmt.Print(digits[space][row], digits[dotes][row], digits[space][row])
	fmt.Print(digits[secondDigits[0]][row], digits[space][row], digits[secondDigits[1]][row])
	fmt.Println()
}

func screenClean() {
	fmt.Print("\033[2J")
}

func moveTopLeft() {
	fmt.Print("\033[H")
}
