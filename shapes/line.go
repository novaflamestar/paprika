package shapes

import (
	"fmt"
	"strings"
	"time"
)

type Line struct {
	icon     string
	aniShape Shape
}

func NewLine(str, shape string) Line {
	c := make(chan int)
	return Line{shape, Shape{str, c}}
}

func (l Line) Stop() {
	l.aniShape.stopper <- 1
}

func (l Line) Animate() {
	i := 0
	size := 10
	initSnake := strings.Repeat("-", size)
	for {
		select {
		case <-l.aniShape.stopper:
			fmt.Printf("\033[H\033[2J%s Initialized! %s\n", strings.Repeat("+", size), strings.Repeat("+", size))
			return
		default:
			before, after := initSnake[:i], ""
			if i != len(initSnake) {
				after = initSnake[i+1:]
			}
			fmt.Printf("\033[H\033[2J%s %s %s\n", strings.Join([]string{before, l.icon, after}, ""), l.aniShape.wrapped, strings.Join([]string{after, l.icon, before}, ""))
			i = (i + 1) % len(initSnake)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
