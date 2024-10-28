// Плюсы:
// 1. Избавляет от множества больших условных операторов машины состояний.
// 2. Легко добавить новые состояния и соответствующее поведение без изменения существующего кода.
// 3. Концентрирует в одном месте код, связанный с определённым состоянием.
// Минусы:
// 1. Может неоправданно усложнить код, если состояний мало и они редко меняются.
// Применимость:
// 1. Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём типов состояний много, и их код часто меняется.
// 2. Когда необходимо, чтобы объект мог переходить между состояниями, не завися от деталей реализации каждого состояния.
// 3. Когда код содержит множество условных операторов, которые проверяют текущее состояние объекта, и нужно избежать их чрезмерного использования.
// Примеры:
// 1. Онлайн-заказы: Заказ может находиться в состоянии “Новый”, “Обрабатывается”, “Отправлен”, “Доставлен”, и действия зависят от статуса.
// 2. Аутентификация пользователя: Пользователь может находиться в состоянии “Неаутентифицирован”, “Аутентифицирован”, “Заблокирован”.

package main

import "fmt"

type State interface {
	Publish()
}

type DraftState struct {
	doc *Document
}

func (d *DraftState) Publish() {
	fmt.Println("Document is being sent for moderation.")
	d.doc.SetState(&ModerationState{doc: d.doc})
}

type ModerationState struct {
	doc *Document
}

func (m *ModerationState) Publish() {
	fmt.Println("Document is approved and published.")
	m.doc.SetState(&PublishedState{doc: m.doc})
}

type PublishedState struct {
	doc *Document
}

func (p *PublishedState) Publish() {
	fmt.Println("Document is already published. No action taken.")
}

type Document struct {
	currentState State
}

func (d *Document) SetState(state State) {
	d.currentState = state
}

func (d *Document) Publish() {
	d.currentState.Publish()
}

func main() {
	document := &Document{}
	document.SetState(&DraftState{doc: document})

	document.Publish()
	document.Publish()
	document.Publish()
}
