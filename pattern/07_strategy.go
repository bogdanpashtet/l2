package pattern

import "fmt"

const (
	A = "Home"
	B = "Airport"
)

// Strategy - интерфейс, который наследуют все стратегии, содержащий метод постройки маршрута
type Strategy interface {
	buildRoute(a, b string)
}

// Bus - конкретная стратегия
type Bus struct {
	time uint16
}

func (bus *Bus) buildRoute(a, b string) {
	fmt.Printf("Время потраченное на перемещение из %s в %s на автобусе заняло %v минут\n", a, b, bus.time)
}

// Walking - конкретная стратегия
type Walking struct {
	time uint16
}

func (wlk *Walking) buildRoute(a, b string) {
	fmt.Printf("Время потраченное на перемещение из %s в %s пешком заняло %v минут\n", a, b, wlk.time)
}

// Navigator - Context
type Navigator struct {
	wayToGetThere Strategy // указатель на конкретную стратегию (способ перемещения)
}

func initNavigaror(str Strategy) *Navigator {
	return &Navigator{
		wayToGetThere: str,
	}
}

func (nav *Navigator) setStrategy(str Strategy) {
	nav.wayToGetThere = str
}

func StrategyFunc() {
	bus := Bus{3}
	walking := Walking{10}
	navigator := initNavigaror(&bus)
	navigator.wayToGetThere.buildRoute(A, B)

	navigator = initNavigaror(&walking)
	navigator.wayToGetThere.buildRoute(A, B)
}
