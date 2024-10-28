// Плюсы:
// 1. Можно заменять алгоритмы на лету.
// 2. Изолирует код и данные алгоритмов от остальных классов.
// 3. Легко добавлять новые стратегии без изменения существующего кода.
// Минусы:
// 1. Усложняет программу за счёт дополнительных объектов.
// 2. Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
// Применимость:
// 1. Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
// 2. Когда существует несколько схожих алгоритмов, которые отличаются небольшими деталями, и нужно избежать множества условных операторов в коде.
// 3. Когда нужно изменять поведение объекта на лету, подставляя один из нескольких возможных алгоритмов.
// Примеры:
// 1. Mожно использовать разные алгоритмы сортировки в зависимости от типа данных и их размера.
// 2. Оплата в интернет-магазине: использование разных стратегий оплаты.

package main

import "fmt"

type RouteStrategy interface {
	buildRoute(A, B string)
}

type RoadStrategy struct{}

func (r *RoadStrategy) buildRoute(A, B string) {
	fmt.Printf("The road route from %s to %s has been built\n", A, B)
}

type WalkingStrategy struct{}

func (w *WalkingStrategy) buildRoute(A, B string) {
	fmt.Printf("The walking route from %s to %s has been built\n", A, B)
}

type PublicTransportStrategy struct{}

func (pt *PublicTransportStrategy) buildRoute(A, B string) {
	fmt.Printf("The public transport route from %s to %s has been built\n", A, B)
}

type Navigator struct {
	strategy RouteStrategy
}

func (n *Navigator) SetStrategy(strategy RouteStrategy) {
	n.strategy = strategy
}

func (n *Navigator) BuildRoute(A, B string) {
	if n.strategy == nil {
		fmt.Println("No strategy set. Please set a route strategy.")
		return
	}
	n.strategy.buildRoute(A, B)
}

func main() {
	navigator := Navigator{}

	navigator.SetStrategy(&RoadStrategy{})
	navigator.BuildRoute("A", "B")

	navigator.SetStrategy(&WalkingStrategy{})
	navigator.BuildRoute("A", "B")

	navigator.SetStrategy(&PublicTransportStrategy{})
	navigator.BuildRoute("A", "B")
}
