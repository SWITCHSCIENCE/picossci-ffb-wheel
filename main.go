package main

import (
	"context"
	"time"

	"github.com/SWITCHSCIENCE/ffb_steering_controller/control"
	"github.com/SWITCHSCIENCE/ffb_steering_controller/settings"

	"github.com/SWITCHSCIENCE/picossci-ffb-wheel/board"
)

var (
	sw [3]bool
)

func update() {
	s := settings.Get()
	now := [3]bool{
		!board.SW1.Get(),
		!board.SW2.Get(),
		!board.SW3.Get(),
	}
	active := [3]bool{
		now[0] && !sw[0],
		now[1] && !sw[1],
		now[2] && !sw[2],
	}
	copy(sw[:], now[:])
	current := s.Lock2Lock
	next := current
	switch {
	case active[2]:
		switch current {
		case 1080:
		case 720:
			next = 1080
		case 540:
			next = 720
		case 360:
			next = 540
		case 180:
			next = 360
		}
	case active[0]:
		switch s.Lock2Lock {
		case 1080:
			next = 720
		case 720:
			next = 540
		case 540:
			next = 360
		case 360:
			next = 180
		case 180:
		}
	}
	switch next {
	case 1080:
		board.LED1.Low()
		board.LED2.High()
	case 720:
		board.LED1.Low()
		board.LED2.Low()
	case 540:
		board.LED1.High()
		board.LED2.Low()
	case 360:
		board.LED1.High()
		board.LED2.Low()
	case 180:
		board.LED1.High()
		board.LED2.High()
	}
	if s.Lock2Lock != next {
		s.Lock2Lock = next
		settings.Update(s)
	}
}

func main() {
	board.LED1.Low()
	can, err := board.NewCan()
	if err != nil {
		println(err)
		return
	}
	js := control.NewWheel(can)
	board.ShowLogo()
	s := settings.Get()
	s.MaxCenteringForce = 50
	settings.Update(s)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		tick := time.NewTicker(20 * time.Millisecond)
		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				update()
			}
		}
	}()
	defer cancel()
	for {
		if err := js.Loop(ctx); err != nil {
			println(err)
			time.Sleep(3 * time.Second)
		}
	}
}
