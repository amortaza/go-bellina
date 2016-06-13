package bl

import "xel"

var Key_Event_Type string = "key event"

type KeyEvent struct {
	Key xel.Key
	Action xel.Action
	Alt, Ctrl, Shift bool
}

func (k *KeyEvent) Type() string {
	return Key_Event_Type
}

func NewKeyEvent(key xel.Key, action xel.Action, alt, ctrl, shift bool) *KeyEvent {
	keyEvent := &KeyEvent{key, action, alt, ctrl, shift}

	return keyEvent
}
