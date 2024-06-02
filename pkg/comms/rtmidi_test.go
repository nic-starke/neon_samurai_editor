/******************************************************************************
 * Date   : June 2024
 * Brief  : Test routines for WebMidi backend.
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package comms

import (
	"testing"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv" // autoregisters driver
)

/* ------------------------------- Types ----------------------------------- */
/* ------------------------------- Constants ------------------------------- */

const deviceName string = "SAMURAI:SAMURAI MIDI"

/* ------------------------------- Variables ------------------------------- */
/* ------------------------------- Functions ------------------------------- */

func init() {

}

func TestNewConnection(t *testing.T) {
	defer midi.CloseDriver()

	conn, err := NewMidiConnection(deviceName)
	if err != nil {
		t.Error(err)
	}

	conn.Close()
}
