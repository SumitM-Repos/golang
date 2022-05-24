package main

import "fmt"

type Electricity_Slabs struct {
	Slab_1 *Slab
	Slab_2 *Slab
	Slab_3 *Slab
	Slab_4 *Slab
}

type Slab struct {
	Min_Unit       int
	Max_Unit       int
	Price_per_unit int
}

var E_S Electricity_Slabs

func Create_Slabs() {

	Slab_1 := Slab{1, 100, 1}
	Slab_2 := Slab{101, 200, 2}
	Slab_3 := Slab{201, 300, 3}
	Slab_4 := Slab{301, 400, 5}

	E_S = Electricity_Slabs{Slab_1: &Slab_1, Slab_2: &Slab_2, Slab_3: &Slab_3, Slab_4: &Slab_4}
}

func Calculate_Bill(e_s Electricity_Slabs, units_used int) int {

	bill_amount := 0
	final_bill := 0

	if units_used < 0 {
		fmt.Println("Units used is not valid number...")
		return -1
	}

	for units_used > 0 {

		billed_units := 0

		switch {
		case units_used <= e_s.Slab_1.Max_Unit:
			billed_units = units_used
			bill_amount = billed_units * e_s.Slab_1.Price_per_unit

		case units_used <= e_s.Slab_2.Max_Unit && units_used >= e_s.Slab_2.Min_Unit:
			billed_units = units_used - e_s.Slab_1.Max_Unit
			bill_amount = billed_units * e_s.Slab_2.Price_per_unit

		case units_used <= e_s.Slab_3.Max_Unit && units_used >= e_s.Slab_3.Min_Unit:
			billed_units = units_used - e_s.Slab_2.Max_Unit
			bill_amount = billed_units * e_s.Slab_3.Price_per_unit

		case units_used >= e_s.Slab_4.Min_Unit:
			billed_units = units_used - e_s.Slab_3.Max_Unit
			bill_amount = billed_units * e_s.Slab_4.Price_per_unit

		}
		units_used = units_used - billed_units
		final_bill = final_bill + bill_amount

	}
	return final_bill

}

func init() {
	Create_Slabs()
}

func main() {

	fmt.Println(Calculate_Bill(E_S, 50))   // 50 * 1 = 50
	fmt.Println(Calculate_Bill(E_S, 150))  // 150 * 2 = 300
	fmt.Println(Calculate_Bill(E_S, 250))  // 250 * 3 = 750
	fmt.Println(Calculate_Bill(E_S, 350))  // 350 * 5 = 1750
	fmt.Println(Calculate_Bill(E_S, -222)) // Units used is not valid number... -1

}
