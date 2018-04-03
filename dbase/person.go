package dbase

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
	//return []Layout{
	//    {
	//        ColName: "name",
	//        ColType: "string",
	//    },
	//    {
	//        ColName: "age",
	//        ColType: "int",
	//    },
	//}
	return GetLayout(&Person{})
}

// Get return Person data.
func (p *Person) Get() interface{} {
	return &Person{
		Name: p.Name,
		Age:  p.Age,
	}
}
