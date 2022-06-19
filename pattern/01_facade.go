package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	The facade pattern (also spelled façade) is a software-design pattern commonly used in object-oriented programming.
	Analogous to a facade in architecture, a facade is an object that serves as a front-facing interface masking more complex underlying or structural code.
	A facade can:

		improve the readability and usability of a software library by masking interaction with more complex components behind a single (and often simplified) API
		provide a context-specific interface to more generic functionality (complete with context-specific input validation)
		serve as a launching point for a broader refactor of monolithic or tightly-coupled systems in favor of more loosely-coupled code


	Developers often use the facade design pattern when a system is very complex or difficult to understand
	because the system has many interdependent classes or because its source code is unavailable.
	This pattern hides the complexities of the larger system and provides a simpler interface to the client.
	It typically involves a single wrapper class that contains a set of members required by the client.
	These members access the system on behalf of the facade client and hide the implementation details.

*/
type Order struct {
	id      string
	data    string
	adm     *Administration
	kitchen *Kitchen
	courier *Courier
}

type Administration struct {
	name string
	id   string
}

func newAdministration(name, id string) *Administration {
	return &Administration{name, id}
}

func (a *Administration) recieveOrder(id, data string) *Order {
	fmt.Printf("Order %s accepted.\nOrder details: %s\n\n", id, data)
	return &Order{id: id, data: data}
}

type Courier struct {
	name string
	id   string
}

func newCourier(name, id string) *Courier {
	return &Courier{name, id}
}

func (c *Courier) deliverOrder(id string) {
	fmt.Printf("Order %s delivered\n", id)
}

type Kitchen struct {
	name string
	id   string
}

func newKitchen(name, id string) *Kitchen {
	return &Kitchen{name, id}
}

func (k *Kitchen) makePizza(id string) {
	fmt.Printf("The order %s is ready.\n\n", id)
}

func newOrder() {

}

func makeOrder(id, data string) {

	adm := newAdministration("adm", "adm1")
	kitchen := newKitchen("kitchen", "kitchen1")
	courier := newCourier("courier", "courier1")
	adm.recieveOrder(id, data)
	kitchen.makePizza(id)
	courier.deliverOrder(id)
}

// func main() {
// 	makeOrder("1", "Neapolitan pizza")
// }
