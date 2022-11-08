package main

import "fmt"

type Transport interface { //定义一个运输工具接口
	deliver()
}
type Car struct { //定义汽车运输工具结构体
}

type Ship struct { //定义船运输工具结构体
}

func (c Car) deliver() {
	fmt.Println("在盒子中以陆路运输")
}

func (s Ship) deliver() {
	fmt.Println("在集装箱中以海路运输")
}

type transport_Factory struct { //定义交通工具创建工厂
}

func (tf transport_Factory) Create(types string) Transport { //工厂创建交通工具实例
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
