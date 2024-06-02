/******************************************************************************
 * Date   : May 2024
 * Brief  : MIDI communication backend implementation.
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package comms

import (
	"fmt"
	"log"
	"strings"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

/* ------------------------------- Types ----------------------------------- */

type RTMidiConnection struct {
	CommsBase
	in  drivers.In
	out drivers.Out
	stopFn func()
}

/* ------------------------------- Constants ------------------------------- */
/* ------------------------------- Variables ------------------------------- */
/* ------------------------------- Functions ------------------------------- */

func NewMidiConnection(device string) (*RTMidiConnection, error) {
	log.Println("Attempting a new midi connection to", device)
	c := &RTMidiConnection{}

	if err := c.StartDriver(); err != nil {
		return nil, fmt.Errorf("start driver failed: %w", err);
	}

	for _, in := range midi.GetInPorts() {
		if strings.Contains(in.String(),device) {
			log.Printf("found %v input port\n", device)
			c.in = in
			break
		}
	}

	if c.in == nil {
		return nil, fmt.Errorf("could not find input port for %v", device)
	}

	for _, out := range midi.GetOutPorts() {
		if strings.Contains(out.String(), device) {
			log.Printf("found %v output port\n", device)
			c.out = out
			break
		}
	}

	if c.out == nil {
		return nil, fmt.Errorf("could not find output port for %v", device)
	}

	log.Println("Listening to input...")

	if stop, err := midi.ListenTo(c.in, onMessage); err != nil {
		return nil, err
	} else {
		c.stopFn = stop
	}

	c.state = connActive
	c.CommsProtocol = c
	return c, nil
}

func onMessage(msg midi.Message, millis int32) {
	fmt.Println("got a new midi message: ", msg)
}

func (c *RTMidiConnection) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (c *RTMidiConnection) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (c *RTMidiConnection) Close() error {
	log.Println("Closing connection")
	c.stopFn()
	c.in = nil
	c.out = nil
	return nil
}

func (c *RTMidiConnection) StartDriver() error {
	return nil
}

func (c *RTMidiConnection) StopDriver() error {
	midi.CloseDriver()
	return nil
}
