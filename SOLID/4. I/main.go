package main

type Room struct {
}

type VacuumCleaner interface {
	Vacuum(r Room)
	Mop(r Room)
	Polish(r Room)
}

// It is ok for multifuncional device
// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
type MultiFunctionVacuumCleaner struct {
}

func (m MultiFunctionVacuumCleaner) Vacuum(r Room) {
}

func (m MultiFunctionVacuumCleaner) Mop(r Room) {
}

func (m MultiFunctionVacuumCleaner) Polish(r Room) {
}

// ...but when we have old vacuum cleaner
// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
type OldVacuumCleaner struct {
}

func (o OldVacuumCleaner) Vacuum(r Room) {
	// OK
}

func (o OldVacuumCleaner) Mop(r Room) {
	panic("Operation not supported")
}

func (o OldVacuumCleaner) Polish(r Room) {
	panic("Operation not supported")
}

// Better approach -> split into several interfaces
// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
type Vaccuming interface {
	Vacuum(r Room)
}

type Mopping interface {
	Mop(r Room)
}

type Polishing interface {
	Polish(r Room)
}

// Vacuum cleaner only
type MyVacuumCleaner struct {
}

func (m MyVacuumCleaner) Vaccum(r Room) {
}

// Combine interfaces
type VacuumingAndMoppingDevice struct{}

func (v VacuumingAndMoppingDevice) Vaccum(r Room) {
}

func (v VacuumingAndMoppingDevice) Mop(r Room) {
}

type MultiFunctionDevice interface {
	Vaccuming
	Mopping
}

// Interface combination + decorator
type MultiFunctionMachine struct {
	vacuumCleaner Vaccuming
	mop           Mopping
}

func (m MultiFunctionMachine) Vaccum(r Room) {
	m.vacuumCleaner.Vacuum(r)
}

func (m MultiFunctionMachine) Mop(r Room) {
	m.mop.Mop(r)
}

func main() {

}