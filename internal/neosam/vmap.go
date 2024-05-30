/******************************************************************************
 * Date   : May 2024
 * Brief  : Vmap (virtual paramater mappings).
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package neon_samurai

/* ------------------------------- Types ----------------------------------- */

type vmap struct {


}

type protocol interface {

}

type protoType uint
type vmapMode uint
type vmapDisplay uint

/* ------------------------------- Constants ------------------------------- */

const (
	none protoType = iota
	midi
	osc
)

const (
	modeToggle vmapMode = iota
	modeOverlay
)

const (
	single vmapDisplay = iota
	overlay
)

/* ------------------------------- Variables ------------------------------- */
/* ------------------------------- Functions ------------------------------- */
