package shapes

import (
	"fmt"
	"time"
)

var (
	ball = []string{"O", "o", ".", "o", "O"}
)

type Ball struct {
	aniShape Shape
}

func NewBall(str string) Ball {
	c := make(chan int)
	return Ball{Shape{str, c}}
}

func (b Ball) Stop() {
	b.aniShape.stopper <- 1
}

func (b Ball) Animate() {
	i := 0
	for {
		select {
		case <-b.aniShape.stopper:
			fmt.Printf("\033[H\033[2J%s Initialized! %s\n", "*", "*")
			return
		default:
			fmt.Printf("\033[H\033[2J%s %s %s\n", ball[i], b.aniShape.wrapped, ball[i])
			i = (i + 1) % len(ball)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
