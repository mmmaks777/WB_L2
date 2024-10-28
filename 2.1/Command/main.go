// Плюсы:
// 1. Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
// 2. Позволяет реализовать простую отмену и повтор операций.
// 3. Позволяет реализовать отложенный запуск операций.
// Минусы:
// 1. Усложняет код программы из-за введения множества дополнительных объектов.
// Применимость:
// 1. Когда нужно отправлять запросы в виде объектов, чтобы передавать их между процессами, откладывать или сохранять для последующего выполнения.
// 2. Когда требуется поддержка операций отмены и повтора.
// 3. Когда нужно сохранять историю действий.
// Примеры:
// 1. Операции (например, копировать, вставить, удалить) можно представить как команды, которые можно отменить или повторить.
// 2. Графические интерфейсы, кнопки и меню могут представлять команды, которые выполняются при нажатии.

package main

import "fmt"

type Command interface {
	Execute()
	Undo()
}

type TurnOnLightCommand struct {
	light *Light
}

func (c *TurnOnLightCommand) Execute() {
	c.light.On()
}

func (c *TurnOnLightCommand) Undo() {
	c.light.Off()
}

type Light struct {
	IsOn bool
}

func (l *Light) On() {
	l.IsOn = true
	fmt.Println("Light on")
}

func (l *Light) Off() {
	l.IsOn = false
	fmt.Println("Light off")
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func (r *RemoteControl) PressUndo() {
	r.command.Undo()
}

func main() {
	light := &Light{}

	command := &TurnOnLightCommand{light: light}
	remote := RemoteControl{}

	remote.SetCommand(command)
	remote.PressButton()
	remote.PressUndo()
}
