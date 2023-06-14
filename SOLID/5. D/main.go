package main

import "fmt"

type Relationship int

const (
	Teacher Relationship = iota
	Friend
	Student
)

type Person struct {
	name string
}

type Info struct {
	from *Person
	relationship Relationship
	to *Person
}

type Relationships struct {
	relations []Info
}

type RelationshipBrowser interface {
	FindAllStudentsOf(name string) []*Person
}

func (rs *Relationships) FindAllStudentsOf(name string) []*Person {

	result := make([]*Person, 0)

	for i, v := range rs.relations {
		if v.relationship == Teacher &&
			v.from.name == name {
			result = append(result, rs.relations[i].to)
		}
	}

	return result
}

func (rs *Relationships) AddTeacherAndStudent(teacher, student *Person) {

	rs.relations = append(rs.relations, Info{teacher, Teacher, student})
	rs.relations = append(rs.relations, Info{student, Student, teacher})
}

type Research struct {
	// WRONG -> high-level module not depend on low-level module
	// vvvvvvvvvvvvvvvvvvvvvvvvvvv
	// relationships Relationships

	// Better solution - use abstraction (interface)
	// vvvvvvvvvvvvvvvvvvvvvvvv
	browser RelationshipBrowser
}

func (r *Research) Check() {

	// Wrong way to use low-level module
	// vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv

	//relations := r.relationships.relations
	//for _, rel := range relations {
	//	if rel.from.name == "Gomez" &&
	//		rel.relationship == Teacher {
	//		fmt.Println("Gomez has a student called", rel.to.name)
	//	}
	//}

	for _, s := range r.browser.FindAllStudentsOf("Gomez") {
		fmt.Println("Gomez has a student called", s.name)
	}
}

func main() {
	teacher := Person{ "Gomez" }
	student1 := Person{ "Matt" }
	student2 := Person{ "Steve" }

	relationships := Relationships{}
	relationships.AddTeacherAndStudent(&teacher, &student1)
	relationships.AddTeacherAndStudent(&teacher, &student2)

	research := Research{&relationships}
	research.Check()
}