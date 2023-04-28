package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee // 型のみを書くことで埋め込み
	Reports  []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	newEmployees := []Employee{
		{
			"Bob",
			"123",
		},
		{
			"Alice",
			"234",
		},
	}
	return newEmployees
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "Todd",
			ID:   "100",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	m.Reports = m.FindNewEmployees()
	fmt.Println(m.Employee)
	fmt.Println(m.Reports)
}
