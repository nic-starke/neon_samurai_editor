package main

import (
	"fmt"

	"gitlab.com/gomidi/midi/v2"

	_ "gitlab.com/gomidi/midi/v2/drivers/webmididrv" // autoregisters driver
)

func main() {
	defer midi.CloseDriver()

	fmt.Printf("Midi In :\n" + midi.GetIn().String() + "\n")
	fmt.Printf("Midi Out :\n" + midi.GetOutPorts().String() + "\n")

}
