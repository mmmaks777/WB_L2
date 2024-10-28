// Плюсы:
// 1. Позволяет создавать продукты пошагово.
// 2. Позволяет использовать один и тот же код для создания различных продуктов.
// 3. Изолирует сложный код сборки продукта от его основной бизнес-логики.
// Минусы:
// 1. Усложняет код программы из-за введения дополнительных строителей.
// 2. Клиент будет привязан к конкретным строителям, так как в интерфейсе директора может не быть метода получения результата.
// Применимость:
// 1. Когда требуется создавать разные варианты объекта, используя один и тот же код строительства.
// 2. Когда нужно обеспечивать пошаговое создание продукта, позволяя прерывать процесс в любой момент.
// Примеры:
// 1. Создание сложных объектов в GUI, построение сложных интерфейсов с разными компонентами.
// 2. Конфигурация серверов или приложений с множеством параметров.

package main

import "fmt"

type House struct {
	Walls   string
	Doors   string
	Windows string
	Roof    string
}

func (h *House) String() string {
	return fmt.Sprintln(h.Walls, h.Doors, h.Windows, h.Roof)
}

type HouseBuilder interface {
	BuildWalls()
	BuildDoors()
	BuildWindows()
	BuildRoof()
	GetHouse() *House
}

type WoodenHouseBuilder struct {
	house *House
}

func (w *WoodenHouseBuilder) BuildWalls() {
	w.house.Walls = "Wooden Walls"
}

func (w *WoodenHouseBuilder) BuildDoors() {
	w.house.Doors = "Wooden Doors"
}

func (w *WoodenHouseBuilder) BuildWindows() {
	w.house.Windows = "Glass Windows"
}

func (w *WoodenHouseBuilder) BuildRoof() {
	w.house.Roof = "Wooden Roof"
}

func (w *WoodenHouseBuilder) GetHouse() *House {
	return w.house
}

func NewWoodenHouseBuilder() *WoodenHouseBuilder {
	return &WoodenHouseBuilder{house: &House{}}
}

type StoneHouseBuilder struct {
	house *House
}

func (w *StoneHouseBuilder) BuildWalls() {
	w.house.Walls = "Stone Walls"
}

func (w *StoneHouseBuilder) BuildDoors() {
	w.house.Doors = "Stone Doors"
}

func (w *StoneHouseBuilder) BuildWindows() {
	w.house.Windows = "Glass Windows"
}

func (w *StoneHouseBuilder) BuildRoof() {
	w.house.Roof = "Stone Roof"
}

func (w *StoneHouseBuilder) GetHouse() *House {
	return w.house
}

func NewStoneHouseBuilder() *StoneHouseBuilder {
	return &StoneHouseBuilder{house: &House{}}
}

type Director struct {
	builder HouseBuilder
}

func (d *Director) SetBuilder(b HouseBuilder) {
	d.builder = b
}

func (d *Director) ConstructHouse() {
	d.builder.BuildWalls()
	d.builder.BuildDoors()
	d.builder.BuildWindows()
	d.builder.BuildRoof()
}

func main() {
	director := &Director{}

	director.SetBuilder(NewWoodenHouseBuilder())
	director.ConstructHouse()
	fmt.Println(director.builder.GetHouse())

	director.SetBuilder(NewStoneHouseBuilder())
	director.ConstructHouse()
	fmt.Println(director.builder.GetHouse())
}
