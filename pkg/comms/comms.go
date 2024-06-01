/******************************************************************************
 * Date   : May 2024
 * Brief  : Custom messaging data structure.
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package comms

import (
	"fmt"
	"io"
	"time"
)

/* ------------------------------- Types ----------------------------------- */

type CommsProtocol interface {
	io.ReadWriteCloser
	StartDriver() error
	StopDriver() error
}

type CommsBase struct {
	CommsProtocol
	txTimeout time.Duration
	rxTimeout time.Duration
	state int
}

/* ------------------------------- Constants ------------------------------- */

const (
	cmdRequest byte = 0x01
	cmdResponse byte = 0x02
)

const (
	connInactive = iota
	connActive
	connError
)

var (
	errActive = fmt.Errorf("connection already active")
	errInactive = fmt.Errorf("connection not active")
	errTimeout = fmt.Errorf("timeout")
)


/* ------------------------------- Variables ------------------------------- */
/* ------------------------------- Functions ------------------------------- */

func (c *CommsBase) SetTimeouts(tx time.Duration, rx time.Duration) {
	c.rxTimeout = rx
	c.txTimeout = tx
}

func (c *CommsBase) Send(data []byte) error {
	if (c.state != connActive) {
		return errInactive
	}

	dataLen := len(data)
	errCh := make(chan error, 1)
	doneCh := make(chan error, 1)

	// This go routine is a non-blocking transmit - it tries to write the entire
	// data buffer and will return any errors.
	go func() {
		defer close(doneCh)
		sentLen := 0
		for sentLen < dataLen {
			n, err := c.Write(data[sentLen:0]);
			if err != nil {
				errCh <- err
				return
			}
			sentLen += n
		}
	}()

	// This select statement receives the error from the non-blocking transmit
	// and it also receives a timeout. If the timeout occurs first then
	// a timeout error is returned to the caller.
	select {
	case err := <- errCh:
		return err
	case <-time.After(c.txTimeout):
		return errTimeout
	case <-doneCh:
			return nil
	}
}

func (c *CommsBase) Receive(expectedLen int) ([]byte,error) {
	if c.state != connActive {
		return nil, errInactive
	}

	receivedLen := 0
	data := make([]byte, expectedLen)
	errCh := make(chan error, 1)
	doneCh := make(chan struct{})

	// Non blocking receive that returns an error to a channel
	go func() {
		defer close(doneCh)
		for receivedLen < expectedLen {
			n, err := c.Read(data[receivedLen:])
			if err != nil {
				errCh <- err
				return
			}
			receivedLen += n
		}
	}()

	select {
	case err := <-errCh:
		return nil, err
	case <-time.After(c.rxTimeout):
		return nil, errTimeout
	case <-doneCh:
		return data, nil
	}

}

func (c *CommsBase) Stop() error {
	if c.state != connActive {
		return errInactive
	}

	if err := c.Close(); err != nil {
		c.state = connError
		return err
	}

	return nil
}
