package dbase

import (
	"bytes"
	"fmt"
)

// Person represents an example for a table layout.
type Person struct {
	Name string
	Age  int
}

// NewPerson creates a new Person instance.
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// PersonLayout returns the layout for Person.
func PersonLayout() []Layout {
	return GetLayout(&Person{})
}

// Get returns Person data.
func (p *Person) Get() interface{} {
	return &Person{
		Name: p.Name,
		Age:  p.Age,
	}
}

// ToString returns Person data in string format.
func (p *Person) ToString() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s, %d", p.Name, p.Age))
	return buf.String()
}
