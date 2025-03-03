package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang/

import "fmt"

// department
type department interface {
	execute(*patient)
	setNext(department)
}

// reception
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

// doctor
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

// medical
type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

// cachier
type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) {
	c.next = next
}

// patient
type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// main
// func main() {
// 	cashier := &cashier{}

// 	//Set next for medical department
// 	medical := &medical{}
// 	medical.setNext(cashier)

// 	//Set next for doctor department
// 	doctor := &doctor{}
// 	doctor.setNext(medical)

// 	//Set next for reception department
// 	reception := &reception{}
// 	reception.setNext(doctor)

// 	patient := &patient{name: "abc"}
// 	//Patient visiting
// 	reception.execute(patient)
// }
