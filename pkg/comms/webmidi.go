/******************************************************************************
 * Date   : May 2024
 * Brief  : MIDI communication backend implementation.
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package comms

import (
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
	_ "gitlab.com/gomidi/midi/v2/drivers/webmididrv" // autoregisters driver
)

/* ------------------------------- Types ----------------------------------- */

type WebMidiConnection struct {
	CommsBase
	in  drivers.In
	out drivers.Out
}

/* ------------------------------- Constants ------------------------------- */
/* ------------------------------- Variables ------------------------------- */
/* ------------------------------- Functions ------------------------------- */

func NewConnection(device string) (*WebMidiConnection, error) {
	c := &WebMidiConnection{}

	for _, in := range midi.GetInPorts() {
		if in.String() == device {
			c.in = in
			break
		}
	}

	for _, out := range midi.GetOutPorts() {
		if out.String() == device {
			c.out = out
			break
		}
	}

	c.state = connActive
	c.CommsProtocol = c
	return c, nil
}

func (c *WebMidiConnection) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (c *WebMidiConnection) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (c *WebMidiConnection) Close() error {
	return nil
}

func (c *WebMidiConnection) StartDriver() error {
	return nil
}

func (c *WebMidiConnection) StopDriver() error {
	return nil
}
