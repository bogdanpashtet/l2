package pattern

import "fmt"

const (
	Toyota = "Toyota"
	Subaru = "Subaru"
)

const (
	Black = "Black"
	White = "White"
)

const Roadster = "Roadster"

// Car - та структура, которую строим
type Car struct {
	brand    string
	color    string
	bodyType string
}

// Car.print - вывод полей в форматированном виде
func (c Car) print() {
	fmt.Printf("--------------\nBrand: %s\nColor: %s\nBody type: %s\n--------------\n",
		c.brand, c.color, c.bodyType)
}

// Assembly - интерфейс строитель
type Assembly interface {
	SetBrand()
	SetColor()
	SetBodyType()
	GetCar() *Car
}

type assembler struct {
	assembly Assembly
}

func StartAssembling(brand string) *Car {
	switch brand {
	case Toyota:
		assembler1 := assembler{&ToyotaBuilder{new(Car)}}
		return assembler1.Construct()
	case Subaru:
		assembler1 := assembler{&SubaruBuilder{new(Car)}}
		return assembler1.Construct()
	default:
		return nil
	}
}

// Construct -собираем необходимую машину
func (d *assembler) Construct() *Car {
	d.assembly.SetBrand()
	d.assembly.SetColor()
	d.assembly.SetBodyType()
	return d.assembly.GetCar()
}

//-------------- TOYOTA --------------

type ToyotaBuilder struct {
	car *Car
}

func (tb *ToyotaBuilder) SetBrand() {
	tb.car.brand = Toyota
}

func (tb *ToyotaBuilder) SetColor() {
	tb.car.color = Black
}

func (tb *ToyotaBuilder) SetBodyType() {
	tb.car.bodyType = Roadster
}

func (tb *ToyotaBuilder) GetCar() *Car {
	return tb.car
}

//-------------- SUBARU --------------

type SubaruBuilder struct {
	car *Car
}

func (tb *SubaruBuilder) SetBrand() {
	tb.car.brand = Subaru
}

func (tb *SubaruBuilder) SetColor() {
	tb.car.color = White
}

func (tb *SubaruBuilder) SetBodyType() {
	tb.car.bodyType = Roadster
}

func (tb *SubaruBuilder) GetCar() *Car {
	return tb.car
}

func Builder() {
	tt := StartAssembling(Toyota)
	tt.print()

	tt = StartAssembling(Subaru)
	tt.print()
}
