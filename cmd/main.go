package main

import "fmt"

type Agent struct {
	name   string
	sex    string
	age    int
	salary int
}

func newAgent(name, sex string, age, salary int) Agent {
	return Agent{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

func (a Agent) getInfo() string {
	return fmt.Sprintf("Agent: %s\nAge: %d\nSalary: %d\n", a.name, a.age, a.salary)
}

func main() {
	agent1 := newAgent("Rick", "M", 21, 100)

	fmt.Printf("%s\n", agent1.getInfo())
}
