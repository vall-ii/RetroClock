package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███ ",
		"█ █ ",
		"█ █ ",
		"█ █ ",
		"███ ",
	}

	one := placeholder{
		"██  ",
		" █  ",
		" █  ",
		" █  ",
		"███ ",
	}

	two := placeholder{
		"███ ",
		"  █ ",
		"███ ",
		"█   ",
		"███ ",
	}

	three := placeholder{
		"███ ",
		"  █ ",
		"███ ",
		"  █ ",
		"███ ",
	}

	four := placeholder{
		"█ █ ",
		"█ █ ",
		"███ ",
		"  █ ",
		"  █ ",
	}

	five := placeholder{
		"███ ",
		"█   ",
		"███ ",
		"  █ ",
		"███ ",
	}

	six := placeholder{
		"███ ",
		"█   ",
		"███ ",
		"█ █ ",
		"███ ",
	}

	seven := placeholder{
		"███ ",
		"  █ ",
		"  █ ",
		"  █ ",
		"  █ ",
	}

	eight := placeholder{
		"███ ",
		"█ █ ",
		"███ ",
		"█ █ ",
		"███ ",
	}

	nine := placeholder{
		"███ ",
		"█ █ ",
		"███ ",
		"  █ ",
		"███ ",
	}

	digits := []placeholder{
		zero,
		one,
		two,
		three,
		four,
		five,
		six,
		seven,
		eight,
		nine,
	}

	separator := placeholder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	separatorEmpty := placeholder{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}
	for {
		timeNow := time.Now()

		var timeStr = [6]string{}

		//получаем время поэлементно, по цифре
		if timeNow.Hour() >= 10 {
			timeStr[0] = string(strconv.Itoa(timeNow.Hour())[0])
			timeStr[1] = string(strconv.Itoa(timeNow.Hour())[1])
		} else {
			timeStr[0] = "0"
			timeStr[1] = string(strconv.Itoa(timeNow.Hour())[0])
		}
		if timeNow.Minute() >= 10 {
			timeStr[2] = string(strconv.Itoa(timeNow.Minute())[0])
			timeStr[3] = string(strconv.Itoa(timeNow.Minute())[1])
		} else {
			timeStr[2] = "0"
			timeStr[3] = string(strconv.Itoa(timeNow.Minute())[0])
		}
		if timeNow.Second() >= 10 {
			timeStr[4] = string(strconv.Itoa(timeNow.Second())[0])
			timeStr[5] = string(strconv.Itoa(timeNow.Second())[1])
		} else {
			timeStr[4] = "0"
			timeStr[5] = string(strconv.Itoa(timeNow.Second())[0])
		}

		timeInt := [6]int{}
		for i, j := range timeStr {
			d, err := strconv.Atoi(j)
			if err != nil {
				return
			} else {
				timeInt[i] = d
			}
		}
		screen.Clear()
		for i := 0; i < 5; i++ {
			currentSep := separator
			if timeNow.Second()%2 == 0 {
				currentSep = separatorEmpty
			}
			for y, t := range timeInt {
				if y == 2 || y == 4 {
					fmt.Print(currentSep[i])
				}
				for x, n := range digits {

					if x == t {

						fmt.Print(n[i])
					}
				}
			}
			fmt.Println()
		}
		time.Sleep(1 * time.Second)

		screen.MoveTopLeft()
	}
}
