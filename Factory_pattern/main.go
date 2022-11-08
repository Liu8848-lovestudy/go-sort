package main

import "fmt"

type Transport interface {
	deliver()
}
type Car struct {
}

type Ship struct {
}

func (c Car) deliver() {
	fmt.Println("在盒子中以陆路运输")
}

func (s Ship) deliver() {
	fmt.Println("在集装箱中以海路运输")
}

type transport_Factory struct {
}

func (tf transport_Factory) Create(types string) Transport {
	if types == "car" {
		return Car{}
	} else if types == "ship" {
		return Ship{}
	} else {
		return nil
	}
}

func main() {

	tf := transport_Factory{}
	car := tf.Create("car")
	ship := tf.Create("ship")
	car.deliver()
	ship.deliver()

}
