package interfaces

import "fmt"

type vehicle interface {
	start() string
}

type car struct {
	name string
}

type motorcycle struct {
	name string
}

func (c car) start() string {
	return "The car " + c.name + " has been started"
}

func (m motorcycle) start() string {
	return "The motorcycle " + m.name + " has been started"
}

func startVehicle(v vehicle) string {
	return v.start()
}

func Describe() {
	interfaces()
	fmt.Printf("\n\n")
	emptyInterfaces()
}

func interfaces() {
	fmt.Printf("Interfaces\n\n")

	porsche := car{"Porsche"}
	bmw := motorcycle{"BMW"}

	fmt.Println(startVehicle(porsche))
	fmt.Println(startVehicle(bmw))
}

type Names []interface{}

func emptyInterfaces() {
	fmt.Printf("Empty Interfaces\n\n")

	var names Names

	names.load()
	names.printNames()
}

func (n *Names) load() {
	*n = Names{
		"Rafael",
		"Papastamatiou",
		"Maia",
	}
}

func (n Names) printNames() {
	fmt.Println(n)
}
