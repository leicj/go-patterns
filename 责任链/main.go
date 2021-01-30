package main

import link "p13/link"

func main() {
	cashier := &link.Cashier{}

	medical := &link.Medical{}
	medical.SetNext(cashier)

	doctor := &link.Doctor{}
	doctor.SetNext(medical)

	reception := &link.Reception{}
	reception.SetNext(doctor)

	patient := &link.Patient{Name: "abc"}
	reception.Execute(patient)
}