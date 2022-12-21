package pattern

import "fmt"

type Command interface {
	Execute()
}

type MoveLeft struct {
	receiver *Receiver
}

func (mL *MoveLeft) Execute() {
	mL.receiver.Action("Переместить влево")
}

type MoveRight struct {
	receiver *Receiver
}

func (mR *MoveRight) Execute() {
	mR.receiver.Action("Переместить вправо")
}

type Receiver struct {
}

func (r *Receiver) Action(m string) {
	fmt.Println(m)
}

// Button - Invoker (исполнитель)
type Button struct {
	command Command
}

func (b *Button) press(cmd Command) {
	cmd.Execute()
}

func InvokerFuncktion() {
	receiver := new(Receiver)
	left := &MoveLeft{receiver}
	right := &MoveRight{receiver}
	invoker := new(Button)
	invoker.press(left)
	invoker.press(right)
}
