package bl

var EventType_Key string = "key-event"

type KeyEvent struct {
	Key KeyboardKey
	Action ButtonAction
	Alt, Ctrl, Shift bool
}

func (k *KeyEvent) Type() string {
	return EventType_Key
}

func NewKeyEvent(key KeyboardKey, action ButtonAction, alt, ctrl, shift bool) *KeyEvent {
	keyEvent := &KeyEvent{key, action, alt, ctrl, shift}

	return keyEvent
}
