package bl

<<<<<<< HEAD
import "github.com/amortaza/go-xel"
=======
import "xel"
>>>>>>> 457139a89938fb2b2ae6a239356e9473e6b022ec

var EventType_Mouse_Button string = "mouse button event"

type MouseButtonEvent struct {
	Button xel.MouseButton
<<<<<<< HEAD
	Action xel.ButtonAction
=======
	Action xel.Action
>>>>>>> 457139a89938fb2b2ae6a239356e9473e6b022ec
	Target *Node // the node the event originated on
	CurrentTarget *Node
	BubbleToParent bool
}

func (m *MouseButtonEvent) Type() string {
	return EventType_Mouse_Button
}

<<<<<<< HEAD
func NewMouseButtonEvent(button xel.MouseButton, action xel.ButtonAction, target *Node) *MouseButtonEvent {
=======
func NewMouseButtonEvent(button xel.MouseButton, action xel.Action, target *Node) *MouseButtonEvent {
>>>>>>> 457139a89938fb2b2ae6a239356e9473e6b022ec
	e := &MouseButtonEvent{button, action, target, target, true}

	return e
}
