/******************************************************************************
 * Date   : June 2024
 * Brief  : NEON_SAMURAI Command-Line Interface
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package main

import "github.com/nic-starke/neon_samurai_editor/pkg/serialcomms"

/* ------------------------------- Types ----------------------------------- */
/* ------------------------------- Constants ------------------------------- */
/* ------------------------------- Variables ------------------------------- */
/* ------------------------------- Functions ------------------------------- */

// func main() {
// 	fmt.Println("Welcome to NEON_SAMURAI")
// 	conn, err := comms.NewMidiConnection("SAMURAI")

// 	if (err != nil) {
// 		return
// 	}

// 	defer conn.Close()
// 	defer conn.StopDriver()

// 	start := time.Now()
// 	duration := 10 * time.Second

// 	for time.Since(start) < duration {
// 		// Your code to be executed within the loop
// 		fmt.Println("Looping...")
// 		time.Sleep(1 * time.Second) // Sleep for 1 second to avoid busy-waiting
// 	}
// }

func main() {
	serialcomms.InitSerial(115200,"/dev/ttyACM0")
	defer serialcomms.CloseSerial()
	// block forever:
	select {}
}
