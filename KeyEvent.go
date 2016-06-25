package bl

<<<<<<< HEAD
import "github.com/amortaza/go-xel"
=======
import "xel"
>>>>>>> 457139a89938fb2b2ae6a239356e9473e6b022ec

var EventType_Key string = "key event"

type KeyEvent struct {
<<<<<<< HEAD
	Key xel.KeyboardKey
	Action xel.ButtonAction
=======
	Key xel.Key
	Action xel.Action
>>>>>>> 457139a89938fb2b2ae6a239356e9473e6b022ec
	Alt, Ctrl, Shift bool
}

func (k *KeyEvent) Type() string {
	return EventType_Key
}

<<<<<<< HEAD
func NewKeyEvent(key xel.KeyboardKey, action xel.ButtonAction, alt, ctrl, shift bool) *KeyEvent {
=======
func NewKeyEvent(key xel.Key, action xel.Action, alt, ctrl, shift bool) *KeyEvent {
>>>>>>> 457139a89938fb2b2ae6a239356e9473e6b022ec
	keyEvent := &KeyEvent{key, action, alt, ctrl, shift}

	return keyEvent
}
