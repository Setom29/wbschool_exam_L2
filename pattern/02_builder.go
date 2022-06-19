package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	The builder pattern is a design pattern designed to provide a flexible solution to various object creation problems in object-oriented programming.
	The intent of the Builder design pattern is to separate the construction of a complex object from its representation.
	It is one of the Gang of Four design patterns.

	The Builder design pattern solves problems like:

		How can a class (the same construction process) create different representations of a complex object?
		How can a class that includes creating a complex object be simplified?
		Creating and assembling the parts of a complex object directly within a class is inflexible. It commits the class to creating a particular representation of the complex object and makes it impossible to change the representation later independently from (without having to change) the class.

	The Builder design pattern describes how to solve such problems:

		Encapsulate creating and assembling the parts of a complex object in a separate Builder object.
		A class delegates object creation to a Builder object instead of creating the objects directly.
		A class (the same construction process) can delegate to different Builder objects to create different representations of a complex object.

	Definition
	The intent of the Builder design pattern is to separate the construction of a complex object from its representation. By doing so, the same construction process can create different representations.[1]

	Advantages of the Builder pattern include:

		Allows you to vary a product's internal representation.
		Encapsulates code for construction and representation.
		Provides control over steps of construction process.

	Disadvantages of the Builder pattern include:

		A distinct ConcreteBuilder must be created for each type of product.
		Builder classes must be mutable.
		May hamper/complicate dependency injection.
*/

// Builder provides a builder interface.
type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Director implements a manager
type Director struct {
	builder Builder
}

// Construct tells the builder what to do and in what order.
func (d *Director) Construct() {
	d.builder.MakeHeader("Header")
	d.builder.MakeBody("Body")
	d.builder.MakeFooter("Footer")
}

// ConcreteBuilder implements Builder interface.
type ConcreteBuilder struct {
	product *Product
}

// MakeHeader builds a header of document..
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.product.Content += "<header>" + str + "</header>"
}

// MakeBody builds a body of document.
func (b *ConcreteBuilder) MakeBody(str string) {
	b.product.Content += "<article>" + str + "</article>"
}

// MakeFooter builds a footer of document.
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.product.Content += "<footer>" + str + "</footer>"
}

// Product implementation.
type Product struct {
	Content string
}

// Show returns product.
func (p *Product) Show() string {
	return p.Content
}
