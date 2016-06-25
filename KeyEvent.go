package bl

import "github.com/amortaza/go-xel"

var EventType_Key string = "key event"

type KeyEvent struct {
	Key xel.KeyboardKey
	Action xel.ButtonAction
	Alt, Ctrl, Shift bool
}

func (k *KeyEvent) Type() string {
	return EventType_Key
}

func NewKeyEvent(key xel.KeyboardKey, action xel.ButtonAction, alt, ctrl, shift bool) *KeyEvent {
	keyEvent := &KeyEvent{key, action, alt, ctrl, shift}

	return keyEvent
}
