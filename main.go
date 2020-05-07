package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type employee struct {
	ID        string
	Manager   *employee
	Employees []*employee
	ManagerID string
	Name      string
}

func printHeirarchy(e *employee, level int) {
	for i := 0; i <= level; i++ {
		fmt.Print("\t")
	}
	fmt.Println(e.Name)
	l := level + 1
	for _, emp := range e.Employees {
		printHeirarchy(emp, l)
	}

}

func main() {

	employeeMap := make(map[string]*employee)
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		e := &employee{Name: strings.TrimSpace(data[0]), ID: strings.TrimSpace(data[1]), ManagerID: strings.TrimSpace(data[2])}
		employeeMap[strings.TrimSpace(data[1])] = e
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var ceo *employee
	for _, emp := range employeeMap {

		m, ok := employeeMap[emp.ManagerID]
		if ok {
			emp.Manager = m
			m.Employees = append(m.Employees, emp)
			employeeMap[emp.ID] = emp
			employeeMap[m.ID] = m
		} else {
			log.Println("can not find manager for employee", emp.Name, emp.ID, emp.ManagerID)
			ceo = emp
		}

	}

	printHeirarchy(ceo, 0)

}
