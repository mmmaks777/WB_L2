// Плюсы:
// 1. Выделяет код производства продуктов в одно место, упрощая поддержку кода.
// 2. Упрощает добавление новых продуктов в программу.
// Минусы:
// 1. Для каждого нового продукта нужно создавать подклассы или реализацию фабричного метода, что может привести к усложнению системы.
// Применимость:
// 1. Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
// 2. Когда нужно давать возможность подклассам выбрать тип создаваемого объекта.
// Примеры:
// 1. СУБД используют фабричный метод для создания соединений с различными базами данных (например, MySQL, PostgreSQL).
// 2. GUI-библиотеки используют фабричный метод для создания интерфейсных элементов, таких как кнопки и окна, адаптированные для разных операционных систем.

package main

import "fmt"

type Transport interface {
	Deliver()
}

type Car struct{}

func (c *Car) Deliver() {
	fmt.Println("Deliver by car")
}

type Ship struct{}

func (s *Ship) Deliver() {
	fmt.Println("Deliver by ship")
}

type TransportFactory interface {
	CreateTransport() Transport
}

type CarFactory struct{}

func (c *CarFactory) CreateTransport() Transport {
	return &Car{}
}

type ShipFactory struct{}

func (s *ShipFactory) CreateTransport() Transport {
	return &Ship{}
}

func main() {
	carFactory := &CarFactory{}
	transport := carFactory.CreateTransport()
	transport.Deliver()

	shipFactory := &ShipFactory{}
	transport = shipFactory.CreateTransport()
	transport.Deliver()
}
