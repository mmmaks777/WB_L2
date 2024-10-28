// Плюсы:
// 1. Уменьшает зависимость между клиентом и обработчиками.
// 2. Можно динамически изменять порядок обработки запросов.
// 3. Легко можно добавить новый обработчик в цепочку без изменений в существующем коде.
// Минусы:
// 1. Запрос может остаться никем не обработанным.
// Применимость:
// 1. Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
// 2. Когда набор объектов, способных обработать запрос, должен задаваться динамически.
// 3. Когда программа должна обрабатывать разнообразные запросы несколькими способами,
// но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
// Примеры:
// 1. В графических интерфейсах события, такие как клики мыши или нажатия клавиш, могут обрабатываться элементами интерфейса в цепочке.
// 2. Разные уровни логирования (ошибки, предупреждения, информация) могут быть обработаны разными объектами.

package main

import "fmt"

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

type BasicSupport struct {
	BaseHandler
}

func (h *BasicSupport) Handle(request string) {
	if request == "basic" {
		fmt.Println("Basic Support: Handling request.")
	} else {
		fmt.Println("Basic Support: Passing request to the next handler.")
		h.BaseHandler.Handle(request)
	}
}

type TechSupport struct {
	BaseHandler
}

func (h *TechSupport) Handle(request string) {
	if request == "technical" {
		fmt.Println("Tech Support: Handling request.")
	} else {
		fmt.Println("Tech Support: Passing request to the next handler.")
		h.BaseHandler.Handle(request)
	}
}

type SeniorManager struct {
	BaseHandler
}

func (h *SeniorManager) Handle(request string) {
	if request == "manager" {
		fmt.Println("Senior Manager: Handling request.")
	} else {
		fmt.Println("Senior Manager: Request can't be handled.")
	}
}

func main() {
	basicSupport := &BasicSupport{}
	techSupport := &TechSupport{}
	managerSupport := &SeniorManager{}

	basicSupport.SetNext(techSupport).SetNext(managerSupport)

	fmt.Println("Sending 'basic' request:")
	basicSupport.Handle("basic")

	fmt.Println("\nSending 'technical' request:")
	basicSupport.Handle("technical")

	fmt.Println("\nSending 'manager' request:")
	basicSupport.Handle("manager")

	fmt.Println("\nSending 'unknown' request:")
	basicSupport.Handle("unknown")
}
