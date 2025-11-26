package main

// MARK: import

import (
	"context"
	"machine"
	"machine/usb"
	"time"

	"github.com/SWITCHSCIENCE/ffb_steering_controller/control"
	"github.com/SWITCHSCIENCE/ffb_steering_controller/settings"

	"github.com/SWITCHSCIENCE/picossci-ffb-wheel/board"
)

const (
	FLASH_TARGET_OFFSET = 0 // 書き込み開始アドレス(例)
)

// Flashに1ページ書き込み
func writeFlashBlock(data []byte) error {
	err := machine.Flash.EraseBlocks(0, 1)
	if err != nil {
		return err
	}
	if _, err := machine.Flash.WriteAt(data, FLASH_TARGET_OFFSET); err != nil {
		return err
	}
	return nil
}

// Flashから読み出し
func readFlashBlock() ([]byte, error) {
	buff := make([]byte, machine.Flash.WriteBlockSize())
	n, err := machine.Flash.ReadAt(buff, FLASH_TARGET_OFFSET)
	if err != nil {
		return nil, err
	}
	return buff[:n], nil
}

// MARK: variables
var (
	sw [3]bool
)

// MARK: functions

func init() {
	//usb.VendorID = 0x2341
	//usb.ProductID = 0x8036
	usb.Product = "DIY Steering Controller"
	usb.Manufacturer = "Switch Science"
	board.LCD.Show(board.Logo)
	board.LCD.Display()
	if false {
		for !machine.Serial.DTR() {
			time.Sleep(100 * time.Millisecond)
		}
		//println("boot")
		//println(machine.FlashDataStart())
	}
}

func update(menu *Menu) {
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
	if active[0] {
		menu.Up()
	} else if active[1] {
		menu.Enter()
	} else if active[2] {
		menu.Down()
	}
	if menu.IsSetting() {
		board.LED1.Low()
	} else {
		board.LED1.Low()
	}
}

func main() {
	board.LED2.Low()
	can, err := board.NewCan()
	if err != nil {
		println(err)
		return
	}
	js := control.NewWheel(can)
	b, err := readFlashBlock()
	if err != nil {
		println(err)
		return
	}
	s, err := settings.Unmarshal(b)
	if err != nil {
		println(err)
		s = settings.Default()
	}
	if err := settings.Update(s); err != nil {
		println(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		menu := NewMenu()
		tick := time.NewTicker(20 * time.Millisecond)
		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				update(menu)
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
