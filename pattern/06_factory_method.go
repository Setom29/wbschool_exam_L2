package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// https://golangbyexample.com/golang-factory-design-pattern/

import "fmt"

// iGun
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// gun
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

// ak47
type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// maverick
type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}

// gunFactory
func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "maverick" {
		return newMaverick(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

// main
// func main() {
// 	ak47, _ := getGun("ak47")
// 	maverick, _ := getGun("maverick")
// 	printDetails(ak47)
// 	printDetails(maverick)
// }

// func printDetails(g iGun) {
// 	fmt.Printf("Gun: %s", g.getName())
// 	fmt.Println()
// 	fmt.Printf("Power: %d", g.getPower())
// 	fmt.Println()
// }
