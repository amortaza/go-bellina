package bl

import "github.com/amortaza/go-hal"

var EventType_Key string = "key-event"

type KeyEvent struct {

	Key hal.KeyboardKey

	Action hal.ButtonAction

	Alt, Ctrl, Shift bool
}

func (k *KeyEvent) Type() string {

	return EventType_Key
}

func NewKeyEvent(
		key hal.KeyboardKey,
		action hal.ButtonAction,
		alt, ctrl, shift bool) *KeyEvent {

	keyEvent := &KeyEvent{key, action, alt, ctrl, shift}

	return keyEvent
}
