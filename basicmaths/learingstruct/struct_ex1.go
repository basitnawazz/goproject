package learingstruct

type Car struct {
	Name  string
	Age   int
	Fuel  string
	Modal string
}

func NewCar(name, fuel, modal string, age int) Car {
	car := Car{Name: name, Age: age, Modal: modal, Fuel: fuel}
	return car

}
