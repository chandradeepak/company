package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

//since it is mentioned the company is small we are gong to go with in memory tree stucture and going
// to use map for quick reference.

type employee struct {
	ID        string
	Manager   *employee
	Employees []*employee
	ManagerID string
	Name      string
}

//prints the heirarchy of the employess
func printCompanyHeirarchy(e *employee, level int) {
	//for each level we need that many tabs
	for i := 0; i <= level; i++ {
		fmt.Print("\t")
	}
	//for each name put a new line
	fmt.Println(e.Name)
	l := level + 1
	for _, emp := range e.Employees {
		printCompanyHeirarchy(emp, l)
	}

}

func parseRecord(record string) (*employee, error) {

	data := strings.Split(record, ",")

	if len(data) < 0 || len(data) > 3 {
		return nil, errors.New(("wrong number of delimeters"))
	}

	name := strings.Trim(strings.TrimSpace(data[0]), "\"")
	id := strings.Trim(strings.TrimSpace(data[1]), "\"")
	managerID := strings.Trim(strings.TrimSpace(data[2]), "\"")

	if id == "" {
		return nil, errors.New("employee ID can not be empty exiting now")
	}

	//make sure your are trimmimg all spaces when putting in map
	e := &employee{Name: name, ID: id, ManagerID: managerID}
	return e, nil

}

func loadData(fileName string) (map[string]*employee, error) {
	employeeMap := make(map[string]*employee)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//read one line at a time
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//read comma delimited data
		emp, err := parseRecord(scanner.Text())
		if err != nil {
			return nil, err
		}
		//update employee map with id and the employee
		de, ok := employeeMap[emp.ID]
		if ok {
			return nil, fmt.Errorf("duplicate employee found Name:%s ID:%s ", de.Name, de.ID)
		}
		employeeMap[emp.ID] = emp

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return employeeMap, nil
}

func main() {
	dataFile := flag.String("data", "data.csv", "file name")

	employeeMap, err := loadData(*dataFile)
	if err != nil {
		log.Fatal(err)
	}

	var ceo *employee

	// Link the employees by iterating through the map and assign each employees manager and his employees
	// this linking can be further optimized when reading the data itself but for code readbility purpose we are
	// iterating in seperately
	for _, emp := range employeeMap {

		//check if employee's manager exists. If he exists create the employee - manager relation ship
		// and add this employee as sub ordinate of the manager.
		m, ok := employeeMap[emp.ManagerID]
		if ok {
			emp.Manager = m
			m.Employees = append(m.Employees, emp)
			employeeMap[emp.ID] = emp
			employeeMap[m.ID] = m
		} else {
			// if we have got only one employee who doesn't have a manager. we are assuming data is given like that
			if ceo != nil {
				log.Fatal("we can not have more than one employee with out manager stopping now")
			}
			log.Println("can not find manager for employee, ceo is:", emp.Name, emp.ID, emp.ManagerID)
			ceo = emp
		}

	}

	//print the heirarchy of employees
	printCompanyHeirarchy(ceo, 0)

}
