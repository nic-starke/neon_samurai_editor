/******************************************************************************
 * Date   : May 2024
 * Brief  : Object definitions.
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package objects

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/* ------------------------------- Types ----------------------------------- */

type objectID uint

/* ------------------------------- Constants ------------------------------- */

const (
	objectSystemConfig objectID = iota

	objectEncoder
	objectEncoderQuad
	objectEncoderVmap
	objectEncoderSwitch
	objectEncoderDisplay

	objectVmap
	objectVmapProtocol

	objectSideSwitch

	objectDisplay

	objectFirmware
)

/* ------------------------------- Variables ------------------------------- */
/* ------------------------------- Functions ------------------------------- */

func Serialise(input interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.LittleEndian, input)
	if err != nil {
		return nil, fmt.Errorf("error serializing data: %v", err)
	}
	return buffer.Bytes(), nil
}

func Deserialise(data []byte, out interface{}) error {
	buffer := bytes.NewReader(data)
	err := binary.Read(buffer, binary.LittleEndian, out)
	if err != nil {
		return fmt.Errorf("error deserializing data: %v", err)
	}
	return nil
}
