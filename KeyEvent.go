package bl

import "xel"

var Key_Event_Type string = "key event"

type KeyEvent struct {
	Key xel.Key
	Action xel.Action
}

func (k *KeyEvent) Type() string {
	return Key_Event_Type
}

func NewKeyEvent(key xel.Key, action xel.Action) *KeyEvent {
	keyEvent := &KeyEvent{key, action}

	return keyEvent
}
