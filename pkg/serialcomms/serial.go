/******************************************************************************
 * Date   : June 2024
 * Brief  : Serial communications parsing
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package serialcomms

import (
	"bufio"
	"log"

	"go.bug.st/serial"
)

/* ------------------------------- Types ----------------------------------- */
/* ------------------------------- Constants ------------------------------- */
/* ------------------------------- Variables ------------------------------- */

var port serial.Port

/* ------------------------------- Functions ------------------------------- */

func InitSerial(baud int, dev string) {
	mode := &serial.Mode{
		BaudRate: baud,
	}

	port, err := serial.Open(dev, mode)
	if err != nil {
		log.Fatalf("Error opening serial port: %v", err)
	}
	go readSerialData(port)
	go HandleEvents()
}

func CloseSerial() {
	if port != nil {
		port.Close()
	}
}

func readSerialData(port serial.Port) {
	scanner := bufio.NewScanner(port)
	for scanner.Scan() {
		line := scanner.Text()
		event, err := parseInput(line)
		if err != nil {
			log.Printf("Error parsing input: %v", err)
			continue
		}
		EventChannel <- event
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from serial port: %v", err)
	}
}
