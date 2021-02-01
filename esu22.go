package esu22

import (
	"github.com/google/gousb"
)

type Input byte

const (
	Microphone       Input = 0x01
	HiZ                    = 0x02
	Line                   = 0x04
	MicrophoneAndHiZ       = 0x03
)

func SetInput(i Input) error  { return send(0x2a, byte(i)) }
func EnableHeadphones() error { return send(0x1a, 0) }

func send(op, x byte) error {
	ctx := gousb.NewContext()
	defer ctx.Close()

	d, err := ctx.OpenDeviceWithVIDPID(0x0a92, 0x0141)
	if err != nil {
		return err
	}
	defer d.Close()

	_, err = d.Control(0x21, 0x09, 0x0200, 0, []byte{
		0:  0x12,
		1:  0x34,
		2:  op,
		4:  1,
		5:  x,
		21: 0x80,
	})
	return err
}
