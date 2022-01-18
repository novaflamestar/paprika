package main

import (
	"fmt"
	"os"
	"time"

	"github.com/novaflamestar/paprika/shapes"
)

var ()

type Animatable interface {
	Animate()
	Stop()
}

func main() {
	//go animate(ball)
	sType := os.Args[1]
	var a Animatable

	switch sType {
	case "Line":
		a = shapes.NewLine("Initializing", "V")
	case "Ball":
		a = shapes.NewBall("Initializing")
	case "Spinner":
		a = shapes.NewSpinner("Initializing")
	default:
		return
	}

	go a.Animate()
	time.Sleep(10 * time.Second)
	a.Stop()

	fmt.Println("Stopping...")
	os.Exit(0)
}
