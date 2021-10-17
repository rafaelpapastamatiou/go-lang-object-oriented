package structs

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name            string `json:"name"`
	Year            int    `json:"year"`
	Color           string `json:"color"`
	colorUnexported string `json:"-"`
}

func (c Car) info() string {
	return fmt.Sprintf("\nCar: %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}

type SuperCar struct {
	Car
	CanFly bool `json:"canFly"`
}

func (c SuperCar) info() string {
	return fmt.Sprintf("\nCar: %s\nYear: %d\nColor: %s\nCan fly? %v", c.Name, c.Year, c.Color, c.CanFly)
}

func Describe() {
	fmt.Printf("Structs\n")
	corolla := Car{
		"Corolla",
		2018,
		"Grey",
		"Grey",
	}

	bmw := Car{
		"BMW",
		2022,
		"Black",
		"Black",
	}

	fmt.Printf("\n%s => %v", corolla.Name, corolla)
	fmt.Printf("\n%s => %v\n", bmw.Name, bmw)

	fmt.Println(Car.info(corolla))
	fmt.Println(Car.info(bmw))

	bugatti := SuperCar{
		Car: Car{
			"Bugatti Chiron",
			2020,
			"Black",
			"Black",
		},
		CanFly: true,
	}

	fmt.Println(SuperCar.info(bugatti))

	bugattiJson, _ := json.Marshal(bugatti)

	fmt.Printf("\nBugatti Struct => JSON (ASCII) => %v", bugattiJson)
	fmt.Printf("\n\nBugatti Struct => JSON (parsed to string) => %v", string(bugattiJson))

	var ferrari SuperCar
	ferrariJSON := []byte(`{"name":"Ferrari Portofino", "year": 2018, "color":"red", "canFly":true}`)

	json.Unmarshal(ferrariJSON, &ferrari)

	fmt.Printf("\n\nFerrari JSON => Struct => %v", ferrari)
}
