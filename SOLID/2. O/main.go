package main

import "fmt"

type Cuisine int
type MealSize int

const (
	polish Cuisine = iota
	italian
	american
)

const (
	small MealSize = iota
	medium
	large
)

type Meal struct {
	name    string
	cuisine Cuisine
	size    MealSize
}

type Filter struct{}

// Basic requirement
// vvvvvvvvvvvvvvvvv
func (f *Filter) filterByCuisine(meals []Meal, cuisine Cuisine) []*Meal {

	result := make([]*Meal, 0)

	for i, v := range meals {
		if v.cuisine == cuisine {
			result = append(result, &meals[i])
		}
	}

	return result
}

// Additional requirements added to Filter type --> WRONG APPROACH!
// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

func (f *Filter) filterBySize(meals []Meal, size MealSize) []*Meal {

	result := make([]*Meal, 0)

	for i, v := range meals {
		if v.size == size {
			result = append(result, &meals[i])
		}
	}

	return result
}

func (f *Filter) filterBySizeAndCuisine(meals []Meal, size MealSize, cuisine Cuisine) []*Meal {

	result := make([]*Meal, 0)

	for i, v := range meals {
		if v.size == size && v.cuisine == cuisine {
			result = append(result, &meals[i])
		}
	}

	return result
}

// Better solution -> Specification pattern
// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

type Specification interface {
	IsSatisfied(m *Meal) bool
}

type CuisineSpecification struct {
	cuisine Cuisine
}

func (spec CuisineSpecification) IsSatisfied(meal *Meal) bool {
	return meal.cuisine == spec.cuisine
}

type SizeSpecification struct {
	size MealSize
}

func (spec SizeSpecification) IsSatisfied(meal *Meal) bool {
	return meal.size == spec.size
}

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(meal *Meal) bool {
	return spec.first.IsSatisfied(meal) && spec.second.IsSatisfied(meal)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(meals []Meal, spec Specification) []*Meal {

	result := make([]*Meal, 0)

	for i, v := range meals {
		if spec.IsSatisfied(&v) {
			result = append(result, &meals[i])
		}
	}

	return result
}

func main() {

	Pasta := Meal{"Pasta", italian, medium}
	Pizza := Meal{"Pizza", italian, large}
	Schabowy := Meal{"Schabowy", polish, large}
	Burger := Meal{"Burger", american, small}

	meals := []Meal{Pasta, Pizza, Schabowy, Burger}

	// First solution - with broken open-closed principle
	// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	fmt.Print("Italian meals (old):\n")
	f := Filter{}
	for _, v := range f.filterByCuisine(meals, italian) {
		fmt.Printf(" - %s is Italian meal\n", v.name)
	}

	// Second solution - specification pattern
	// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
	fmt.Print("Italian meals (new):\n")
	italianSpec := CuisineSpecification{italian}
	bf := BetterFilter{}
	for _, v := range bf.Filter(meals, italianSpec) {
		fmt.Printf(" - %s is Italian meal\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	largeItalianSpec := AndSpecification{largeSpec, italianSpec}

	fmt.Print("Large Italian items:\n")
	for _, v := range bf.Filter(meals, largeItalianSpec) {
		fmt.Printf(" - %s is large and Italian meal\n", v.name)
	}
}
