package bl

import "github.com/amortaza/go-bellina/constants"

var EventType_Key string = "key event"

type KeyEvent struct {
	Key bl.KeyboardKey
	Action bl.ButtonAction
	Alt, Ctrl, Shift bool
}

func (k *KeyEvent) Type() string {
	return EventType_Key
}

func NewKeyEvent(key bl.KeyboardKey, action bl.ButtonAction, alt, ctrl, shift bool) *KeyEvent {
	keyEvent := &KeyEvent{key, action, alt, ctrl, shift}

	return keyEvent
}
