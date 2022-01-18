package shapes

import (
	"fmt"
	"time"
)

var (
	spinner = []string{"|", "/", "-", "\\"}
)

type Spinner struct {
	aniShape Shape
}

func NewSpinner(str string) Spinner {
	c := make(chan int)
	return Spinner{Shape{str, c}}
}

func (s Spinner) Stop() {
	s.aniShape.stopper <- 1
}

func (s Spinner) Animate() {
	i := 0
	for {
		select {
		case <-s.aniShape.stopper:
			fmt.Printf("\033[H\033[2J%s Initialized! %s\n", "*", "*")
			return
		default:
			fmt.Printf("\033[H\033[2J%s %s %s\n", spinner[i], s.aniShape.wrapped, spinner[i])
			i = (i + 1) % len(spinner)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
