package bl

import "xel"

var EventType_Key string = "key event"

type KeyEvent struct {
	Key xel.Key
	Action xel.Action
	Alt, Ctrl, Shift bool
}

func (k *KeyEvent) Type() string {
	return EventType_Key
}

func NewKeyEvent(key xel.Key, action xel.Action, alt, ctrl, shift bool) *KeyEvent {
	keyEvent := &KeyEvent{key, action, alt, ctrl, shift}

	return keyEvent
}
