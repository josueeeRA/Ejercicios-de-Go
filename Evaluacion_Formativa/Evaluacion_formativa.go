package main

import (
	"fmt"
)

type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate int
	shelfLife       int
}

func NewMedicine(name string, manufacturer string, manufactureDate int, shelfLife int) Medicine {
	medicina := Medicine{}
	medicina.SetName(name)
	medicina.SetManufacturer(manufacturer)
	medicina.SetManufactureDate(manufactureDate)
	medicina.SetShelfLife(shelfLife)
	return medicina
}

func (medicina Medicine) GetName() string {
	return medicina.name
}

func (medicina *Medicine) SetName(name string) {
	if name != "" {
		medicina.name = name
	}
}

func (medicina Medicine) GetManufacturer() string {
	return medicina.manufacturer
}

func (medicina *Medicine) SetManufacturer(manufacturer string) {
	if manufacturer != "" {
		medicina.manufacturer = manufacturer
	}
}

func (medicina Medicine) GetManufactureDate() int {
	return medicina.manufactureDate
}

func (medicina *Medicine) SetManufactureDate(manufactureDate int) {
	if manufactureDate > 0 {
		medicina.manufactureDate = manufactureDate
	}
}

func (medicina Medicine) GetShelfLife() int {
	return medicina.shelfLife
}

func (medicina *Medicine) SetShelfLife(shelfLife int) {
	if shelfLife > 0 {
		medicina.shelfLife = shelfLife
	}
}

type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}

func NewTablet(name string, manufacturer string, manufactureDate int, shelfLife int, dosePerTablet float64, isPrescription bool) Tablet {
	tableta := Tablet{
		Medicine:       NewMedicine(name, manufacturer, manufactureDate, shelfLife),
		isPrescription: isPrescription,
	}
	if dosePerTablet > 0 {
		tableta.dosePerTablet = dosePerTablet
	}
	return tableta
}