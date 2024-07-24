package patterns

import "fmt"

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

type Command interface {
	Execute()
}

type LightObj struct{}

func (obj LightObj) On() {
	fmt.Println("LightObj On")
}

func (obj LightObj) Off() {
	fmt.Println("LightObj Off")
}

type LightOnCommand struct {
	lo *LightObj
}

func (loc *LightOnCommand) Execute() {
	loc.lo.On()
}

type LightOffCommand struct {
	lo *LightObj
}

func (loc *LightOffCommand) Execute() {
	loc.lo.Off()
}
