package board

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/mcp2515"
)

func init() {
	CAN_INT.Configure(machine.PinConfig{Mode: machine.PinInput})
	CAN_RST.Configure(machine.PinConfig{Mode: machine.PinOutput})
	CAN_RST.Low()
	time.Sleep(10 * time.Millisecond)
	CAN_RST.High()
	time.Sleep(10 * time.Millisecond)
	if err := machine.SPI0.Configure(
		machine.SPIConfig{
			Frequency: 500000,
			SCK:       CAN_SCK,
			SDO:       CAN_TX,
			SDI:       CAN_RX,
			Mode:      0,
		},
	); err != nil {
		panic(err)
	}
}

func NewCan() (*mcp2515.Device, error) {
	can := mcp2515.New(machine.SPI0, CAN_CS)
	can.Configure()
	if err := can.Begin(mcp2515.CAN500kBps, mcp2515.Clock8MHz); err != nil {
		return nil, err
	}
	return can, nil
}
