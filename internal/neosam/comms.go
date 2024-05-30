/******************************************************************************
 * Date   : May 2024
 * Brief  : Custom messaging data structure.
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package neon_samurai

import (
	"time"
)

/* ------------------------------- Types ----------------------------------- */
/* ------------------------------- Constants ------------------------------- */

const (
	cmdRequest byte = 0x01
	cmdResponse byte = 0x02
)

/* ------------------------------- Variables ------------------------------- */
/* ------------------------------- Functions ------------------------------- */

func tx([]byte) error {
	// transmit the data...
	return nil
}

func rx(timeout time.Duration) ([]byte, error) {
	return nil,nil
}

func RequestObject(id objectID, local interface{}) error {
	tx([]byte{cmdRequest, byte(id)})

	response, err := rx(time.Millisecond * 500)
	if err != nil {
		return err
	}

	return Deserialise(response, local)
}

func RespondObject (id objectID, local interface{}) error {
	data, err := Serialise(local)
	if err != nil {
		return err
	}

	return tx(append([]byte{cmdResponse, byte(id)}, data...))
}
