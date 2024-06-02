/******************************************************************************
 * Date   : June 2024
 * Brief  : Process and transmit events received via serial
 *
 * Copyright (c) 2024 by Nicolaus Starke | All rights reserved.
 * SPDX-License-Identifier: MIT
 *****************************************************************************/

package serialcomms

import (
	"fmt"
	"regexp"
)

/* ------------------------------- Types ----------------------------------- */

type EventType string

type Event struct {
	Type EventType
	Data interface{}
}


type EncoderDirectionEvent struct {
	Index     int
	Direction int
}

type EncoderSwitchEvent struct {
	Index int
	State int
}

type SideSwitchEvent struct {
	Side  int
	State int
}

type CurrentConfigEvent struct {
	Config string
}

type LEDEvent struct {
	LEDID  int
	State  int
	Color  string
}

type MIDIEvent struct {
	Data string
}

/* ------------------------------- Constants ------------------------------- */

const (
	EncoderDirectionType EventType = "ed"
	EncoderSwitchType    EventType = "es"
	SideSwitchType       EventType = "ss"
	CurrentConfigType    EventType = "cf"
	LEDEventType         EventType = "ld"
	MIDIInType           EventType = "mi"
	MIDIOutType          EventType = "mo"
)

/* ------------------------------- Variables ------------------------------- */

var EventChannel chan Event

/* ------------------------------- Functions ------------------------------- */

func init() {
	EventChannel = make(chan Event)
}

func HandleEvents() {
	for event := range EventChannel {
		switch event.Type {
		case EncoderDirectionType:
			if data, ok := event.Data.(EncoderDirectionEvent); ok {
				fmt.Printf("Received encoder direction event: %+v\n", data)
			} else {
				fmt.Println("Invalid data for EncoderDirectionType")
			}
		case EncoderSwitchType:
			if data, ok := event.Data.(EncoderSwitchEvent); ok {
				fmt.Printf("Received encoder switch event: %+v\n", data)
			} else {
				fmt.Println("Invalid data for EncoderSwitchType")
			}
		case SideSwitchType:
			if data, ok := event.Data.(SideSwitchEvent); ok {
				fmt.Printf("Received side switch event: %+v\n", data)
			} else {
				fmt.Println("Invalid data for SideSwitchType")
			}
		case CurrentConfigType:
			if data, ok := event.Data.(CurrentConfigEvent); ok {
				fmt.Printf("Received current config event: %+v\n", data)
			} else {
				fmt.Println("Invalid data for CurrentConfigType")
			}
		case LEDEventType:
			if data, ok := event.Data.(LEDEvent); ok {
				fmt.Printf("Received LED event: %+v\n", data)
			} else {
				fmt.Println("Invalid data for LEDEventType")
			}
		case MIDIInType:
			if data, ok := event.Data.(MIDIEvent); ok {
				fmt.Printf("Received MIDI in event: %+v\n", data)
			} else {
				fmt.Println("Invalid data for MIDIInType")
			}
		case MIDIOutType:
			if data, ok := event.Data.(MIDIEvent); ok {
				fmt.Printf("Received MIDI out event: %+v\n", data)
			} else {
				fmt.Println("Invalid data for MIDIOutType")
			}
		default:
			fmt.Println("Unknown event type")
		}
	}
}

func parseInput(input string) (Event, error) {
	eventType := input[:2]
	data := input[2:]

	switch eventType {
	case string(EncoderDirectionType):
		return parseEncoderDirection(data)
	case string(EncoderSwitchType):
		return parseEncoderSwitch(data)
	case string(SideSwitchType):
		return parseSideSwitch(data)
	case string(CurrentConfigType):
		return parseCurrentConfig(data)
	case string(LEDEventType):
		return parseLEDEvent(data)
	case string(MIDIInType), string(MIDIOutType):
		return parseMIDIEvent(eventType, data)
	default:
		return Event{}, fmt.Errorf("unknown event type: %s", eventType)
	}
}

func parseEncoderDirection(data string) (Event, error) {
	re := regexp.MustCompile(`\[(\d+)\]\[(-?\d+)\]`)
	matches := re.FindStringSubmatch(data)

	if len(matches) != 3 {
		return Event{}, fmt.Errorf("invalid encoder direction format: %s", data)
	}

	var event EncoderDirectionEvent
	_, err := fmt.Sscanf(matches[1], "%d", &event.Index)
	if err != nil {
		return Event{}, err
	}
	_, err = fmt.Sscanf(matches[2], "%d", &event.Direction)
	if err != nil {
		return Event{}, err
	}

	return Event{Type: EncoderDirectionType, Data: event}, nil
}

func parseEncoderSwitch(data string) (Event, error) {
	re := regexp.MustCompile(`\[(\d+)\]\[(\d+)\]`)
	matches := re.FindStringSubmatch(data)

	if len(matches) != 3 {
		return Event{}, fmt.Errorf("invalid encoder switch format: %s", data)
	}

	var event EncoderSwitchEvent
	_, err := fmt.Sscanf(matches[1], "%d", &event.Index)
	if err != nil {
		return Event{}, err
	}
	_, err = fmt.Sscanf(matches[2], "%d", &event.State)
	if err != nil {
		return Event{}, err
	}

	return Event{Type: EncoderSwitchType, Data: event}, nil
}

func parseSideSwitch(data string) (Event, error) {
	re := regexp.MustCompile(`\[(\d+)\]\[(\d+)\]`)
	matches := re.FindStringSubmatch(data)

	if len(matches) != 3 {
		return Event{}, fmt.Errorf("invalid side switch format: %s", data)
	}

	var event SideSwitchEvent
	_, err := fmt.Sscanf(matches[1], "%d", &event.Side)
	if err != nil {
		return Event{}, err
	}
	_, err = fmt.Sscanf(matches[2], "%d", &event.State)
	if err != nil {
		return Event{}, err
	}

	return Event{Type: SideSwitchType, Data: event}, nil
}

func parseCurrentConfig(data string) (Event, error) {
	return Event{Type: CurrentConfigType, Data: CurrentConfigEvent{Config: data}}, nil
}

func parseLEDEvent(data string) (Event, error) {
	re := regexp.MustCompile(`\[(\d+)\]\[(\d+)\]\[(\w+)\]`)
	matches := re.FindStringSubmatch(data)

	if len(matches) != 4 {
		return Event{}, fmt.Errorf("invalid LED event format: %s", data)
	}

	var event LEDEvent
	_, err := fmt.Sscanf(matches[1], "%d", &event.LEDID)
	if err != nil {
		return Event{}, err
	}
	_, err = fmt.Sscanf(matches[2], "%d", &event.State)
	if err != nil {
		return Event{}, err
	}
	event.Color = matches[3]

	return Event{Type: LEDEventType, Data: event}, nil
}

func parseMIDIEvent(eventType, data string) (Event, error) {
	return Event{Type: EventType(eventType), Data: MIDIEvent{Data: data}}, nil
}
